/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { BridgeValidator } from "../gravity/types";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "gravity";

/**
 * MsgSetOrchestratorAddress
 * this message allows validators to delegate their voting responsibilities
 * to a given key. This key is then used as an optional authentication method
 * for sigining oracle claims
 * VALIDATOR
 * The validator field is a cosmosvaloper1... string (i.e. sdk.ValAddress)
 * that references a validator in the active set
 * ORCHESTRATOR
 * The orchestrator field is a cosmos1... string  (i.e. sdk.AccAddress) that
 * references the key that is being delegated to
 * ETH_ADDRESS
 * This is a hex encoded 0x Ethereum public key that will be used by this validator
 * on Ethereum
 */
export interface MsgSetOrchestratorAddress {
  validator: string;
  orchestrator: string;
  eth_address: string;
}

export interface MsgSetOrchestratorAddressResponse {}

/**
 * MsgValsetConfirm
 * this is the message sent by the validators when they wish to submit their
 * signatures over the validator set at a given block height. A validator must
 * first call MsgSetEthAddress to set their Ethereum address to be used for
 * signing. Then someone (anyone) must make a ValsetRequest, the request is
 * essentially a messaging mechanism to determine which block all validators
 * should submit signatures over. Finally validators sign the validator set,
 * powers, and Ethereum addresses of the entire validator set at the height of a
 * ValsetRequest and submit that signature with this message.
 *
 * If a sufficient number of validators (66% of voting power) (A) have set
 * Ethereum addresses and (B) submit ValsetConfirm messages with their
 * signatures it is then possible for anyone to view these signatures in the
 * chain store and submit them to Ethereum to update the validator set
 * -------------
 */
export interface MsgValsetConfirm {
  nonce: number;
  orchestrator: string;
  eth_address: string;
  signature: string;
}

export interface MsgValsetConfirmResponse {}

/**
 * MsgSendToEth
 * This is the message that a user calls when they want to bridge an asset
 * it will later be removed when it is included in a batch and successfully
 * submitted tokens are removed from the users balance immediately
 * -------------
 * AMOUNT:
 * the coin to send across the bridge, note the restriction that this is a
 * single coin not a set of coins that is normal in other Cosmos messages
 * FEE:
 * the fee paid for the bridge, distinct from the fee paid to the chain to
 * actually send this message in the first place. So a successful send has
 * two layers of fees for the user
 */
export interface MsgSendToEth {
  sender: string;
  eth_dest: string;
  amount: Coin | undefined;
  bridge_fee: Coin | undefined;
}

export interface MsgSendToEthResponse {}

/**
 * MsgRequestBatch
 * this is a message anyone can send that requests a batch of transactions to
 * send across the bridge be created for whatever block height this message is
 * included in. This acts as a coordination point, the handler for this message
 * looks at the AddToOutgoingPool tx's in the store and generates a batch, also
 * available in the store tied to this message. The validators then grab this
 * batch, sign it, submit the signatures with a MsgConfirmBatch before a relayer
 * can finally submit the batch
 * -------------
 */
export interface MsgRequestBatch {
  sender: string;
  denom: string;
}

export interface MsgRequestBatchResponse {}

/**
 * MsgConfirmBatch
 * When validators observe a MsgRequestBatch they form a batch by ordering
 * transactions currently in the txqueue in order of highest to lowest fee,
 * cutting off when the batch either reaches a hardcoded maximum size (to be
 * decided, probably around 100) or when transactions stop being profitable
 * (TODO determine this without nondeterminism) This message includes the batch
 * as well as an Ethereum signature over this batch by the validator
 * -------------
 */
export interface MsgConfirmBatch {
  nonce: number;
  token_contract: string;
  eth_signer: string;
  orchestrator: string;
  signature: string;
}

export interface MsgConfirmBatchResponse {}

/**
 * MsgConfirmLogicCall
 * When validators observe a MsgRequestBatch they form a batch by ordering
 * transactions currently in the txqueue in order of highest to lowest fee,
 * cutting off when the batch either reaches a hardcoded maximum size (to be
 * decided, probably around 100) or when transactions stop being profitable
 * (TODO determine this without nondeterminism) This message includes the batch
 * as well as an Ethereum signature over this batch by the validator
 * -------------
 */
export interface MsgConfirmLogicCall {
  invalidation_id: string;
  invalidation_nonce: number;
  eth_signer: string;
  orchestrator: string;
  signature: string;
}

export interface MsgConfirmLogicCallResponse {}

/**
 * MsgSendToCosmosClaim
 * When more than 66% of the active validator set has
 * claimed to have seen the deposit enter the ethereum blockchain coins are
 * issued to the Cosmos address in question
 * -------------
 */
export interface MsgSendToCosmosClaim {
  event_nonce: number;
  block_height: number;
  token_contract: string;
  amount: string;
  ethereum_sender: string;
  cosmos_receiver: string;
  orchestrator: string;
}

export interface MsgSendToCosmosClaimResponse {}

/**
 * MsgExecuteIbcAutoForwards
 * Prompts the forwarding of Pending IBC Auto-Forwards in the queue
 * The Pending forwards will be executed in order of their original SendToCosmos.EventNonce
 * The funds in the queue will be sent to a local gravity-prefixed address if IBC transfer is not possible
 */
export interface MsgExecuteIbcAutoForwards {
  /** How many queued forwards to clear, be careful about gas limits */
  forwards_to_clear: number;
  /** This message's sender */
  executor: string;
}

export interface MsgExecuteIbcAutoForwardsResponse {}

/**
 * BatchSendToEthClaim claims that a batch of send to eth
 * operations on the bridge contract was executed.
 */
export interface MsgBatchSendToEthClaim {
  event_nonce: number;
  block_height: number;
  batch_nonce: number;
  token_contract: string;
  orchestrator: string;
}

export interface MsgBatchSendToEthClaimResponse {}

/**
 * ERC20DeployedClaim allows the Cosmos module
 * to learn about an ERC20 that someone deployed
 * to represent a Cosmos asset
 */
export interface MsgERC20DeployedClaim {
  event_nonce: number;
  block_height: number;
  cosmos_denom: string;
  token_contract: string;
  name: string;
  symbol: string;
  decimals: number;
  orchestrator: string;
}

export interface MsgERC20DeployedClaimResponse {}

/**
 * This informs the Cosmos module that a logic
 * call has been executed
 */
export interface MsgLogicCallExecutedClaim {
  event_nonce: number;
  block_height: number;
  invalidation_id: Uint8Array;
  invalidation_nonce: number;
  orchestrator: string;
}

export interface MsgLogicCallExecutedClaimResponse {}

/**
 * This informs the Cosmos module that a validator
 * set has been updated.
 */
export interface MsgValsetUpdatedClaim {
  event_nonce: number;
  valset_nonce: number;
  block_height: number;
  members: BridgeValidator[];
  reward_amount: string;
  reward_token: string;
  orchestrator: string;
}

export interface MsgValsetUpdatedClaimResponse {}

/**
 * This call allows the sender (and only the sender)
 * to cancel a given MsgSendToEth and recieve a refund
 * of the tokens
 */
export interface MsgCancelSendToEth {
  transaction_id: number;
  sender: string;
}

export interface MsgCancelSendToEthResponse {}

/**
 * This call allows anyone to submit evidence that a
 * validator has signed a valset, batch, or logic call that never
 * existed on the Cosmos chain.
 * Subject contains the batch, valset, or logic call.
 */
export interface MsgSubmitBadSignatureEvidence {
  subject: Any | undefined;
  signature: string;
  sender: string;
}

export interface MsgSubmitBadSignatureEvidenceResponse {}

export interface EventSetOperatorAddress {
  message: string;
  address: string;
}

export interface EventValsetConfirmKey {
  message: string;
  key: string;
}

export interface EventBatchCreated {
  message: string;
  batch_nonce: string;
}

export interface EventBatchConfirmKey {
  message: string;
  batch_confirm_key: string;
}

export interface EventBatchSendToEthClaim {
  nonce: string;
}

export interface EventClaim {
  message: string;
  claim_hash: string;
  attestation_id: string;
}

export interface EventBadSignatureEvidence {
  message: string;
  bad_eth_signature: string;
  bad_eth_signature_subject: string;
}

export interface EventERC20DeployedClaim {
  token: string;
  nonce: string;
}

export interface EventValsetUpdatedClaim {
  nonce: string;
}

export interface EventMultisigUpdateRequest {
  bridge_contract: string;
  bridge_chain_id: string;
  multisig_id: string;
  nonce: string;
}

export interface EventOutgoingLogicCallCanceled {
  logic_call_invalidation_id: string;
  logic_call_invalidation_nonce: string;
}

