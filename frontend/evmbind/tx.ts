/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.evmbind";

export interface MsgCreateBinding {
  creator: string;
  ethAddress: string;
  ethSignature: string;
  signMessage: string;
}

export interface MsgCreateBindingResponse {}

export interface MsgUpdateBinding {
  creator: string;
  ethAddress: string;
  ethSignature: string;
  signMessage: string;
}

export interface MsgUpdateBindingResponse {}

export interface MsgDeleteBinding {
  creator: string;
  ethAddress: string;
}

export interface MsgDeleteBindingResponse {}

export interface MsgEthSend {
  creator: string;
  fromEth: string;
  toEth: string;
  amount: string;
}

export interface MsgEthSendResponse {}

const baseMsgCreateBinding: object = {
  creator: "",
  ethAddress: "",
  ethSignature: "",
  signMessage: "",
};

export const MsgCreateBinding = {
  encode(message: MsgCreateBinding, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ethAddress !== "") {
      writer.uint32(18).string(message.ethAddress);
    }
    if (message.ethSignature !== "") {
      writer.uint32(26).string(message.ethSignature);
    }
    if (message.signMessage !== "") {
      writer.uint32(34).string(message.signMessage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBinding {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateBinding } as MsgCreateBinding;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ethAddress = reader.string();
          break;
        case 3:
          message.ethSignature = reader.string();
          break;
        case 4:
          message.signMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateBinding {
    const message = { ...baseMsgCreateBinding } as MsgCreateBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgCreateBinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    message.ethSignature !== undefined &&
      (obj.ethSignature = message.ethSignature);
    message.signMessage !== undefined &&
      (obj.signMessage = message.signMessage);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateBinding>): MsgCreateBinding {
    const message = { ...baseMsgCreateBinding } as MsgCreateBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgCreateBindingResponse: object = {};

export const MsgCreateBindingResponse = {
  encode(
    _: MsgCreateBindingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateBindingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateBindingResponse,
    } as MsgCreateBindingResponse;
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

  fromJSON(_: any): MsgCreateBindingResponse {
    const message = {
      ...baseMsgCreateBindingResponse,
    } as MsgCreateBindingResponse;
    return message;
  },

  toJSON(_: MsgCreateBindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateBindingResponse>
  ): MsgCreateBindingResponse {
    const message = {
      ...baseMsgCreateBindingResponse,
    } as MsgCreateBindingResponse;
    return message;
  },
};

const baseMsgUpdateBinding: object = {
  creator: "",
  ethAddress: "",
  ethSignature: "",
  signMessage: "",
};

export const MsgUpdateBinding = {
  encode(message: MsgUpdateBinding, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ethAddress !== "") {
      writer.uint32(18).string(message.ethAddress);
    }
    if (message.ethSignature !== "") {
      writer.uint32(26).string(message.ethSignature);
    }
    if (message.signMessage !== "") {
      writer.uint32(34).string(message.signMessage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBinding {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateBinding } as MsgUpdateBinding;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ethAddress = reader.string();
          break;
        case 3:
          message.ethSignature = reader.string();
          break;
        case 4:
          message.signMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateBinding {
    const message = { ...baseMsgUpdateBinding } as MsgUpdateBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgUpdateBinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    message.ethSignature !== undefined &&
      (obj.ethSignature = message.ethSignature);
    message.signMessage !== undefined &&
      (obj.signMessage = message.signMessage);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateBinding>): MsgUpdateBinding {
    const message = { ...baseMsgUpdateBinding } as MsgUpdateBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgUpdateBindingResponse: object = {};

export const MsgUpdateBindingResponse = {
  encode(
    _: MsgUpdateBindingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateBindingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateBindingResponse,
    } as MsgUpdateBindingResponse;
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

  fromJSON(_: any): MsgUpdateBindingResponse {
    const message = {
      ...baseMsgUpdateBindingResponse,
    } as MsgUpdateBindingResponse;
    return message;
  },

  toJSON(_: MsgUpdateBindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateBindingResponse>
  ): MsgUpdateBindingResponse {
    const message = {
      ...baseMsgUpdateBindingResponse,
    } as MsgUpdateBindingResponse;
    return message;
  },
};

const baseMsgDeleteBinding: object = { creator: "", ethAddress: "" };

export const MsgDeleteBinding = {
  encode(message: MsgDeleteBinding, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ethAddress !== "") {
      writer.uint32(18).string(message.ethAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBinding {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteBinding } as MsgDeleteBinding;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ethAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteBinding {
    const message = { ...baseMsgDeleteBinding } as MsgDeleteBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = String(object.ethAddress);
    } else {
      message.ethAddress = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteBinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteBinding>): MsgDeleteBinding {
    const message = { ...baseMsgDeleteBinding } as MsgDeleteBinding;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = object.ethAddress;
    } else {
      message.ethAddress = "";
    }
    return message;
  },
};

const baseMsgDeleteBindingResponse: object = {};

export const MsgDeleteBindingResponse = {
  encode(
    _: MsgDeleteBindingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteBindingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteBindingResponse,
    } as MsgDeleteBindingResponse;
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

  fromJSON(_: any): MsgDeleteBindingResponse {
    const message = {
      ...baseMsgDeleteBindingResponse,
    } as MsgDeleteBindingResponse;
    return message;
  },

  toJSON(_: MsgDeleteBindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteBindingResponse>
  ): MsgDeleteBindingResponse {
    const message = {
      ...baseMsgDeleteBindingResponse,
    } as MsgDeleteBindingResponse;
    return message;
  },
};

const baseMsgEthSend: object = {
  creator: "",
  fromEth: "",
  toEth: "",
  amount: "",
};

export const MsgEthSend = {
  encode(message: MsgEthSend, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fromEth !== "") {
      writer.uint32(18).string(message.fromEth);
    }
    if (message.toEth !== "") {
      writer.uint32(26).string(message.toEth);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEthSend {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEthSend } as MsgEthSend;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fromEth = reader.string();
          break;
        case 3:
          message.toEth = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEthSend {
    const message = { ...baseMsgEthSend } as MsgEthSend;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fromEth !== undefined && object.fromEth !== null) {
      message.fromEth = String(object.fromEth);
    } else {
      message.fromEth = "";
    }
    if (object.toEth !== undefined && object.toEth !== null) {
      message.toEth = String(object.toEth);
    } else {
      message.toEth = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: MsgEthSend): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fromEth !== undefined && (obj.fromEth = message.fromEth);
    message.toEth !== undefined && (obj.toEth = message.toEth);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEthSend>): MsgEthSend {
    const message = { ...baseMsgEthSend } as MsgEthSend;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fromEth !== undefined && object.fromEth !== null) {
      message.fromEth = object.fromEth;
    } else {
      message.fromEth = "";
    }
    if (object.toEth !== undefined && object.toEth !== null) {
      message.toEth = object.toEth;
    } else {
      message.toEth = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgEthSendResponse: object = {};

export const MsgEthSendResponse = {
  encode(_: MsgEthSendResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEthSendResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEthSendResponse } as MsgEthSendResponse;
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

  fromJSON(_: any): MsgEthSendResponse {
    const message = { ...baseMsgEthSendResponse } as MsgEthSendResponse;
    return message;
  },

  toJSON(_: MsgEthSendResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgEthSendResponse>): MsgEthSendResponse {
    const message = { ...baseMsgEthSendResponse } as MsgEthSendResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateBinding(request: MsgCreateBinding): Promise<MsgCreateBindingResponse>;
  UpdateBinding(request: MsgUpdateBinding): Promise<MsgUpdateBindingResponse>;
  DeleteBinding(request: MsgDeleteBinding): Promise<MsgDeleteBindingResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EthSend(request: MsgEthSend): Promise<MsgEthSendResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateBinding(request: MsgCreateBinding): Promise<MsgCreateBindingResponse> {
    const data = MsgCreateBinding.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Msg",
      "CreateBinding",
      data
    );
    return promise.then((data) =>
      MsgCreateBindingResponse.decode(new Reader(data))
    );
  }

  UpdateBinding(request: MsgUpdateBinding): Promise<MsgUpdateBindingResponse> {
    const data = MsgUpdateBinding.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Msg",
      "UpdateBinding",
      data
    );
    return promise.then((data) =>
      MsgUpdateBindingResponse.decode(new Reader(data))
    );
  }

  DeleteBinding(request: MsgDeleteBinding): Promise<MsgDeleteBindingResponse> {
    const data = MsgDeleteBinding.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Msg",
      "DeleteBinding",
      data
    );
    return promise.then((data) =>
      MsgDeleteBindingResponse.decode(new Reader(data))
    );
  }

  EthSend(request: MsgEthSend): Promise<MsgEthSendResponse> {
    const data = MsgEthSend.encode(request).finish();
    const promise = this.rpc.request(
      "thesixnetwork.sixprotocol.evmbind.Msg",
      "EthSend",
      data
    );
    return promise.then((data) => MsgEthSendResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
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
