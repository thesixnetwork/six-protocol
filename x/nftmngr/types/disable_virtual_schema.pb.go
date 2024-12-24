// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/disable_virtual_schema.proto

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

type DisableVirtualSchema struct {
	Id                   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VirtualNftSchemaCode string `protobuf:"bytes,2,opt,name=virtualNftSchemaCode,proto3" json:"virtualNftSchemaCode,omitempty"`
	ProposalExpiredBlock string `protobuf:"bytes,3,opt,name=proposalExpiredBlock,proto3" json:"proposalExpiredBlock,omitempty"`
}

func (m *DisableVirtualSchema) Reset()         { *m = DisableVirtualSchema{} }
func (m *DisableVirtualSchema) String() string { return proto.CompactTextString(m) }
func (*DisableVirtualSchema) ProtoMessage()    {}
func (*DisableVirtualSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{0}
}
func (m *DisableVirtualSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisableVirtualSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisableVirtualSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisableVirtualSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableVirtualSchema.Merge(m, src)
}
func (m *DisableVirtualSchema) XXX_Size() int {
	return m.Size()
}
func (m *DisableVirtualSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableVirtualSchema.DiscardUnknown(m)
}

var xxx_messageInfo_DisableVirtualSchema proto.InternalMessageInfo

func (m *DisableVirtualSchema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DisableVirtualSchema) GetVirtualNftSchemaCode() string {
	if m != nil {
		return m.VirtualNftSchemaCode
	}
	return ""
}

func (m *DisableVirtualSchema) GetProposalExpiredBlock() string {
	if m != nil {
		return m.ProposalExpiredBlock
	}
	return ""
}

type VirtualSchemaDisableRequest struct {
	Id                   string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VirtualNftSchemaCode string                          `protobuf:"bytes,2,opt,name=virtualNftSchemaCode,proto3" json:"virtualNftSchemaCode,omitempty"`
	Registry             []*VirtualSchemaDisableRegistry `protobuf:"bytes,3,rep,name=registry,proto3" json:"registry,omitempty"`
	ProposalExpiredBlock string                          `protobuf:"bytes,4,opt,name=proposalExpiredBlock,proto3" json:"proposalExpiredBlock,omitempty"`
	Creator              string                          `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *VirtualSchemaDisableRequest) Reset()         { *m = VirtualSchemaDisableRequest{} }
func (m *VirtualSchemaDisableRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaDisableRequest) ProtoMessage()    {}
func (*VirtualSchemaDisableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{1}
}
func (m *VirtualSchemaDisableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VirtualSchemaDisableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaDisableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VirtualSchemaDisableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaDisableRequest.Merge(m, src)
}
func (m *VirtualSchemaDisableRequest) XXX_Size() int {
	return m.Size()
}
func (m *VirtualSchemaDisableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaDisableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaDisableRequest proto.InternalMessageInfo

func (m *VirtualSchemaDisableRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualSchemaDisableRequest) GetVirtualNftSchemaCode() string {
	if m != nil {
		return m.VirtualNftSchemaCode
	}
	return ""
}

func (m *VirtualSchemaDisableRequest) GetRegistry() []*VirtualSchemaDisableRegistry {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *VirtualSchemaDisableRequest) GetProposalExpiredBlock() string {
	if m != nil {
		return m.ProposalExpiredBlock
	}
	return ""
}

func (m *VirtualSchemaDisableRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type VirtualSchemaDisableRegistry struct {
	NftSchemaCode string         `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Status        RegistryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=thesixnetwork.sixprotocol.nftmngr.RegistryStatus" json:"status,omitempty"`
}

func (m *VirtualSchemaDisableRegistry) Reset()         { *m = VirtualSchemaDisableRegistry{} }
func (m *VirtualSchemaDisableRegistry) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaDisableRegistry) ProtoMessage()    {}
func (*VirtualSchemaDisableRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{2}
}
func (m *VirtualSchemaDisableRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VirtualSchemaDisableRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaDisableRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VirtualSchemaDisableRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaDisableRegistry.Merge(m, src)
}
func (m *VirtualSchemaDisableRegistry) XXX_Size() int {
	return m.Size()
}
func (m *VirtualSchemaDisableRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaDisableRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaDisableRegistry proto.InternalMessageInfo

func (m *VirtualSchemaDisableRegistry) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *VirtualSchemaDisableRegistry) GetStatus() RegistryStatus {
	if m != nil {
		return m.Status
	}
	return RegistryStatus_PENDING
}

