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
	Status        RegistryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=thesixnetwork.sixnft.nftmngr.RegistryStatus" json:"status,omitempty"`
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
	proto.RegisterType((*DisableVirtualSchema)(nil), "thesixnetwork.sixnft.nftmngr.DisableVirtualSchema")
	proto.RegisterType((*VirtualSchemaDisableRequest)(nil), "thesixnetwork.sixnft.nftmngr.VirtualSchemaDisableRequest")
	proto.RegisterType((*VirtualSchemaDisableRegistry)(nil), "thesixnetwork.sixnft.nftmngr.VirtualSchemaDisableRegistry")
}

func init() {
	proto.RegisterFile("nftmngr/disable_virtual_schema.proto", fileDescriptor_d2dd6dee3450cf76)
}

var fileDescriptor_d2dd6dee3450cf76 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4f, 0xc2, 0x40,
	0x14, 0xc6, 0xb9, 0xa2, 0xa8, 0x67, 0x64, 0x68, 0x18, 0x1a, 0x25, 0x0d, 0x21, 0x0c, 0x0c, 0xda,
	0x26, 0x18, 0x17, 0x47, 0xc4, 0x55, 0x93, 0x92, 0x30, 0xb8, 0x90, 0xd2, 0x1e, 0x70, 0xa1, 0xf4,
	0xea, 0xdd, 0xab, 0x96, 0xbf, 0xc0, 0xc4, 0xc1, 0xf8, 0x67, 0x39, 0x32, 0x3a, 0x1a, 0xf8, 0x47,
	0x0c, 0x77, 0xad, 0x49, 0x93, 0xd2, 0xc9, 0xad, 0xed, 0xf7, 0xbe, 0x5f, 0xbf, 0xef, 0xe5, 0xe1,
	0x4e, 0x38, 0x85, 0x65, 0x38, 0xe3, 0xb6, 0x4f, 0x85, 0x3b, 0x09, 0xc8, 0xf8, 0x85, 0x72, 0x88,
	0xdd, 0x60, 0x2c, 0xbc, 0x39, 0x59, 0xba, 0x56, 0xc4, 0x19, 0x30, 0xbd, 0x09, 0x73, 0x22, 0x68,
	0x12, 0x12, 0x78, 0x65, 0x7c, 0x61, 0xed, 0x1e, 0xa7, 0x60, 0xa5, 0xd6, 0xf3, 0x66, 0xc6, 0x28,
	0xf2, 0xb6, 0x3f, 0x10, 0x6e, 0x0c, 0x14, 0x7c, 0xa4, 0xf4, 0xa1, 0x94, 0xf5, 0x3a, 0xd6, 0xa8,
	0x6f, 0xa0, 0x16, 0xea, 0x9e, 0x38, 0x1a, 0xf5, 0xf5, 0x1e, 0x6e, 0xa4, 0x80, 0x87, 0x29, 0xa8,
	0x99, 0x3b, 0xe6, 0x13, 0x43, 0x93, 0x13, 0x85, 0xda, 0xce, 0x13, 0x71, 0x16, 0x31, 0xe1, 0x06,
	0xf7, 0x49, 0x44, 0x39, 0xf1, 0xfb, 0x01, 0xf3, 0x16, 0x46, 0x55, 0x79, 0x8a, 0xb4, 0xf6, 0x9b,
	0x86, 0x2f, 0x72, 0x49, 0xd2, 0x74, 0x0e, 0x79, 0x8e, 0x89, 0x80, 0x7f, 0xc9, 0x35, 0xc2, 0xc7,
	0x9c, 0xcc, 0xa8, 0x00, 0xbe, 0x32, 0xaa, 0xad, 0x6a, 0xf7, 0xb4, 0x77, 0x6b, 0x95, 0xed, 0xd0,
	0x2a, 0x0e, 0xa4, 0x08, 0xce, 0x1f, 0x6b, 0x6f, 0xdf, 0x83, 0xfd, 0x7d, 0x75, 0x03, 0x1f, 0x79,
	0x9c, 0xb8, 0xc0, 0xb8, 0x71, 0x28, 0xc7, 0xb2, 0xd7, 0xf6, 0x3b, 0xc2, 0xcd, 0xb2, 0x1f, 0xeb,
	0x1d, 0x7c, 0x16, 0xe6, 0x3a, 0xab, 0xad, 0xe4, 0x3f, 0xea, 0x03, 0x5c, 0x13, 0xe0, 0x42, 0x2c,
	0xe4, 0x4a, 0xea, 0xbd, 0xcb, 0xf2, 0xaa, 0x19, 0x7d, 0x28, 0x3d, 0x4e, 0xea, 0xed, 0x3f, 0x7e,
	0x6d, 0x4c, 0xb4, 0xde, 0x98, 0xe8, 0x67, 0x63, 0xa2, 0xcf, 0xad, 0x59, 0x59, 0x6f, 0xcd, 0xca,
	0xf7, 0xd6, 0xac, 0x3c, 0xdd, 0xcc, 0x28, 0xcc, 0xe3, 0x89, 0xe5, 0xb1, 0xa5, 0x9d, 0x23, 0xdb,
	0x82, 0x26, 0x57, 0xf2, 0xca, 0x3c, 0x16, 0xd8, 0x89, 0x9d, 0xdd, 0x21, 0xac, 0x22, 0x22, 0x26,
	0x35, 0xa9, 0x5c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x40, 0xc8, 0xd9, 0xe3, 0x02, 0x00,
	0x00,
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
