/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Action } from "../nftmngr/action";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixprotocol.nftmngr";

export enum RegistryStatus {
  PENDING = 0,
  REJECT = 1,
  ACCEPT = 2,
  UNRECOGNIZED = -1,
}

export function registryStatusFromJSON(object: any): RegistryStatus {
  switch (object) {
    case 0:
    case "PENDING":
      return RegistryStatus.PENDING;
    case 1:
    case "REJECT":
      return RegistryStatus.REJECT;
    case 2:
    case "ACCEPT":
      return RegistryStatus.ACCEPT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RegistryStatus.UNRECOGNIZED;
  }
}

export function registryStatusToJSON(object: RegistryStatus): string {
  switch (object) {
    case RegistryStatus.PENDING:
      return "PENDING";
    case RegistryStatus.REJECT:
      return "REJECT";
    case RegistryStatus.ACCEPT:
      return "ACCEPT";
    default:
      return "UNKNOWN";
  }
}

export enum ProposalType {
  CREATE = 0,
  EDIT = 1,
  UNRECOGNIZED = -1,
}

export function proposalTypeFromJSON(object: any): ProposalType {
  switch (object) {
    case 0:
    case "CREATE":
      return ProposalType.CREATE;
    case 1:
    case "EDIT":
      return ProposalType.EDIT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ProposalType.UNRECOGNIZED;
  }
}

export function proposalTypeToJSON(object: ProposalType): string {
  switch (object) {
    case ProposalType.CREATE:
      return "CREATE";
    case ProposalType.EDIT:
      return "EDIT";
    default:
      return "UNKNOWN";
  }
}

export interface VirtualSchemaProposal {
  id: string;
  proposalType: ProposalType;
  virtualSchema: VirtualSchema | undefined;
  actions: Action[];
  executors: string[];
  submitTime: Date | undefined;
  votingStartTime: Date | undefined;
  votingEndTime: Date | undefined;
}

export interface VirtualSchemaProposalRequest {
  virtualSchemaCode: string;
  virtualSchemaRegistry: string[];
  actions: Action[];
  executors: string[];
  enable: boolean;
}

export interface VirtualSchema {
  virtualNftSchemaCode: string;
  registry: VirtualSchemaRegistry[];
  /** link wih virtual action. We will not put in here to reduce redundance info */
  enable: boolean;
}

export interface VirtualSchemaRegistry {
  nftSchemaCode: string;
  decision: RegistryStatus;
}

export interface VirtualSchemaRegistryRequest {
  nftSchemaCode: string;
}

export interface ActiveVirtualSchemaProposal {
  id: string;
}

export interface InactiveVirtualSchemaProposal {
  id: string;
}

const baseVirtualSchemaProposal: object = {
  id: "",
  proposalType: 0,
  executors: "",
};

