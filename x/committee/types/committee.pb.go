// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/committee/v1beta1/committee.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TallyOption enumerates the valid types of a tally.
type TallyOption int32

const (
	// TALLY_OPTION_UNSPECIFIED defines a null tally option.
	TALLY_OPTION_UNSPECIFIED TallyOption = 0
	// Votes are tallied each block and the proposal passes as soon as the vote threshold is reached
	TALLY_OPTION_FIRST_PAST_THE_POST TallyOption = 1
	// Votes are tallied exactly once, when the deadline time is reached
	TALLY_OPTION_DEADLINE TallyOption = 2
)

var TallyOption_name = map[int32]string{
	0: "TALLY_OPTION_UNSPECIFIED",
	1: "TALLY_OPTION_FIRST_PAST_THE_POST",
	2: "TALLY_OPTION_DEADLINE",
}

var TallyOption_value = map[string]int32{
	"TALLY_OPTION_UNSPECIFIED":         0,
	"TALLY_OPTION_FIRST_PAST_THE_POST": 1,
	"TALLY_OPTION_DEADLINE":            2,
}

func (TallyOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a2549fd9d70ca349, []int{0}
}

// BaseCommittee is a common type shared by all Committees
type BaseCommittee struct {
	Id          uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Members     []string     `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Permissions []*types.Any `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Smallest percentage that must vote for a proposal to pass
	VoteThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"vote_threshold"`
	// The length of time a proposal remains active for. Proposals will close earlier if they get enough votes.
	ProposalDuration time.Duration `protobuf:"bytes,6,opt,name=proposal_duration,json=proposalDuration,proto3,stdduration" json:"proposal_duration"`
	TallyOption      TallyOption   `protobuf:"varint,7,opt,name=tally_option,json=tallyOption,proto3,enum=kava.committee.v1beta1.TallyOption" json:"tally_option,omitempty"`
}

func (m *BaseCommittee) Reset()         { *m = BaseCommittee{} }
func (m *BaseCommittee) String() string { return proto.CompactTextString(m) }
func (*BaseCommittee) ProtoMessage()    {}
func (*BaseCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2549fd9d70ca349, []int{0}
}
func (m *BaseCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseCommittee.Merge(m, src)
}
func (m *BaseCommittee) XXX_Size() int {
	return m.Size()
}
func (m *BaseCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_BaseCommittee proto.InternalMessageInfo

// MemberCommittee is an alias of BaseCommittee
type MemberCommittee struct {
	*BaseCommittee `protobuf:"bytes,1,opt,name=base_committee,json=baseCommittee,proto3,embedded=base_committee" json:"base_committee,omitempty"`
}

func (m *MemberCommittee) Reset()         { *m = MemberCommittee{} }
func (m *MemberCommittee) String() string { return proto.CompactTextString(m) }
func (*MemberCommittee) ProtoMessage()    {}
func (*MemberCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2549fd9d70ca349, []int{1}
}
func (m *MemberCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberCommittee.Merge(m, src)
}
func (m *MemberCommittee) XXX_Size() int {
	return m.Size()
}
func (m *MemberCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_MemberCommittee proto.InternalMessageInfo

// TokenCommittee supports voting on proposals by token holders
type TokenCommittee struct {
	*BaseCommittee `protobuf:"bytes,1,opt,name=base_committee,json=baseCommittee,proto3,embedded=base_committee" json:"base_committee,omitempty"`
	Quorum         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=quorum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quorum"`
	TallyDenom     string                                 `protobuf:"bytes,3,opt,name=tally_denom,json=tallyDenom,proto3" json:"tally_denom,omitempty"`
}

