import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../tokenmngr/params";
import { Token } from "../tokenmngr/token";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Mintperm } from "../tokenmngr/mintperm";
import { Options } from "../tokenmngr/options";
export declare const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
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
export interface QueryGetOptionsRequest {
}
export interface QueryGetOptionsResponse {
    Options: Options | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetTokenRequest: {
    encode(message: QueryGetTokenRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTokenRequest;
    fromJSON(object: any): QueryGetTokenRequest;
    toJSON(message: QueryGetTokenRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTokenRequest>): QueryGetTokenRequest;
};
export declare const QueryGetTokenResponse: {
    encode(message: QueryGetTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTokenResponse;
    fromJSON(object: any): QueryGetTokenResponse;
    toJSON(message: QueryGetTokenResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTokenResponse>): QueryGetTokenResponse;
};
export declare const QueryAllTokenRequest: {
    encode(message: QueryAllTokenRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTokenRequest;
    fromJSON(object: any): QueryAllTokenRequest;
    toJSON(message: QueryAllTokenRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTokenRequest>): QueryAllTokenRequest;
};
export declare const QueryAllTokenResponse: {
    encode(message: QueryAllTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTokenResponse;
    fromJSON(object: any): QueryAllTokenResponse;
    toJSON(message: QueryAllTokenResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTokenResponse>): QueryAllTokenResponse;
};
export declare const QueryGetMintpermRequest: {
    encode(message: QueryGetMintpermRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMintpermRequest;
    fromJSON(object: any): QueryGetMintpermRequest;
    toJSON(message: QueryGetMintpermRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetMintpermRequest>): QueryGetMintpermRequest;
};
export declare const QueryGetMintpermResponse: {
    encode(message: QueryGetMintpermResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMintpermResponse;
    fromJSON(object: any): QueryGetMintpermResponse;
    toJSON(message: QueryGetMintpermResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetMintpermResponse>): QueryGetMintpermResponse;
};
export declare const QueryAllMintpermRequest: {
    encode(message: QueryAllMintpermRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMintpermRequest;
    fromJSON(object: any): QueryAllMintpermRequest;
    toJSON(message: QueryAllMintpermRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllMintpermRequest>): QueryAllMintpermRequest;
};
export declare const QueryAllMintpermResponse: {
    encode(message: QueryAllMintpermResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMintpermResponse;
    fromJSON(object: any): QueryAllMintpermResponse;
    toJSON(message: QueryAllMintpermResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllMintpermResponse>): QueryAllMintpermResponse;
};
export declare const QueryGetOptionsRequest: {
    encode(_: QueryGetOptionsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOptionsRequest;
    fromJSON(_: any): QueryGetOptionsRequest;
    toJSON(_: QueryGetOptionsRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetOptionsRequest>): QueryGetOptionsRequest;
};
export declare const QueryGetOptionsResponse: {
    encode(message: QueryGetOptionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOptionsResponse;
    fromJSON(object: any): QueryGetOptionsResponse;
    toJSON(message: QueryGetOptionsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetOptionsResponse>): QueryGetOptionsResponse;
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
    MintpermAll(request: QueryAllMintpermRequest): Promise<QueryAllMintpermResponse>;
    /** Queries a Options by index. */
    Options(request: QueryGetOptionsRequest): Promise<QueryGetOptionsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Token(request: QueryGetTokenRequest): Promise<QueryGetTokenResponse>;
    TokenAll(request: QueryAllTokenRequest): Promise<QueryAllTokenResponse>;
    Mintperm(request: QueryGetMintpermRequest): Promise<QueryGetMintpermResponse>;
    MintpermAll(request: QueryAllMintpermRequest): Promise<QueryAllMintpermResponse>;
    Options(request: QueryGetOptionsRequest): Promise<QueryGetOptionsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
