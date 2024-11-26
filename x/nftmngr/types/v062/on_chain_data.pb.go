// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v062/on_chain_data.proto

package v062

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

type OnChainData struct {
	RevealRequired     bool                   `protobuf:"varint,1,opt,name=reveal_required,json=revealRequired,proto3" json:"reveal_required,omitempty"`
	RevealSecret       []byte                 `protobuf:"bytes,2,opt,name=reveal_secret,json=revealSecret,proto3" json:"reveal_secret,omitempty"`
	NftAttributes      []*AttributeDefinition `protobuf:"bytes,3,rep,name=nft_attributes,json=nftAttributes,proto3" json:"nft_attributes,omitempty"`
	TokenAttributes    []*AttributeDefinition `protobuf:"bytes,4,rep,name=token_attributes,json=tokenAttributes,proto3" json:"token_attributes,omitempty"`
	Actions            []*Action              `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	Status             map[string]bool        `protobuf:"bytes,6,rep,name=status,proto3" json:"status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	OnOffSwitch        map[string]bool        `protobuf:"bytes,7,rep,name=on_off_switch,json=onOffSwitch,proto3" json:"on_off_switch,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	NftAttributesValue []*NftAttributeValue   `protobuf:"bytes,8,rep,name=nft_attributes_value,json=nftAttributesValue,proto3" json:"nft_attributes_value,omitempty"`
}

func (m *OnChainData) Reset()         { *m = OnChainData{} }
func (m *OnChainData) String() string { return proto.CompactTextString(m) }
func (*OnChainData) ProtoMessage()    {}
func (*OnChainData) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c070fd2f41c105, []int{0}
}
func (m *OnChainData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OnChainData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OnChainData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OnChainData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnChainData.Merge(m, src)
}
func (m *OnChainData) XXX_Size() int {
	return m.Size()
}
func (m *OnChainData) XXX_DiscardUnknown() {
	xxx_messageInfo_OnChainData.DiscardUnknown(m)
}

var xxx_messageInfo_OnChainData proto.InternalMessageInfo

func (m *OnChainData) GetRevealRequired() bool {
	if m != nil {
		return m.RevealRequired
	}
	return false
}

func (m *OnChainData) GetRevealSecret() []byte {
	if m != nil {
		return m.RevealSecret
	}
	return nil
}

func (m *OnChainData) GetNftAttributes() []*AttributeDefinition {
	if m != nil {
		return m.NftAttributes
	}
	return nil
}

func (m *OnChainData) GetTokenAttributes() []*AttributeDefinition {
	if m != nil {
		return m.TokenAttributes
	}
	return nil
}

func (m *OnChainData) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *OnChainData) GetStatus() map[string]bool {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OnChainData) GetOnOffSwitch() map[string]bool {
	if m != nil {
		return m.OnOffSwitch
	}
	return nil
}

func (m *OnChainData) GetNftAttributesValue() []*NftAttributeValue {
	if m != nil {
		return m.NftAttributesValue
	}
	return nil
}

func init() {
	proto.RegisterType((*OnChainData)(nil), "thesixnetwork.sixprotocol.nftmngr.v062.OnChainData")
	proto.RegisterMapType((map[string]bool)(nil), "thesixnetwork.sixprotocol.nftmngr.v062.OnChainData.OnOffSwitchEntry")
	proto.RegisterMapType((map[string]bool)(nil), "thesixnetwork.sixprotocol.nftmngr.v062.OnChainData.StatusEntry")
}

func init() { proto.RegisterFile("nftmngr/v062/on_chain_data.proto", fileDescriptor_05c070fd2f41c105) }

