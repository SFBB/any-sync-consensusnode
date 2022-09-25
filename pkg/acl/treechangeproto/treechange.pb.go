// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/acl/treechangeproto/protos/treechange.proto

package treechangeproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RootChange is a root of a tree
type RootChange struct {
	// AclHeadId is a cid of latest acl record at the time of tree creation
	AclHeadId string `protobuf:"bytes,1,opt,name=aclHeadId,proto3" json:"aclHeadId,omitempty"`
	// SpaceId is an id of space where the document is placed
	SpaceId string `protobuf:"bytes,2,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	// ChangeType is a type of tree which this RootChange is a root of
	ChangeType string `protobuf:"bytes,3,opt,name=changeType,proto3" json:"changeType,omitempty"`
	// Timestamp is this change creation timestamp
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Seed is a random bytes to make root change unique
	Seed []byte `protobuf:"bytes,5,opt,name=seed,proto3" json:"seed,omitempty"`
	// Identity is a public key of the tree's creator
	Identity []byte `protobuf:"bytes,6,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (m *RootChange) Reset()         { *m = RootChange{} }
func (m *RootChange) String() string { return proto.CompactTextString(m) }
func (*RootChange) ProtoMessage()    {}
func (*RootChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f177d8514fae978f, []int{0}
}
func (m *RootChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RootChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RootChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RootChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootChange.Merge(m, src)
}
func (m *RootChange) XXX_Size() int {
	return m.Size()
}
func (m *RootChange) XXX_DiscardUnknown() {
	xxx_messageInfo_RootChange.DiscardUnknown(m)
}

var xxx_messageInfo_RootChange proto.InternalMessageInfo

func (m *RootChange) GetAclHeadId() string {
	if m != nil {
		return m.AclHeadId
	}
	return ""
}

func (m *RootChange) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

func (m *RootChange) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *RootChange) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RootChange) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *RootChange) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// TreeChange is a change of a tree
type TreeChange struct {
	// TreeHeadIds are previous ids for this TreeChange
	TreeHeadIds []string `protobuf:"bytes,1,rep,name=treeHeadIds,proto3" json:"treeHeadIds,omitempty"`
	// AclHeadId is a cid of latest acl record at the time of this change
	AclHeadId string `protobuf:"bytes,2,opt,name=aclHeadId,proto3" json:"aclHeadId,omitempty"`
	// SnapshotBaseId is a snapshot (root) of the tree where this change is added
	SnapshotBaseId string `protobuf:"bytes,3,opt,name=snapshotBaseId,proto3" json:"snapshotBaseId,omitempty"`
	// ChangesData is an arbitrary payload to be read by the client
	ChangesData []byte `protobuf:"bytes,4,opt,name=changesData,proto3" json:"changesData,omitempty"`
	// CurrentReadKeyHash is the hash of the read key which is used to encrypt this change
	CurrentReadKeyHash uint64 `protobuf:"varint,5,opt,name=currentReadKeyHash,proto3" json:"currentReadKeyHash,omitempty"`
	// Timestamp is this change creation timestamp
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Identity is a public key with which the raw payload of this change is signed
	Identity []byte `protobuf:"bytes,7,opt,name=identity,proto3" json:"identity,omitempty"`
	// IsSnapshot indicates whether this change contains a snapshot of state
	IsSnapshot bool `protobuf:"varint,8,opt,name=isSnapshot,proto3" json:"isSnapshot,omitempty"`
}

func (m *TreeChange) Reset()         { *m = TreeChange{} }
func (m *TreeChange) String() string { return proto.CompactTextString(m) }
func (*TreeChange) ProtoMessage()    {}
func (*TreeChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f177d8514fae978f, []int{1}
}
func (m *TreeChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TreeChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TreeChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TreeChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeChange.Merge(m, src)
}
func (m *TreeChange) XXX_Size() int {
	return m.Size()
}
func (m *TreeChange) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeChange.DiscardUnknown(m)
}

var xxx_messageInfo_TreeChange proto.InternalMessageInfo

func (m *TreeChange) GetTreeHeadIds() []string {
	if m != nil {
		return m.TreeHeadIds
	}
	return nil
}

func (m *TreeChange) GetAclHeadId() string {
	if m != nil {
		return m.AclHeadId
	}
	return ""
}

