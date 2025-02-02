// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/lock_schema_fee.proto

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

type LockSchemaFee struct {
	Id                string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VirtualSchemaCode string     `protobuf:"bytes,2,opt,name=virtualSchemaCode,proto3" json:"virtualSchemaCode,omitempty"`
	Amount            types.Coin `protobuf:"bytes,3,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	Proposer          string     `protobuf:"bytes,4,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (m *LockSchemaFee) Reset()         { *m = LockSchemaFee{} }
func (m *LockSchemaFee) String() string { return proto.CompactTextString(m) }
func (*LockSchemaFee) ProtoMessage()    {}
func (*LockSchemaFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_f27fff5cf2b0bbaf, []int{0}
}

func (m *LockSchemaFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LockSchemaFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockSchemaFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LockSchemaFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockSchemaFee.Merge(m, src)
}

func (m *LockSchemaFee) XXX_Size() int {
	return m.Size()
}

func (m *LockSchemaFee) XXX_DiscardUnknown() {
	xxx_messageInfo_LockSchemaFee.DiscardUnknown(m)
}

var xxx_messageInfo_LockSchemaFee proto.InternalMessageInfo

func (m *LockSchemaFee) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LockSchemaFee) GetVirtualSchemaCode() string {
	if m != nil {
		return m.VirtualSchemaCode
	}
	return ""
}

func (m *LockSchemaFee) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *LockSchemaFee) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func init() {
	proto.RegisterType((*LockSchemaFee)(nil), "thesixnetwork.sixprotocol.nftmngr.LockSchemaFee")
}

func init() { proto.RegisterFile("nftmngr/lock_schema_fee.proto", fileDescriptor_f27fff5cf2b0bbaf) }

var fileDescriptor_f27fff5cf2b0bbaf = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x9b, 0xfd, 0xff, 0x0c, 0xad, 0x28, 0x58, 0x3c, 0xd4, 0x81, 0xd9, 0xf4, 0xe2, 0x0e,
	0x2e, 0x61, 0x8a, 0x5f, 0x60, 0x03, 0x4f, 0x82, 0x30, 0x6f, 0x5e, 0x46, 0x9b, 0xbe, 0xeb, 0x42,
	0xd7, 0xbe, 0x25, 0xc9, 0x66, 0xfd, 0x16, 0x7e, 0x0e, 0x3f, 0xc9, 0x2e, 0xc2, 0x8e, 0x9e, 0x54,
	0xd6, 0x2f, 0x22, 0x4b, 0x3b, 0x51, 0x3c, 0x25, 0x79, 0x9f, 0x3c, 0x79, 0x7e, 0x3c, 0x71, 0x4f,
	0xb2, 0x89, 0x49, 0xb3, 0x58, 0xf1, 0x19, 0x8a, 0x64, 0xac, 0xc5, 0x14, 0xd2, 0x60, 0x3c, 0x01,
	0x60, 0xb9, 0x42, 0x83, 0xde, 0xa9, 0x99, 0x82, 0x96, 0x45, 0x06, 0xe6, 0x11, 0x55, 0xc2, 0xb4,
	0x2c, 0xec, 0x5c, 0xe0, 0x8c, 0xd5, 0xc6, 0xd6, 0x51, 0x8c, 0x31, 0xda, 0x29, 0xdf, 0xec, 0x2a,
	0x63, 0x8b, 0x0a, 0xd4, 0x29, 0x6a, 0x1e, 0x06, 0x1a, 0xf8, 0xa2, 0x1f, 0x82, 0x09, 0xfa, 0x5c,
	0xa0, 0xcc, 0x2a, 0xfd, 0xec, 0x95, 0xb8, 0xfb, 0xb7, 0x28, 0x92, 0x7b, 0x9b, 0x78, 0x03, 0xe0,
	0x1d, 0xb8, 0x0d, 0x19, 0xf9, 0xa4, 0x43, 0xba, 0xbb, 0xa3, 0x86, 0x8c, 0xbc, 0x0b, 0xf7, 0x70,
	0x21, 0x95, 0x99, 0x07, 0xb3, 0xea, 0xce, 0x10, 0x23, 0xf0, 0x1b, 0x56, 0xfe, 0x2b, 0x78, 0xa1,
	0xdb, 0x0c, 0x52, 0x9c, 0x67, 0xc6, 0xff, 0xd7, 0x21, 0xdd, 0xbd, 0xcb, 0x63, 0x56, 0x01, 0xb0,
	0x0d, 0x00, 0xab, 0x01, 0xd8, 0x10, 0x65, 0x36, 0xe0, 0xcb, 0xf7, 0xb6, 0xf3, 0xf2, 0xd1, 0x3e,
	0x8f, 0xa5, 0x99, 0xce, 0x43, 0x26, 0x30, 0xe5, 0x35, 0x6d, 0xb5, 0xf4, 0x74, 0x94, 0x70, 0xf3,
	0x94, 0x83, 0xb6, 0x86, 0x51, 0xfd, 0xb2, 0xd7, 0x72, 0x77, 0x72, 0x85, 0x39, 0x6a, 0x50, 0xfe,
	0x7f, 0x0b, 0xf2, 0x7d, 0x1e, 0xdc, 0x2d, 0xd7, 0x94, 0xac, 0xd6, 0x94, 0x7c, 0xae, 0x29, 0x79,
	0x2e, 0xa9, 0xb3, 0x2a, 0xa9, 0xf3, 0x56, 0x52, 0xe7, 0xe1, 0xfa, 0x47, 0xcc, 0xaf, 0x36, 0xb9,
	0x96, 0x45, 0x6f, 0x5b, 0x27, 0x2f, 0xf8, 0xf6, 0x27, 0x6c, 0x72, 0xd8, 0xb4, 0xca, 0xd5, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xc4, 0x25, 0x13, 0xa1, 0x01, 0x00, 0x00,
}

func (m *LockSchemaFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockSchemaFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockSchemaFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintLockSchemaFee(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLockSchemaFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.VirtualSchemaCode) > 0 {
		i -= len(m.VirtualSchemaCode)
		copy(dAtA[i:], m.VirtualSchemaCode)
		i = encodeVarintLockSchemaFee(dAtA, i, uint64(len(m.VirtualSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintLockSchemaFee(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLockSchemaFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovLockSchemaFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *LockSchemaFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovLockSchemaFee(uint64(l))
	}
	l = len(m.VirtualSchemaCode)
	if l > 0 {
		n += 1 + l + sovLockSchemaFee(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLockSchemaFee(uint64(l))
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovLockSchemaFee(uint64(l))
	}
	return n
}

func sovLockSchemaFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozLockSchemaFee(x uint64) (n int) {
	return sovLockSchemaFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *LockSchemaFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockSchemaFee
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
			return fmt.Errorf("proto: LockSchemaFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockSchemaFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockSchemaFee
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
				return ErrInvalidLengthLockSchemaFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLockSchemaFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockSchemaFee
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
				return ErrInvalidLengthLockSchemaFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLockSchemaFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockSchemaFee
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
				return ErrInvalidLengthLockSchemaFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockSchemaFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockSchemaFee
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
				return ErrInvalidLengthLockSchemaFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLockSchemaFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockSchemaFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockSchemaFee
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

func skipLockSchemaFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLockSchemaFee
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
					return 0, ErrIntOverflowLockSchemaFee
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
					return 0, ErrIntOverflowLockSchemaFee
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
				return 0, ErrInvalidLengthLockSchemaFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLockSchemaFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLockSchemaFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLockSchemaFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLockSchemaFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLockSchemaFee = fmt.Errorf("proto: unexpected end of group")
)
