// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/virtual_action.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

// TODO:
// 1. Policy of action (who can use action)
// 2. Proposal to enable and disable action
type VirtualAction struct {
	NftSchemaCode string   `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc          string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Disable       bool     `protobuf:"varint,4,opt,name=disable,proto3" json:"disable,omitempty"`
	When          string   `protobuf:"bytes,5,opt,name=when,proto3" json:"when,omitempty"`
	Then          []string `protobuf:"bytes,6,rep,name=then,proto3" json:"then,omitempty"`
	// NOTE: Policy of virtual action
	AllowedActioner AllowedActioner `protobuf:"varint,7,opt,name=allowedActioner,proto3,enum=thesixnetwork.sixnft.nftmngr.AllowedActioner" json:"allowedActioner,omitempty"`
	Params          []*ActionParams `protobuf:"bytes,8,rep,name=params,proto3" json:"params,omitempty"`
}

func (m *VirtualAction) Reset()         { *m = VirtualAction{} }
func (m *VirtualAction) String() string { return proto.CompactTextString(m) }
func (*VirtualAction) ProtoMessage()    {}
func (*VirtualAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a2c3f5c3a847225, []int{0}
}

func (m *VirtualAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *VirtualAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *VirtualAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualAction.Merge(m, src)
}

func (m *VirtualAction) XXX_Size() int {
	return m.Size()
}

func (m *VirtualAction) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualAction.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualAction proto.InternalMessageInfo

func (m *VirtualAction) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *VirtualAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualAction) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *VirtualAction) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *VirtualAction) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *VirtualAction) GetThen() []string {
	if m != nil {
		return m.Then
	}
	return nil
}

func (m *VirtualAction) GetAllowedActioner() AllowedActioner {
	if m != nil {
		return m.AllowedActioner
	}
	return AllowedActioner_ALLOWED_ACTIONER_ALL
}

func (m *VirtualAction) GetParams() []*ActionParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualAction)(nil), "thesixnetwork.sixnft.nftmngr.VirtualAction")
}

func init() { proto.RegisterFile("nftmngr/virtual_action.proto", fileDescriptor_9a2c3f5c3a847225) }

var fileDescriptor_9a2c3f5c3a847225 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x97, 0x6d, 0x6e, 0x33, 0x32, 0x85, 0xb0, 0x43, 0x18, 0xa3, 0x16, 0xf1, 0x50, 0x84,
	0x35, 0x30, 0xf1, 0x01, 0x36, 0x2f, 0xde, 0x94, 0x09, 0x1e, 0x04, 0x91, 0xac, 0xcd, 0xda, 0x60,
	0x9b, 0x8c, 0x24, 0x73, 0xf3, 0x2d, 0x7c, 0xac, 0x1d, 0x77, 0xf4, 0x24, 0xb2, 0xbd, 0x88, 0x24,
	0xdd, 0xc0, 0x7a, 0xd0, 0xdb, 0xbf, 0xbf, 0xef, 0xfb, 0xfd, 0x5b, 0xfa, 0xc1, 0x9e, 0x98, 0x9a,
	0x5c, 0x24, 0x8a, 0xbc, 0x72, 0x65, 0xe6, 0x34, 0x7b, 0xa6, 0x91, 0xe1, 0x52, 0x84, 0x33, 0x25,
	0x8d, 0x44, 0x3d, 0x93, 0x32, 0xcd, 0x97, 0x82, 0x99, 0x85, 0x54, 0x2f, 0xa1, 0x8d, 0x53, 0x13,
	0xee, 0x94, 0x6e, 0x67, 0xef, 0xfe, 0x74, 0xba, 0x9d, 0x44, 0x26, 0xd2, 0x45, 0x62, 0x53, 0x41,
	0xcf, 0x56, 0x55, 0xd8, 0x7e, 0x28, 0x5e, 0x31, 0x74, 0xdb, 0xe8, 0x1c, 0xb6, 0xc5, 0xd4, 0xdc,
	0x47, 0x29, 0xcb, 0xe9, 0xb5, 0x8c, 0x19, 0x06, 0x3e, 0x08, 0x0e, 0xc7, 0x65, 0x88, 0x10, 0xac,
	0x0b, 0x9a, 0x33, 0x5c, 0x75, 0x43, 0x97, 0x2d, 0x8b, 0x99, 0x8e, 0x70, 0xad, 0x60, 0x36, 0x23,
	0x0c, 0x9b, 0x31, 0xd7, 0x74, 0x92, 0x31, 0x5c, 0xf7, 0x41, 0xd0, 0x1a, 0xef, 0x1f, 0xed, 0xf6,
	0x22, 0x65, 0x02, 0x1f, 0x14, 0xdb, 0x36, 0x5b, 0x66, 0x2c, 0x6b, 0xf8, 0x35, 0xcb, 0x6c, 0x46,
	0x4f, 0xf0, 0x84, 0x66, 0x99, 0x5c, 0xb0, 0xb8, 0xf8, 0x40, 0xa6, 0x70, 0xd3, 0x07, 0xc1, 0xf1,
	0xa0, 0x1f, 0xfe, 0xf5, 0x17, 0xc2, 0x61, 0x59, 0x1a, 0xd5, 0x57, 0x9f, 0xa7, 0x60, 0xfc, 0xbb,
	0x0b, 0xdd, 0xc0, 0xc6, 0x8c, 0x2a, 0x9a, 0x6b, 0xdc, 0xf2, 0x6b, 0xc1, 0xd1, 0xe0, 0xe2, 0x9f,
	0x56, 0xe7, 0xdd, 0x39, 0x63, 0x57, 0xb9, 0xf3, 0x47, 0xb7, 0xab, 0x8d, 0x07, 0xd6, 0x1b, 0x0f,
	0x7c, 0x6d, 0x3c, 0xf0, 0xbe, 0xf5, 0x2a, 0xeb, 0xad, 0x57, 0xf9, 0xd8, 0x7a, 0x95, 0xc7, 0xab,
	0x84, 0x9b, 0x74, 0x3e, 0x09, 0x23, 0x99, 0x93, 0x52, 0x3b, 0xd1, 0x7c, 0xd9, 0x77, 0x87, 0x88,
	0x64, 0x46, 0x96, 0x64, 0x7f, 0x38, 0xf3, 0x36, 0x63, 0x7a, 0xd2, 0x70, 0x93, 0xcb, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x05, 0x10, 0xb8, 0x0c, 0x02, 0x00, 0x00,
}

func (m *VirtualAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Params) > 0 {
		for iNdEx := len(m.Params) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Params[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVirtualAction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.AllowedActioner != 0 {
		i = encodeVarintVirtualAction(dAtA, i, uint64(m.AllowedActioner))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Then) > 0 {
		for iNdEx := len(m.Then) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Then[iNdEx])
			copy(dAtA[i:], m.Then[iNdEx])
			i = encodeVarintVirtualAction(dAtA, i, uint64(len(m.Then[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.When) > 0 {
		i -= len(m.When)
		copy(dAtA[i:], m.When)
		i = encodeVarintVirtualAction(dAtA, i, uint64(len(m.When)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Disable {
		i--
		if m.Disable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintVirtualAction(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVirtualAction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintVirtualAction(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVirtualAction(dAtA []byte, offset int, v uint64) int {
	offset -= sovVirtualAction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *VirtualAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovVirtualAction(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVirtualAction(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovVirtualAction(uint64(l))
	}
	if m.Disable {
		n += 2
	}
	l = len(m.When)
	if l > 0 {
		n += 1 + l + sovVirtualAction(uint64(l))
	}
	if len(m.Then) > 0 {
		for _, s := range m.Then {
			l = len(s)
			n += 1 + l + sovVirtualAction(uint64(l))
		}
	}
	if m.AllowedActioner != 0 {
		n += 1 + sovVirtualAction(uint64(m.AllowedActioner))
	}
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovVirtualAction(uint64(l))
		}
	}
	return n
}

func sovVirtualAction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozVirtualAction(x uint64) (n int) {
	return sovVirtualAction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *VirtualAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVirtualAction
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
			return fmt.Errorf("proto: VirtualAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disable = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field When", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.When = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Then", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Then = append(m.Then, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedActioner", wireType)
			}
			m.AllowedActioner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowedActioner |= AllowedActioner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualAction
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
				return ErrInvalidLengthVirtualAction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, &ActionParams{})
			if err := m.Params[len(m.Params)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVirtualAction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVirtualAction
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

func skipVirtualAction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVirtualAction
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
					return 0, ErrIntOverflowVirtualAction
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
					return 0, ErrIntOverflowVirtualAction
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
				return 0, ErrInvalidLengthVirtualAction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVirtualAction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVirtualAction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVirtualAction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVirtualAction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVirtualAction = fmt.Errorf("proto: unexpected end of group")
)