func (m *TreeChange) GetSnapshotBaseId() string {
	if m != nil {
		return m.SnapshotBaseId
	}
	return ""
}

func (m *TreeChange) GetChangesData() []byte {
	if m != nil {
		return m.ChangesData
	}
	return nil
}

func (m *TreeChange) GetCurrentReadKeyHash() uint64 {
	if m != nil {
		return m.CurrentReadKeyHash
	}
	return 0
}

func (m *TreeChange) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TreeChange) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *TreeChange) GetIsSnapshot() bool {
	if m != nil {
		return m.IsSnapshot
	}
	return false
}

// RawTreeChange is a marshalled TreeChange (or RootChange) payload and a signature of this payload
type RawTreeChange struct {
	// Payload is a byte payload containing TreeChange
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// Signature is a signature made by identity indicated in the TreeChange payload
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RawTreeChange) Reset()         { *m = RawTreeChange{} }
func (m *RawTreeChange) String() string { return proto.CompactTextString(m) }
func (*RawTreeChange) ProtoMessage()    {}
func (*RawTreeChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f177d8514fae978f, []int{2}
}
func (m *RawTreeChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawTreeChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawTreeChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawTreeChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTreeChange.Merge(m, src)
}
func (m *RawTreeChange) XXX_Size() int {
	return m.Size()
}
func (m *RawTreeChange) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTreeChange.DiscardUnknown(m)
}

var xxx_messageInfo_RawTreeChange proto.InternalMessageInfo

func (m *RawTreeChange) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RawTreeChange) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// RawTreeChangeWithId is a marshalled RawTreeChange with CID
type RawTreeChangeWithId struct {
	// RawChange is a byte payload of RawTreeChange
	RawChange []byte `protobuf:"bytes,1,opt,name=rawChange,proto3" json:"rawChange,omitempty"`
	// Id is a cid made from rawChange payload
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *RawTreeChangeWithId) Reset()         { *m = RawTreeChangeWithId{} }
func (m *RawTreeChangeWithId) String() string { return proto.CompactTextString(m) }
func (*RawTreeChangeWithId) ProtoMessage()    {}
func (*RawTreeChangeWithId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f177d8514fae978f, []int{3}
}
func (m *RawTreeChangeWithId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawTreeChangeWithId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawTreeChangeWithId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawTreeChangeWithId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTreeChangeWithId.Merge(m, src)
}
func (m *RawTreeChangeWithId) XXX_Size() int {
	return m.Size()
}
func (m *RawTreeChangeWithId) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTreeChangeWithId.DiscardUnknown(m)
}

var xxx_messageInfo_RawTreeChangeWithId proto.InternalMessageInfo

func (m *RawTreeChangeWithId) GetRawChange() []byte {
	if m != nil {
		return m.RawChange
	}
	return nil
}

func (m *RawTreeChangeWithId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*RootChange)(nil), "treechange.RootChange")
	proto.RegisterType((*TreeChange)(nil), "treechange.TreeChange")
	proto.RegisterType((*RawTreeChange)(nil), "treechange.RawTreeChange")
	proto.RegisterType((*RawTreeChangeWithId)(nil), "treechange.RawTreeChangeWithId")
}

func init() {
	proto.RegisterFile("pkg/acl/treechangeproto/protos/treechange.proto", fileDescriptor_f177d8514fae978f)
}