export interface EventSignatureSlashing {
  type: string;
  address: string;
}

export interface EventOutgoingTxId {
  message: string;
  tx_id: string;
}

const baseMsgSetOrchestratorAddress: object = {
  validator: "",
  orchestrator: "",
  eth_address: "",
};

export const MsgSetOrchestratorAddress = {
  encode(
    message: MsgSetOrchestratorAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.validator !== "") {
      writer.uint32(10).string(message.validator);
    }
    if (message.orchestrator !== "") {
      writer.uint32(18).string(message.orchestrator);
    }
    if (message.eth_address !== "") {
      writer.uint32(26).string(message.eth_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetOrchestratorAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetOrchestratorAddress,
    } as MsgSetOrchestratorAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = reader.string();
          break;
        case 2:
          message.orchestrator = reader.string();
          break;
        case 3:
          message.eth_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetOrchestratorAddress {
    const message = {
      ...baseMsgSetOrchestratorAddress,
    } as MsgSetOrchestratorAddress;
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = String(object.validator);
    } else {
      message.validator = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = String(object.eth_address);
    } else {
      message.eth_address = "";
    }
    return message;
  },

  toJSON(message: MsgSetOrchestratorAddress): unknown {
    const obj: any = {};
    message.validator !== undefined && (obj.validator = message.validator);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    message.eth_address !== undefined &&
      (obj.eth_address = message.eth_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetOrchestratorAddress>
  ): MsgSetOrchestratorAddress {
    const message = {
      ...baseMsgSetOrchestratorAddress,
    } as MsgSetOrchestratorAddress;
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = object.validator;
    } else {
      message.validator = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = object.eth_address;
    } else {
      message.eth_address = "";
    }
    return message;
  },
};

const baseMsgSetOrchestratorAddressResponse: object = {};

export const MsgSetOrchestratorAddressResponse = {
  encode(
    _: MsgSetOrchestratorAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetOrchestratorAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetOrchestratorAddressResponse,
    } as MsgSetOrchestratorAddressResponse;
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

  fromJSON(_: any): MsgSetOrchestratorAddressResponse {
    const message = {
      ...baseMsgSetOrchestratorAddressResponse,
    } as MsgSetOrchestratorAddressResponse;
    return message;
  },

  toJSON(_: MsgSetOrchestratorAddressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSetOrchestratorAddressResponse>
  ): MsgSetOrchestratorAddressResponse {
    const message = {
      ...baseMsgSetOrchestratorAddressResponse,
    } as MsgSetOrchestratorAddressResponse;
    return message;
  },
};

const baseMsgValsetConfirm: object = {
  nonce: 0,
  orchestrator: "",
  eth_address: "",
  signature: "",
};

