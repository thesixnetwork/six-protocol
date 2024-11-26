// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v094/nft_fee_config.proto

package v094

import (
	encoding_binary "encoding/binary"
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

type FeeDistributionMethod int32

const (
	FeeDistributionMethod_BURN        FeeDistributionMethod = 0
	FeeDistributionMethod_REWARD_POOL FeeDistributionMethod = 1
	FeeDistributionMethod_TRANSFER    FeeDistributionMethod = 2
)

var FeeDistributionMethod_name = map[int32]string{
	0: "BURN",
	1: "REWARD_POOL",
	2: "TRANSFER",
}

var FeeDistributionMethod_value = map[string]int32{
	"BURN":        0,
	"REWARD_POOL": 1,
	"TRANSFER":    2,
}

func (x FeeDistributionMethod) String() string {
	return proto.EnumName(FeeDistributionMethod_name, int32(x))
}

func (FeeDistributionMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ba2cda3bf4b38ef2, []int{0}
}

type FeeDistribution struct {
	Method  FeeDistributionMethod `protobuf:"varint,1,opt,name=method,proto3,enum=thesixnetwork.sixprotocol.nftmngr.v094.FeeDistributionMethod" json:"method,omitempty"`
	Portion float32               `protobuf:"fixed32,2,opt,name=portion,proto3" json:"portion,omitempty"`
}

func (m *FeeDistribution) Reset()         { *m = FeeDistribution{} }
func (m *FeeDistribution) String() string { return proto.CompactTextString(m) }
func (*FeeDistribution) ProtoMessage()    {}
func (*FeeDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba2cda3bf4b38ef2, []int{0}
}
func (m *FeeDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeDistribution.Merge(m, src)
}
func (m *FeeDistribution) XXX_Size() int {
	return m.Size()
}
func (m *FeeDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_FeeDistribution proto.InternalMessageInfo

func (m *FeeDistribution) GetMethod() FeeDistributionMethod {
	if m != nil {
		return m.Method
	}
	return FeeDistributionMethod_BURN
}

func (m *FeeDistribution) GetPortion() float32 {
	if m != nil {
		return m.Portion
	}
	return 0
}

type FeeConfig struct {
	FeeAmount        string             `protobuf:"bytes,1,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`
	FeeDistributions []*FeeDistribution `protobuf:"bytes,2,rep,name=fee_distributions,json=feeDistributions,proto3" json:"fee_distributions,omitempty"`
}

func (m *FeeConfig) Reset()         { *m = FeeConfig{} }
func (m *FeeConfig) String() string { return proto.CompactTextString(m) }
func (*FeeConfig) ProtoMessage()    {}
func (*FeeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba2cda3bf4b38ef2, []int{1}
}
func (m *FeeConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeConfig.Merge(m, src)
}
func (m *FeeConfig) XXX_Size() int {
	return m.Size()
}
func (m *FeeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FeeConfig proto.InternalMessageInfo

func (m *FeeConfig) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

func (m *FeeConfig) GetFeeDistributions() []*FeeDistribution {
	if m != nil {
		return m.FeeDistributions
	}
	return nil
}

type NFTFeeConfig struct {
	SchemaFee *FeeConfig `protobuf:"bytes,1,opt,name=schema_fee,json=schemaFee,proto3" json:"schema_fee,omitempty"`
}

func (m *NFTFeeConfig) Reset()         { *m = NFTFeeConfig{} }
func (m *NFTFeeConfig) String() string { return proto.CompactTextString(m) }
func (*NFTFeeConfig) ProtoMessage()    {}
func (*NFTFeeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba2cda3bf4b38ef2, []int{2}
}
func (m *NFTFeeConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTFeeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTFeeConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTFeeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTFeeConfig.Merge(m, src)
}
func (m *NFTFeeConfig) XXX_Size() int {
	return m.Size()
}
func (m *NFTFeeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTFeeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NFTFeeConfig proto.InternalMessageInfo

func (m *NFTFeeConfig) GetSchemaFee() *FeeConfig {
	if m != nil {
		return m.SchemaFee
	}
	return nil
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixprotocol.nftmngr.v094.FeeDistributionMethod", FeeDistributionMethod_name, FeeDistributionMethod_value)
	proto.RegisterType((*FeeDistribution)(nil), "thesixnetwork.sixprotocol.nftmngr.v094.FeeDistribution")
	proto.RegisterType((*FeeConfig)(nil), "thesixnetwork.sixprotocol.nftmngr.v094.FeeConfig")
	proto.RegisterType((*NFTFeeConfig)(nil), "thesixnetwork.sixprotocol.nftmngr.v094.NFTFeeConfig")
}

func init() { proto.RegisterFile("nftmngr/v094/nft_fee_config.proto", fileDescriptor_ba2cda3bf4b38ef2) }

var fileDescriptor_ba2cda3bf4b38ef2 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0xbd, 0x37, 0x5c, 0x3a, 0x90, 0x4b, 0x9d, 0xc4, 0xa4, 0x1b, 0x1b, 0x64, 0x61,
	0x88, 0x89, 0xad, 0xa2, 0x89, 0xd1, 0xc4, 0x44, 0x10, 0xba, 0x52, 0x20, 0x23, 0xc4, 0xc4, 0x0d,
	0x42, 0x99, 0xd2, 0x89, 0xb6, 0x43, 0x3a, 0x83, 0xd6, 0xad, 0x4f, 0xc0, 0x63, 0xb9, 0x64, 0xe9,
	0xd2, 0xc0, 0x8b, 0x98, 0x0e, 0x10, 0xc5, 0xb8, 0x50, 0x97, 0x73, 0xfe, 0xcc, 0x77, 0xbe, 0x73,
	0x72, 0xe0, 0x66, 0xe8, 0x89, 0x20, 0x1c, 0x44, 0xf6, 0xfd, 0xee, 0xd1, 0x81, 0x1d, 0x7a, 0xa2,
	0xe3, 0x11, 0xd2, 0x71, 0x59, 0xe8, 0xd1, 0x81, 0x35, 0x8c, 0x98, 0x60, 0x68, 0x4b, 0xf8, 0x84,
	0xd3, 0x38, 0x24, 0xe2, 0x81, 0x45, 0xb7, 0x16, 0xa7, 0xb1, 0xac, 0xbb, 0xec, 0xce, 0x5a, 0x7c,
	0xb6, 0x92, 0xcf, 0x85, 0x27, 0x00, 0x73, 0x0e, 0x21, 0x55, 0xca, 0x45, 0x44, 0x7b, 0x23, 0x41,
	0x59, 0x88, 0xda, 0x30, 0x15, 0x10, 0xe1, 0xb3, 0xbe, 0x01, 0xf2, 0xa0, 0xf8, 0xbf, 0x74, 0x62,
	0x7d, 0x0f, 0x66, 0x7d, 0x02, 0x5d, 0x48, 0x08, 0x5e, 0xc0, 0x90, 0x01, 0xff, 0x0d, 0x59, 0x94,
	0x04, 0x86, 0x9a, 0x07, 0x45, 0x15, 0x2f, 0x9f, 0x85, 0x31, 0x80, 0x9a, 0x43, 0xc8, 0x99, 0x1c,
	0x00, 0x6d, 0x40, 0x98, 0x8c, 0xd3, 0x0d, 0xd8, 0x28, 0x14, 0x52, 0x41, 0xc3, 0x9a, 0x47, 0x48,
	0x59, 0x16, 0x50, 0x1f, 0xae, 0x25, 0x71, 0xff, 0x43, 0x23, 0x6e, 0xa8, 0xf9, 0x3f, 0xc5, 0x4c,
	0xe9, 0xf0, 0x97, 0xa2, 0x58, 0xf7, 0x56, 0x0b, 0xbc, 0x70, 0x03, 0xb3, 0x75, 0xa7, 0xf5, 0x2e,
	0xd5, 0x84, 0x90, 0xbb, 0x3e, 0x09, 0xba, 0xc9, 0xaa, 0xa5, 0x54, 0xa6, 0xb4, 0xf7, 0x83, 0x76,
	0x73, 0x0c, 0xd6, 0xe6, 0x10, 0x87, 0x90, 0xed, 0x53, 0xb8, 0xfe, 0xe5, 0xbe, 0x50, 0x1a, 0xfe,
	0xad, 0xb4, 0x71, 0x5d, 0x57, 0x50, 0x0e, 0x66, 0x70, 0xed, 0xaa, 0x8c, 0xab, 0x9d, 0x66, 0xa3,
	0x71, 0xae, 0x03, 0x94, 0x85, 0xe9, 0x16, 0x2e, 0xd7, 0x2f, 0x9d, 0x1a, 0xd6, 0xd5, 0x4a, 0xeb,
	0x79, 0x6a, 0x82, 0xc9, 0xd4, 0x04, 0xaf, 0x53, 0x13, 0x8c, 0x67, 0xa6, 0x32, 0x99, 0x99, 0xca,
	0xcb, 0xcc, 0x54, 0xae, 0x8f, 0x07, 0x54, 0xf8, 0xa3, 0x9e, 0xe5, 0xb2, 0xc0, 0x5e, 0x71, 0xb4,
	0x39, 0x8d, 0x77, 0x96, 0x92, 0x76, 0x6c, 0x2f, 0x0f, 0x49, 0x3c, 0x0e, 0x09, 0x97, 0xe7, 0xd4,
	0x4b, 0xc9, 0x78, 0xff, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x08, 0x2c, 0x52, 0x98, 0x65, 0x02, 0x00,
	0x00,
}

func (m *FeeDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Portion != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Portion))))
		i--
		dAtA[i] = 0x15
	}
	if m.Method != 0 {
		i = encodeVarintNftFeeConfig(dAtA, i, uint64(m.Method))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FeeConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeDistributions) > 0 {
		for iNdEx := len(m.FeeDistributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeDistributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftFeeConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.FeeAmount) > 0 {
		i -= len(m.FeeAmount)
		copy(dAtA[i:], m.FeeAmount)
		i = encodeVarintNftFeeConfig(dAtA, i, uint64(len(m.FeeAmount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTFeeConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTFeeConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTFeeConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaFee != nil {
		{
			size, err := m.SchemaFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftFeeConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftFeeConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftFeeConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovNftFeeConfig(uint64(m.Method))
	}
	if m.Portion != 0 {
		n += 5
	}
	return n
}

func (m *FeeConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeeAmount)
	if l > 0 {
		n += 1 + l + sovNftFeeConfig(uint64(l))
	}
	if len(m.FeeDistributions) > 0 {
		for _, e := range m.FeeDistributions {
			l = e.Size()
			n += 1 + l + sovNftFeeConfig(uint64(l))
		}
	}
	return n
}

func (m *NFTFeeConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SchemaFee != nil {
		l = m.SchemaFee.Size()
		n += 1 + l + sovNftFeeConfig(uint64(l))
	}
	return n
}

func sovNftFeeConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftFeeConfig(x uint64) (n int) {
	return sovNftFeeConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftFeeConfig
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
			return fmt.Errorf("proto: FeeDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftFeeConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= FeeDistributionMethod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Portion", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Portion = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipNftFeeConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftFeeConfig
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
func (m *FeeConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftFeeConfig
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
			return fmt.Errorf("proto: FeeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftFeeConfig
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
				return ErrInvalidLengthNftFeeConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftFeeConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDistributions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftFeeConfig
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
				return ErrInvalidLengthNftFeeConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftFeeConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeDistributions = append(m.FeeDistributions, &FeeDistribution{})
			if err := m.FeeDistributions[len(m.FeeDistributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftFeeConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftFeeConfig
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
func (m *NFTFeeConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftFeeConfig
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
			return fmt.Errorf("proto: NFTFeeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTFeeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftFeeConfig
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
				return ErrInvalidLengthNftFeeConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftFeeConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SchemaFee == nil {
				m.SchemaFee = &FeeConfig{}
			}
			if err := m.SchemaFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftFeeConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftFeeConfig
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
func skipNftFeeConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftFeeConfig
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
					return 0, ErrIntOverflowNftFeeConfig
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
					return 0, ErrIntOverflowNftFeeConfig
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
				return 0, ErrInvalidLengthNftFeeConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftFeeConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftFeeConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftFeeConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftFeeConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftFeeConfig = fmt.Errorf("proto: unexpected end of group")
)
