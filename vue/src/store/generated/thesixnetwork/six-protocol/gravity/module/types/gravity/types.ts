/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Metadata } from "../cosmos/bank/v1beta1/bank";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "gravity";

/** BridgeValidator represents a validator's ETH address and its power */
export interface BridgeValidator {
  power: number;
  ethereum_address: string;
}

/**
 * Valset is the Ethereum Bridge Multsig Set, each gravity validator also
 * maintains an ETH key to sign messages, these are used to check signatures on
 * ETH because of the significant gas savings
 */
export interface Valset {
  nonce: number;
  members: BridgeValidator[];
  height: number;
  reward_amount: string;
  /** the reward token in it's Ethereum hex address representation */
  reward_token: string;
}

/**
 * LastObservedEthereumBlockHeight stores the last observed
 * Ethereum block height along with the Cosmos block height that
 * it was observed at. These two numbers can be used to project
 * outward and always produce batches with timeouts in the future
 * even if no Ethereum block height has been relayed for a long time
 */
export interface LastObservedEthereumBlockHeight {
  cosmos_block_height: number;
  ethereum_block_height: number;
}

/**
 * This records the relationship between an ERC20 token and the denom
 * of the corresponding Cosmos originated asset
 */
export interface ERC20ToDenom {
  erc20: string;
  denom: string;
}

/**
 * UnhaltBridgeProposal defines a custom governance proposal useful for restoring
 * the bridge after a oracle disagreement. Once this proposal is passed bridge state will roll back events
 * to the nonce provided in target_nonce if and only if those events have not yet been observed (executed on the Cosmos chain). This allows for easy
 * handling of cases where for example an Ethereum hardfork has occured and more than 1/3 of the vlaidtor set
 * disagrees with the rest. Normally this would require a chain halt, manual genesis editing and restar to resolve
 * with this feature a governance proposal can be used instead
 */
export interface UnhaltBridgeProposal {
  title: string;
  description: string;
  target_nonce: number;
}

/**
 * AirdropProposal defines a custom governance proposal type that allows an airdrop to occur in a decentralized
 * fashion. A list of destination addresses and an amount per airdrop recipient is provided. The funds for this
 * airdrop are removed from the Community Pool, if the community pool does not have sufficient funding to perform
 * the airdrop to all provided recipients nothing will occur
 */
export interface AirdropProposal {
  title: string;
  description: string;
  denom: string;
  recipients: Uint8Array;
  amounts: number[];
}

/**
 * IBCMetadataProposal defines a custom governance proposal type that allows governance to set the
 * metadata for an IBC token, this will allow Gravity to deploy an ERC20 representing this token on
 * Ethereum
 * Name: the token name
 * Symbol: the token symbol
 * Description: the token description, not sent to ETH at all, only used on Cosmos
 * Display: the token display name (only used on Cosmos to decide ERC20 Decimals)
 * Deicmals: the decimals for the display unit
 * ibc_denom is the denom of the token in question on this chain
 */
export interface IBCMetadataProposal {
  title: string;
  description: string;
  metadata: Metadata | undefined;
  ibc_denom: string;
}

/**
 * PendingIbcAutoForward represents a SendToCosmos transaction with a foreign CosmosReceiver which will be added to the
 * PendingIbcAutoForward queue in attestation_handler and sent over IBC on some submission of a MsgExecuteIbcAutoForwards
 */
export interface PendingIbcAutoForward {
  /** the destination address. sdk.AccAddress does not preserve foreign prefixes */
  foreign_receiver: string;
  /** the token sent from ethereum to the ibc-enabled chain over `IbcChannel` */
  token: Coin | undefined;
  /** the IBC channel to send `Amount` over via ibc-transfer module */
  ibc_channel: string;
  /** the EventNonce from the MsgSendToCosmosClaim, used for ordering the queue */
  event_nonce: number;
}

const baseBridgeValidator: object = { power: 0, ethereum_address: "" };

