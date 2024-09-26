// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/v2/options.proto

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

type Options struct {
	DefaultMintee string `protobuf:"bytes,1,opt,name=defaultMintee,proto3" json:"defaultMintee,omitempty"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_d336aa4e6eee3602, []int{0}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Options.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return m.Size()
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetDefaultMintee() string {
	if m != nil {
		return m.DefaultMintee
	}
	return ""
}

func init() {
	proto.RegisterType((*Options)(nil), "thesixnetwork.sixprotocol.tokenmngr.v2.Options")
}

func init() { proto.RegisterFile("tokenmngr/v2/options.proto", fileDescriptor_d336aa4e6eee3602) }

var fileDescriptor_d336aa4e6eee3602 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc9, 0xcf, 0x4e,
	0xcd, 0xcb, 0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xd2, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2b, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b,
	0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0xc1,
	0x75, 0xe9, 0x95, 0x19, 0x29, 0xe9, 0x73, 0xb1, 0xfb, 0x43, 0x34, 0x0a, 0xa9, 0x70, 0xf1, 0xa6,
	0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0xf8, 0x66, 0xe6, 0x95, 0xa4, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x3a, 0x05, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x8a, 0xed,
	0xfa, 0xc5, 0x99, 0x15, 0xba, 0x30, 0xeb, 0xf5, 0x2b, 0xf4, 0x11, 0xce, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0xcb, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x53, 0x2d, 0x01, 0xc5,
	0xd0, 0x00, 0x00, 0x00,
}

func (m *Options) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Options) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Options) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DefaultMintee) > 0 {
		i -= len(m.DefaultMintee)
		copy(dAtA[i:], m.DefaultMintee)
		i = encodeVarintOptions(dAtA, i, uint64(len(m.DefaultMintee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOptions(dAtA []byte, offset int, v uint64) int {
	offset -= sovOptions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Options) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DefaultMintee)
	if l > 0 {
		n += 1 + l + sovOptions(uint64(l))
	}
	return n
}

func sovOptions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOptions(x uint64) (n int) {
	return sovOptions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Options) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOptions
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
			return fmt.Errorf("proto: Options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMintee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOptions
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
				return ErrInvalidLengthOptions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOptions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultMintee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOptions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOptions
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
func skipOptions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOptions
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
					return 0, ErrIntOverflowOptions
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
					return 0, ErrIntOverflowOptions
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
				return 0, ErrInvalidLengthOptions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOptions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOptions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOptions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOptions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOptions = fmt.Errorf("proto: unexpected end of group")
)
