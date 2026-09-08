package legacyv3

import (
	fmt "fmt"
	io "io"

	proto "github.com/cosmos/gogoproto/proto"
)

// MsgCreateOptions and MsgUpdateOptions changed field numbers between v3 and
// v4: `defaultMintee` was field 2 in v3 and became field 3 in v4. Embedding the
// current type would silently drop defaultMintee when decoding old txs, so
// these types carry the v3 field layout (creator=1, defaultMintee=2) with their
// own marshaling.

type MsgCreateOptions struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	DefaultMintee string `protobuf:"bytes,2,opt,name=defaultMintee,proto3" json:"defaultMintee,omitempty"`
}

func (m *MsgCreateOptions) Reset()               { *m = MsgCreateOptions{} }
func (m *MsgCreateOptions) String() string       { return proto.CompactTextString(m) }
func (*MsgCreateOptions) ProtoMessage()          {}
func (*MsgCreateOptions) XXX_MessageName() string { return legacyPkg + "MsgCreateOptions" }
func (m *MsgCreateOptions) Marshal() ([]byte, error) {
	return marshalOptions(m.Creator, m.DefaultMintee)
}
func (m *MsgCreateOptions) Unmarshal(dAtA []byte) error {
	return unmarshalOptions(dAtA, &m.Creator, &m.DefaultMintee)
}
func (m *MsgCreateOptions) Size() int { return sizeOptions(m.Creator, m.DefaultMintee) }

type MsgUpdateOptions struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	DefaultMintee string `protobuf:"bytes,2,opt,name=defaultMintee,proto3" json:"defaultMintee,omitempty"`
}

func (m *MsgUpdateOptions) Reset()               { *m = MsgUpdateOptions{} }
func (m *MsgUpdateOptions) String() string       { return proto.CompactTextString(m) }
func (*MsgUpdateOptions) ProtoMessage()          {}
func (*MsgUpdateOptions) XXX_MessageName() string { return legacyPkg + "MsgUpdateOptions" }
func (m *MsgUpdateOptions) Marshal() ([]byte, error) {
	return marshalOptions(m.Creator, m.DefaultMintee)
}
func (m *MsgUpdateOptions) Unmarshal(dAtA []byte) error {
	return unmarshalOptions(dAtA, &m.Creator, &m.DefaultMintee)
}
func (m *MsgUpdateOptions) Size() int { return sizeOptions(m.Creator, m.DefaultMintee) }

// --- shared proto3 marshaling for {creator=1, defaultMintee=2} ---

func sizeOptions(creator, defaultMintee string) int {
	var n int
	if l := len(creator); l > 0 {
		n += 1 + l + sovVarint(uint64(l))
	}
	if l := len(defaultMintee); l > 0 {
		n += 1 + l + sovVarint(uint64(l))
	}
	return n
}

func marshalOptions(creator, defaultMintee string) ([]byte, error) {
	size := sizeOptions(creator, defaultMintee)
	dAtA := make([]byte, size)
	i := len(dAtA)
	// Encoded in reverse (gogoproto MarshalToSizedBuffer convention).
	if len(defaultMintee) > 0 {
		i -= len(defaultMintee)
		copy(dAtA[i:], defaultMintee)
		i = encodeVarint(dAtA, i, uint64(len(defaultMintee)))
		i--
		dAtA[i] = 0x12 // field 2, wire type 2
	}
	if len(creator) > 0 {
		i -= len(creator)
		copy(dAtA[i:], creator)
		i = encodeVarint(dAtA, i, uint64(len(creator)))
		i--
		dAtA[i] = 0xa // field 1, wire type 2
	}
	return dAtA[i:], nil
}

func unmarshalOptions(dAtA []byte, creator, defaultMintee *string) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		// read field key varint
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return errIntOverflow
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
		if wireType != 2 {
			// only string fields expected; skip anything else defensively
			return fmt.Errorf("legacyv3: unexpected wire type %d for field %d", wireType, fieldNum)
		}
		// read length
		var strLen uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return errIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			strLen |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		if strLen > uint64(l-iNdEx) {
			return io.ErrUnexpectedEOF
		}
		postIndex := iNdEx + int(strLen)
		switch fieldNum {
		case 1:
			*creator = string(dAtA[iNdEx:postIndex])
		case 2:
			*defaultMintee = string(dAtA[iNdEx:postIndex])
		}
		iNdEx = postIndex
	}
	return nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sovVarint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovVarint(x uint64) (n int) {
	return (bits64Len(x|1) + 6) / 7
}

func bits64Len(x uint64) int {
	n := 0
	for x != 0 {
		n++
		x >>= 1
	}
	return n
}

var errIntOverflow = fmt.Errorf("legacyv3: integer overflow")