export const VirtualSchemaProposal = {
  encode(
    message: VirtualSchemaProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.proposalType !== 0) {
      writer.uint32(16).int32(message.proposalType);
    }
    if (message.virtualSchema !== undefined) {
      VirtualSchema.encode(
        message.virtualSchema,
        writer.uint32(26).fork()
      ).ldelim();
    }
    for (const v of message.actions) {
      Action.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.executors) {
      writer.uint32(42).string(v!);
    }
    if (message.submitTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.submitTime),
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.votingStartTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.votingStartTime),
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.votingEndTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.votingEndTime),
        writer.uint32(66).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualSchemaProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualSchemaProposal } as VirtualSchemaProposal;
    message.actions = [];
    message.executors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.proposalType = reader.int32() as any;
          break;
        case 3:
          message.virtualSchema = VirtualSchema.decode(reader, reader.uint32());
          break;
        case 4:
          message.actions.push(Action.decode(reader, reader.uint32()));
          break;
        case 5:
          message.executors.push(reader.string());
          break;
        case 6:
          message.submitTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.votingStartTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.votingEndTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualSchemaProposal {
    const message = { ...baseVirtualSchemaProposal } as VirtualSchemaProposal;
    message.actions = [];
    message.executors = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.proposalType !== undefined && object.proposalType !== null) {
      message.proposalType = proposalTypeFromJSON(object.proposalType);
    } else {
      message.proposalType = 0;
    }
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      message.virtualSchema = VirtualSchema.fromJSON(object.virtualSchema);
    } else {
      message.virtualSchema = undefined;
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromJSON(e));
      }
    }
    if (object.executors !== undefined && object.executors !== null) {
      for (const e of object.executors) {
        message.executors.push(String(e));
      }
    }
    if (object.submitTime !== undefined && object.submitTime !== null) {
      message.submitTime = fromJsonTimestamp(object.submitTime);
    } else {
      message.submitTime = undefined;
    }
    if (
      object.votingStartTime !== undefined &&
      object.votingStartTime !== null
    ) {
      message.votingStartTime = fromJsonTimestamp(object.votingStartTime);
    } else {
      message.votingStartTime = undefined;
    }
    if (object.votingEndTime !== undefined && object.votingEndTime !== null) {
      message.votingEndTime = fromJsonTimestamp(object.votingEndTime);
    } else {
      message.votingEndTime = undefined;
    }
    return message;
  },

  toJSON(message: VirtualSchemaProposal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.proposalType !== undefined &&
      (obj.proposalType = proposalTypeToJSON(message.proposalType));
    message.virtualSchema !== undefined &&
      (obj.virtualSchema = message.virtualSchema
        ? VirtualSchema.toJSON(message.virtualSchema)
        : undefined);
    if (message.actions) {
      obj.actions = message.actions.map((e) =>
        e ? Action.toJSON(e) : undefined
      );
    } else {
      obj.actions = [];
    }
    if (message.executors) {
      obj.executors = message.executors.map((e) => e);
    } else {
      obj.executors = [];
    }
    message.submitTime !== undefined &&
      (obj.submitTime =
        message.submitTime !== undefined
          ? message.submitTime.toISOString()
          : null);
    message.votingStartTime !== undefined &&
      (obj.votingStartTime =
        message.votingStartTime !== undefined
          ? message.votingStartTime.toISOString()
          : null);
    message.votingEndTime !== undefined &&
      (obj.votingEndTime =
        message.votingEndTime !== undefined
          ? message.votingEndTime.toISOString()
          : null);
    return obj;
  },

  fromPartial(
    object: DeepPartial<VirtualSchemaProposal>
  ): VirtualSchemaProposal {
    const message = { ...baseVirtualSchemaProposal } as VirtualSchemaProposal;
    message.actions = [];
    message.executors = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.proposalType !== undefined && object.proposalType !== null) {
      message.proposalType = object.proposalType;
    } else {
      message.proposalType = 0;
    }
    if (object.virtualSchema !== undefined && object.virtualSchema !== null) {
      message.virtualSchema = VirtualSchema.fromPartial(object.virtualSchema);
    } else {
      message.virtualSchema = undefined;
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromPartial(e));
      }
    }
    if (object.executors !== undefined && object.executors !== null) {
      for (const e of object.executors) {
        message.executors.push(e);
      }
    }
    if (object.submitTime !== undefined && object.submitTime !== null) {
      message.submitTime = object.submitTime;
    } else {
      message.submitTime = undefined;
    }
    if (
      object.votingStartTime !== undefined &&
      object.votingStartTime !== null
    ) {
      message.votingStartTime = object.votingStartTime;
    } else {
      message.votingStartTime = undefined;
    }
    if (object.votingEndTime !== undefined && object.votingEndTime !== null) {
      message.votingEndTime = object.votingEndTime;
    } else {
      message.votingEndTime = undefined;
    }
    return message;
  },
};

const baseVirtualSchemaProposalRequest: object = {
  virtualSchemaCode: "",
  virtualSchemaRegistry: "",
  executors: "",
  enable: false,
};

