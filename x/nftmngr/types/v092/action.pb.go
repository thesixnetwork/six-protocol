// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v092/action.proto

package v092

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type AllowedActioner int32

const (
	AllowedActioner_ALLOWED_ACTIONER_ALL         AllowedActioner = 0
	AllowedActioner_ALLOWED_ACTIONER_SYSTEM_ONLY AllowedActioner = 1
	AllowedActioner_ALLOWED_ACTIONER_USER_ONLY   AllowedActioner = 2
)

var AllowedActioner_name = map[int32]string{
	0: "ALLOWED_ACTIONER_ALL",
	1: "ALLOWED_ACTIONER_SYSTEM_ONLY",
	2: "ALLOWED_ACTIONER_USER_ONLY",
}

var AllowedActioner_value = map[string]int32{
	"ALLOWED_ACTIONER_ALL":         0,
	"ALLOWED_ACTIONER_SYSTEM_ONLY": 1,
	"ALLOWED_ACTIONER_USER_ONLY":   2,
}

func (x AllowedActioner) String() string {
	return proto.EnumName(AllowedActioner_name, int32(x))
}

func (AllowedActioner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3def40d856f8ce60, []int{0}
}

type ActionParams struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc         string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	DataType     string `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Required     bool   `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	DefaultValue string `protobuf:"bytes,5,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (m *ActionParams) Reset()         { *m = ActionParams{} }
func (m *ActionParams) String() string { return proto.CompactTextString(m) }
func (*ActionParams) ProtoMessage()    {}
func (*ActionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3def40d856f8ce60, []int{0}
}
func (m *ActionParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionParams.Merge(m, src)
}
func (m *ActionParams) XXX_Size() int {
	return m.Size()
}
func (m *ActionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionParams.DiscardUnknown(m)
}

var xxx_messageInfo_ActionParams proto.InternalMessageInfo

func (m *ActionParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionParams) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ActionParams) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *ActionParams) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *ActionParams) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type Action struct {
	Name            string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc            string          `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Disable         bool            `protobuf:"varint,3,opt,name=disable,proto3" json:"disable,omitempty"`
	When            string          `protobuf:"bytes,4,opt,name=when,proto3" json:"when,omitempty"`
	Then            []string        `protobuf:"bytes,5,rep,name=then,proto3" json:"then,omitempty"`
	AllowedActioner AllowedActioner `protobuf:"varint,6,opt,name=allowed_actioner,json=allowedActioner,proto3,enum=thesixnetwork.sixprotocol.nftmngr.v092.AllowedActioner" json:"allowed_actioner,omitempty"`
	Params          []*ActionParams `protobuf:"bytes,7,rep,name=params,proto3" json:"params,omitempty"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_3def40d856f8ce60, []int{1}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Action.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return m.Size()
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Action) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *Action) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *Action) GetThen() []string {
	if m != nil {
		return m.Then
	}
	return nil
}

func (m *Action) GetAllowedActioner() AllowedActioner {
	if m != nil {
		return m.AllowedActioner
	}
	return AllowedActioner_ALLOWED_ACTIONER_ALL
}

func (m *Action) GetParams() []*ActionParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixprotocol.nftmngr.v092.AllowedActioner", AllowedActioner_name, AllowedActioner_value)
	proto.RegisterType((*ActionParams)(nil), "thesixnetwork.sixprotocol.nftmngr.v092.action_params")
	proto.RegisterType((*Action)(nil), "thesixnetwork.sixprotocol.nftmngr.v092.Action")
}

func init() { proto.RegisterFile("nftmngr/v092/action.proto", fileDescriptor_3def40d856f8ce60) }

var fileDescriptor_3def40d856f8ce60 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0xb6, 0xeb, 0x5a, 0xc3, 0x58, 0x65, 0xed, 0x60, 0x0a, 0x0a, 0xd1, 0x90, 0x50,
	0x85, 0x44, 0x82, 0x8a, 0x10, 0x82, 0x5b, 0x06, 0x3d, 0x20, 0x85, 0x55, 0x4a, 0x03, 0x68, 0x5c,
	0x22, 0x37, 0xf1, 0x92, 0x88, 0x24, 0x0e, 0x89, 0xb3, 0x66, 0x6f, 0xc1, 0x85, 0x07, 0xe1, 0x2d,
	0x76, 0xdc, 0x91, 0x13, 0x42, 0xed, 0x8b, 0x20, 0xdb, 0x29, 0x52, 0xe9, 0xa5, 0xb7, 0xbf, 0x7f,
	0x7f, 0xff, 0xbf, 0x7c, 0xfe, 0xf2, 0xc1, 0xfb, 0xd9, 0x25, 0x4f, 0xb3, 0xb0, 0x30, 0xaf, 0x9e,
	0xbf, 0x9e, 0x98, 0xc4, 0xe7, 0x31, 0xcb, 0x8c, 0xbc, 0x60, 0x9c, 0xa1, 0x27, 0x3c, 0xa2, 0x65,
	0x5c, 0x67, 0x94, 0x2f, 0x59, 0xf1, 0xd5, 0x28, 0xe3, 0x5a, 0x72, 0x9f, 0x25, 0x46, 0x13, 0x32,
	0x44, 0x68, 0x74, 0x12, 0xb2, 0x90, 0x49, 0xcb, 0x14, 0x4a, 0xa5, 0x4f, 0x7f, 0x00, 0x78, 0xa4,
	0xca, 0x79, 0x39, 0x29, 0x48, 0x5a, 0x22, 0x04, 0xbb, 0x19, 0x49, 0x29, 0x06, 0x3a, 0x18, 0x0f,
	0x1c, 0xa9, 0x05, 0x0b, 0x68, 0xe9, 0xe3, 0xb6, 0x62, 0x42, 0xa3, 0x07, 0x70, 0x10, 0x10, 0x4e,
	0x3c, 0x7e, 0x9d, 0x53, 0xdc, 0x91, 0x46, 0x5f, 0x00, 0xf7, 0x3a, 0xa7, 0x68, 0x04, 0xfb, 0x05,
	0xfd, 0x56, 0xc5, 0x05, 0x0d, 0x70, 0x57, 0x07, 0xe3, 0xbe, 0xf3, 0xef, 0x8c, 0x1e, 0xc3, 0xa3,
	0x80, 0x5e, 0x92, 0x2a, 0xe1, 0xde, 0x15, 0x49, 0x2a, 0x8a, 0x0f, 0x64, 0xf8, 0x6e, 0x03, 0x3f,
	0x09, 0x76, 0xfa, 0xb3, 0x0d, 0x7b, 0x96, 0xec, 0x6b, 0xef, 0x86, 0x30, 0x3c, 0x0c, 0xe2, 0x92,
	0x2c, 0x12, 0xd5, 0x4e, 0xdf, 0xd9, 0x1c, 0xc5, 0xed, 0x65, 0x44, 0x33, 0xd9, 0xc9, 0xc0, 0x91,
	0x5a, 0x30, 0x2e, 0xd8, 0x81, 0xde, 0x11, 0x4c, 0x68, 0x14, 0xc1, 0x21, 0x49, 0x12, 0xb6, 0xa4,
	0x81, 0xa7, 0x66, 0x42, 0x0b, 0xdc, 0xd3, 0xc1, 0xf8, 0xde, 0xe4, 0x95, 0xb1, 0xdf, 0x94, 0x0d,
	0x4b, 0xe5, 0xad, 0x26, 0x7e, 0xd6, 0xbd, 0xf9, 0xfd, 0x08, 0x38, 0xc7, 0x64, 0x1b, 0xa3, 0x39,
	0xec, 0xa9, 0x71, 0xe3, 0x43, 0xbd, 0x33, 0xbe, 0x33, 0x79, 0xb9, 0x6f, 0xfd, 0xad, 0x7f, 0xd5,
	0x54, 0x6f, 0x4a, 0x3d, 0x4d, 0xe1, 0xf1, 0x7f, 0x9f, 0x47, 0x18, 0x9e, 0x58, 0xb6, 0x3d, 0xfb,
	0x3c, 0x7d, 0xe7, 0x59, 0x6f, 0xdd, 0xf7, 0xb3, 0xf3, 0xa9, 0xe3, 0x59, 0xb6, 0x3d, 0x6c, 0x21,
	0x1d, 0x3e, 0xdc, 0x71, 0xe6, 0x17, 0x73, 0x77, 0xfa, 0xc1, 0x9b, 0x9d, 0xdb, 0x17, 0x43, 0x80,
	0x34, 0x38, 0xda, 0xb9, 0xf1, 0x71, 0x3e, 0x75, 0x94, 0xdf, 0x3e, 0x73, 0x6f, 0x56, 0x1a, 0xb8,
	0x5d, 0x69, 0xe0, 0xcf, 0x4a, 0x03, 0xdf, 0xd7, 0x5a, 0xeb, 0x76, 0xad, 0xb5, 0x7e, 0xad, 0xb5,
	0xd6, 0x97, 0x37, 0x61, 0xcc, 0xa3, 0x6a, 0x61, 0xf8, 0x2c, 0x35, 0xb7, 0xde, 0x65, 0x96, 0x71,
	0xfd, 0x6c, 0xf3, 0x30, 0xb3, 0x36, 0x37, 0x5b, 0x2d, 0x76, 0xa8, 0x94, 0xbb, 0xbd, 0xe8, 0x49,
	0xfb, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xbb, 0x1a, 0x3c, 0xf2, 0x02, 0x00, 0x00,
}

func (m *ActionParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DefaultValue) > 0 {
		i -= len(m.DefaultValue)
		copy(dAtA[i:], m.DefaultValue)
		i = encodeVarintAction(dAtA, i, uint64(len(m.DefaultValue)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Required {
		i--
		if m.Required {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.DataType) > 0 {
		i -= len(m.DataType)
		copy(dAtA[i:], m.DataType)
		i = encodeVarintAction(dAtA, i, uint64(len(m.DataType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Action) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Action) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Action) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintAction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.AllowedActioner != 0 {
		i = encodeVarintAction(dAtA, i, uint64(m.AllowedActioner))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Then) > 0 {
		for iNdEx := len(m.Then) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Then[iNdEx])
			copy(dAtA[i:], m.Then[iNdEx])
			i = encodeVarintAction(dAtA, i, uint64(len(m.Then[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.When) > 0 {
		i -= len(m.When)
		copy(dAtA[i:], m.When)
		i = encodeVarintAction(dAtA, i, uint64(len(m.When)))
		i--
		dAtA[i] = 0x22
	}
	if m.Disable {
		i--
		if m.Disable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	l = len(m.DataType)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	if m.Required {
		n += 2
	}
	l = len(m.DefaultValue)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	return n
}

func (m *Action) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	if m.Disable {
		n += 2
	}
	l = len(m.When)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	if len(m.Then) > 0 {
		for _, s := range m.Then {
			l = len(s)
			n += 1 + l + sovAction(uint64(l))
		}
	}
	if m.AllowedActioner != 0 {
		n += 1 + sovAction(uint64(m.AllowedActioner))
	}
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovAction(uint64(l))
		}
	}
	return n
}

func sovAction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAction(x uint64) (n int) {
	return sovAction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAction
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
			return fmt.Errorf("proto: action_params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: action_params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
			m.Required = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAction
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
func (m *Action) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAction
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
			return fmt.Errorf("proto: Action: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Action: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field When", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.When = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Then", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Then = append(m.Then, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedActioner", wireType)
			}
			m.AllowedActioner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAction
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
			skippy, err := skipAction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAction
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
func skipAction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAction
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
					return 0, ErrIntOverflowAction
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
					return 0, ErrIntOverflowAction
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
				return 0, ErrInvalidLengthAction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAction = fmt.Errorf("proto: unexpected end of group")
)
