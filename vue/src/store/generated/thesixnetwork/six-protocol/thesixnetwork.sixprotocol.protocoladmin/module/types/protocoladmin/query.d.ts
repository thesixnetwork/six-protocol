import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../protocoladmin/params";
import { Group } from "../protocoladmin/group";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Admin } from "../protocoladmin/admin";
export declare const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetGroupRequest {
    name: string;
}
export interface QueryGetGroupResponse {
    group: Group | undefined;
}
export interface QueryAllGroupRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllGroupResponse {
    group: Group[];
    pagination: PageResponse | undefined;
}
export interface QueryGetAdminRequest {
    group: string;
    admin: string;
}
export interface QueryGetAdminResponse {
    admin: Admin | undefined;
}
export interface QueryAllAdminRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllAdminResponse {
    admin: Admin[];
    pagination: PageResponse | undefined;
}
export interface QueryListAdminOfGroupRequest {
    group: string;
    pagination: PageRequest | undefined;
}
export interface QueryListAdminOfGroupResponse {
    admin: string[];
    pagination: PageResponse | undefined;
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
export declare const QueryGetGroupRequest: {
    encode(message: QueryGetGroupRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGroupRequest;
    fromJSON(object: any): QueryGetGroupRequest;
    toJSON(message: QueryGetGroupRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetGroupRequest>): QueryGetGroupRequest;
};
export declare const QueryGetGroupResponse: {
    encode(message: QueryGetGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGroupResponse;
    fromJSON(object: any): QueryGetGroupResponse;
    toJSON(message: QueryGetGroupResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetGroupResponse>): QueryGetGroupResponse;
};
export declare const QueryAllGroupRequest: {
    encode(message: QueryAllGroupRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGroupRequest;
    fromJSON(object: any): QueryAllGroupRequest;
    toJSON(message: QueryAllGroupRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllGroupRequest>): QueryAllGroupRequest;
};
export declare const QueryAllGroupResponse: {
    encode(message: QueryAllGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGroupResponse;
    fromJSON(object: any): QueryAllGroupResponse;
    toJSON(message: QueryAllGroupResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllGroupResponse>): QueryAllGroupResponse;
};
export declare const QueryGetAdminRequest: {
    encode(message: QueryGetAdminRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAdminRequest;
    fromJSON(object: any): QueryGetAdminRequest;
    toJSON(message: QueryGetAdminRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAdminRequest>): QueryGetAdminRequest;
};
export declare const QueryGetAdminResponse: {
    encode(message: QueryGetAdminResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAdminResponse;
    fromJSON(object: any): QueryGetAdminResponse;
    toJSON(message: QueryGetAdminResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAdminResponse>): QueryGetAdminResponse;
};
export declare const QueryAllAdminRequest: {
    encode(message: QueryAllAdminRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAdminRequest;
    fromJSON(object: any): QueryAllAdminRequest;
    toJSON(message: QueryAllAdminRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAdminRequest>): QueryAllAdminRequest;
};
export declare const QueryAllAdminResponse: {
    encode(message: QueryAllAdminResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAdminResponse;
    fromJSON(object: any): QueryAllAdminResponse;
    toJSON(message: QueryAllAdminResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAdminResponse>): QueryAllAdminResponse;
};
export declare const QueryListAdminOfGroupRequest: {
    encode(message: QueryListAdminOfGroupRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryListAdminOfGroupRequest;
    fromJSON(object: any): QueryListAdminOfGroupRequest;
    toJSON(message: QueryListAdminOfGroupRequest): unknown;
    fromPartial(object: DeepPartial<QueryListAdminOfGroupRequest>): QueryListAdminOfGroupRequest;
};
export declare const QueryListAdminOfGroupResponse: {
    encode(message: QueryListAdminOfGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryListAdminOfGroupResponse;
    fromJSON(object: any): QueryListAdminOfGroupResponse;
    toJSON(message: QueryListAdminOfGroupResponse): unknown;
    fromPartial(object: DeepPartial<QueryListAdminOfGroupResponse>): QueryListAdminOfGroupResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Group by index. */
    Group(request: QueryGetGroupRequest): Promise<QueryGetGroupResponse>;
    /** Queries a list of Group items. */
    GroupAll(request: QueryAllGroupRequest): Promise<QueryAllGroupResponse>;
    /** Queries a Admin by index. */
    Admin(request: QueryGetAdminRequest): Promise<QueryGetAdminResponse>;
    /** Queries a list of Admin items. */
    AdminAll(request: QueryAllAdminRequest): Promise<QueryAllAdminResponse>;
    /** Queries a list of ListAdminOfGroup items. */
    ListAdminOfGroup(request: QueryListAdminOfGroupRequest): Promise<QueryListAdminOfGroupResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Group(request: QueryGetGroupRequest): Promise<QueryGetGroupResponse>;
    GroupAll(request: QueryAllGroupRequest): Promise<QueryAllGroupResponse>;
    Admin(request: QueryGetAdminRequest): Promise<QueryGetAdminResponse>;
    AdminAll(request: QueryAllAdminRequest): Promise<QueryAllAdminResponse>;
    ListAdminOfGroup(request: QueryListAdminOfGroupRequest): Promise<QueryListAdminOfGroupResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
