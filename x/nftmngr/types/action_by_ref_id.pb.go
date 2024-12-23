// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/action_by_ref_id.proto

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

type ActionByRefId struct {
	RefId         string `protobuf:"bytes,1,opt,name=refId,proto3" json:"refId,omitempty"`
	Creator       string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	NftSchemaCode string `protobuf:"bytes,3,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	TokenId       string `protobuf:"bytes,4,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Action        string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
}

func (m *ActionByRefId) Reset()         { *m = ActionByRefId{} }
func (m *ActionByRefId) String() string { return proto.CompactTextString(m) }
func (*ActionByRefId) ProtoMessage()    {}
func (*ActionByRefId) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf414553fb0f05e, []int{0}
}
func (m *ActionByRefId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionByRefId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionByRefId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionByRefId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionByRefId.Merge(m, src)
}
func (m *ActionByRefId) XXX_Size() int {
	return m.Size()
}
func (m *ActionByRefId) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionByRefId.DiscardUnknown(m)
}

var xxx_messageInfo_ActionByRefId proto.InternalMessageInfo

func (m *ActionByRefId) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *ActionByRefId) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ActionByRefId) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionByRefId) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *ActionByRefId) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionByRefId)(nil), "thesixnetwork.sixnft.nftmngr.ActionByRefId")
}

func init() { proto.RegisterFile("nftmngr/action_by_ref_id.proto", fileDescriptor_8bf414553fb0f05e) }

var fileDescriptor_8bf414553fb0f05e = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0xaa, 0x8c, 0x2f, 0x4a,
	0x4d, 0x8b, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xc9, 0x48, 0x2d,
	0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0x31, 0xd3, 0x4a, 0xf4, 0xa0,
	0x9a, 0x94, 0xa6, 0x33, 0x72, 0xf1, 0x3a, 0x82, 0x35, 0x3a, 0x55, 0x06, 0xa5, 0xa6, 0x79, 0xa6,
	0x08, 0x89, 0x70, 0xb1, 0x16, 0x81, 0x18, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e,
	0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x13, 0x58, 0x1c, 0xc6,
	0x15, 0x52, 0xe1, 0xe2, 0xcd, 0x4b, 0x2b, 0x09, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x74, 0xce, 0x4f,
	0x49, 0x95, 0x60, 0x06, 0xcb, 0xa3, 0x0a, 0x82, 0xf4, 0x97, 0xe4, 0x67, 0xa7, 0xe6, 0x79, 0xa6,
	0x48, 0xb0, 0x40, 0xf4, 0x43, 0xb9, 0x42, 0x62, 0x5c, 0x6c, 0x10, 0x97, 0x4b, 0xb0, 0x82, 0x25,
	0xa0, 0x3c, 0x27, 0xff, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xf1, 0x9c, 0x7e, 0x71, 0x66,
	0x85, 0x2e, 0xd8, 0xd7, 0xc9, 0xf9, 0x39, 0xfa, 0x15, 0xfa, 0xb0, 0x90, 0x29, 0xa9, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0xcb, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x34, 0x94, 0x1b,
	0x31, 0x01, 0x00, 0x00,
}

func (m *ActionByRefId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionByRefId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionByRefId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintActionByRefId(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintActionByRefId(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionByRefId(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintActionByRefId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RefId) > 0 {
		i -= len(m.RefId)
		copy(dAtA[i:], m.RefId)
		i = encodeVarintActionByRefId(dAtA, i, uint64(len(m.RefId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionByRefId(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionByRefId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionByRefId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RefId)
	if l > 0 {
		n += 1 + l + sovActionByRefId(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovActionByRefId(uint64(l))
	}
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionByRefId(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovActionByRefId(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovActionByRefId(uint64(l))
	}
	return n
}

func sovActionByRefId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionByRefId(x uint64) (n int) {
	return sovActionByRefId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionByRefId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionByRefId
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
			return fmt.Errorf("proto: ActionByRefId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionByRefId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionByRefId
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
				return ErrInvalidLengthActionByRefId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionByRefId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionByRefId
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
				return ErrInvalidLengthActionByRefId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionByRefId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionByRefId
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
				return ErrInvalidLengthActionByRefId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionByRefId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionByRefId
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
				return ErrInvalidLengthActionByRefId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionByRefId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionByRefId
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
				return ErrInvalidLengthActionByRefId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionByRefId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionByRefId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionByRefId
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
func skipActionByRefId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionByRefId
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
					return 0, ErrIntOverflowActionByRefId
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
					return 0, ErrIntOverflowActionByRefId
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
				return 0, ErrInvalidLengthActionByRefId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionByRefId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionByRefId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionByRefId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionByRefId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionByRefId = fmt.Errorf("proto: unexpected end of group")
)