export const VirtualSchemaProposalRequest = {
  encode(
    message: VirtualSchemaProposalRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.virtualSchemaCode !== "") {
      writer.uint32(10).string(message.virtualSchemaCode);
    }
    for (const v of message.virtualSchemaRegistry) {
      writer.uint32(18).string(v!);
    }
    for (const v of message.actions) {
      Action.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.executors) {
      writer.uint32(34).string(v!);
    }
    if (message.enable === true) {
      writer.uint32(40).bool(message.enable);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): VirtualSchemaProposalRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseVirtualSchemaProposalRequest,
    } as VirtualSchemaProposalRequest;
    message.virtualSchemaRegistry = [];
    message.actions = [];
    message.executors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualSchemaCode = reader.string();
          break;
        case 2:
          message.virtualSchemaRegistry.push(reader.string());
          break;
        case 3:
          message.actions.push(Action.decode(reader, reader.uint32()));
          break;
        case 4:
          message.executors.push(reader.string());
          break;
        case 5:
          message.enable = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualSchemaProposalRequest {
    const message = {
      ...baseVirtualSchemaProposalRequest,
    } as VirtualSchemaProposalRequest;
    message.virtualSchemaRegistry = [];
    message.actions = [];
    message.executors = [];
    if (
      object.virtualSchemaCode !== undefined &&
      object.virtualSchemaCode !== null
    ) {
      message.virtualSchemaCode = String(object.virtualSchemaCode);
    } else {
      message.virtualSchemaCode = "";
    }
    if (
      object.virtualSchemaRegistry !== undefined &&
      object.virtualSchemaRegistry !== null
    ) {
      for (const e of object.virtualSchemaRegistry) {
        message.virtualSchemaRegistry.push(String(e));
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromJSON(e));
      }
    }
    if (object.executors !== undefined && object.executors !== null) {
      for (const e of object.executors) {
        message.executors.push(String(e));
      }
    }
    if (object.enable !== undefined && object.enable !== null) {
      message.enable = Boolean(object.enable);
    } else {
      message.enable = false;
    }
    return message;
  },

  toJSON(message: VirtualSchemaProposalRequest): unknown {
    const obj: any = {};
    message.virtualSchemaCode !== undefined &&
      (obj.virtualSchemaCode = message.virtualSchemaCode);
    if (message.virtualSchemaRegistry) {
      obj.virtualSchemaRegistry = message.virtualSchemaRegistry.map((e) => e);
    } else {
      obj.virtualSchemaRegistry = [];
    }
    if (message.actions) {
      obj.actions = message.actions.map((e) =>
        e ? Action.toJSON(e) : undefined
      );
    } else {
      obj.actions = [];
    }
    if (message.executors) {
      obj.executors = message.executors.map((e) => e);
    } else {
      obj.executors = [];
    }
    message.enable !== undefined && (obj.enable = message.enable);
    return obj;
  },

  fromPartial(
    object: DeepPartial<VirtualSchemaProposalRequest>
  ): VirtualSchemaProposalRequest {
    const message = {
      ...baseVirtualSchemaProposalRequest,
    } as VirtualSchemaProposalRequest;
    message.virtualSchemaRegistry = [];
    message.actions = [];
    message.executors = [];
    if (
      object.virtualSchemaCode !== undefined &&
      object.virtualSchemaCode !== null
    ) {
      message.virtualSchemaCode = object.virtualSchemaCode;
    } else {
      message.virtualSchemaCode = "";
    }
    if (
      object.virtualSchemaRegistry !== undefined &&
      object.virtualSchemaRegistry !== null
    ) {
      for (const e of object.virtualSchemaRegistry) {
        message.virtualSchemaRegistry.push(e);
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromPartial(e));
      }
    }
    if (object.executors !== undefined && object.executors !== null) {
      for (const e of object.executors) {
        message.executors.push(e);
      }
    }
    if (object.enable !== undefined && object.enable !== null) {
      message.enable = object.enable;
    } else {
      message.enable = false;
    }
    return message;
  },
};

const baseVirtualSchema: object = { virtualNftSchemaCode: "", enable: false };

