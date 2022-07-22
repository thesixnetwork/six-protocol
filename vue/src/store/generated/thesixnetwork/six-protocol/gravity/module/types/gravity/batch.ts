/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { ERC20Token } from "../gravity/attestation";

export const protobufPackage = "gravity";

/** OutgoingTxBatch represents a batch of transactions going from gravity to ETH */
export interface OutgoingTxBatch {
  batch_nonce: number;
  batch_timeout: number;
  transactions: OutgoingTransferTx[];
  token_contract: string;
  block: number;
}

/** OutgoingTransferTx represents an individual send from gravity to ETH */
export interface OutgoingTransferTx {
  id: number;
  sender: string;
  dest_address: string;
  erc20_token: ERC20Token | undefined;
  erc20_fee: ERC20Token | undefined;
}

/** OutgoingLogicCall represents an individual logic call from gravity to ETH */
export interface OutgoingLogicCall {
  transfers: ERC20Token[];
  fees: ERC20Token[];
  logic_contract_address: string;
  payload: Uint8Array;
  timeout: number;
  invalidation_id: Uint8Array;
  invalidation_nonce: number;
  block: number;
}

export interface EventOutgoingBatchCanceled {
  bridge_contract: string;
  bridge_chain_id: string;
  batch_id: string;
  nonce: string;
}

export interface EventOutgoingBatch {
  bridge_contract: string;
  bridge_chain_id: string;
  batch_id: string;
  nonce: string;
}

const baseOutgoingTxBatch: object = {
  batch_nonce: 0,
  batch_timeout: 0,
  token_contract: "",
  block: 0,
};

