/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../evmbind/params";
import { Binding } from "../evmbind/binding";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "thesixnetwork.sixprotocol.evmbind";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetBindingRequest {
  ethAddress: string;
}

export interface QueryGetBindingResponse {
  binding: Binding | undefined;
}

export interface QueryAllBindingRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBindingResponse {
  binding: Binding[];
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

const baseQueryGetBindingRequest: object = { ethAddress: "" };

export const QueryGetBindingRequest = {
  encode(
    message: QueryGetBindingRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ethAddress !== "") {
      writer.uint32(10).string(message.ethAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBindingRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBindingRequest } as QueryGetBindingRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBindingRequest {
    const message = { ...baseQueryGetBindingRequest } as QueryGetBindingRequest;
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = String(object.ethAddress);
    } else {
      message.ethAddress = "";
    }
    return message;
  },

  toJSON(message: QueryGetBindingRequest): unknown {
    const obj: any = {};
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBindingRequest>
  ): QueryGetBindingRequest {
    const message = { ...baseQueryGetBindingRequest } as QueryGetBindingRequest;
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = object.ethAddress;
    } else {
      message.ethAddress = "";
    }
    return message;
  },
};

const baseQueryGetBindingResponse: object = {};

export const QueryGetBindingResponse = {
  encode(
    message: QueryGetBindingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.binding !== undefined) {
      Binding.encode(message.binding, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBindingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBindingResponse,
    } as QueryGetBindingResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.binding = Binding.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBindingResponse {
    const message = {
      ...baseQueryGetBindingResponse,
    } as QueryGetBindingResponse;
    if (object.binding !== undefined && object.binding !== null) {
      message.binding = Binding.fromJSON(object.binding);
    } else {
      message.binding = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBindingResponse): unknown {
    const obj: any = {};
    message.binding !== undefined &&
      (obj.binding = message.binding
        ? Binding.toJSON(message.binding)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBindingResponse>
  ): QueryGetBindingResponse {
    const message = {
      ...baseQueryGetBindingResponse,
    } as QueryGetBindingResponse;
    if (object.binding !== undefined && object.binding !== null) {
      message.binding = Binding.fromPartial(object.binding);
    } else {
      message.binding = undefined;
    }
    return message;
  },
};

const baseQueryAllBindingRequest: object = {};

export const QueryAllBindingRequest = {
  encode(
    message: QueryAllBindingRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBindingRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBindingRequest } as QueryAllBindingRequest;
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

  fromJSON(object: any): QueryAllBindingRequest {
    const message = { ...baseQueryAllBindingRequest } as QueryAllBindingRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBindingRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBindingRequest>
  ): QueryAllBindingRequest {
    const message = { ...baseQueryAllBindingRequest } as QueryAllBindingRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBindingResponse: object = {};

export const QueryAllBindingResponse = {
  encode(
    message: QueryAllBindingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.binding) {
      Binding.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBindingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBindingResponse,
    } as QueryAllBindingResponse;
    message.binding = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.binding.push(Binding.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBindingResponse {
    const message = {
      ...baseQueryAllBindingResponse,
    } as QueryAllBindingResponse;
    message.binding = [];
    if (object.binding !== undefined && object.binding !== null) {
      for (const e of object.binding) {
        message.binding.push(Binding.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBindingResponse): unknown {
    const obj: any = {};
    if (message.binding) {
      obj.binding = message.binding.map((e) =>
        e ? Binding.toJSON(e) : undefined
      );
    } else {
      obj.binding = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBindingResponse>
  ): QueryAllBindingResponse {
    const message = {
      ...baseQueryAllBindingResponse,
    } as QueryAllBindingResponse;
    message.binding = [];
    if (object.binding !== undefined && object.binding !== null) {
      for (const e of object.binding) {
        message.binding.push(Binding.fromPartial(e));
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
  /** Queries a Binding by index. */
  Binding(request: QueryGetBindingRequest): Promise<QueryGetBindingResponse>;
  /** Queries a list of Binding items. */
  BindingAll(request: QueryAllBindingRequest): Promise<QueryAllBindingResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Binding(request: QueryGetBindingRequest): Promise<QueryGetBindingResponse> {
    const data = QueryGetBindingRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Query",
      "Binding",
      data
    );
    return promise.then((data) =>
      QueryGetBindingResponse.decode(new Reader(data))
    );
  }

  BindingAll(
    request: QueryAllBindingRequest
  ): Promise<QueryAllBindingResponse> {
    const data = QueryAllBindingRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Query",
      "BindingAll",
      data
    );
    return promise.then((data) =>
      QueryAllBindingResponse.decode(new Reader(data))
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
