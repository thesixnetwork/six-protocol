import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
export interface ProtocoladminPacketData {
    /** this line is used by starport scaffolding # ibc/packet/proto/field */
    noData: NoData | undefined;
}
export interface NoData {
}
export declare const ProtocoladminPacketData: {
    encode(message: ProtocoladminPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProtocoladminPacketData;
    fromJSON(object: any): ProtocoladminPacketData;
    toJSON(message: ProtocoladminPacketData): unknown;
    fromPartial(object: DeepPartial<ProtocoladminPacketData>): ProtocoladminPacketData;
};
export declare const NoData: {
    encode(_: NoData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NoData;
    fromJSON(_: any): NoData;
    toJSON(_: NoData): unknown;
    fromPartial(_: DeepPartial<NoData>): NoData;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
