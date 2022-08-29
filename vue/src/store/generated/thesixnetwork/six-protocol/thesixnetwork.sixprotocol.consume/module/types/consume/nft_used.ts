/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.consume";

export interface NftUsed {
  token: string;
  amount: number;
  updateAt: string;
}

const baseNftUsed: object = { token: "", amount: 0, updateAt: "" };

export const NftUsed = {
  encode(message: NftUsed, writer: Writer = Writer.create()): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.amount !== 0) {
      writer.uint32(16).uint64(message.amount);
    }
    if (message.updateAt !== "") {
      writer.uint32(26).string(message.updateAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftUsed {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftUsed } as NftUsed;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.updateAt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftUsed {
    const message = { ...baseNftUsed } as NftUsed;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.updateAt !== undefined && object.updateAt !== null) {
      message.updateAt = String(object.updateAt);
    } else {
      message.updateAt = "";
    }
    return message;
  },

  toJSON(message: NftUsed): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.amount !== undefined && (obj.amount = message.amount);
    message.updateAt !== undefined && (obj.updateAt = message.updateAt);
    return obj;
  },

  fromPartial(object: DeepPartial<NftUsed>): NftUsed {
    const message = { ...baseNftUsed } as NftUsed;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.updateAt !== undefined && object.updateAt !== null) {
      message.updateAt = object.updateAt;
    } else {
      message.updateAt = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
