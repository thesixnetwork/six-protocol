/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "gravity.v1";

/** IDSet represents a set of IDs */
export interface IDSet {
  ids: number[];
}

export interface BatchFees {
  token: string;
  total_fees: string;
  tx_count: number;
}

export interface EventWithdrawalReceived {
  bridge_contract: string;
  bridge_chain_id: string;
  outgoing_tx_id: string;
  nonce: string;
}

export interface EventWithdrawCanceled {
  sender: string;
  tx_id: string;
  bridge_contract: string;
  bridge_chain_id: string;
}

const baseIDSet: object = { ids: 0 };

export const IDSet = {
  encode(message: IDSet, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IDSet {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIDSet } as IDSet;
    message.ids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ids.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IDSet {
    const message = { ...baseIDSet } as IDSet;
    message.ids = [];
    if (object.ids !== undefined && object.ids !== null) {
      for (const e of object.ids) {
        message.ids.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: IDSet): unknown {
    const obj: any = {};
    if (message.ids) {
      obj.ids = message.ids.map((e) => e);
    } else {
      obj.ids = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<IDSet>): IDSet {
    const message = { ...baseIDSet } as IDSet;
    message.ids = [];
    if (object.ids !== undefined && object.ids !== null) {
      for (const e of object.ids) {
        message.ids.push(e);
      }
    }
    return message;
  },
};

const baseBatchFees: object = { token: "", total_fees: "", tx_count: 0 };

export const BatchFees = {
  encode(message: BatchFees, writer: Writer = Writer.create()): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.total_fees !== "") {
      writer.uint32(18).string(message.total_fees);
    }
    if (message.tx_count !== 0) {
      writer.uint32(24).uint64(message.tx_count);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BatchFees {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBatchFees } as BatchFees;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.total_fees = reader.string();
          break;
        case 3:
          message.tx_count = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BatchFees {
    const message = { ...baseBatchFees } as BatchFees;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.total_fees !== undefined && object.total_fees !== null) {
      message.total_fees = String(object.total_fees);
    } else {
      message.total_fees = "";
    }
    if (object.tx_count !== undefined && object.tx_count !== null) {
      message.tx_count = Number(object.tx_count);
    } else {
      message.tx_count = 0;
    }
    return message;
  },

  toJSON(message: BatchFees): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.total_fees !== undefined && (obj.total_fees = message.total_fees);
    message.tx_count !== undefined && (obj.tx_count = message.tx_count);
    return obj;
  },

  fromPartial(object: DeepPartial<BatchFees>): BatchFees {
    const message = { ...baseBatchFees } as BatchFees;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.total_fees !== undefined && object.total_fees !== null) {
      message.total_fees = object.total_fees;
    } else {
      message.total_fees = "";
    }
    if (object.tx_count !== undefined && object.tx_count !== null) {
      message.tx_count = object.tx_count;
    } else {
      message.tx_count = 0;
    }
    return message;
  },
};

const baseEventWithdrawalReceived: object = {
  bridge_contract: "",
  bridge_chain_id: "",
  outgoing_tx_id: "",
  nonce: "",
};

export const EventWithdrawalReceived = {
  encode(
    message: EventWithdrawalReceived,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bridge_contract !== "") {
      writer.uint32(10).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(18).string(message.bridge_chain_id);
    }
    if (message.outgoing_tx_id !== "") {
      writer.uint32(26).string(message.outgoing_tx_id);
    }
    if (message.nonce !== "") {
      writer.uint32(34).string(message.nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventWithdrawalReceived {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventWithdrawalReceived,
    } as EventWithdrawalReceived;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bridge_contract = reader.string();
          break;
        case 2:
          message.bridge_chain_id = reader.string();
          break;
        case 3:
          message.outgoing_tx_id = reader.string();
          break;
        case 4:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventWithdrawalReceived {
    const message = {
      ...baseEventWithdrawalReceived,
    } as EventWithdrawalReceived;
    if (
      object.bridge_contract !== undefined &&
      object.bridge_contract !== null
    ) {
      message.bridge_contract = String(object.bridge_contract);
    } else {
      message.bridge_contract = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = String(object.bridge_chain_id);
    } else {
      message.bridge_chain_id = "";
    }
    if (object.outgoing_tx_id !== undefined && object.outgoing_tx_id !== null) {
      message.outgoing_tx_id = String(object.outgoing_tx_id);
    } else {
      message.outgoing_tx_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventWithdrawalReceived): unknown {
    const obj: any = {};
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.outgoing_tx_id !== undefined &&
      (obj.outgoing_tx_id = message.outgoing_tx_id);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventWithdrawalReceived>
  ): EventWithdrawalReceived {
    const message = {
      ...baseEventWithdrawalReceived,
    } as EventWithdrawalReceived;
    if (
      object.bridge_contract !== undefined &&
      object.bridge_contract !== null
    ) {
      message.bridge_contract = object.bridge_contract;
    } else {
      message.bridge_contract = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = object.bridge_chain_id;
    } else {
      message.bridge_chain_id = "";
    }
    if (object.outgoing_tx_id !== undefined && object.outgoing_tx_id !== null) {
      message.outgoing_tx_id = object.outgoing_tx_id;
    } else {
      message.outgoing_tx_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventWithdrawCanceled: object = {
  sender: "",
  tx_id: "",
  bridge_contract: "",
  bridge_chain_id: "",
};

export const EventWithdrawCanceled = {
  encode(
    message: EventWithdrawCanceled,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.tx_id !== "") {
      writer.uint32(18).string(message.tx_id);
    }
    if (message.bridge_contract !== "") {
      writer.uint32(26).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(34).string(message.bridge_chain_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventWithdrawCanceled {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventWithdrawCanceled } as EventWithdrawCanceled;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.tx_id = reader.string();
          break;
        case 3:
          message.bridge_contract = reader.string();
          break;
        case 4:
          message.bridge_chain_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventWithdrawCanceled {
    const message = { ...baseEventWithdrawCanceled } as EventWithdrawCanceled;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.tx_id !== undefined && object.tx_id !== null) {
      message.tx_id = String(object.tx_id);
    } else {
      message.tx_id = "";
    }
    if (
      object.bridge_contract !== undefined &&
      object.bridge_contract !== null
    ) {
      message.bridge_contract = String(object.bridge_contract);
    } else {
      message.bridge_contract = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = String(object.bridge_chain_id);
    } else {
      message.bridge_chain_id = "";
    }
    return message;
  },

  toJSON(message: EventWithdrawCanceled): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.tx_id !== undefined && (obj.tx_id = message.tx_id);
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventWithdrawCanceled>
  ): EventWithdrawCanceled {
    const message = { ...baseEventWithdrawCanceled } as EventWithdrawCanceled;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.tx_id !== undefined && object.tx_id !== null) {
      message.tx_id = object.tx_id;
    } else {
      message.tx_id = "";
    }
    if (
      object.bridge_contract !== undefined &&
      object.bridge_contract !== null
    ) {
      message.bridge_contract = object.bridge_contract;
    } else {
      message.bridge_contract = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = object.bridge_chain_id;
    } else {
      message.bridge_chain_id = "";
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