export const VirtualSchema = {
  encode(message: VirtualSchema, writer: Writer = Writer.create()): Writer {
    if (message.virtualNftSchemaCode !== "") {
      writer.uint32(10).string(message.virtualNftSchemaCode);
    }
    for (const v of message.registry) {
      VirtualSchemaRegistry.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.enable === true) {
      writer.uint32(24).bool(message.enable);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualSchema {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualSchema } as VirtualSchema;
    message.registry = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualNftSchemaCode = reader.string();
          break;
        case 2:
          message.registry.push(
            VirtualSchemaRegistry.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.enable = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualSchema {
    const message = { ...baseVirtualSchema } as VirtualSchema;
    message.registry = [];
    if (
      object.virtualNftSchemaCode !== undefined &&
      object.virtualNftSchemaCode !== null
    ) {
      message.virtualNftSchemaCode = String(object.virtualNftSchemaCode);
    } else {
      message.virtualNftSchemaCode = "";
    }
    if (object.registry !== undefined && object.registry !== null) {
      for (const e of object.registry) {
        message.registry.push(VirtualSchemaRegistry.fromJSON(e));
      }
    }
    if (object.enable !== undefined && object.enable !== null) {
      message.enable = Boolean(object.enable);
    } else {
      message.enable = false;
    }
    return message;
  },

  toJSON(message: VirtualSchema): unknown {
    const obj: any = {};
    message.virtualNftSchemaCode !== undefined &&
      (obj.virtualNftSchemaCode = message.virtualNftSchemaCode);
    if (message.registry) {
      obj.registry = message.registry.map((e) =>
        e ? VirtualSchemaRegistry.toJSON(e) : undefined
      );
    } else {
      obj.registry = [];
    }
    message.enable !== undefined && (obj.enable = message.enable);
    return obj;
  },

  fromPartial(object: DeepPartial<VirtualSchema>): VirtualSchema {
    const message = { ...baseVirtualSchema } as VirtualSchema;
    message.registry = [];
    if (
      object.virtualNftSchemaCode !== undefined &&
      object.virtualNftSchemaCode !== null
    ) {
      message.virtualNftSchemaCode = object.virtualNftSchemaCode;
    } else {
      message.virtualNftSchemaCode = "";
    }
    if (object.registry !== undefined && object.registry !== null) {
      for (const e of object.registry) {
        message.registry.push(VirtualSchemaRegistry.fromPartial(e));
      }
    }
    if (object.enable !== undefined && object.enable !== null) {
      message.enable = object.enable;
    } else {
      message.enable = false;
    }
    return message;
  },
};

const baseVirtualSchemaRegistry: object = { nftSchemaCode: "", decision: 0 };

export const VirtualSchemaRegistry = {
  encode(
    message: VirtualSchemaRegistry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.decision !== 0) {
      writer.uint32(16).int32(message.decision);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualSchemaRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualSchemaRegistry } as VirtualSchemaRegistry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.decision = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualSchemaRegistry {
    const message = { ...baseVirtualSchemaRegistry } as VirtualSchemaRegistry;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.decision !== undefined && object.decision !== null) {
      message.decision = registryStatusFromJSON(object.decision);
    } else {
      message.decision = 0;
    }
    return message;
  },

  toJSON(message: VirtualSchemaRegistry): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.decision !== undefined &&
      (obj.decision = registryStatusToJSON(message.decision));
    return obj;
  },

  fromPartial(
    object: DeepPartial<VirtualSchemaRegistry>
  ): VirtualSchemaRegistry {
    const message = { ...baseVirtualSchemaRegistry } as VirtualSchemaRegistry;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.decision !== undefined && object.decision !== null) {
      message.decision = object.decision;
    } else {
      message.decision = 0;
    }
    return message;
  },
};

const baseVirtualSchemaRegistryRequest: object = { nftSchemaCode: "" };

export const VirtualSchemaRegistryRequest = {
  encode(
    message: VirtualSchemaRegistryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): VirtualSchemaRegistryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseVirtualSchemaRegistryRequest,
    } as VirtualSchemaRegistryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualSchemaRegistryRequest {
    const message = {
      ...baseVirtualSchemaRegistryRequest,
    } as VirtualSchemaRegistryRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },

  toJSON(message: VirtualSchemaRegistryRequest): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    return obj;
  },

  fromPartial(
    object: DeepPartial<VirtualSchemaRegistryRequest>
  ): VirtualSchemaRegistryRequest {
    const message = {
      ...baseVirtualSchemaRegistryRequest,
    } as VirtualSchemaRegistryRequest;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    return message;
  },
};

const baseActiveVirtualSchemaProposal: object = { id: "" };

export const ActiveVirtualSchemaProposal = {
  encode(
    message: ActiveVirtualSchemaProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ActiveVirtualSchemaProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseActiveVirtualSchemaProposal,
    } as ActiveVirtualSchemaProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ActiveVirtualSchemaProposal {
    const message = {
      ...baseActiveVirtualSchemaProposal,
    } as ActiveVirtualSchemaProposal;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: ActiveVirtualSchemaProposal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ActiveVirtualSchemaProposal>
  ): ActiveVirtualSchemaProposal {
    const message = {
      ...baseActiveVirtualSchemaProposal,
    } as ActiveVirtualSchemaProposal;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseInactiveVirtualSchemaProposal: object = { id: "" };

export const InactiveVirtualSchemaProposal = {
  encode(
    message: InactiveVirtualSchemaProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): InactiveVirtualSchemaProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseInactiveVirtualSchemaProposal,
    } as InactiveVirtualSchemaProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InactiveVirtualSchemaProposal {
    const message = {
      ...baseInactiveVirtualSchemaProposal,
    } as InactiveVirtualSchemaProposal;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: InactiveVirtualSchemaProposal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<InactiveVirtualSchemaProposal>
  ): InactiveVirtualSchemaProposal {
    const message = {
      ...baseInactiveVirtualSchemaProposal,
    } as InactiveVirtualSchemaProposal;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}