var fileDescriptor_f177d8514fae978f = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x0e, 0x93, 0x40,
	0x10, 0x87, 0xbb, 0xb4, 0xf6, 0xcf, 0x88, 0x3d, 0xac, 0x07, 0x37, 0xc6, 0x10, 0xc2, 0xc1, 0x70,
	0x2a, 0x31, 0xbe, 0x41, 0x6b, 0x62, 0x1b, 0x6f, 0x6b, 0x13, 0x13, 0x6f, 0x23, 0x4c, 0xca, 0xc6,
	0x16, 0x08, 0xbb, 0x4d, 0xc3, 0x5b, 0xf8, 0x08, 0x3e, 0x83, 0x4f, 0xe1, 0xb1, 0x47, 0x8f, 0xa6,
	0x7d, 0x11, 0xc3, 0xd2, 0x0a, 0x34, 0x7a, 0x01, 0xe6, 0x9b, 0x30, 0xcc, 0xef, 0x63, 0x21, 0x2a,
	0xbe, 0xee, 0x22, 0x8c, 0xf7, 0x91, 0x29, 0x89, 0xe2, 0x14, 0xb3, 0x1d, 0x15, 0x65, 0x6e, 0xf2,
	0xc8, 0x5e, 0x75, 0x07, 0x2f, 0x2c, 0xe1, 0xd0, 0x92, 0xe0, 0x07, 0x03, 0x90, 0x79, 0x6e, 0x56,
	0xb6, 0xe4, 0xaf, 0x60, 0x86, 0xf1, 0x7e, 0x4d, 0x98, 0x6c, 0x12, 0xc1, 0x7c, 0x16, 0xce, 0x64,
	0x0b, 0xb8, 0x80, 0x89, 0x2e, 0x30, 0xa6, 0x4d, 0x22, 0x1c, 0xdb, 0xbb, 0x97, 0xdc, 0x03, 0x68,
	0x06, 0x6e, 0xab, 0x82, 0xc4, 0xd0, 0x36, 0x3b, 0xa4, 0x9e, 0x6b, 0xd4, 0x81, 0xb4, 0xc1, 0x43,
	0x21, 0x46, 0x3e, 0x0b, 0x87, 0xb2, 0x05, 0x9c, 0xc3, 0x48, 0x13, 0x25, 0xe2, 0x89, 0xcf, 0x42,
	0x57, 0xda, 0x67, 0xfe, 0x12, 0xa6, 0x2a, 0xa1, 0xcc, 0x28, 0x53, 0x89, 0xb1, 0xe5, 0x7f, 0xeb,
	0xe0, 0xbb, 0x03, 0xb0, 0x2d, 0x89, 0x6e, 0x4b, 0xfb, 0xf0, 0xb4, 0x4e, 0xd4, 0x2c, 0xa9, 0x05,
	0xf3, 0x87, 0xe1, 0x4c, 0x76, 0x51, 0x3f, 0x96, 0xf3, 0x18, 0xeb, 0x35, 0xcc, 0x75, 0x86, 0x85,
	0x4e, 0x73, 0xb3, 0x44, 0x5d, 0xa7, 0x6b, 0x02, 0x3c, 0xd0, 0xfa, 0x3b, 0x4d, 0x24, 0xfd, 0x0e,
	0x0d, 0xda, 0x18, 0xae, 0xec, 0x22, 0xbe, 0x00, 0x1e, 0x1f, 0xcb, 0x92, 0x32, 0x23, 0x09, 0x93,
	0x0f, 0x54, 0xad, 0x51, 0xa7, 0x36, 0xd6, 0x48, 0xfe, 0xa3, 0xd3, 0xd7, 0x32, 0x7e, 0xd4, 0xd2,
	0x55, 0x30, 0xe9, 0x2b, 0xa8, 0x85, 0x2b, 0xfd, 0xf1, 0xb6, 0x9f, 0x98, 0xfa, 0x2c, 0x9c, 0xca,
	0x0e, 0x09, 0xde, 0xc3, 0x33, 0x89, 0xa7, 0x8e, 0x24, 0x01, 0x93, 0x02, 0xab, 0x7d, 0x8e, 0xcd,
	0x7f, 0x75, 0xe5, 0xbd, 0xac, 0x97, 0xd0, 0x6a, 0x97, 0xa1, 0x39, 0x96, 0x64, 0xe5, 0xb8, 0xb2,
	0x05, 0xc1, 0x0a, 0x9e, 0xf7, 0x06, 0x7d, 0x52, 0x26, 0xdd, 0xd8, 0x97, 0x4a, 0x3c, 0x35, 0xe8,
	0x36, 0xb0, 0x05, 0x7c, 0x0e, 0x8e, 0xba, 0x8b, 0x76, 0x54, 0xb2, 0x7c, 0xf3, 0xf3, 0xe2, 0xb1,
	0xf3, 0xc5, 0x63, 0xbf, 0x2f, 0x1e, 0xfb, 0x76, 0xf5, 0x06, 0xe7, 0xab, 0x37, 0xf8, 0x75, 0xf5,
	0x06, 0x9f, 0x5f, 0xfc, 0xe7, 0xf0, 0x7e, 0x19, 0xdb, 0xdb, 0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x78, 0xc6, 0x1e, 0xde, 0x02, 0x00, 0x00,
}

