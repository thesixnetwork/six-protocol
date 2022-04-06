import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
export interface Group {
    name: string;
    owner: string;
}
export declare const Group: {
    encode(message: Group, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Group;
    fromJSON(object: any): Group;
    toJSON(message: Group): unknown;
    fromPartial(object: DeepPartial<Group>): Group;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
