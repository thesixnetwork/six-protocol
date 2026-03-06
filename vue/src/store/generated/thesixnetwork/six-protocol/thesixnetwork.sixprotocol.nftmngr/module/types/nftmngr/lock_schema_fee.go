/* eslint-disable */
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.nftmngr";

export interface LockSchemaFee {
  id: string;
  virtualSchemaCode: string;
  amount: Coin | undefined;
  proposer: string;
}

const baseLockSchemaFee: object = {
  id: "",
  virtualSchemaCode: "",
  proposer: "",
};

export const LockSchemaFee = {
  encode(message: LockSchemaFee, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.virtualSchemaCode !== "") {
      writer.uint32(18).string(message.virtualSchemaCode);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    if (message.proposer !== "") {
      writer.uint32(34).string(message.proposer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LockSchemaFee {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLockSchemaFee } as LockSchemaFee;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.virtualSchemaCode = reader.string();
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.proposer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LockSchemaFee {
    const message = { ...baseLockSchemaFee } as LockSchemaFee;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (
      object.virtualSchemaCode !== undefined &&
      object.virtualSchemaCode !== null
    ) {
      message.virtualSchemaCode = String(object.virtualSchemaCode);
    } else {
      message.virtualSchemaCode = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer);
    } else {
      message.proposer = "";
    }
    return message;
  },

  toJSON(message: LockSchemaFee): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.virtualSchemaCode !== undefined &&
      (obj.virtualSchemaCode = message.virtualSchemaCode);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.proposer !== undefined && (obj.proposer = message.proposer);
    return obj;
  },

  fromPartial(object: DeepPartial<LockSchemaFee>): LockSchemaFee {
    const message = { ...baseLockSchemaFee } as LockSchemaFee;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (
      object.virtualSchemaCode !== undefined &&
      object.virtualSchemaCode !== null
    ) {
      message.virtualSchemaCode = object.virtualSchemaCode;
    } else {
      message.virtualSchemaCode = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer;
    } else {
      message.proposer = "";
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