func init() {
	proto.RegisterType((*DisableVirtualSchema)(nil), "thesixnetwork.sixprotocol.nftmngr.DisableVirtualSchema")
	proto.RegisterType((*VirtualSchemaDisableRequest)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaDisableRequest")
	proto.RegisterType((*VirtualSchemaDisableRegistry)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaDisableRegistry")
}

func init() {
	proto.RegisterFile("nftmngr/disable_virtual_schema.proto", fileDescriptor_d2dd6dee3450cf76)
}

var fileDescriptor_d2dd6dee3450cf76 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4f, 0xfa, 0x40,
	0x18, 0xc6, 0xb9, 0xf2, 0xff, 0xa3, 0x9e, 0x91, 0xa1, 0x61, 0x68, 0x94, 0x34, 0x48, 0x18, 0x58,
	0x6c, 0x23, 0xc6, 0xd9, 0x04, 0x75, 0x70, 0xd1, 0xa4, 0x24, 0x0e, 0x3a, 0x90, 0xd2, 0x1e, 0x70,
	0xa1, 0xf4, 0xea, 0xdd, 0x5b, 0x2d, 0xdf, 0xc1, 0xa8, 0x1f, 0xcb, 0x91, 0xd1, 0xd1, 0xd0, 0x2f,
	0x62, 0xbc, 0x6b, 0x4d, 0x9a, 0x14, 0x75, 0x70, 0xec, 0x3d, 0xef, 0xf3, 0x7b, 0x9f, 0xe7, 0x4d,
	0x71, 0x27, 0x1c, 0xc3, 0x3c, 0x9c, 0x70, 0xdb, 0xa7, 0xc2, 0x1d, 0x05, 0x64, 0x78, 0x4f, 0x39,
	0xc4, 0x6e, 0x30, 0x14, 0xde, 0x94, 0xcc, 0x5d, 0x2b, 0xe2, 0x0c, 0x98, 0xbe, 0x0f, 0x53, 0x22,
	0x68, 0x12, 0x12, 0x78, 0x60, 0x7c, 0x66, 0x09, 0x9a, 0xc8, 0x77, 0x8f, 0x05, 0x56, 0xe6, 0xdf,
	0x6d, 0xe6, 0xa0, 0x32, 0x40, 0xfb, 0x09, 0xe1, 0xc6, 0x99, 0xda, 0x70, 0xad, 0xf4, 0x81, 0x94,
	0xf5, 0x3a, 0xd6, 0xa8, 0x6f, 0xa0, 0x16, 0xea, 0x6e, 0x39, 0x1a, 0xf5, 0xf5, 0x1e, 0x6e, 0x64,
	0x80, 0xcb, 0x31, 0xa8, 0x99, 0x53, 0xe6, 0x13, 0x43, 0x93, 0x13, 0xa5, 0xda, 0xa7, 0x27, 0xe2,
	0x2c, 0x62, 0xc2, 0x0d, 0xce, 0x93, 0x88, 0x72, 0xe2, 0xf7, 0x03, 0xe6, 0xcd, 0x8c, 0xaa, 0xf2,
	0x94, 0x69, 0xed, 0x47, 0x0d, 0xef, 0x15, 0x92, 0x64, 0xe9, 0x1c, 0x72, 0x17, 0x13, 0x01, 0x7f,
	0x92, 0xeb, 0x16, 0x6f, 0x72, 0x32, 0xa1, 0x02, 0xf8, 0xc2, 0xa8, 0xb6, 0xaa, 0xdd, 0xed, 0xde,
	0x89, 0xf5, 0xe3, 0x21, 0xad, 0xf2, 0x54, 0x0a, 0xe3, 0x7c, 0x01, 0xd7, 0x96, 0xfe, 0xb7, 0xbe,
	0xb4, 0x6e, 0xe0, 0x0d, 0x8f, 0x13, 0x17, 0x18, 0x37, 0xfe, 0xcb, 0xb1, 0xfc, 0xb3, 0xfd, 0x8c,
	0x70, 0xf3, 0xbb, 0xc5, 0x7a, 0x07, 0xef, 0x84, 0x85, 0xe2, 0xea, 0x34, 0xc5, 0x47, 0xfd, 0x02,
	0xd7, 0x04, 0xb8, 0x10, 0x0b, 0x79, 0x97, 0x7a, 0xef, 0xf0, 0x17, 0x7d, 0xf3, 0x15, 0x03, 0x69,
	0x74, 0x32, 0x40, 0xff, 0xea, 0x75, 0x65, 0xa2, 0xe5, 0xca, 0x44, 0xef, 0x2b, 0x13, 0xbd, 0xa4,
	0x66, 0x65, 0x99, 0x9a, 0x95, 0xb7, 0xd4, 0xac, 0xdc, 0x1c, 0x4f, 0x28, 0x4c, 0xe3, 0x91, 0xe5,
	0xb1, 0xb9, 0x5d, 0xc0, 0xdb, 0x82, 0x26, 0x07, 0x39, 0xdf, 0x4e, 0xec, 0xfc, 0x8f, 0x84, 0x45,
	0x44, 0xc4, 0xa8, 0x26, 0x95, 0xa3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xf5, 0xe3, 0xe5,
	0xf2, 0x02, 0x00, 0x00,
}

func (m *DisableVirtualSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisableVirtualSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisableVirtualSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProposalExpiredBlock) > 0 {
		i -= len(m.ProposalExpiredBlock)
		copy(dAtA[i:], m.ProposalExpiredBlock)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.ProposalExpiredBlock)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VirtualNftSchemaCode) > 0 {
		i -= len(m.VirtualNftSchemaCode)
		copy(dAtA[i:], m.VirtualNftSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.VirtualNftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchemaDisableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaDisableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaDisableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProposalExpiredBlock) > 0 {
		i -= len(m.ProposalExpiredBlock)
		copy(dAtA[i:], m.ProposalExpiredBlock)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.ProposalExpiredBlock)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VirtualNftSchemaCode) > 0 {
		i -= len(m.VirtualNftSchemaCode)
		copy(dAtA[i:], m.VirtualNftSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.VirtualNftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchemaDisableRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaDisableRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaDisableRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDisableVirtualSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovDisableVirtualSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DisableVirtualSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.VirtualNftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.ProposalExpiredBlock)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	return n
}

