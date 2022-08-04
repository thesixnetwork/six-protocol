/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Valset, ERC20ToDenom } from "../../thesixnetwork/six-protocol/gravity/v1/types";
import {
  MsgValsetConfirm,
  MsgConfirmBatch,
  MsgConfirmLogicCall,
  MsgSetOrchestratorAddress,
} from "../../thesixnetwork/six-protocol/gravity/v1/msgs";
import {
  OutgoingTxBatch,
  OutgoingLogicCall,
  OutgoingTransferTx,
} from "../../thesixnetwork/six-protocol/gravity/v1/batch";
import { Attestation } from "../../thesixnetwork/six-protocol/gravity/v1/attestation";

export const protobufPackage = "gravity.v1";

/**
 * The slashing fractions for the various gravity related slashing conditions. The first three
 * refer to not submitting a particular message, the third for submitting a different claim
 * for the same Ethereum event
 *
 * unbond_slashing_valsets_window
 *
 * The unbond slashing valsets window is used to determine how many blocks after starting to unbond
 * a validator needs to continue signing blocks. The goal of this paramater is that when a validator leaves
 * the set, if their leaving creates enough change in the validator set to justify an update they will sign
 * a validator set update for the Ethereum bridge that does not include themselves. Allowing us to remove them
 * from the Ethereum bridge and replace them with the new set gracefully.
 *
 * valset_reward
 *
 * These parameters allow for the bridge oracle to resolve a fork on the Ethereum chain without halting
 * the chain. Once set reset bridge state will roll back events to the nonce provided in reset_bridge_nonce
 * if and only if those events have not yet been observed (executed on the Cosmos chain). This allows for easy
 * handling of cases where for example an Ethereum hardfork has occured and more than 1/3 of the vlaidtor set
 * disagrees with the rest. Normally this would require a chain halt, manual genesis editing and restar to resolve
 * with this feature a governance proposal can be used instead
 *
 * bridge_active
 *
 * This boolean flag can be used by governance to temporarily halt the bridge due to a vulnerability or other issue
 * In this context halting the bridge means prevent the execution of any oracle events from Ethereum and preventing
 * the creation of new batches that may be relayed to Ethereum.
 * This does not prevent the creation of validator sets
 * or slashing for not submitting validator set signatures as either of these might allow key signers to leave the validator
 * set and steal funds on Ethereum without consequence.
 * The practical outcome of this flag being set to 'false' is that deposits from Ethereum will not show up and withdraws from
 * Cosmos will not execute on Ethereum.
 */
export interface Params {
  gravity_id: string;
  contract_source_hash: string;
  bridge_ethereum_address: string;
  bridge_chain_id: number;
  signed_valsets_window: number;
  signed_batches_window: number;
  signed_logic_calls_window: number;
  target_batch_timeout: number;
  average_block_time: number;
  average_ethereum_block_time: number;
  slash_fraction_valset: Uint8Array;
  slash_fraction_batch: Uint8Array;
  slash_fraction_logic_call: Uint8Array;
  unbond_slashing_valsets_window: number;
  slash_fraction_bad_eth_signature: Uint8Array;
  valset_reward: Coin | undefined;
  bridge_active: boolean;
  /**
   * addresses on this blacklist are forbidden from depositing or withdrawing
   * from Ethereum to the bridge
   */
  ethereum_blacklist: string[];
}

/** GenesisState struct, containing all persistant data required by the Gravity module */
export interface GenesisState {
  params: Params | undefined;
  gravity_nonces: GravityNonces | undefined;
  valsets: Valset[];
  valset_confirms: MsgValsetConfirm[];
  batches: OutgoingTxBatch[];
  batch_confirms: MsgConfirmBatch[];
  logic_calls: OutgoingLogicCall[];
  logic_call_confirms: MsgConfirmLogicCall[];
  attestations: Attestation[];
  delegate_keys: MsgSetOrchestratorAddress[];
  erc20_to_denoms: ERC20ToDenom[];
  unbatched_transfers: OutgoingTransferTx[];
}

