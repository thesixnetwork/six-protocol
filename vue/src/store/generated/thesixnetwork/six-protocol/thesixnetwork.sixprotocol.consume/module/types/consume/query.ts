/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../consume/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { UseNft } from "../consume/use_nft";
import { NftUsed } from "../consume/nft_used";

export const protobufPackage = "thesixnetwork.sixprotocol.consume";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryConsumeNftsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryConsumeNftsResponse {
  UseNft: UseNft[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNftUsedRequest {
  token: string;
  creator: string;
}

export interface QueryGetNftUsedResponse {
  nftUsed: NftUsed | undefined;
}

export interface QueryAllNftUsedRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNftUsedResponse {
  nftUsed: NftUsed[];
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

const baseQueryConsumeNftsRequest: object = {};

export const QueryConsumeNftsRequest = {
  encode(
    message: QueryConsumeNftsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryConsumeNftsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryConsumeNftsRequest,
    } as QueryConsumeNftsRequest;
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

  fromJSON(object: any): QueryConsumeNftsRequest {
    const message = {
      ...baseQueryConsumeNftsRequest,
    } as QueryConsumeNftsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryConsumeNftsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryConsumeNftsRequest>
  ): QueryConsumeNftsRequest {
    const message = {
      ...baseQueryConsumeNftsRequest,
    } as QueryConsumeNftsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryConsumeNftsResponse: object = {};

export const QueryConsumeNftsResponse = {
  encode(
    message: QueryConsumeNftsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.UseNft) {
      UseNft.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryConsumeNftsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryConsumeNftsResponse,
    } as QueryConsumeNftsResponse;
    message.UseNft = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.UseNft.push(UseNft.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryConsumeNftsResponse {
    const message = {
      ...baseQueryConsumeNftsResponse,
    } as QueryConsumeNftsResponse;
    message.UseNft = [];
    if (object.UseNft !== undefined && object.UseNft !== null) {
      for (const e of object.UseNft) {
        message.UseNft.push(UseNft.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryConsumeNftsResponse): unknown {
    const obj: any = {};
    if (message.UseNft) {
      obj.UseNft = message.UseNft.map((e) =>
        e ? UseNft.toJSON(e) : undefined
      );
    } else {
      obj.UseNft = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryConsumeNftsResponse>
  ): QueryConsumeNftsResponse {
    const message = {
      ...baseQueryConsumeNftsResponse,
    } as QueryConsumeNftsResponse;
    message.UseNft = [];
    if (object.UseNft !== undefined && object.UseNft !== null) {
      for (const e of object.UseNft) {
        message.UseNft.push(UseNft.fromPartial(e));
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

const baseQueryGetNftUsedRequest: object = { token: "", creator: "" };

export const QueryGetNftUsedRequest = {
  encode(
    message: QueryGetNftUsedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNftUsedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNftUsedRequest } as QueryGetNftUsedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftUsedRequest {
    const message = { ...baseQueryGetNftUsedRequest } as QueryGetNftUsedRequest;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: QueryGetNftUsedRequest): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftUsedRequest>
  ): QueryGetNftUsedRequest {
    const message = { ...baseQueryGetNftUsedRequest } as QueryGetNftUsedRequest;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseQueryGetNftUsedResponse: object = {};

export const QueryGetNftUsedResponse = {
  encode(
    message: QueryGetNftUsedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftUsed !== undefined) {
      NftUsed.encode(message.nftUsed, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNftUsedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNftUsedResponse,
    } as QueryGetNftUsedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftUsed = NftUsed.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNftUsedResponse {
    const message = {
      ...baseQueryGetNftUsedResponse,
    } as QueryGetNftUsedResponse;
    if (object.nftUsed !== undefined && object.nftUsed !== null) {
      message.nftUsed = NftUsed.fromJSON(object.nftUsed);
    } else {
      message.nftUsed = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNftUsedResponse): unknown {
    const obj: any = {};
    message.nftUsed !== undefined &&
      (obj.nftUsed = message.nftUsed
        ? NftUsed.toJSON(message.nftUsed)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNftUsedResponse>
  ): QueryGetNftUsedResponse {
    const message = {
      ...baseQueryGetNftUsedResponse,
    } as QueryGetNftUsedResponse;
    if (object.nftUsed !== undefined && object.nftUsed !== null) {
      message.nftUsed = NftUsed.fromPartial(object.nftUsed);
    } else {
      message.nftUsed = undefined;
    }
    return message;
  },
};

const baseQueryAllNftUsedRequest: object = {};

export const QueryAllNftUsedRequest = {
  encode(
    message: QueryAllNftUsedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNftUsedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNftUsedRequest } as QueryAllNftUsedRequest;
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

  fromJSON(object: any): QueryAllNftUsedRequest {
    const message = { ...baseQueryAllNftUsedRequest } as QueryAllNftUsedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNftUsedRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNftUsedRequest>
  ): QueryAllNftUsedRequest {
    const message = { ...baseQueryAllNftUsedRequest } as QueryAllNftUsedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNftUsedResponse: object = {};

export const QueryAllNftUsedResponse = {
  encode(
    message: QueryAllNftUsedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nftUsed) {
      NftUsed.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNftUsedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNftUsedResponse,
    } as QueryAllNftUsedResponse;
    message.nftUsed = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftUsed.push(NftUsed.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllNftUsedResponse {
    const message = {
      ...baseQueryAllNftUsedResponse,
    } as QueryAllNftUsedResponse;
    message.nftUsed = [];
    if (object.nftUsed !== undefined && object.nftUsed !== null) {
      for (const e of object.nftUsed) {
        message.nftUsed.push(NftUsed.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNftUsedResponse): unknown {
    const obj: any = {};
    if (message.nftUsed) {
      obj.nftUsed = message.nftUsed.map((e) =>
        e ? NftUsed.toJSON(e) : undefined
      );
    } else {
      obj.nftUsed = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNftUsedResponse>
  ): QueryAllNftUsedResponse {
    const message = {
      ...baseQueryAllNftUsedResponse,
    } as QueryAllNftUsedResponse;
    message.nftUsed = [];
    if (object.nftUsed !== undefined && object.nftUsed !== null) {
      for (const e of object.nftUsed) {
        message.nftUsed.push(NftUsed.fromPartial(e));
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
  /** Queries a list of ConsumeNfts items. */
  ConsumeNfts(
    request: QueryConsumeNftsRequest
  ): Promise<QueryConsumeNftsResponse>;
  /** Queries a NftUsed by index. */
  NftUsed(request: QueryGetNftUsedRequest): Promise<QueryGetNftUsedResponse>;
  /** Queries a list of NftUsed items. */
  NftUsedAll(request: QueryAllNftUsedRequest): Promise<QueryAllNftUsedResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  ConsumeNfts(
    request: QueryConsumeNftsRequest
  ): Promise<QueryConsumeNftsResponse> {
    const data = QueryConsumeNftsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Query",
      "ConsumeNfts",
      data
    );
    return promise.then((data) =>
      QueryConsumeNftsResponse.decode(new Reader(data))
    );
  }

  NftUsed(request: QueryGetNftUsedRequest): Promise<QueryGetNftUsedResponse> {
    const data = QueryGetNftUsedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Query",
      "NftUsed",
      data
    );
    return promise.then((data) =>
      QueryGetNftUsedResponse.decode(new Reader(data))
    );
  }

  NftUsedAll(
    request: QueryAllNftUsedRequest
  ): Promise<QueryAllNftUsedResponse> {
    const data = QueryAllNftUsedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Query",
      "NftUsedAll",
      data
    );
    return promise.then((data) =>
      QueryAllNftUsedResponse.decode(new Reader(data))
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
