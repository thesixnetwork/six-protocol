// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v080/nft_collection.proto

package v080

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

type NftCollection struct {
	NftSchemaCode string     `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Total         uint64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	NftDatas      []*NftData `protobuf:"bytes,3,rep,name=nftDatas,proto3" json:"nftDatas,omitempty"`
}

func (m *NftCollection) Reset()         { *m = NftCollection{} }
func (m *NftCollection) String() string { return proto.CompactTextString(m) }
func (*NftCollection) ProtoMessage()    {}
func (*NftCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a42ca52407abed3, []int{0}
}
func (m *NftCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftCollection.Merge(m, src)
}
func (m *NftCollection) XXX_Size() int {
	return m.Size()
}
func (m *NftCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_NftCollection.DiscardUnknown(m)
}

var xxx_messageInfo_NftCollection proto.InternalMessageInfo

func (m *NftCollection) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *NftCollection) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *NftCollection) GetNftDatas() []*NftData {
	if m != nil {
		return m.NftDatas
	}
	return nil
}

func init() {
	proto.RegisterType((*NftCollection)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.NftCollection")
}

func init() { proto.RegisterFile("nftmngr/v080/nft_collection.proto", fileDescriptor_3a42ca52407abed3) }

var fileDescriptor_3a42ca52407abed3 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xb0, 0x30, 0xd0, 0xcf, 0x4b, 0x2b, 0x89, 0x4f, 0xce, 0xcf,
	0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2b,
	0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0xce, 0xac,
	0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0x41, 0x35, 0xeb, 0x81, 0x34, 0x4b, 0x49, 0x63, 0x18, 0x95,
	0x92, 0x58, 0x92, 0x08, 0x31, 0x44, 0x69, 0x06, 0x23, 0x17, 0xaf, 0x5f, 0x5a, 0x89, 0x33, 0xdc,
	0x70, 0x21, 0x15, 0x2e, 0xde, 0xbc, 0xb4, 0x92, 0xe0, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0xe7, 0xfc,
	0x94, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x54, 0x41, 0x21, 0x11, 0x2e, 0xd6, 0x92,
	0xfc, 0x92, 0xc4, 0x1c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0xc8, 0x9b, 0x8b,
	0x23, 0x2f, 0xad, 0xc4, 0x25, 0xb1, 0x24, 0xb1, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48,
	0x5f, 0x8f, 0x38, 0x57, 0xea, 0xf9, 0x41, 0xf4, 0x05, 0xc1, 0x0d, 0x70, 0x0a, 0x39, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xab, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0x14, 0xe3, 0xf5, 0x8b, 0x33, 0x2b, 0x74, 0x61, 0xe6, 0xeb, 0x57, 0xe8,
	0xc3, 0x7c, 0x5e, 0x52, 0x59, 0x90, 0x5a, 0x0c, 0xf6, 0x7f, 0x12, 0x1b, 0x58, 0xda, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x6d, 0x8f, 0x5c, 0x61, 0x01, 0x00, 0x00,
}

func (m *NftCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftDatas) > 0 {
		for iNdEx := len(m.NftDatas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftDatas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Total != 0 {
		i = encodeVarintNftCollection(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovNftCollection(uint64(m.Total))
	}
	if len(m.NftDatas) > 0 {
		for _, e := range m.NftDatas {
			l = e.Size()
			n += 1 + l + sovNftCollection(uint64(l))
		}
	}
	return n
}

func sovNftCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftCollection(x uint64) (n int) {
	return sovNftCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftCollection
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
			return fmt.Errorf("proto: NftCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftDatas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftDatas = append(m.NftDatas, &NftData{})
			if err := m.NftDatas[len(m.NftDatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftCollection
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
func skipNftCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftCollection
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
					return 0, ErrIntOverflowNftCollection
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
					return 0, ErrIntOverflowNftCollection
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
				return 0, ErrInvalidLengthNftCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftCollection = fmt.Errorf("proto: unexpected end of group")
)
