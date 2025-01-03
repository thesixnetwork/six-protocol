// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/collection_owner_request.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type OriginContractParam struct {
	Chain           string    `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	ContractAddress string    `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ContractOwner   string    `protobuf:"bytes,3,opt,name=contract_owner,json=contractOwner,proto3" json:"contract_owner,omitempty"`
	RequestExpire   time.Time `protobuf:"bytes,4,opt,name=request_expire,json=requestExpire,proto3,stdtime" json:"request_expire"`
}

func (m *OriginContractParam) Reset()         { *m = OriginContractParam{} }
func (m *OriginContractParam) String() string { return proto.CompactTextString(m) }
func (*OriginContractParam) ProtoMessage()    {}
func (*OriginContractParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{0}
}
func (m *OriginContractParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginContractParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginContractParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginContractParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginContractParam.Merge(m, src)
}
func (m *OriginContractParam) XXX_Size() int {
	return m.Size()
}
func (m *OriginContractParam) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginContractParam.DiscardUnknown(m)
}

var xxx_messageInfo_OriginContractParam proto.InternalMessageInfo

func (m *OriginContractParam) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *OriginContractParam) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *OriginContractParam) GetContractOwner() string {
	if m != nil {
		return m.ContractOwner
	}
	return ""
}

func (m *OriginContractParam) GetRequestExpire() time.Time {
	if m != nil {
		return m.RequestExpire
	}
	return time.Time{}
}