func (m *TokenCommittee) Reset()         { *m = TokenCommittee{} }
func (m *TokenCommittee) String() string { return proto.CompactTextString(m) }
func (*TokenCommittee) ProtoMessage()    {}
func (*TokenCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2549fd9d70ca349, []int{2}
}
func (m *TokenCommittee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenCommittee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenCommittee.Merge(m, src)
}
func (m *TokenCommittee) XXX_Size() int {
	return m.Size()
}
func (m *TokenCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_TokenCommittee proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("kava.committee.v1beta1.TallyOption", TallyOption_name, TallyOption_value)
	proto.RegisterType((*BaseCommittee)(nil), "kava.committee.v1beta1.BaseCommittee")
	proto.RegisterType((*MemberCommittee)(nil), "kava.committee.v1beta1.MemberCommittee")
	proto.RegisterType((*TokenCommittee)(nil), "kava.committee.v1beta1.TokenCommittee")
}

func init() {
	proto.RegisterFile("kava/committee/v1beta1/committee.proto", fileDescriptor_a2549fd9d70ca349)
}

var fileDescriptor_a2549fd9d70ca349 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xf5, 0x24, 0xfd, 0xfa, 0x33, 0x6e, 0xf2, 0x85, 0x51, 0x41, 0x6e, 0x85, 0x1c, 0xab, 0x40,
	0x65, 0x81, 0x62, 0xab, 0x61, 0xc7, 0x2e, 0x69, 0x12, 0x11, 0xa9, 0x34, 0x96, 0xe3, 0x2e, 0x60,
	0x63, 0xd9, 0xf1, 0x90, 0x5a, 0xb5, 0x3d, 0xc6, 0x33, 0x89, 0xc8, 0x1b, 0xb0, 0x64, 0xc9, 0xbe,
	0x2c, 0x78, 0x80, 0x3e, 0x44, 0xc5, 0xaa, 0x62, 0x85, 0x10, 0x2a, 0xd0, 0xbe, 0x08, 0xf2, 0x5f,
	0xe3, 0x06, 0x58, 0xb0, 0x60, 0xe5, 0xb9, 0xf7, 0x9c, 0x3b, 0xf7, 0xce, 0x39, 0x57, 0x86, 0x3b,
	0xc7, 0xd6, 0xd4, 0x52, 0x47, 0xc4, 0xf7, 0x5d, 0xc6, 0x30, 0x56, 0xa7, 0xbb, 0x36, 0x66, 0xd6,
	0xee, 0x3c, 0xa3, 0x84, 0x11, 0x61, 0x04, 0xdd, 0x89, 0x79, 0xca, 0x3c, 0x9b, 0xf1, 0xb6, 0x36,
	0xc6, 0x64, 0x4c, 0x12, 0x8a, 0x1a, 0x9f, 0x52, 0xf6, 0xd6, 0xe6, 0x88, 0x50, 0x9f, 0x50, 0x33,
	0x05, 0xd2, 0x20, 0x87, 0xc6, 0x84, 0x8c, 0x3d, 0xac, 0x26, 0x91, 0x3d, 0x79, 0xa9, 0x5a, 0xc1,
	0x2c, 0x83, 0xc4, 0x45, 0xc8, 0x99, 0x44, 0x16, 0x73, 0x49, 0x90, 0xe2, 0xdb, 0x27, 0x65, 0x58,
	0x69, 0x5b, 0x14, 0xef, 0xe5, 0x53, 0xa0, 0x2a, 0x2c, 0xb9, 0x8e, 0x00, 0x24, 0x20, 0x2f, 0xe9,
	0x25, 0xd7, 0x41, 0x12, 0xe4, 0x1d, 0x4c, 0x47, 0x91, 0x1b, 0xc6, 0x65, 0x42, 0x49, 0x02, 0xf2,
	0x9a, 0x5e, 0x4c, 0x21, 0x01, 0xae, 0xf8, 0xd8, 0xb7, 0x71, 0x44, 0x85, 0xb2, 0x54, 0x96, 0xd7,
	0xf4, 0x3c, 0x44, 0x3d, 0xc8, 0x87, 0x38, 0xf2, 0x5d, 0x4a, 0x5d, 0x12, 0x50, 0x61, 0x49, 0x2a,
	0xcb, 0x7c, 0x73, 0x43, 0x49, 0x67, 0x52, 0xf2, 0x99, 0x94, 0x56, 0x30, 0x6b, 0x57, 0x3f, 0x9e,
	0x36, 0xa0, 0x76, 0x4d, 0xd6, 0x8b, 0x85, 0xe8, 0x10, 0x56, 0xa7, 0x84, 0x61, 0x93, 0x1d, 0x45,
	0x98, 0x1e, 0x11, 0xcf, 0x11, 0xfe, 0x93, 0x80, 0xbc, 0xde, 0x56, 0xce, 0x2e, 0xea, 0xdc, 0x97,
	0x8b, 0xfa, 0xce, 0xd8, 0x65, 0x47, 0x13, 0x3b, 0xd6, 0x32, 0x53, 0x26, 0xfb, 0x34, 0xa8, 0x73,
	0xac, 0xb2, 0x59, 0x88, 0xa9, 0xd2, 0xc1, 0x23, 0xbd, 0x12, 0xdf, 0x62, 0xe4, 0x97, 0x20, 0x0d,
	0xde, 0x0a, 0x23, 0x12, 0x12, 0x6a, 0x79, 0x66, 0xae, 0x8b, 0xb0, 0x2c, 0x01, 0x99, 0x6f, 0x6e,
	0xfe, 0x32, 0x64, 0x27, 0x23, 0xb4, 0x57, 0xe3, 0xa6, 0xef, 0xbe, 0xd5, 0x81, 0x5e, 0xcb, 0xab,
	0x73, 0x0c, 0xf5, 0xe0, 0x3a, 0xb3, 0x3c, 0x6f, 0x66, 0x92, 0x54, 0xad, 0x15, 0x09, 0xc8, 0xd5,
	0xe6, 0x3d, 0xe5, 0xf7, 0x4e, 0x2b, 0x46, 0xcc, 0x1d, 0x24, 0x54, 0x9d, 0x67, 0xf3, 0xe0, 0x49,
	0xe5, 0xd3, 0x69, 0x63, 0xed, 0xda, 0x93, 0x6d, 0x06, 0xff, 0x7f, 0x96, 0x48, 0x3a, 0xb7, 0x49,
	0x87, 0x55, 0xdb, 0xa2, 0xd8, 0xbc, 0xbe, 0x34, 0xb1, 0x8c, 0x6f, 0x3e, 0xf8, 0x53, 0xaf, 0x1b,
	0x2e, 0xb7, 0x97, 0xce, 0x2f, 0xea, 0x40, 0xaf, 0xd8, 0xc5, 0xe4, 0x62, 0xd7, 0xaf, 0x00, 0x56,
	0x0d, 0x72, 0x8c, 0x83, 0x7f, 0xda, 0x15, 0xf5, 0xe0, 0xf2, 0xab, 0x09, 0x89, 0x26, 0x7e, 0xb2,
	0x5b, 0x7f, 0x6f, 0x6a, 0x56, 0x8d, 0xea, 0x30, 0x95, 0xd0, 0x74, 0x70, 0x40, 0x7c, 0xa1, 0x9c,
	0x2c, 0x2a, 0x4c, 0x52, 0x9d, 0x38, 0xb3, 0xf0, 0xbc, 0x87, 0x53, 0xc8, 0x17, 0xf4, 0x47, 0x77,
	0xa1, 0x60, 0xb4, 0xf6, 0xf7, 0x9f, 0x9b, 0x03, 0xcd, 0xe8, 0x0f, 0x0e, 0xcc, 0xc3, 0x83, 0xa1,
	0xd6, 0xdd, 0xeb, 0xf7, 0xfa, 0xdd, 0x4e, 0x8d, 0x43, 0xf7, 0xa1, 0x74, 0x03, 0xed, 0xf5, 0xf5,
	0xa1, 0x61, 0x6a, 0xad, 0xa1, 0x61, 0x1a, 0x4f, 0xbb, 0xa6, 0x36, 0x18, 0x1a, 0x35, 0x80, 0x36,
	0xe1, 0xed, 0x1b, 0xac, 0x4e, 0xb7, 0xd5, 0xd9, 0xef, 0x1f, 0x74, 0x6b, 0xa5, 0xad, 0xd5, 0x37,
	0x27, 0x22, 0xf7, 0xe1, 0xbd, 0xc8, 0xb5, 0xfb, 0x67, 0x3f, 0x44, 0xee, 0xec, 0x52, 0x04, 0xe7,
	0x97, 0x22, 0xf8, 0x7e, 0x29, 0x82, 0xb7, 0x57, 0x22, 0x77, 0x7e, 0x25, 0x72, 0x9f, 0xaf, 0x44,
	0xee, 0xc5, 0xa3, 0xc2, 0xab, 0x63, 0x4d, 0x1b, 0x9e, 0x65, 0xd3, 0xe4, 0xa4, 0xbe, 0x2e, 0xfc,
	0x53, 0x92, 0xe7, 0xdb, 0xcb, 0xc9, 0x76, 0x3e, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x66,
	0xff, 0x61, 0x72, 0x04, 0x00, 0x00,
}