func (m *RootChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RootChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RootChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Timestamp != 0 {
		i = encodeVarintTreechange(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChangeType) > 0 {
		i -= len(m.ChangeType)
		copy(dAtA[i:], m.ChangeType)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.ChangeType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AclHeadId) > 0 {
		i -= len(m.AclHeadId)
		copy(dAtA[i:], m.AclHeadId)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.AclHeadId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TreeChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TreeChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TreeChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSnapshot {
		i--
		if m.IsSnapshot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintTreechange(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.CurrentReadKeyHash != 0 {
		i = encodeVarintTreechange(dAtA, i, uint64(m.CurrentReadKeyHash))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ChangesData) > 0 {
		i -= len(m.ChangesData)
		copy(dAtA[i:], m.ChangesData)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.ChangesData)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SnapshotBaseId) > 0 {
		i -= len(m.SnapshotBaseId)
		copy(dAtA[i:], m.SnapshotBaseId)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.SnapshotBaseId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AclHeadId) > 0 {
		i -= len(m.AclHeadId)
		copy(dAtA[i:], m.AclHeadId)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.AclHeadId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TreeHeadIds) > 0 {
		for iNdEx := len(m.TreeHeadIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TreeHeadIds[iNdEx])
			copy(dAtA[i:], m.TreeHeadIds[iNdEx])
			i = encodeVarintTreechange(dAtA, i, uint64(len(m.TreeHeadIds[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RawTreeChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawTreeChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawTreeChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RawTreeChangeWithId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawTreeChangeWithId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawTreeChangeWithId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RawChange) > 0 {
		i -= len(m.RawChange)
		copy(dAtA[i:], m.RawChange)
		i = encodeVarintTreechange(dAtA, i, uint64(len(m.RawChange)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreechange(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreechange(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RootChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AclHeadId)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.ChangeType)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTreechange(uint64(m.Timestamp))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	return n
}

func (m *TreeChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TreeHeadIds) > 0 {
		for _, s := range m.TreeHeadIds {
			l = len(s)
			n += 1 + l + sovTreechange(uint64(l))
		}
	}
	l = len(m.AclHeadId)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.SnapshotBaseId)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.ChangesData)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	if m.CurrentReadKeyHash != 0 {
		n += 1 + sovTreechange(uint64(m.CurrentReadKeyHash))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTreechange(uint64(m.Timestamp))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	if m.IsSnapshot {
		n += 2
	}
	return n
}

func (m *RawTreeChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	return n
}

func (m *RawTreeChangeWithId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawChange)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTreechange(uint64(l))
	}
	return n
}

func sovTreechange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreechange(x uint64) (n int) {
	return sovTreechange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RootChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreechange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RootChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RootChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AclHeadId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AclHeadId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChangeType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed[:0], dAtA[iNdEx:postIndex]...)
			if m.Seed == nil {
				m.Seed = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = append(m.Identity[:0], dAtA[iNdEx:postIndex]...)
			if m.Identity == nil {
				m.Identity = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreechange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreechange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TreeChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreechange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TreeChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TreeChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeHeadIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeHeadIds = append(m.TreeHeadIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AclHeadId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AclHeadId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotBaseId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotBaseId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangesData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChangesData = append(m.ChangesData[:0], dAtA[iNdEx:postIndex]...)
			if m.ChangesData == nil {
				m.ChangesData = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentReadKeyHash", wireType)
			}
			m.CurrentReadKeyHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentReadKeyHash |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = append(m.Identity[:0], dAtA[iNdEx:postIndex]...)
			if m.Identity == nil {
				m.Identity = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSnapshot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSnapshot = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTreechange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreechange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RawTreeChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreechange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawTreeChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawTreeChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreechange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreechange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RawTreeChangeWithId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreechange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawTreeChangeWithId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawTreeChangeWithId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawChange", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawChange = append(m.RawChange[:0], dAtA[iNdEx:postIndex]...)
			if m.RawChange == nil {
				m.RawChange = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreechange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreechange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreechange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreechange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTreechange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreechange
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTreechange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTreechange
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreechange
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreechange
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreechange        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreechange          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreechange = fmt.Errorf("proto: unexpected end of group")
)