type CollectionOwnerRequest struct {
	Id              uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftSchemaCode   string                `protobuf:"bytes,2,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Signer          string                `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	RequiredConfirm uint64                `protobuf:"varint,4,opt,name=required_confirm,json=requiredConfirm,proto3" json:"required_confirm,omitempty"`
	Status          RequestStatus         `protobuf:"varint,5,opt,name=status,proto3,enum=thesixnetwork.sixprotocol.nftoracle.RequestStatus" json:"status,omitempty"`
	CurrentConfirm  uint64                `protobuf:"varint,6,opt,name=current_confirm,json=currentConfirm,proto3" json:"current_confirm,omitempty"`
	Confirmers      []string              `protobuf:"bytes,7,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
	CreatedAt       time.Time             `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ValidUntil      time.Time             `protobuf:"bytes,9,opt,name=valid_until,json=validUntil,proto3,stdtime" json:"valid_until"`
	ContractInfo    []*OriginContractInfo `protobuf:"bytes,10,rep,name=contract_info,json=contractInfo,proto3" json:"contract_info,omitempty"`
	ExpiredHeight   int64                 `protobuf:"varint,11,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
}

func (m *CollectionOwnerRequest) Reset()         { *m = CollectionOwnerRequest{} }
func (m *CollectionOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionOwnerRequest) ProtoMessage()    {}
func (*CollectionOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{1}
}
func (m *CollectionOwnerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectionOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectionOwnerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectionOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionOwnerRequest.Merge(m, src)
}
func (m *CollectionOwnerRequest) XXX_Size() int {
	return m.Size()
}
func (m *CollectionOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionOwnerRequest proto.InternalMessageInfo

func (m *CollectionOwnerRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CollectionOwnerRequest) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *CollectionOwnerRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *CollectionOwnerRequest) GetRequiredConfirm() uint64 {
	if m != nil {
		return m.RequiredConfirm
	}
	return 0
}

func (m *CollectionOwnerRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_PENDING
}

func (m *CollectionOwnerRequest) GetCurrentConfirm() uint64 {
	if m != nil {
		return m.CurrentConfirm
	}
	return 0
}

func (m *CollectionOwnerRequest) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func (m *CollectionOwnerRequest) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *CollectionOwnerRequest) GetValidUntil() time.Time {
	if m != nil {
		return m.ValidUntil
	}
	return time.Time{}
}

func (m *CollectionOwnerRequest) GetContractInfo() []*OriginContractInfo {
	if m != nil {
		return m.ContractInfo
	}
	return nil
}

func (m *CollectionOwnerRequest) GetExpiredHeight() int64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

type OriginContractInfo struct {
	ContractOriginDataInfo *OriginContractParam `protobuf:"bytes,1,opt,name=contractOriginDataInfo,proto3" json:"contractOriginDataInfo,omitempty"`
	Hash                   []byte               `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Confirmers             []string             `protobuf:"bytes,3,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
}

func (m *OriginContractInfo) Reset()         { *m = OriginContractInfo{} }
func (m *OriginContractInfo) String() string { return proto.CompactTextString(m) }
func (*OriginContractInfo) ProtoMessage()    {}
func (*OriginContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{2}
}
func (m *OriginContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginContractInfo.Merge(m, src)
}
func (m *OriginContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *OriginContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OriginContractInfo proto.InternalMessageInfo

func (m *OriginContractInfo) GetContractOriginDataInfo() *OriginContractParam {
	if m != nil {
		return m.ContractOriginDataInfo
	}
	return nil
}

func (m *OriginContractInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *OriginContractInfo) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func init() {
	proto.RegisterType((*OriginContractParam)(nil), "thesixnetwork.sixprotocol.nftoracle.OriginContractParam")
	proto.RegisterType((*CollectionOwnerRequest)(nil), "thesixnetwork.sixprotocol.nftoracle.CollectionOwnerRequest")
	proto.RegisterType((*OriginContractInfo)(nil), "thesixnetwork.sixprotocol.nftoracle.OriginContractInfo")
}

func init() {
	proto.RegisterFile("nftoracle/collection_owner_request.proto", fileDescriptor_5b5e5fd2fa665471)
}

var fileDescriptor_5b5e5fd2fa665471 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x6f, 0xd6, 0xae, 0xac, 0xee, 0xda, 0x21, 0x33, 0x8d, 0x68, 0x87, 0xac, 0x1a, 0x20, 0xc2,
	0x81, 0x44, 0x2a, 0x87, 0x71, 0xdd, 0xca, 0x24, 0xfe, 0x1c, 0x06, 0x19, 0x5c, 0x10, 0x52, 0xe4,
	0x39, 0x4e, 0x62, 0x91, 0xd8, 0xc5, 0x76, 0x58, 0xf9, 0x16, 0xfb, 0x36, 0x48, 0x7c, 0x82, 0x89,
	0xd3, 0x8e, 0x9c, 0x00, 0x6d, 0x5f, 0x04, 0xc5, 0x71, 0xd2, 0x75, 0x08, 0x69, 0xe3, 0xe6, 0xf7,
	0xf3, 0xfb, 0xbd, 0xf7, 0x7b, 0xcf, 0xbf, 0x04, 0xb8, 0x2c, 0x56, 0x5c, 0x20, 0x9c, 0x11, 0x1f,
	0xf3, 0x2c, 0x23, 0x58, 0x51, 0xce, 0x42, 0x7e, 0xcc, 0x88, 0x08, 0x05, 0xf9, 0x54, 0x10, 0xa9,
	0xbc, 0xa9, 0xe0, 0x8a, 0xc3, 0x7b, 0x2a, 0x25, 0x92, 0xce, 0x18, 0x51, 0xc7, 0x5c, 0x7c, 0xf4,
	0x24, 0x9d, 0x69, 0x1c, 0xf3, 0xcc, 0x6b, 0x6a, 0x6c, 0xde, 0x9d, 0x97, 0x5b, 0x60, 0x6f, 0xae,
	0x27, 0x3c, 0xe1, 0xfa, 0xe8, 0x97, 0x27, 0x83, 0x6e, 0x25, 0x9c, 0x27, 0x19, 0xf1, 0x75, 0x74,
	0x54, 0xc4, 0xbe, 0xa2, 0x39, 0x91, 0x0a, 0xe5, 0xd3, 0x2a, 0x61, 0xfb, 0xbb, 0x05, 0xee, 0x1c,
	0x08, 0x9a, 0x50, 0x36, 0xe1, 0x4c, 0x09, 0x84, 0xd5, 0x6b, 0x24, 0x50, 0x0e, 0xd7, 0xc1, 0x32,
	0x4e, 0x11, 0x65, 0xb6, 0x35, 0xb2, 0xdc, 0x5e, 0x50, 0x05, 0xf0, 0x11, 0xb8, 0x8d, 0x4d, 0x5a,
	0x88, 0xa2, 0x48, 0x10, 0x29, 0xed, 0x25, 0x9d, 0xb0, 0x56, 0xe3, 0xbb, 0x15, 0x0c, 0x1f, 0x80,
	0x61, 0x93, 0xaa, 0xa7, 0xb5, 0xdb, 0x3a, 0x71, 0x50, 0xa3, 0x07, 0x25, 0x08, 0x5f, 0x81, 0xa1,
	0x99, 0x23, 0x24, 0xb3, 0x29, 0x15, 0xc4, 0xee, 0x8c, 0x2c, 0xb7, 0x3f, 0xde, 0xf4, 0x2a, 0xe5,
	0x5e, 0xad, 0xdc, 0x7b, 0x5b, 0x2b, 0xdf, 0x5b, 0x39, 0xfd, 0xb9, 0xd5, 0x3a, 0xf9, 0xb5, 0x65,
	0x05, 0x03, 0xc3, 0xdd, 0xd7, 0xd4, 0xed, 0xaf, 0x1d, 0xb0, 0x31, 0x69, 0x96, 0xac, 0x1b, 0x04,
	0x55, 0x02, 0x1c, 0x82, 0x25, 0x1a, 0xe9, 0x61, 0x3a, 0xc1, 0x12, 0x8d, 0xe0, 0x7d, 0x30, 0x60,
	0xb1, 0x3a, 0xc4, 0x29, 0xc9, 0xd1, 0x84, 0x47, 0xc4, 0x8c, 0xb1, 0x08, 0xc2, 0x0d, 0xd0, 0x95,
	0x34, 0x99, 0x8b, 0x37, 0x51, 0xb9, 0x87, 0xb2, 0x33, 0x15, 0x24, 0x0a, 0x31, 0x67, 0x31, 0x15,
	0xb9, 0xd6, 0xdd, 0x09, 0xd6, 0x6a, 0x7c, 0x52, 0xc1, 0xf0, 0x25, 0xe8, 0x4a, 0x85, 0x54, 0x21,
	0xed, 0xe5, 0x91, 0xe5, 0x0e, 0xc7, 0x63, 0xef, 0x1a, 0xcf, 0xec, 0x19, 0xd9, 0x87, 0x9a, 0x19,
	0x98, 0x0a, 0xf0, 0x21, 0x58, 0xc3, 0x85, 0x10, 0x84, 0xa9, 0xa6, 0x6b, 0x57, 0x77, 0x1d, 0x1a,
	0xb8, 0x6e, 0xea, 0x00, 0x60, 0x12, 0x88, 0x90, 0xf6, 0xad, 0x51, 0xdb, 0xed, 0x05, 0x97, 0x10,
	0x38, 0x01, 0x00, 0x0b, 0x82, 0x14, 0x89, 0x42, 0xa4, 0xec, 0x95, 0x1b, 0x6c, 0xbc, 0x67, 0x78,
	0xbb, 0x0a, 0xee, 0x83, 0xfe, 0x67, 0x94, 0xd1, 0x28, 0x2c, 0x98, 0xa2, 0x99, 0xdd, 0xbb, 0x41,
	0x15, 0xa0, 0x89, 0xef, 0x4a, 0x1e, 0xfc, 0x00, 0x1a, 0x4b, 0x84, 0x94, 0xc5, 0xdc, 0x06, 0xa3,
	0xb6, 0xdb, 0x1f, 0xef, 0x5c, 0x6b, 0x4f, 0x8b, 0xd6, 0x7d, 0xc1, 0x62, 0x1e, 0xac, 0xe2, 0x4b,
	0x51, 0x69, 0xc3, 0xca, 0x57, 0x51, 0x98, 0x12, 0x9a, 0xa4, 0xca, 0xee, 0x8f, 0x2c, 0xb7, 0x1d,
	0x0c, 0x0c, 0xfa, 0x5c, 0x83, 0xdb, 0xdf, 0x2c, 0x00, 0xff, 0xae, 0x05, 0xa7, 0x60, 0xa3, 0xb1,
	0xab, 0xbe, 0x7d, 0x86, 0x14, 0x2a, 0x6f, 0xb4, 0x93, 0xfa, 0xe3, 0xa7, 0xff, 0x21, 0x52, 0x7f,
	0x5f, 0xc1, 0x3f, 0xea, 0x42, 0x08, 0x3a, 0x29, 0x92, 0xa9, 0xb6, 0xe3, 0x6a, 0xa0, 0xcf, 0x57,
	0x5e, 0xb3, 0x7d, 0xf5, 0x35, 0xf7, 0xde, 0x9c, 0x9e, 0x3b, 0xd6, 0xd9, 0xb9, 0x63, 0xfd, 0x3e,
	0x77, 0xac, 0x93, 0x0b, 0xa7, 0x75, 0x76, 0xe1, 0xb4, 0x7e, 0x5c, 0x38, 0xad, 0xf7, 0x3b, 0x09,
	0x55, 0x69, 0x71, 0xe4, 0x61, 0x9e, 0xfb, 0x0b, 0x4a, 0x7d, 0x49, 0x67, 0x8f, 0x6b, 0xa9, 0xfe,
	0xcc, 0x9f, 0xff, 0x55, 0xd4, 0x97, 0x29, 0x91, 0x47, 0x5d, 0x7d, 0xf7, 0xe4, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x2e, 0x7f, 0x03, 0xbe, 0x04, 0x00, 0x00,
}

func (m *OriginContractParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginContractParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginContractParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RequestExpire, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestExpire):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.ContractOwner) > 0 {
		i -= len(m.ContractOwner)
		copy(dAtA[i:], m.ContractOwner)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.ContractOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectionOwnerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectionOwnerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectionOwnerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.ContractInfo) > 0 {
		for iNdEx := len(m.ContractInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ValidUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x42
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.CurrentConfirm != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.CurrentConfirm))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.RequiredConfirm != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.RequiredConfirm))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OriginContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.ContractOriginDataInfo != nil {
		{
			size, err := m.ContractOriginDataInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollectionOwnerRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollectionOwnerRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OriginContractParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.ContractOwner)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestExpire)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	return n
}

func (m *CollectionOwnerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.Id))
	}
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	if m.RequiredConfirm != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.RequiredConfirm))
	}
	if m.Status != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.Status))
	}
	if m.CurrentConfirm != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.CurrentConfirm))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	if len(m.ContractInfo) > 0 {
		for _, e := range m.ContractInfo {
			l = e.Size()
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.ExpiredHeight))
	}
	return n
}

func (m *OriginContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ContractOriginDataInfo != nil {
		l = m.ContractOriginDataInfo.Size()
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	return n
}

func sovCollectionOwnerRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollectionOwnerRequest(x uint64) (n int) {
	return sovCollectionOwnerRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginContractParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: OriginContractParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginContractParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestExpire", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RequestExpire, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func (m *CollectionOwnerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: CollectionOwnerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectionOwnerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredConfirm", wireType)
			}
			m.RequiredConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequiredConfirm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentConfirm", wireType)
			}
			m.CurrentConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentConfirm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ValidUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractInfo = append(m.ContractInfo, &OriginContractInfo{})
			if err := m.ContractInfo[len(m.ContractInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func (m *OriginContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: OriginContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOriginDataInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContractOriginDataInfo == nil {
				m.ContractOriginDataInfo = &OriginContractParam{}
			}
			if err := m.ContractOriginDataInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func skipCollectionOwnerRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollectionOwnerRequest
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
					return 0, ErrIntOverflowCollectionOwnerRequest
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
					return 0, ErrIntOverflowCollectionOwnerRequest
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
				return 0, ErrInvalidLengthCollectionOwnerRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollectionOwnerRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollectionOwnerRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollectionOwnerRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollectionOwnerRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollectionOwnerRequest = fmt.Errorf("proto: unexpected end of group")
)
