/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Any } from "../../google/protobuf/any";

export const protobufPackage = "gravity.v1";

export enum ClaimType {
  CLAIM_TYPE_UNSPECIFIED = 0,
  CLAIM_TYPE_SEND_TO_COSMOS = 1,
  CLAIM_TYPE_BATCH_SEND_TO_ETH = 2,
  CLAIM_TYPE_ERC20_DEPLOYED = 3,
  CLAIM_TYPE_LOGIC_CALL_EXECUTED = 4,
  CLAIM_TYPE_VALSET_UPDATED = 5,
  UNRECOGNIZED = -1,
}

export function claimTypeFromJSON(object: any): ClaimType {
  switch (object) {
    case 0:
    case "CLAIM_TYPE_UNSPECIFIED":
      return ClaimType.CLAIM_TYPE_UNSPECIFIED;
    case 1:
    case "CLAIM_TYPE_SEND_TO_COSMOS":
      return ClaimType.CLAIM_TYPE_SEND_TO_COSMOS;
    case 2:
    case "CLAIM_TYPE_BATCH_SEND_TO_ETH":
      return ClaimType.CLAIM_TYPE_BATCH_SEND_TO_ETH;
    case 3:
    case "CLAIM_TYPE_ERC20_DEPLOYED":
      return ClaimType.CLAIM_TYPE_ERC20_DEPLOYED;
    case 4:
    case "CLAIM_TYPE_LOGIC_CALL_EXECUTED":
      return ClaimType.CLAIM_TYPE_LOGIC_CALL_EXECUTED;
    case 5:
    case "CLAIM_TYPE_VALSET_UPDATED":
      return ClaimType.CLAIM_TYPE_VALSET_UPDATED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ClaimType.UNRECOGNIZED;
  }
}

export function claimTypeToJSON(object: ClaimType): string {
  switch (object) {
    case ClaimType.CLAIM_TYPE_UNSPECIFIED:
      return "CLAIM_TYPE_UNSPECIFIED";
    case ClaimType.CLAIM_TYPE_SEND_TO_COSMOS:
      return "CLAIM_TYPE_SEND_TO_COSMOS";
    case ClaimType.CLAIM_TYPE_BATCH_SEND_TO_ETH:
      return "CLAIM_TYPE_BATCH_SEND_TO_ETH";
    case ClaimType.CLAIM_TYPE_ERC20_DEPLOYED:
      return "CLAIM_TYPE_ERC20_DEPLOYED";
    case ClaimType.CLAIM_TYPE_LOGIC_CALL_EXECUTED:
      return "CLAIM_TYPE_LOGIC_CALL_EXECUTED";
    case ClaimType.CLAIM_TYPE_VALSET_UPDATED:
      return "CLAIM_TYPE_VALSET_UPDATED";
    default:
      return "UNKNOWN";
  }
}

/**
 * Attestation is an aggregate of `claims` that eventually becomes `observed` by
 * all orchestrators
 * EVENT_NONCE:
 * EventNonce a nonce provided by the gravity contract that is unique per event fired
 * These event nonces must be relayed in order. This is a correctness issue,
 * if relaying out of order transaction replay attacks become possible
 * OBSERVED:
 * Observed indicates that >67% of validators have attested to the event,
 * and that the event should be executed by the gravity state machine
 *
 * The actual content of the claims is passed in with the transaction making the claim
 * and then passed through the call stack alongside the attestation while it is processed
 * the key in which the attestation is stored is keyed on the exact details of the claim
 * but there is no reason to store those exact details becuause the next message sender
 * will kindly provide you with them.
 */
export interface Attestation {
  observed: boolean;
  votes: string[];
  height: number;
  claim: Any | undefined;
}

/**
 * ERC20Token unique identifier for an Ethereum ERC20 token.
 * CONTRACT:
 * The contract address on ETH of the token, this could be a Cosmos
 * originated token, if so it will be the ERC20 address of the representation
 * (note: developers should look up the token symbol using the address on ETH to display for UI)
 */
export interface ERC20Token {
  contract: string;
  amount: string;
}

export interface EventObservation {
  attestation_type: string;
  bridge_contract: string;
  bridge_chain_id: string;
  attestation_id: string;
  nonce: string;
}

