/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../nftmngr/params";
import { NFTSchemaQueryResult } from "../nftmngr/nft_schema";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { NftData } from "../nftmngr/nft_data";
import { ActionByRefId } from "../nftmngr/action_by_ref_id";
import { Organization } from "../nftmngr/organization";
import { NFTSchemaByContract } from "../nftmngr/nft_schema_by_contract";
import { NFTFeeConfig } from "../nftmngr/nft_fee_config";
import { NFTFeeBalance } from "../nftmngr/nft_fee_balance";
import { MetadataCreator } from "../nftmngr/metadata_creator";
import { ActionExecutor } from "../nftmngr/action_executor";
import { SchemaAttribute } from "../nftmngr/schema_attribute";
import { ActionOfSchema } from "../nftmngr/action_of_schema";
import { ExecutorOfSchema } from "../nftmngr/executor_of_schema";
import { VirtualAction } from "../nftmngr/virtual_action";
import {
  VirtualSchema,
  VirtualSchemaProposal,
} from "../nftmngr/virtual_schema";
import { LockSchemaFee } from "../nftmngr/lock_schema_fee";

export const protobufPackage = "thesixnetwork.sixprotocol.nftmngr";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNFTSchemaRequest {
  code: string;
}

export interface QueryGetNFTSchemaResponse {
  nFTSchema: NFTSchemaQueryResult | undefined;
}

export interface QueryAllNFTSchemaRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNFTSchemaResponse {
  nFTSchema: NFTSchemaQueryResult[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNftDataRequest {
  nftSchemaCode: string;
  tokenId: string;
  withGlobal: boolean;
}

export interface QueryGetNftDataResponse {
  nftData: NftData | undefined;
}

export interface QueryAllNftDataRequest {
  withGlobal: boolean;
  pagination: PageRequest | undefined;
}

export interface QueryAllNftDataResponse {
  nftData: NftData[];
  pagination: PageResponse | undefined;
}

export interface QueryGetActionByRefIdRequest {
  refId: string;
}

export interface QueryGetActionByRefIdResponse {
  actionByRefId: ActionByRefId | undefined;
}

export interface QueryAllActionByRefIdRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllActionByRefIdResponse {
  actionByRefId: ActionByRefId[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOrganizationRequest {
  name: string;
}

export interface QueryGetOrganizationResponse {
  organization: Organization | undefined;
}

export interface QueryAllOrganizationRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOrganizationResponse {
  organization: Organization[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNftCollectionRequest {
  nftSchemaCode: string;
  pagination: PageRequest | undefined;
}

export interface QueryGetNftCollectionResponse {
  nftCollection: NftData[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNFTSchemaByContractRequest {
  originContractAddress: string;
}

export interface QueryGetNFTSchemaByContractResponse {
  nFTSchemaByContract: NFTSchemaByContract | undefined;
}

export interface QueryAllNFTSchemaByContractRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNFTSchemaByContractResponse {
  nFTSchemaByContract: NFTSchemaByContract[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNFTFeeConfigRequest {}

export interface QueryGetNFTFeeConfigResponse {
  NFTFeeConfig: NFTFeeConfig | undefined;
}

export interface QueryGetNFTFeeBalanceRequest {}

export interface QueryGetNFTFeeBalanceResponse {
  NFTFeeBalance: NFTFeeBalance | undefined;
}

export interface QueryGetMetadataCreatorRequest {
  nftSchemaCode: string;
}

export interface QueryGetMetadataCreatorResponse {
  metadataCreator: MetadataCreator | undefined;
}

export interface QueryAllMetadataCreatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMetadataCreatorResponse {
  metadataCreator: MetadataCreator[];
  pagination: PageResponse | undefined;
}

export interface QueryGetActionExecutorRequest {
  nftSchemaCode: string;
  executorAddress: string;
}

export interface QueryGetActionExecutorResponse {
  actionExecutor: ActionExecutor | undefined;
}

export interface QueryAllActionExecutorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllActionExecutorResponse {
  actionExecutor: ActionExecutor[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSchemaAttributeRequest {
  nftSchemaCode: string;
  name: string;
}

export interface QueryGetSchemaAttributeResponse {
  schemaAttribute: SchemaAttribute | undefined;
}

export interface QueryAllSchemaAttributeRequest {
  nftSchemaCode: string;
  pagination: PageRequest | undefined;
}

export interface QueryAllSchemaAttributeResponse {
  schemaAttribute: SchemaAttribute[];
  pagination: PageResponse | undefined;
}

export interface QueryListAttributeBySchemaRequest {
  nftSchemaCode: string;
}

export interface QueryListAttributeBySchemaResponse {
  schemaAttribute: SchemaAttribute[];
}

export interface QueryGetActionOfSchemaRequest {
  nftSchemaCode: string;
  name: string;
}

export interface QueryGetActionOfSchemaResponse {
  actionOfSchema: ActionOfSchema | undefined;
}

export interface QueryAllActionOfSchemaRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllActionOfSchemaResponse {
  actionOfSchema: ActionOfSchema[];
  pagination: PageResponse | undefined;
}

export interface QueryGetExecutorOfSchemaRequest {
  nftSchemaCode: string;
}

export interface QueryGetExecutorOfSchemaResponse {
  executorOfSchema: ExecutorOfSchema | undefined;
}

export interface QueryAllExecutorOfSchemaRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllExecutorOfSchemaResponse {
  executorOfSchema: ExecutorOfSchema[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVirtualActionRequest {
  nftSchemaCode: string;
  name: string;
}

export interface QueryGetVirtualActionResponse {
  virtualAction: VirtualAction | undefined;
}

export interface QueryAllVirtualActionRequest {
  nftSchemaCode: string;
  pagination: PageRequest | undefined;
}

export interface QueryAllVirtualActionResponse {
  virtualAction: VirtualAction[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVirtualSchemaRequest {
  nftSchemaCode: string;
}

export interface QueryGetVirtualSchemaResponse {
  virtualSchema: VirtualSchema | undefined;
}

export interface QueryAllVirtualSchemaRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVirtualSchemaResponse {
  virtualSchema: VirtualSchema[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVirtualSchemaProposalRequest {
  index: string;
}

export interface QueryGetVirtualSchemaProposalResponse {
  virtualSchemaProposal: VirtualSchemaProposal | undefined;
}

export interface QueryAllVirtualSchemaProposalRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVirtualSchemaProposalResponse {
  virtualSchemaProposal: VirtualSchemaProposal[];
  pagination: PageResponse | undefined;
}

export interface QueryListActiveProposalRequest {}

export interface QueryListActiveProposalResponse {
  virtualSchemaProposal: VirtualSchemaProposal[];
}

export interface QueryGetLockSchemaFeeRequest {
  index: string;
}

export interface QueryGetLockSchemaFeeResponse {
  lockSchemaFee: LockSchemaFee | undefined;
}

export interface QueryAllLockSchemaFeeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLockSchemaFeeResponse {
  lockSchemaFee: LockSchemaFee[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetNFTSchemaRequest: object = { code: "" };

export const QueryGetNFTSchemaRequest = {
  encode(
    message: QueryGetNFTSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaRequest {
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaRequest): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaRequest>
  ): QueryGetNFTSchemaRequest {
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    return message;
  },
};

const baseQueryGetNFTSchemaResponse: object = {};

export const QueryGetNFTSchemaResponse = {
  encode(
    message: QueryGetNFTSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nFTSchema !== undefined) {
      NFTSchemaQueryResult.encode(
        message.nFTSchema,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchema = NFTSchemaQueryResult.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaResponse {
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      message.nFTSchema = NFTSchemaQueryResult.fromJSON(object.nFTSchema);
    } else {
      message.nFTSchema = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaResponse): unknown {
    const obj: any = {};
    message.nFTSchema !== undefined &&
      (obj.nFTSchema = message.nFTSchema
        ? NFTSchemaQueryResult.toJSON(message.nFTSchema)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaResponse>
  ): QueryGetNFTSchemaResponse {
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      message.nFTSchema = NFTSchemaQueryResult.fromPartial(object.nFTSchema);
    } else {
      message.nFTSchema = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaRequest: object = {};

export const QueryAllNFTSchemaRequest = {
  encode(
    message: QueryAllNFTSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaRequest {
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaRequest>
  ): QueryAllNFTSchemaRequest {
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaResponse: object = {};

export const QueryAllNFTSchemaResponse = {
  encode(
    message: QueryAllNFTSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nFTSchema) {
      NFTSchemaQueryResult.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchema.push(
            NFTSchemaQueryResult.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaResponse {
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      for (const e of object.nFTSchema) {
        message.nFTSchema.push(NFTSchemaQueryResult.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaResponse): unknown {
    const obj: any = {};
    if (message.nFTSchema) {
      obj.nFTSchema = message.nFTSchema.map((e) =>
        e ? NFTSchemaQueryResult.toJSON(e) : undefined
      );
    } else {
      obj.nFTSchema = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaResponse>
  ): QueryAllNFTSchemaResponse {
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      for (const e of object.nFTSchema) {
        message.nFTSchema.push(NFTSchemaQueryResult.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNftDataRequest: object = {
  nftSchemaCode: "",
  tokenId: "",
  withGlobal: false,
};

export const QueryGetNftDataRequest = {
  encode(
    message: QueryGetNftDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.withGlobal === true) {
      writer.uint32(24).bool(message.withGlobal);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNftDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNftDataRequest } as QueryGetNftDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        case 3:
          message.withGlobal = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftDataRequest {
    const message = { ...baseQueryGetNftDataRequest } as QueryGetNftDataRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.withGlobal !== undefined && object.withGlobal !== null) {
      message.withGlobal = Boolean(object.withGlobal);
    } else {
      message.withGlobal = false;
    }
    return message;
  },

  toJSON(message: QueryGetNftDataRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.withGlobal !== undefined && (obj.withGlobal = message.withGlobal);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftDataRequest>
  ): QueryGetNftDataRequest {
    const message = { ...baseQueryGetNftDataRequest } as QueryGetNftDataRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.withGlobal !== undefined && object.withGlobal !== null) {
      message.withGlobal = object.withGlobal;
    } else {
      message.withGlobal = false;
    }
    return message;
  },
};

const baseQueryGetNftDataResponse: object = {};

export const QueryGetNftDataResponse = {
  encode(
    message: QueryGetNftDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftData !== undefined) {
      NftData.encode(message.nftData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNftDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNftDataResponse,
    } as QueryGetNftDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftData = NftData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftDataResponse {
    const message = {
      ...baseQueryGetNftDataResponse,
    } as QueryGetNftDataResponse;
    if (object.nftData !== undefined && object.nftData !== null) {
      message.nftData = NftData.fromJSON(object.nftData);
    } else {
      message.nftData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNftDataResponse): unknown {
    const obj: any = {};
    message.nftData !== undefined &&
      (obj.nftData = message.nftData
        ? NftData.toJSON(message.nftData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftDataResponse>
  ): QueryGetNftDataResponse {
    const message = {
      ...baseQueryGetNftDataResponse,
    } as QueryGetNftDataResponse;
    if (object.nftData !== undefined && object.nftData !== null) {
      message.nftData = NftData.fromPartial(object.nftData);
    } else {
      message.nftData = undefined;
    }
    return message;
  },
};

const baseQueryAllNftDataRequest: object = { withGlobal: false };

export const QueryAllNftDataRequest = {
  encode(
    message: QueryAllNftDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.withGlobal === true) {
      writer.uint32(8).bool(message.withGlobal);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNftDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNftDataRequest } as QueryAllNftDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.withGlobal = reader.bool();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNftDataRequest {
    const message = { ...baseQueryAllNftDataRequest } as QueryAllNftDataRequest;
    if (object.withGlobal !== undefined && object.withGlobal !== null) {
      message.withGlobal = Boolean(object.withGlobal);
    } else {
      message.withGlobal = false;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNftDataRequest): unknown {
    const obj: any = {};
    message.withGlobal !== undefined && (obj.withGlobal = message.withGlobal);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNftDataRequest>
  ): QueryAllNftDataRequest {
    const message = { ...baseQueryAllNftDataRequest } as QueryAllNftDataRequest;
    if (object.withGlobal !== undefined && object.withGlobal !== null) {
      message.withGlobal = object.withGlobal;
    } else {
      message.withGlobal = false;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNftDataResponse: object = {};

export const QueryAllNftDataResponse = {
  encode(
    message: QueryAllNftDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nftData) {
      NftData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNftDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNftDataResponse,
    } as QueryAllNftDataResponse;
    message.nftData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftData.push(NftData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNftDataResponse {
    const message = {
      ...baseQueryAllNftDataResponse,
    } as QueryAllNftDataResponse;
    message.nftData = [];
    if (object.nftData !== undefined && object.nftData !== null) {
      for (const e of object.nftData) {
        message.nftData.push(NftData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNftDataResponse): unknown {
    const obj: any = {};
    if (message.nftData) {
      obj.nftData = message.nftData.map((e) =>
        e ? NftData.toJSON(e) : undefined
      );
    } else {
      obj.nftData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNftDataResponse>
  ): QueryAllNftDataResponse {
    const message = {
      ...baseQueryAllNftDataResponse,
    } as QueryAllNftDataResponse;
    message.nftData = [];
    if (object.nftData !== undefined && object.nftData !== null) {
      for (const e of object.nftData) {
        message.nftData.push(NftData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetActionByRefIdRequest: object = { refId: "" };

export const QueryGetActionByRefIdRequest = {
  encode(
    message: QueryGetActionByRefIdRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.refId !== "") {
      writer.uint32(10).string(message.refId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionByRefIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionByRefIdRequest,
    } as QueryGetActionByRefIdRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.refId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionByRefIdRequest {
    const message = {
      ...baseQueryGetActionByRefIdRequest,
    } as QueryGetActionByRefIdRequest;
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = String(object.refId);
    } else {
      message.refId = "";
    }
    return message;
  },

  toJSON(message: QueryGetActionByRefIdRequest): unknown {
    const obj: any = {};
    message.refId !== undefined && (obj.refId = message.refId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionByRefIdRequest>
  ): QueryGetActionByRefIdRequest {
    const message = {
      ...baseQueryGetActionByRefIdRequest,
    } as QueryGetActionByRefIdRequest;
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = object.refId;
    } else {
      message.refId = "";
    }
    return message;
  },
};

const baseQueryGetActionByRefIdResponse: object = {};

export const QueryGetActionByRefIdResponse = {
  encode(
    message: QueryGetActionByRefIdResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.actionByRefId !== undefined) {
      ActionByRefId.encode(
        message.actionByRefId,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionByRefIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionByRefIdResponse,
    } as QueryGetActionByRefIdResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionByRefId = ActionByRefId.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionByRefIdResponse {
    const message = {
      ...baseQueryGetActionByRefIdResponse,
    } as QueryGetActionByRefIdResponse;
    if (object.actionByRefId !== undefined && object.actionByRefId !== null) {
      message.actionByRefId = ActionByRefId.fromJSON(object.actionByRefId);
    } else {
      message.actionByRefId = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetActionByRefIdResponse): unknown {
    const obj: any = {};
    message.actionByRefId !== undefined &&
      (obj.actionByRefId = message.actionByRefId
        ? ActionByRefId.toJSON(message.actionByRefId)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionByRefIdResponse>
  ): QueryGetActionByRefIdResponse {
    const message = {
      ...baseQueryGetActionByRefIdResponse,
    } as QueryGetActionByRefIdResponse;
    if (object.actionByRefId !== undefined && object.actionByRefId !== null) {
      message.actionByRefId = ActionByRefId.fromPartial(object.actionByRefId);
    } else {
      message.actionByRefId = undefined;
    }
    return message;
  },
};

const baseQueryAllActionByRefIdRequest: object = {};

export const QueryAllActionByRefIdRequest = {
  encode(
    message: QueryAllActionByRefIdRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionByRefIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionByRefIdRequest,
    } as QueryAllActionByRefIdRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionByRefIdRequest {
    const message = {
      ...baseQueryAllActionByRefIdRequest,
    } as QueryAllActionByRefIdRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionByRefIdRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionByRefIdRequest>
  ): QueryAllActionByRefIdRequest {
    const message = {
      ...baseQueryAllActionByRefIdRequest,
    } as QueryAllActionByRefIdRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllActionByRefIdResponse: object = {};

export const QueryAllActionByRefIdResponse = {
  encode(
    message: QueryAllActionByRefIdResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.actionByRefId) {
      ActionByRefId.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionByRefIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionByRefIdResponse,
    } as QueryAllActionByRefIdResponse;
    message.actionByRefId = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionByRefId.push(
            ActionByRefId.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionByRefIdResponse {
    const message = {
      ...baseQueryAllActionByRefIdResponse,
    } as QueryAllActionByRefIdResponse;
    message.actionByRefId = [];
    if (object.actionByRefId !== undefined && object.actionByRefId !== null) {
      for (const e of object.actionByRefId) {
        message.actionByRefId.push(ActionByRefId.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionByRefIdResponse): unknown {
    const obj: any = {};
    if (message.actionByRefId) {
      obj.actionByRefId = message.actionByRefId.map((e) =>
        e ? ActionByRefId.toJSON(e) : undefined
      );
    } else {
      obj.actionByRefId = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionByRefIdResponse>
  ): QueryAllActionByRefIdResponse {
    const message = {
      ...baseQueryAllActionByRefIdResponse,
    } as QueryAllActionByRefIdResponse;
    message.actionByRefId = [];
    if (object.actionByRefId !== undefined && object.actionByRefId !== null) {
      for (const e of object.actionByRefId) {
        message.actionByRefId.push(ActionByRefId.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetOrganizationRequest: object = { name: "" };

export const QueryGetOrganizationRequest = {
  encode(
    message: QueryGetOrganizationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOrganizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOrganizationRequest,
    } as QueryGetOrganizationRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrganizationRequest {
    const message = {
      ...baseQueryGetOrganizationRequest,
    } as QueryGetOrganizationRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetOrganizationRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOrganizationRequest>
  ): QueryGetOrganizationRequest {
    const message = {
      ...baseQueryGetOrganizationRequest,
    } as QueryGetOrganizationRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetOrganizationResponse: object = {};

export const QueryGetOrganizationResponse = {
  encode(
    message: QueryGetOrganizationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.organization !== undefined) {
      Organization.encode(
        message.organization,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOrganizationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOrganizationResponse,
    } as QueryGetOrganizationResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.organization = Organization.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrganizationResponse {
    const message = {
      ...baseQueryGetOrganizationResponse,
    } as QueryGetOrganizationResponse;
    if (object.organization !== undefined && object.organization !== null) {
      message.organization = Organization.fromJSON(object.organization);
    } else {
      message.organization = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOrganizationResponse): unknown {
    const obj: any = {};
    message.organization !== undefined &&
      (obj.organization = message.organization
        ? Organization.toJSON(message.organization)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOrganizationResponse>
  ): QueryGetOrganizationResponse {
    const message = {
      ...baseQueryGetOrganizationResponse,
    } as QueryGetOrganizationResponse;
    if (object.organization !== undefined && object.organization !== null) {
      message.organization = Organization.fromPartial(object.organization);
    } else {
      message.organization = undefined;
    }
    return message;
  },
};

const baseQueryAllOrganizationRequest: object = {};

export const QueryAllOrganizationRequest = {
  encode(
    message: QueryAllOrganizationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOrganizationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOrganizationRequest,
    } as QueryAllOrganizationRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrganizationRequest {
    const message = {
      ...baseQueryAllOrganizationRequest,
    } as QueryAllOrganizationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOrganizationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOrganizationRequest>
  ): QueryAllOrganizationRequest {
    const message = {
      ...baseQueryAllOrganizationRequest,
    } as QueryAllOrganizationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOrganizationResponse: object = {};

export const QueryAllOrganizationResponse = {
  encode(
    message: QueryAllOrganizationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.organization) {
      Organization.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOrganizationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOrganizationResponse,
    } as QueryAllOrganizationResponse;
    message.organization = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.organization.push(
            Organization.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrganizationResponse {
    const message = {
      ...baseQueryAllOrganizationResponse,
    } as QueryAllOrganizationResponse;
    message.organization = [];
    if (object.organization !== undefined && object.organization !== null) {
      for (const e of object.organization) {
        message.organization.push(Organization.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOrganizationResponse): unknown {
    const obj: any = {};
    if (message.organization) {
      obj.organization = message.organization.map((e) =>
        e ? Organization.toJSON(e) : undefined
      );
    } else {
      obj.organization = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOrganizationResponse>
  ): QueryAllOrganizationResponse {
    const message = {
      ...baseQueryAllOrganizationResponse,
    } as QueryAllOrganizationResponse;
    message.organization = [];
    if (object.organization !== undefined && object.organization !== null) {
      for (const e of object.organization) {
        message.organization.push(Organization.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNftCollectionRequest: object = { nftSchemaCode: "" };

export const QueryGetNftCollectionRequest = {
  encode(
    message: QueryGetNftCollectionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNftCollectionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNftCollectionRequest,
    } as QueryGetNftCollectionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftCollectionRequest {
    const message = {
      ...baseQueryGetNftCollectionRequest,
    } as QueryGetNftCollectionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNftCollectionRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftCollectionRequest>
  ): QueryGetNftCollectionRequest {
    const message = {
      ...baseQueryGetNftCollectionRequest,
    } as QueryGetNftCollectionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNftCollectionResponse: object = {};

export const QueryGetNftCollectionResponse = {
  encode(
    message: QueryGetNftCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nftCollection) {
      NftData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNftCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNftCollectionResponse,
    } as QueryGetNftCollectionResponse;
    message.nftCollection = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftCollection.push(NftData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftCollectionResponse {
    const message = {
      ...baseQueryGetNftCollectionResponse,
    } as QueryGetNftCollectionResponse;
    message.nftCollection = [];
    if (object.nftCollection !== undefined && object.nftCollection !== null) {
      for (const e of object.nftCollection) {
        message.nftCollection.push(NftData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNftCollectionResponse): unknown {
    const obj: any = {};
    if (message.nftCollection) {
      obj.nftCollection = message.nftCollection.map((e) =>
        e ? NftData.toJSON(e) : undefined
      );
    } else {
      obj.nftCollection = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftCollectionResponse>
  ): QueryGetNftCollectionResponse {
    const message = {
      ...baseQueryGetNftCollectionResponse,
    } as QueryGetNftCollectionResponse;
    message.nftCollection = [];
    if (object.nftCollection !== undefined && object.nftCollection !== null) {
      for (const e of object.nftCollection) {
        message.nftCollection.push(NftData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNFTSchemaByContractRequest: object = {
  originContractAddress: "",
};

export const QueryGetNFTSchemaByContractRequest = {
  encode(
    message: QueryGetNFTSchemaByContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.originContractAddress !== "") {
      writer.uint32(10).string(message.originContractAddress);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaByContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaByContractRequest,
    } as QueryGetNFTSchemaByContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.originContractAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaByContractRequest {
    const message = {
      ...baseQueryGetNFTSchemaByContractRequest,
    } as QueryGetNFTSchemaByContractRequest;
    if (
      object.originContractAddress !== undefined &&
      object.originContractAddress !== null
    ) {
      message.originContractAddress = String(object.originContractAddress);
    } else {
      message.originContractAddress = "";
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaByContractRequest): unknown {
    const obj: any = {};
    message.originContractAddress !== undefined &&
      (obj.originContractAddress = message.originContractAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaByContractRequest>
  ): QueryGetNFTSchemaByContractRequest {
    const message = {
      ...baseQueryGetNFTSchemaByContractRequest,
    } as QueryGetNFTSchemaByContractRequest;
    if (
      object.originContractAddress !== undefined &&
      object.originContractAddress !== null
    ) {
      message.originContractAddress = object.originContractAddress;
    } else {
      message.originContractAddress = "";
    }
    return message;
  },
};

const baseQueryGetNFTSchemaByContractResponse: object = {};

export const QueryGetNFTSchemaByContractResponse = {
  encode(
    message: QueryGetNFTSchemaByContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nFTSchemaByContract !== undefined) {
      NFTSchemaByContract.encode(
        message.nFTSchemaByContract,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaByContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaByContractResponse,
    } as QueryGetNFTSchemaByContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchemaByContract = NFTSchemaByContract.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaByContractResponse {
    const message = {
      ...baseQueryGetNFTSchemaByContractResponse,
    } as QueryGetNFTSchemaByContractResponse;
    if (
      object.nFTSchemaByContract !== undefined &&
      object.nFTSchemaByContract !== null
    ) {
      message.nFTSchemaByContract = NFTSchemaByContract.fromJSON(
        object.nFTSchemaByContract
      );
    } else {
      message.nFTSchemaByContract = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaByContractResponse): unknown {
    const obj: any = {};
    message.nFTSchemaByContract !== undefined &&
      (obj.nFTSchemaByContract = message.nFTSchemaByContract
        ? NFTSchemaByContract.toJSON(message.nFTSchemaByContract)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaByContractResponse>
  ): QueryGetNFTSchemaByContractResponse {
    const message = {
      ...baseQueryGetNFTSchemaByContractResponse,
    } as QueryGetNFTSchemaByContractResponse;
    if (
      object.nFTSchemaByContract !== undefined &&
      object.nFTSchemaByContract !== null
    ) {
      message.nFTSchemaByContract = NFTSchemaByContract.fromPartial(
        object.nFTSchemaByContract
      );
    } else {
      message.nFTSchemaByContract = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaByContractRequest: object = {};

export const QueryAllNFTSchemaByContractRequest = {
  encode(
    message: QueryAllNFTSchemaByContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaByContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaByContractRequest,
    } as QueryAllNFTSchemaByContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaByContractRequest {
    const message = {
      ...baseQueryAllNFTSchemaByContractRequest,
    } as QueryAllNFTSchemaByContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaByContractRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaByContractRequest>
  ): QueryAllNFTSchemaByContractRequest {
    const message = {
      ...baseQueryAllNFTSchemaByContractRequest,
    } as QueryAllNFTSchemaByContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaByContractResponse: object = {};

export const QueryAllNFTSchemaByContractResponse = {
  encode(
    message: QueryAllNFTSchemaByContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nFTSchemaByContract) {
      NFTSchemaByContract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaByContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaByContractResponse,
    } as QueryAllNFTSchemaByContractResponse;
    message.nFTSchemaByContract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchemaByContract.push(
            NFTSchemaByContract.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaByContractResponse {
    const message = {
      ...baseQueryAllNFTSchemaByContractResponse,
    } as QueryAllNFTSchemaByContractResponse;
    message.nFTSchemaByContract = [];
    if (
      object.nFTSchemaByContract !== undefined &&
      object.nFTSchemaByContract !== null
    ) {
      for (const e of object.nFTSchemaByContract) {
        message.nFTSchemaByContract.push(NFTSchemaByContract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaByContractResponse): unknown {
    const obj: any = {};
    if (message.nFTSchemaByContract) {
      obj.nFTSchemaByContract = message.nFTSchemaByContract.map((e) =>
        e ? NFTSchemaByContract.toJSON(e) : undefined
      );
    } else {
      obj.nFTSchemaByContract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaByContractResponse>
  ): QueryAllNFTSchemaByContractResponse {
    const message = {
      ...baseQueryAllNFTSchemaByContractResponse,
    } as QueryAllNFTSchemaByContractResponse;
    message.nFTSchemaByContract = [];
    if (
      object.nFTSchemaByContract !== undefined &&
      object.nFTSchemaByContract !== null
    ) {
      for (const e of object.nFTSchemaByContract) {
        message.nFTSchemaByContract.push(NFTSchemaByContract.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNFTFeeConfigRequest: object = {};

export const QueryGetNFTFeeConfigRequest = {
  encode(
    _: QueryGetNFTFeeConfigRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTFeeConfigRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTFeeConfigRequest,
    } as QueryGetNFTFeeConfigRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetNFTFeeConfigRequest {
    const message = {
      ...baseQueryGetNFTFeeConfigRequest,
    } as QueryGetNFTFeeConfigRequest;
    return message;
  },

  toJSON(_: QueryGetNFTFeeConfigRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetNFTFeeConfigRequest>
  ): QueryGetNFTFeeConfigRequest {
    const message = {
      ...baseQueryGetNFTFeeConfigRequest,
    } as QueryGetNFTFeeConfigRequest;
    return message;
  },
};

const baseQueryGetNFTFeeConfigResponse: object = {};

export const QueryGetNFTFeeConfigResponse = {
  encode(
    message: QueryGetNFTFeeConfigResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NFTFeeConfig !== undefined) {
      NFTFeeConfig.encode(
        message.NFTFeeConfig,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTFeeConfigResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTFeeConfigResponse,
    } as QueryGetNFTFeeConfigResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NFTFeeConfig = NFTFeeConfig.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTFeeConfigResponse {
    const message = {
      ...baseQueryGetNFTFeeConfigResponse,
    } as QueryGetNFTFeeConfigResponse;
    if (object.NFTFeeConfig !== undefined && object.NFTFeeConfig !== null) {
      message.NFTFeeConfig = NFTFeeConfig.fromJSON(object.NFTFeeConfig);
    } else {
      message.NFTFeeConfig = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNFTFeeConfigResponse): unknown {
    const obj: any = {};
    message.NFTFeeConfig !== undefined &&
      (obj.NFTFeeConfig = message.NFTFeeConfig
        ? NFTFeeConfig.toJSON(message.NFTFeeConfig)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTFeeConfigResponse>
  ): QueryGetNFTFeeConfigResponse {
    const message = {
      ...baseQueryGetNFTFeeConfigResponse,
    } as QueryGetNFTFeeConfigResponse;
    if (object.NFTFeeConfig !== undefined && object.NFTFeeConfig !== null) {
      message.NFTFeeConfig = NFTFeeConfig.fromPartial(object.NFTFeeConfig);
    } else {
      message.NFTFeeConfig = undefined;
    }
    return message;
  },
};

const baseQueryGetNFTFeeBalanceRequest: object = {};

export const QueryGetNFTFeeBalanceRequest = {
  encode(
    _: QueryGetNFTFeeBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTFeeBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTFeeBalanceRequest,
    } as QueryGetNFTFeeBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetNFTFeeBalanceRequest {
    const message = {
      ...baseQueryGetNFTFeeBalanceRequest,
    } as QueryGetNFTFeeBalanceRequest;
    return message;
  },

  toJSON(_: QueryGetNFTFeeBalanceRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetNFTFeeBalanceRequest>
  ): QueryGetNFTFeeBalanceRequest {
    const message = {
      ...baseQueryGetNFTFeeBalanceRequest,
    } as QueryGetNFTFeeBalanceRequest;
    return message;
  },
};

const baseQueryGetNFTFeeBalanceResponse: object = {};

export const QueryGetNFTFeeBalanceResponse = {
  encode(
    message: QueryGetNFTFeeBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NFTFeeBalance !== undefined) {
      NFTFeeBalance.encode(
        message.NFTFeeBalance,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTFeeBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTFeeBalanceResponse,
    } as QueryGetNFTFeeBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NFTFeeBalance = NFTFeeBalance.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTFeeBalanceResponse {
    const message = {
      ...baseQueryGetNFTFeeBalanceResponse,
    } as QueryGetNFTFeeBalanceResponse;
    if (object.NFTFeeBalance !== undefined && object.NFTFeeBalance !== null) {
      message.NFTFeeBalance = NFTFeeBalance.fromJSON(object.NFTFeeBalance);
    } else {
      message.NFTFeeBalance = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNFTFeeBalanceResponse): unknown {
    const obj: any = {};
    message.NFTFeeBalance !== undefined &&
      (obj.NFTFeeBalance = message.NFTFeeBalance
        ? NFTFeeBalance.toJSON(message.NFTFeeBalance)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTFeeBalanceResponse>
  ): QueryGetNFTFeeBalanceResponse {
    const message = {
      ...baseQueryGetNFTFeeBalanceResponse,
    } as QueryGetNFTFeeBalanceResponse;
    if (object.NFTFeeBalance !== undefined && object.NFTFeeBalance !== null) {
      message.NFTFeeBalance = NFTFeeBalance.fromPartial(object.NFTFeeBalance);
    } else {
      message.NFTFeeBalance = undefined;
    }
    return message;
  },
};

const baseQueryGetMetadataCreatorRequest: object = { nftSchemaCode: "" };

export const QueryGetMetadataCreatorRequest = {
  encode(
    message: QueryGetMetadataCreatorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMetadataCreatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMetadataCreatorRequest,
    } as QueryGetMetadataCreatorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMetadataCreatorRequest {
    const message = {
      ...baseQueryGetMetadataCreatorRequest,
    } as QueryGetMetadataCreatorRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },

  toJSON(message: QueryGetMetadataCreatorRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMetadataCreatorRequest>
  ): QueryGetMetadataCreatorRequest {
    const message = {
      ...baseQueryGetMetadataCreatorRequest,
    } as QueryGetMetadataCreatorRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },
};

const baseQueryGetMetadataCreatorResponse: object = {};

export const QueryGetMetadataCreatorResponse = {
  encode(
    message: QueryGetMetadataCreatorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.metadataCreator !== undefined) {
      MetadataCreator.encode(
        message.metadataCreator,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMetadataCreatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMetadataCreatorResponse,
    } as QueryGetMetadataCreatorResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadataCreator = MetadataCreator.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMetadataCreatorResponse {
    const message = {
      ...baseQueryGetMetadataCreatorResponse,
    } as QueryGetMetadataCreatorResponse;
    if (
      object.metadataCreator !== undefined &&
      object.metadataCreator !== null
    ) {
      message.metadataCreator = MetadataCreator.fromJSON(
        object.metadataCreator
      );
    } else {
      message.metadataCreator = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMetadataCreatorResponse): unknown {
    const obj: any = {};
    message.metadataCreator !== undefined &&
      (obj.metadataCreator = message.metadataCreator
        ? MetadataCreator.toJSON(message.metadataCreator)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMetadataCreatorResponse>
  ): QueryGetMetadataCreatorResponse {
    const message = {
      ...baseQueryGetMetadataCreatorResponse,
    } as QueryGetMetadataCreatorResponse;
    if (
      object.metadataCreator !== undefined &&
      object.metadataCreator !== null
    ) {
      message.metadataCreator = MetadataCreator.fromPartial(
        object.metadataCreator
      );
    } else {
      message.metadataCreator = undefined;
    }
    return message;
  },
};

const baseQueryAllMetadataCreatorRequest: object = {};

export const QueryAllMetadataCreatorRequest = {
  encode(
    message: QueryAllMetadataCreatorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllMetadataCreatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMetadataCreatorRequest,
    } as QueryAllMetadataCreatorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMetadataCreatorRequest {
    const message = {
      ...baseQueryAllMetadataCreatorRequest,
    } as QueryAllMetadataCreatorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMetadataCreatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMetadataCreatorRequest>
  ): QueryAllMetadataCreatorRequest {
    const message = {
      ...baseQueryAllMetadataCreatorRequest,
    } as QueryAllMetadataCreatorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMetadataCreatorResponse: object = {};

export const QueryAllMetadataCreatorResponse = {
  encode(
    message: QueryAllMetadataCreatorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.metadataCreator) {
      MetadataCreator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllMetadataCreatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMetadataCreatorResponse,
    } as QueryAllMetadataCreatorResponse;
    message.metadataCreator = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadataCreator.push(
            MetadataCreator.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMetadataCreatorResponse {
    const message = {
      ...baseQueryAllMetadataCreatorResponse,
    } as QueryAllMetadataCreatorResponse;
    message.metadataCreator = [];
    if (
      object.metadataCreator !== undefined &&
      object.metadataCreator !== null
    ) {
      for (const e of object.metadataCreator) {
        message.metadataCreator.push(MetadataCreator.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMetadataCreatorResponse): unknown {
    const obj: any = {};
    if (message.metadataCreator) {
      obj.metadataCreator = message.metadataCreator.map((e) =>
        e ? MetadataCreator.toJSON(e) : undefined
      );
    } else {
      obj.metadataCreator = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMetadataCreatorResponse>
  ): QueryAllMetadataCreatorResponse {
    const message = {
      ...baseQueryAllMetadataCreatorResponse,
    } as QueryAllMetadataCreatorResponse;
    message.metadataCreator = [];
    if (
      object.metadataCreator !== undefined &&
      object.metadataCreator !== null
    ) {
      for (const e of object.metadataCreator) {
        message.metadataCreator.push(MetadataCreator.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetActionExecutorRequest: object = {
  nftSchemaCode: "",
  executorAddress: "",
};

export const QueryGetActionExecutorRequest = {
  encode(
    message: QueryGetActionExecutorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.executorAddress !== "") {
      writer.uint32(18).string(message.executorAddress);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionExecutorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionExecutorRequest,
    } as QueryGetActionExecutorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.executorAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionExecutorRequest {
    const message = {
      ...baseQueryGetActionExecutorRequest,
    } as QueryGetActionExecutorRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.executorAddress !== undefined &&
      object.executorAddress !== null
    ) {
      message.executorAddress = String(object.executorAddress);
    } else {
      message.executorAddress = "";
    }
    return message;
  },

  toJSON(message: QueryGetActionExecutorRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.executorAddress !== undefined &&
      (obj.executorAddress = message.executorAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionExecutorRequest>
  ): QueryGetActionExecutorRequest {
    const message = {
      ...baseQueryGetActionExecutorRequest,
    } as QueryGetActionExecutorRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.executorAddress !== undefined &&
      object.executorAddress !== null
    ) {
      message.executorAddress = object.executorAddress;
    } else {
      message.executorAddress = "";
    }
    return message;
  },
};

const baseQueryGetActionExecutorResponse: object = {};

export const QueryGetActionExecutorResponse = {
  encode(
    message: QueryGetActionExecutorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.actionExecutor !== undefined) {
      ActionExecutor.encode(
        message.actionExecutor,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionExecutorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionExecutorResponse,
    } as QueryGetActionExecutorResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionExecutor = ActionExecutor.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionExecutorResponse {
    const message = {
      ...baseQueryGetActionExecutorResponse,
    } as QueryGetActionExecutorResponse;
    if (object.actionExecutor !== undefined && object.actionExecutor !== null) {
      message.actionExecutor = ActionExecutor.fromJSON(object.actionExecutor);
    } else {
      message.actionExecutor = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetActionExecutorResponse): unknown {
    const obj: any = {};
    message.actionExecutor !== undefined &&
      (obj.actionExecutor = message.actionExecutor
        ? ActionExecutor.toJSON(message.actionExecutor)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionExecutorResponse>
  ): QueryGetActionExecutorResponse {
    const message = {
      ...baseQueryGetActionExecutorResponse,
    } as QueryGetActionExecutorResponse;
    if (object.actionExecutor !== undefined && object.actionExecutor !== null) {
      message.actionExecutor = ActionExecutor.fromPartial(
        object.actionExecutor
      );
    } else {
      message.actionExecutor = undefined;
    }
    return message;
  },
};

const baseQueryAllActionExecutorRequest: object = {};

export const QueryAllActionExecutorRequest = {
  encode(
    message: QueryAllActionExecutorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionExecutorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionExecutorRequest,
    } as QueryAllActionExecutorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionExecutorRequest {
    const message = {
      ...baseQueryAllActionExecutorRequest,
    } as QueryAllActionExecutorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionExecutorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionExecutorRequest>
  ): QueryAllActionExecutorRequest {
    const message = {
      ...baseQueryAllActionExecutorRequest,
    } as QueryAllActionExecutorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllActionExecutorResponse: object = {};

export const QueryAllActionExecutorResponse = {
  encode(
    message: QueryAllActionExecutorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.actionExecutor) {
      ActionExecutor.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionExecutorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionExecutorResponse,
    } as QueryAllActionExecutorResponse;
    message.actionExecutor = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionExecutor.push(
            ActionExecutor.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionExecutorResponse {
    const message = {
      ...baseQueryAllActionExecutorResponse,
    } as QueryAllActionExecutorResponse;
    message.actionExecutor = [];
    if (object.actionExecutor !== undefined && object.actionExecutor !== null) {
      for (const e of object.actionExecutor) {
        message.actionExecutor.push(ActionExecutor.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionExecutorResponse): unknown {
    const obj: any = {};
    if (message.actionExecutor) {
      obj.actionExecutor = message.actionExecutor.map((e) =>
        e ? ActionExecutor.toJSON(e) : undefined
      );
    } else {
      obj.actionExecutor = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionExecutorResponse>
  ): QueryAllActionExecutorResponse {
    const message = {
      ...baseQueryAllActionExecutorResponse,
    } as QueryAllActionExecutorResponse;
    message.actionExecutor = [];
    if (object.actionExecutor !== undefined && object.actionExecutor !== null) {
      for (const e of object.actionExecutor) {
        message.actionExecutor.push(ActionExecutor.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetSchemaAttributeRequest: object = {
  nftSchemaCode: "",
  name: "",
};

export const QueryGetSchemaAttributeRequest = {
  encode(
    message: QueryGetSchemaAttributeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSchemaAttributeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSchemaAttributeRequest,
    } as QueryGetSchemaAttributeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSchemaAttributeRequest {
    const message = {
      ...baseQueryGetSchemaAttributeRequest,
    } as QueryGetSchemaAttributeRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetSchemaAttributeRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSchemaAttributeRequest>
  ): QueryGetSchemaAttributeRequest {
    const message = {
      ...baseQueryGetSchemaAttributeRequest,
    } as QueryGetSchemaAttributeRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetSchemaAttributeResponse: object = {};

export const QueryGetSchemaAttributeResponse = {
  encode(
    message: QueryGetSchemaAttributeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.schemaAttribute !== undefined) {
      SchemaAttribute.encode(
        message.schemaAttribute,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSchemaAttributeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSchemaAttributeResponse,
    } as QueryGetSchemaAttributeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaAttribute = SchemaAttribute.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSchemaAttributeResponse {
    const message = {
      ...baseQueryGetSchemaAttributeResponse,
    } as QueryGetSchemaAttributeResponse;
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      message.schemaAttribute = SchemaAttribute.fromJSON(
        object.schemaAttribute
      );
    } else {
      message.schemaAttribute = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSchemaAttributeResponse): unknown {
    const obj: any = {};
    message.schemaAttribute !== undefined &&
      (obj.schemaAttribute = message.schemaAttribute
        ? SchemaAttribute.toJSON(message.schemaAttribute)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSchemaAttributeResponse>
  ): QueryGetSchemaAttributeResponse {
    const message = {
      ...baseQueryGetSchemaAttributeResponse,
    } as QueryGetSchemaAttributeResponse;
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      message.schemaAttribute = SchemaAttribute.fromPartial(
        object.schemaAttribute
      );
    } else {
      message.schemaAttribute = undefined;
    }
    return message;
  },
};

const baseQueryAllSchemaAttributeRequest: object = { nftSchemaCode: "" };

export const QueryAllSchemaAttributeRequest = {
  encode(
    message: QueryAllSchemaAttributeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSchemaAttributeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSchemaAttributeRequest,
    } as QueryAllSchemaAttributeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSchemaAttributeRequest {
    const message = {
      ...baseQueryAllSchemaAttributeRequest,
    } as QueryAllSchemaAttributeRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSchemaAttributeRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSchemaAttributeRequest>
  ): QueryAllSchemaAttributeRequest {
    const message = {
      ...baseQueryAllSchemaAttributeRequest,
    } as QueryAllSchemaAttributeRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSchemaAttributeResponse: object = {};

export const QueryAllSchemaAttributeResponse = {
  encode(
    message: QueryAllSchemaAttributeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.schemaAttribute) {
      SchemaAttribute.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSchemaAttributeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSchemaAttributeResponse,
    } as QueryAllSchemaAttributeResponse;
    message.schemaAttribute = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaAttribute.push(
            SchemaAttribute.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSchemaAttributeResponse {
    const message = {
      ...baseQueryAllSchemaAttributeResponse,
    } as QueryAllSchemaAttributeResponse;
    message.schemaAttribute = [];
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      for (const e of object.schemaAttribute) {
        message.schemaAttribute.push(SchemaAttribute.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSchemaAttributeResponse): unknown {
    const obj: any = {};
    if (message.schemaAttribute) {
      obj.schemaAttribute = message.schemaAttribute.map((e) =>
        e ? SchemaAttribute.toJSON(e) : undefined
      );
    } else {
      obj.schemaAttribute = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSchemaAttributeResponse>
  ): QueryAllSchemaAttributeResponse {
    const message = {
      ...baseQueryAllSchemaAttributeResponse,
    } as QueryAllSchemaAttributeResponse;
    message.schemaAttribute = [];
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      for (const e of object.schemaAttribute) {
        message.schemaAttribute.push(SchemaAttribute.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryListAttributeBySchemaRequest: object = { nftSchemaCode: "" };

export const QueryListAttributeBySchemaRequest = {
  encode(
    message: QueryListAttributeBySchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListAttributeBySchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListAttributeBySchemaRequest,
    } as QueryListAttributeBySchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryListAttributeBySchemaRequest {
    const message = {
      ...baseQueryListAttributeBySchemaRequest,
    } as QueryListAttributeBySchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },

  toJSON(message: QueryListAttributeBySchemaRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryListAttributeBySchemaRequest>
  ): QueryListAttributeBySchemaRequest {
    const message = {
      ...baseQueryListAttributeBySchemaRequest,
    } as QueryListAttributeBySchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },
};

const baseQueryListAttributeBySchemaResponse: object = {};

export const QueryListAttributeBySchemaResponse = {
  encode(
    message: QueryListAttributeBySchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.schemaAttribute) {
      SchemaAttribute.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListAttributeBySchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListAttributeBySchemaResponse,
    } as QueryListAttributeBySchemaResponse;
    message.schemaAttribute = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaAttribute.push(
            SchemaAttribute.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryListAttributeBySchemaResponse {
    const message = {
      ...baseQueryListAttributeBySchemaResponse,
    } as QueryListAttributeBySchemaResponse;
    message.schemaAttribute = [];
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      for (const e of object.schemaAttribute) {
        message.schemaAttribute.push(SchemaAttribute.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryListAttributeBySchemaResponse): unknown {
    const obj: any = {};
    if (message.schemaAttribute) {
      obj.schemaAttribute = message.schemaAttribute.map((e) =>
        e ? SchemaAttribute.toJSON(e) : undefined
      );
    } else {
      obj.schemaAttribute = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryListAttributeBySchemaResponse>
  ): QueryListAttributeBySchemaResponse {
    const message = {
      ...baseQueryListAttributeBySchemaResponse,
    } as QueryListAttributeBySchemaResponse;
    message.schemaAttribute = [];
    if (
      object.schemaAttribute !== undefined &&
      object.schemaAttribute !== null
    ) {
      for (const e of object.schemaAttribute) {
        message.schemaAttribute.push(SchemaAttribute.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryGetActionOfSchemaRequest: object = {
  nftSchemaCode: "",
  name: "",
};

export const QueryGetActionOfSchemaRequest = {
  encode(
    message: QueryGetActionOfSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionOfSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionOfSchemaRequest,
    } as QueryGetActionOfSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionOfSchemaRequest {
    const message = {
      ...baseQueryGetActionOfSchemaRequest,
    } as QueryGetActionOfSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetActionOfSchemaRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionOfSchemaRequest>
  ): QueryGetActionOfSchemaRequest {
    const message = {
      ...baseQueryGetActionOfSchemaRequest,
    } as QueryGetActionOfSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetActionOfSchemaResponse: object = {};

export const QueryGetActionOfSchemaResponse = {
  encode(
    message: QueryGetActionOfSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.actionOfSchema !== undefined) {
      ActionOfSchema.encode(
        message.actionOfSchema,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActionOfSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActionOfSchemaResponse,
    } as QueryGetActionOfSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionOfSchema = ActionOfSchema.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActionOfSchemaResponse {
    const message = {
      ...baseQueryGetActionOfSchemaResponse,
    } as QueryGetActionOfSchemaResponse;
    if (object.actionOfSchema !== undefined && object.actionOfSchema !== null) {
      message.actionOfSchema = ActionOfSchema.fromJSON(object.actionOfSchema);
    } else {
      message.actionOfSchema = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetActionOfSchemaResponse): unknown {
    const obj: any = {};
    message.actionOfSchema !== undefined &&
      (obj.actionOfSchema = message.actionOfSchema
        ? ActionOfSchema.toJSON(message.actionOfSchema)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActionOfSchemaResponse>
  ): QueryGetActionOfSchemaResponse {
    const message = {
      ...baseQueryGetActionOfSchemaResponse,
    } as QueryGetActionOfSchemaResponse;
    if (object.actionOfSchema !== undefined && object.actionOfSchema !== null) {
      message.actionOfSchema = ActionOfSchema.fromPartial(
        object.actionOfSchema
      );
    } else {
      message.actionOfSchema = undefined;
    }
    return message;
  },
};

const baseQueryAllActionOfSchemaRequest: object = {};

export const QueryAllActionOfSchemaRequest = {
  encode(
    message: QueryAllActionOfSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionOfSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionOfSchemaRequest,
    } as QueryAllActionOfSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionOfSchemaRequest {
    const message = {
      ...baseQueryAllActionOfSchemaRequest,
    } as QueryAllActionOfSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionOfSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionOfSchemaRequest>
  ): QueryAllActionOfSchemaRequest {
    const message = {
      ...baseQueryAllActionOfSchemaRequest,
    } as QueryAllActionOfSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllActionOfSchemaResponse: object = {};

export const QueryAllActionOfSchemaResponse = {
  encode(
    message: QueryAllActionOfSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.actionOfSchema) {
      ActionOfSchema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActionOfSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActionOfSchemaResponse,
    } as QueryAllActionOfSchemaResponse;
    message.actionOfSchema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionOfSchema.push(
            ActionOfSchema.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActionOfSchemaResponse {
    const message = {
      ...baseQueryAllActionOfSchemaResponse,
    } as QueryAllActionOfSchemaResponse;
    message.actionOfSchema = [];
    if (object.actionOfSchema !== undefined && object.actionOfSchema !== null) {
      for (const e of object.actionOfSchema) {
        message.actionOfSchema.push(ActionOfSchema.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActionOfSchemaResponse): unknown {
    const obj: any = {};
    if (message.actionOfSchema) {
      obj.actionOfSchema = message.actionOfSchema.map((e) =>
        e ? ActionOfSchema.toJSON(e) : undefined
      );
    } else {
      obj.actionOfSchema = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActionOfSchemaResponse>
  ): QueryAllActionOfSchemaResponse {
    const message = {
      ...baseQueryAllActionOfSchemaResponse,
    } as QueryAllActionOfSchemaResponse;
    message.actionOfSchema = [];
    if (object.actionOfSchema !== undefined && object.actionOfSchema !== null) {
      for (const e of object.actionOfSchema) {
        message.actionOfSchema.push(ActionOfSchema.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetExecutorOfSchemaRequest: object = { nftSchemaCode: "" };

export const QueryGetExecutorOfSchemaRequest = {
  encode(
    message: QueryGetExecutorOfSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetExecutorOfSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetExecutorOfSchemaRequest,
    } as QueryGetExecutorOfSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExecutorOfSchemaRequest {
    const message = {
      ...baseQueryGetExecutorOfSchemaRequest,
    } as QueryGetExecutorOfSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },

  toJSON(message: QueryGetExecutorOfSchemaRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetExecutorOfSchemaRequest>
  ): QueryGetExecutorOfSchemaRequest {
    const message = {
      ...baseQueryGetExecutorOfSchemaRequest,
    } as QueryGetExecutorOfSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },
};

const baseQueryGetExecutorOfSchemaResponse: object = {};

export const QueryGetExecutorOfSchemaResponse = {
  encode(
    message: QueryGetExecutorOfSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.executorOfSchema !== undefined) {
      ExecutorOfSchema.encode(
        message.executorOfSchema,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetExecutorOfSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetExecutorOfSchemaResponse,
    } as QueryGetExecutorOfSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.executorOfSchema = ExecutorOfSchema.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExecutorOfSchemaResponse {
    const message = {
      ...baseQueryGetExecutorOfSchemaResponse,
    } as QueryGetExecutorOfSchemaResponse;
    if (
      object.executorOfSchema !== undefined &&
      object.executorOfSchema !== null
    ) {
      message.executorOfSchema = ExecutorOfSchema.fromJSON(
        object.executorOfSchema
      );
    } else {
      message.executorOfSchema = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetExecutorOfSchemaResponse): unknown {
    const obj: any = {};
    message.executorOfSchema !== undefined &&
      (obj.executorOfSchema = message.executorOfSchema
        ? ExecutorOfSchema.toJSON(message.executorOfSchema)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetExecutorOfSchemaResponse>
  ): QueryGetExecutorOfSchemaResponse {
    const message = {
      ...baseQueryGetExecutorOfSchemaResponse,
    } as QueryGetExecutorOfSchemaResponse;
    if (
      object.executorOfSchema !== undefined &&
      object.executorOfSchema !== null
    ) {
      message.executorOfSchema = ExecutorOfSchema.fromPartial(
        object.executorOfSchema
      );
    } else {
      message.executorOfSchema = undefined;
    }
    return message;
  },
};

const baseQueryAllExecutorOfSchemaRequest: object = {};

export const QueryAllExecutorOfSchemaRequest = {
  encode(
    message: QueryAllExecutorOfSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllExecutorOfSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllExecutorOfSchemaRequest,
    } as QueryAllExecutorOfSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExecutorOfSchemaRequest {
    const message = {
      ...baseQueryAllExecutorOfSchemaRequest,
    } as QueryAllExecutorOfSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllExecutorOfSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllExecutorOfSchemaRequest>
  ): QueryAllExecutorOfSchemaRequest {
    const message = {
      ...baseQueryAllExecutorOfSchemaRequest,
    } as QueryAllExecutorOfSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllExecutorOfSchemaResponse: object = {};

export const QueryAllExecutorOfSchemaResponse = {
  encode(
    message: QueryAllExecutorOfSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.executorOfSchema) {
      ExecutorOfSchema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllExecutorOfSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllExecutorOfSchemaResponse,
    } as QueryAllExecutorOfSchemaResponse;
    message.executorOfSchema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.executorOfSchema.push(
            ExecutorOfSchema.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExecutorOfSchemaResponse {
    const message = {
      ...baseQueryAllExecutorOfSchemaResponse,
    } as QueryAllExecutorOfSchemaResponse;
    message.executorOfSchema = [];
    if (
      object.executorOfSchema !== undefined &&
      object.executorOfSchema !== null
    ) {
      for (const e of object.executorOfSchema) {
        message.executorOfSchema.push(ExecutorOfSchema.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllExecutorOfSchemaResponse): unknown {
    const obj: any = {};
    if (message.executorOfSchema) {
      obj.executorOfSchema = message.executorOfSchema.map((e) =>
        e ? ExecutorOfSchema.toJSON(e) : undefined
      );
    } else {
      obj.executorOfSchema = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllExecutorOfSchemaResponse>
  ): QueryAllExecutorOfSchemaResponse {
    const message = {
      ...baseQueryAllExecutorOfSchemaResponse,
    } as QueryAllExecutorOfSchemaResponse;
    message.executorOfSchema = [];
    if (
      object.executorOfSchema !== undefined &&
      object.executorOfSchema !== null
    ) {
      for (const e of object.executorOfSchema) {
        message.executorOfSchema.push(ExecutorOfSchema.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetVirtualActionRequest: object = {
  nftSchemaCode: "",
  name: "",
};

export const QueryGetVirtualActionRequest = {
  encode(
    message: QueryGetVirtualActionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualActionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualActionRequest,
    } as QueryGetVirtualActionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualActionRequest {
    const message = {
      ...baseQueryGetVirtualActionRequest,
    } as QueryGetVirtualActionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetVirtualActionRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualActionRequest>
  ): QueryGetVirtualActionRequest {
    const message = {
      ...baseQueryGetVirtualActionRequest,
    } as QueryGetVirtualActionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetVirtualActionResponse: object = {};

export const QueryGetVirtualActionResponse = {
  encode(
    message: QueryGetVirtualActionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.virtualAction !== undefined) {
      VirtualAction.encode(
        message.virtualAction,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualActionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualActionResponse,
    } as QueryGetVirtualActionResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualAction = VirtualAction.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualActionResponse {
    const message = {
      ...baseQueryGetVirtualActionResponse,
    } as QueryGetVirtualActionResponse;
    if (object.virtualAction !== undefined && object.virtualAction !== null) {
      message.virtualAction = VirtualAction.fromJSON(object.virtualAction);
    } else {
      message.virtualAction = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualActionResponse): unknown {
    const obj: any = {};
    message.virtualAction !== undefined &&
      (obj.virtualAction = message.virtualAction
        ? VirtualAction.toJSON(message.virtualAction)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualActionResponse>
  ): QueryGetVirtualActionResponse {
    const message = {
      ...baseQueryGetVirtualActionResponse,
    } as QueryGetVirtualActionResponse;
    if (object.virtualAction !== undefined && object.virtualAction !== null) {
      message.virtualAction = VirtualAction.fromPartial(object.virtualAction);
    } else {
      message.virtualAction = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualActionRequest: object = { nftSchemaCode: "" };

export const QueryAllVirtualActionRequest = {
  encode(
    message: QueryAllVirtualActionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualActionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualActionRequest,
    } as QueryAllVirtualActionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualActionRequest {
    const message = {
      ...baseQueryAllVirtualActionRequest,
    } as QueryAllVirtualActionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualActionRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualActionRequest>
  ): QueryAllVirtualActionRequest {
    const message = {
      ...baseQueryAllVirtualActionRequest,
    } as QueryAllVirtualActionRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualActionResponse: object = {};

export const QueryAllVirtualActionResponse = {
  encode(
    message: QueryAllVirtualActionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.virtualAction) {
      VirtualAction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualActionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualActionResponse,
    } as QueryAllVirtualActionResponse;
    message.virtualAction = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualAction.push(
            VirtualAction.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualActionResponse {
    const message = {
      ...baseQueryAllVirtualActionResponse,
    } as QueryAllVirtualActionResponse;
    message.virtualAction = [];
    if (object.virtualAction !== undefined && object.virtualAction !== null) {
      for (const e of object.virtualAction) {
        message.virtualAction.push(VirtualAction.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualActionResponse): unknown {
    const obj: any = {};
    if (message.virtualAction) {
      obj.virtualAction = message.virtualAction.map((e) =>
        e ? VirtualAction.toJSON(e) : undefined
      );
    } else {
      obj.virtualAction = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualActionResponse>
  ): QueryAllVirtualActionResponse {
    const message = {
      ...baseQueryAllVirtualActionResponse,
    } as QueryAllVirtualActionResponse;
    message.virtualAction = [];
    if (object.virtualAction !== undefined && object.virtualAction !== null) {
      for (const e of object.virtualAction) {
        message.virtualAction.push(VirtualAction.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetVirtualSchemaRequest: object = { nftSchemaCode: "" };

export const QueryGetVirtualSchemaRequest = {
  encode(
    message: QueryGetVirtualSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualSchemaRequest,
    } as QueryGetVirtualSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualSchemaRequest {
    const message = {
      ...baseQueryGetVirtualSchemaRequest,
    } as QueryGetVirtualSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },

  toJSON(message: QueryGetVirtualSchemaRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualSchemaRequest>
  ): QueryGetVirtualSchemaRequest {
    const message = {
      ...baseQueryGetVirtualSchemaRequest,
    } as QueryGetVirtualSchemaRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },
};

const baseQueryGetVirtualSchemaResponse: object = {};

export const QueryGetVirtualSchemaResponse = {
  encode(
    message: QueryGetVirtualSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.virtualSchema !== undefined) {
      VirtualSchema.encode(
        message.virtualSchema,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualSchemaResponse,
    } as QueryGetVirtualSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchema = VirtualSchema.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualSchemaResponse {
    const message = {
      ...baseQueryGetVirtualSchemaResponse,
    } as QueryGetVirtualSchemaResponse;
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      message.virtualSchema = VirtualSchema.fromJSON(object.virtualSchema);
    } else {
      message.virtualSchema = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualSchemaResponse): unknown {
    const obj: any = {};
    message.virtualSchema !== undefined &&
      (obj.virtualSchema = message.virtualSchema
        ? VirtualSchema.toJSON(message.virtualSchema)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualSchemaResponse>
  ): QueryGetVirtualSchemaResponse {
    const message = {
      ...baseQueryGetVirtualSchemaResponse,
    } as QueryGetVirtualSchemaResponse;
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      message.virtualSchema = VirtualSchema.fromPartial(object.virtualSchema);
    } else {
      message.virtualSchema = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualSchemaRequest: object = {};

export const QueryAllVirtualSchemaRequest = {
  encode(
    message: QueryAllVirtualSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualSchemaRequest,
    } as QueryAllVirtualSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualSchemaRequest {
    const message = {
      ...baseQueryAllVirtualSchemaRequest,
    } as QueryAllVirtualSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualSchemaRequest>
  ): QueryAllVirtualSchemaRequest {
    const message = {
      ...baseQueryAllVirtualSchemaRequest,
    } as QueryAllVirtualSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualSchemaResponse: object = {};

export const QueryAllVirtualSchemaResponse = {
  encode(
    message: QueryAllVirtualSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.virtualSchema) {
      VirtualSchema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualSchemaResponse,
    } as QueryAllVirtualSchemaResponse;
    message.virtualSchema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchema.push(
            VirtualSchema.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualSchemaResponse {
    const message = {
      ...baseQueryAllVirtualSchemaResponse,
    } as QueryAllVirtualSchemaResponse;
    message.virtualSchema = [];
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      for (const e of object.virtualSchema) {
        message.virtualSchema.push(VirtualSchema.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualSchemaResponse): unknown {
    const obj: any = {};
    if (message.virtualSchema) {
      obj.virtualSchema = message.virtualSchema.map((e) =>
        e ? VirtualSchema.toJSON(e) : undefined
      );
    } else {
      obj.virtualSchema = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualSchemaResponse>
  ): QueryAllVirtualSchemaResponse {
    const message = {
      ...baseQueryAllVirtualSchemaResponse,
    } as QueryAllVirtualSchemaResponse;
    message.virtualSchema = [];
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      for (const e of object.virtualSchema) {
        message.virtualSchema.push(VirtualSchema.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetVirtualSchemaProposalRequest: object = { index: "" };

export const QueryGetVirtualSchemaProposalRequest = {
  encode(
    message: QueryGetVirtualSchemaProposalRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualSchemaProposalRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualSchemaProposalRequest,
    } as QueryGetVirtualSchemaProposalRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualSchemaProposalRequest {
    const message = {
      ...baseQueryGetVirtualSchemaProposalRequest,
    } as QueryGetVirtualSchemaProposalRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetVirtualSchemaProposalRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualSchemaProposalRequest>
  ): QueryGetVirtualSchemaProposalRequest {
    const message = {
      ...baseQueryGetVirtualSchemaProposalRequest,
    } as QueryGetVirtualSchemaProposalRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetVirtualSchemaProposalResponse: object = {};

export const QueryGetVirtualSchemaProposalResponse = {
  encode(
    message: QueryGetVirtualSchemaProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.virtualSchemaProposal !== undefined) {
      VirtualSchemaProposal.encode(
        message.virtualSchemaProposal,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualSchemaProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualSchemaProposalResponse,
    } as QueryGetVirtualSchemaProposalResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchemaProposal = VirtualSchemaProposal.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualSchemaProposalResponse {
    const message = {
      ...baseQueryGetVirtualSchemaProposalResponse,
    } as QueryGetVirtualSchemaProposalResponse;
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      message.virtualSchemaProposal = VirtualSchemaProposal.fromJSON(
        object.virtualSchemaProposal
      );
    } else {
      message.virtualSchemaProposal = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualSchemaProposalResponse): unknown {
    const obj: any = {};
    message.virtualSchemaProposal !== undefined &&
      (obj.virtualSchemaProposal = message.virtualSchemaProposal
        ? VirtualSchemaProposal.toJSON(message.virtualSchemaProposal)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualSchemaProposalResponse>
  ): QueryGetVirtualSchemaProposalResponse {
    const message = {
      ...baseQueryGetVirtualSchemaProposalResponse,
    } as QueryGetVirtualSchemaProposalResponse;
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      message.virtualSchemaProposal = VirtualSchemaProposal.fromPartial(
        object.virtualSchemaProposal
      );
    } else {
      message.virtualSchemaProposal = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualSchemaProposalRequest: object = {};

export const QueryAllVirtualSchemaProposalRequest = {
  encode(
    message: QueryAllVirtualSchemaProposalRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualSchemaProposalRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualSchemaProposalRequest,
    } as QueryAllVirtualSchemaProposalRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualSchemaProposalRequest {
    const message = {
      ...baseQueryAllVirtualSchemaProposalRequest,
    } as QueryAllVirtualSchemaProposalRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualSchemaProposalRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualSchemaProposalRequest>
  ): QueryAllVirtualSchemaProposalRequest {
    const message = {
      ...baseQueryAllVirtualSchemaProposalRequest,
    } as QueryAllVirtualSchemaProposalRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualSchemaProposalResponse: object = {};

export const QueryAllVirtualSchemaProposalResponse = {
  encode(
    message: QueryAllVirtualSchemaProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.virtualSchemaProposal) {
      VirtualSchemaProposal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVirtualSchemaProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualSchemaProposalResponse,
    } as QueryAllVirtualSchemaProposalResponse;
    message.virtualSchemaProposal = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchemaProposal.push(
            VirtualSchemaProposal.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVirtualSchemaProposalResponse {
    const message = {
      ...baseQueryAllVirtualSchemaProposalResponse,
    } as QueryAllVirtualSchemaProposalResponse;
    message.virtualSchemaProposal = [];
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      for (const e of object.virtualSchemaProposal) {
        message.virtualSchemaProposal.push(VirtualSchemaProposal.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualSchemaProposalResponse): unknown {
    const obj: any = {};
    if (message.virtualSchemaProposal) {
      obj.virtualSchemaProposal = message.virtualSchemaProposal.map((e) =>
        e ? VirtualSchemaProposal.toJSON(e) : undefined
      );
    } else {
      obj.virtualSchemaProposal = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualSchemaProposalResponse>
  ): QueryAllVirtualSchemaProposalResponse {
    const message = {
      ...baseQueryAllVirtualSchemaProposalResponse,
    } as QueryAllVirtualSchemaProposalResponse;
    message.virtualSchemaProposal = [];
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      for (const e of object.virtualSchemaProposal) {
        message.virtualSchemaProposal.push(
          VirtualSchemaProposal.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryListActiveProposalRequest: object = {};

export const QueryListActiveProposalRequest = {
  encode(
    _: QueryListActiveProposalRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListActiveProposalRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListActiveProposalRequest,
    } as QueryListActiveProposalRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryListActiveProposalRequest {
    const message = {
      ...baseQueryListActiveProposalRequest,
    } as QueryListActiveProposalRequest;
    return message;
  },

  toJSON(_: QueryListActiveProposalRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryListActiveProposalRequest>
  ): QueryListActiveProposalRequest {
    const message = {
      ...baseQueryListActiveProposalRequest,
    } as QueryListActiveProposalRequest;
    return message;
  },
};

const baseQueryListActiveProposalResponse: object = {};

export const QueryListActiveProposalResponse = {
  encode(
    message: QueryListActiveProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.virtualSchemaProposal) {
      VirtualSchemaProposal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListActiveProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListActiveProposalResponse,
    } as QueryListActiveProposalResponse;
    message.virtualSchemaProposal = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchemaProposal.push(
            VirtualSchemaProposal.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryListActiveProposalResponse {
    const message = {
      ...baseQueryListActiveProposalResponse,
    } as QueryListActiveProposalResponse;
    message.virtualSchemaProposal = [];
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      for (const e of object.virtualSchemaProposal) {
        message.virtualSchemaProposal.push(VirtualSchemaProposal.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryListActiveProposalResponse): unknown {
    const obj: any = {};
    if (message.virtualSchemaProposal) {
      obj.virtualSchemaProposal = message.virtualSchemaProposal.map((e) =>
        e ? VirtualSchemaProposal.toJSON(e) : undefined
      );
    } else {
      obj.virtualSchemaProposal = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryListActiveProposalResponse>
  ): QueryListActiveProposalResponse {
    const message = {
      ...baseQueryListActiveProposalResponse,
    } as QueryListActiveProposalResponse;
    message.virtualSchemaProposal = [];
    if (
      object.virtualSchemaProposal !== undefined &&
      object.virtualSchemaProposal !== null
    ) {
      for (const e of object.virtualSchemaProposal) {
        message.virtualSchemaProposal.push(
          VirtualSchemaProposal.fromPartial(e)
        );
      }
    }
    return message;
  },
};

const baseQueryGetLockSchemaFeeRequest: object = { index: "" };

export const QueryGetLockSchemaFeeRequest = {
  encode(
    message: QueryGetLockSchemaFeeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLockSchemaFeeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLockSchemaFeeRequest,
    } as QueryGetLockSchemaFeeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLockSchemaFeeRequest {
    const message = {
      ...baseQueryGetLockSchemaFeeRequest,
    } as QueryGetLockSchemaFeeRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetLockSchemaFeeRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLockSchemaFeeRequest>
  ): QueryGetLockSchemaFeeRequest {
    const message = {
      ...baseQueryGetLockSchemaFeeRequest,
    } as QueryGetLockSchemaFeeRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetLockSchemaFeeResponse: object = {};

export const QueryGetLockSchemaFeeResponse = {
  encode(
    message: QueryGetLockSchemaFeeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.lockSchemaFee !== undefined) {
      LockSchemaFee.encode(
        message.lockSchemaFee,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLockSchemaFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLockSchemaFeeResponse,
    } as QueryGetLockSchemaFeeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lockSchemaFee = LockSchemaFee.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLockSchemaFeeResponse {
    const message = {
      ...baseQueryGetLockSchemaFeeResponse,
    } as QueryGetLockSchemaFeeResponse;
    if (object.lockSchemaFee !== undefined && object.lockSchemaFee !== null) {
      message.lockSchemaFee = LockSchemaFee.fromJSON(object.lockSchemaFee);
    } else {
      message.lockSchemaFee = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLockSchemaFeeResponse): unknown {
    const obj: any = {};
    message.lockSchemaFee !== undefined &&
      (obj.lockSchemaFee = message.lockSchemaFee
        ? LockSchemaFee.toJSON(message.lockSchemaFee)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLockSchemaFeeResponse>
  ): QueryGetLockSchemaFeeResponse {
    const message = {
      ...baseQueryGetLockSchemaFeeResponse,
    } as QueryGetLockSchemaFeeResponse;
    if (object.lockSchemaFee !== undefined && object.lockSchemaFee !== null) {
      message.lockSchemaFee = LockSchemaFee.fromPartial(object.lockSchemaFee);
    } else {
      message.lockSchemaFee = undefined;
    }
    return message;
  },
};

const baseQueryAllLockSchemaFeeRequest: object = {};

export const QueryAllLockSchemaFeeRequest = {
  encode(
    message: QueryAllLockSchemaFeeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLockSchemaFeeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLockSchemaFeeRequest,
    } as QueryAllLockSchemaFeeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLockSchemaFeeRequest {
    const message = {
      ...baseQueryAllLockSchemaFeeRequest,
    } as QueryAllLockSchemaFeeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLockSchemaFeeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLockSchemaFeeRequest>
  ): QueryAllLockSchemaFeeRequest {
    const message = {
      ...baseQueryAllLockSchemaFeeRequest,
    } as QueryAllLockSchemaFeeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLockSchemaFeeResponse: object = {};

export const QueryAllLockSchemaFeeResponse = {
  encode(
    message: QueryAllLockSchemaFeeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.lockSchemaFee) {
      LockSchemaFee.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLockSchemaFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLockSchemaFeeResponse,
    } as QueryAllLockSchemaFeeResponse;
    message.lockSchemaFee = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lockSchemaFee.push(
            LockSchemaFee.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLockSchemaFeeResponse {
    const message = {
      ...baseQueryAllLockSchemaFeeResponse,
    } as QueryAllLockSchemaFeeResponse;
    message.lockSchemaFee = [];
    if (object.lockSchemaFee !== undefined && object.lockSchemaFee !== null) {
      for (const e of object.lockSchemaFee) {
        message.lockSchemaFee.push(LockSchemaFee.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLockSchemaFeeResponse): unknown {
    const obj: any = {};
    if (message.lockSchemaFee) {
      obj.lockSchemaFee = message.lockSchemaFee.map((e) =>
        e ? LockSchemaFee.toJSON(e) : undefined
      );
    } else {
      obj.lockSchemaFee = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLockSchemaFeeResponse>
  ): QueryAllLockSchemaFeeResponse {
    const message = {
      ...baseQueryAllLockSchemaFeeResponse,
    } as QueryAllLockSchemaFeeResponse;
    message.lockSchemaFee = [];
    if (object.lockSchemaFee !== undefined && object.lockSchemaFee !== null) {
      for (const e of object.lockSchemaFee) {
        message.lockSchemaFee.push(LockSchemaFee.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a NFTSchema by index. */
  NFTSchema(
    request: QueryGetNFTSchemaRequest
  ): Promise<QueryGetNFTSchemaResponse>;
  /** Queries a list of NFTSchema items. */
  NFTSchemaAll(
    request: QueryAllNFTSchemaRequest
  ): Promise<QueryAllNFTSchemaResponse>;
  /** Queries a NftData by index. */
  NftData(request: QueryGetNftDataRequest): Promise<QueryGetNftDataResponse>;
  /** Queries a list of NftData items. */
  NftDataAll(request: QueryAllNftDataRequest): Promise<QueryAllNftDataResponse>;
  /** Queries a ActionByRefId by index. */
  ActionByRefId(
    request: QueryGetActionByRefIdRequest
  ): Promise<QueryGetActionByRefIdResponse>;
  /** Queries a list of ActionByRefId items. */
  ActionByRefIdAll(
    request: QueryAllActionByRefIdRequest
  ): Promise<QueryAllActionByRefIdResponse>;
  /** Queries a Organization by index. */
  Organization(
    request: QueryGetOrganizationRequest
  ): Promise<QueryGetOrganizationResponse>;
  /** Queries a list of Organization items. */
  OrganizationAll(
    request: QueryAllOrganizationRequest
  ): Promise<QueryAllOrganizationResponse>;
  /** Queries a NftCollection by index. */
  NftCollection(
    request: QueryGetNftCollectionRequest
  ): Promise<QueryGetNftCollectionResponse>;
  /** Queries a NFTSchemaByContract by index. */
  NFTSchemaByContract(
    request: QueryGetNFTSchemaByContractRequest
  ): Promise<QueryGetNFTSchemaByContractResponse>;
  /** Queries a list of NFTSchemaByContract items. */
  NFTSchemaByContractAll(
    request: QueryAllNFTSchemaByContractRequest
  ): Promise<QueryAllNFTSchemaByContractResponse>;
  /** Queries a NFTFeeConfig by index. */
  NFTFeeConfig(
    request: QueryGetNFTFeeConfigRequest
  ): Promise<QueryGetNFTFeeConfigResponse>;
  /** Queries a NFTFeeBalance by index. */
  NFTFeeBalance(
    request: QueryGetNFTFeeBalanceRequest
  ): Promise<QueryGetNFTFeeBalanceResponse>;
  /** Queries a MetadataCreator by index. */
  MetadataCreator(
    request: QueryGetMetadataCreatorRequest
  ): Promise<QueryGetMetadataCreatorResponse>;
  /** Queries a list of MetadataCreator items. */
  MetadataCreatorAll(
    request: QueryAllMetadataCreatorRequest
  ): Promise<QueryAllMetadataCreatorResponse>;
  /** Queries a ActionExecutor by index. */
  ActionExecutor(
    request: QueryGetActionExecutorRequest
  ): Promise<QueryGetActionExecutorResponse>;
  /** Queries a list of ActionExecutor items. */
  ActionExecutorAll(
    request: QueryAllActionExecutorRequest
  ): Promise<QueryAllActionExecutorResponse>;
  /** Queries a SchemaAttribute by index. */
  SchemaAttribute(
    request: QueryGetSchemaAttributeRequest
  ): Promise<QueryGetSchemaAttributeResponse>;
  /** Queries a list of SchemaAttribute items. */
  SchemaAttributeAll(
    request: QueryAllSchemaAttributeRequest
  ): Promise<QueryAllSchemaAttributeResponse>;
  /** Queries a list of ListAttributeBySchema items. */
  ListAttributeBySchema(
    request: QueryListAttributeBySchemaRequest
  ): Promise<QueryListAttributeBySchemaResponse>;
  /** Queries a ActionOfSchema by index. */
  ActionOfSchema(
    request: QueryGetActionOfSchemaRequest
  ): Promise<QueryGetActionOfSchemaResponse>;
  /** Queries a list of ActionOfSchema items. */
  ActionOfSchemaAll(
    request: QueryAllActionOfSchemaRequest
  ): Promise<QueryAllActionOfSchemaResponse>;
  /** Queries a ExecutorOfSchema by index. */
  ExecutorOfSchema(
    request: QueryGetExecutorOfSchemaRequest
  ): Promise<QueryGetExecutorOfSchemaResponse>;
  /** Queries a list of ExecutorOfSchema items. */
  ExecutorOfSchemaAll(
    request: QueryAllExecutorOfSchemaRequest
  ): Promise<QueryAllExecutorOfSchemaResponse>;
  /** Queries a Virtual by index */
  VirtualAction(
    request: QueryGetVirtualActionRequest
  ): Promise<QueryGetVirtualActionResponse>;
  /** Queries a list of Virtual items. */
  VirtualActionAll(
    request: QueryAllVirtualActionRequest
  ): Promise<QueryAllVirtualActionResponse>;
  /** Queries a VirtualSchema by index. */
  VirtualSchema(
    request: QueryGetVirtualSchemaRequest
  ): Promise<QueryGetVirtualSchemaResponse>;
  /** Queries a list of VirtualSchema items. */
  VirtualSchemaAll(
    request: QueryAllVirtualSchemaRequest
  ): Promise<QueryAllVirtualSchemaResponse>;
  /** Queries a VirtualSchemaProposal by index. */
  VirtualSchemaProposal(
    request: QueryGetVirtualSchemaProposalRequest
  ): Promise<QueryGetVirtualSchemaProposalResponse>;
  /** Queries a list of VirtualSchemaProposal items. */
  VirtualSchemaProposalAll(
    request: QueryAllVirtualSchemaProposalRequest
  ): Promise<QueryAllVirtualSchemaProposalResponse>;
  /** Queries a list of ListActiveProposal items. */
  ListActiveProposal(
    request: QueryListActiveProposalRequest
  ): Promise<QueryListActiveProposalResponse>;
  /** Queries a LockSchemaFee by index. */
  LockSchemaFee(
    request: QueryGetLockSchemaFeeRequest
  ): Promise<QueryGetLockSchemaFeeResponse>;
  /** Queries a list of LockSchemaFee items. */
  LockSchemaFeeAll(
    request: QueryAllLockSchemaFeeRequest
  ): Promise<QueryAllLockSchemaFeeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NFTSchema(
    request: QueryGetNFTSchemaRequest
  ): Promise<QueryGetNFTSchemaResponse> {
    const data = QueryGetNFTSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTSchema",
      data
    );
    return promise.then((data) =>
      QueryGetNFTSchemaResponse.decode(new Reader(data))
    );
  }

  NFTSchemaAll(
    request: QueryAllNFTSchemaRequest
  ): Promise<QueryAllNFTSchemaResponse> {
    const data = QueryAllNFTSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTSchemaAll",
      data
    );
    return promise.then((data) =>
      QueryAllNFTSchemaResponse.decode(new Reader(data))
    );
  }

  NftData(request: QueryGetNftDataRequest): Promise<QueryGetNftDataResponse> {
    const data = QueryGetNftDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NftData",
      data
    );
    return promise.then((data) =>
      QueryGetNftDataResponse.decode(new Reader(data))
    );
  }

  NftDataAll(
    request: QueryAllNftDataRequest
  ): Promise<QueryAllNftDataResponse> {
    const data = QueryAllNftDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NftDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllNftDataResponse.decode(new Reader(data))
    );
  }

  ActionByRefId(
    request: QueryGetActionByRefIdRequest
  ): Promise<QueryGetActionByRefIdResponse> {
    const data = QueryGetActionByRefIdRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionByRefId",
      data
    );
    return promise.then((data) =>
      QueryGetActionByRefIdResponse.decode(new Reader(data))
    );
  }

  ActionByRefIdAll(
    request: QueryAllActionByRefIdRequest
  ): Promise<QueryAllActionByRefIdResponse> {
    const data = QueryAllActionByRefIdRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionByRefIdAll",
      data
    );
    return promise.then((data) =>
      QueryAllActionByRefIdResponse.decode(new Reader(data))
    );
  }

  Organization(
    request: QueryGetOrganizationRequest
  ): Promise<QueryGetOrganizationResponse> {
    const data = QueryGetOrganizationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "Organization",
      data
    );
    return promise.then((data) =>
      QueryGetOrganizationResponse.decode(new Reader(data))
    );
  }

  OrganizationAll(
    request: QueryAllOrganizationRequest
  ): Promise<QueryAllOrganizationResponse> {
    const data = QueryAllOrganizationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "OrganizationAll",
      data
    );
    return promise.then((data) =>
      QueryAllOrganizationResponse.decode(new Reader(data))
    );
  }

  NftCollection(
    request: QueryGetNftCollectionRequest
  ): Promise<QueryGetNftCollectionResponse> {
    const data = QueryGetNftCollectionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NftCollection",
      data
    );
    return promise.then((data) =>
      QueryGetNftCollectionResponse.decode(new Reader(data))
    );
  }

  NFTSchemaByContract(
    request: QueryGetNFTSchemaByContractRequest
  ): Promise<QueryGetNFTSchemaByContractResponse> {
    const data = QueryGetNFTSchemaByContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTSchemaByContract",
      data
    );
    return promise.then((data) =>
      QueryGetNFTSchemaByContractResponse.decode(new Reader(data))
    );
  }

  NFTSchemaByContractAll(
    request: QueryAllNFTSchemaByContractRequest
  ): Promise<QueryAllNFTSchemaByContractResponse> {
    const data = QueryAllNFTSchemaByContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTSchemaByContractAll",
      data
    );
    return promise.then((data) =>
      QueryAllNFTSchemaByContractResponse.decode(new Reader(data))
    );
  }

  NFTFeeConfig(
    request: QueryGetNFTFeeConfigRequest
  ): Promise<QueryGetNFTFeeConfigResponse> {
    const data = QueryGetNFTFeeConfigRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTFeeConfig",
      data
    );
    return promise.then((data) =>
      QueryGetNFTFeeConfigResponse.decode(new Reader(data))
    );
  }

  NFTFeeBalance(
    request: QueryGetNFTFeeBalanceRequest
  ): Promise<QueryGetNFTFeeBalanceResponse> {
    const data = QueryGetNFTFeeBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "NFTFeeBalance",
      data
    );
    return promise.then((data) =>
      QueryGetNFTFeeBalanceResponse.decode(new Reader(data))
    );
  }

  MetadataCreator(
    request: QueryGetMetadataCreatorRequest
  ): Promise<QueryGetMetadataCreatorResponse> {
    const data = QueryGetMetadataCreatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "MetadataCreator",
      data
    );
    return promise.then((data) =>
      QueryGetMetadataCreatorResponse.decode(new Reader(data))
    );
  }

  MetadataCreatorAll(
    request: QueryAllMetadataCreatorRequest
  ): Promise<QueryAllMetadataCreatorResponse> {
    const data = QueryAllMetadataCreatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "MetadataCreatorAll",
      data
    );
    return promise.then((data) =>
      QueryAllMetadataCreatorResponse.decode(new Reader(data))
    );
  }

  ActionExecutor(
    request: QueryGetActionExecutorRequest
  ): Promise<QueryGetActionExecutorResponse> {
    const data = QueryGetActionExecutorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionExecutor",
      data
    );
    return promise.then((data) =>
      QueryGetActionExecutorResponse.decode(new Reader(data))
    );
  }

  ActionExecutorAll(
    request: QueryAllActionExecutorRequest
  ): Promise<QueryAllActionExecutorResponse> {
    const data = QueryAllActionExecutorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionExecutorAll",
      data
    );
    return promise.then((data) =>
      QueryAllActionExecutorResponse.decode(new Reader(data))
    );
  }

  SchemaAttribute(
    request: QueryGetSchemaAttributeRequest
  ): Promise<QueryGetSchemaAttributeResponse> {
    const data = QueryGetSchemaAttributeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "SchemaAttribute",
      data
    );
    return promise.then((data) =>
      QueryGetSchemaAttributeResponse.decode(new Reader(data))
    );
  }

  SchemaAttributeAll(
    request: QueryAllSchemaAttributeRequest
  ): Promise<QueryAllSchemaAttributeResponse> {
    const data = QueryAllSchemaAttributeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "SchemaAttributeAll",
      data
    );
    return promise.then((data) =>
      QueryAllSchemaAttributeResponse.decode(new Reader(data))
    );
  }

  ListAttributeBySchema(
    request: QueryListAttributeBySchemaRequest
  ): Promise<QueryListAttributeBySchemaResponse> {
    const data = QueryListAttributeBySchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ListAttributeBySchema",
      data
    );
    return promise.then((data) =>
      QueryListAttributeBySchemaResponse.decode(new Reader(data))
    );
  }

  ActionOfSchema(
    request: QueryGetActionOfSchemaRequest
  ): Promise<QueryGetActionOfSchemaResponse> {
    const data = QueryGetActionOfSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionOfSchema",
      data
    );
    return promise.then((data) =>
      QueryGetActionOfSchemaResponse.decode(new Reader(data))
    );
  }

  ActionOfSchemaAll(
    request: QueryAllActionOfSchemaRequest
  ): Promise<QueryAllActionOfSchemaResponse> {
    const data = QueryAllActionOfSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ActionOfSchemaAll",
      data
    );
    return promise.then((data) =>
      QueryAllActionOfSchemaResponse.decode(new Reader(data))
    );
  }

  ExecutorOfSchema(
    request: QueryGetExecutorOfSchemaRequest
  ): Promise<QueryGetExecutorOfSchemaResponse> {
    const data = QueryGetExecutorOfSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ExecutorOfSchema",
      data
    );
    return promise.then((data) =>
      QueryGetExecutorOfSchemaResponse.decode(new Reader(data))
    );
  }

  ExecutorOfSchemaAll(
    request: QueryAllExecutorOfSchemaRequest
  ): Promise<QueryAllExecutorOfSchemaResponse> {
    const data = QueryAllExecutorOfSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ExecutorOfSchemaAll",
      data
    );
    return promise.then((data) =>
      QueryAllExecutorOfSchemaResponse.decode(new Reader(data))
    );
  }

  VirtualAction(
    request: QueryGetVirtualActionRequest
  ): Promise<QueryGetVirtualActionResponse> {
    const data = QueryGetVirtualActionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualAction",
      data
    );
    return promise.then((data) =>
      QueryGetVirtualActionResponse.decode(new Reader(data))
    );
  }

  VirtualActionAll(
    request: QueryAllVirtualActionRequest
  ): Promise<QueryAllVirtualActionResponse> {
    const data = QueryAllVirtualActionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualActionAll",
      data
    );
    return promise.then((data) =>
      QueryAllVirtualActionResponse.decode(new Reader(data))
    );
  }

  VirtualSchema(
    request: QueryGetVirtualSchemaRequest
  ): Promise<QueryGetVirtualSchemaResponse> {
    const data = QueryGetVirtualSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualSchema",
      data
    );
    return promise.then((data) =>
      QueryGetVirtualSchemaResponse.decode(new Reader(data))
    );
  }

  VirtualSchemaAll(
    request: QueryAllVirtualSchemaRequest
  ): Promise<QueryAllVirtualSchemaResponse> {
    const data = QueryAllVirtualSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualSchemaAll",
      data
    );
    return promise.then((data) =>
      QueryAllVirtualSchemaResponse.decode(new Reader(data))
    );
  }

  VirtualSchemaProposal(
    request: QueryGetVirtualSchemaProposalRequest
  ): Promise<QueryGetVirtualSchemaProposalResponse> {
    const data = QueryGetVirtualSchemaProposalRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualSchemaProposal",
      data
    );
    return promise.then((data) =>
      QueryGetVirtualSchemaProposalResponse.decode(new Reader(data))
    );
  }

  VirtualSchemaProposalAll(
    request: QueryAllVirtualSchemaProposalRequest
  ): Promise<QueryAllVirtualSchemaProposalResponse> {
    const data = QueryAllVirtualSchemaProposalRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "VirtualSchemaProposalAll",
      data
    );
    return promise.then((data) =>
      QueryAllVirtualSchemaProposalResponse.decode(new Reader(data))
    );
  }

  ListActiveProposal(
    request: QueryListActiveProposalRequest
  ): Promise<QueryListActiveProposalResponse> {
    const data = QueryListActiveProposalRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "ListActiveProposal",
      data
    );
    return promise.then((data) =>
      QueryListActiveProposalResponse.decode(new Reader(data))
    );
  }

  LockSchemaFee(
    request: QueryGetLockSchemaFeeRequest
  ): Promise<QueryGetLockSchemaFeeResponse> {
    const data = QueryGetLockSchemaFeeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "LockSchemaFee",
      data
    );
    return promise.then((data) =>
      QueryGetLockSchemaFeeResponse.decode(new Reader(data))
    );
  }

  LockSchemaFeeAll(
    request: QueryAllLockSchemaFeeRequest
  ): Promise<QueryAllLockSchemaFeeResponse> {
    const data = QueryAllLockSchemaFeeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.nftmngr.Query",
      "LockSchemaFeeAll",
      data
    );
    return promise.then((data) =>
      QueryAllLockSchemaFeeResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
