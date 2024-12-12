// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/genesis.proto

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

// GenesisState defines the nftmngr module's genesis state.
type GenesisState struct {
	Params                   Params                        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	NFTSchemaList            []NFTSchema                   `protobuf:"bytes,2,rep,name=nFTSchemaList,proto3" json:"nFTSchemaList"`
	NftDataList              []NftData                     `protobuf:"bytes,3,rep,name=nftDataList,proto3" json:"nftDataList"`
	ActionByRefIdList        []ActionByRefId               `protobuf:"bytes,4,rep,name=actionByRefIdList,proto3" json:"actionByRefIdList"`
	OrganizationList         []Organization                `protobuf:"bytes,5,rep,name=organizationList,proto3" json:"organizationList"`
	NFTSchemaByContractList  []NFTSchemaByContract         `protobuf:"bytes,7,rep,name=nFTSchemaByContractList,proto3" json:"nFTSchemaByContractList"`
	NftFeeConfig             *NFTFeeConfig                 `protobuf:"bytes,8,opt,name=nft_fee_config,json=nftFeeConfig,proto3" json:"nft_fee_config,omitempty"`
	NFTFeeBalance            *NFTFeeBalance                `protobuf:"bytes,9,opt,name=nFTFeeBalance,proto3" json:"nFTFeeBalance,omitempty"`
	MetadataCreatorList      []MetadataCreator             `protobuf:"bytes,10,rep,name=metadataCreatorList,proto3" json:"metadataCreatorList"`
	NftCollectionList        []NftCollection               `protobuf:"bytes,11,rep,name=nftCollectionList,proto3" json:"nftCollectionList"`
	ActionExecutorList       []ActionExecutor              `protobuf:"bytes,12,rep,name=actionExecutorList,proto3" json:"actionExecutorList"`
	SchemaAttributeList      []SchemaAttribute             `protobuf:"bytes,13,rep,name=schemaAttributeList,proto3" json:"schemaAttributeList"`
	ActionOfSchemaList       []ActionOfSchema              `protobuf:"bytes,14,rep,name=actionOfSchemaList,proto3" json:"actionOfSchemaList"`
	ExecutorOfSchemaList     []ExecutorOfSchema            `protobuf:"bytes,15,rep,name=executorOfSchemaList,proto3" json:"executorOfSchemaList"`
	VirtualActionList        []VirtualAction               `protobuf:"bytes,16,rep,name=virtualActionList,proto3" json:"virtualActionList"`
	VirtualSchemaList        []VirtualSchema               `protobuf:"bytes,17,rep,name=virtualSchemaList,proto3" json:"virtualSchemaList"`
	DisableVirtualSchemaList []VirtualSchemaDisableRequest `protobuf:"bytes,18,rep,name=disableVirtualSchemaList,proto3" json:"disableVirtualSchemaList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c26d098aac64c1a, []int{0}
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

func (m *GenesisState) GetNFTSchemaList() []NFTSchema {
	if m != nil {
		return m.NFTSchemaList
	}
	return nil
}

func (m *GenesisState) GetNftDataList() []NftData {
	if m != nil {
		return m.NftDataList
	}
	return nil
}

func (m *GenesisState) GetActionByRefIdList() []ActionByRefId {
	if m != nil {
		return m.ActionByRefIdList
	}
	return nil
}

func (m *GenesisState) GetOrganizationList() []Organization {
	if m != nil {
		return m.OrganizationList
	}
	return nil
}

func (m *GenesisState) GetNFTSchemaByContractList() []NFTSchemaByContract {
	if m != nil {
		return m.NFTSchemaByContractList
	}
	return nil
}

func (m *GenesisState) GetNftFeeConfig() *NFTFeeConfig {
	if m != nil {
		return m.NftFeeConfig
	}
	return nil
}

func (m *GenesisState) GetNFTFeeBalance() *NFTFeeBalance {
	if m != nil {
		return m.NFTFeeBalance
	}
	return nil
}

func (m *GenesisState) GetMetadataCreatorList() []MetadataCreator {
	if m != nil {
		return m.MetadataCreatorList
	}
	return nil
}

func (m *GenesisState) GetNftCollectionList() []NftCollection {
	if m != nil {
		return m.NftCollectionList
	}
	return nil
}

func (m *GenesisState) GetActionExecutorList() []ActionExecutor {
	if m != nil {
		return m.ActionExecutorList
	}
	return nil
}

func (m *GenesisState) GetSchemaAttributeList() []SchemaAttribute {
	if m != nil {
		return m.SchemaAttributeList
	}
	return nil
}

func (m *GenesisState) GetActionOfSchemaList() []ActionOfSchema {
	if m != nil {
		return m.ActionOfSchemaList
	}
	return nil
}

func (m *GenesisState) GetExecutorOfSchemaList() []ExecutorOfSchema {
	if m != nil {
		return m.ExecutorOfSchemaList
	}
	return nil
}

func (m *GenesisState) GetVirtualActionList() []VirtualAction {
	if m != nil {
		return m.VirtualActionList
	}
	return nil
}

func (m *GenesisState) GetVirtualSchemaList() []VirtualSchema {
	if m != nil {
		return m.VirtualSchemaList
	}
	return nil
}

func (m *GenesisState) GetDisableVirtualSchemaList() []VirtualSchemaDisableRequest {
	if m != nil {
		return m.DisableVirtualSchemaList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "thesixnetwork.sixprotocol.nftmngr.GenesisState")
}

func init() { proto.RegisterFile("nftmngr/genesis.proto", fileDescriptor_0c26d098aac64c1a) }

var fileDescriptor_0c26d098aac64c1a = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4f, 0x4f, 0x13, 0x41,
	0x18, 0xc6, 0x5b, 0x41, 0xd4, 0x29, 0x20, 0x8c, 0xa8, 0x4d, 0x23, 0x2b, 0x1a, 0x0f, 0x68, 0xb4,
	0x8b, 0x10, 0x3d, 0x9a, 0xd0, 0xf2, 0x27, 0x26, 0x0a, 0xa6, 0x28, 0x31, 0x5e, 0x9a, 0xe9, 0xf6,
	0xdd, 0x65, 0xb4, 0x9d, 0xc1, 0xdd, 0x29, 0xb6, 0x9e, 0xfc, 0x08, 0x7e, 0x27, 0x2f, 0x1c, 0x39,
	0x7a, 0x32, 0x06, 0xbe, 0x88, 0xd9, 0xd9, 0x99, 0x65, 0xb6, 0xdb, 0xc6, 0xe9, 0x0d, 0xe6, 0x9d,
	0xe7, 0xf7, 0xcc, 0xbc, 0xfb, 0xcc, 0x5b, 0x74, 0x9b, 0xf9, 0xa2, 0xcb, 0x82, 0xd0, 0x0d, 0x80,
	0x41, 0x44, 0xa3, 0xea, 0x71, 0xc8, 0x05, 0xc7, 0x0f, 0xc4, 0x11, 0x44, 0xb4, 0xcf, 0x40, 0x7c,
	0xe3, 0xe1, 0x97, 0x6a, 0x44, 0xfb, 0x72, 0xdd, 0xe3, 0x9d, 0xaa, 0x12, 0x54, 0x96, 0x02, 0x1e,
	0x70, 0xb9, 0xea, 0xc6, 0x7f, 0x25, 0xc2, 0xca, 0x92, 0xe6, 0x1d, 0x93, 0x90, 0x74, 0x15, 0xae,
	0x52, 0xd6, 0xab, 0xcc, 0x17, 0xcd, 0xc8, 0x3b, 0x82, 0x2e, 0x51, 0x95, 0x3b, 0x66, 0xa5, 0x4d,
	0x84, 0x5e, 0x77, 0xf4, 0x3a, 0xf1, 0x04, 0xe5, 0xac, 0xd9, 0x1a, 0x34, 0x43, 0xf0, 0x9b, 0xb4,
	0xad, 0xea, 0x15, 0x5d, 0xe7, 0x61, 0x40, 0x18, 0xfd, 0x4e, 0xe2, 0x5d, 0xaa, 0xf6, 0x28, 0xef,
	0x16, 0xeb, 0x3d, 0xce, 0x44, 0x48, 0x3c, 0xa1, 0x76, 0xdd, 0x33, 0x77, 0xf9, 0x00, 0x71, 0xdd,
	0xa7, 0x81, 0xaa, 0x2e, 0x0f, 0x57, 0x5b, 0xa4, 0x43, 0x98, 0x07, 0xc3, 0xc7, 0xeb, 0x82, 0x20,
	0xf1, 0xb1, 0x9b, 0x5e, 0x08, 0x44, 0xf0, 0x70, 0x14, 0xdc, 0xe3, 0x9d, 0x0e, 0x78, 0xc6, 0x01,
	0x97, 0x87, 0x2e, 0x07, 0x7d, 0xf0, 0x7a, 0x97, 0xe2, 0x14, 0xae, 0xce, 0x4e, 0x84, 0x08, 0x69,
	0xab, 0x27, 0x60, 0x4c, 0x6f, 0xb8, 0x9f, 0xed, 0xe9, 0x8a, 0xae, 0x6b, 0x6e, 0x6e, 0x47, 0x7a,
	0xbc, 0x13, 0x1a, 0x8a, 0x1e, 0xe9, 0x34, 0x89, 0x79, 0xbc, 0x5c, 0xd5, 0xd4, 0x3e, 0xfc, 0x35,
	0x87, 0x66, 0x77, 0x93, 0xb0, 0x1c, 0x08, 0x22, 0x00, 0xef, 0xa2, 0x99, 0xe4, 0x63, 0x97, 0x8b,
	0x2b, 0xc5, 0xd5, 0xd2, 0xfa, 0xe3, 0xea, 0x7f, 0xc3, 0x53, 0x7d, 0x27, 0x05, 0xb5, 0xe9, 0xd3,
	0x3f, 0xf7, 0x0b, 0x0d, 0x25, 0xc7, 0x1f, 0xd1, 0x1c, 0xdb, 0x79, 0x7f, 0x20, 0xcd, 0xde, 0xd0,
	0x48, 0x94, 0xaf, 0xac, 0x4c, 0xad, 0x96, 0xd6, 0x9f, 0x5a, 0xf0, 0xf6, 0xb4, 0x4e, 0x21, 0xb3,
	0x20, 0xdc, 0x40, 0x25, 0xe6, 0x8b, 0x2d, 0x22, 0x12, 0xee, 0x94, 0xe4, 0x3e, 0xb1, 0xe1, 0x26,
	0x2a, 0x45, 0x35, 0x21, 0xb8, 0x8d, 0x16, 0x93, 0xae, 0xd5, 0x06, 0x0d, 0xf0, 0x5f, 0xb7, 0x25,
	0x79, 0x5a, 0x92, 0xd7, 0x2c, 0xc8, 0x9b, 0xa6, 0x56, 0xf1, 0xf3, 0x40, 0x4c, 0xd0, 0x82, 0x99,
	0x70, 0x69, 0x72, 0x55, 0x9a, 0xb8, 0x16, 0x26, 0xfb, 0x86, 0x54, 0x79, 0xe4, 0x70, 0xf8, 0x04,
	0xdd, 0x4d, 0xbb, 0x55, 0x1b, 0xd4, 0xd5, 0x2b, 0x91, 0x4e, 0xd7, 0xa4, 0xd3, 0xcb, 0x89, 0x3e,
	0x40, 0x4a, 0x50, 0x86, 0xe3, 0xe0, 0xf8, 0x03, 0x9a, 0xcf, 0x3e, 0xbd, 0xf2, 0x75, 0x99, 0x1f,
	0xd7, 0xce, 0x6e, 0x07, 0xa0, 0x2e, 0x65, 0x8d, 0x59, 0xe6, 0x8b, 0xf4, 0x3f, 0x7c, 0x28, 0x53,
	0xb4, 0x03, 0x50, 0x4b, 0x5e, 0x6c, 0xf9, 0x86, 0xa4, 0xae, 0x59, 0x53, 0x95, 0xae, 0x91, 0xc5,
	0xe0, 0xcf, 0xe8, 0x96, 0x7e, 0xec, 0xf5, 0xe4, 0xad, 0xcb, 0x16, 0x21, 0xd9, 0xa2, 0x75, 0x0b,
	0xfa, 0xdb, 0xac, 0x5a, 0xb5, 0x67, 0x14, 0x34, 0xce, 0x16, 0xf3, 0x45, 0x3d, 0x9d, 0x1b, 0xd2,
	0xa9, 0x64, 0x9d, 0xad, 0x3d, 0x53, 0xab, 0xb3, 0x95, 0x03, 0xe2, 0x00, 0xe1, 0x24, 0x70, 0xdb,
	0x6a, 0x4e, 0x48, 0x9b, 0x59, 0x69, 0xf3, 0xdc, 0x3a, 0xc2, 0x5a, 0xac, 0x7c, 0x46, 0x20, 0xe3,
	0xd6, 0x25, 0x23, 0x64, 0x53, 0x4f, 0x32, 0xe9, 0x34, 0x67, 0xdd, 0xba, 0x83, 0xac, 0x5a, 0xb7,
	0x6e, 0x04, 0xf4, 0xf2, 0x52, 0xfb, 0xbe, 0x31, 0x49, 0xe6, 0x27, 0xbc, 0x94, 0x16, 0x67, 0x2f,
	0x65, 0x22, 0x71, 0x17, 0x2d, 0xe9, 0xf9, 0x9a, 0xb1, 0xba, 0x29, 0xad, 0x36, 0x2c, 0xac, 0xb6,
	0x87, 0xe4, 0xca, 0x6c, 0x24, 0x36, 0x8e, 0x84, 0x1a, 0xc7, 0x9b, 0x97, 0x91, 0x58, 0xb0, 0x8e,
	0xc4, 0xa1, 0xa9, 0xd5, 0x91, 0xc8, 0x01, 0x0d, 0x17, 0xe3, 0x46, 0x8b, 0x93, 0xba, 0x64, 0xae,
	0x93, 0x07, 0xe2, 0x1f, 0x45, 0x54, 0x6e, 0xd3, 0x88, 0xb4, 0x3a, 0x70, 0x98, 0x73, 0xc3, 0xd2,
	0xed, 0xd5, 0xa4, 0x6e, 0x5b, 0x09, 0xaf, 0x01, 0x5f, 0x7b, 0x10, 0xe9, 0xd9, 0x33, 0xd6, 0xa5,
	0xb6, 0x7f, 0x7a, 0xee, 0x14, 0xcf, 0xce, 0x9d, 0xe2, 0xdf, 0x73, 0xa7, 0xf8, 0xf3, 0xc2, 0x29,
	0x9c, 0x5d, 0x38, 0x85, 0xdf, 0x17, 0x4e, 0xe1, 0xd3, 0x8b, 0x80, 0x8a, 0xa3, 0x5e, 0xab, 0xea,
	0xf1, 0xae, 0x9b, 0x39, 0x83, 0x1b, 0xd1, 0xfe, 0x33, 0x7d, 0x08, 0xb7, 0xef, 0xea, 0x5f, 0x49,
	0x31, 0x38, 0x86, 0xa8, 0x35, 0x23, 0x2b, 0x1b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x36,
	0x9c, 0x18, 0x51, 0x09, 0x00, 0x00,
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
	if len(m.DisableVirtualSchemaList) > 0 {
		for iNdEx := len(m.DisableVirtualSchemaList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DisableVirtualSchemaList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x92
		}
	}
	if len(m.VirtualSchemaList) > 0 {
		for iNdEx := len(m.VirtualSchemaList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VirtualSchemaList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.VirtualActionList) > 0 {
		for iNdEx := len(m.VirtualActionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VirtualActionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.ExecutorOfSchemaList) > 0 {
		for iNdEx := len(m.ExecutorOfSchemaList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExecutorOfSchemaList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.ActionOfSchemaList) > 0 {
		for iNdEx := len(m.ActionOfSchemaList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionOfSchemaList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.SchemaAttributeList) > 0 {
		for iNdEx := len(m.SchemaAttributeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SchemaAttributeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.ActionExecutorList) > 0 {
		for iNdEx := len(m.ActionExecutorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionExecutorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.NftCollectionList) > 0 {
		for iNdEx := len(m.NftCollectionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftCollectionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.MetadataCreatorList) > 0 {
		for iNdEx := len(m.MetadataCreatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MetadataCreatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.NFTFeeBalance != nil {
		{
			size, err := m.NFTFeeBalance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.NftFeeConfig != nil {
		{
			size, err := m.NftFeeConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.NFTSchemaByContractList) > 0 {
		for iNdEx := len(m.NFTSchemaByContractList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NFTSchemaByContractList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.OrganizationList) > 0 {
		for iNdEx := len(m.OrganizationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrganizationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ActionByRefIdList) > 0 {
		for iNdEx := len(m.ActionByRefIdList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionByRefIdList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.NftDataList) > 0 {
		for iNdEx := len(m.NftDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.NFTSchemaList) > 0 {
		for iNdEx := len(m.NFTSchemaList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NFTSchemaList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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
	if len(m.NFTSchemaList) > 0 {
		for _, e := range m.NFTSchemaList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NftDataList) > 0 {
		for _, e := range m.NftDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionByRefIdList) > 0 {
		for _, e := range m.ActionByRefIdList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrganizationList) > 0 {
		for _, e := range m.OrganizationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NFTSchemaByContractList) > 0 {
		for _, e := range m.NFTSchemaByContractList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NftFeeConfig != nil {
		l = m.NftFeeConfig.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NFTFeeBalance != nil {
		l = m.NFTFeeBalance.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MetadataCreatorList) > 0 {
		for _, e := range m.MetadataCreatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NftCollectionList) > 0 {
		for _, e := range m.NftCollectionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionExecutorList) > 0 {
		for _, e := range m.ActionExecutorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SchemaAttributeList) > 0 {
		for _, e := range m.SchemaAttributeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionOfSchemaList) > 0 {
		for _, e := range m.ActionOfSchemaList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExecutorOfSchemaList) > 0 {
		for _, e := range m.ExecutorOfSchemaList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VirtualActionList) > 0 {
		for _, e := range m.VirtualActionList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VirtualSchemaList) > 0 {
		for _, e := range m.VirtualSchemaList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DisableVirtualSchemaList) > 0 {
		for _, e := range m.DisableVirtualSchemaList {
			l = e.Size()
			n += 2 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field NFTSchemaList", wireType)
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
			m.NFTSchemaList = append(m.NFTSchemaList, NFTSchema{})
			if err := m.NFTSchemaList[len(m.NFTSchemaList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftDataList", wireType)
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
			m.NftDataList = append(m.NftDataList, NftData{})
			if err := m.NftDataList[len(m.NftDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionByRefIdList", wireType)
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
			m.ActionByRefIdList = append(m.ActionByRefIdList, ActionByRefId{})
			if err := m.ActionByRefIdList[len(m.ActionByRefIdList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationList", wireType)
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
			m.OrganizationList = append(m.OrganizationList, Organization{})
			if err := m.OrganizationList[len(m.OrganizationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTSchemaByContractList", wireType)
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
			m.NFTSchemaByContractList = append(m.NFTSchemaByContractList, NFTSchemaByContract{})
			if err := m.NFTSchemaByContractList[len(m.NFTSchemaByContractList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftFeeConfig", wireType)
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
			if m.NftFeeConfig == nil {
				m.NftFeeConfig = &NFTFeeConfig{}
			}
			if err := m.NftFeeConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTFeeBalance", wireType)
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
			if m.NFTFeeBalance == nil {
				m.NFTFeeBalance = &NFTFeeBalance{}
			}
			if err := m.NFTFeeBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataCreatorList", wireType)
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
			m.MetadataCreatorList = append(m.MetadataCreatorList, MetadataCreator{})
			if err := m.MetadataCreatorList[len(m.MetadataCreatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftCollectionList", wireType)
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
			m.NftCollectionList = append(m.NftCollectionList, NftCollection{})
			if err := m.NftCollectionList[len(m.NftCollectionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionExecutorList", wireType)
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
			m.ActionExecutorList = append(m.ActionExecutorList, ActionExecutor{})
			if err := m.ActionExecutorList[len(m.ActionExecutorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaAttributeList", wireType)
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
			m.SchemaAttributeList = append(m.SchemaAttributeList, SchemaAttribute{})
			if err := m.SchemaAttributeList[len(m.SchemaAttributeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionOfSchemaList", wireType)
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
			m.ActionOfSchemaList = append(m.ActionOfSchemaList, ActionOfSchema{})
			if err := m.ActionOfSchemaList[len(m.ActionOfSchemaList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorOfSchemaList", wireType)
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
			m.ExecutorOfSchemaList = append(m.ExecutorOfSchemaList, ExecutorOfSchema{})
			if err := m.ExecutorOfSchemaList[len(m.ExecutorOfSchemaList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualActionList", wireType)
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
			m.VirtualActionList = append(m.VirtualActionList, VirtualAction{})
			if err := m.VirtualActionList[len(m.VirtualActionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualSchemaList", wireType)
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
			m.VirtualSchemaList = append(m.VirtualSchemaList, VirtualSchema{})
			if err := m.VirtualSchemaList[len(m.VirtualSchemaList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableVirtualSchemaList", wireType)
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
			m.DisableVirtualSchemaList = append(m.DisableVirtualSchemaList, VirtualSchemaDisableRequest{})
			if err := m.DisableVirtualSchemaList[len(m.DisableVirtualSchemaList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