func (m *VirtualSchemaDisableRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.VirtualNftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovDisableVirtualSchema(uint64(l))
		}
	}
	l = len(m.ProposalExpiredBlock)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	return n
}

func (m *VirtualSchemaDisableRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDisableVirtualSchema(uint64(m.Status))
	}
	return n
}

func sovDisableVirtualSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDisableVirtualSchema(x uint64) (n int) {
	return sovDisableVirtualSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DisableVirtualSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisableVirtualSchema
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
			return fmt.Errorf("proto: DisableVirtualSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisableVirtualSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualNftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualNftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalExpiredBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalExpiredBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDisableVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisableVirtualSchema
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
func (m *VirtualSchemaDisableRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisableVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchemaDisableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaDisableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualNftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualNftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, &VirtualSchemaDisableRegistry{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalExpiredBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalExpiredBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDisableVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisableVirtualSchema
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
func (m *VirtualSchemaDisableRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisableVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchemaDisableRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaDisableRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RegistryStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDisableVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisableVirtualSchema
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
func skipDisableVirtualSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDisableVirtualSchema
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
					return 0, ErrIntOverflowDisableVirtualSchema
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
					return 0, ErrIntOverflowDisableVirtualSchema
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
				return 0, ErrInvalidLengthDisableVirtualSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDisableVirtualSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDisableVirtualSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDisableVirtualSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDisableVirtualSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDisableVirtualSchema = fmt.Errorf("proto: unexpected end of group")
)