export const BridgeValidator = {
  encode(message: BridgeValidator, writer: Writer = Writer.create()): Writer {
    if (message.power !== 0) {
      writer.uint32(8).uint64(message.power);
    }
    if (message.ethereum_address !== "") {
      writer.uint32(18).string(message.ethereum_address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BridgeValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBridgeValidator } as BridgeValidator;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.power = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.ethereum_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BridgeValidator {
    const message = { ...baseBridgeValidator } as BridgeValidator;
    if (object.power !== undefined && object.power !== null) {
      message.power = Number(object.power);
    } else {
      message.power = 0;
    }
    if (
      object.ethereum_address !== undefined &&
      object.ethereum_address !== null
    ) {
      message.ethereum_address = String(object.ethereum_address);
    } else {
      message.ethereum_address = "";
    }
    return message;
  },

  toJSON(message: BridgeValidator): unknown {
    const obj: any = {};
    message.power !== undefined && (obj.power = message.power);
    message.ethereum_address !== undefined &&
      (obj.ethereum_address = message.ethereum_address);
    return obj;
  },

  fromPartial(object: DeepPartial<BridgeValidator>): BridgeValidator {
    const message = { ...baseBridgeValidator } as BridgeValidator;
    if (object.power !== undefined && object.power !== null) {
      message.power = object.power;
    } else {
      message.power = 0;
    }
    if (
      object.ethereum_address !== undefined &&
      object.ethereum_address !== null
    ) {
      message.ethereum_address = object.ethereum_address;
    } else {
      message.ethereum_address = "";
    }
    return message;
  },
};

const baseValset: object = {
  nonce: 0,
  height: 0,
  reward_amount: "",
  reward_token: "",
};

export const Valset = {
  encode(message: Valset, writer: Writer = Writer.create()): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    for (const v of message.members) {
      BridgeValidator.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.height !== 0) {
      writer.uint32(24).uint64(message.height);
    }
    if (message.reward_amount !== "") {
      writer.uint32(34).string(message.reward_amount);
    }
    if (message.reward_token !== "") {
      writer.uint32(42).string(message.reward_token);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Valset {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseValset } as Valset;
    message.members = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.members.push(BridgeValidator.decode(reader, reader.uint32()));
          break;
        case 3:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.reward_amount = reader.string();
          break;
        case 5:
          message.reward_token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Valset {
    const message = { ...baseValset } as Valset;
    message.members = [];
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(BridgeValidator.fromJSON(e));
      }
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    if (object.reward_amount !== undefined && object.reward_amount !== null) {
      message.reward_amount = String(object.reward_amount);
    } else {
      message.reward_amount = "";
    }
    if (object.reward_token !== undefined && object.reward_token !== null) {
      message.reward_token = String(object.reward_token);
    } else {
      message.reward_token = "";
    }
    return message;
  },

  toJSON(message: Valset): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    if (message.members) {
      obj.members = message.members.map((e) =>
        e ? BridgeValidator.toJSON(e) : undefined
      );
    } else {
      obj.members = [];
    }
    message.height !== undefined && (obj.height = message.height);
    message.reward_amount !== undefined &&
      (obj.reward_amount = message.reward_amount);
    message.reward_token !== undefined &&
      (obj.reward_token = message.reward_token);
    return obj;
  },

  fromPartial(object: DeepPartial<Valset>): Valset {
    const message = { ...baseValset } as Valset;
    message.members = [];
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(BridgeValidator.fromPartial(e));
      }
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    if (object.reward_amount !== undefined && object.reward_amount !== null) {
      message.reward_amount = object.reward_amount;
    } else {
      message.reward_amount = "";
    }
    if (object.reward_token !== undefined && object.reward_token !== null) {
      message.reward_token = object.reward_token;
    } else {
      message.reward_token = "";
    }
    return message;
  },
};

const baseLastObservedEthereumBlockHeight: object = {
  cosmos_block_height: 0,
  ethereum_block_height: 0,
};