var fileDescriptor_05c070fd2f41c105 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x96, 0x75, 0xc3, 0x5d, 0xb7, 0xca, 0xda, 0x21, 0xf4, 0x10, 0x55, 0x20, 0x6d,
	0xbd, 0xe0, 0xa0, 0x21, 0x21, 0x36, 0x24, 0x10, 0x50, 0x24, 0x4e, 0x54, 0x4a, 0x11, 0x48, 0x5c,
	0x22, 0x37, 0xb5, 0x17, 0xab, 0x9d, 0x3d, 0xec, 0x97, 0xae, 0xfd, 0x16, 0x7c, 0x00, 0x3e, 0x10,
	0xc7, 0x1d, 0x39, 0xa2, 0xf6, 0x8b, 0xa0, 0xda, 0xe9, 0x9a, 0x80, 0x90, 0x0a, 0xdc, 0x92, 0xff,
	0x7b, 0xef, 0xf7, 0xf7, 0xf3, 0x7b, 0x46, 0x1d, 0xc9, 0xe1, 0x52, 0x5e, 0xe8, 0x70, 0xfa, 0xe8,
	0xc9, 0x69, 0xa8, 0x64, 0x9c, 0xa4, 0x54, 0xc8, 0x78, 0x44, 0x81, 0x92, 0x2b, 0xad, 0x40, 0xe1,
	0x63, 0x48, 0x99, 0x11, 0x33, 0xc9, 0xe0, 0x5a, 0xe9, 0x31, 0x31, 0x62, 0x66, 0xf5, 0x44, 0x4d,
	0x48, 0x5e, 0x4b, 0x56, 0xb5, 0xed, 0x93, 0x12, 0x89, 0x02, 0x68, 0x31, 0xcc, 0x80, 0xc5, 0x23,
	0xc6, 0x85, 0x14, 0x20, 0x94, 0x74, 0xc0, 0xf6, 0xbd, 0x72, 0x62, 0xf2, 0xc7, 0x90, 0x01, 0x0a,
	0x99, 0xc9, 0x43, 0xbf, 0x1d, 0x54, 0x71, 0x1e, 0x9b, 0x6b, 0x01, 0x49, 0x9a, 0x67, 0x1c, 0x97,
	0x32, 0x24, 0x87, 0x78, 0x73, 0x88, 0x29, 0x9d, 0x64, 0xcc, 0xe5, 0xdd, 0xff, 0x5a, 0x47, 0x8d,
	0xbe, 0x7c, 0xbd, 0xea, 0xb3, 0x47, 0x81, 0xe2, 0x13, 0x74, 0xa8, 0xd9, 0x94, 0xd1, 0x49, 0xac,
	0xd9, 0xe7, 0x4c, 0x68, 0x36, 0xf2, 0xbd, 0x8e, 0xd7, 0xdd, 0x8b, 0x0e, 0x9c, 0x1c, 0xe5, 0x2a,
	0x7e, 0x80, 0x9a, 0x79, 0xa2, 0x61, 0x89, 0x66, 0xe0, 0x57, 0x3b, 0x5e, 0x77, 0x3f, 0xda, 0x77,
	0xe2, 0xc0, 0x6a, 0x78, 0x88, 0x0e, 0x4a, 0xd6, 0xc6, 0xaf, 0x75, 0x6a, 0xdd, 0xc6, 0xe9, 0x33,
	0xb2, 0xdd, 0x3d, 0x92, 0x97, 0xeb, 0xca, 0xde, 0xed, 0xc5, 0x45, 0x4d, 0xc9, 0xe1, 0x56, 0x37,
	0x98, 0xa3, 0x16, 0xa8, 0x31, 0x93, 0x45, 0x97, 0x3b, 0xff, 0xef, 0x72, 0x68, 0xa1, 0x05, 0x9f,
	0xb7, 0x68, 0xd7, 0x8d, 0xc7, 0xf8, 0x3b, 0x16, 0x4f, 0xb6, 0xc6, 0xdb, 0xb2, 0x68, 0x5d, 0x8e,
	0x3f, 0xa2, 0xba, 0x9b, 0xa6, 0x5f, 0xb7, 0xa0, 0x17, 0xdb, 0x82, 0x0a, 0x83, 0x22, 0x03, 0x4b,
	0x78, 0x23, 0x41, 0xcf, 0xa3, 0x1c, 0x87, 0x53, 0xd4, 0x2c, 0xed, 0x82, 0xbf, 0x6b, 0xf9, 0xbd,
	0x7f, 0xe1, 0xf7, 0x65, 0x9f, 0xf3, 0x81, 0xc5, 0x38, 0x93, 0x86, 0xda, 0x28, 0x78, 0x8c, 0x8e,
	0xca, 0x83, 0x75, 0x4b, 0xe5, 0xef, 0x59, 0xc3, 0xb3, 0x6d, 0x0d, 0xdf, 0x15, 0x26, 0xf9, 0x61,
	0x05, 0x88, 0x70, 0x69, 0xb8, 0x56, 0x6b, 0x9f, 0xa1, 0x46, 0xa1, 0x5b, 0xdc, 0x42, 0xb5, 0x31,
	0x9b, 0xdb, 0xb5, 0xbc, 0x1b, 0xad, 0x3e, 0xf1, 0x11, 0xda, 0x71, 0xf6, 0x55, 0xbb, 0xaa, 0xee,
	0xe7, 0xbc, 0xfa, 0xd4, 0x6b, 0x3f, 0x47, 0xad, 0x5f, 0x1b, 0xf9, 0x9b, 0xfa, 0x57, 0xef, 0xbf,
	0x2d, 0x02, 0xef, 0x66, 0x11, 0x78, 0x3f, 0x16, 0x81, 0xf7, 0x65, 0x19, 0x54, 0x6e, 0x96, 0x41,
	0xe5, 0xfb, 0x32, 0xa8, 0x7c, 0x3a, 0xbf, 0x10, 0x90, 0x66, 0x43, 0x92, 0xa8, 0xcb, 0xb0, 0xd4,
	0x6d, 0x68, 0xc4, 0xec, 0xe1, 0xba, 0xdd, 0x70, 0x16, 0xae, 0x1f, 0x22, 0xcc, 0xaf, 0x98, 0xb1,
	0xcf, 0x71, 0x58, 0xb7, 0xe1, 0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2f, 0xe4, 0x25,
	0x70, 0x04, 0x00, 0x00,
}

