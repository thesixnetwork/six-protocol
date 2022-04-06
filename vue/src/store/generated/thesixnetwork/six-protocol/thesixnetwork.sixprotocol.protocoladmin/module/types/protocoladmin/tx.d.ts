import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
export interface MsgCreateGroup {
    creator: string;
    name: string;
}
export interface MsgCreateGroupResponse {
}
export interface MsgUpdateGroup {
    creator: string;
    name: string;
}
export interface MsgUpdateGroupResponse {
}
export interface MsgDeleteGroup {
    creator: string;
    name: string;
}
export interface MsgDeleteGroupResponse {
}
export interface MsgAddAdminToGroup {
    creator: string;
    name: string;
    address: string;
}
export interface MsgAddAdminToGroupResponse {
}
export interface MsgRemoveAdminFromGroup {
    creator: string;
    name: string;
    address: string;
}
export interface MsgRemoveAdminFromGroupResponse {
}
export declare const MsgCreateGroup: {
    encode(message: MsgCreateGroup, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGroup;
    fromJSON(object: any): MsgCreateGroup;
    toJSON(message: MsgCreateGroup): unknown;
    fromPartial(object: DeepPartial<MsgCreateGroup>): MsgCreateGroup;
};
export declare const MsgCreateGroupResponse: {
    encode(_: MsgCreateGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGroupResponse;
    fromJSON(_: any): MsgCreateGroupResponse;
    toJSON(_: MsgCreateGroupResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateGroupResponse>): MsgCreateGroupResponse;
};
export declare const MsgUpdateGroup: {
    encode(message: MsgUpdateGroup, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateGroup;
    fromJSON(object: any): MsgUpdateGroup;
    toJSON(message: MsgUpdateGroup): unknown;
    fromPartial(object: DeepPartial<MsgUpdateGroup>): MsgUpdateGroup;
};
export declare const MsgUpdateGroupResponse: {
    encode(_: MsgUpdateGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateGroupResponse;
    fromJSON(_: any): MsgUpdateGroupResponse;
    toJSON(_: MsgUpdateGroupResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateGroupResponse>): MsgUpdateGroupResponse;
};
export declare const MsgDeleteGroup: {
    encode(message: MsgDeleteGroup, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteGroup;
    fromJSON(object: any): MsgDeleteGroup;
    toJSON(message: MsgDeleteGroup): unknown;
    fromPartial(object: DeepPartial<MsgDeleteGroup>): MsgDeleteGroup;
};
export declare const MsgDeleteGroupResponse: {
    encode(_: MsgDeleteGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteGroupResponse;
    fromJSON(_: any): MsgDeleteGroupResponse;
    toJSON(_: MsgDeleteGroupResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteGroupResponse>): MsgDeleteGroupResponse;
};
export declare const MsgAddAdminToGroup: {
    encode(message: MsgAddAdminToGroup, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddAdminToGroup;
    fromJSON(object: any): MsgAddAdminToGroup;
    toJSON(message: MsgAddAdminToGroup): unknown;
    fromPartial(object: DeepPartial<MsgAddAdminToGroup>): MsgAddAdminToGroup;
};
export declare const MsgAddAdminToGroupResponse: {
    encode(_: MsgAddAdminToGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddAdminToGroupResponse;
    fromJSON(_: any): MsgAddAdminToGroupResponse;
    toJSON(_: MsgAddAdminToGroupResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddAdminToGroupResponse>): MsgAddAdminToGroupResponse;
};
export declare const MsgRemoveAdminFromGroup: {
    encode(message: MsgRemoveAdminFromGroup, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRemoveAdminFromGroup;
    fromJSON(object: any): MsgRemoveAdminFromGroup;
    toJSON(message: MsgRemoveAdminFromGroup): unknown;
    fromPartial(object: DeepPartial<MsgRemoveAdminFromGroup>): MsgRemoveAdminFromGroup;
};
export declare const MsgRemoveAdminFromGroupResponse: {
    encode(_: MsgRemoveAdminFromGroupResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRemoveAdminFromGroupResponse;
    fromJSON(_: any): MsgRemoveAdminFromGroupResponse;
    toJSON(_: MsgRemoveAdminFromGroupResponse): unknown;
    fromPartial(_: DeepPartial<MsgRemoveAdminFromGroupResponse>): MsgRemoveAdminFromGroupResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateGroup(request: MsgCreateGroup): Promise<MsgCreateGroupResponse>;
    UpdateGroup(request: MsgUpdateGroup): Promise<MsgUpdateGroupResponse>;
    DeleteGroup(request: MsgDeleteGroup): Promise<MsgDeleteGroupResponse>;
    AddAdminToGroup(request: MsgAddAdminToGroup): Promise<MsgAddAdminToGroupResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RemoveAdminFromGroup(request: MsgRemoveAdminFromGroup): Promise<MsgRemoveAdminFromGroupResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateGroup(request: MsgCreateGroup): Promise<MsgCreateGroupResponse>;
    UpdateGroup(request: MsgUpdateGroup): Promise<MsgUpdateGroupResponse>;
    DeleteGroup(request: MsgDeleteGroup): Promise<MsgDeleteGroupResponse>;
    AddAdminToGroup(request: MsgAddAdminToGroup): Promise<MsgAddAdminToGroupResponse>;
    RemoveAdminFromGroup(request: MsgRemoveAdminFromGroup): Promise<MsgRemoveAdminFromGroupResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