func (m *BaseCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TallyOption != 0 {
		i = encodeVarintCommittee(dAtA, i, uint64(m.TallyOption))
		i--
		dAtA[i] = 0x38
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ProposalDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ProposalDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCommittee(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommittee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommittee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintCommittee(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCommittee(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCommittee(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseCommittee != nil {
		{
			size, err := m.BaseCommittee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommittee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenCommittee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenCommittee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenCommittee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TallyDenom) > 0 {
		i -= len(m.TallyDenom)
		copy(dAtA[i:], m.TallyDenom)
		i = encodeVarintCommittee(dAtA, i, uint64(len(m.TallyDenom)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Quorum.Size()
		i -= size
		if _, err := m.Quorum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommittee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.BaseCommittee != nil {
		{
			size, err := m.BaseCommittee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommittee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommittee(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommittee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCommittee(uint64(m.Id))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCommittee(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovCommittee(uint64(l))
		}
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovCommittee(uint64(l))
		}
	}
	l = m.VoteThreshold.Size()
	n += 1 + l + sovCommittee(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ProposalDuration)
	n += 1 + l + sovCommittee(uint64(l))
	if m.TallyOption != 0 {
		n += 1 + sovCommittee(uint64(m.TallyOption))
	}
	return n
}

func (m *MemberCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseCommittee != nil {
		l = m.BaseCommittee.Size()
		n += 1 + l + sovCommittee(uint64(l))
	}
	return n
}

func (m *TokenCommittee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseCommittee != nil {
		l = m.BaseCommittee.Size()
		n += 1 + l + sovCommittee(uint64(l))
	}
	l = m.Quorum.Size()
	n += 1 + l + sovCommittee(uint64(l))
	l = len(m.TallyDenom)
	if l > 0 {
		n += 1 + l + sovCommittee(uint64(l))
	}
	return n
}

func sovCommittee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommittee(x uint64) (n int) {
	return sovCommittee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: BaseCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &types.Any{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ProposalDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyOption", wireType)
			}
			m.TallyOption = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TallyOption |= TallyOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func (m *MemberCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: MemberCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCommittee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseCommittee == nil {
				m.BaseCommittee = &BaseCommittee{}
			}
			if err := m.BaseCommittee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func (m *TokenCommittee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittee
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
			return fmt.Errorf("proto: TokenCommittee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenCommittee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCommittee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseCommittee == nil {
				m.BaseCommittee = &BaseCommittee{}
			}
			if err := m.BaseCommittee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quorum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittee
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
				return ErrInvalidLengthCommittee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TallyDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommittee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittee
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
func skipCommittee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommittee
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
					return 0, ErrIntOverflowCommittee
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
					return 0, ErrIntOverflowCommittee
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
				return 0, ErrInvalidLengthCommittee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommittee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommittee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommittee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommittee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommittee = fmt.Errorf("proto: unexpected end of group")
)
