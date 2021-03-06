// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/group.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/ovrclk/akash/types"
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

type Group_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	GroupStateInvalid      Group_State = 0
	GroupOpen              Group_State = 1
	GroupOrdered           Group_State = 2
	GroupMatched           Group_State = 3
	GroupInsufficientFunds Group_State = 4
	GroupClosed            Group_State = 5
)

var Group_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "ordered",
	3: "matched",
	4: "insufficient_funds",
	5: "closed",
}

var Group_State_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"ordered":            2,
	"matched":            3,
	"insufficient_funds": 4,
	"closed":             5,
}

func (x Group_State) String() string {
	return proto.EnumName(Group_State_name, int32(x))
}

func (Group_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{3, 0}
}

// MsgCloseGroup defines SDK message to close a single Group within a Deployment.
type MsgCloseGroup struct {
	ID GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseGroup) Reset()         { *m = MsgCloseGroup{} }
func (m *MsgCloseGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCloseGroup) ProtoMessage()    {}
func (*MsgCloseGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{0}
}
func (m *MsgCloseGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseGroup.Merge(m, src)
}
func (m *MsgCloseGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseGroup proto.InternalMessageInfo

func (m *MsgCloseGroup) GetID() GroupID {
	if m != nil {
		return m.ID
	}
	return GroupID{}
}

// GroupID stores owner, deployment sequence number and group sequence number
type GroupID struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq  uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
}

func (m *GroupID) Reset()      { *m = GroupID{} }
func (*GroupID) ProtoMessage() {}
func (*GroupID) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{1}
}
func (m *GroupID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupID.Merge(m, src)
}
func (m *GroupID) XXX_Size() int {
	return m.Size()
}
func (m *GroupID) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupID.DiscardUnknown(m)
}

var xxx_messageInfo_GroupID proto.InternalMessageInfo

func (m *GroupID) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GroupID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *GroupID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

type GroupSpec struct {
	Name             string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Requirements     []types.Attribute `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements" yaml:"requirements"`
	Resources        []Resource        `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources" yaml:"resources"`
	OrderBidDuration int64             `protobuf:"varint,4,opt,name=orderBidDuration,proto3" json:"order-bid-duration" yaml:"order-bid-duration"`
}

