/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.evmbind";

export interface Binding {
  ethAddress: string;
  ethSignature: string;
  signMessage: string;
  creator: string;
}

const baseBinding: object = {
  ethAddress: "",
  ethSignature: "",
  signMessage: "",
  creator: "",
};

export const Binding = {
  encode(message: Binding, writer: Writer = Writer.create()): Writer {
    if (message.ethAddress !== "") {
      writer.uint32(10).string(message.ethAddress);
    }
    if (message.ethSignature !== "") {
      writer.uint32(18).string(message.ethSignature);
    }
    if (message.signMessage !== "") {
      writer.uint32(26).string(message.signMessage);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Binding {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBinding } as Binding;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethAddress = reader.string();
          break;
        case 2:
          message.ethSignature = reader.string();
          break;
        case 3:
          message.signMessage = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Binding {
    const message = { ...baseBinding } as Binding;
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = String(object.ethAddress);
    } else {
      message.ethAddress = "";
    }
    if (object.ethSignature !== undefined && object.ethSignature !== null) {
      message.ethSignature = String(object.ethSignature);
    } else {
      message.ethSignature = "";
    }
    if (object.signMessage !== undefined && object.signMessage !== null) {
      message.signMessage = String(object.signMessage);
    } else {
      message.signMessage = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Binding): unknown {
    const obj: any = {};
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    message.ethSignature !== undefined &&
      (obj.ethSignature = message.ethSignature);
    message.signMessage !== undefined &&
      (obj.signMessage = message.signMessage);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Binding>): Binding {
    const message = { ...baseBinding } as Binding;
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = object.ethAddress;
    } else {
      message.ethAddress = "";
    }
    if (object.ethSignature !== undefined && object.ethSignature !== null) {
      message.ethSignature = object.ethSignature;
    } else {
      message.ethSignature = "";
    }
    if (object.signMessage !== undefined && object.signMessage !== null) {
      message.signMessage = object.signMessage;
    } else {
      message.signMessage = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

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