export const LastObservedEthereumBlockHeight = {
  encode(
    message: LastObservedEthereumBlockHeight,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cosmos_block_height !== 0) {
      writer.uint32(8).uint64(message.cosmos_block_height);
    }
    if (message.ethereum_block_height !== 0) {
      writer.uint32(16).uint64(message.ethereum_block_height);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LastObservedEthereumBlockHeight {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLastObservedEthereumBlockHeight,
    } as LastObservedEthereumBlockHeight;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cosmos_block_height = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.ethereum_block_height = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LastObservedEthereumBlockHeight {
    const message = {
      ...baseLastObservedEthereumBlockHeight,
    } as LastObservedEthereumBlockHeight;
    if (
      object.cosmos_block_height !== undefined &&
      object.cosmos_block_height !== null
    ) {
      message.cosmos_block_height = Number(object.cosmos_block_height);
    } else {
      message.cosmos_block_height = 0;
    }
    if (
      object.ethereum_block_height !== undefined &&
      object.ethereum_block_height !== null
    ) {
      message.ethereum_block_height = Number(object.ethereum_block_height);
    } else {
      message.ethereum_block_height = 0;
    }
    return message;
  },

  toJSON(message: LastObservedEthereumBlockHeight): unknown {
    const obj: any = {};
    message.cosmos_block_height !== undefined &&
      (obj.cosmos_block_height = message.cosmos_block_height);
    message.ethereum_block_height !== undefined &&
      (obj.ethereum_block_height = message.ethereum_block_height);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LastObservedEthereumBlockHeight>
  ): LastObservedEthereumBlockHeight {
    const message = {
      ...baseLastObservedEthereumBlockHeight,
    } as LastObservedEthereumBlockHeight;
    if (
      object.cosmos_block_height !== undefined &&
      object.cosmos_block_height !== null
    ) {
      message.cosmos_block_height = object.cosmos_block_height;
    } else {
      message.cosmos_block_height = 0;
    }
    if (
      object.ethereum_block_height !== undefined &&
      object.ethereum_block_height !== null
    ) {
      message.ethereum_block_height = object.ethereum_block_height;
    } else {
      message.ethereum_block_height = 0;
    }
    return message;
  },
};

const baseERC20ToDenom: object = { erc20: "", denom: "" };

export const ERC20ToDenom = {
  encode(message: ERC20ToDenom, writer: Writer = Writer.create()): Writer {
    if (message.erc20 !== "") {
      writer.uint32(10).string(message.erc20);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ERC20ToDenom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseERC20ToDenom } as ERC20ToDenom;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erc20 = reader.string();
          break;
        case 2:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ERC20ToDenom {
    const message = { ...baseERC20ToDenom } as ERC20ToDenom;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = String(object.erc20);
    } else {
      message.erc20 = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: ERC20ToDenom): unknown {
    const obj: any = {};
    message.erc20 !== undefined && (obj.erc20 = message.erc20);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(object: DeepPartial<ERC20ToDenom>): ERC20ToDenom {
    const message = { ...baseERC20ToDenom } as ERC20ToDenom;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = object.erc20;
    } else {
      message.erc20 = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseUnhaltBridgeProposal: object = {
  title: "",
  description: "",
  target_nonce: 0,
};

export const UnhaltBridgeProposal = {
  encode(
    message: UnhaltBridgeProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.target_nonce !== 0) {
      writer.uint32(32).uint64(message.target_nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UnhaltBridgeProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUnhaltBridgeProposal } as UnhaltBridgeProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 4:
          message.target_nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnhaltBridgeProposal {
    const message = { ...baseUnhaltBridgeProposal } as UnhaltBridgeProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.target_nonce !== undefined && object.target_nonce !== null) {
      message.target_nonce = Number(object.target_nonce);
    } else {
      message.target_nonce = 0;
    }
    return message;
  },

  toJSON(message: UnhaltBridgeProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.target_nonce !== undefined &&
      (obj.target_nonce = message.target_nonce);
    return obj;
  },

  fromPartial(object: DeepPartial<UnhaltBridgeProposal>): UnhaltBridgeProposal {
    const message = { ...baseUnhaltBridgeProposal } as UnhaltBridgeProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.target_nonce !== undefined && object.target_nonce !== null) {
      message.target_nonce = object.target_nonce;
    } else {
      message.target_nonce = 0;
    }
    return message;
  },
};

const baseAirdropProposal: object = {
  title: "",
  description: "",
  denom: "",
  amounts: 0,
};

export const AirdropProposal = {
  encode(message: AirdropProposal, writer: Writer = Writer.create()): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.recipients.length !== 0) {
      writer.uint32(34).bytes(message.recipients);
    }
    writer.uint32(42).fork();
    for (const v of message.amounts) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AirdropProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAirdropProposal } as AirdropProposal;
    message.amounts = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.recipients = reader.bytes();
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.amounts.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.amounts.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AirdropProposal {
    const message = { ...baseAirdropProposal } as AirdropProposal;
    message.amounts = [];
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.recipients !== undefined && object.recipients !== null) {
      message.recipients = bytesFromBase64(object.recipients);
    }
    if (object.amounts !== undefined && object.amounts !== null) {
      for (const e of object.amounts) {
        message.amounts.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: AirdropProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.denom !== undefined && (obj.denom = message.denom);
    message.recipients !== undefined &&
      (obj.recipients = base64FromBytes(
        message.recipients !== undefined ? message.recipients : new Uint8Array()
      ));
    if (message.amounts) {
      obj.amounts = message.amounts.map((e) => e);
    } else {
      obj.amounts = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AirdropProposal>): AirdropProposal {
    const message = { ...baseAirdropProposal } as AirdropProposal;
    message.amounts = [];
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.recipients !== undefined && object.recipients !== null) {
      message.recipients = object.recipients;
    } else {
      message.recipients = new Uint8Array();
    }
    if (object.amounts !== undefined && object.amounts !== null) {
      for (const e of object.amounts) {
        message.amounts.push(e);
      }
    }
    return message;
  },
};

const baseIBCMetadataProposal: object = {
  title: "",
  description: "",
  ibc_denom: "",
};

export const IBCMetadataProposal = {
  encode(
    message: IBCMetadataProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(26).fork()).ldelim();
    }
    if (message.ibc_denom !== "") {
      writer.uint32(34).string(message.ibc_denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IBCMetadataProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIBCMetadataProposal } as IBCMetadataProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 4:
          message.ibc_denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IBCMetadataProposal {
    const message = { ...baseIBCMetadataProposal } as IBCMetadataProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.ibc_denom !== undefined && object.ibc_denom !== null) {
      message.ibc_denom = String(object.ibc_denom);
    } else {
      message.ibc_denom = "";
    }
    return message;
  },

  toJSON(message: IBCMetadataProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    message.ibc_denom !== undefined && (obj.ibc_denom = message.ibc_denom);
    return obj;
  },

  fromPartial(object: DeepPartial<IBCMetadataProposal>): IBCMetadataProposal {
    const message = { ...baseIBCMetadataProposal } as IBCMetadataProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.ibc_denom !== undefined && object.ibc_denom !== null) {
      message.ibc_denom = object.ibc_denom;
    } else {
      message.ibc_denom = "";
    }
    return message;
  },
};

const basePendingIbcAutoForward: object = {
  foreign_receiver: "",
  ibc_channel: "",
  event_nonce: 0,
};

export const PendingIbcAutoForward = {
  encode(
    message: PendingIbcAutoForward,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.foreign_receiver !== "") {
      writer.uint32(10).string(message.foreign_receiver);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(18).fork()).ldelim();
    }
    if (message.ibc_channel !== "") {
      writer.uint32(26).string(message.ibc_channel);
    }
    if (message.event_nonce !== 0) {
      writer.uint32(32).uint64(message.event_nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PendingIbcAutoForward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePendingIbcAutoForward } as PendingIbcAutoForward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.foreign_receiver = reader.string();
          break;
        case 2:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.ibc_channel = reader.string();
          break;
        case 4:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingIbcAutoForward {
    const message = { ...basePendingIbcAutoForward } as PendingIbcAutoForward;
    if (
      object.foreign_receiver !== undefined &&
      object.foreign_receiver !== null
    ) {
      message.foreign_receiver = String(object.foreign_receiver);
    } else {
      message.foreign_receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Coin.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    if (object.ibc_channel !== undefined && object.ibc_channel !== null) {
      message.ibc_channel = String(object.ibc_channel);
    } else {
      message.ibc_channel = "";
    }
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    return message;
  },

  toJSON(message: PendingIbcAutoForward): unknown {
    const obj: any = {};
    message.foreign_receiver !== undefined &&
      (obj.foreign_receiver = message.foreign_receiver);
    message.token !== undefined &&
      (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    message.ibc_channel !== undefined &&
      (obj.ibc_channel = message.ibc_channel);
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<PendingIbcAutoForward>
  ): PendingIbcAutoForward {
    const message = { ...basePendingIbcAutoForward } as PendingIbcAutoForward;
    if (
      object.foreign_receiver !== undefined &&
      object.foreign_receiver !== null
    ) {
      message.foreign_receiver = object.foreign_receiver;
    } else {
      message.foreign_receiver = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Coin.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    if (object.ibc_channel !== undefined && object.ibc_channel !== null) {
      message.ibc_channel = object.ibc_channel;
    } else {
      message.ibc_channel = "";
    }
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
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