export const OutgoingTxBatch = {
  encode(message: OutgoingTxBatch, writer: Writer = Writer.create()): Writer {
    if (message.batch_nonce !== 0) {
      writer.uint32(8).uint64(message.batch_nonce);
    }
    if (message.batch_timeout !== 0) {
      writer.uint32(16).uint64(message.batch_timeout);
    }
    for (const v of message.transactions) {
      OutgoingTransferTx.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.token_contract !== "") {
      writer.uint32(34).string(message.token_contract);
    }
    if (message.block !== 0) {
      writer.uint32(40).uint64(message.block);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OutgoingTxBatch {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOutgoingTxBatch } as OutgoingTxBatch;
    message.transactions = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batch_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.batch_timeout = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.transactions.push(
            OutgoingTransferTx.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.token_contract = reader.string();
          break;
        case 5:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutgoingTxBatch {
    const message = { ...baseOutgoingTxBatch } as OutgoingTxBatch;
    message.transactions = [];
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = Number(object.batch_nonce);
    } else {
      message.batch_nonce = 0;
    }
    if (object.batch_timeout !== undefined && object.batch_timeout !== null) {
      message.batch_timeout = Number(object.batch_timeout);
    } else {
      message.batch_timeout = 0;
    }
    if (object.transactions !== undefined && object.transactions !== null) {
      for (const e of object.transactions) {
        message.transactions.push(OutgoingTransferTx.fromJSON(e));
      }
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = String(object.token_contract);
    } else {
      message.token_contract = "";
    }
    if (object.block !== undefined && object.block !== null) {
      message.block = Number(object.block);
    } else {
      message.block = 0;
    }
    return message;
  },

  toJSON(message: OutgoingTxBatch): unknown {
    const obj: any = {};
    message.batch_nonce !== undefined &&
      (obj.batch_nonce = message.batch_nonce);
    message.batch_timeout !== undefined &&
      (obj.batch_timeout = message.batch_timeout);
    if (message.transactions) {
      obj.transactions = message.transactions.map((e) =>
        e ? OutgoingTransferTx.toJSON(e) : undefined
      );
    } else {
      obj.transactions = [];
    }
    message.token_contract !== undefined &&
      (obj.token_contract = message.token_contract);
    message.block !== undefined && (obj.block = message.block);
    return obj;
  },

  fromPartial(object: DeepPartial<OutgoingTxBatch>): OutgoingTxBatch {
    const message = { ...baseOutgoingTxBatch } as OutgoingTxBatch;
    message.transactions = [];
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = object.batch_nonce;
    } else {
      message.batch_nonce = 0;
    }
    if (object.batch_timeout !== undefined && object.batch_timeout !== null) {
      message.batch_timeout = object.batch_timeout;
    } else {
      message.batch_timeout = 0;
    }
    if (object.transactions !== undefined && object.transactions !== null) {
      for (const e of object.transactions) {
        message.transactions.push(OutgoingTransferTx.fromPartial(e));
      }
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = object.token_contract;
    } else {
      message.token_contract = "";
    }
    if (object.block !== undefined && object.block !== null) {
      message.block = object.block;
    } else {
      message.block = 0;
    }
    return message;
  },
};

const baseOutgoingTransferTx: object = { id: 0, sender: "", dest_address: "" };

export const OutgoingTransferTx = {
  encode(
    message: OutgoingTransferTx,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    if (message.dest_address !== "") {
      writer.uint32(26).string(message.dest_address);
    }
    if (message.erc20_token !== undefined) {
      ERC20Token.encode(message.erc20_token, writer.uint32(34).fork()).ldelim();
    }
    if (message.erc20_fee !== undefined) {
      ERC20Token.encode(message.erc20_fee, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OutgoingTransferTx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOutgoingTransferTx } as OutgoingTransferTx;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.sender = reader.string();
          break;
        case 3:
          message.dest_address = reader.string();
          break;
        case 4:
          message.erc20_token = ERC20Token.decode(reader, reader.uint32());
          break;
        case 5:
          message.erc20_fee = ERC20Token.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutgoingTransferTx {
    const message = { ...baseOutgoingTransferTx } as OutgoingTransferTx;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.dest_address !== undefined && object.dest_address !== null) {
      message.dest_address = String(object.dest_address);
    } else {
      message.dest_address = "";
    }
    if (object.erc20_token !== undefined && object.erc20_token !== null) {
      message.erc20_token = ERC20Token.fromJSON(object.erc20_token);
    } else {
      message.erc20_token = undefined;
    }
    if (object.erc20_fee !== undefined && object.erc20_fee !== null) {
      message.erc20_fee = ERC20Token.fromJSON(object.erc20_fee);
    } else {
      message.erc20_fee = undefined;
    }
    return message;
  },

  toJSON(message: OutgoingTransferTx): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sender !== undefined && (obj.sender = message.sender);
    message.dest_address !== undefined &&
      (obj.dest_address = message.dest_address);
    message.erc20_token !== undefined &&
      (obj.erc20_token = message.erc20_token
        ? ERC20Token.toJSON(message.erc20_token)
        : undefined);
    message.erc20_fee !== undefined &&
      (obj.erc20_fee = message.erc20_fee
        ? ERC20Token.toJSON(message.erc20_fee)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<OutgoingTransferTx>): OutgoingTransferTx {
    const message = { ...baseOutgoingTransferTx } as OutgoingTransferTx;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.dest_address !== undefined && object.dest_address !== null) {
      message.dest_address = object.dest_address;
    } else {
      message.dest_address = "";
    }
    if (object.erc20_token !== undefined && object.erc20_token !== null) {
      message.erc20_token = ERC20Token.fromPartial(object.erc20_token);
    } else {
      message.erc20_token = undefined;
    }
    if (object.erc20_fee !== undefined && object.erc20_fee !== null) {
      message.erc20_fee = ERC20Token.fromPartial(object.erc20_fee);
    } else {
      message.erc20_fee = undefined;
    }
    return message;
  },
};

const baseOutgoingLogicCall: object = {
  logic_contract_address: "",
  timeout: 0,
  invalidation_nonce: 0,
  block: 0,
};

export const OutgoingLogicCall = {
  encode(message: OutgoingLogicCall, writer: Writer = Writer.create()): Writer {
    for (const v of message.transfers) {
      ERC20Token.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.fees) {
      ERC20Token.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.logic_contract_address !== "") {
      writer.uint32(26).string(message.logic_contract_address);
    }
    if (message.payload.length !== 0) {
      writer.uint32(34).bytes(message.payload);
    }
    if (message.timeout !== 0) {
      writer.uint32(40).uint64(message.timeout);
    }
    if (message.invalidation_id.length !== 0) {
      writer.uint32(50).bytes(message.invalidation_id);
    }
    if (message.invalidation_nonce !== 0) {
      writer.uint32(56).uint64(message.invalidation_nonce);
    }
    if (message.block !== 0) {
      writer.uint32(64).uint64(message.block);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OutgoingLogicCall {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOutgoingLogicCall } as OutgoingLogicCall;
    message.transfers = [];
    message.fees = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transfers.push(ERC20Token.decode(reader, reader.uint32()));
          break;
        case 2:
          message.fees.push(ERC20Token.decode(reader, reader.uint32()));
          break;
        case 3:
          message.logic_contract_address = reader.string();
          break;
        case 4:
          message.payload = reader.bytes();
          break;
        case 5:
          message.timeout = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.invalidation_id = reader.bytes();
          break;
        case 7:
          message.invalidation_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutgoingLogicCall {
    const message = { ...baseOutgoingLogicCall } as OutgoingLogicCall;
    message.transfers = [];
    message.fees = [];
    if (object.transfers !== undefined && object.transfers !== null) {
      for (const e of object.transfers) {
        message.transfers.push(ERC20Token.fromJSON(e));
      }
    }
    if (object.fees !== undefined && object.fees !== null) {
      for (const e of object.fees) {
        message.fees.push(ERC20Token.fromJSON(e));
      }
    }
    if (
      object.logic_contract_address !== undefined &&
      object.logic_contract_address !== null
    ) {
      message.logic_contract_address = String(object.logic_contract_address);
    } else {
      message.logic_contract_address = "";
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = bytesFromBase64(object.payload);
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = Number(object.timeout);
    } else {
      message.timeout = 0;
    }
    if (
      object.invalidation_id !== undefined &&
      object.invalidation_id !== null
    ) {
      message.invalidation_id = bytesFromBase64(object.invalidation_id);
    }
    if (
      object.invalidation_nonce !== undefined &&
      object.invalidation_nonce !== null
    ) {
      message.invalidation_nonce = Number(object.invalidation_nonce);
    } else {
      message.invalidation_nonce = 0;
    }
    if (object.block !== undefined && object.block !== null) {
      message.block = Number(object.block);
    } else {
      message.block = 0;
    }
    return message;
  },

  toJSON(message: OutgoingLogicCall): unknown {
    const obj: any = {};
    if (message.transfers) {
      obj.transfers = message.transfers.map((e) =>
        e ? ERC20Token.toJSON(e) : undefined
      );
    } else {
      obj.transfers = [];
    }
    if (message.fees) {
      obj.fees = message.fees.map((e) =>
        e ? ERC20Token.toJSON(e) : undefined
      );
    } else {
      obj.fees = [];
    }
    message.logic_contract_address !== undefined &&
      (obj.logic_contract_address = message.logic_contract_address);
    message.payload !== undefined &&
      (obj.payload = base64FromBytes(
        message.payload !== undefined ? message.payload : new Uint8Array()
      ));
    message.timeout !== undefined && (obj.timeout = message.timeout);
    message.invalidation_id !== undefined &&
      (obj.invalidation_id = base64FromBytes(
        message.invalidation_id !== undefined
          ? message.invalidation_id
          : new Uint8Array()
      ));
    message.invalidation_nonce !== undefined &&
      (obj.invalidation_nonce = message.invalidation_nonce);
    message.block !== undefined && (obj.block = message.block);
    return obj;
  },

  fromPartial(object: DeepPartial<OutgoingLogicCall>): OutgoingLogicCall {
    const message = { ...baseOutgoingLogicCall } as OutgoingLogicCall;
    message.transfers = [];
    message.fees = [];
    if (object.transfers !== undefined && object.transfers !== null) {
      for (const e of object.transfers) {
        message.transfers.push(ERC20Token.fromPartial(e));
      }
    }
    if (object.fees !== undefined && object.fees !== null) {
      for (const e of object.fees) {
        message.fees.push(ERC20Token.fromPartial(e));
      }
    }
    if (
      object.logic_contract_address !== undefined &&
      object.logic_contract_address !== null
    ) {
      message.logic_contract_address = object.logic_contract_address;
    } else {
      message.logic_contract_address = "";
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = object.payload;
    } else {
      message.payload = new Uint8Array();
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = object.timeout;
    } else {
      message.timeout = 0;
    }
    if (
      object.invalidation_id !== undefined &&
      object.invalidation_id !== null
    ) {
      message.invalidation_id = object.invalidation_id;
    } else {
      message.invalidation_id = new Uint8Array();
    }
    if (
      object.invalidation_nonce !== undefined &&
      object.invalidation_nonce !== null
    ) {
      message.invalidation_nonce = object.invalidation_nonce;
    } else {
      message.invalidation_nonce = 0;
    }
    if (object.block !== undefined && object.block !== null) {
      message.block = object.block;
    } else {
      message.block = 0;
    }
    return message;
  },
};

const baseEventOutgoingBatchCanceled: object = {
  bridge_contract: "",
  bridge_chain_id: "",
  batch_id: "",
  nonce: "",
};

export const EventOutgoingBatchCanceled = {
  encode(
    message: EventOutgoingBatchCanceled,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bridge_contract !== "") {
      writer.uint32(10).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(18).string(message.bridge_chain_id);
    }
    if (message.batch_id !== "") {
      writer.uint32(26).string(message.batch_id);
    }
    if (message.nonce !== "") {
      writer.uint32(34).string(message.nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventOutgoingBatchCanceled {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventOutgoingBatchCanceled,
    } as EventOutgoingBatchCanceled;
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
          message.batch_id = reader.string();
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

  fromJSON(object: any): EventOutgoingBatchCanceled {
    const message = {
      ...baseEventOutgoingBatchCanceled,
    } as EventOutgoingBatchCanceled;
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
    if (object.batch_id !== undefined && object.batch_id !== null) {
      message.batch_id = String(object.batch_id);
    } else {
      message.batch_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventOutgoingBatchCanceled): unknown {
    const obj: any = {};
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.batch_id !== undefined && (obj.batch_id = message.batch_id);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventOutgoingBatchCanceled>
  ): EventOutgoingBatchCanceled {
    const message = {
      ...baseEventOutgoingBatchCanceled,
    } as EventOutgoingBatchCanceled;
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
    if (object.batch_id !== undefined && object.batch_id !== null) {
      message.batch_id = object.batch_id;
    } else {
      message.batch_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventOutgoingBatch: object = {
  bridge_contract: "",
  bridge_chain_id: "",
  batch_id: "",
  nonce: "",
};

export const EventOutgoingBatch = {
  encode(
    message: EventOutgoingBatch,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bridge_contract !== "") {
      writer.uint32(10).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(18).string(message.bridge_chain_id);
    }
    if (message.batch_id !== "") {
      writer.uint32(26).string(message.batch_id);
    }
    if (message.nonce !== "") {
      writer.uint32(34).string(message.nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventOutgoingBatch {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventOutgoingBatch } as EventOutgoingBatch;
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
          message.batch_id = reader.string();
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

  fromJSON(object: any): EventOutgoingBatch {
    const message = { ...baseEventOutgoingBatch } as EventOutgoingBatch;
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
    if (object.batch_id !== undefined && object.batch_id !== null) {
      message.batch_id = String(object.batch_id);
    } else {
      message.batch_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventOutgoingBatch): unknown {
    const obj: any = {};
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.batch_id !== undefined && (obj.batch_id = message.batch_id);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(object: DeepPartial<EventOutgoingBatch>): EventOutgoingBatch {
    const message = { ...baseEventOutgoingBatch } as EventOutgoingBatch;
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
    if (object.batch_id !== undefined && object.batch_id !== null) {
      message.batch_id = object.batch_id;
    } else {
      message.batch_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

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