/** GravityCounters contains the many noces and counters required to maintain the bridge state in the genesis */
export interface GravityNonces {
  /** the nonce of the last generated validator set */
  latest_valset_nonce: number;
  /** the last observed Gravity.sol contract event nonce */
  last_observed_nonce: number;
  /** the last valset nonce we have slashed, to prevent double slashing */
  last_slashed_valset_nonce: number;
  /**
   * the last batch Cosmos chain block that batch slashing has completed for
   * there is an individual batch nonce for each token type so this removes
   * the need to store them all
   */
  last_slashed_batch_block: number;
  /** the last cosmos block that logic call slashing has completed for */
  last_slashed_logic_call_block: number;
  /**
   * the last transaction id from the Gravity TX pool, this prevents ID
   * duplication during chain upgrades
   */
  last_tx_pool_id: number;
  /**
   * the last batch id from the Gravity batch pool, this prevents ID duplication
   * during chain upgrades
   */
  last_batch_id: number;
}

const baseParams: object = {
  gravity_id: "",
  contract_source_hash: "",
  bridge_ethereum_address: "",
  bridge_chain_id: 0,
  signed_valsets_window: 0,
  signed_batches_window: 0,
  signed_logic_calls_window: 0,
  target_batch_timeout: 0,
  average_block_time: 0,
  average_ethereum_block_time: 0,
  unbond_slashing_valsets_window: 0,
  bridge_active: false,
  ethereum_blacklist: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.gravity_id !== "") {
      writer.uint32(10).string(message.gravity_id);
    }
    if (message.contract_source_hash !== "") {
      writer.uint32(18).string(message.contract_source_hash);
    }
    if (message.bridge_ethereum_address !== "") {
      writer.uint32(34).string(message.bridge_ethereum_address);
    }
    if (message.bridge_chain_id !== 0) {
      writer.uint32(40).uint64(message.bridge_chain_id);
    }
    if (message.signed_valsets_window !== 0) {
      writer.uint32(48).uint64(message.signed_valsets_window);
    }
    if (message.signed_batches_window !== 0) {
      writer.uint32(56).uint64(message.signed_batches_window);
    }
    if (message.signed_logic_calls_window !== 0) {
      writer.uint32(64).uint64(message.signed_logic_calls_window);
    }
    if (message.target_batch_timeout !== 0) {
      writer.uint32(72).uint64(message.target_batch_timeout);
    }
    if (message.average_block_time !== 0) {
      writer.uint32(80).uint64(message.average_block_time);
    }
    if (message.average_ethereum_block_time !== 0) {
      writer.uint32(88).uint64(message.average_ethereum_block_time);
    }
    if (message.slash_fraction_valset.length !== 0) {
      writer.uint32(98).bytes(message.slash_fraction_valset);
    }
    if (message.slash_fraction_batch.length !== 0) {
      writer.uint32(106).bytes(message.slash_fraction_batch);
    }
    if (message.slash_fraction_logic_call.length !== 0) {
      writer.uint32(114).bytes(message.slash_fraction_logic_call);
    }
    if (message.unbond_slashing_valsets_window !== 0) {
      writer.uint32(120).uint64(message.unbond_slashing_valsets_window);
    }
    if (message.slash_fraction_bad_eth_signature.length !== 0) {
      writer.uint32(130).bytes(message.slash_fraction_bad_eth_signature);
    }
    if (message.valset_reward !== undefined) {
      Coin.encode(message.valset_reward, writer.uint32(138).fork()).ldelim();
    }
    if (message.bridge_active === true) {
      writer.uint32(144).bool(message.bridge_active);
    }
    for (const v of message.ethereum_blacklist) {
      writer.uint32(154).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.ethereum_blacklist = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gravity_id = reader.string();
          break;
        case 2:
          message.contract_source_hash = reader.string();
          break;
        case 4:
          message.bridge_ethereum_address = reader.string();
          break;
        case 5:
          message.bridge_chain_id = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.signed_valsets_window = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.signed_batches_window = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.signed_logic_calls_window = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 9:
          message.target_batch_timeout = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.average_block_time = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.average_ethereum_block_time = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 12:
          message.slash_fraction_valset = reader.bytes();
          break;
        case 13:
          message.slash_fraction_batch = reader.bytes();
          break;
        case 14:
          message.slash_fraction_logic_call = reader.bytes();
          break;
        case 15:
          message.unbond_slashing_valsets_window = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 16:
          message.slash_fraction_bad_eth_signature = reader.bytes();
          break;
        case 17:
          message.valset_reward = Coin.decode(reader, reader.uint32());
          break;
        case 18:
          message.bridge_active = reader.bool();
          break;
        case 19:
          message.ethereum_blacklist.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.ethereum_blacklist = [];
    if (object.gravity_id !== undefined && object.gravity_id !== null) {
      message.gravity_id = String(object.gravity_id);
    } else {
      message.gravity_id = "";
    }
    if (
      object.contract_source_hash !== undefined &&
      object.contract_source_hash !== null
    ) {
      message.contract_source_hash = String(object.contract_source_hash);
    } else {
      message.contract_source_hash = "";
    }
    if (
      object.bridge_ethereum_address !== undefined &&
      object.bridge_ethereum_address !== null
    ) {
      message.bridge_ethereum_address = String(object.bridge_ethereum_address);
    } else {
      message.bridge_ethereum_address = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = Number(object.bridge_chain_id);
    } else {
      message.bridge_chain_id = 0;
    }
    if (
      object.signed_valsets_window !== undefined &&
      object.signed_valsets_window !== null
    ) {
      message.signed_valsets_window = Number(object.signed_valsets_window);
    } else {
      message.signed_valsets_window = 0;
    }
    if (
      object.signed_batches_window !== undefined &&
      object.signed_batches_window !== null
    ) {
      message.signed_batches_window = Number(object.signed_batches_window);
    } else {
      message.signed_batches_window = 0;
    }
    if (
      object.signed_logic_calls_window !== undefined &&
      object.signed_logic_calls_window !== null
    ) {
      message.signed_logic_calls_window = Number(
        object.signed_logic_calls_window
      );
    } else {
      message.signed_logic_calls_window = 0;
    }
    if (
      object.target_batch_timeout !== undefined &&
      object.target_batch_timeout !== null
    ) {
      message.target_batch_timeout = Number(object.target_batch_timeout);
    } else {
      message.target_batch_timeout = 0;
    }
    if (
      object.average_block_time !== undefined &&
      object.average_block_time !== null
    ) {
      message.average_block_time = Number(object.average_block_time);
    } else {
      message.average_block_time = 0;
    }
    if (
      object.average_ethereum_block_time !== undefined &&
      object.average_ethereum_block_time !== null
    ) {
      message.average_ethereum_block_time = Number(
        object.average_ethereum_block_time
      );
    } else {
      message.average_ethereum_block_time = 0;
    }
    if (
      object.slash_fraction_valset !== undefined &&
      object.slash_fraction_valset !== null
    ) {
      message.slash_fraction_valset = bytesFromBase64(
        object.slash_fraction_valset
      );
    }
    if (
      object.slash_fraction_batch !== undefined &&
      object.slash_fraction_batch !== null
    ) {
      message.slash_fraction_batch = bytesFromBase64(
        object.slash_fraction_batch
      );
    }
    if (
      object.slash_fraction_logic_call !== undefined &&
      object.slash_fraction_logic_call !== null
    ) {
      message.slash_fraction_logic_call = bytesFromBase64(
        object.slash_fraction_logic_call
      );
    }
    if (
      object.unbond_slashing_valsets_window !== undefined &&
      object.unbond_slashing_valsets_window !== null
    ) {
      message.unbond_slashing_valsets_window = Number(
        object.unbond_slashing_valsets_window
      );
    } else {
      message.unbond_slashing_valsets_window = 0;
    }
    if (
      object.slash_fraction_bad_eth_signature !== undefined &&
      object.slash_fraction_bad_eth_signature !== null
    ) {
      message.slash_fraction_bad_eth_signature = bytesFromBase64(
        object.slash_fraction_bad_eth_signature
      );
    }
    if (object.valset_reward !== undefined && object.valset_reward !== null) {
      message.valset_reward = Coin.fromJSON(object.valset_reward);
    } else {
      message.valset_reward = undefined;
    }
    if (object.bridge_active !== undefined && object.bridge_active !== null) {
      message.bridge_active = Boolean(object.bridge_active);
    } else {
      message.bridge_active = false;
    }
    if (
      object.ethereum_blacklist !== undefined &&
      object.ethereum_blacklist !== null
    ) {
      for (const e of object.ethereum_blacklist) {
        message.ethereum_blacklist.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.gravity_id !== undefined && (obj.gravity_id = message.gravity_id);
    message.contract_source_hash !== undefined &&
      (obj.contract_source_hash = message.contract_source_hash);
    message.bridge_ethereum_address !== undefined &&
      (obj.bridge_ethereum_address = message.bridge_ethereum_address);
    message.bridge_chain_id !== undefined &&
      (obj.bridge_chain_id = message.bridge_chain_id);
    message.signed_valsets_window !== undefined &&
      (obj.signed_valsets_window = message.signed_valsets_window);
    message.signed_batches_window !== undefined &&
      (obj.signed_batches_window = message.signed_batches_window);
    message.signed_logic_calls_window !== undefined &&
      (obj.signed_logic_calls_window = message.signed_logic_calls_window);
    message.target_batch_timeout !== undefined &&
      (obj.target_batch_timeout = message.target_batch_timeout);
    message.average_block_time !== undefined &&
      (obj.average_block_time = message.average_block_time);
    message.average_ethereum_block_time !== undefined &&
      (obj.average_ethereum_block_time = message.average_ethereum_block_time);
    message.slash_fraction_valset !== undefined &&
      (obj.slash_fraction_valset = base64FromBytes(
        message.slash_fraction_valset !== undefined
          ? message.slash_fraction_valset
          : new Uint8Array()
      ));
    message.slash_fraction_batch !== undefined &&
      (obj.slash_fraction_batch = base64FromBytes(
        message.slash_fraction_batch !== undefined
          ? message.slash_fraction_batch
          : new Uint8Array()
      ));
    message.slash_fraction_logic_call !== undefined &&
      (obj.slash_fraction_logic_call = base64FromBytes(
        message.slash_fraction_logic_call !== undefined
          ? message.slash_fraction_logic_call
          : new Uint8Array()
      ));
    message.unbond_slashing_valsets_window !== undefined &&
      (obj.unbond_slashing_valsets_window =
        message.unbond_slashing_valsets_window);
    message.slash_fraction_bad_eth_signature !== undefined &&
      (obj.slash_fraction_bad_eth_signature = base64FromBytes(
        message.slash_fraction_bad_eth_signature !== undefined
          ? message.slash_fraction_bad_eth_signature
          : new Uint8Array()
      ));
    message.valset_reward !== undefined &&
      (obj.valset_reward = message.valset_reward
        ? Coin.toJSON(message.valset_reward)
        : undefined);
    message.bridge_active !== undefined &&
      (obj.bridge_active = message.bridge_active);
    if (message.ethereum_blacklist) {
      obj.ethereum_blacklist = message.ethereum_blacklist.map((e) => e);
    } else {
      obj.ethereum_blacklist = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.ethereum_blacklist = [];
    if (object.gravity_id !== undefined && object.gravity_id !== null) {
      message.gravity_id = object.gravity_id;
    } else {
      message.gravity_id = "";
    }
    if (
      object.contract_source_hash !== undefined &&
      object.contract_source_hash !== null
    ) {
      message.contract_source_hash = object.contract_source_hash;
    } else {
      message.contract_source_hash = "";
    }
    if (
      object.bridge_ethereum_address !== undefined &&
      object.bridge_ethereum_address !== null
    ) {
      message.bridge_ethereum_address = object.bridge_ethereum_address;
    } else {
      message.bridge_ethereum_address = "";
    }
    if (
      object.bridge_chain_id !== undefined &&
      object.bridge_chain_id !== null
    ) {
      message.bridge_chain_id = object.bridge_chain_id;
    } else {
      message.bridge_chain_id = 0;
    }
    if (
      object.signed_valsets_window !== undefined &&
      object.signed_valsets_window !== null
    ) {
      message.signed_valsets_window = object.signed_valsets_window;
    } else {
      message.signed_valsets_window = 0;
    }
    if (
      object.signed_batches_window !== undefined &&
      object.signed_batches_window !== null
    ) {
      message.signed_batches_window = object.signed_batches_window;
    } else {
      message.signed_batches_window = 0;
    }
    if (
      object.signed_logic_calls_window !== undefined &&
      object.signed_logic_calls_window !== null
    ) {
      message.signed_logic_calls_window = object.signed_logic_calls_window;
    } else {
      message.signed_logic_calls_window = 0;
    }
    if (
      object.target_batch_timeout !== undefined &&
      object.target_batch_timeout !== null
    ) {
      message.target_batch_timeout = object.target_batch_timeout;
    } else {
      message.target_batch_timeout = 0;
    }
    if (
      object.average_block_time !== undefined &&
      object.average_block_time !== null
    ) {
      message.average_block_time = object.average_block_time;
    } else {
      message.average_block_time = 0;
    }
    if (
      object.average_ethereum_block_time !== undefined &&
      object.average_ethereum_block_time !== null
    ) {
      message.average_ethereum_block_time = object.average_ethereum_block_time;
    } else {
      message.average_ethereum_block_time = 0;
    }
    if (
      object.slash_fraction_valset !== undefined &&
      object.slash_fraction_valset !== null
    ) {
      message.slash_fraction_valset = object.slash_fraction_valset;
    } else {
      message.slash_fraction_valset = new Uint8Array();
    }
    if (
      object.slash_fraction_batch !== undefined &&
      object.slash_fraction_batch !== null
    ) {
      message.slash_fraction_batch = object.slash_fraction_batch;
    } else {
      message.slash_fraction_batch = new Uint8Array();
    }
    if (
      object.slash_fraction_logic_call !== undefined &&
      object.slash_fraction_logic_call !== null
    ) {
      message.slash_fraction_logic_call = object.slash_fraction_logic_call;
    } else {
      message.slash_fraction_logic_call = new Uint8Array();
    }
    if (
      object.unbond_slashing_valsets_window !== undefined &&
      object.unbond_slashing_valsets_window !== null
    ) {
      message.unbond_slashing_valsets_window =
        object.unbond_slashing_valsets_window;
    } else {
      message.unbond_slashing_valsets_window = 0;
    }
    if (
      object.slash_fraction_bad_eth_signature !== undefined &&
      object.slash_fraction_bad_eth_signature !== null
    ) {
      message.slash_fraction_bad_eth_signature =
        object.slash_fraction_bad_eth_signature;
    } else {
      message.slash_fraction_bad_eth_signature = new Uint8Array();
    }
    if (object.valset_reward !== undefined && object.valset_reward !== null) {
      message.valset_reward = Coin.fromPartial(object.valset_reward);
    } else {
      message.valset_reward = undefined;
    }
    if (object.bridge_active !== undefined && object.bridge_active !== null) {
      message.bridge_active = object.bridge_active;
    } else {
      message.bridge_active = false;
    }
    if (
      object.ethereum_blacklist !== undefined &&
      object.ethereum_blacklist !== null
    ) {
      for (const e of object.ethereum_blacklist) {
        message.ethereum_blacklist.push(e);
      }
    }
    return message;
  },
};

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.gravity_nonces !== undefined) {
      GravityNonces.encode(
        message.gravity_nonces,
        writer.uint32(18).fork()
      ).ldelim();
    }
    for (const v of message.valsets) {
      Valset.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.valset_confirms) {
      MsgValsetConfirm.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.batches) {
      OutgoingTxBatch.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.batch_confirms) {
      MsgConfirmBatch.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.logic_calls) {
      OutgoingLogicCall.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.logic_call_confirms) {
      MsgConfirmLogicCall.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.attestations) {
      Attestation.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    for (const v of message.delegate_keys) {
      MsgSetOrchestratorAddress.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.erc20_to_denoms) {
      ERC20ToDenom.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    for (const v of message.unbatched_transfers) {
      OutgoingTransferTx.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.valsets = [];
    message.valset_confirms = [];
    message.batches = [];
    message.batch_confirms = [];
    message.logic_calls = [];
    message.logic_call_confirms = [];
    message.attestations = [];
    message.delegate_keys = [];
    message.erc20_to_denoms = [];
    message.unbatched_transfers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.gravity_nonces = GravityNonces.decode(
            reader,
            reader.uint32()
          );
          break;
        case 3:
          message.valsets.push(Valset.decode(reader, reader.uint32()));
          break;
        case 4:
          message.valset_confirms.push(
            MsgValsetConfirm.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.batches.push(OutgoingTxBatch.decode(reader, reader.uint32()));
          break;
        case 6:
          message.batch_confirms.push(
            MsgConfirmBatch.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.logic_calls.push(
            OutgoingLogicCall.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.logic_call_confirms.push(
            MsgConfirmLogicCall.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.attestations.push(
            Attestation.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.delegate_keys.push(
            MsgSetOrchestratorAddress.decode(reader, reader.uint32())
          );
          break;
        case 11:
          message.erc20_to_denoms.push(
            ERC20ToDenom.decode(reader, reader.uint32())
          );
          break;
        case 12:
          message.unbatched_transfers.push(
            OutgoingTransferTx.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.valsets = [];
    message.valset_confirms = [];
    message.batches = [];
    message.batch_confirms = [];
    message.logic_calls = [];
    message.logic_call_confirms = [];
    message.attestations = [];
    message.delegate_keys = [];
    message.erc20_to_denoms = [];
    message.unbatched_transfers = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.gravity_nonces !== undefined && object.gravity_nonces !== null) {
      message.gravity_nonces = GravityNonces.fromJSON(object.gravity_nonces);
    } else {
      message.gravity_nonces = undefined;
    }
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromJSON(e));
      }
    }
    if (
      object.valset_confirms !== undefined &&
      object.valset_confirms !== null
    ) {
      for (const e of object.valset_confirms) {
        message.valset_confirms.push(MsgValsetConfirm.fromJSON(e));
      }
    }
    if (object.batches !== undefined && object.batches !== null) {
      for (const e of object.batches) {
        message.batches.push(OutgoingTxBatch.fromJSON(e));
      }
    }
    if (object.batch_confirms !== undefined && object.batch_confirms !== null) {
      for (const e of object.batch_confirms) {
        message.batch_confirms.push(MsgConfirmBatch.fromJSON(e));
      }
    }
    if (object.logic_calls !== undefined && object.logic_calls !== null) {
      for (const e of object.logic_calls) {
        message.logic_calls.push(OutgoingLogicCall.fromJSON(e));
      }
    }
    if (
      object.logic_call_confirms !== undefined &&
      object.logic_call_confirms !== null
    ) {
      for (const e of object.logic_call_confirms) {
        message.logic_call_confirms.push(MsgConfirmLogicCall.fromJSON(e));
      }
    }
    if (object.attestations !== undefined && object.attestations !== null) {
      for (const e of object.attestations) {
        message.attestations.push(Attestation.fromJSON(e));
      }
    }
    if (object.delegate_keys !== undefined && object.delegate_keys !== null) {
      for (const e of object.delegate_keys) {
        message.delegate_keys.push(MsgSetOrchestratorAddress.fromJSON(e));
      }
    }
    if (
      object.erc20_to_denoms !== undefined &&
      object.erc20_to_denoms !== null
    ) {
      for (const e of object.erc20_to_denoms) {
        message.erc20_to_denoms.push(ERC20ToDenom.fromJSON(e));
      }
    }
    if (
      object.unbatched_transfers !== undefined &&
      object.unbatched_transfers !== null
    ) {
      for (const e of object.unbatched_transfers) {
        message.unbatched_transfers.push(OutgoingTransferTx.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.gravity_nonces !== undefined &&
      (obj.gravity_nonces = message.gravity_nonces
        ? GravityNonces.toJSON(message.gravity_nonces)
        : undefined);
    if (message.valsets) {
      obj.valsets = message.valsets.map((e) =>
        e ? Valset.toJSON(e) : undefined
      );
    } else {
      obj.valsets = [];
    }
    if (message.valset_confirms) {
      obj.valset_confirms = message.valset_confirms.map((e) =>
        e ? MsgValsetConfirm.toJSON(e) : undefined
      );
    } else {
      obj.valset_confirms = [];
    }
    if (message.batches) {
      obj.batches = message.batches.map((e) =>
        e ? OutgoingTxBatch.toJSON(e) : undefined
      );
    } else {
      obj.batches = [];
    }
    if (message.batch_confirms) {
      obj.batch_confirms = message.batch_confirms.map((e) =>
        e ? MsgConfirmBatch.toJSON(e) : undefined
      );
    } else {
      obj.batch_confirms = [];
    }
    if (message.logic_calls) {
      obj.logic_calls = message.logic_calls.map((e) =>
        e ? OutgoingLogicCall.toJSON(e) : undefined
      );
    } else {
      obj.logic_calls = [];
    }
    if (message.logic_call_confirms) {
      obj.logic_call_confirms = message.logic_call_confirms.map((e) =>
        e ? MsgConfirmLogicCall.toJSON(e) : undefined
      );
    } else {
      obj.logic_call_confirms = [];
    }
    if (message.attestations) {
      obj.attestations = message.attestations.map((e) =>
        e ? Attestation.toJSON(e) : undefined
      );
    } else {
      obj.attestations = [];
    }
    if (message.delegate_keys) {
      obj.delegate_keys = message.delegate_keys.map((e) =>
        e ? MsgSetOrchestratorAddress.toJSON(e) : undefined
      );
    } else {
      obj.delegate_keys = [];
    }
    if (message.erc20_to_denoms) {
      obj.erc20_to_denoms = message.erc20_to_denoms.map((e) =>
        e ? ERC20ToDenom.toJSON(e) : undefined
      );
    } else {
      obj.erc20_to_denoms = [];
    }
    if (message.unbatched_transfers) {
      obj.unbatched_transfers = message.unbatched_transfers.map((e) =>
        e ? OutgoingTransferTx.toJSON(e) : undefined
      );
    } else {
      obj.unbatched_transfers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.valsets = [];
    message.valset_confirms = [];
    message.batches = [];
    message.batch_confirms = [];
    message.logic_calls = [];
    message.logic_call_confirms = [];
    message.attestations = [];
    message.delegate_keys = [];
    message.erc20_to_denoms = [];
    message.unbatched_transfers = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.gravity_nonces !== undefined && object.gravity_nonces !== null) {
      message.gravity_nonces = GravityNonces.fromPartial(object.gravity_nonces);
    } else {
      message.gravity_nonces = undefined;
    }
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromPartial(e));
      }
    }
    if (
      object.valset_confirms !== undefined &&
      object.valset_confirms !== null
    ) {
      for (const e of object.valset_confirms) {
        message.valset_confirms.push(MsgValsetConfirm.fromPartial(e));
      }
    }
    if (object.batches !== undefined && object.batches !== null) {
      for (const e of object.batches) {
        message.batches.push(OutgoingTxBatch.fromPartial(e));
      }
    }
    if (object.batch_confirms !== undefined && object.batch_confirms !== null) {
      for (const e of object.batch_confirms) {
        message.batch_confirms.push(MsgConfirmBatch.fromPartial(e));
      }
    }
    if (object.logic_calls !== undefined && object.logic_calls !== null) {
      for (const e of object.logic_calls) {
        message.logic_calls.push(OutgoingLogicCall.fromPartial(e));
      }
    }
    if (
      object.logic_call_confirms !== undefined &&
      object.logic_call_confirms !== null
    ) {
      for (const e of object.logic_call_confirms) {
        message.logic_call_confirms.push(MsgConfirmLogicCall.fromPartial(e));
      }
    }
    if (object.attestations !== undefined && object.attestations !== null) {
      for (const e of object.attestations) {
        message.attestations.push(Attestation.fromPartial(e));
      }
    }
    if (object.delegate_keys !== undefined && object.delegate_keys !== null) {
      for (const e of object.delegate_keys) {
        message.delegate_keys.push(MsgSetOrchestratorAddress.fromPartial(e));
      }
    }
    if (
      object.erc20_to_denoms !== undefined &&
      object.erc20_to_denoms !== null
    ) {
      for (const e of object.erc20_to_denoms) {
        message.erc20_to_denoms.push(ERC20ToDenom.fromPartial(e));
      }
    }
    if (
      object.unbatched_transfers !== undefined &&
      object.unbatched_transfers !== null
    ) {
      for (const e of object.unbatched_transfers) {
        message.unbatched_transfers.push(OutgoingTransferTx.fromPartial(e));
      }
    }
    return message;
  },
};

const baseGravityNonces: object = {
  latest_valset_nonce: 0,
  last_observed_nonce: 0,
  last_slashed_valset_nonce: 0,
  last_slashed_batch_block: 0,
  last_slashed_logic_call_block: 0,
  last_tx_pool_id: 0,
  last_batch_id: 0,
};

export const GravityNonces = {
  encode(message: GravityNonces, writer: Writer = Writer.create()): Writer {
    if (message.latest_valset_nonce !== 0) {
      writer.uint32(8).uint64(message.latest_valset_nonce);
    }
    if (message.last_observed_nonce !== 0) {
      writer.uint32(16).uint64(message.last_observed_nonce);
    }
    if (message.last_slashed_valset_nonce !== 0) {
      writer.uint32(24).uint64(message.last_slashed_valset_nonce);
    }
    if (message.last_slashed_batch_block !== 0) {
      writer.uint32(32).uint64(message.last_slashed_batch_block);
    }
    if (message.last_slashed_logic_call_block !== 0) {
      writer.uint32(40).uint64(message.last_slashed_logic_call_block);
    }
    if (message.last_tx_pool_id !== 0) {
      writer.uint32(48).uint64(message.last_tx_pool_id);
    }
    if (message.last_batch_id !== 0) {
      writer.uint32(56).uint64(message.last_batch_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GravityNonces {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGravityNonces } as GravityNonces;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.latest_valset_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.last_observed_nonce = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.last_slashed_valset_nonce = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 4:
          message.last_slashed_batch_block = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 5:
          message.last_slashed_logic_call_block = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 6:
          message.last_tx_pool_id = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.last_batch_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GravityNonces {
    const message = { ...baseGravityNonces } as GravityNonces;
    if (
      object.latest_valset_nonce !== undefined &&
      object.latest_valset_nonce !== null
    ) {
      message.latest_valset_nonce = Number(object.latest_valset_nonce);
    } else {
      message.latest_valset_nonce = 0;
    }
    if (
      object.last_observed_nonce !== undefined &&
      object.last_observed_nonce !== null
    ) {
      message.last_observed_nonce = Number(object.last_observed_nonce);
    } else {
      message.last_observed_nonce = 0;
    }
    if (
      object.last_slashed_valset_nonce !== undefined &&
      object.last_slashed_valset_nonce !== null
    ) {
      message.last_slashed_valset_nonce = Number(
        object.last_slashed_valset_nonce
      );
    } else {
      message.last_slashed_valset_nonce = 0;
    }
    if (
      object.last_slashed_batch_block !== undefined &&
      object.last_slashed_batch_block !== null
    ) {
      message.last_slashed_batch_block = Number(
        object.last_slashed_batch_block
      );
    } else {
      message.last_slashed_batch_block = 0;
    }
    if (
      object.last_slashed_logic_call_block !== undefined &&
      object.last_slashed_logic_call_block !== null
    ) {
      message.last_slashed_logic_call_block = Number(
        object.last_slashed_logic_call_block
      );
    } else {
      message.last_slashed_logic_call_block = 0;
    }
    if (
      object.last_tx_pool_id !== undefined &&
      object.last_tx_pool_id !== null
    ) {
      message.last_tx_pool_id = Number(object.last_tx_pool_id);
    } else {
      message.last_tx_pool_id = 0;
    }
    if (object.last_batch_id !== undefined && object.last_batch_id !== null) {
      message.last_batch_id = Number(object.last_batch_id);
    } else {
      message.last_batch_id = 0;
    }
    return message;
  },

  toJSON(message: GravityNonces): unknown {
    const obj: any = {};
    message.latest_valset_nonce !== undefined &&
      (obj.latest_valset_nonce = message.latest_valset_nonce);
    message.last_observed_nonce !== undefined &&
      (obj.last_observed_nonce = message.last_observed_nonce);
    message.last_slashed_valset_nonce !== undefined &&
      (obj.last_slashed_valset_nonce = message.last_slashed_valset_nonce);
    message.last_slashed_batch_block !== undefined &&
      (obj.last_slashed_batch_block = message.last_slashed_batch_block);
    message.last_slashed_logic_call_block !== undefined &&
      (obj.last_slashed_logic_call_block =
        message.last_slashed_logic_call_block);
    message.last_tx_pool_id !== undefined &&
      (obj.last_tx_pool_id = message.last_tx_pool_id);
    message.last_batch_id !== undefined &&
      (obj.last_batch_id = message.last_batch_id);
    return obj;
  },

  fromPartial(object: DeepPartial<GravityNonces>): GravityNonces {
    const message = { ...baseGravityNonces } as GravityNonces;
    if (
      object.latest_valset_nonce !== undefined &&
      object.latest_valset_nonce !== null
    ) {
      message.latest_valset_nonce = object.latest_valset_nonce;
    } else {
      message.latest_valset_nonce = 0;
    }
    if (
      object.last_observed_nonce !== undefined &&
      object.last_observed_nonce !== null
    ) {
      message.last_observed_nonce = object.last_observed_nonce;
    } else {
      message.last_observed_nonce = 0;
    }
    if (
      object.last_slashed_valset_nonce !== undefined &&
      object.last_slashed_valset_nonce !== null
    ) {
      message.last_slashed_valset_nonce = object.last_slashed_valset_nonce;
    } else {
      message.last_slashed_valset_nonce = 0;
    }
    if (
      object.last_slashed_batch_block !== undefined &&
      object.last_slashed_batch_block !== null
    ) {
      message.last_slashed_batch_block = object.last_slashed_batch_block;
    } else {
      message.last_slashed_batch_block = 0;
    }
    if (
      object.last_slashed_logic_call_block !== undefined &&
      object.last_slashed_logic_call_block !== null
    ) {
      message.last_slashed_logic_call_block =
        object.last_slashed_logic_call_block;
    } else {
      message.last_slashed_logic_call_block = 0;
    }
    if (
      object.last_tx_pool_id !== undefined &&
      object.last_tx_pool_id !== null
    ) {
      message.last_tx_pool_id = object.last_tx_pool_id;
    } else {
      message.last_tx_pool_id = 0;
    }
    if (object.last_batch_id !== undefined && object.last_batch_id !== null) {
      message.last_batch_id = object.last_batch_id;
    } else {
      message.last_batch_id = 0;
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
