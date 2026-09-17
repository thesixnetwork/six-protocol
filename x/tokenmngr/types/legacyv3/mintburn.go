package legacyv3

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/jsonpb"
)

// MsgMint and MsgBurn carry two different on-chain wire layouts under the same
// legacy type URL /thesixnetwork.sixprotocol.tokenmngr.Msg{Mint,Burn}:
//
//	v1 (tags v2.0.x–v2.1.x): creator=1 string, amount=2 uint64, token=3 string
//	v2/v3 (tags v2.2.0+):    creator=1 string, amount=2 Coin
//
// The two are wire-distinguishable: in v1 field 2 is a varint (wire type 0);
// in v2/v3 it is a length-delimited Coin (wire type 2). A single custom
// Unmarshal branches on that wire type, so both eras decode without needing to
// know the block height. IsLegacy records which layout was found so display and
// re-marshaling stay faithful to the original bytes.
type mintBurn struct {
	Creator      string
	Amount       sdk.Coin // set for the v2/v3 (Coin) layout
	LegacyAmount uint64   // set for the v1 (uint64) layout
	Token        string   // set for the v1 layout only (field 3)
	IsLegacy     bool
}

func (m *mintBurn) Reset()         { *m = mintBurn{} }
func (m *mintBurn) ProtoMessage()  {}
func (m *mintBurn) String() string {
	if m.IsLegacy {
		return fmt.Sprintf("creator:%q amount:%d token:%q", m.Creator, m.LegacyAmount, m.Token)
	}
	return fmt.Sprintf("creator:%q amount:%s", m.Creator, m.Amount.String())
}

// Unmarshal decodes either the v1 (uint64) or the v2/v3 (Coin) layout.
func (m *mintBurn) Unmarshal(data []byte) error {
	m.Reset()
	i := 0
	for i < len(data) {
		key, n, err := readVarint(data, i)
		if err != nil {
			return err
		}
		i += n
		fieldNum := key >> 3
		wireType := key & 0x7
		switch {
		case fieldNum == 1 && wireType == 2: // creator
			s, ni, err := readBytes(data, i)
			if err != nil {
				return err
			}
			m.Creator = string(s)
			i = ni
		case fieldNum == 2 && wireType == 0: // v1 amount (uint64)
			v, n, err := readVarint(data, i)
			if err != nil {
				return err
			}
			m.LegacyAmount = v
			m.IsLegacy = true
			i += n
		case fieldNum == 2 && wireType == 2: // v2/v3 amount (Coin)
			s, ni, err := readBytes(data, i)
			if err != nil {
				return err
			}
			if err := m.Amount.Unmarshal(s); err != nil {
				return err
			}
			i = ni
		case fieldNum == 3 && wireType == 2: // v1 token
			s, ni, err := readBytes(data, i)
			if err != nil {
				return err
			}
			m.Token = string(s)
			m.IsLegacy = true
			i = ni
		default:
			ni, err := skipField(data, i, wireType)
			if err != nil {
				return err
			}
			i = ni
		}
	}
	return nil
}

func (m *mintBurn) Size() int {
	var n int
	if l := len(m.Creator); l > 0 {
		n += 1 + l + sovVarint(uint64(l))
	}
	if m.IsLegacy {
		n += 1 + sovVarint(m.LegacyAmount)
		if l := len(m.Token); l > 0 {
			n += 1 + l + sovVarint(uint64(l))
		}
		return n
	}
	l := m.Amount.Size()
	n += 1 + l + sovVarint(uint64(l))
	return n
}

// Marshal reproduces the original wire bytes for whichever layout was decoded.
func (m *mintBurn) Marshal() ([]byte, error) {
	dAtA := make([]byte, m.Size())
	i := len(dAtA)
	if m.IsLegacy {
		if len(m.Token) > 0 {
			i -= len(m.Token)
			copy(dAtA[i:], m.Token)
			i = encodeVarint(dAtA, i, uint64(len(m.Token)))
			i--
			dAtA[i] = 0x1a // field 3, wire type 2
		}
		i = encodeVarint(dAtA, i, m.LegacyAmount)
		i--
		dAtA[i] = 0x10 // field 2, wire type 0
	} else {
		sz := m.Amount.Size()
		i -= sz
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return nil, err
		}
		i = encodeVarint(dAtA, i, uint64(sz))
		i--
		dAtA[i] = 0x12 // field 2, wire type 2
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarint(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa // field 1, wire type 2
	}
	return dAtA[i:], nil
}

// MarshalJSONPB renders the original data so `sixd q tx` shows real values.
func (m *mintBurn) MarshalJSONPB(_ *jsonpb.Marshaler) ([]byte, error) {
	if m.IsLegacy {
		obj := map[string]string{"creator": m.Creator, "amount": strconv.FormatUint(m.LegacyAmount, 10)}
		if m.Token != "" {
			obj["token"] = m.Token
		}
		return json.Marshal(obj)
	}
	amount := map[string]string{"denom": m.Amount.Denom, "amount": intString(m.Amount.Amount)}
	return json.Marshal(map[string]interface{}{"creator": m.Creator, "amount": amount})
}

func intString(i sdkmath.Int) string {
	if i.IsNil() {
		return "0"
	}
	return i.String()
}

// MsgMint is the legacy (v1..v3) mint message.
type MsgMint struct{ mintBurn }

func (*MsgMint) XXX_MessageName() string { return legacyPkg + "MsgMint" }

// MsgBurn is the legacy (v1..v3) burn message.
type MsgBurn struct{ mintBurn }

func (*MsgBurn) XXX_MessageName() string { return legacyPkg + "MsgBurn" }

// --- minimal protobuf wire readers ---

func readVarint(data []byte, i int) (uint64, int, error) {
	var v uint64
	start := i
	for shift := uint(0); ; shift += 7 {
		if shift >= 64 {
			return 0, 0, errIntOverflow
		}
		if i >= len(data) {
			return 0, 0, fmt.Errorf("legacyv3: unexpected EOF reading varint")
		}
		b := data[i]
		i++
		v |= uint64(b&0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	return v, i - start, nil
}

func readBytes(data []byte, i int) ([]byte, int, error) {
	l, n, err := readVarint(data, i)
	if err != nil {
		return nil, 0, err
	}
	i += n
	end := i + int(l)
	if int(l) < 0 || end > len(data) {
		return nil, 0, fmt.Errorf("legacyv3: length-delimited field out of range")
	}
	return data[i:end], end, nil
}

func skipField(data []byte, i int, wireType uint64) (int, error) {
	switch wireType {
	case 0:
		_, n, err := readVarint(data, i)
		return i + n, err
	case 1:
		return i + 8, nil
	case 2:
		_, ni, err := readBytes(data, i)
		return ni, err
	case 5:
		return i + 4, nil
	default:
		return 0, fmt.Errorf("legacyv3: unsupported wire type %d", wireType)
	}
}
