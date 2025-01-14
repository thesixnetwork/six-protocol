package types

import (
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (p *ActionParameter) MustGetNumber(key string) (uint64, error) {
	v, err := strconv.ParseUint(p.Value, 10, 64)
	if err != nil {
		return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, key)
	}
	return v, nil
}

func (p *ActionParameter) GetNumber() uint64 {
	v, err := p.MustGetNumber(p.Name)
	if err != nil {
		panic(err)
	}
	return v
}

func (p *ActionParameter) MustGetFloat(key string) (float64, error) {
	v, err := strconv.ParseFloat(p.Value, 64)
	if err != nil {
		return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, key)
	}
	return v, nil
}

func (p *ActionParameter) GetFloat() float64 {
	v, err := p.MustGetFloat(p.Name)
	if err != nil {
		panic(err)
	}
	return v
}

func (p *ActionParameter) MustGetBool(key string) (bool, error) {
	v, err := strconv.ParseBool(p.Value)
	if err != nil {
		return false, sdkerrors.Wrap(ErrAttributeTypeNotMatch, key)
	}
	return v, nil
}

func (p *ActionParameter) GetBoolean() bool {
	v, err := p.MustGetBool(p.Name)
	if err != nil {
		panic(err)
	}
	return v
}

func (p *ActionParameter) GetString() string {
	return p.Value
}

// return substring of string from start to end of parameter
func (p *ActionParameter) GetSubString(start int64, end int64) string {
	val := p.Value
	if end > int64(len(val)) {
		panic(sdkerrors.Wrap(ErrInvaliActionParameter, "end can not be greater than string length"))
	}
	if start == end {
		return ""
	}
	if start < 0 {
		start = int64(len(val)) + (start + 1)
	}
	if end < 0 {
		end = int64(len(val)) + (end + 1)
	}
	if start > end {
		panic(sdkerrors.Wrap(ErrInvaliActionParameter, "start can not be greater than end"))
	}
	return val[start:end]
}

// return LowerCase of parameter
func (p *ActionParameter) ToLowerCase() string {
	return strings.ToLower(p.Value)
}

// return UpperCase of parameter
func (p *ActionParameter) ToUpperCase() string {
	return strings.ToUpper(p.Value)
}
