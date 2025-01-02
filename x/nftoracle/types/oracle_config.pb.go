// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/oracle_config.proto

package types

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

type OracleConfig struct {
	MinimumConfirmation int32 `protobuf:"varint,1,opt,name=minimum_confirmation,json=minimumConfirmation,proto3" json:"minimum_confirmation,omitempty"`
}

func (m *OracleConfig) Reset()         { *m = OracleConfig{} }
func (m *OracleConfig) String() string { return proto.CompactTextString(m) }
func (*OracleConfig) ProtoMessage()    {}
func (*OracleConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_da61cb4fbc7889cc, []int{0}
}
func (m *OracleConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleConfig.Merge(m, src)
}
func (m *OracleConfig) XXX_Size() int {
	return m.Size()
}
func (m *OracleConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OracleConfig proto.InternalMessageInfo

func (m *OracleConfig) GetMinimumConfirmation() int32 {
	if m != nil {
		return m.MinimumConfirmation
	}
	return 0
}

func init() {
	proto.RegisterType((*OracleConfig)(nil), "thesixnetwork.sixprotocol.nftoracle.OracleConfig")
}

func init() { proto.RegisterFile("nftoracle/oracle_config.proto", fileDescriptor_da61cb4fbc7889cc) }

var fileDescriptor_da61cb4fbc7889cc = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x4b, 0x2b, 0xc9,
	0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x87, 0x50, 0xf1, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xca, 0x25, 0x19, 0xa9, 0xc5, 0x99, 0x15, 0x79, 0xa9, 0x25,
	0xe5, 0xf9, 0x45, 0xd9, 0x7a, 0xc5, 0x99, 0x15, 0x60, 0xf1, 0xe4, 0xfc, 0x1c, 0x3d, 0xb8, 0x46,
	0x25, 0x47, 0x2e, 0x1e, 0x7f, 0x30, 0xcb, 0x19, 0xac, 0x55, 0xc8, 0x90, 0x4b, 0x24, 0x37, 0x33,
	0x2f, 0x33, 0xb7, 0x34, 0x17, 0x62, 0x58, 0x51, 0x6e, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x6b, 0x90, 0x30, 0x54, 0xce, 0x19, 0x49, 0xca, 0x29, 0xf0, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xcc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x51, 0x1c, 0xa3, 0x5f, 0x9c, 0x59, 0xa1, 0x0b, 0x73, 0x8d, 0x7e, 0x85, 0x3e,
	0xc2, 0x23, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x39, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x66, 0xd3, 0x0f, 0x17, 0xe2, 0x00, 0x00, 0x00,
}

func (m *OracleConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinimumConfirmation != 0 {
		i = encodeVarintOracleConfig(dAtA, i, uint64(m.MinimumConfirmation))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracleConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracleConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinimumConfirmation != 0 {
		n += 1 + sovOracleConfig(uint64(m.MinimumConfirmation))
	}
	return n
}

func sovOracleConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracleConfig(x uint64) (n int) {
	return sovOracleConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracleConfig
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
			return fmt.Errorf("proto: OracleConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumConfirmation", wireType)
			}
			m.MinimumConfirmation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumConfirmation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracleConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracleConfig
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
func skipOracleConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracleConfig
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
					return 0, ErrIntOverflowOracleConfig
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
					return 0, ErrIntOverflowOracleConfig
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
				return 0, ErrInvalidLengthOracleConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracleConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracleConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracleConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracleConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracleConfig = fmt.Errorf("proto: unexpected end of group")
)