func (m *GroupSpec) Reset()         { *m = GroupSpec{} }
func (m *GroupSpec) String() string { return proto.CompactTextString(m) }
func (*GroupSpec) ProtoMessage()    {}
func (*GroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{2}
}
func (m *GroupSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSpec.Merge(m, src)
}
func (m *GroupSpec) XXX_Size() int {
	return m.Size()
}
func (m *GroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSpec proto.InternalMessageInfo

type Group struct {
	GroupID   GroupID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Group_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.Group_State" json:"state" yaml:"state"`
	GroupSpec GroupSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{3}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGroupID() GroupID {
	if m != nil {
		return m.GroupID
	}
	return GroupID{}
}

func (m *Group) GetState() Group_State {
	if m != nil {
		return m.State
	}
	return GroupStateInvalid
}

func (m *Group) GetGroupSpec() GroupSpec {
	if m != nil {
		return m.GroupSpec
	}
	return GroupSpec{}
}

type Resource struct {
	Unit  types1.Unit `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit" yaml:"unit"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price types.Coin  `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetUnit() types1.Unit {
	if m != nil {
		return m.Unit
	}
	return types1.Unit{}
}

func (m *Resource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Resource) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("akash.deployment.Group_State", Group_State_name, Group_State_value)
	proto.RegisterType((*MsgCloseGroup)(nil), "akash.deployment.MsgCloseGroup")
	proto.RegisterType((*GroupID)(nil), "akash.deployment.GroupID")
	proto.RegisterType((*GroupSpec)(nil), "akash.deployment.GroupSpec")
	proto.RegisterType((*Group)(nil), "akash.deployment.Group")
	proto.RegisterType((*Resource)(nil), "akash.deployment.Resource")
}

func init() { proto.RegisterFile("akash/deployment/group.proto", fileDescriptor_940c80f4f1e11e54) }

var fileDescriptor_940c80f4f1e11e54 = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0x13, 0x67, 0xdb, 0x4e, 0x5a, 0x30, 0xc3, 0xc2, 0xa6, 0x2e, 0xcd, 0x58, 0xb3, 0x42,
	0x8a, 0x58, 0xd5, 0xd6, 0x66, 0x6f, 0x15, 0x97, 0x7a, 0x23, 0x4a, 0x91, 0x96, 0x95, 0x5c, 0x21,
	0x10, 0x97, 0xe2, 0x78, 0xa6, 0xe9, 0xa8, 0x89, 0xc7, 0xf5, 0xd8, 0x85, 0xfe, 0x07, 0xab, 0x1c,
	0x10, 0x47, 0x2e, 0x91, 0x2a, 0xf1, 0xcf, 0xec, 0x8d, 0x3d, 0xc2, 0xc5, 0xa0, 0xf4, 0xb2, 0x8a,
	0x38, 0xe5, 0xc8, 0x09, 0xcd, 0x8c, 0x83, 0xd3, 0x2d, 0x2b, 0x71, 0x4a, 0xde, 0xf7, 0x7d, 0xef,
	0xcd, 0x8f, 0xf7, 0xcd, 0x33, 0xf8, 0x28, 0x3c, 0x0f, 0xc5, 0x99, 0x47, 0x68, 0x32, 0xe2, 0x57,
	0x63, 0x1a, 0x67, 0xde, 0x30, 0xe5, 0x79, 0xe2, 0x26, 0x29, 0xcf, 0x38, 0xb4, 0x14, 0xeb, 0x56,
	0xac, 0x7d, 0x7f, 0xc8, 0x87, 0x5c, 0x91, 0x9e, 0xfc, 0xa7, 0x75, 0xf6, 0x7d, 0x5d, 0x25, 0xa5,
	0x82, 0xe7, 0x69, 0x44, 0x4b, 0xb4, 0x13, 0x71, 0x31, 0xe6, 0xc2, 0x1b, 0x84, 0x82, 0x7a, 0x97,
	0x8f, 0x07, 0x34, 0x0b, 0x1f, 0x7b, 0x11, 0x67, 0x71, 0xc9, 0x3f, 0x5c, 0xe5, 0xc3, 0x41, 0xc4,
	0xfe, 0x15, 0xc9, 0x40, 0x8b, 0xf0, 0x09, 0xd8, 0x7a, 0x26, 0x86, 0x4f, 0x47, 0x5c, 0xd0, 0x43,
	0xb9, 0x33, 0xf8, 0x39, 0xa8, 0x33, 0xd2, 0x36, 0x1c, 0xa3, 0xdb, 0xea, 0x6d, 0xbb, 0x6f, 0x6e,
	0xd0, 0x55, 0xa2, 0xa3, 0xbe, 0xbf, 0xfb, 0xb2, 0x40, 0xb5, 0x59, 0x81, 0xea, 0x47, 0xfd, 0x79,
	0x81, 0xea, 0x8c, 0x2c, 0x0a, 0xb4, 0x71, 0x15, 0x8e, 0x47, 0xfb, 0x98, 0x11, 0x1c, 0xd4, 0x19,
	0xd9, 0x37, 0x5f, 0x5f, 0xa3, 0x1a, 0xfe, 0xc3, 0x00, 0x6b, 0x65, 0x12, 0xfc, 0x0e, 0x34, 0xf9,
	0xf7, 0x31, 0x4d, 0x55, 0xf9, 0x4d, 0xff, 0x8b, 0x79, 0x81, 0x34, 0xb0, 0x28, 0xd0, 0xa6, 0x4e,
	0x56, 0x21, 0xfe, 0xbb, 0x40, 0x7b, 0x43, 0x96, 0x9d, 0xe5, 0x03, 0x37, 0xe2, 0x63, 0xaf, 0x3c,
	0x88, 0xfe, 0xd9, 0x13, 0xe4, 0xdc, 0xcb, 0xae, 0x12, 0x2a, 0xdc, 0x83, 0x28, 0x3a, 0x20, 0x24,
	0xa5, 0x42, 0x04, 0xba, 0x0e, 0x7c, 0x02, 0x4c, 0x22, 0xe8, 0x45, 0xbb, 0xee, 0x18, 0x5d, 0xd3,
	0x47, 0xb3, 0x02, 0x99, 0xfd, 0x63, 0x7a, 0x31, 0x2f, 0x90, 0xc2, 0x17, 0x05, 0x6a, 0xe9, 0x75,
	0x64, 0x84, 0x03, 0x05, 0xca, 0xa4, 0xa1, 0x4c, 0x6a, 0x38, 0x46, 0x77, 0x4b, 0x27, 0x1d, 0x96,
	0x49, 0xc3, 0x5b, 0x49, 0x43, 0x9d, 0x24, 0x7f, 0xf6, 0xd7, 0x7f, 0xbe, 0x46, 0xb5, 0xd7, 0xd7,
	0xc8, 0xc0, 0x3f, 0x36, 0xc0, 0x86, 0x3a, 0xe1, 0x71, 0x42, 0x23, 0xf8, 0x08, 0x98, 0x71, 0x38,
	0xa6, 0xea, 0x88, 0x1b, 0xfe, 0x03, 0x59, 0x44, 0xc6, 0x55, 0x11, 0x19, 0xe1, 0x40, 0x81, 0x30,
	0x03, 0x9b, 0x29, 0xbd, 0xc8, 0x59, 0x4a, 0xe5, 0xe5, 0x8a, 0x76, 0xdd, 0x69, 0x74, 0x5b, 0xbd,
	0x87, 0xae, 0x3e, 0xa9, 0x2b, 0x3b, 0xe7, 0xaa, 0x66, 0x95, 0x9d, 0x73, 0x0f, 0xb2, 0x2c, 0x65,
	0x83, 0x3c, 0xa3, 0xfe, 0x23, 0xd9, 0x80, 0x79, 0x81, 0x6e, 0x15, 0x58, 0x14, 0xe8, 0x7d, 0xbd,
	0xca, 0x2a, 0x8a, 0x83, 0x5b, 0x22, 0x78, 0x02, 0x36, 0x96, 0x56, 0x12, 0xed, 0x86, 0x5a, 0xd2,
	0xbe, 0xdb, 0xe9, 0xa0, 0x94, 0xf8, 0x1f, 0x97, 0x2b, 0x55, 0x49, 0x8b, 0x02, 0x59, 0xcb, 0x65,
	0x4a, 0x08, 0x07, 0x15, 0x0d, 0x05, 0xb0, 0x78, 0x4a, 0x68, 0xea, 0x33, 0xd2, 0xcf, 0xd3, 0x30,
	0x63, 0x3c, 0x6e, 0x9b, 0x8e, 0xd1, 0x6d, 0xf8, 0x87, 0xb3, 0x02, 0x59, 0xcf, 0xdf, 0xe0, 0xe6,
	0x05, 0x82, 0x4a, 0xbf, 0x37, 0x60, 0x64, 0x8f, 0x94, 0xe8, 0xa2, 0x40, 0xdb, 0xa5, 0x27, 0xee,
	0x70, 0x38, 0xb8, 0xb3, 0xc0, 0xfe, 0xfa, 0x0b, 0xdd, 0x90, 0x1a, 0xfe, 0xab, 0x01, 0x9a, 0xda,
	0xcc, 0xdf, 0x80, 0x75, 0xf5, 0xde, 0x4e, 0xfe, 0x8f, 0xa5, 0x71, 0x69, 0xe9, 0xa5, 0x5d, 0xff,
	0xcb, 0xd7, 0x6b, 0xaa, 0xdc, 0x11, 0x81, 0x5f, 0x82, 0xa6, 0xc8, 0xc2, 0x8c, 0x2a, 0xa7, 0xbd,
	0xd3, 0xdb, 0x7d, 0x4b, 0x59, 0xf7, 0x58, 0x8a, 0xfc, 0x6d, 0xe9, 0x74, 0xa5, 0xaf, 0x9c, 0xae,
	0x42, 0x1c, 0x68, 0x18, 0x7e, 0x0d, 0x80, 0xde, 0xa9, 0x48, 0x68, 0xa4, 0x9c, 0xd8, 0xea, 0xed,
	0xbc, 0xa5, 0xa8, 0xf4, 0x99, 0xbf, 0x53, 0x76, 0xc5, 0x94, 0x09, 0x95, 0xbb, 0x64, 0x84, 0x83,
	0x8d, 0xe1, 0x52, 0x87, 0x7f, 0x37, 0x40, 0x53, 0x6d, 0x02, 0x62, 0xb0, 0xc6, 0xe2, 0xcb, 0x70,
	0xc4, 0x88, 0x55, 0xb3, 0x3f, 0x98, 0x4c, 0x9d, 0xf7, 0x74, 0x35, 0x49, 0x1e, 0x69, 0x02, 0x3e,
	0x00, 0x26, 0x4f, 0x68, 0x6c, 0x19, 0xf6, 0xd6, 0x64, 0xea, 0x68, 0x5b, 0x3f, 0x4f, 0x68, 0x0c,
	0x77, 0xc1, 0x9a, 0xba, 0x71, 0x4a, 0xac, 0xba, 0x6d, 0x4d, 0xa6, 0xce, 0xa6, 0xe6, 0x34, 0x26,
	0xe9, 0x71, 0x98, 0x45, 0x67, 0x94, 0x58, 0x8d, 0x15, 0xfa, 0x99, 0xc6, 0x60, 0x0f, 0x40, 0x16,
	0x8b, 0xfc, 0xf4, 0x94, 0x45, 0x8c, 0xc6, 0xd9, 0xc9, 0x69, 0x1e, 0x13, 0x61, 0x99, 0xb6, 0x3d,
	0x99, 0x3a, 0x1f, 0xea, 0xeb, 0x5e, 0xa1, 0x3f, 0x93, 0x2c, 0xdc, 0x01, 0xf7, 0x22, 0x39, 0x96,
	0x88, 0xd5, 0xb4, 0xdf, 0x9d, 0x4c, 0x9d, 0x96, 0xd2, 0xa9, 0x49, 0x45, 0x6c, 0xf3, 0xc5, 0x2f,
	0x9d, 0x5a, 0x39, 0x61, 0x7e, 0x35, 0xc0, 0xfa, 0xd2, 0xac, 0xf0, 0x53, 0x60, 0xe6, 0x31, 0xcb,
	0xca, 0x6e, 0xb7, 0xca, 0x1b, 0xfc, 0x2a, 0x66, 0x59, 0x75, 0x63, 0x52, 0x50, 0xdd, 0x98, 0x8c,
	0x70, 0xa0, 0x40, 0xe8, 0x81, 0x66, 0xc4, 0xf3, 0x38, 0x53, 0x5d, 0xdd, 0xd2, 0x6d, 0x53, 0x40,
	0xd5, 0x36, 0x15, 0xe2, 0x40, 0xc3, 0xd2, 0x06, 0x49, 0xca, 0x22, 0x5a, 0x76, 0x6c, 0xfb, 0xd6,
	0xcb, 0x5d, 0x3e, 0xda, 0xa7, 0x9c, 0xc5, 0x7a, 0x60, 0xca, 0x7a, 0x4a, 0x5f, 0xd5, 0x53, 0x21,
	0x0e, 0x34, 0xac, 0x4e, 0x64, 0xf8, 0xfd, 0x97, 0xb3, 0x8e, 0xf1, 0x6a, 0xd6, 0x31, 0xfe, 0x9c,
	0x75, 0x8c, 0x9f, 0x6e, 0x3a, 0xb5, 0x57, 0x37, 0x9d, 0xda, 0x6f, 0x37, 0x9d, 0xda, 0xb7, 0x9f,
	0xac, 0x4c, 0x45, 0x7e, 0x99, 0x46, 0xa3, 0x73, 0x4f, 0x7f, 0x1b, 0x7e, 0x58, 0xfd, 0xc6, 0xa8,
	0xe9, 0x38, 0xb8, 0xa7, 0x26, 0xfc, 0x93, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x0e, 0x6a,
	0xdd, 0x84, 0x06, 0x00, 0x00,
}

func (this *GroupID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GroupID)
	if !ok {
		that2, ok := that.(GroupID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if this.DSeq != that1.DSeq {
		return false
	}
	if this.GSeq != that1.GSeq {
		return false
	}
	return true
}
func (this *Resource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Resource)
	if !ok {
		that2, ok := that.(Resource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Unit.Equal(&that1.Unit) {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	return true
}
func (m *MsgCloseGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GroupID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderBidDuration != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.OrderBidDuration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GroupSpec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Unit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCloseGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *GroupID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovGroup(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovGroup(uint64(m.GSeq))
	}
	return n
}

func (m *GroupSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	if m.OrderBidDuration != 0 {
		n += 1 + sovGroup(uint64(m.OrderBidDuration))
	}
	return n
}

func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GroupID.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.State != 0 {
		n += 1 + sovGroup(uint64(m.State))
	}
	l = m.GroupSpec.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Unit.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.Count != 0 {
		n += 1 + sovGroup(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCloseGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: MsgCloseGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *GroupID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: GroupID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *GroupSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: GroupSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, types.Attribute{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, Resource{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBidDuration", wireType)
			}
			m.OrderBidDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderBidDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Group_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
