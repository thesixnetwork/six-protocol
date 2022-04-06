/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../tokenmngr/params";
import { Token } from "../tokenmngr/token";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Mintperm } from "../tokenmngr/mintperm";
import { Options } from "../tokenmngr/options";

export const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTokenRequest {
  name: string;
}

export interface QueryGetTokenResponse {
  token: Token | undefined;
}

export interface QueryAllTokenRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokenResponse {
  token: Token[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMintpermRequest {
  token: string;
  address: string;
}

export interface QueryGetMintpermResponse {
  mintperm: Mintperm | undefined;
}

export interface QueryAllMintpermRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMintpermResponse {
  mintperm: Mintperm[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOptionsRequest {}

export interface QueryGetOptionsResponse {
  Options: Options | undefined;
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

const baseQueryGetTokenRequest: object = { name: "" };

export const QueryGetTokenRequest = {
  encode(
    message: QueryGetTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
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

  fromJSON(object: any): QueryGetTokenRequest {
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetTokenRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetTokenRequest>): QueryGetTokenRequest {
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetTokenResponse: object = {};

export const QueryGetTokenResponse = {
  encode(
    message: QueryGetTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token !== undefined) {
      Token.encode(message.token, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = Token.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenResponse {
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokenResponse): unknown {
    const obj: any = {};
    message.token !== undefined &&
      (obj.token = message.token ? Token.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenResponse>
  ): QueryGetTokenResponse {
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenRequest: object = {};

export const QueryAllTokenRequest = {
  encode(
    message: QueryAllTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
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

  fromJSON(object: any): QueryAllTokenRequest {
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllTokenRequest>): QueryAllTokenRequest {
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenResponse: object = {};

export const QueryAllTokenResponse = {
  encode(
    message: QueryAllTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.token) {
      Token.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token.push(Token.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTokenResponse {
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    if (object.token !== undefined && object.token !== null) {
      for (const e of object.token) {
        message.token.push(Token.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenResponse): unknown {
    const obj: any = {};
    if (message.token) {
      obj.token = message.token.map((e) => (e ? Token.toJSON(e) : undefined));
    } else {
      obj.token = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenResponse>
  ): QueryAllTokenResponse {
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    if (object.token !== undefined && object.token !== null) {
      for (const e of object.token) {
        message.token.push(Token.fromPartial(e));
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

const baseQueryGetMintpermRequest: object = { token: "", address: "" };

export const QueryGetMintpermRequest = {
  encode(
    message: QueryGetMintpermRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMintpermRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintpermRequest,
    } as QueryGetMintpermRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMintpermRequest {
    const message = {
      ...baseQueryGetMintpermRequest,
    } as QueryGetMintpermRequest;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetMintpermRequest): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMintpermRequest>
  ): QueryGetMintpermRequest {
    const message = {
      ...baseQueryGetMintpermRequest,
    } as QueryGetMintpermRequest;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetMintpermResponse: object = {};

export const QueryGetMintpermResponse = {
  encode(
    message: QueryGetMintpermResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mintperm !== undefined) {
      Mintperm.encode(message.mintperm, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMintpermResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintpermResponse,
    } as QueryGetMintpermResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mintperm = Mintperm.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMintpermResponse {
    const message = {
      ...baseQueryGetMintpermResponse,
    } as QueryGetMintpermResponse;
    if (object.mintperm !== undefined && object.mintperm !== null) {
      message.mintperm = Mintperm.fromJSON(object.mintperm);
    } else {
      message.mintperm = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMintpermResponse): unknown {
    const obj: any = {};
    message.mintperm !== undefined &&
      (obj.mintperm = message.mintperm
        ? Mintperm.toJSON(message.mintperm)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMintpermResponse>
  ): QueryGetMintpermResponse {
    const message = {
      ...baseQueryGetMintpermResponse,
    } as QueryGetMintpermResponse;
    if (object.mintperm !== undefined && object.mintperm !== null) {
      message.mintperm = Mintperm.fromPartial(object.mintperm);
    } else {
      message.mintperm = undefined;
    }
    return message;
  },
};

const baseQueryAllMintpermRequest: object = {};

export const QueryAllMintpermRequest = {
  encode(
    message: QueryAllMintpermRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMintpermRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMintpermRequest,
    } as QueryAllMintpermRequest;
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

  fromJSON(object: any): QueryAllMintpermRequest {
    const message = {
      ...baseQueryAllMintpermRequest,
    } as QueryAllMintpermRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMintpermRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMintpermRequest>
  ): QueryAllMintpermRequest {
    const message = {
      ...baseQueryAllMintpermRequest,
    } as QueryAllMintpermRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMintpermResponse: object = {};

export const QueryAllMintpermResponse = {
  encode(
    message: QueryAllMintpermResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.mintperm) {
      Mintperm.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllMintpermResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMintpermResponse,
    } as QueryAllMintpermResponse;
    message.mintperm = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mintperm.push(Mintperm.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllMintpermResponse {
    const message = {
      ...baseQueryAllMintpermResponse,
    } as QueryAllMintpermResponse;
    message.mintperm = [];
    if (object.mintperm !== undefined && object.mintperm !== null) {
      for (const e of object.mintperm) {
        message.mintperm.push(Mintperm.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMintpermResponse): unknown {
    const obj: any = {};
    if (message.mintperm) {
      obj.mintperm = message.mintperm.map((e) =>
        e ? Mintperm.toJSON(e) : undefined
      );
    } else {
      obj.mintperm = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMintpermResponse>
  ): QueryAllMintpermResponse {
    const message = {
      ...baseQueryAllMintpermResponse,
    } as QueryAllMintpermResponse;
    message.mintperm = [];
    if (object.mintperm !== undefined && object.mintperm !== null) {
      for (const e of object.mintperm) {
        message.mintperm.push(Mintperm.fromPartial(e));
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

const baseQueryGetOptionsRequest: object = {};

export const QueryGetOptionsRequest = {
  encode(_: QueryGetOptionsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOptionsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOptionsRequest } as QueryGetOptionsRequest;
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

  fromJSON(_: any): QueryGetOptionsRequest {
    const message = { ...baseQueryGetOptionsRequest } as QueryGetOptionsRequest;
    return message;
  },

  toJSON(_: QueryGetOptionsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetOptionsRequest>): QueryGetOptionsRequest {
    const message = { ...baseQueryGetOptionsRequest } as QueryGetOptionsRequest;
    return message;
  },
};

const baseQueryGetOptionsResponse: object = {};

export const QueryGetOptionsResponse = {
  encode(
    message: QueryGetOptionsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Options !== undefined) {
      Options.encode(message.Options, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOptionsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOptionsResponse,
    } as QueryGetOptionsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Options = Options.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOptionsResponse {
    const message = {
      ...baseQueryGetOptionsResponse,
    } as QueryGetOptionsResponse;
    if (object.Options !== undefined && object.Options !== null) {
      message.Options = Options.fromJSON(object.Options);
    } else {
      message.Options = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOptionsResponse): unknown {
    const obj: any = {};
    message.Options !== undefined &&
      (obj.Options = message.Options
        ? Options.toJSON(message.Options)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOptionsResponse>
  ): QueryGetOptionsResponse {
    const message = {
      ...baseQueryGetOptionsResponse,
    } as QueryGetOptionsResponse;
    if (object.Options !== undefined && object.Options !== null) {
      message.Options = Options.fromPartial(object.Options);
    } else {
      message.Options = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Token by index. */
  Token(request: QueryGetTokenRequest): Promise<QueryGetTokenResponse>;
  /** Queries a list of Token items. */
  TokenAll(request: QueryAllTokenRequest): Promise<QueryAllTokenResponse>;
  /** Queries a Mintperm by index. */
  Mintperm(request: QueryGetMintpermRequest): Promise<QueryGetMintpermResponse>;
  /** Queries a list of Mintperm items. */
  MintpermAll(
    request: QueryAllMintpermRequest
  ): Promise<QueryAllMintpermResponse>;
  /** Queries a Options by index. */
  Options(request: QueryGetOptionsRequest): Promise<QueryGetOptionsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Token(request: QueryGetTokenRequest): Promise<QueryGetTokenResponse> {
    const data = QueryGetTokenRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "Token",
      data
    );
    return promise.then((data) =>
      QueryGetTokenResponse.decode(new Reader(data))
    );
  }

  TokenAll(request: QueryAllTokenRequest): Promise<QueryAllTokenResponse> {
    const data = QueryAllTokenRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "TokenAll",
      data
    );
    return promise.then((data) =>
      QueryAllTokenResponse.decode(new Reader(data))
    );
  }

  Mintperm(
    request: QueryGetMintpermRequest
  ): Promise<QueryGetMintpermResponse> {
    const data = QueryGetMintpermRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "Mintperm",
      data
    );
    return promise.then((data) =>
      QueryGetMintpermResponse.decode(new Reader(data))
    );
  }

  MintpermAll(
    request: QueryAllMintpermRequest
  ): Promise<QueryAllMintpermResponse> {
    const data = QueryAllMintpermRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "MintpermAll",
      data
    );
    return promise.then((data) =>
      QueryAllMintpermResponse.decode(new Reader(data))
    );
  }

  Options(request: QueryGetOptionsRequest): Promise<QueryGetOptionsResponse> {
    const data = QueryGetOptionsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.tokenmngr.Query",
      "Options",
      data
    );
    return promise.then((data) =>
      QueryGetOptionsResponse.decode(new Reader(data))
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
