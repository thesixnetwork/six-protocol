/* eslint-disable */
import {
  AllowedActioner,
  ActionParams,
  allowedActionerFromJSON,
  allowedActionerToJSON,
} from "../nftmngr/action";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.nftmngr";

export interface VirtualAction {
  virtualNftSchemaCode: string;
  name: string;
  desc: string;
  disable: boolean;
  when: string;
  then: string[];
  /** NOTE: Policy of virtual action */
  allowed_actioner: AllowedActioner;
  params: ActionParams[];
}

const baseVirtualAction: object = {
  virtualNftSchemaCode: "",
  name: "",
  desc: "",
  disable: false,
  when: "",
  then: "",
  allowed_actioner: 0,
};

export const VirtualAction = {
  encode(message: VirtualAction, writer: Writer = Writer.create()): Writer {
    if (message.virtualNftSchemaCode !== "") {
      writer.uint32(10).string(message.virtualNftSchemaCode);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.desc !== "") {
      writer.uint32(26).string(message.desc);
    }
    if (message.disable === true) {
      writer.uint32(32).bool(message.disable);
    }
    if (message.when !== "") {
      writer.uint32(42).string(message.when);
    }
    for (const v of message.then) {
      writer.uint32(50).string(v!);
    }
    if (message.allowed_actioner !== 0) {
      writer.uint32(56).int32(message.allowed_actioner);
    }
    for (const v of message.params) {
      ActionParams.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualAction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualAction } as VirtualAction;
    message.then = [];
    message.params = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualNftSchemaCode = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.desc = reader.string();
          break;
        case 4:
          message.disable = reader.bool();
          break;
        case 5:
          message.when = reader.string();
          break;
        case 6:
          message.then.push(reader.string());
          break;
        case 7:
          message.allowed_actioner = reader.int32() as any;
          break;
        case 8:
          message.params.push(ActionParams.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualAction {
    const message = { ...baseVirtualAction } as VirtualAction;
    message.then = [];
    message.params = [];
    if (
      object.virtualNftSchemaCode !== undefined &&
      object.virtualNftSchemaCode !== null
    ) {
      message.virtualNftSchemaCode = String(object.virtualNftSchemaCode);
    } else {
      message.virtualNftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.desc !== undefined && object.desc !== null) {
      message.desc = String(object.desc);
    } else {
      message.desc = "";
    }
    if (object.disable !== undefined && object.disable !== null) {
      message.disable = Boolean(object.disable);
    } else {
      message.disable = false;
    }
    if (object.when !== undefined && object.when !== null) {
      message.when = String(object.when);
    } else {
      message.when = "";
    }
    if (object.then !== undefined && object.then !== null) {
      for (const e of object.then) {
        message.then.push(String(e));
      }
    }
    if (
      object.allowed_actioner !== undefined &&
      object.allowed_actioner !== null
    ) {
      message.allowed_actioner = allowedActionerFromJSON(
        object.allowed_actioner
      );
    } else {
      message.allowed_actioner = 0;
    }
    if (object.params !== undefined && object.params !== null) {
      for (const e of object.params) {
        message.params.push(ActionParams.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: VirtualAction): unknown {
    const obj: any = {};
    message.virtualNftSchemaCode !== undefined &&
      (obj.virtualNftSchemaCode = message.virtualNftSchemaCode);
    message.name !== undefined && (obj.name = message.name);
    message.desc !== undefined && (obj.desc = message.desc);
    message.disable !== undefined && (obj.disable = message.disable);
    message.when !== undefined && (obj.when = message.when);
    if (message.then) {
      obj.then = message.then.map((e) => e);
    } else {
      obj.then = [];
    }
    message.allowed_actioner !== undefined &&
      (obj.allowed_actioner = allowedActionerToJSON(message.allowed_actioner));
    if (message.params) {
      obj.params = message.params.map((e) =>
        e ? ActionParams.toJSON(e) : undefined
      );
    } else {
      obj.params = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<VirtualAction>): VirtualAction {
    const message = { ...baseVirtualAction } as VirtualAction;
    message.then = [];
    message.params = [];
    if (
      object.virtualNftSchemaCode !== undefined &&
      object.virtualNftSchemaCode !== null
    ) {
      message.virtualNftSchemaCode = object.virtualNftSchemaCode;
    } else {
      message.virtualNftSchemaCode = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.desc !== undefined && object.desc !== null) {
      message.desc = object.desc;
    } else {
      message.desc = "";
    }
    if (object.disable !== undefined && object.disable !== null) {
      message.disable = object.disable;
    } else {
      message.disable = false;
    }
    if (object.when !== undefined && object.when !== null) {
      message.when = object.when;
    } else {
      message.when = "";
    }
    if (object.then !== undefined && object.then !== null) {
      for (const e of object.then) {
        message.then.push(e);
      }
    }
    if (
      object.allowed_actioner !== undefined &&
      object.allowed_actioner !== null
    ) {
      message.allowed_actioner = object.allowed_actioner;
    } else {
      message.allowed_actioner = 0;
    }
    if (object.params !== undefined && object.params !== null) {
      for (const e of object.params) {
        message.params.push(ActionParams.fromPartial(e));
      }
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
