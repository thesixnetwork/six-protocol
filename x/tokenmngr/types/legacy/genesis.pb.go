// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/v1/genesis.proto

package legacy

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
	return fileDescriptor_d709ea449ee46bdf, []int{0}
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
	proto.RegisterType((*GenesisState)(nil), "thesixnetwork.sixprotocol.tokenmngr.GenesisState")
}

func init() { proto.RegisterFile("tokenmngr/v1/genesis.proto", fileDescriptor_d709ea449ee46bdf) }

var fileDescriptor_d709ea449ee46bdf = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0xc3, 0x30,
	0x1c, 0xc6, 0x5b, 0x37, 0x3a, 0x96, 0xcd, 0x4b, 0x11, 0xac, 0x15, 0xeb, 0xd0, 0xcb, 0x50, 0x97,
	0xb0, 0x79, 0xf3, 0xb8, 0x83, 0x32, 0xf0, 0x8d, 0x39, 0x10, 0x76, 0x19, 0x7b, 0x09, 0x5d, 0x98,
	0x6d, 0x4a, 0x92, 0xce, 0xfa, 0x2d, 0xfc, 0x46, 0x5e, 0x77, 0xdc, 0xd1, 0x93, 0xc8, 0xf6, 0x45,
	0x64, 0x69, 0xea, 0x08, 0x5e, 0xea, 0x2d, 0xc9, 0xc3, 0xf3, 0x7b, 0xfe, 0xcf, 0x9f, 0x00, 0x57,
	0xd0, 0x19, 0x0e, 0x83, 0xd0, 0x67, 0x68, 0xde, 0x44, 0x3e, 0x0e, 0x31, 0x27, 0x1c, 0x46, 0x8c,
	0x0a, 0x6a, 0x9f, 0x8a, 0x29, 0xe6, 0x24, 0x09, 0xb1, 0x78, 0xa5, 0x6c, 0x06, 0x39, 0x49, 0xe4,
	0xfb, 0x98, 0xbe, 0xc0, 0x5f, 0x97, 0xbb, 0xe7, 0x53, 0x9f, 0xca, 0x77, 0xb4, 0x39, 0xa5, 0x56,
	0xf7, 0x40, 0xc3, 0x46, 0x43, 0x36, 0x0c, 0x14, 0xd5, 0x75, 0x34, 0x49, 0x5e, 0x94, 0x72, 0xa8,
	0x29, 0x01, 0x09, 0x45, 0x84, 0x59, 0xa0, 0x44, 0x7d, 0x50, 0x1a, 0x09, 0x42, 0xc3, 0x0c, 0x79,
	0xf4, 0x17, 0x39, 0x18, 0xc5, 0x4c, 0x71, 0x4f, 0x3e, 0x0a, 0xa0, 0x7a, 0x93, 0x36, 0x7b, 0x12,
	0x43, 0x81, 0xed, 0x0e, 0xb0, 0xd2, 0x91, 0x1c, 0xb3, 0x66, 0xd6, 0x2b, 0xad, 0x73, 0x98, 0xa3,
	0x29, 0x7c, 0x94, 0x96, 0x76, 0x71, 0xf1, 0x75, 0x6c, 0x74, 0x15, 0xc0, 0xde, 0x07, 0xa5, 0x88,
	0x32, 0x31, 0x20, 0x13, 0x67, 0xa7, 0x66, 0xd6, 0xcb, 0x5d, 0x6b, 0x73, 0xed, 0x4c, 0xec, 0x7b,
	0x50, 0x96, 0xd6, 0x5b, 0xc2, 0x85, 0x53, 0xa8, 0x15, 0xea, 0x95, 0xd6, 0x59, 0xae, 0x98, 0xde,
	0xe6, 0xa4, 0x52, 0xb6, 0x08, 0xfb, 0x19, 0x54, 0xb3, 0x8d, 0x48, 0x64, 0x51, 0x22, 0x1b, 0xb9,
	0x90, 0x77, 0xca, 0xa8, 0xa8, 0x1a, 0xc8, 0xbe, 0x06, 0x25, 0xb5, 0x4d, 0xc7, 0x92, 0xdb, 0xb8,
	0xc8, 0xc5, 0x7c, 0x48, 0x3d, 0xdd, 0xcc, 0x6c, 0xf7, 0xc1, 0xae, 0x54, 0xdb, 0x31, 0x4b, 0x4b,
	0x97, 0xe4, 0x84, 0xf0, 0x1f, 0xa5, 0x63, 0x96, 0x15, 0xd7, 0x51, 0xed, 0xde, 0x62, 0xe5, 0x99,
	0xcb, 0x95, 0x67, 0x7e, 0xaf, 0x3c, 0xf3, 0x7d, 0xed, 0x19, 0xcb, 0xb5, 0x67, 0x7c, 0xae, 0x3d,
	0xa3, 0x7f, 0xe5, 0x13, 0x31, 0x8d, 0x47, 0x70, 0x4c, 0x03, 0xa4, 0x05, 0x21, 0x4e, 0x92, 0x46,
	0x96, 0x84, 0x12, 0xb4, 0xfd, 0x22, 0xe2, 0x2d, 0xc2, 0x1c, 0xcd, 0x9b, 0x23, 0x4b, 0xca, 0x97,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x6e, 0xf2, 0x6f, 0x04, 0x03, 0x00, 0x00,
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