func (m *OnChainData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OnChainData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OnChainData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftAttributesValue) > 0 {
		for iNdEx := len(m.NftAttributesValue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftAttributesValue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.OnOffSwitch) > 0 {
		for k := range m.OnOffSwitch {
			v := m.OnOffSwitch[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintOnChainData(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintOnChainData(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Status) > 0 {
		for k := range m.Status {
			v := m.Status[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintOnChainData(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintOnChainData(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TokenAttributes) > 0 {
		for iNdEx := len(m.TokenAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NftAttributes) > 0 {
		for iNdEx := len(m.NftAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RevealSecret) > 0 {
		i -= len(m.RevealSecret)
		copy(dAtA[i:], m.RevealSecret)
		i = encodeVarintOnChainData(dAtA, i, uint64(len(m.RevealSecret)))
		i--
		dAtA[i] = 0x12
	}
	if m.RevealRequired {
		i--
		if m.RevealRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOnChainData(dAtA []byte, offset int, v uint64) int {
	offset -= sovOnChainData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OnChainData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RevealRequired {
		n += 2
	}
	l = len(m.RevealSecret)
	if l > 0 {
		n += 1 + l + sovOnChainData(uint64(l))
	}
	if len(m.NftAttributes) > 0 {
		for _, e := range m.NftAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.TokenAttributes) > 0 {
		for _, e := range m.TokenAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Status) > 0 {
		for k, v := range m.Status {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovOnChainData(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovOnChainData(uint64(mapEntrySize))
		}
	}
	if len(m.OnOffSwitch) > 0 {
		for k, v := range m.OnOffSwitch {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovOnChainData(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovOnChainData(uint64(mapEntrySize))
		}
	}
	if len(m.NftAttributesValue) > 0 {
		for _, e := range m.NftAttributesValue {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	return n
}

func sovOnChainData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOnChainData(x uint64) (n int) {
	return sovOnChainData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OnChainData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnChainData
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
			return fmt.Errorf("proto: OnChainData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OnChainData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
			m.RevealRequired = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevealSecret = append(m.RevealSecret[:0], dAtA[iNdEx:postIndex]...)
			if m.RevealSecret == nil {
				m.RevealSecret = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftAttributes = append(m.NftAttributes, &AttributeDefinition{})
			if err := m.NftAttributes[len(m.NftAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenAttributes = append(m.TokenAttributes, &AttributeDefinition{})
			if err := m.TokenAttributes[len(m.TokenAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &Action{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnChainData
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnChainData
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthOnChainData
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthOnChainData
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnChainData
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOnChainData(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthOnChainData
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Status[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnOffSwitch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnOffSwitch == nil {
				m.OnOffSwitch = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOnChainData
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnChainData
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthOnChainData
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthOnChainData
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOnChainData
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOnChainData(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthOnChainData
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.OnOffSwitch[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftAttributesValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftAttributesValue = append(m.NftAttributesValue, &NftAttributeValue{})
			if err := m.NftAttributesValue[len(m.NftAttributesValue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOnChainData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOnChainData
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
func skipOnChainData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOnChainData
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
					return 0, ErrIntOverflowOnChainData
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
					return 0, ErrIntOverflowOnChainData
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
				return 0, ErrInvalidLengthOnChainData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOnChainData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOnChainData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOnChainData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOnChainData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOnChainData = fmt.Errorf("proto: unexpected end of group")
)
