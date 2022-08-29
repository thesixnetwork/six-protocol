/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "thesixnetwork.sixprotocol.consume";

export interface MsgUseNft {
  creator: string;
  token: string;
  timestamp: string;
}

export interface MsgUseNftResponse {
  Id: number;
}

export interface MsgUseNftByEVM {
  creator: string;
  token: string;
  ethSignature: string;
  signMessage: string;
  ethAddress: string;
}

export interface MsgUseNftByEVMResponse {}

const baseMsgUseNft: object = { creator: "", token: "", timestamp: "" };

export const MsgUseNft = {
  encode(message: MsgUseNft, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    if (message.timestamp !== "") {
      writer.uint32(26).string(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUseNft {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUseNft } as MsgUseNft;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.token = reader.string();
          break;
        case 3:
          message.timestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUseNft {
    const message = { ...baseMsgUseNft } as MsgUseNft;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    return message;
  },

  toJSON(message: MsgUseNft): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.token !== undefined && (obj.token = message.token);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUseNft>): MsgUseNft {
    const message = { ...baseMsgUseNft } as MsgUseNft;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    return message;
  },
};

const baseMsgUseNftResponse: object = { Id: 0 };

export const MsgUseNftResponse = {
  encode(message: MsgUseNftResponse, writer: Writer = Writer.create()): Writer {
    if (message.Id !== 0) {
      writer.uint32(8).uint64(message.Id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUseNftResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUseNftResponse } as MsgUseNftResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUseNftResponse {
    const message = { ...baseMsgUseNftResponse } as MsgUseNftResponse;
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = Number(object.Id);
    } else {
      message.Id = 0;
    }
    return message;
  },

  toJSON(message: MsgUseNftResponse): unknown {
    const obj: any = {};
    message.Id !== undefined && (obj.Id = message.Id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUseNftResponse>): MsgUseNftResponse {
    const message = { ...baseMsgUseNftResponse } as MsgUseNftResponse;
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = object.Id;
    } else {
      message.Id = 0;
    }
    return message;
  },
};

const baseMsgUseNftByEVM: object = {
  creator: "",
  token: "",
  ethSignature: "",
  signMessage: "",
  ethAddress: "",
};

export const MsgUseNftByEVM = {
  encode(message: MsgUseNftByEVM, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    if (message.ethSignature !== "") {
      writer.uint32(26).string(message.ethSignature);
    }
    if (message.signMessage !== "") {
      writer.uint32(34).string(message.signMessage);
    }
    if (message.ethAddress !== "") {
      writer.uint32(42).string(message.ethAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUseNftByEVM {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUseNftByEVM } as MsgUseNftByEVM;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.token = reader.string();
          break;
        case 3:
          message.ethSignature = reader.string();
          break;
        case 4:
          message.signMessage = reader.string();
          break;
        case 5:
          message.ethAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUseNftByEVM {
    const message = { ...baseMsgUseNftByEVM } as MsgUseNftByEVM;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
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
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = String(object.ethAddress);
    } else {
      message.ethAddress = "";
    }
    return message;
  },

  toJSON(message: MsgUseNftByEVM): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.token !== undefined && (obj.token = message.token);
    message.ethSignature !== undefined &&
      (obj.ethSignature = message.ethSignature);
    message.signMessage !== undefined &&
      (obj.signMessage = message.signMessage);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUseNftByEVM>): MsgUseNftByEVM {
    const message = { ...baseMsgUseNftByEVM } as MsgUseNftByEVM;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
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
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = object.ethAddress;
    } else {
      message.ethAddress = "";
    }
    return message;
  },
};

const baseMsgUseNftByEVMResponse: object = {};

export const MsgUseNftByEVMResponse = {
  encode(_: MsgUseNftByEVMResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUseNftByEVMResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUseNftByEVMResponse } as MsgUseNftByEVMResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUseNftByEVMResponse {
    const message = { ...baseMsgUseNftByEVMResponse } as MsgUseNftByEVMResponse;
    return message;
  },

  toJSON(_: MsgUseNftByEVMResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUseNftByEVMResponse>): MsgUseNftByEVMResponse {
    const message = { ...baseMsgUseNftByEVMResponse } as MsgUseNftByEVMResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  UseNft(request: MsgUseNft): Promise<MsgUseNftResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UseNftByEVM(request: MsgUseNftByEVM): Promise<MsgUseNftByEVMResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  UseNft(request: MsgUseNft): Promise<MsgUseNftResponse> {
    const data = MsgUseNft.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Msg",
      "UseNft",
      data
    );
    return promise.then((data) => MsgUseNftResponse.decode(new Reader(data)));
  }

  UseNftByEVM(request: MsgUseNftByEVM): Promise<MsgUseNftByEVMResponse> {
    const data = MsgUseNftByEVM.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.consume.Msg",
      "UseNftByEVM",
      data
    );
    return promise.then((data) =>
      MsgUseNftByEVMResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
