import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
export interface Options {
    defaultMintee: string;
}
export declare const Options: {
    encode(message: Options, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Options;
    fromJSON(object: any): Options;
    toJSON(message: Options): unknown;
    fromPartial(object: DeepPartial<Options>): Options;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