export interface EventInvalidSendToCosmosReceiver {
  amount: string;
  nonce: string;
  token: string;
  sender: string;
}

export interface EventSendToCosmos {
  amount: string;
  nonce: string;
  token: string;
}

export interface EventSendToCosmosLocal {
  nonce: string;
  receiver: string;
  token: string;
  amount: string;
}

export interface EventSendToCosmosPendingIbcAutoForward {
  nonce: string;
  receiver: string;
  token: string;
  amount: string;
  channel: string;
}

export interface EventSendToCosmosExecutedIbcAutoForward {
  nonce: string;
  receiver: string;
  token: string;
  amount: string;
  channel: string;
  timeout_time: string;
  timeout_height: string;
}

const baseAttestation: object = { observed: false, votes: "", height: 0 };

export const Attestation = {
  encode(message: Attestation, writer: Writer = Writer.create()): Writer {
    if (message.observed === true) {
      writer.uint32(8).bool(message.observed);
    }
    for (const v of message.votes) {
      writer.uint32(18).string(v!);
    }
    if (message.height !== 0) {
      writer.uint32(24).uint64(message.height);
    }
    if (message.claim !== undefined) {
      Any.encode(message.claim, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Attestation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAttestation } as Attestation;
    message.votes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.observed = reader.bool();
          break;
        case 2:
          message.votes.push(reader.string());
          break;
        case 3:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.claim = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Attestation {
    const message = { ...baseAttestation } as Attestation;
    message.votes = [];
    if (object.observed !== undefined && object.observed !== null) {
      message.observed = Boolean(object.observed);
    } else {
      message.observed = false;
    }
    if (object.votes !== undefined && object.votes !== null) {
      for (const e of object.votes) {
        message.votes.push(String(e));
      }
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    if (object.claim !== undefined && object.claim !== null) {
      message.claim = Any.fromJSON(object.claim);
    } else {
      message.claim = undefined;
    }
    return message;
  },

  toJSON(message: Attestation): unknown {
    const obj: any = {};
    message.observed !== undefined && (obj.observed = message.observed);
    if (message.votes) {
      obj.votes = message.votes.map((e) => e);
    } else {
      obj.votes = [];
    }
    message.height !== undefined && (obj.height = message.height);
    message.claim !== undefined &&
      (obj.claim = message.claim ? Any.toJSON(message.claim) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Attestation>): Attestation {
    const message = { ...baseAttestation } as Attestation;
    message.votes = [];
    if (object.observed !== undefined && object.observed !== null) {
      message.observed = object.observed;
    } else {
      message.observed = false;
    }
    if (object.votes !== undefined && object.votes !== null) {
      for (const e of object.votes) {
        message.votes.push(e);
      }
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    if (object.claim !== undefined && object.claim !== null) {
      message.claim = Any.fromPartial(object.claim);
    } else {
      message.claim = undefined;
    }
    return message;
  },
};

const baseERC20Token: object = { contract: "", amount: "" };

export const ERC20Token = {
  encode(message: ERC20Token, writer: Writer = Writer.create()): Writer {
    if (message.contract !== "") {
      writer.uint32(10).string(message.contract);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ERC20Token {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseERC20Token } as ERC20Token;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contract = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ERC20Token {
    const message = { ...baseERC20Token } as ERC20Token;
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = String(object.contract);
    } else {
      message.contract = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: ERC20Token): unknown {
    const obj: any = {};
    message.contract !== undefined && (obj.contract = message.contract);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<ERC20Token>): ERC20Token {
    const message = { ...baseERC20Token } as ERC20Token;
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = object.contract;
    } else {
      message.contract = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseEventObservation: object = {
  attestation_type: "",
  bridge_contract: "",
  bridge_chain_id: "",
  attestation_id: "",
  nonce: "",
};

export const EventObservation = {
  encode(message: EventObservation, writer: Writer = Writer.create()): Writer {
    if (message.attestation_type !== "") {
      writer.uint32(10).string(message.attestation_type);
    }
    if (message.bridge_contract !== "") {
      writer.uint32(18).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(26).string(message.bridge_chain_id);
    }
    if (message.attestation_id !== "") {
      writer.uint32(34).string(message.attestation_id);
    }
    if (message.nonce !== "") {
      writer.uint32(42).string(message.nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventObservation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventObservation } as EventObservation;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.attestation_type = reader.string();
          break;
        case 2:
          message.bridge_contract = reader.string();
          break;
        case 3:
          message.bridge_chain_id = reader.string();
          break;
        case 4:
          message.attestation_id = reader.string();
          break;
        case 5:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventObservation {
    const message = { ...baseEventObservation } as EventObservation;
    if (
      object.attestation_type !== undefined &&
      object.attestation_type !== null
    ) {
      message.attestation_type = String(object.attestation_type);
    } else {
      message.attestation_type = "";
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
    if (object.attestation_id !== undefined && object.attestation_id !== null) {
      message.attestation_id = String(object.attestation_id);
    } else {
      message.attestation_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventObservation): unknown {
    const obj: any = {};
    message.attestation_type !== undefined &&
      (obj.attestation_type = message.attestation_type);
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.attestation_id !== undefined &&
      (obj.attestation_id = message.attestation_id);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(object: DeepPartial<EventObservation>): EventObservation {
    const message = { ...baseEventObservation } as EventObservation;
    if (
      object.attestation_type !== undefined &&
      object.attestation_type !== null
    ) {
      message.attestation_type = object.attestation_type;
    } else {
      message.attestation_type = "";
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
    if (object.attestation_id !== undefined && object.attestation_id !== null) {
      message.attestation_id = object.attestation_id;
    } else {
      message.attestation_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventInvalidSendToCosmosReceiver: object = {
  amount: "",
  nonce: "",
  token: "",
  sender: "",
};

export const EventInvalidSendToCosmosReceiver = {
  encode(
    message: EventInvalidSendToCosmosReceiver,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    if (message.nonce !== "") {
      writer.uint32(18).string(message.nonce);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.sender !== "") {
      writer.uint32(34).string(message.sender);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventInvalidSendToCosmosReceiver {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventInvalidSendToCosmosReceiver,
    } as EventInvalidSendToCosmosReceiver;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;
        case 2:
          message.nonce = reader.string();
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventInvalidSendToCosmosReceiver {
    const message = {
      ...baseEventInvalidSendToCosmosReceiver,
    } as EventInvalidSendToCosmosReceiver;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: EventInvalidSendToCosmosReceiver): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.token !== undefined && (obj.token = message.token);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventInvalidSendToCosmosReceiver>
  ): EventInvalidSendToCosmosReceiver {
    const message = {
      ...baseEventInvalidSendToCosmosReceiver,
    } as EventInvalidSendToCosmosReceiver;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseEventSendToCosmos: object = { amount: "", nonce: "", token: "" };

export const EventSendToCosmos = {
  encode(message: EventSendToCosmos, writer: Writer = Writer.create()): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    if (message.nonce !== "") {
      writer.uint32(18).string(message.nonce);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventSendToCosmos {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventSendToCosmos } as EventSendToCosmos;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;
        case 2:
          message.nonce = reader.string();
          break;
        case 3:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventSendToCosmos {
    const message = { ...baseEventSendToCosmos } as EventSendToCosmos;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: EventSendToCosmos): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(object: DeepPartial<EventSendToCosmos>): EventSendToCosmos {
    const message = { ...baseEventSendToCosmos } as EventSendToCosmos;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseEventSendToCosmosLocal: object = {
  nonce: "",
  receiver: "",
  token: "",
  amount: "",
};

export const EventSendToCosmosLocal = {
  encode(
    message: EventSendToCosmosLocal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== "") {
      writer.uint32(10).string(message.nonce);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventSendToCosmosLocal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventSendToCosmosLocal } as EventSendToCosmosLocal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.token = reader.string();
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

  fromJSON(object: any): EventSendToCosmosLocal {
    const message = { ...baseEventSendToCosmosLocal } as EventSendToCosmosLocal;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: EventSendToCosmosLocal): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.token !== undefined && (obj.token = message.token);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventSendToCosmosLocal>
  ): EventSendToCosmosLocal {
    const message = { ...baseEventSendToCosmosLocal } as EventSendToCosmosLocal;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseEventSendToCosmosPendingIbcAutoForward: object = {
  nonce: "",
  receiver: "",
  token: "",
  amount: "",
  channel: "",
};

export const EventSendToCosmosPendingIbcAutoForward = {
  encode(
    message: EventSendToCosmosPendingIbcAutoForward,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== "") {
      writer.uint32(10).string(message.nonce);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.channel !== "") {
      writer.uint32(42).string(message.channel);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventSendToCosmosPendingIbcAutoForward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventSendToCosmosPendingIbcAutoForward,
    } as EventSendToCosmosPendingIbcAutoForward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.channel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventSendToCosmosPendingIbcAutoForward {
    const message = {
      ...baseEventSendToCosmosPendingIbcAutoForward,
    } as EventSendToCosmosPendingIbcAutoForward;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    return message;
  },

  toJSON(message: EventSendToCosmosPendingIbcAutoForward): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.token !== undefined && (obj.token = message.token);
    message.amount !== undefined && (obj.amount = message.amount);
    message.channel !== undefined && (obj.channel = message.channel);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventSendToCosmosPendingIbcAutoForward>
  ): EventSendToCosmosPendingIbcAutoForward {
    const message = {
      ...baseEventSendToCosmosPendingIbcAutoForward,
    } as EventSendToCosmosPendingIbcAutoForward;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    return message;
  },
};

const baseEventSendToCosmosExecutedIbcAutoForward: object = {
  nonce: "",
  receiver: "",
  token: "",
  amount: "",
  channel: "",
  timeout_time: "",
  timeout_height: "",
};

export const EventSendToCosmosExecutedIbcAutoForward = {
  encode(
    message: EventSendToCosmosExecutedIbcAutoForward,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== "") {
      writer.uint32(10).string(message.nonce);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.channel !== "") {
      writer.uint32(42).string(message.channel);
    }
    if (message.timeout_time !== "") {
      writer.uint32(50).string(message.timeout_time);
    }
    if (message.timeout_height !== "") {
      writer.uint32(58).string(message.timeout_height);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventSendToCosmosExecutedIbcAutoForward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventSendToCosmosExecutedIbcAutoForward,
    } as EventSendToCosmosExecutedIbcAutoForward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.channel = reader.string();
          break;
        case 6:
          message.timeout_time = reader.string();
          break;
        case 7:
          message.timeout_height = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventSendToCosmosExecutedIbcAutoForward {
    const message = {
      ...baseEventSendToCosmosExecutedIbcAutoForward,
    } as EventSendToCosmosExecutedIbcAutoForward;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = String(object.channel);
    } else {
      message.channel = "";
    }
    if (object.timeout_time !== undefined && object.timeout_time !== null) {
      message.timeout_time = String(object.timeout_time);
    } else {
      message.timeout_time = "";
    }
    if (object.timeout_height !== undefined && object.timeout_height !== null) {
      message.timeout_height = String(object.timeout_height);
    } else {
      message.timeout_height = "";
    }
    return message;
  },

  toJSON(message: EventSendToCosmosExecutedIbcAutoForward): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.token !== undefined && (obj.token = message.token);
    message.amount !== undefined && (obj.amount = message.amount);
    message.channel !== undefined && (obj.channel = message.channel);
    message.timeout_time !== undefined &&
      (obj.timeout_time = message.timeout_time);
    message.timeout_height !== undefined &&
      (obj.timeout_height = message.timeout_height);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventSendToCosmosExecutedIbcAutoForward>
  ): EventSendToCosmosExecutedIbcAutoForward {
    const message = {
      ...baseEventSendToCosmosExecutedIbcAutoForward,
    } as EventSendToCosmosExecutedIbcAutoForward;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.channel !== undefined && object.channel !== null) {
      message.channel = object.channel;
    } else {
      message.channel = "";
    }
    if (object.timeout_time !== undefined && object.timeout_time !== null) {
      message.timeout_time = object.timeout_time;
    } else {
      message.timeout_time = "";
    }
    if (object.timeout_height !== undefined && object.timeout_height !== null) {
      message.timeout_height = object.timeout_height;
    } else {
      message.timeout_height = "";
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
