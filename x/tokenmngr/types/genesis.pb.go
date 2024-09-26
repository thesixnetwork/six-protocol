// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/v2/genesis.proto

package types

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

// GenesisState defines the tokenmngr module's genesis state.
type GenesisState struct {
	Params        Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId        string      `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	TokenList     []Token     `protobuf:"bytes,3,rep,name=tokenList,proto3" json:"tokenList"`
	MintpermList  []Mintperm  `protobuf:"bytes,4,rep,name=mintpermList,proto3" json:"mintpermList"`
	Options       *Options    `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	TokenBurnList []TokenBurn `protobuf:"bytes,7,rep,name=tokenBurnList,proto3" json:"tokenBurnList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ffe7365b897bc6d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetTokenList() []Token {
	if m != nil {
		return m.TokenList
	}
	return nil
}

func (m *GenesisState) GetMintpermList() []Mintperm {
	if m != nil {
		return m.MintpermList
	}
	return nil
}

func (m *GenesisState) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *GenesisState) GetTokenBurnList() []TokenBurn {
	if m != nil {
		return m.TokenBurnList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "thesixnetwork.sixprotocol.tokenmngr.v2.GenesisState")
}

func init() { proto.RegisterFile("tokenmngr/v2/genesis.proto", fileDescriptor_3ffe7365b897bc6d) }

var fileDescriptor_3ffe7365b897bc6d = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x18, 0xc5, 0xdb, 0x0b, 0x29, 0x61, 0xe0, 0x6e, 0x26, 0x37, 0xb9, 0xbd, 0xbd, 0xb1, 0x12, 0x17,
	0x86, 0x0d, 0x33, 0x5a, 0x17, 0xee, 0xd9, 0x18, 0x12, 0x8c, 0x82, 0xae, 0x48, 0x0c, 0xe1, 0xcf,
	0xa4, 0x4c, 0xb0, 0x9d, 0x66, 0x66, 0x8a, 0xf5, 0x2d, 0x7c, 0x2c, 0x56, 0x86, 0xa5, 0x2b, 0x63,
	0xe0, 0x45, 0x0c, 0xd3, 0xa9, 0x64, 0xe2, 0xa6, 0xee, 0xfa, 0xf5, 0xf4, 0xfc, 0xbe, 0xef, 0x9c,
	0x14, 0x78, 0x92, 0x2d, 0x49, 0x1c, 0xc5, 0x21, 0xc7, 0xab, 0x00, 0x87, 0x24, 0x26, 0x82, 0x0a,
	0x94, 0x70, 0x26, 0x19, 0x3c, 0x95, 0x0b, 0x22, 0x68, 0x16, 0x13, 0xf9, 0xc4, 0xf8, 0x12, 0x09,
	0x9a, 0xa9, 0xf7, 0x33, 0xf6, 0x88, 0xbe, 0x5c, 0x68, 0x15, 0x78, 0x7f, 0x42, 0x16, 0x32, 0x25,
	0xe1, 0xfd, 0x53, 0xee, 0xf6, 0xfe, 0x19, 0xe4, 0x64, 0xc2, 0x27, 0x91, 0x06, 0x7b, 0xae, 0x21,
	0xa9, 0x41, 0x2b, 0xff, 0x0d, 0x25, 0xa2, 0xb1, 0x4c, 0x08, 0x8f, 0xb4, 0x68, 0xde, 0xca, 0x12,
	0x49, 0x59, 0x5c, 0x20, 0x8f, 0xbe, 0x23, 0xc7, 0xd3, 0x94, 0x6b, 0xee, 0xc9, 0x6b, 0x05, 0x34,
	0xaf, 0xf2, 0x70, 0x77, 0x72, 0x22, 0x09, 0xec, 0x03, 0x27, 0x3f, 0xc9, 0xb5, 0x5b, 0x76, 0xbb,
	0x11, 0x20, 0x54, 0x2e, 0x2c, 0xba, 0x55, 0xae, 0x6e, 0x75, 0xfd, 0x7e, 0x6c, 0x0d, 0x35, 0x03,
	0xfe, 0x05, 0xb5, 0x84, 0x71, 0x39, 0xa6, 0x73, 0xf7, 0x57, 0xcb, 0x6e, 0xd7, 0x87, 0xce, 0x7e,
	0xec, 0xcd, 0xe1, 0x00, 0xd4, 0x95, 0xbb, 0x4f, 0x85, 0x74, 0x2b, 0xad, 0x4a, 0xbb, 0x11, 0x74,
	0xca, 0x6e, 0xba, 0xdf, 0x0f, 0x7a, 0xd1, 0x81, 0x02, 0x47, 0xa0, 0x59, 0xf4, 0xa2, 0xa8, 0x55,
	0x45, 0x3d, 0x2b, 0x4b, 0xbd, 0xd6, 0x5e, 0x0d, 0x36, 0x58, 0xb0, 0x07, 0x6a, 0xba, 0x56, 0xd7,
	0x51, 0xb5, 0xe0, 0xb2, 0xd8, 0x9b, 0xdc, 0x36, 0x2c, 0xfc, 0xf0, 0x01, 0xfc, 0x56, 0x1f, 0x74,
	0x53, 0x9e, 0xa7, 0xaf, 0xa9, 0x3b, 0xcf, 0x7f, 0x96, 0x3e, 0xe5, 0x45, 0x03, 0x26, 0xad, 0x3b,
	0x58, 0x6f, 0x7d, 0x7b, 0xb3, 0xf5, 0xed, 0x8f, 0xad, 0x6f, 0xbf, 0xec, 0x7c, 0x6b, 0xb3, 0xf3,
	0xad, 0xb7, 0x9d, 0x6f, 0x8d, 0x2e, 0x43, 0x2a, 0x17, 0xe9, 0x14, 0xcd, 0x58, 0x84, 0x8d, 0x5d,
	0x58, 0xd0, 0xac, 0x53, 0x2c, 0xc3, 0x19, 0x3e, 0xfc, 0x31, 0xf2, 0x39, 0x21, 0x62, 0xea, 0x28,
	0xed, 0xe2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x74, 0x21, 0x95, 0x85, 0x13, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenBurnList) > 0 {
		for iNdEx := len(m.TokenBurnList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenBurnList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Options != nil {
		{
			size, err := m.Options.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.MintpermList) > 0 {
		for iNdEx := len(m.MintpermList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintpermList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TokenList) > 0 {
		for iNdEx := len(m.TokenList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenList) > 0 {
		for _, e := range m.TokenList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MintpermList) > 0 {
		for _, e := range m.MintpermList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenBurnList) > 0 {
		for _, e := range m.TokenBurnList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenList = append(m.TokenList, Token{})
			if err := m.TokenList[len(m.TokenList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintpermList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintpermList = append(m.MintpermList, Mintperm{})
			if err := m.MintpermList[len(m.MintpermList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &Options{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenBurnList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenBurnList = append(m.TokenBurnList, TokenBurn{})
			if err := m.TokenBurnList[len(m.TokenBurnList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