export const MsgValsetConfirm = {
  encode(message: MsgValsetConfirm, writer: Writer = Writer.create()): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    if (message.orchestrator !== "") {
      writer.uint32(18).string(message.orchestrator);
    }
    if (message.eth_address !== "") {
      writer.uint32(26).string(message.eth_address);
    }
    if (message.signature !== "") {
      writer.uint32(34).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgValsetConfirm {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgValsetConfirm } as MsgValsetConfirm;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.orchestrator = reader.string();
          break;
        case 3:
          message.eth_address = reader.string();
          break;
        case 4:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgValsetConfirm {
    const message = { ...baseMsgValsetConfirm } as MsgValsetConfirm;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = String(object.eth_address);
    } else {
      message.eth_address = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgValsetConfirm): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    message.eth_address !== undefined &&
      (obj.eth_address = message.eth_address);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgValsetConfirm>): MsgValsetConfirm {
    const message = { ...baseMsgValsetConfirm } as MsgValsetConfirm;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = object.eth_address;
    } else {
      message.eth_address = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgValsetConfirmResponse: object = {};

export const MsgValsetConfirmResponse = {
  encode(
    _: MsgValsetConfirmResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgValsetConfirmResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgValsetConfirmResponse,
    } as MsgValsetConfirmResponse;
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

  fromJSON(_: any): MsgValsetConfirmResponse {
    const message = {
      ...baseMsgValsetConfirmResponse,
    } as MsgValsetConfirmResponse;
    return message;
  },

  toJSON(_: MsgValsetConfirmResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgValsetConfirmResponse>
  ): MsgValsetConfirmResponse {
    const message = {
      ...baseMsgValsetConfirmResponse,
    } as MsgValsetConfirmResponse;
    return message;
  },
};

const baseMsgSendToEth: object = { sender: "", eth_dest: "" };

export const MsgSendToEth = {
  encode(message: MsgSendToEth, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.eth_dest !== "") {
      writer.uint32(18).string(message.eth_dest);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    if (message.bridge_fee !== undefined) {
      Coin.encode(message.bridge_fee, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendToEth {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendToEth } as MsgSendToEth;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.eth_dest = reader.string();
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.bridge_fee = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendToEth {
    const message = { ...baseMsgSendToEth } as MsgSendToEth;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.eth_dest !== undefined && object.eth_dest !== null) {
      message.eth_dest = String(object.eth_dest);
    } else {
      message.eth_dest = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.bridge_fee !== undefined && object.bridge_fee !== null) {
      message.bridge_fee = Coin.fromJSON(object.bridge_fee);
    } else {
      message.bridge_fee = undefined;
    }
    return message;
  },

  toJSON(message: MsgSendToEth): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.eth_dest !== undefined && (obj.eth_dest = message.eth_dest);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.bridge_fee !== undefined &&
      (obj.bridge_fee = message.bridge_fee
        ? Coin.toJSON(message.bridge_fee)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendToEth>): MsgSendToEth {
    const message = { ...baseMsgSendToEth } as MsgSendToEth;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.eth_dest !== undefined && object.eth_dest !== null) {
      message.eth_dest = object.eth_dest;
    } else {
      message.eth_dest = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    if (object.bridge_fee !== undefined && object.bridge_fee !== null) {
      message.bridge_fee = Coin.fromPartial(object.bridge_fee);
    } else {
      message.bridge_fee = undefined;
    }
    return message;
  },
};

const baseMsgSendToEthResponse: object = {};

export const MsgSendToEthResponse = {
  encode(_: MsgSendToEthResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendToEthResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendToEthResponse } as MsgSendToEthResponse;
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

  fromJSON(_: any): MsgSendToEthResponse {
    const message = { ...baseMsgSendToEthResponse } as MsgSendToEthResponse;
    return message;
  },

  toJSON(_: MsgSendToEthResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSendToEthResponse>): MsgSendToEthResponse {
    const message = { ...baseMsgSendToEthResponse } as MsgSendToEthResponse;
    return message;
  },
};

const baseMsgRequestBatch: object = { sender: "", denom: "" };

export const MsgRequestBatch = {
  encode(message: MsgRequestBatch, writer: Writer = Writer.create()): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestBatch {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestBatch } as MsgRequestBatch;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
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

  fromJSON(object: any): MsgRequestBatch {
    const message = { ...baseMsgRequestBatch } as MsgRequestBatch;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: MsgRequestBatch): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequestBatch>): MsgRequestBatch {
    const message = { ...baseMsgRequestBatch } as MsgRequestBatch;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseMsgRequestBatchResponse: object = {};

export const MsgRequestBatchResponse = {
  encode(_: MsgRequestBatchResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestBatchResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestBatchResponse,
    } as MsgRequestBatchResponse;
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

  fromJSON(_: any): MsgRequestBatchResponse {
    const message = {
      ...baseMsgRequestBatchResponse,
    } as MsgRequestBatchResponse;
    return message;
  },

  toJSON(_: MsgRequestBatchResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestBatchResponse>
  ): MsgRequestBatchResponse {
    const message = {
      ...baseMsgRequestBatchResponse,
    } as MsgRequestBatchResponse;
    return message;
  },
};

const baseMsgConfirmBatch: object = {
  nonce: 0,
  token_contract: "",
  eth_signer: "",
  orchestrator: "",
  signature: "",
};

export const MsgConfirmBatch = {
  encode(message: MsgConfirmBatch, writer: Writer = Writer.create()): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    if (message.token_contract !== "") {
      writer.uint32(18).string(message.token_contract);
    }
    if (message.eth_signer !== "") {
      writer.uint32(26).string(message.eth_signer);
    }
    if (message.orchestrator !== "") {
      writer.uint32(34).string(message.orchestrator);
    }
    if (message.signature !== "") {
      writer.uint32(42).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfirmBatch {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfirmBatch } as MsgConfirmBatch;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.token_contract = reader.string();
          break;
        case 3:
          message.eth_signer = reader.string();
          break;
        case 4:
          message.orchestrator = reader.string();
          break;
        case 5:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfirmBatch {
    const message = { ...baseMsgConfirmBatch } as MsgConfirmBatch;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = String(object.token_contract);
    } else {
      message.token_contract = "";
    }
    if (object.eth_signer !== undefined && object.eth_signer !== null) {
      message.eth_signer = String(object.eth_signer);
    } else {
      message.eth_signer = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgConfirmBatch): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.token_contract !== undefined &&
      (obj.token_contract = message.token_contract);
    message.eth_signer !== undefined && (obj.eth_signer = message.eth_signer);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConfirmBatch>): MsgConfirmBatch {
    const message = { ...baseMsgConfirmBatch } as MsgConfirmBatch;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = object.token_contract;
    } else {
      message.token_contract = "";
    }
    if (object.eth_signer !== undefined && object.eth_signer !== null) {
      message.eth_signer = object.eth_signer;
    } else {
      message.eth_signer = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgConfirmBatchResponse: object = {};

export const MsgConfirmBatchResponse = {
  encode(_: MsgConfirmBatchResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfirmBatchResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgConfirmBatchResponse,
    } as MsgConfirmBatchResponse;
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

  fromJSON(_: any): MsgConfirmBatchResponse {
    const message = {
      ...baseMsgConfirmBatchResponse,
    } as MsgConfirmBatchResponse;
    return message;
  },

  toJSON(_: MsgConfirmBatchResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgConfirmBatchResponse>
  ): MsgConfirmBatchResponse {
    const message = {
      ...baseMsgConfirmBatchResponse,
    } as MsgConfirmBatchResponse;
    return message;
  },
};

const baseMsgConfirmLogicCall: object = {
  invalidation_id: "",
  invalidation_nonce: 0,
  eth_signer: "",
  orchestrator: "",
  signature: "",
};

export const MsgConfirmLogicCall = {
  encode(
    message: MsgConfirmLogicCall,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.invalidation_id !== "") {
      writer.uint32(10).string(message.invalidation_id);
    }
    if (message.invalidation_nonce !== 0) {
      writer.uint32(16).uint64(message.invalidation_nonce);
    }
    if (message.eth_signer !== "") {
      writer.uint32(26).string(message.eth_signer);
    }
    if (message.orchestrator !== "") {
      writer.uint32(34).string(message.orchestrator);
    }
    if (message.signature !== "") {
      writer.uint32(42).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfirmLogicCall {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfirmLogicCall } as MsgConfirmLogicCall;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invalidation_id = reader.string();
          break;
        case 2:
          message.invalidation_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.eth_signer = reader.string();
          break;
        case 4:
          message.orchestrator = reader.string();
          break;
        case 5:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfirmLogicCall {
    const message = { ...baseMsgConfirmLogicCall } as MsgConfirmLogicCall;
    if (
      object.invalidation_id !== undefined &&
      object.invalidation_id !== null
    ) {
      message.invalidation_id = String(object.invalidation_id);
    } else {
      message.invalidation_id = "";
    }
    if (
      object.invalidation_nonce !== undefined &&
      object.invalidation_nonce !== null
    ) {
      message.invalidation_nonce = Number(object.invalidation_nonce);
    } else {
      message.invalidation_nonce = 0;
    }
    if (object.eth_signer !== undefined && object.eth_signer !== null) {
      message.eth_signer = String(object.eth_signer);
    } else {
      message.eth_signer = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgConfirmLogicCall): unknown {
    const obj: any = {};
    message.invalidation_id !== undefined &&
      (obj.invalidation_id = message.invalidation_id);
    message.invalidation_nonce !== undefined &&
      (obj.invalidation_nonce = message.invalidation_nonce);
    message.eth_signer !== undefined && (obj.eth_signer = message.eth_signer);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConfirmLogicCall>): MsgConfirmLogicCall {
    const message = { ...baseMsgConfirmLogicCall } as MsgConfirmLogicCall;
    if (
      object.invalidation_id !== undefined &&
      object.invalidation_id !== null
    ) {
      message.invalidation_id = object.invalidation_id;
    } else {
      message.invalidation_id = "";
    }
    if (
      object.invalidation_nonce !== undefined &&
      object.invalidation_nonce !== null
    ) {
      message.invalidation_nonce = object.invalidation_nonce;
    } else {
      message.invalidation_nonce = 0;
    }
    if (object.eth_signer !== undefined && object.eth_signer !== null) {
      message.eth_signer = object.eth_signer;
    } else {
      message.eth_signer = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgConfirmLogicCallResponse: object = {};

export const MsgConfirmLogicCallResponse = {
  encode(
    _: MsgConfirmLogicCallResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgConfirmLogicCallResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgConfirmLogicCallResponse,
    } as MsgConfirmLogicCallResponse;
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

  fromJSON(_: any): MsgConfirmLogicCallResponse {
    const message = {
      ...baseMsgConfirmLogicCallResponse,
    } as MsgConfirmLogicCallResponse;
    return message;
  },

  toJSON(_: MsgConfirmLogicCallResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgConfirmLogicCallResponse>
  ): MsgConfirmLogicCallResponse {
    const message = {
      ...baseMsgConfirmLogicCallResponse,
    } as MsgConfirmLogicCallResponse;
    return message;
  },
};

const baseMsgSendToCosmosClaim: object = {
  event_nonce: 0,
  block_height: 0,
  token_contract: "",
  amount: "",
  ethereum_sender: "",
  cosmos_receiver: "",
  orchestrator: "",
};

export const MsgSendToCosmosClaim = {
  encode(
    message: MsgSendToCosmosClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    if (message.block_height !== 0) {
      writer.uint32(16).uint64(message.block_height);
    }
    if (message.token_contract !== "") {
      writer.uint32(26).string(message.token_contract);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.ethereum_sender !== "") {
      writer.uint32(42).string(message.ethereum_sender);
    }
    if (message.cosmos_receiver !== "") {
      writer.uint32(50).string(message.cosmos_receiver);
    }
    if (message.orchestrator !== "") {
      writer.uint32(58).string(message.orchestrator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendToCosmosClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendToCosmosClaim } as MsgSendToCosmosClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.block_height = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.token_contract = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.ethereum_sender = reader.string();
          break;
        case 6:
          message.cosmos_receiver = reader.string();
          break;
        case 7:
          message.orchestrator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendToCosmosClaim {
    const message = { ...baseMsgSendToCosmosClaim } as MsgSendToCosmosClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = Number(object.block_height);
    } else {
      message.block_height = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = String(object.token_contract);
    } else {
      message.token_contract = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (
      object.ethereum_sender !== undefined &&
      object.ethereum_sender !== null
    ) {
      message.ethereum_sender = String(object.ethereum_sender);
    } else {
      message.ethereum_sender = "";
    }
    if (
      object.cosmos_receiver !== undefined &&
      object.cosmos_receiver !== null
    ) {
      message.cosmos_receiver = String(object.cosmos_receiver);
    } else {
      message.cosmos_receiver = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    return message;
  },

  toJSON(message: MsgSendToCosmosClaim): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    message.block_height !== undefined &&
      (obj.block_height = message.block_height);
    message.token_contract !== undefined &&
      (obj.token_contract = message.token_contract);
    message.amount !== undefined && (obj.amount = message.amount);
    message.ethereum_sender !== undefined &&
      (obj.ethereum_sender = message.ethereum_sender);
    message.cosmos_receiver !== undefined &&
      (obj.cosmos_receiver = message.cosmos_receiver);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendToCosmosClaim>): MsgSendToCosmosClaim {
    const message = { ...baseMsgSendToCosmosClaim } as MsgSendToCosmosClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = object.block_height;
    } else {
      message.block_height = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = object.token_contract;
    } else {
      message.token_contract = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (
      object.ethereum_sender !== undefined &&
      object.ethereum_sender !== null
    ) {
      message.ethereum_sender = object.ethereum_sender;
    } else {
      message.ethereum_sender = "";
    }
    if (
      object.cosmos_receiver !== undefined &&
      object.cosmos_receiver !== null
    ) {
      message.cosmos_receiver = object.cosmos_receiver;
    } else {
      message.cosmos_receiver = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    return message;
  },
};

const baseMsgSendToCosmosClaimResponse: object = {};

export const MsgSendToCosmosClaimResponse = {
  encode(
    _: MsgSendToCosmosClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendToCosmosClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendToCosmosClaimResponse,
    } as MsgSendToCosmosClaimResponse;
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

  fromJSON(_: any): MsgSendToCosmosClaimResponse {
    const message = {
      ...baseMsgSendToCosmosClaimResponse,
    } as MsgSendToCosmosClaimResponse;
    return message;
  },

  toJSON(_: MsgSendToCosmosClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendToCosmosClaimResponse>
  ): MsgSendToCosmosClaimResponse {
    const message = {
      ...baseMsgSendToCosmosClaimResponse,
    } as MsgSendToCosmosClaimResponse;
    return message;
  },
};

const baseMsgExecuteIbcAutoForwards: object = {
  forwards_to_clear: 0,
  executor: "",
};

export const MsgExecuteIbcAutoForwards = {
  encode(
    message: MsgExecuteIbcAutoForwards,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.forwards_to_clear !== 0) {
      writer.uint32(8).uint64(message.forwards_to_clear);
    }
    if (message.executor !== "") {
      writer.uint32(18).string(message.executor);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgExecuteIbcAutoForwards {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgExecuteIbcAutoForwards,
    } as MsgExecuteIbcAutoForwards;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.forwards_to_clear = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.executor = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExecuteIbcAutoForwards {
    const message = {
      ...baseMsgExecuteIbcAutoForwards,
    } as MsgExecuteIbcAutoForwards;
    if (
      object.forwards_to_clear !== undefined &&
      object.forwards_to_clear !== null
    ) {
      message.forwards_to_clear = Number(object.forwards_to_clear);
    } else {
      message.forwards_to_clear = 0;
    }
    if (object.executor !== undefined && object.executor !== null) {
      message.executor = String(object.executor);
    } else {
      message.executor = "";
    }
    return message;
  },

  toJSON(message: MsgExecuteIbcAutoForwards): unknown {
    const obj: any = {};
    message.forwards_to_clear !== undefined &&
      (obj.forwards_to_clear = message.forwards_to_clear);
    message.executor !== undefined && (obj.executor = message.executor);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgExecuteIbcAutoForwards>
  ): MsgExecuteIbcAutoForwards {
    const message = {
      ...baseMsgExecuteIbcAutoForwards,
    } as MsgExecuteIbcAutoForwards;
    if (
      object.forwards_to_clear !== undefined &&
      object.forwards_to_clear !== null
    ) {
      message.forwards_to_clear = object.forwards_to_clear;
    } else {
      message.forwards_to_clear = 0;
    }
    if (object.executor !== undefined && object.executor !== null) {
      message.executor = object.executor;
    } else {
      message.executor = "";
    }
    return message;
  },
};

const baseMsgExecuteIbcAutoForwardsResponse: object = {};

export const MsgExecuteIbcAutoForwardsResponse = {
  encode(
    _: MsgExecuteIbcAutoForwardsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgExecuteIbcAutoForwardsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgExecuteIbcAutoForwardsResponse,
    } as MsgExecuteIbcAutoForwardsResponse;
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

  fromJSON(_: any): MsgExecuteIbcAutoForwardsResponse {
    const message = {
      ...baseMsgExecuteIbcAutoForwardsResponse,
    } as MsgExecuteIbcAutoForwardsResponse;
    return message;
  },

  toJSON(_: MsgExecuteIbcAutoForwardsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgExecuteIbcAutoForwardsResponse>
  ): MsgExecuteIbcAutoForwardsResponse {
    const message = {
      ...baseMsgExecuteIbcAutoForwardsResponse,
    } as MsgExecuteIbcAutoForwardsResponse;
    return message;
  },
};

const baseMsgBatchSendToEthClaim: object = {
  event_nonce: 0,
  block_height: 0,
  batch_nonce: 0,
  token_contract: "",
  orchestrator: "",
};

export const MsgBatchSendToEthClaim = {
  encode(
    message: MsgBatchSendToEthClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    if (message.block_height !== 0) {
      writer.uint32(16).uint64(message.block_height);
    }
    if (message.batch_nonce !== 0) {
      writer.uint32(24).uint64(message.batch_nonce);
    }
    if (message.token_contract !== "") {
      writer.uint32(34).string(message.token_contract);
    }
    if (message.orchestrator !== "") {
      writer.uint32(42).string(message.orchestrator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBatchSendToEthClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBatchSendToEthClaim } as MsgBatchSendToEthClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.block_height = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.batch_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.token_contract = reader.string();
          break;
        case 5:
          message.orchestrator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBatchSendToEthClaim {
    const message = { ...baseMsgBatchSendToEthClaim } as MsgBatchSendToEthClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = Number(object.block_height);
    } else {
      message.block_height = 0;
    }
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = Number(object.batch_nonce);
    } else {
      message.batch_nonce = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = String(object.token_contract);
    } else {
      message.token_contract = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    return message;
  },

  toJSON(message: MsgBatchSendToEthClaim): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    message.block_height !== undefined &&
      (obj.block_height = message.block_height);
    message.batch_nonce !== undefined &&
      (obj.batch_nonce = message.batch_nonce);
    message.token_contract !== undefined &&
      (obj.token_contract = message.token_contract);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgBatchSendToEthClaim>
  ): MsgBatchSendToEthClaim {
    const message = { ...baseMsgBatchSendToEthClaim } as MsgBatchSendToEthClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = object.block_height;
    } else {
      message.block_height = 0;
    }
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = object.batch_nonce;
    } else {
      message.batch_nonce = 0;
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = object.token_contract;
    } else {
      message.token_contract = "";
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    return message;
  },
};

const baseMsgBatchSendToEthClaimResponse: object = {};

export const MsgBatchSendToEthClaimResponse = {
  encode(
    _: MsgBatchSendToEthClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgBatchSendToEthClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBatchSendToEthClaimResponse,
    } as MsgBatchSendToEthClaimResponse;
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

  fromJSON(_: any): MsgBatchSendToEthClaimResponse {
    const message = {
      ...baseMsgBatchSendToEthClaimResponse,
    } as MsgBatchSendToEthClaimResponse;
    return message;
  },

  toJSON(_: MsgBatchSendToEthClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgBatchSendToEthClaimResponse>
  ): MsgBatchSendToEthClaimResponse {
    const message = {
      ...baseMsgBatchSendToEthClaimResponse,
    } as MsgBatchSendToEthClaimResponse;
    return message;
  },
};

const baseMsgERC20DeployedClaim: object = {
  event_nonce: 0,
  block_height: 0,
  cosmos_denom: "",
  token_contract: "",
  name: "",
  symbol: "",
  decimals: 0,
  orchestrator: "",
};

export const MsgERC20DeployedClaim = {
  encode(
    message: MsgERC20DeployedClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    if (message.block_height !== 0) {
      writer.uint32(16).uint64(message.block_height);
    }
    if (message.cosmos_denom !== "") {
      writer.uint32(26).string(message.cosmos_denom);
    }
    if (message.token_contract !== "") {
      writer.uint32(34).string(message.token_contract);
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(50).string(message.symbol);
    }
    if (message.decimals !== 0) {
      writer.uint32(56).uint64(message.decimals);
    }
    if (message.orchestrator !== "") {
      writer.uint32(66).string(message.orchestrator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgERC20DeployedClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgERC20DeployedClaim } as MsgERC20DeployedClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.block_height = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.cosmos_denom = reader.string();
          break;
        case 4:
          message.token_contract = reader.string();
          break;
        case 5:
          message.name = reader.string();
          break;
        case 6:
          message.symbol = reader.string();
          break;
        case 7:
          message.decimals = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.orchestrator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgERC20DeployedClaim {
    const message = { ...baseMsgERC20DeployedClaim } as MsgERC20DeployedClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = Number(object.block_height);
    } else {
      message.block_height = 0;
    }
    if (object.cosmos_denom !== undefined && object.cosmos_denom !== null) {
      message.cosmos_denom = String(object.cosmos_denom);
    } else {
      message.cosmos_denom = "";
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = String(object.token_contract);
    } else {
      message.token_contract = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = Number(object.decimals);
    } else {
      message.decimals = 0;
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    return message;
  },

  toJSON(message: MsgERC20DeployedClaim): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    message.block_height !== undefined &&
      (obj.block_height = message.block_height);
    message.cosmos_denom !== undefined &&
      (obj.cosmos_denom = message.cosmos_denom);
    message.token_contract !== undefined &&
      (obj.token_contract = message.token_contract);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.decimals !== undefined && (obj.decimals = message.decimals);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgERC20DeployedClaim>
  ): MsgERC20DeployedClaim {
    const message = { ...baseMsgERC20DeployedClaim } as MsgERC20DeployedClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = object.block_height;
    } else {
      message.block_height = 0;
    }
    if (object.cosmos_denom !== undefined && object.cosmos_denom !== null) {
      message.cosmos_denom = object.cosmos_denom;
    } else {
      message.cosmos_denom = "";
    }
    if (object.token_contract !== undefined && object.token_contract !== null) {
      message.token_contract = object.token_contract;
    } else {
      message.token_contract = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = object.decimals;
    } else {
      message.decimals = 0;
    }
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    return message;
  },
};

const baseMsgERC20DeployedClaimResponse: object = {};

export const MsgERC20DeployedClaimResponse = {
  encode(
    _: MsgERC20DeployedClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgERC20DeployedClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgERC20DeployedClaimResponse,
    } as MsgERC20DeployedClaimResponse;
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

  fromJSON(_: any): MsgERC20DeployedClaimResponse {
    const message = {
      ...baseMsgERC20DeployedClaimResponse,
    } as MsgERC20DeployedClaimResponse;
    return message;
  },

  toJSON(_: MsgERC20DeployedClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgERC20DeployedClaimResponse>
  ): MsgERC20DeployedClaimResponse {
    const message = {
      ...baseMsgERC20DeployedClaimResponse,
    } as MsgERC20DeployedClaimResponse;
    return message;
  },
};

const baseMsgLogicCallExecutedClaim: object = {
  event_nonce: 0,
  block_height: 0,
  invalidation_nonce: 0,
  orchestrator: "",
};

export const MsgLogicCallExecutedClaim = {
  encode(
    message: MsgLogicCallExecutedClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    if (message.block_height !== 0) {
      writer.uint32(16).uint64(message.block_height);
    }
    if (message.invalidation_id.length !== 0) {
      writer.uint32(26).bytes(message.invalidation_id);
    }
    if (message.invalidation_nonce !== 0) {
      writer.uint32(32).uint64(message.invalidation_nonce);
    }
    if (message.orchestrator !== "") {
      writer.uint32(42).string(message.orchestrator);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgLogicCallExecutedClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgLogicCallExecutedClaim,
    } as MsgLogicCallExecutedClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.block_height = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.invalidation_id = reader.bytes();
          break;
        case 4:
          message.invalidation_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.orchestrator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLogicCallExecutedClaim {
    const message = {
      ...baseMsgLogicCallExecutedClaim,
    } as MsgLogicCallExecutedClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = Number(object.block_height);
    } else {
      message.block_height = 0;
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
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    return message;
  },

  toJSON(message: MsgLogicCallExecutedClaim): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    message.block_height !== undefined &&
      (obj.block_height = message.block_height);
    message.invalidation_id !== undefined &&
      (obj.invalidation_id = base64FromBytes(
        message.invalidation_id !== undefined
          ? message.invalidation_id
          : new Uint8Array()
      ));
    message.invalidation_nonce !== undefined &&
      (obj.invalidation_nonce = message.invalidation_nonce);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgLogicCallExecutedClaim>
  ): MsgLogicCallExecutedClaim {
    const message = {
      ...baseMsgLogicCallExecutedClaim,
    } as MsgLogicCallExecutedClaim;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = object.block_height;
    } else {
      message.block_height = 0;
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
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    return message;
  },
};

const baseMsgLogicCallExecutedClaimResponse: object = {};

export const MsgLogicCallExecutedClaimResponse = {
  encode(
    _: MsgLogicCallExecutedClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgLogicCallExecutedClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgLogicCallExecutedClaimResponse,
    } as MsgLogicCallExecutedClaimResponse;
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

  fromJSON(_: any): MsgLogicCallExecutedClaimResponse {
    const message = {
      ...baseMsgLogicCallExecutedClaimResponse,
    } as MsgLogicCallExecutedClaimResponse;
    return message;
  },

  toJSON(_: MsgLogicCallExecutedClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgLogicCallExecutedClaimResponse>
  ): MsgLogicCallExecutedClaimResponse {
    const message = {
      ...baseMsgLogicCallExecutedClaimResponse,
    } as MsgLogicCallExecutedClaimResponse;
    return message;
  },
};

const baseMsgValsetUpdatedClaim: object = {
  event_nonce: 0,
  valset_nonce: 0,
  block_height: 0,
  reward_amount: "",
  reward_token: "",
  orchestrator: "",
};

export const MsgValsetUpdatedClaim = {
  encode(
    message: MsgValsetUpdatedClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    if (message.valset_nonce !== 0) {
      writer.uint32(16).uint64(message.valset_nonce);
    }
    if (message.block_height !== 0) {
      writer.uint32(24).uint64(message.block_height);
    }
    for (const v of message.members) {
      BridgeValidator.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.reward_amount !== "") {
      writer.uint32(42).string(message.reward_amount);
    }
    if (message.reward_token !== "") {
      writer.uint32(50).string(message.reward_token);
    }
    if (message.orchestrator !== "") {
      writer.uint32(58).string(message.orchestrator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgValsetUpdatedClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgValsetUpdatedClaim } as MsgValsetUpdatedClaim;
    message.members = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.valset_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.block_height = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.members.push(BridgeValidator.decode(reader, reader.uint32()));
          break;
        case 5:
          message.reward_amount = reader.string();
          break;
        case 6:
          message.reward_token = reader.string();
          break;
        case 7:
          message.orchestrator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgValsetUpdatedClaim {
    const message = { ...baseMsgValsetUpdatedClaim } as MsgValsetUpdatedClaim;
    message.members = [];
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    if (object.valset_nonce !== undefined && object.valset_nonce !== null) {
      message.valset_nonce = Number(object.valset_nonce);
    } else {
      message.valset_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = Number(object.block_height);
    } else {
      message.block_height = 0;
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(BridgeValidator.fromJSON(e));
      }
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
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = String(object.orchestrator);
    } else {
      message.orchestrator = "";
    }
    return message;
  },

  toJSON(message: MsgValsetUpdatedClaim): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    message.valset_nonce !== undefined &&
      (obj.valset_nonce = message.valset_nonce);
    message.block_height !== undefined &&
      (obj.block_height = message.block_height);
    if (message.members) {
      obj.members = message.members.map((e) =>
        e ? BridgeValidator.toJSON(e) : undefined
      );
    } else {
      obj.members = [];
    }
    message.reward_amount !== undefined &&
      (obj.reward_amount = message.reward_amount);
    message.reward_token !== undefined &&
      (obj.reward_token = message.reward_token);
    message.orchestrator !== undefined &&
      (obj.orchestrator = message.orchestrator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgValsetUpdatedClaim>
  ): MsgValsetUpdatedClaim {
    const message = { ...baseMsgValsetUpdatedClaim } as MsgValsetUpdatedClaim;
    message.members = [];
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    if (object.valset_nonce !== undefined && object.valset_nonce !== null) {
      message.valset_nonce = object.valset_nonce;
    } else {
      message.valset_nonce = 0;
    }
    if (object.block_height !== undefined && object.block_height !== null) {
      message.block_height = object.block_height;
    } else {
      message.block_height = 0;
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(BridgeValidator.fromPartial(e));
      }
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
    if (object.orchestrator !== undefined && object.orchestrator !== null) {
      message.orchestrator = object.orchestrator;
    } else {
      message.orchestrator = "";
    }
    return message;
  },
};

const baseMsgValsetUpdatedClaimResponse: object = {};

export const MsgValsetUpdatedClaimResponse = {
  encode(
    _: MsgValsetUpdatedClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgValsetUpdatedClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgValsetUpdatedClaimResponse,
    } as MsgValsetUpdatedClaimResponse;
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

  fromJSON(_: any): MsgValsetUpdatedClaimResponse {
    const message = {
      ...baseMsgValsetUpdatedClaimResponse,
    } as MsgValsetUpdatedClaimResponse;
    return message;
  },

  toJSON(_: MsgValsetUpdatedClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgValsetUpdatedClaimResponse>
  ): MsgValsetUpdatedClaimResponse {
    const message = {
      ...baseMsgValsetUpdatedClaimResponse,
    } as MsgValsetUpdatedClaimResponse;
    return message;
  },
};

const baseMsgCancelSendToEth: object = { transaction_id: 0, sender: "" };

export const MsgCancelSendToEth = {
  encode(
    message: MsgCancelSendToEth,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.transaction_id !== 0) {
      writer.uint32(8).uint64(message.transaction_id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelSendToEth {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelSendToEth } as MsgCancelSendToEth;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transaction_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelSendToEth {
    const message = { ...baseMsgCancelSendToEth } as MsgCancelSendToEth;
    if (object.transaction_id !== undefined && object.transaction_id !== null) {
      message.transaction_id = Number(object.transaction_id);
    } else {
      message.transaction_id = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: MsgCancelSendToEth): unknown {
    const obj: any = {};
    message.transaction_id !== undefined &&
      (obj.transaction_id = message.transaction_id);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelSendToEth>): MsgCancelSendToEth {
    const message = { ...baseMsgCancelSendToEth } as MsgCancelSendToEth;
    if (object.transaction_id !== undefined && object.transaction_id !== null) {
      message.transaction_id = object.transaction_id;
    } else {
      message.transaction_id = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseMsgCancelSendToEthResponse: object = {};

export const MsgCancelSendToEthResponse = {
  encode(
    _: MsgCancelSendToEthResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelSendToEthResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelSendToEthResponse,
    } as MsgCancelSendToEthResponse;
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

  fromJSON(_: any): MsgCancelSendToEthResponse {
    const message = {
      ...baseMsgCancelSendToEthResponse,
    } as MsgCancelSendToEthResponse;
    return message;
  },

  toJSON(_: MsgCancelSendToEthResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelSendToEthResponse>
  ): MsgCancelSendToEthResponse {
    const message = {
      ...baseMsgCancelSendToEthResponse,
    } as MsgCancelSendToEthResponse;
    return message;
  },
};

const baseMsgSubmitBadSignatureEvidence: object = { signature: "", sender: "" };

export const MsgSubmitBadSignatureEvidence = {
  encode(
    message: MsgSubmitBadSignatureEvidence,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.subject !== undefined) {
      Any.encode(message.subject, writer.uint32(10).fork()).ldelim();
    }
    if (message.signature !== "") {
      writer.uint32(18).string(message.signature);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitBadSignatureEvidence {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitBadSignatureEvidence,
    } as MsgSubmitBadSignatureEvidence;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = Any.decode(reader, reader.uint32());
          break;
        case 2:
          message.signature = reader.string();
          break;
        case 3:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitBadSignatureEvidence {
    const message = {
      ...baseMsgSubmitBadSignatureEvidence,
    } as MsgSubmitBadSignatureEvidence;
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = Any.fromJSON(object.subject);
    } else {
      message.subject = undefined;
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitBadSignatureEvidence): unknown {
    const obj: any = {};
    message.subject !== undefined &&
      (obj.subject = message.subject ? Any.toJSON(message.subject) : undefined);
    message.signature !== undefined && (obj.signature = message.signature);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitBadSignatureEvidence>
  ): MsgSubmitBadSignatureEvidence {
    const message = {
      ...baseMsgSubmitBadSignatureEvidence,
    } as MsgSubmitBadSignatureEvidence;
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = Any.fromPartial(object.subject);
    } else {
      message.subject = undefined;
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseMsgSubmitBadSignatureEvidenceResponse: object = {};

export const MsgSubmitBadSignatureEvidenceResponse = {
  encode(
    _: MsgSubmitBadSignatureEvidenceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitBadSignatureEvidenceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitBadSignatureEvidenceResponse,
    } as MsgSubmitBadSignatureEvidenceResponse;
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

  fromJSON(_: any): MsgSubmitBadSignatureEvidenceResponse {
    const message = {
      ...baseMsgSubmitBadSignatureEvidenceResponse,
    } as MsgSubmitBadSignatureEvidenceResponse;
    return message;
  },

  toJSON(_: MsgSubmitBadSignatureEvidenceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitBadSignatureEvidenceResponse>
  ): MsgSubmitBadSignatureEvidenceResponse {
    const message = {
      ...baseMsgSubmitBadSignatureEvidenceResponse,
    } as MsgSubmitBadSignatureEvidenceResponse;
    return message;
  },
};

const baseEventSetOperatorAddress: object = { message: "", address: "" };

export const EventSetOperatorAddress = {
  encode(
    message: EventSetOperatorAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventSetOperatorAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventSetOperatorAddress,
    } as EventSetOperatorAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventSetOperatorAddress {
    const message = {
      ...baseEventSetOperatorAddress,
    } as EventSetOperatorAddress;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: EventSetOperatorAddress): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventSetOperatorAddress>
  ): EventSetOperatorAddress {
    const message = {
      ...baseEventSetOperatorAddress,
    } as EventSetOperatorAddress;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseEventValsetConfirmKey: object = { message: "", key: "" };

export const EventValsetConfirmKey = {
  encode(
    message: EventValsetConfirmKey,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventValsetConfirmKey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventValsetConfirmKey } as EventValsetConfirmKey;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventValsetConfirmKey {
    const message = { ...baseEventValsetConfirmKey } as EventValsetConfirmKey;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    return message;
  },

  toJSON(message: EventValsetConfirmKey): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventValsetConfirmKey>
  ): EventValsetConfirmKey {
    const message = { ...baseEventValsetConfirmKey } as EventValsetConfirmKey;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    return message;
  },
};

const baseEventBatchCreated: object = { message: "", batch_nonce: "" };

export const EventBatchCreated = {
  encode(message: EventBatchCreated, writer: Writer = Writer.create()): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.batch_nonce !== "") {
      writer.uint32(18).string(message.batch_nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventBatchCreated {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventBatchCreated } as EventBatchCreated;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.batch_nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventBatchCreated {
    const message = { ...baseEventBatchCreated } as EventBatchCreated;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = String(object.batch_nonce);
    } else {
      message.batch_nonce = "";
    }
    return message;
  },

  toJSON(message: EventBatchCreated): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.batch_nonce !== undefined &&
      (obj.batch_nonce = message.batch_nonce);
    return obj;
  },

  fromPartial(object: DeepPartial<EventBatchCreated>): EventBatchCreated {
    const message = { ...baseEventBatchCreated } as EventBatchCreated;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.batch_nonce !== undefined && object.batch_nonce !== null) {
      message.batch_nonce = object.batch_nonce;
    } else {
      message.batch_nonce = "";
    }
    return message;
  },
};

const baseEventBatchConfirmKey: object = { message: "", batch_confirm_key: "" };

export const EventBatchConfirmKey = {
  encode(
    message: EventBatchConfirmKey,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.batch_confirm_key !== "") {
      writer.uint32(18).string(message.batch_confirm_key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventBatchConfirmKey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventBatchConfirmKey } as EventBatchConfirmKey;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.batch_confirm_key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventBatchConfirmKey {
    const message = { ...baseEventBatchConfirmKey } as EventBatchConfirmKey;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (
      object.batch_confirm_key !== undefined &&
      object.batch_confirm_key !== null
    ) {
      message.batch_confirm_key = String(object.batch_confirm_key);
    } else {
      message.batch_confirm_key = "";
    }
    return message;
  },

  toJSON(message: EventBatchConfirmKey): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.batch_confirm_key !== undefined &&
      (obj.batch_confirm_key = message.batch_confirm_key);
    return obj;
  },

  fromPartial(object: DeepPartial<EventBatchConfirmKey>): EventBatchConfirmKey {
    const message = { ...baseEventBatchConfirmKey } as EventBatchConfirmKey;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (
      object.batch_confirm_key !== undefined &&
      object.batch_confirm_key !== null
    ) {
      message.batch_confirm_key = object.batch_confirm_key;
    } else {
      message.batch_confirm_key = "";
    }
    return message;
  },
};

const baseEventBatchSendToEthClaim: object = { nonce: "" };

export const EventBatchSendToEthClaim = {
  encode(
    message: EventBatchSendToEthClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== "") {
      writer.uint32(10).string(message.nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventBatchSendToEthClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventBatchSendToEthClaim,
    } as EventBatchSendToEthClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventBatchSendToEthClaim {
    const message = {
      ...baseEventBatchSendToEthClaim,
    } as EventBatchSendToEthClaim;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventBatchSendToEthClaim): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventBatchSendToEthClaim>
  ): EventBatchSendToEthClaim {
    const message = {
      ...baseEventBatchSendToEthClaim,
    } as EventBatchSendToEthClaim;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventClaim: object = {
  message: "",
  claim_hash: "",
  attestation_id: "",
};

export const EventClaim = {
  encode(message: EventClaim, writer: Writer = Writer.create()): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.claim_hash !== "") {
      writer.uint32(18).string(message.claim_hash);
    }
    if (message.attestation_id !== "") {
      writer.uint32(26).string(message.attestation_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventClaim } as EventClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.claim_hash = reader.string();
          break;
        case 3:
          message.attestation_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventClaim {
    const message = { ...baseEventClaim } as EventClaim;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.claim_hash !== undefined && object.claim_hash !== null) {
      message.claim_hash = String(object.claim_hash);
    } else {
      message.claim_hash = "";
    }
    if (object.attestation_id !== undefined && object.attestation_id !== null) {
      message.attestation_id = String(object.attestation_id);
    } else {
      message.attestation_id = "";
    }
    return message;
  },

  toJSON(message: EventClaim): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.claim_hash !== undefined && (obj.claim_hash = message.claim_hash);
    message.attestation_id !== undefined &&
      (obj.attestation_id = message.attestation_id);
    return obj;
  },

  fromPartial(object: DeepPartial<EventClaim>): EventClaim {
    const message = { ...baseEventClaim } as EventClaim;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.claim_hash !== undefined && object.claim_hash !== null) {
      message.claim_hash = object.claim_hash;
    } else {
      message.claim_hash = "";
    }
    if (object.attestation_id !== undefined && object.attestation_id !== null) {
      message.attestation_id = object.attestation_id;
    } else {
      message.attestation_id = "";
    }
    return message;
  },
};

const baseEventBadSignatureEvidence: object = {
  message: "",
  bad_eth_signature: "",
  bad_eth_signature_subject: "",
};

export const EventBadSignatureEvidence = {
  encode(
    message: EventBadSignatureEvidence,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.bad_eth_signature !== "") {
      writer.uint32(18).string(message.bad_eth_signature);
    }
    if (message.bad_eth_signature_subject !== "") {
      writer.uint32(26).string(message.bad_eth_signature_subject);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventBadSignatureEvidence {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventBadSignatureEvidence,
    } as EventBadSignatureEvidence;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.bad_eth_signature = reader.string();
          break;
        case 3:
          message.bad_eth_signature_subject = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventBadSignatureEvidence {
    const message = {
      ...baseEventBadSignatureEvidence,
    } as EventBadSignatureEvidence;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (
      object.bad_eth_signature !== undefined &&
      object.bad_eth_signature !== null
    ) {
      message.bad_eth_signature = String(object.bad_eth_signature);
    } else {
      message.bad_eth_signature = "";
    }
    if (
      object.bad_eth_signature_subject !== undefined &&
      object.bad_eth_signature_subject !== null
    ) {
      message.bad_eth_signature_subject = String(
        object.bad_eth_signature_subject
      );
    } else {
      message.bad_eth_signature_subject = "";
    }
    return message;
  },

  toJSON(message: EventBadSignatureEvidence): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.bad_eth_signature !== undefined &&
      (obj.bad_eth_signature = message.bad_eth_signature);
    message.bad_eth_signature_subject !== undefined &&
      (obj.bad_eth_signature_subject = message.bad_eth_signature_subject);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventBadSignatureEvidence>
  ): EventBadSignatureEvidence {
    const message = {
      ...baseEventBadSignatureEvidence,
    } as EventBadSignatureEvidence;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (
      object.bad_eth_signature !== undefined &&
      object.bad_eth_signature !== null
    ) {
      message.bad_eth_signature = object.bad_eth_signature;
    } else {
      message.bad_eth_signature = "";
    }
    if (
      object.bad_eth_signature_subject !== undefined &&
      object.bad_eth_signature_subject !== null
    ) {
      message.bad_eth_signature_subject = object.bad_eth_signature_subject;
    } else {
      message.bad_eth_signature_subject = "";
    }
    return message;
  },
};

const baseEventERC20DeployedClaim: object = { token: "", nonce: "" };

export const EventERC20DeployedClaim = {
  encode(
    message: EventERC20DeployedClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.nonce !== "") {
      writer.uint32(18).string(message.nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventERC20DeployedClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventERC20DeployedClaim,
    } as EventERC20DeployedClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventERC20DeployedClaim {
    const message = {
      ...baseEventERC20DeployedClaim,
    } as EventERC20DeployedClaim;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventERC20DeployedClaim): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventERC20DeployedClaim>
  ): EventERC20DeployedClaim {
    const message = {
      ...baseEventERC20DeployedClaim,
    } as EventERC20DeployedClaim;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventValsetUpdatedClaim: object = { nonce: "" };

export const EventValsetUpdatedClaim = {
  encode(
    message: EventValsetUpdatedClaim,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== "") {
      writer.uint32(10).string(message.nonce);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventValsetUpdatedClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventValsetUpdatedClaim,
    } as EventValsetUpdatedClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventValsetUpdatedClaim {
    const message = {
      ...baseEventValsetUpdatedClaim,
    } as EventValsetUpdatedClaim;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventValsetUpdatedClaim): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventValsetUpdatedClaim>
  ): EventValsetUpdatedClaim {
    const message = {
      ...baseEventValsetUpdatedClaim,
    } as EventValsetUpdatedClaim;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventMultisigUpdateRequest: object = {
  bridge_contract: "",
  bridge_chain_id: "",
  multisig_id: "",
  nonce: "",
};

export const EventMultisigUpdateRequest = {
  encode(
    message: EventMultisigUpdateRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bridge_contract !== "") {
      writer.uint32(10).string(message.bridge_contract);
    }
    if (message.bridge_chain_id !== "") {
      writer.uint32(18).string(message.bridge_chain_id);
    }
    if (message.multisig_id !== "") {
      writer.uint32(26).string(message.multisig_id);
    }
    if (message.nonce !== "") {
      writer.uint32(34).string(message.nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventMultisigUpdateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventMultisigUpdateRequest,
    } as EventMultisigUpdateRequest;
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
          message.multisig_id = reader.string();
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

  fromJSON(object: any): EventMultisigUpdateRequest {
    const message = {
      ...baseEventMultisigUpdateRequest,
    } as EventMultisigUpdateRequest;
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
    if (object.multisig_id !== undefined && object.multisig_id !== null) {
      message.multisig_id = String(object.multisig_id);
    } else {
      message.multisig_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = String(object.nonce);
    } else {
      message.nonce = "";
    }
    return message;
  },

  toJSON(message: EventMultisigUpdateRequest): unknown {
    const obj: any = {};
    message.bridge_contract !== undefined &&
      (obj.bridge_contract = message.bridge_contract);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.multisig_id !== undefined &&
      (obj.multisig_id = message.multisig_id);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventMultisigUpdateRequest>
  ): EventMultisigUpdateRequest {
    const message = {
      ...baseEventMultisigUpdateRequest,
    } as EventMultisigUpdateRequest;
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
    if (object.multisig_id !== undefined && object.multisig_id !== null) {
      message.multisig_id = object.multisig_id;
    } else {
      message.multisig_id = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = "";
    }
    return message;
  },
};

const baseEventOutgoingLogicCallCanceled: object = {
  logic_call_invalidation_id: "",
  logic_call_invalidation_nonce: "",
};

export const EventOutgoingLogicCallCanceled = {
  encode(
    message: EventOutgoingLogicCallCanceled,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.logic_call_invalidation_id !== "") {
      writer.uint32(10).string(message.logic_call_invalidation_id);
    }
    if (message.logic_call_invalidation_nonce !== "") {
      writer.uint32(18).string(message.logic_call_invalidation_nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventOutgoingLogicCallCanceled {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventOutgoingLogicCallCanceled,
    } as EventOutgoingLogicCallCanceled;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.logic_call_invalidation_id = reader.string();
          break;
        case 2:
          message.logic_call_invalidation_nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventOutgoingLogicCallCanceled {
    const message = {
      ...baseEventOutgoingLogicCallCanceled,
    } as EventOutgoingLogicCallCanceled;
    if (
      object.logic_call_invalidation_id !== undefined &&
      object.logic_call_invalidation_id !== null
    ) {
      message.logic_call_invalidation_id = String(
        object.logic_call_invalidation_id
      );
    } else {
      message.logic_call_invalidation_id = "";
    }
    if (
      object.logic_call_invalidation_nonce !== undefined &&
      object.logic_call_invalidation_nonce !== null
    ) {
      message.logic_call_invalidation_nonce = String(
        object.logic_call_invalidation_nonce
      );
    } else {
      message.logic_call_invalidation_nonce = "";
    }
    return message;
  },

  toJSON(message: EventOutgoingLogicCallCanceled): unknown {
    const obj: any = {};
    message.logic_call_invalidation_id !== undefined &&
      (obj.logic_call_invalidation_id = message.logic_call_invalidation_id);
    message.logic_call_invalidation_nonce !== undefined &&
      (obj.logic_call_invalidation_nonce =
        message.logic_call_invalidation_nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventOutgoingLogicCallCanceled>
  ): EventOutgoingLogicCallCanceled {
    const message = {
      ...baseEventOutgoingLogicCallCanceled,
    } as EventOutgoingLogicCallCanceled;
    if (
      object.logic_call_invalidation_id !== undefined &&
      object.logic_call_invalidation_id !== null
    ) {
      message.logic_call_invalidation_id = object.logic_call_invalidation_id;
    } else {
      message.logic_call_invalidation_id = "";
    }
    if (
      object.logic_call_invalidation_nonce !== undefined &&
      object.logic_call_invalidation_nonce !== null
    ) {
      message.logic_call_invalidation_nonce =
        object.logic_call_invalidation_nonce;
    } else {
      message.logic_call_invalidation_nonce = "";
    }
    return message;
  },
};

const baseEventSignatureSlashing: object = { type: "", address: "" };

export const EventSignatureSlashing = {
  encode(
    message: EventSignatureSlashing,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventSignatureSlashing {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventSignatureSlashing } as EventSignatureSlashing;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventSignatureSlashing {
    const message = { ...baseEventSignatureSlashing } as EventSignatureSlashing;
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: EventSignatureSlashing): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventSignatureSlashing>
  ): EventSignatureSlashing {
    const message = { ...baseEventSignatureSlashing } as EventSignatureSlashing;
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseEventOutgoingTxId: object = { message: "", tx_id: "" };

export const EventOutgoingTxId = {
  encode(message: EventOutgoingTxId, writer: Writer = Writer.create()): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.tx_id !== "") {
      writer.uint32(18).string(message.tx_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventOutgoingTxId {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventOutgoingTxId } as EventOutgoingTxId;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 2:
          message.tx_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventOutgoingTxId {
    const message = { ...baseEventOutgoingTxId } as EventOutgoingTxId;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.tx_id !== undefined && object.tx_id !== null) {
      message.tx_id = String(object.tx_id);
    } else {
      message.tx_id = "";
    }
    return message;
  },

  toJSON(message: EventOutgoingTxId): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.tx_id !== undefined && (obj.tx_id = message.tx_id);
    return obj;
  },

  fromPartial(object: DeepPartial<EventOutgoingTxId>): EventOutgoingTxId {
    const message = { ...baseEventOutgoingTxId } as EventOutgoingTxId;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.tx_id !== undefined && object.tx_id !== null) {
      message.tx_id = object.tx_id;
    } else {
      message.tx_id = "";
    }
    return message;
  },
};

/** Msg defines the state transitions possible within gravity */
export interface Msg {
  ValsetConfirm(request: MsgValsetConfirm): Promise<MsgValsetConfirmResponse>;
  SendToEth(request: MsgSendToEth): Promise<MsgSendToEthResponse>;
  RequestBatch(request: MsgRequestBatch): Promise<MsgRequestBatchResponse>;
  ConfirmBatch(request: MsgConfirmBatch): Promise<MsgConfirmBatchResponse>;
  ConfirmLogicCall(
    request: MsgConfirmLogicCall
  ): Promise<MsgConfirmLogicCallResponse>;
  SendToCosmosClaim(
    request: MsgSendToCosmosClaim
  ): Promise<MsgSendToCosmosClaimResponse>;
  ExecuteIbcAutoForwards(
    request: MsgExecuteIbcAutoForwards
  ): Promise<MsgExecuteIbcAutoForwardsResponse>;
  BatchSendToEthClaim(
    request: MsgBatchSendToEthClaim
  ): Promise<MsgBatchSendToEthClaimResponse>;
  ValsetUpdateClaim(
    request: MsgValsetUpdatedClaim
  ): Promise<MsgValsetUpdatedClaimResponse>;
  ERC20DeployedClaim(
    request: MsgERC20DeployedClaim
  ): Promise<MsgERC20DeployedClaimResponse>;
  LogicCallExecutedClaim(
    request: MsgLogicCallExecutedClaim
  ): Promise<MsgLogicCallExecutedClaimResponse>;
  SetOrchestratorAddress(
    request: MsgSetOrchestratorAddress
  ): Promise<MsgSetOrchestratorAddressResponse>;
  CancelSendToEth(
    request: MsgCancelSendToEth
  ): Promise<MsgCancelSendToEthResponse>;
  SubmitBadSignatureEvidence(
    request: MsgSubmitBadSignatureEvidence
  ): Promise<MsgSubmitBadSignatureEvidenceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ValsetConfirm(request: MsgValsetConfirm): Promise<MsgValsetConfirmResponse> {
    const data = MsgValsetConfirm.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "ValsetConfirm", data);
    return promise.then((data) =>
      MsgValsetConfirmResponse.decode(new Reader(data))
    );
  }

  SendToEth(request: MsgSendToEth): Promise<MsgSendToEthResponse> {
    const data = MsgSendToEth.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "SendToEth", data);
    return promise.then((data) =>
      MsgSendToEthResponse.decode(new Reader(data))
    );
  }

  RequestBatch(request: MsgRequestBatch): Promise<MsgRequestBatchResponse> {
    const data = MsgRequestBatch.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "RequestBatch", data);
    return promise.then((data) =>
      MsgRequestBatchResponse.decode(new Reader(data))
    );
  }

  ConfirmBatch(request: MsgConfirmBatch): Promise<MsgConfirmBatchResponse> {
    const data = MsgConfirmBatch.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "ConfirmBatch", data);
    return promise.then((data) =>
      MsgConfirmBatchResponse.decode(new Reader(data))
    );
  }

  ConfirmLogicCall(
    request: MsgConfirmLogicCall
  ): Promise<MsgConfirmLogicCallResponse> {
    const data = MsgConfirmLogicCall.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "ConfirmLogicCall", data);
    return promise.then((data) =>
      MsgConfirmLogicCallResponse.decode(new Reader(data))
    );
  }

  SendToCosmosClaim(
    request: MsgSendToCosmosClaim
  ): Promise<MsgSendToCosmosClaimResponse> {
    const data = MsgSendToCosmosClaim.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "SendToCosmosClaim", data);
    return promise.then((data) =>
      MsgSendToCosmosClaimResponse.decode(new Reader(data))
    );
  }

  ExecuteIbcAutoForwards(
    request: MsgExecuteIbcAutoForwards
  ): Promise<MsgExecuteIbcAutoForwardsResponse> {
    const data = MsgExecuteIbcAutoForwards.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Msg",
      "ExecuteIbcAutoForwards",
      data
    );
    return promise.then((data) =>
      MsgExecuteIbcAutoForwardsResponse.decode(new Reader(data))
    );
  }

  BatchSendToEthClaim(
    request: MsgBatchSendToEthClaim
  ): Promise<MsgBatchSendToEthClaimResponse> {
    const data = MsgBatchSendToEthClaim.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Msg",
      "BatchSendToEthClaim",
      data
    );
    return promise.then((data) =>
      MsgBatchSendToEthClaimResponse.decode(new Reader(data))
    );
  }

  ValsetUpdateClaim(
    request: MsgValsetUpdatedClaim
  ): Promise<MsgValsetUpdatedClaimResponse> {
    const data = MsgValsetUpdatedClaim.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "ValsetUpdateClaim", data);
    return promise.then((data) =>
      MsgValsetUpdatedClaimResponse.decode(new Reader(data))
    );
  }

  ERC20DeployedClaim(
    request: MsgERC20DeployedClaim
  ): Promise<MsgERC20DeployedClaimResponse> {
    const data = MsgERC20DeployedClaim.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "ERC20DeployedClaim", data);
    return promise.then((data) =>
      MsgERC20DeployedClaimResponse.decode(new Reader(data))
    );
  }

  LogicCallExecutedClaim(
    request: MsgLogicCallExecutedClaim
  ): Promise<MsgLogicCallExecutedClaimResponse> {
    const data = MsgLogicCallExecutedClaim.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Msg",
      "LogicCallExecutedClaim",
      data
    );
    return promise.then((data) =>
      MsgLogicCallExecutedClaimResponse.decode(new Reader(data))
    );
  }

  SetOrchestratorAddress(
    request: MsgSetOrchestratorAddress
  ): Promise<MsgSetOrchestratorAddressResponse> {
    const data = MsgSetOrchestratorAddress.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Msg",
      "SetOrchestratorAddress",
      data
    );
    return promise.then((data) =>
      MsgSetOrchestratorAddressResponse.decode(new Reader(data))
    );
  }

  CancelSendToEth(
    request: MsgCancelSendToEth
  ): Promise<MsgCancelSendToEthResponse> {
    const data = MsgCancelSendToEth.encode(request).finish();
    const promise = this.rpc.request("gravity.Msg", "CancelSendToEth", data);
    return promise.then((data) =>
      MsgCancelSendToEthResponse.decode(new Reader(data))
    );
  }

  SubmitBadSignatureEvidence(
    request: MsgSubmitBadSignatureEvidence
  ): Promise<MsgSubmitBadSignatureEvidenceResponse> {
    const data = MsgSubmitBadSignatureEvidence.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Msg",
      "SubmitBadSignatureEvidence",
      data
    );
    return promise.then((data) =>
      MsgSubmitBadSignatureEvidenceResponse.decode(new Reader(data))
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
