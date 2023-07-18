package db

import (
	"context"
	consensus "github.com/anyproto/any-sync-consensusnode"
	"github.com/anyproto/any-sync-consensusnode/config"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/consensus/consensusproto/consensuserr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var ctx = context.Background()

func TestService_AddLog(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		log := consensus.Log{
			Id: "logOne",
			Records: []consensus.Record{
				{
					Id:      "recordOne",
					PrevId:  "",
					Payload: []byte("payload"),
					Created: time.Now().Truncate(time.Second).UTC(),
				},
			},
		}
		require.NoError(t, fx.AddLog(ctx, log))
		fetched, err := fx.FetchLog(ctx, log.Id)
		require.NoError(t, err)
		assert.Equal(t, log, fetched)
	})
	t.Run("duplicate error", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		log := consensus.Log{
			Id: "logOne",
		}
		require.NoError(t, fx.AddLog(ctx, log))
		// TODO: check for specified error
		require.Error(t, fx.AddLog(ctx, log))
	})
}

func TestService_AddRecord(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		var records = []consensus.Record{
			{
				Id:     "2",
				PrevId: "1",
			},
			{
				Id:     "3",
				PrevId: "2",
			},
			{
				Id:     "4",
				PrevId: "3",
			},
		}
		l := consensus.Log{
			Id: "logTestRecords",
			Records: []consensus.Record{
				{
					Id: "1",
				},
			},
		}
		require.NoError(t, fx.AddLog(ctx, l))
		for _, rec := range records {
			require.NoError(t, fx.AddRecord(ctx, l.Id, rec))
		}
		fx.assertLogValid(t, l.Id, 4)
	})
	t.Run("conflict", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		log := consensus.Log{
			Id: "logTestRecords",
			Records: []consensus.Record{
				{
					Id: "1",
				},
			},
		}
		require.NoError(t, fx.AddLog(ctx, log))
		assert.EqualError(t, fx.AddRecord(ctx, log.Id, consensus.Record{Id: "2", PrevId: "3"}), consensuserr.ErrConflict.Error())
	})
}

func TestService_FetchLog(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		l, err := fx.FetchLog(ctx, "not exists")
		assert.Empty(t, l)
		assert.ErrorIs(t, err, consensuserr.ErrLogNotFound)
	})
}

func TestService_ChangeReceive(t *testing.T) {
	t.Run("set after run", func(t *testing.T) {
		fx := newFixture(t, nil)
		defer fx.Finish(t)
		assert.Error(t, fx.SetChangeReceiver(func(logId string, records []consensus.Record) {}))
	})
	t.Run("receive changes", func(t *testing.T) {
		var logs = make(chan consensus.Log, 10)
		var count int
		fx := newFixture(t, func(logId string, records []consensus.Record) {
			logs <- consensus.Log{Id: logId, Records: records}
			count++
		})
		defer fx.Finish(t)
		var l = consensus.Log{
			Id: "logTestStream",
			Records: []consensus.Record{
				{
					Id: "1",
				},
			},
		}
		var records = []consensus.Record{
			{
				Id:     "2",
				PrevId: "1",
			},
			{
				Id:     "3",
				PrevId: "2",
			},
			{
				Id:     "4",
				PrevId: "3",
			},
		}
		require.NoError(t, fx.AddLog(ctx, l))
		assert.Empty(t, count)

		for _, rec := range records {
			require.NoError(t, fx.AddRecord(ctx, l.Id, rec))
		}

		timeout := time.After(time.Second)
		for i := 0; i < len(records); i++ {
			select {
			case resLog := <-logs:
				assertLogValid(t, resLog, i+2)
			case <-timeout:
				require.False(t, true)
			}
		}
	})
}

func newFixture(t *testing.T, cr ChangeReceiver) *fixture {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	fx := &fixture{
		Service: New(),
		cancel:  cancel,
		a:       new(app.App),
	}
	fx.a.Register(&config.Config{
		Mongo: config.Mongo{
			Connect:       "mongodb://localhost:27017/?w=majority",
			Database:      "consensus_test",
			LogCollection: "log",
		},
	})
	fx.a.Register(fx.Service)
	require.NoError(t, fx.Service.SetChangeReceiver(cr))
	err := fx.a.Start(ctx)
	if err != nil {
		fx.cancel()
	}
	require.NoError(t, err)
	return fx
}

type fixture struct {
	Service
	a      *app.App
	cancel context.CancelFunc
}

func (fx *fixture) Finish(t *testing.T) {
	if fx.cancel != nil {
		fx.cancel()
	}
	coll := fx.Service.(*service).logColl
	t.Log(coll.Drop(ctx))
	assert.NoError(t, fx.a.Close(ctx))
}

func (fx *fixture) assertLogValid(t *testing.T, logId string, count int) {
	log, err := fx.FetchLog(ctx, logId)
	require.NoError(t, err)
	assertLogValid(t, log, count)
}

func assertLogValid(t *testing.T, log consensus.Log, count int) {
	if count >= 0 {
		assert.Len(t, log.Records, count)
	}
	var prevId string
	for _, rec := range log.Records {
		if len(prevId) != 0 {
			assert.Equal(t, prevId, rec.Id)
		}
		prevId = rec.PrevId
	}
}
