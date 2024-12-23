// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/token_burn.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TokenBurn struct {
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *TokenBurn) Reset()         { *m = TokenBurn{} }
func (m *TokenBurn) String() string { return proto.CompactTextString(m) }
func (*TokenBurn) ProtoMessage()    {}
func (*TokenBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_658df29e28c1c287, []int{0}
}

func (m *TokenBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *TokenBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *TokenBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenBurn.Merge(m, src)
}

func (m *TokenBurn) XXX_Size() int {
	return m.Size()
}

func (m *TokenBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenBurn.DiscardUnknown(m)
}

var xxx_messageInfo_TokenBurn proto.InternalMessageInfo

func (m *TokenBurn) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*TokenBurn)(nil), "thesixnetwork.sixprotocol.tokenmngr.TokenBurn")
}

func init() { proto.RegisterFile("tokenmngr/token_burn.proto", fileDescriptor_658df29e28c1c287) }

var fileDescriptor_658df29e28c1c287 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0xa5, 0x12, 0x61, 0xab, 0x18, 0x20, 0x83, 0x8b, 0x60, 0x80, 0xa5, 0x7e, 0x2a,
	0x0c, 0xec, 0xe1, 0x04, 0x20, 0x26, 0x16, 0x14, 0x07, 0x2b, 0xb5, 0x42, 0xfc, 0x2a, 0xbf, 0x17,
	0x08, 0xb7, 0xe0, 0x1c, 0x9c, 0xa4, 0x63, 0x47, 0x26, 0x40, 0xc9, 0x45, 0x50, 0x9c, 0x80, 0xc2,
	0xe4, 0x5f, 0xbf, 0xf5, 0x7d, 0xd6, 0xef, 0x38, 0x61, 0x2c, 0x8d, 0xab, 0x5c, 0xe1, 0x21, 0xa4,
	0x07, 0x5d, 0x7b, 0xa7, 0x36, 0x1e, 0x19, 0xe7, 0xa7, 0xbc, 0x36, 0x64, 0x1b, 0x67, 0xf8, 0x05,
	0x7d, 0xa9, 0xc8, 0x36, 0xa1, 0xcf, 0xf1, 0x49, 0xfd, 0x51, 0xc9, 0x41, 0x81, 0x05, 0x86, 0x1e,
	0xfa, 0x34, 0xa0, 0x89, 0xcc, 0x91, 0x2a, 0x24, 0xd0, 0x19, 0x19, 0x78, 0x5e, 0x69, 0xc3, 0xd9,
	0x0a, 0x72, 0xb4, 0xa3, 0xfa, 0x04, 0xe3, 0xbd, 0xbb, 0x5e, 0x91, 0xd6, 0xde, 0xcd, 0x75, 0x3c,
	0xcb, 0x2a, 0xac, 0x1d, 0x1f, 0x8a, 0x63, 0x71, 0xbe, 0x7f, 0x71, 0xa4, 0x06, 0x5a, 0xf5, 0xb4,
	0x1a, 0x69, 0x75, 0x8d, 0xd6, 0xa5, 0xb0, 0xfd, 0x5c, 0x44, 0xef, 0x5f, 0x8b, 0xb3, 0xc2, 0xf2,
	0xba, 0xd6, 0x2a, 0xc7, 0x0a, 0xc6, 0xa7, 0x86, 0x63, 0x49, 0x8f, 0x25, 0xf0, 0xeb, 0xc6, 0x50,
	0x00, 0x6e, 0x47, 0x73, 0x7a, 0xb3, 0x6d, 0xa5, 0xd8, 0xb5, 0x52, 0x7c, 0xb7, 0x52, 0xbc, 0x75,
	0x32, 0xda, 0x75, 0x32, 0xfa, 0xe8, 0x64, 0x74, 0x7f, 0x35, 0x51, 0xfd, 0x1b, 0x0c, 0x64, 0x9b,
	0xe5, 0xef, 0x62, 0x68, 0x60, 0xf2, 0x53, 0xbd, 0x5f, 0xcf, 0xc2, 0xdd, 0xe5, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x29, 0x95, 0xe6, 0x9b, 0x43, 0x01, 0x00, 0x00,
}

func (m *TokenBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTokenBurn(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTokenBurn(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenBurn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *TokenBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovTokenBurn(uint64(l))
	return n
}

func sovTokenBurn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTokenBurn(x uint64) (n int) {
	return sovTokenBurn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *TokenBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenBurn
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
			return fmt.Errorf("proto: TokenBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBurn
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
				return ErrInvalidLengthTokenBurn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenBurn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenBurn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenBurn
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

func skipTokenBurn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenBurn
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
					return 0, ErrIntOverflowTokenBurn
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
					return 0, ErrIntOverflowTokenBurn
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
				return 0, ErrInvalidLengthTokenBurn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenBurn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenBurn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenBurn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenBurn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenBurn = fmt.Errorf("proto: unexpected end of group")
)
