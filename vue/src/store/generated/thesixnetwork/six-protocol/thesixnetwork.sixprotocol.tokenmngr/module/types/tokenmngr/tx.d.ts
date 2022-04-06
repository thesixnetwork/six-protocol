import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
export interface MsgCreateToken {
    creator: string;
    name: string;
    maxSupply: number;
    denomMetaData: string;
    mintee: string;
}
export interface MsgCreateTokenResponse {
}
export interface MsgUpdateToken {
    creator: string;
    name: string;
    maxSupply: number;
    mintee: string;
}
export interface MsgUpdateTokenResponse {
}
export interface MsgDeleteToken {
    creator: string;
    name: string;
}
export interface MsgDeleteTokenResponse {
}
export interface MsgCreateMintperm {
    creator: string;
    token: string;
    address: string;
}
export interface MsgCreateMintpermResponse {
}
export interface MsgUpdateMintperm {
    creator: string;
    token: string;
    address: string;
}
export interface MsgUpdateMintpermResponse {
}
export interface MsgDeleteMintperm {
    creator: string;
    token: string;
    address: string;
}
export interface MsgDeleteMintpermResponse {
}
export interface MsgMint {
    creator: string;
    amount: number;
    token: string;
}
export interface MsgMintResponse {
}
export interface MsgCreateOptions {
    creator: string;
    defaultMintee: string;
}
export interface MsgCreateOptionsResponse {
}
export interface MsgUpdateOptions {
    creator: string;
    defaultMintee: string;
}
export interface MsgUpdateOptionsResponse {
}
export interface MsgDeleteOptions {
    creator: string;
}
export interface MsgDeleteOptionsResponse {
}
export declare const MsgCreateToken: {
    encode(message: MsgCreateToken, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateToken;
    fromJSON(object: any): MsgCreateToken;
    toJSON(message: MsgCreateToken): unknown;
    fromPartial(object: DeepPartial<MsgCreateToken>): MsgCreateToken;
};
export declare const MsgCreateTokenResponse: {
    encode(_: MsgCreateTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTokenResponse;
    fromJSON(_: any): MsgCreateTokenResponse;
    toJSON(_: MsgCreateTokenResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateTokenResponse>): MsgCreateTokenResponse;
};
export declare const MsgUpdateToken: {
    encode(message: MsgUpdateToken, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateToken;
    fromJSON(object: any): MsgUpdateToken;
    toJSON(message: MsgUpdateToken): unknown;
    fromPartial(object: DeepPartial<MsgUpdateToken>): MsgUpdateToken;
};
export declare const MsgUpdateTokenResponse: {
    encode(_: MsgUpdateTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTokenResponse;
    fromJSON(_: any): MsgUpdateTokenResponse;
    toJSON(_: MsgUpdateTokenResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTokenResponse>): MsgUpdateTokenResponse;
};
export declare const MsgDeleteToken: {
    encode(message: MsgDeleteToken, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteToken;
    fromJSON(object: any): MsgDeleteToken;
    toJSON(message: MsgDeleteToken): unknown;
    fromPartial(object: DeepPartial<MsgDeleteToken>): MsgDeleteToken;
};
export declare const MsgDeleteTokenResponse: {
    encode(_: MsgDeleteTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTokenResponse;
    fromJSON(_: any): MsgDeleteTokenResponse;
    toJSON(_: MsgDeleteTokenResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTokenResponse>): MsgDeleteTokenResponse;
};
export declare const MsgCreateMintperm: {
    encode(message: MsgCreateMintperm, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMintperm;
    fromJSON(object: any): MsgCreateMintperm;
    toJSON(message: MsgCreateMintperm): unknown;
    fromPartial(object: DeepPartial<MsgCreateMintperm>): MsgCreateMintperm;
};
export declare const MsgCreateMintpermResponse: {
    encode(_: MsgCreateMintpermResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMintpermResponse;
    fromJSON(_: any): MsgCreateMintpermResponse;
    toJSON(_: MsgCreateMintpermResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateMintpermResponse>): MsgCreateMintpermResponse;
};
export declare const MsgUpdateMintperm: {
    encode(message: MsgUpdateMintperm, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMintperm;
    fromJSON(object: any): MsgUpdateMintperm;
    toJSON(message: MsgUpdateMintperm): unknown;
    fromPartial(object: DeepPartial<MsgUpdateMintperm>): MsgUpdateMintperm;
};
export declare const MsgUpdateMintpermResponse: {
    encode(_: MsgUpdateMintpermResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMintpermResponse;
    fromJSON(_: any): MsgUpdateMintpermResponse;
    toJSON(_: MsgUpdateMintpermResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateMintpermResponse>): MsgUpdateMintpermResponse;
};
export declare const MsgDeleteMintperm: {
    encode(message: MsgDeleteMintperm, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMintperm;
    fromJSON(object: any): MsgDeleteMintperm;
    toJSON(message: MsgDeleteMintperm): unknown;
    fromPartial(object: DeepPartial<MsgDeleteMintperm>): MsgDeleteMintperm;
};
export declare const MsgDeleteMintpermResponse: {
    encode(_: MsgDeleteMintpermResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMintpermResponse;
    fromJSON(_: any): MsgDeleteMintpermResponse;
    toJSON(_: MsgDeleteMintpermResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteMintpermResponse>): MsgDeleteMintpermResponse;
};
export declare const MsgMint: {
    encode(message: MsgMint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMint;
    fromJSON(object: any): MsgMint;
    toJSON(message: MsgMint): unknown;
    fromPartial(object: DeepPartial<MsgMint>): MsgMint;
};
export declare const MsgMintResponse: {
    encode(_: MsgMintResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMintResponse;
    fromJSON(_: any): MsgMintResponse;
    toJSON(_: MsgMintResponse): unknown;
    fromPartial(_: DeepPartial<MsgMintResponse>): MsgMintResponse;
};
export declare const MsgCreateOptions: {
    encode(message: MsgCreateOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateOptions;
    fromJSON(object: any): MsgCreateOptions;
    toJSON(message: MsgCreateOptions): unknown;
    fromPartial(object: DeepPartial<MsgCreateOptions>): MsgCreateOptions;
};
export declare const MsgCreateOptionsResponse: {
    encode(_: MsgCreateOptionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateOptionsResponse;
    fromJSON(_: any): MsgCreateOptionsResponse;
    toJSON(_: MsgCreateOptionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateOptionsResponse>): MsgCreateOptionsResponse;
};
export declare const MsgUpdateOptions: {
    encode(message: MsgUpdateOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateOptions;
    fromJSON(object: any): MsgUpdateOptions;
    toJSON(message: MsgUpdateOptions): unknown;
    fromPartial(object: DeepPartial<MsgUpdateOptions>): MsgUpdateOptions;
};
export declare const MsgUpdateOptionsResponse: {
    encode(_: MsgUpdateOptionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateOptionsResponse;
    fromJSON(_: any): MsgUpdateOptionsResponse;
    toJSON(_: MsgUpdateOptionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateOptionsResponse>): MsgUpdateOptionsResponse;
};
export declare const MsgDeleteOptions: {
    encode(message: MsgDeleteOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteOptions;
    fromJSON(object: any): MsgDeleteOptions;
    toJSON(message: MsgDeleteOptions): unknown;
    fromPartial(object: DeepPartial<MsgDeleteOptions>): MsgDeleteOptions;
};
export declare const MsgDeleteOptionsResponse: {
    encode(_: MsgDeleteOptionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteOptionsResponse;
    fromJSON(_: any): MsgDeleteOptionsResponse;
    toJSON(_: MsgDeleteOptionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteOptionsResponse>): MsgDeleteOptionsResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateToken(request: MsgCreateToken): Promise<MsgCreateTokenResponse>;
    UpdateToken(request: MsgUpdateToken): Promise<MsgUpdateTokenResponse>;
    DeleteToken(request: MsgDeleteToken): Promise<MsgDeleteTokenResponse>;
    CreateMintperm(request: MsgCreateMintperm): Promise<MsgCreateMintpermResponse>;
    UpdateMintperm(request: MsgUpdateMintperm): Promise<MsgUpdateMintpermResponse>;
    DeleteMintperm(request: MsgDeleteMintperm): Promise<MsgDeleteMintpermResponse>;
    Mint(request: MsgMint): Promise<MsgMintResponse>;
    CreateOptions(request: MsgCreateOptions): Promise<MsgCreateOptionsResponse>;
    UpdateOptions(request: MsgUpdateOptions): Promise<MsgUpdateOptionsResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteOptions(request: MsgDeleteOptions): Promise<MsgDeleteOptionsResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateToken(request: MsgCreateToken): Promise<MsgCreateTokenResponse>;
    UpdateToken(request: MsgUpdateToken): Promise<MsgUpdateTokenResponse>;
    DeleteToken(request: MsgDeleteToken): Promise<MsgDeleteTokenResponse>;
    CreateMintperm(request: MsgCreateMintperm): Promise<MsgCreateMintpermResponse>;
    UpdateMintperm(request: MsgUpdateMintperm): Promise<MsgUpdateMintpermResponse>;
    DeleteMintperm(request: MsgDeleteMintperm): Promise<MsgDeleteMintpermResponse>;
    Mint(request: MsgMint): Promise<MsgMintResponse>;
    CreateOptions(request: MsgCreateOptions): Promise<MsgCreateOptionsResponse>;
    UpdateOptions(request: MsgUpdateOptions): Promise<MsgUpdateOptionsResponse>;
    DeleteOptions(request: MsgDeleteOptions): Promise<MsgDeleteOptionsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
