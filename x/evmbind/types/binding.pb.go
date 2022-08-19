// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmbind/binding.proto

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

type Binding struct {
	EthAddress   string `protobuf:"bytes,1,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
	EthSignature string `protobuf:"bytes,2,opt,name=ethSignature,proto3" json:"ethSignature,omitempty"`
	SignMessage  string `protobuf:"bytes,3,opt,name=signMessage,proto3" json:"signMessage,omitempty"`
	Creator      string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Binding) Reset()         { *m = Binding{} }
func (m *Binding) String() string { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()    {}
func (*Binding) Descriptor() ([]byte, []int) {
	return fileDescriptor_13a6e2b67d765777, []int{0}
}
func (m *Binding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Binding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Binding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Binding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binding.Merge(m, src)
}
func (m *Binding) XXX_Size() int {
	return m.Size()
}
func (m *Binding) XXX_DiscardUnknown() {
	xxx_messageInfo_Binding.DiscardUnknown(m)
}

var xxx_messageInfo_Binding proto.InternalMessageInfo

func (m *Binding) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *Binding) GetEthSignature() string {
	if m != nil {
		return m.EthSignature
	}
	return ""
}

func (m *Binding) GetSignMessage() string {
	if m != nil {
		return m.SignMessage
	}
	return ""
}

func (m *Binding) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Binding)(nil), "thesixnetwork.sixprotocol.evmbind.Binding")
}

func init() { proto.RegisterFile("evmbind/binding.proto", fileDescriptor_13a6e2b67d765777) }

var fileDescriptor_13a6e2b67d765777 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2d, 0xcb, 0x4d,
	0xca, 0xcc, 0x4b, 0xd1, 0x07, 0x11, 0x99, 0x79, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x8a, 0x25, 0x19, 0xa9, 0xc5, 0x99, 0x15, 0x79, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0x7a, 0xc5,
	0x99, 0x15, 0x60, 0xf1, 0xe4, 0xfc, 0x1c, 0x3d, 0xa8, 0x06, 0xa5, 0x4e, 0x46, 0x2e, 0x76, 0x27,
	0x88, 0x26, 0x21, 0x39, 0x2e, 0xae, 0xd4, 0x92, 0x0c, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24, 0x11, 0x21, 0x25, 0x2e, 0x9e, 0xd4, 0x92, 0x8c,
	0xe0, 0xcc, 0xf4, 0xbc, 0xc4, 0x92, 0xd2, 0xa2, 0x54, 0x09, 0x26, 0xb0, 0x0a, 0x14, 0x31, 0x21,
	0x05, 0x2e, 0xee, 0xe2, 0xcc, 0xf4, 0x3c, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x66,
	0xb0, 0x12, 0x64, 0x21, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09,
	0x16, 0xb0, 0x2c, 0x8c, 0xeb, 0xe4, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0xa6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0x7e, 0xd2,
	0x2f, 0xce, 0xac, 0xd0, 0x85, 0x79, 0x4a, 0xbf, 0x42, 0x1f, 0x16, 0x0e, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x59, 0xeb, 0x67,
	0x1f, 0x01, 0x00, 0x00,
}

func (m *Binding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Binding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Binding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBinding(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SignMessage) > 0 {
		i -= len(m.SignMessage)
		copy(dAtA[i:], m.SignMessage)
		i = encodeVarintBinding(dAtA, i, uint64(len(m.SignMessage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthSignature) > 0 {
		i -= len(m.EthSignature)
		copy(dAtA[i:], m.EthSignature)
		i = encodeVarintBinding(dAtA, i, uint64(len(m.EthSignature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintBinding(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBinding(dAtA []byte, offset int, v uint64) int {
	offset -= sovBinding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Binding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovBinding(uint64(l))
	}
	l = len(m.EthSignature)
	if l > 0 {
		n += 1 + l + sovBinding(uint64(l))
	}
	l = len(m.SignMessage)
	if l > 0 {
		n += 1 + l + sovBinding(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBinding(uint64(l))
	}
	return n
}

func sovBinding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBinding(x uint64) (n int) {
	return sovBinding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Binding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinding
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
			return fmt.Errorf("proto: Binding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Binding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinding
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
				return ErrInvalidLengthBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinding
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
				return ErrInvalidLengthBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinding
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
				return ErrInvalidLengthBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinding
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
				return ErrInvalidLengthBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinding
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
func skipBinding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBinding
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
					return 0, ErrIntOverflowBinding
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
					return 0, ErrIntOverflowBinding
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
				return 0, ErrInvalidLengthBinding
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBinding
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBinding
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBinding        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBinding          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBinding = fmt.Errorf("proto: unexpected end of group")
)
