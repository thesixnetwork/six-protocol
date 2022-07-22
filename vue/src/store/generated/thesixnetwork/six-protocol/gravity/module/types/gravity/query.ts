/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../gravity/genesis";
import { Valset, PendingIbcAutoForward } from "../gravity/types";
import {
  MsgValsetConfirm,
  MsgConfirmBatch,
  MsgConfirmLogicCall,
} from "../gravity/msgs";
import { BatchFees } from "../gravity/pool";
import {
  OutgoingTxBatch,
  OutgoingLogicCall,
  OutgoingTransferTx,
} from "../gravity/batch";
import { Attestation } from "../gravity/attestation";

export const protobufPackage = "gravity";

export interface QueryParamsRequest {}

export interface QueryParamsResponse {
  params: Params | undefined;
}

export interface QueryCurrentValsetRequest {}

export interface QueryCurrentValsetResponse {
  valset: Valset | undefined;
}

export interface QueryValsetRequestRequest {
  nonce: number;
}

export interface QueryValsetRequestResponse {
  valset: Valset | undefined;
}

export interface QueryValsetConfirmRequest {
  nonce: number;
  address: string;
}

export interface QueryValsetConfirmResponse {
  confirm: MsgValsetConfirm | undefined;
}

export interface QueryValsetConfirmsByNonceRequest {
  nonce: number;
}

export interface QueryValsetConfirmsByNonceResponse {
  confirms: MsgValsetConfirm[];
}

export interface QueryLastValsetRequestsRequest {}

export interface QueryLastValsetRequestsResponse {
  valsets: Valset[];
}

export interface QueryLastPendingValsetRequestByAddrRequest {
  address: string;
}

export interface QueryLastPendingValsetRequestByAddrResponse {
  valsets: Valset[];
}

export interface QueryBatchFeeRequest {}

export interface QueryBatchFeeResponse {
  batch_fees: BatchFees[];
}

export interface QueryLastPendingBatchRequestByAddrRequest {
  address: string;
}

export interface QueryLastPendingBatchRequestByAddrResponse {
  batch: OutgoingTxBatch[];
}

export interface QueryLastPendingLogicCallByAddrRequest {
  address: string;
}

export interface QueryLastPendingLogicCallByAddrResponse {
  call: OutgoingLogicCall[];
}

export interface QueryOutgoingTxBatchesRequest {}

export interface QueryOutgoingTxBatchesResponse {
  batches: OutgoingTxBatch[];
}

export interface QueryOutgoingLogicCallsRequest {}

export interface QueryOutgoingLogicCallsResponse {
  calls: OutgoingLogicCall[];
}

export interface QueryBatchRequestByNonceRequest {
  nonce: number;
  contract_address: string;
}

export interface QueryBatchRequestByNonceResponse {
  batch: OutgoingTxBatch | undefined;
}

export interface QueryBatchConfirmsRequest {
  nonce: number;
  contract_address: string;
}

export interface QueryBatchConfirmsResponse {
  confirms: MsgConfirmBatch[];
}

export interface QueryLogicConfirmsRequest {
  invalidation_id: Uint8Array;
  invalidation_nonce: number;
}

export interface QueryLogicConfirmsResponse {
  confirms: MsgConfirmLogicCall[];
}

export interface QueryLastEventNonceByAddrRequest {
  address: string;
}

export interface QueryLastEventNonceByAddrResponse {
  event_nonce: number;
}

export interface QueryERC20ToDenomRequest {
  erc20: string;
}

export interface QueryERC20ToDenomResponse {
  denom: string;
  cosmos_originated: boolean;
}

export interface QueryDenomToERC20Request {
  denom: string;
}

export interface QueryDenomToERC20Response {
  erc20: string;
  cosmos_originated: boolean;
}

/**
 * QueryAttestationsRequest defines the request structure for getting recent
 * attestations with optional query parameters. By default, a limited set of
 * recent attestations will be returned, defined by 'limit'. These attestations
 * can be ordered ascending or descending by nonce, that defaults to ascending.
 * Filtering criteria may also be provided, including nonce, claim type, and
 * height. Note, that an attestation will be returned if it matches ANY of the
 * filter query parameters provided.
 */
export interface QueryAttestationsRequest {
  /** limit defines how many attestations to limit in the response. */
  limit: number;
  /**
   * order_by provides ordering of atteststions by nonce in the response. Either
   * 'asc' or 'desc' can be provided. If no value is provided, it defaults to
   * 'asc'.
   */
  order_by: string;
  /** claim_type allows filtering attestations by Ethereum claim type. */
  claim_type: string;
  /** nonce allows filtering attestations by Ethereum claim nonce. */
  nonce: number;
  /** height allows filtering attestations by Ethereum claim height. */
  height: number;
}

export interface QueryAttestationsResponse {
  attestations: Attestation[];
}

export interface QueryDelegateKeysByValidatorAddress {
  validator_address: string;
}

export interface QueryDelegateKeysByValidatorAddressResponse {
  eth_address: string;
  orchestrator_address: string;
}

export interface QueryDelegateKeysByEthAddress {
  eth_address: string;
}

export interface QueryDelegateKeysByEthAddressResponse {
  validator_address: string;
  orchestrator_address: string;
}

export interface QueryDelegateKeysByOrchestratorAddress {
  orchestrator_address: string;
}

export interface QueryDelegateKeysByOrchestratorAddressResponse {
  validator_address: string;
  eth_address: string;
}

export interface QueryPendingSendToEth {
  sender_address: string;
}

export interface QueryPendingSendToEthResponse {
  transfers_in_batches: OutgoingTransferTx[];
  unbatched_transfers: OutgoingTransferTx[];
}

export interface QueryPendingIbcAutoForwards {
  /** limit defines the number of pending forwards to return, in order of their SendToCosmos.EventNonce */
  limit: number;
}

export interface QueryPendingIbcAutoForwardsResponse {
  pending_ibc_auto_forwards: PendingIbcAutoForward[];
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryCurrentValsetRequest: object = {};

export const QueryCurrentValsetRequest = {
  encode(
    _: QueryCurrentValsetRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCurrentValsetRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCurrentValsetRequest,
    } as QueryCurrentValsetRequest;
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

  fromJSON(_: any): QueryCurrentValsetRequest {
    const message = {
      ...baseQueryCurrentValsetRequest,
    } as QueryCurrentValsetRequest;
    return message;
  },

  toJSON(_: QueryCurrentValsetRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryCurrentValsetRequest>
  ): QueryCurrentValsetRequest {
    const message = {
      ...baseQueryCurrentValsetRequest,
    } as QueryCurrentValsetRequest;
    return message;
  },
};

const baseQueryCurrentValsetResponse: object = {};

export const QueryCurrentValsetResponse = {
  encode(
    message: QueryCurrentValsetResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.valset !== undefined) {
      Valset.encode(message.valset, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCurrentValsetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCurrentValsetResponse,
    } as QueryCurrentValsetResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.valset = Valset.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCurrentValsetResponse {
    const message = {
      ...baseQueryCurrentValsetResponse,
    } as QueryCurrentValsetResponse;
    if (object.valset !== undefined && object.valset !== null) {
      message.valset = Valset.fromJSON(object.valset);
    } else {
      message.valset = undefined;
    }
    return message;
  },

  toJSON(message: QueryCurrentValsetResponse): unknown {
    const obj: any = {};
    message.valset !== undefined &&
      (obj.valset = message.valset ? Valset.toJSON(message.valset) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCurrentValsetResponse>
  ): QueryCurrentValsetResponse {
    const message = {
      ...baseQueryCurrentValsetResponse,
    } as QueryCurrentValsetResponse;
    if (object.valset !== undefined && object.valset !== null) {
      message.valset = Valset.fromPartial(object.valset);
    } else {
      message.valset = undefined;
    }
    return message;
  },
};

const baseQueryValsetRequestRequest: object = { nonce: 0 };

export const QueryValsetRequestRequest = {
  encode(
    message: QueryValsetRequestRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetRequestRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetRequestRequest,
    } as QueryValsetRequestRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValsetRequestRequest {
    const message = {
      ...baseQueryValsetRequestRequest,
    } as QueryValsetRequestRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    return message;
  },

  toJSON(message: QueryValsetRequestRequest): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetRequestRequest>
  ): QueryValsetRequestRequest {
    const message = {
      ...baseQueryValsetRequestRequest,
    } as QueryValsetRequestRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    return message;
  },
};

const baseQueryValsetRequestResponse: object = {};

export const QueryValsetRequestResponse = {
  encode(
    message: QueryValsetRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.valset !== undefined) {
      Valset.encode(message.valset, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetRequestResponse,
    } as QueryValsetRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.valset = Valset.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValsetRequestResponse {
    const message = {
      ...baseQueryValsetRequestResponse,
    } as QueryValsetRequestResponse;
    if (object.valset !== undefined && object.valset !== null) {
      message.valset = Valset.fromJSON(object.valset);
    } else {
      message.valset = undefined;
    }
    return message;
  },

  toJSON(message: QueryValsetRequestResponse): unknown {
    const obj: any = {};
    message.valset !== undefined &&
      (obj.valset = message.valset ? Valset.toJSON(message.valset) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetRequestResponse>
  ): QueryValsetRequestResponse {
    const message = {
      ...baseQueryValsetRequestResponse,
    } as QueryValsetRequestResponse;
    if (object.valset !== undefined && object.valset !== null) {
      message.valset = Valset.fromPartial(object.valset);
    } else {
      message.valset = undefined;
    }
    return message;
  },
};

const baseQueryValsetConfirmRequest: object = { nonce: 0, address: "" };

export const QueryValsetConfirmRequest = {
  encode(
    message: QueryValsetConfirmRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetConfirmRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetConfirmRequest,
    } as QueryValsetConfirmRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): QueryValsetConfirmRequest {
    const message = {
      ...baseQueryValsetConfirmRequest,
    } as QueryValsetConfirmRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryValsetConfirmRequest): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetConfirmRequest>
  ): QueryValsetConfirmRequest {
    const message = {
      ...baseQueryValsetConfirmRequest,
    } as QueryValsetConfirmRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryValsetConfirmResponse: object = {};

export const QueryValsetConfirmResponse = {
  encode(
    message: QueryValsetConfirmResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.confirm !== undefined) {
      MsgValsetConfirm.encode(
        message.confirm,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetConfirmResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetConfirmResponse,
    } as QueryValsetConfirmResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirm = MsgValsetConfirm.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValsetConfirmResponse {
    const message = {
      ...baseQueryValsetConfirmResponse,
    } as QueryValsetConfirmResponse;
    if (object.confirm !== undefined && object.confirm !== null) {
      message.confirm = MsgValsetConfirm.fromJSON(object.confirm);
    } else {
      message.confirm = undefined;
    }
    return message;
  },

  toJSON(message: QueryValsetConfirmResponse): unknown {
    const obj: any = {};
    message.confirm !== undefined &&
      (obj.confirm = message.confirm
        ? MsgValsetConfirm.toJSON(message.confirm)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetConfirmResponse>
  ): QueryValsetConfirmResponse {
    const message = {
      ...baseQueryValsetConfirmResponse,
    } as QueryValsetConfirmResponse;
    if (object.confirm !== undefined && object.confirm !== null) {
      message.confirm = MsgValsetConfirm.fromPartial(object.confirm);
    } else {
      message.confirm = undefined;
    }
    return message;
  },
};

const baseQueryValsetConfirmsByNonceRequest: object = { nonce: 0 };

export const QueryValsetConfirmsByNonceRequest = {
  encode(
    message: QueryValsetConfirmsByNonceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetConfirmsByNonceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetConfirmsByNonceRequest,
    } as QueryValsetConfirmsByNonceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValsetConfirmsByNonceRequest {
    const message = {
      ...baseQueryValsetConfirmsByNonceRequest,
    } as QueryValsetConfirmsByNonceRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    return message;
  },

  toJSON(message: QueryValsetConfirmsByNonceRequest): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetConfirmsByNonceRequest>
  ): QueryValsetConfirmsByNonceRequest {
    const message = {
      ...baseQueryValsetConfirmsByNonceRequest,
    } as QueryValsetConfirmsByNonceRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    return message;
  },
};

const baseQueryValsetConfirmsByNonceResponse: object = {};

export const QueryValsetConfirmsByNonceResponse = {
  encode(
    message: QueryValsetConfirmsByNonceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.confirms) {
      MsgValsetConfirm.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValsetConfirmsByNonceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValsetConfirmsByNonceResponse,
    } as QueryValsetConfirmsByNonceResponse;
    message.confirms = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirms.push(
            MsgValsetConfirm.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValsetConfirmsByNonceResponse {
    const message = {
      ...baseQueryValsetConfirmsByNonceResponse,
    } as QueryValsetConfirmsByNonceResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgValsetConfirm.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryValsetConfirmsByNonceResponse): unknown {
    const obj: any = {};
    if (message.confirms) {
      obj.confirms = message.confirms.map((e) =>
        e ? MsgValsetConfirm.toJSON(e) : undefined
      );
    } else {
      obj.confirms = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValsetConfirmsByNonceResponse>
  ): QueryValsetConfirmsByNonceResponse {
    const message = {
      ...baseQueryValsetConfirmsByNonceResponse,
    } as QueryValsetConfirmsByNonceResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgValsetConfirm.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLastValsetRequestsRequest: object = {};

export const QueryLastValsetRequestsRequest = {
  encode(
    _: QueryLastValsetRequestsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastValsetRequestsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastValsetRequestsRequest,
    } as QueryLastValsetRequestsRequest;
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

  fromJSON(_: any): QueryLastValsetRequestsRequest {
    const message = {
      ...baseQueryLastValsetRequestsRequest,
    } as QueryLastValsetRequestsRequest;
    return message;
  },

  toJSON(_: QueryLastValsetRequestsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryLastValsetRequestsRequest>
  ): QueryLastValsetRequestsRequest {
    const message = {
      ...baseQueryLastValsetRequestsRequest,
    } as QueryLastValsetRequestsRequest;
    return message;
  },
};

const baseQueryLastValsetRequestsResponse: object = {};

export const QueryLastValsetRequestsResponse = {
  encode(
    message: QueryLastValsetRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.valsets) {
      Valset.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastValsetRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastValsetRequestsResponse,
    } as QueryLastValsetRequestsResponse;
    message.valsets = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.valsets.push(Valset.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastValsetRequestsResponse {
    const message = {
      ...baseQueryLastValsetRequestsResponse,
    } as QueryLastValsetRequestsResponse;
    message.valsets = [];
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryLastValsetRequestsResponse): unknown {
    const obj: any = {};
    if (message.valsets) {
      obj.valsets = message.valsets.map((e) =>
        e ? Valset.toJSON(e) : undefined
      );
    } else {
      obj.valsets = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastValsetRequestsResponse>
  ): QueryLastValsetRequestsResponse {
    const message = {
      ...baseQueryLastValsetRequestsResponse,
    } as QueryLastValsetRequestsResponse;
    message.valsets = [];
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLastPendingValsetRequestByAddrRequest: object = { address: "" };

export const QueryLastPendingValsetRequestByAddrRequest = {
  encode(
    message: QueryLastPendingValsetRequestByAddrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingValsetRequestByAddrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrRequest,
    } as QueryLastPendingValsetRequestByAddrRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingValsetRequestByAddrRequest {
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrRequest,
    } as QueryLastPendingValsetRequestByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryLastPendingValsetRequestByAddrRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingValsetRequestByAddrRequest>
  ): QueryLastPendingValsetRequestByAddrRequest {
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrRequest,
    } as QueryLastPendingValsetRequestByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryLastPendingValsetRequestByAddrResponse: object = {};

export const QueryLastPendingValsetRequestByAddrResponse = {
  encode(
    message: QueryLastPendingValsetRequestByAddrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.valsets) {
      Valset.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingValsetRequestByAddrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrResponse,
    } as QueryLastPendingValsetRequestByAddrResponse;
    message.valsets = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.valsets.push(Valset.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingValsetRequestByAddrResponse {
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrResponse,
    } as QueryLastPendingValsetRequestByAddrResponse;
    message.valsets = [];
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryLastPendingValsetRequestByAddrResponse): unknown {
    const obj: any = {};
    if (message.valsets) {
      obj.valsets = message.valsets.map((e) =>
        e ? Valset.toJSON(e) : undefined
      );
    } else {
      obj.valsets = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingValsetRequestByAddrResponse>
  ): QueryLastPendingValsetRequestByAddrResponse {
    const message = {
      ...baseQueryLastPendingValsetRequestByAddrResponse,
    } as QueryLastPendingValsetRequestByAddrResponse;
    message.valsets = [];
    if (object.valsets !== undefined && object.valsets !== null) {
      for (const e of object.valsets) {
        message.valsets.push(Valset.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryBatchFeeRequest: object = {};

export const QueryBatchFeeRequest = {
  encode(_: QueryBatchFeeRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBatchFeeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBatchFeeRequest } as QueryBatchFeeRequest;
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

  fromJSON(_: any): QueryBatchFeeRequest {
    const message = { ...baseQueryBatchFeeRequest } as QueryBatchFeeRequest;
    return message;
  },

  toJSON(_: QueryBatchFeeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryBatchFeeRequest>): QueryBatchFeeRequest {
    const message = { ...baseQueryBatchFeeRequest } as QueryBatchFeeRequest;
    return message;
  },
};

const baseQueryBatchFeeResponse: object = {};

export const QueryBatchFeeResponse = {
  encode(
    message: QueryBatchFeeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.batch_fees) {
      BatchFees.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBatchFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBatchFeeResponse } as QueryBatchFeeResponse;
    message.batch_fees = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batch_fees.push(BatchFees.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBatchFeeResponse {
    const message = { ...baseQueryBatchFeeResponse } as QueryBatchFeeResponse;
    message.batch_fees = [];
    if (object.batch_fees !== undefined && object.batch_fees !== null) {
      for (const e of object.batch_fees) {
        message.batch_fees.push(BatchFees.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryBatchFeeResponse): unknown {
    const obj: any = {};
    if (message.batch_fees) {
      obj.batch_fees = message.batch_fees.map((e) =>
        e ? BatchFees.toJSON(e) : undefined
      );
    } else {
      obj.batch_fees = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBatchFeeResponse>
  ): QueryBatchFeeResponse {
    const message = { ...baseQueryBatchFeeResponse } as QueryBatchFeeResponse;
    message.batch_fees = [];
    if (object.batch_fees !== undefined && object.batch_fees !== null) {
      for (const e of object.batch_fees) {
        message.batch_fees.push(BatchFees.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLastPendingBatchRequestByAddrRequest: object = { address: "" };

export const QueryLastPendingBatchRequestByAddrRequest = {
  encode(
    message: QueryLastPendingBatchRequestByAddrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingBatchRequestByAddrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrRequest,
    } as QueryLastPendingBatchRequestByAddrRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingBatchRequestByAddrRequest {
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrRequest,
    } as QueryLastPendingBatchRequestByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryLastPendingBatchRequestByAddrRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingBatchRequestByAddrRequest>
  ): QueryLastPendingBatchRequestByAddrRequest {
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrRequest,
    } as QueryLastPendingBatchRequestByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryLastPendingBatchRequestByAddrResponse: object = {};

export const QueryLastPendingBatchRequestByAddrResponse = {
  encode(
    message: QueryLastPendingBatchRequestByAddrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.batch) {
      OutgoingTxBatch.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingBatchRequestByAddrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrResponse,
    } as QueryLastPendingBatchRequestByAddrResponse;
    message.batch = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batch.push(OutgoingTxBatch.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingBatchRequestByAddrResponse {
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrResponse,
    } as QueryLastPendingBatchRequestByAddrResponse;
    message.batch = [];
    if (object.batch !== undefined && object.batch !== null) {
      for (const e of object.batch) {
        message.batch.push(OutgoingTxBatch.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryLastPendingBatchRequestByAddrResponse): unknown {
    const obj: any = {};
    if (message.batch) {
      obj.batch = message.batch.map((e) =>
        e ? OutgoingTxBatch.toJSON(e) : undefined
      );
    } else {
      obj.batch = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingBatchRequestByAddrResponse>
  ): QueryLastPendingBatchRequestByAddrResponse {
    const message = {
      ...baseQueryLastPendingBatchRequestByAddrResponse,
    } as QueryLastPendingBatchRequestByAddrResponse;
    message.batch = [];
    if (object.batch !== undefined && object.batch !== null) {
      for (const e of object.batch) {
        message.batch.push(OutgoingTxBatch.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLastPendingLogicCallByAddrRequest: object = { address: "" };

export const QueryLastPendingLogicCallByAddrRequest = {
  encode(
    message: QueryLastPendingLogicCallByAddrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingLogicCallByAddrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingLogicCallByAddrRequest,
    } as QueryLastPendingLogicCallByAddrRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingLogicCallByAddrRequest {
    const message = {
      ...baseQueryLastPendingLogicCallByAddrRequest,
    } as QueryLastPendingLogicCallByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryLastPendingLogicCallByAddrRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingLogicCallByAddrRequest>
  ): QueryLastPendingLogicCallByAddrRequest {
    const message = {
      ...baseQueryLastPendingLogicCallByAddrRequest,
    } as QueryLastPendingLogicCallByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryLastPendingLogicCallByAddrResponse: object = {};

export const QueryLastPendingLogicCallByAddrResponse = {
  encode(
    message: QueryLastPendingLogicCallByAddrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.call) {
      OutgoingLogicCall.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastPendingLogicCallByAddrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastPendingLogicCallByAddrResponse,
    } as QueryLastPendingLogicCallByAddrResponse;
    message.call = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.call.push(OutgoingLogicCall.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPendingLogicCallByAddrResponse {
    const message = {
      ...baseQueryLastPendingLogicCallByAddrResponse,
    } as QueryLastPendingLogicCallByAddrResponse;
    message.call = [];
    if (object.call !== undefined && object.call !== null) {
      for (const e of object.call) {
        message.call.push(OutgoingLogicCall.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryLastPendingLogicCallByAddrResponse): unknown {
    const obj: any = {};
    if (message.call) {
      obj.call = message.call.map((e) =>
        e ? OutgoingLogicCall.toJSON(e) : undefined
      );
    } else {
      obj.call = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastPendingLogicCallByAddrResponse>
  ): QueryLastPendingLogicCallByAddrResponse {
    const message = {
      ...baseQueryLastPendingLogicCallByAddrResponse,
    } as QueryLastPendingLogicCallByAddrResponse;
    message.call = [];
    if (object.call !== undefined && object.call !== null) {
      for (const e of object.call) {
        message.call.push(OutgoingLogicCall.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryOutgoingTxBatchesRequest: object = {};

export const QueryOutgoingTxBatchesRequest = {
  encode(
    _: QueryOutgoingTxBatchesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOutgoingTxBatchesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOutgoingTxBatchesRequest,
    } as QueryOutgoingTxBatchesRequest;
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

  fromJSON(_: any): QueryOutgoingTxBatchesRequest {
    const message = {
      ...baseQueryOutgoingTxBatchesRequest,
    } as QueryOutgoingTxBatchesRequest;
    return message;
  },

  toJSON(_: QueryOutgoingTxBatchesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryOutgoingTxBatchesRequest>
  ): QueryOutgoingTxBatchesRequest {
    const message = {
      ...baseQueryOutgoingTxBatchesRequest,
    } as QueryOutgoingTxBatchesRequest;
    return message;
  },
};

const baseQueryOutgoingTxBatchesResponse: object = {};

export const QueryOutgoingTxBatchesResponse = {
  encode(
    message: QueryOutgoingTxBatchesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.batches) {
      OutgoingTxBatch.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOutgoingTxBatchesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOutgoingTxBatchesResponse,
    } as QueryOutgoingTxBatchesResponse;
    message.batches = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batches.push(OutgoingTxBatch.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOutgoingTxBatchesResponse {
    const message = {
      ...baseQueryOutgoingTxBatchesResponse,
    } as QueryOutgoingTxBatchesResponse;
    message.batches = [];
    if (object.batches !== undefined && object.batches !== null) {
      for (const e of object.batches) {
        message.batches.push(OutgoingTxBatch.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryOutgoingTxBatchesResponse): unknown {
    const obj: any = {};
    if (message.batches) {
      obj.batches = message.batches.map((e) =>
        e ? OutgoingTxBatch.toJSON(e) : undefined
      );
    } else {
      obj.batches = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOutgoingTxBatchesResponse>
  ): QueryOutgoingTxBatchesResponse {
    const message = {
      ...baseQueryOutgoingTxBatchesResponse,
    } as QueryOutgoingTxBatchesResponse;
    message.batches = [];
    if (object.batches !== undefined && object.batches !== null) {
      for (const e of object.batches) {
        message.batches.push(OutgoingTxBatch.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryOutgoingLogicCallsRequest: object = {};

export const QueryOutgoingLogicCallsRequest = {
  encode(
    _: QueryOutgoingLogicCallsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOutgoingLogicCallsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOutgoingLogicCallsRequest,
    } as QueryOutgoingLogicCallsRequest;
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

  fromJSON(_: any): QueryOutgoingLogicCallsRequest {
    const message = {
      ...baseQueryOutgoingLogicCallsRequest,
    } as QueryOutgoingLogicCallsRequest;
    return message;
  },

  toJSON(_: QueryOutgoingLogicCallsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryOutgoingLogicCallsRequest>
  ): QueryOutgoingLogicCallsRequest {
    const message = {
      ...baseQueryOutgoingLogicCallsRequest,
    } as QueryOutgoingLogicCallsRequest;
    return message;
  },
};

const baseQueryOutgoingLogicCallsResponse: object = {};

export const QueryOutgoingLogicCallsResponse = {
  encode(
    message: QueryOutgoingLogicCallsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.calls) {
      OutgoingLogicCall.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOutgoingLogicCallsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOutgoingLogicCallsResponse,
    } as QueryOutgoingLogicCallsResponse;
    message.calls = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.calls.push(OutgoingLogicCall.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOutgoingLogicCallsResponse {
    const message = {
      ...baseQueryOutgoingLogicCallsResponse,
    } as QueryOutgoingLogicCallsResponse;
    message.calls = [];
    if (object.calls !== undefined && object.calls !== null) {
      for (const e of object.calls) {
        message.calls.push(OutgoingLogicCall.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryOutgoingLogicCallsResponse): unknown {
    const obj: any = {};
    if (message.calls) {
      obj.calls = message.calls.map((e) =>
        e ? OutgoingLogicCall.toJSON(e) : undefined
      );
    } else {
      obj.calls = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOutgoingLogicCallsResponse>
  ): QueryOutgoingLogicCallsResponse {
    const message = {
      ...baseQueryOutgoingLogicCallsResponse,
    } as QueryOutgoingLogicCallsResponse;
    message.calls = [];
    if (object.calls !== undefined && object.calls !== null) {
      for (const e of object.calls) {
        message.calls.push(OutgoingLogicCall.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryBatchRequestByNonceRequest: object = {
  nonce: 0,
  contract_address: "",
};

export const QueryBatchRequestByNonceRequest = {
  encode(
    message: QueryBatchRequestByNonceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    if (message.contract_address !== "") {
      writer.uint32(18).string(message.contract_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryBatchRequestByNonceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryBatchRequestByNonceRequest,
    } as QueryBatchRequestByNonceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.contract_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBatchRequestByNonceRequest {
    const message = {
      ...baseQueryBatchRequestByNonceRequest,
    } as QueryBatchRequestByNonceRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = String(object.contract_address);
    } else {
      message.contract_address = "";
    }
    return message;
  },

  toJSON(message: QueryBatchRequestByNonceRequest): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.contract_address !== undefined &&
      (obj.contract_address = message.contract_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBatchRequestByNonceRequest>
  ): QueryBatchRequestByNonceRequest {
    const message = {
      ...baseQueryBatchRequestByNonceRequest,
    } as QueryBatchRequestByNonceRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = object.contract_address;
    } else {
      message.contract_address = "";
    }
    return message;
  },
};

const baseQueryBatchRequestByNonceResponse: object = {};

export const QueryBatchRequestByNonceResponse = {
  encode(
    message: QueryBatchRequestByNonceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.batch !== undefined) {
      OutgoingTxBatch.encode(message.batch, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryBatchRequestByNonceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryBatchRequestByNonceResponse,
    } as QueryBatchRequestByNonceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batch = OutgoingTxBatch.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBatchRequestByNonceResponse {
    const message = {
      ...baseQueryBatchRequestByNonceResponse,
    } as QueryBatchRequestByNonceResponse;
    if (object.batch !== undefined && object.batch !== null) {
      message.batch = OutgoingTxBatch.fromJSON(object.batch);
    } else {
      message.batch = undefined;
    }
    return message;
  },

  toJSON(message: QueryBatchRequestByNonceResponse): unknown {
    const obj: any = {};
    message.batch !== undefined &&
      (obj.batch = message.batch
        ? OutgoingTxBatch.toJSON(message.batch)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBatchRequestByNonceResponse>
  ): QueryBatchRequestByNonceResponse {
    const message = {
      ...baseQueryBatchRequestByNonceResponse,
    } as QueryBatchRequestByNonceResponse;
    if (object.batch !== undefined && object.batch !== null) {
      message.batch = OutgoingTxBatch.fromPartial(object.batch);
    } else {
      message.batch = undefined;
    }
    return message;
  },
};

const baseQueryBatchConfirmsRequest: object = {
  nonce: 0,
  contract_address: "",
};

export const QueryBatchConfirmsRequest = {
  encode(
    message: QueryBatchConfirmsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nonce !== 0) {
      writer.uint32(8).uint64(message.nonce);
    }
    if (message.contract_address !== "") {
      writer.uint32(18).string(message.contract_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryBatchConfirmsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryBatchConfirmsRequest,
    } as QueryBatchConfirmsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.contract_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBatchConfirmsRequest {
    const message = {
      ...baseQueryBatchConfirmsRequest,
    } as QueryBatchConfirmsRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = String(object.contract_address);
    } else {
      message.contract_address = "";
    }
    return message;
  },

  toJSON(message: QueryBatchConfirmsRequest): unknown {
    const obj: any = {};
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.contract_address !== undefined &&
      (obj.contract_address = message.contract_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBatchConfirmsRequest>
  ): QueryBatchConfirmsRequest {
    const message = {
      ...baseQueryBatchConfirmsRequest,
    } as QueryBatchConfirmsRequest;
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = object.contract_address;
    } else {
      message.contract_address = "";
    }
    return message;
  },
};

const baseQueryBatchConfirmsResponse: object = {};

export const QueryBatchConfirmsResponse = {
  encode(
    message: QueryBatchConfirmsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.confirms) {
      MsgConfirmBatch.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryBatchConfirmsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryBatchConfirmsResponse,
    } as QueryBatchConfirmsResponse;
    message.confirms = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirms.push(
            MsgConfirmBatch.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBatchConfirmsResponse {
    const message = {
      ...baseQueryBatchConfirmsResponse,
    } as QueryBatchConfirmsResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgConfirmBatch.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryBatchConfirmsResponse): unknown {
    const obj: any = {};
    if (message.confirms) {
      obj.confirms = message.confirms.map((e) =>
        e ? MsgConfirmBatch.toJSON(e) : undefined
      );
    } else {
      obj.confirms = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBatchConfirmsResponse>
  ): QueryBatchConfirmsResponse {
    const message = {
      ...baseQueryBatchConfirmsResponse,
    } as QueryBatchConfirmsResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgConfirmBatch.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLogicConfirmsRequest: object = { invalidation_nonce: 0 };

export const QueryLogicConfirmsRequest = {
  encode(
    message: QueryLogicConfirmsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.invalidation_id.length !== 0) {
      writer.uint32(10).bytes(message.invalidation_id);
    }
    if (message.invalidation_nonce !== 0) {
      writer.uint32(16).uint64(message.invalidation_nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLogicConfirmsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLogicConfirmsRequest,
    } as QueryLogicConfirmsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invalidation_id = reader.bytes();
          break;
        case 2:
          message.invalidation_nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLogicConfirmsRequest {
    const message = {
      ...baseQueryLogicConfirmsRequest,
    } as QueryLogicConfirmsRequest;
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
    return message;
  },

  toJSON(message: QueryLogicConfirmsRequest): unknown {
    const obj: any = {};
    message.invalidation_id !== undefined &&
      (obj.invalidation_id = base64FromBytes(
        message.invalidation_id !== undefined
          ? message.invalidation_id
          : new Uint8Array()
      ));
    message.invalidation_nonce !== undefined &&
      (obj.invalidation_nonce = message.invalidation_nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLogicConfirmsRequest>
  ): QueryLogicConfirmsRequest {
    const message = {
      ...baseQueryLogicConfirmsRequest,
    } as QueryLogicConfirmsRequest;
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
    return message;
  },
};

const baseQueryLogicConfirmsResponse: object = {};

export const QueryLogicConfirmsResponse = {
  encode(
    message: QueryLogicConfirmsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.confirms) {
      MsgConfirmLogicCall.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLogicConfirmsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLogicConfirmsResponse,
    } as QueryLogicConfirmsResponse;
    message.confirms = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirms.push(
            MsgConfirmLogicCall.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLogicConfirmsResponse {
    const message = {
      ...baseQueryLogicConfirmsResponse,
    } as QueryLogicConfirmsResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgConfirmLogicCall.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryLogicConfirmsResponse): unknown {
    const obj: any = {};
    if (message.confirms) {
      obj.confirms = message.confirms.map((e) =>
        e ? MsgConfirmLogicCall.toJSON(e) : undefined
      );
    } else {
      obj.confirms = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLogicConfirmsResponse>
  ): QueryLogicConfirmsResponse {
    const message = {
      ...baseQueryLogicConfirmsResponse,
    } as QueryLogicConfirmsResponse;
    message.confirms = [];
    if (object.confirms !== undefined && object.confirms !== null) {
      for (const e of object.confirms) {
        message.confirms.push(MsgConfirmLogicCall.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryLastEventNonceByAddrRequest: object = { address: "" };

export const QueryLastEventNonceByAddrRequest = {
  encode(
    message: QueryLastEventNonceByAddrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastEventNonceByAddrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastEventNonceByAddrRequest,
    } as QueryLastEventNonceByAddrRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastEventNonceByAddrRequest {
    const message = {
      ...baseQueryLastEventNonceByAddrRequest,
    } as QueryLastEventNonceByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryLastEventNonceByAddrRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastEventNonceByAddrRequest>
  ): QueryLastEventNonceByAddrRequest {
    const message = {
      ...baseQueryLastEventNonceByAddrRequest,
    } as QueryLastEventNonceByAddrRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryLastEventNonceByAddrResponse: object = { event_nonce: 0 };

export const QueryLastEventNonceByAddrResponse = {
  encode(
    message: QueryLastEventNonceByAddrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.event_nonce !== 0) {
      writer.uint32(8).uint64(message.event_nonce);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastEventNonceByAddrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastEventNonceByAddrResponse,
    } as QueryLastEventNonceByAddrResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.event_nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastEventNonceByAddrResponse {
    const message = {
      ...baseQueryLastEventNonceByAddrResponse,
    } as QueryLastEventNonceByAddrResponse;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = Number(object.event_nonce);
    } else {
      message.event_nonce = 0;
    }
    return message;
  },

  toJSON(message: QueryLastEventNonceByAddrResponse): unknown {
    const obj: any = {};
    message.event_nonce !== undefined &&
      (obj.event_nonce = message.event_nonce);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastEventNonceByAddrResponse>
  ): QueryLastEventNonceByAddrResponse {
    const message = {
      ...baseQueryLastEventNonceByAddrResponse,
    } as QueryLastEventNonceByAddrResponse;
    if (object.event_nonce !== undefined && object.event_nonce !== null) {
      message.event_nonce = object.event_nonce;
    } else {
      message.event_nonce = 0;
    }
    return message;
  },
};

const baseQueryERC20ToDenomRequest: object = { erc20: "" };

export const QueryERC20ToDenomRequest = {
  encode(
    message: QueryERC20ToDenomRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erc20 !== "") {
      writer.uint32(10).string(message.erc20);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryERC20ToDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryERC20ToDenomRequest,
    } as QueryERC20ToDenomRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erc20 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryERC20ToDenomRequest {
    const message = {
      ...baseQueryERC20ToDenomRequest,
    } as QueryERC20ToDenomRequest;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = String(object.erc20);
    } else {
      message.erc20 = "";
    }
    return message;
  },

  toJSON(message: QueryERC20ToDenomRequest): unknown {
    const obj: any = {};
    message.erc20 !== undefined && (obj.erc20 = message.erc20);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryERC20ToDenomRequest>
  ): QueryERC20ToDenomRequest {
    const message = {
      ...baseQueryERC20ToDenomRequest,
    } as QueryERC20ToDenomRequest;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = object.erc20;
    } else {
      message.erc20 = "";
    }
    return message;
  },
};

const baseQueryERC20ToDenomResponse: object = {
  denom: "",
  cosmos_originated: false,
};

export const QueryERC20ToDenomResponse = {
  encode(
    message: QueryERC20ToDenomResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.cosmos_originated === true) {
      writer.uint32(16).bool(message.cosmos_originated);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryERC20ToDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryERC20ToDenomResponse,
    } as QueryERC20ToDenomResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.cosmos_originated = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryERC20ToDenomResponse {
    const message = {
      ...baseQueryERC20ToDenomResponse,
    } as QueryERC20ToDenomResponse;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (
      object.cosmos_originated !== undefined &&
      object.cosmos_originated !== null
    ) {
      message.cosmos_originated = Boolean(object.cosmos_originated);
    } else {
      message.cosmos_originated = false;
    }
    return message;
  },

  toJSON(message: QueryERC20ToDenomResponse): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.cosmos_originated !== undefined &&
      (obj.cosmos_originated = message.cosmos_originated);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryERC20ToDenomResponse>
  ): QueryERC20ToDenomResponse {
    const message = {
      ...baseQueryERC20ToDenomResponse,
    } as QueryERC20ToDenomResponse;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (
      object.cosmos_originated !== undefined &&
      object.cosmos_originated !== null
    ) {
      message.cosmos_originated = object.cosmos_originated;
    } else {
      message.cosmos_originated = false;
    }
    return message;
  },
};

const baseQueryDenomToERC20Request: object = { denom: "" };

export const QueryDenomToERC20Request = {
  encode(
    message: QueryDenomToERC20Request,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDenomToERC20Request {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDenomToERC20Request,
    } as QueryDenomToERC20Request;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDenomToERC20Request {
    const message = {
      ...baseQueryDenomToERC20Request,
    } as QueryDenomToERC20Request;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: QueryDenomToERC20Request): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDenomToERC20Request>
  ): QueryDenomToERC20Request {
    const message = {
      ...baseQueryDenomToERC20Request,
    } as QueryDenomToERC20Request;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseQueryDenomToERC20Response: object = {
  erc20: "",
  cosmos_originated: false,
};

export const QueryDenomToERC20Response = {
  encode(
    message: QueryDenomToERC20Response,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erc20 !== "") {
      writer.uint32(10).string(message.erc20);
    }
    if (message.cosmos_originated === true) {
      writer.uint32(16).bool(message.cosmos_originated);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDenomToERC20Response {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDenomToERC20Response,
    } as QueryDenomToERC20Response;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erc20 = reader.string();
          break;
        case 2:
          message.cosmos_originated = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDenomToERC20Response {
    const message = {
      ...baseQueryDenomToERC20Response,
    } as QueryDenomToERC20Response;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = String(object.erc20);
    } else {
      message.erc20 = "";
    }
    if (
      object.cosmos_originated !== undefined &&
      object.cosmos_originated !== null
    ) {
      message.cosmos_originated = Boolean(object.cosmos_originated);
    } else {
      message.cosmos_originated = false;
    }
    return message;
  },

  toJSON(message: QueryDenomToERC20Response): unknown {
    const obj: any = {};
    message.erc20 !== undefined && (obj.erc20 = message.erc20);
    message.cosmos_originated !== undefined &&
      (obj.cosmos_originated = message.cosmos_originated);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDenomToERC20Response>
  ): QueryDenomToERC20Response {
    const message = {
      ...baseQueryDenomToERC20Response,
    } as QueryDenomToERC20Response;
    if (object.erc20 !== undefined && object.erc20 !== null) {
      message.erc20 = object.erc20;
    } else {
      message.erc20 = "";
    }
    if (
      object.cosmos_originated !== undefined &&
      object.cosmos_originated !== null
    ) {
      message.cosmos_originated = object.cosmos_originated;
    } else {
      message.cosmos_originated = false;
    }
    return message;
  },
};

const baseQueryAttestationsRequest: object = {
  limit: 0,
  order_by: "",
  claim_type: "",
  nonce: 0,
  height: 0,
};

export const QueryAttestationsRequest = {
  encode(
    message: QueryAttestationsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limit !== 0) {
      writer.uint32(8).uint64(message.limit);
    }
    if (message.order_by !== "") {
      writer.uint32(18).string(message.order_by);
    }
    if (message.claim_type !== "") {
      writer.uint32(26).string(message.claim_type);
    }
    if (message.nonce !== 0) {
      writer.uint32(32).uint64(message.nonce);
    }
    if (message.height !== 0) {
      writer.uint32(40).uint64(message.height);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAttestationsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAttestationsRequest,
    } as QueryAttestationsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limit = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.order_by = reader.string();
          break;
        case 3:
          message.claim_type = reader.string();
          break;
        case 4:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAttestationsRequest {
    const message = {
      ...baseQueryAttestationsRequest,
    } as QueryAttestationsRequest;
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Number(object.limit);
    } else {
      message.limit = 0;
    }
    if (object.order_by !== undefined && object.order_by !== null) {
      message.order_by = String(object.order_by);
    } else {
      message.order_by = "";
    }
    if (object.claim_type !== undefined && object.claim_type !== null) {
      message.claim_type = String(object.claim_type);
    } else {
      message.claim_type = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height);
    } else {
      message.height = 0;
    }
    return message;
  },

  toJSON(message: QueryAttestationsRequest): unknown {
    const obj: any = {};
    message.limit !== undefined && (obj.limit = message.limit);
    message.order_by !== undefined && (obj.order_by = message.order_by);
    message.claim_type !== undefined && (obj.claim_type = message.claim_type);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.height !== undefined && (obj.height = message.height);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAttestationsRequest>
  ): QueryAttestationsRequest {
    const message = {
      ...baseQueryAttestationsRequest,
    } as QueryAttestationsRequest;
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit;
    } else {
      message.limit = 0;
    }
    if (object.order_by !== undefined && object.order_by !== null) {
      message.order_by = object.order_by;
    } else {
      message.order_by = "";
    }
    if (object.claim_type !== undefined && object.claim_type !== null) {
      message.claim_type = object.claim_type;
    } else {
      message.claim_type = "";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = 0;
    }
    return message;
  },
};

const baseQueryAttestationsResponse: object = {};

export const QueryAttestationsResponse = {
  encode(
    message: QueryAttestationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.attestations) {
      Attestation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAttestationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAttestationsResponse,
    } as QueryAttestationsResponse;
    message.attestations = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.attestations.push(
            Attestation.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAttestationsResponse {
    const message = {
      ...baseQueryAttestationsResponse,
    } as QueryAttestationsResponse;
    message.attestations = [];
    if (object.attestations !== undefined && object.attestations !== null) {
      for (const e of object.attestations) {
        message.attestations.push(Attestation.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryAttestationsResponse): unknown {
    const obj: any = {};
    if (message.attestations) {
      obj.attestations = message.attestations.map((e) =>
        e ? Attestation.toJSON(e) : undefined
      );
    } else {
      obj.attestations = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAttestationsResponse>
  ): QueryAttestationsResponse {
    const message = {
      ...baseQueryAttestationsResponse,
    } as QueryAttestationsResponse;
    message.attestations = [];
    if (object.attestations !== undefined && object.attestations !== null) {
      for (const e of object.attestations) {
        message.attestations.push(Attestation.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryDelegateKeysByValidatorAddress: object = {
  validator_address: "",
};

export const QueryDelegateKeysByValidatorAddress = {
  encode(
    message: QueryDelegateKeysByValidatorAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.validator_address !== "") {
      writer.uint32(10).string(message.validator_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByValidatorAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByValidatorAddress,
    } as QueryDelegateKeysByValidatorAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByValidatorAddress {
    const message = {
      ...baseQueryDelegateKeysByValidatorAddress,
    } as QueryDelegateKeysByValidatorAddress;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByValidatorAddress): unknown {
    const obj: any = {};
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByValidatorAddress>
  ): QueryDelegateKeysByValidatorAddress {
    const message = {
      ...baseQueryDelegateKeysByValidatorAddress,
    } as QueryDelegateKeysByValidatorAddress;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    return message;
  },
};

const baseQueryDelegateKeysByValidatorAddressResponse: object = {
  eth_address: "",
  orchestrator_address: "",
};

export const QueryDelegateKeysByValidatorAddressResponse = {
  encode(
    message: QueryDelegateKeysByValidatorAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.eth_address !== "") {
      writer.uint32(10).string(message.eth_address);
    }
    if (message.orchestrator_address !== "") {
      writer.uint32(18).string(message.orchestrator_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByValidatorAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByValidatorAddressResponse,
    } as QueryDelegateKeysByValidatorAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.eth_address = reader.string();
          break;
        case 2:
          message.orchestrator_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByValidatorAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByValidatorAddressResponse,
    } as QueryDelegateKeysByValidatorAddressResponse;
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = String(object.eth_address);
    } else {
      message.eth_address = "";
    }
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = String(object.orchestrator_address);
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByValidatorAddressResponse): unknown {
    const obj: any = {};
    message.eth_address !== undefined &&
      (obj.eth_address = message.eth_address);
    message.orchestrator_address !== undefined &&
      (obj.orchestrator_address = message.orchestrator_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByValidatorAddressResponse>
  ): QueryDelegateKeysByValidatorAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByValidatorAddressResponse,
    } as QueryDelegateKeysByValidatorAddressResponse;
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = object.eth_address;
    } else {
      message.eth_address = "";
    }
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = object.orchestrator_address;
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },
};

const baseQueryDelegateKeysByEthAddress: object = { eth_address: "" };

export const QueryDelegateKeysByEthAddress = {
  encode(
    message: QueryDelegateKeysByEthAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.eth_address !== "") {
      writer.uint32(10).string(message.eth_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByEthAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByEthAddress,
    } as QueryDelegateKeysByEthAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.eth_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByEthAddress {
    const message = {
      ...baseQueryDelegateKeysByEthAddress,
    } as QueryDelegateKeysByEthAddress;
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = String(object.eth_address);
    } else {
      message.eth_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByEthAddress): unknown {
    const obj: any = {};
    message.eth_address !== undefined &&
      (obj.eth_address = message.eth_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByEthAddress>
  ): QueryDelegateKeysByEthAddress {
    const message = {
      ...baseQueryDelegateKeysByEthAddress,
    } as QueryDelegateKeysByEthAddress;
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = object.eth_address;
    } else {
      message.eth_address = "";
    }
    return message;
  },
};

const baseQueryDelegateKeysByEthAddressResponse: object = {
  validator_address: "",
  orchestrator_address: "",
};

export const QueryDelegateKeysByEthAddressResponse = {
  encode(
    message: QueryDelegateKeysByEthAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.validator_address !== "") {
      writer.uint32(10).string(message.validator_address);
    }
    if (message.orchestrator_address !== "") {
      writer.uint32(18).string(message.orchestrator_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByEthAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByEthAddressResponse,
    } as QueryDelegateKeysByEthAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator_address = reader.string();
          break;
        case 2:
          message.orchestrator_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByEthAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByEthAddressResponse,
    } as QueryDelegateKeysByEthAddressResponse;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = String(object.orchestrator_address);
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByEthAddressResponse): unknown {
    const obj: any = {};
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    message.orchestrator_address !== undefined &&
      (obj.orchestrator_address = message.orchestrator_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByEthAddressResponse>
  ): QueryDelegateKeysByEthAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByEthAddressResponse,
    } as QueryDelegateKeysByEthAddressResponse;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = object.orchestrator_address;
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },
};

const baseQueryDelegateKeysByOrchestratorAddress: object = {
  orchestrator_address: "",
};

export const QueryDelegateKeysByOrchestratorAddress = {
  encode(
    message: QueryDelegateKeysByOrchestratorAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.orchestrator_address !== "") {
      writer.uint32(10).string(message.orchestrator_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByOrchestratorAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddress,
    } as QueryDelegateKeysByOrchestratorAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orchestrator_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByOrchestratorAddress {
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddress,
    } as QueryDelegateKeysByOrchestratorAddress;
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = String(object.orchestrator_address);
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByOrchestratorAddress): unknown {
    const obj: any = {};
    message.orchestrator_address !== undefined &&
      (obj.orchestrator_address = message.orchestrator_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByOrchestratorAddress>
  ): QueryDelegateKeysByOrchestratorAddress {
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddress,
    } as QueryDelegateKeysByOrchestratorAddress;
    if (
      object.orchestrator_address !== undefined &&
      object.orchestrator_address !== null
    ) {
      message.orchestrator_address = object.orchestrator_address;
    } else {
      message.orchestrator_address = "";
    }
    return message;
  },
};

const baseQueryDelegateKeysByOrchestratorAddressResponse: object = {
  validator_address: "",
  eth_address: "",
};

export const QueryDelegateKeysByOrchestratorAddressResponse = {
  encode(
    message: QueryDelegateKeysByOrchestratorAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.validator_address !== "") {
      writer.uint32(10).string(message.validator_address);
    }
    if (message.eth_address !== "") {
      writer.uint32(18).string(message.eth_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDelegateKeysByOrchestratorAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddressResponse,
    } as QueryDelegateKeysByOrchestratorAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator_address = reader.string();
          break;
        case 2:
          message.eth_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDelegateKeysByOrchestratorAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddressResponse,
    } as QueryDelegateKeysByOrchestratorAddressResponse;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = String(object.eth_address);
    } else {
      message.eth_address = "";
    }
    return message;
  },

  toJSON(message: QueryDelegateKeysByOrchestratorAddressResponse): unknown {
    const obj: any = {};
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    message.eth_address !== undefined &&
      (obj.eth_address = message.eth_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDelegateKeysByOrchestratorAddressResponse>
  ): QueryDelegateKeysByOrchestratorAddressResponse {
    const message = {
      ...baseQueryDelegateKeysByOrchestratorAddressResponse,
    } as QueryDelegateKeysByOrchestratorAddressResponse;
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    if (object.eth_address !== undefined && object.eth_address !== null) {
      message.eth_address = object.eth_address;
    } else {
      message.eth_address = "";
    }
    return message;
  },
};

const baseQueryPendingSendToEth: object = { sender_address: "" };

export const QueryPendingSendToEth = {
  encode(
    message: QueryPendingSendToEth,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender_address !== "") {
      writer.uint32(10).string(message.sender_address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPendingSendToEth {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryPendingSendToEth } as QueryPendingSendToEth;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingSendToEth {
    const message = { ...baseQueryPendingSendToEth } as QueryPendingSendToEth;
    if (object.sender_address !== undefined && object.sender_address !== null) {
      message.sender_address = String(object.sender_address);
    } else {
      message.sender_address = "";
    }
    return message;
  },

  toJSON(message: QueryPendingSendToEth): unknown {
    const obj: any = {};
    message.sender_address !== undefined &&
      (obj.sender_address = message.sender_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryPendingSendToEth>
  ): QueryPendingSendToEth {
    const message = { ...baseQueryPendingSendToEth } as QueryPendingSendToEth;
    if (object.sender_address !== undefined && object.sender_address !== null) {
      message.sender_address = object.sender_address;
    } else {
      message.sender_address = "";
    }
    return message;
  },
};

const baseQueryPendingSendToEthResponse: object = {};

export const QueryPendingSendToEthResponse = {
  encode(
    message: QueryPendingSendToEthResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.transfers_in_batches) {
      OutgoingTransferTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.unbatched_transfers) {
      OutgoingTransferTx.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryPendingSendToEthResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryPendingSendToEthResponse,
    } as QueryPendingSendToEthResponse;
    message.transfers_in_batches = [];
    message.unbatched_transfers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transfers_in_batches.push(
            OutgoingTransferTx.decode(reader, reader.uint32())
          );
          break;
        case 2:
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

  fromJSON(object: any): QueryPendingSendToEthResponse {
    const message = {
      ...baseQueryPendingSendToEthResponse,
    } as QueryPendingSendToEthResponse;
    message.transfers_in_batches = [];
    message.unbatched_transfers = [];
    if (
      object.transfers_in_batches !== undefined &&
      object.transfers_in_batches !== null
    ) {
      for (const e of object.transfers_in_batches) {
        message.transfers_in_batches.push(OutgoingTransferTx.fromJSON(e));
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

  toJSON(message: QueryPendingSendToEthResponse): unknown {
    const obj: any = {};
    if (message.transfers_in_batches) {
      obj.transfers_in_batches = message.transfers_in_batches.map((e) =>
        e ? OutgoingTransferTx.toJSON(e) : undefined
      );
    } else {
      obj.transfers_in_batches = [];
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

  fromPartial(
    object: DeepPartial<QueryPendingSendToEthResponse>
  ): QueryPendingSendToEthResponse {
    const message = {
      ...baseQueryPendingSendToEthResponse,
    } as QueryPendingSendToEthResponse;
    message.transfers_in_batches = [];
    message.unbatched_transfers = [];
    if (
      object.transfers_in_batches !== undefined &&
      object.transfers_in_batches !== null
    ) {
      for (const e of object.transfers_in_batches) {
        message.transfers_in_batches.push(OutgoingTransferTx.fromPartial(e));
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

const baseQueryPendingIbcAutoForwards: object = { limit: 0 };

export const QueryPendingIbcAutoForwards = {
  encode(
    message: QueryPendingIbcAutoForwards,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limit !== 0) {
      writer.uint32(8).uint64(message.limit);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryPendingIbcAutoForwards {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryPendingIbcAutoForwards,
    } as QueryPendingIbcAutoForwards;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limit = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingIbcAutoForwards {
    const message = {
      ...baseQueryPendingIbcAutoForwards,
    } as QueryPendingIbcAutoForwards;
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Number(object.limit);
    } else {
      message.limit = 0;
    }
    return message;
  },

  toJSON(message: QueryPendingIbcAutoForwards): unknown {
    const obj: any = {};
    message.limit !== undefined && (obj.limit = message.limit);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryPendingIbcAutoForwards>
  ): QueryPendingIbcAutoForwards {
    const message = {
      ...baseQueryPendingIbcAutoForwards,
    } as QueryPendingIbcAutoForwards;
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit;
    } else {
      message.limit = 0;
    }
    return message;
  },
};

const baseQueryPendingIbcAutoForwardsResponse: object = {};

export const QueryPendingIbcAutoForwardsResponse = {
  encode(
    message: QueryPendingIbcAutoForwardsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pending_ibc_auto_forwards) {
      PendingIbcAutoForward.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryPendingIbcAutoForwardsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryPendingIbcAutoForwardsResponse,
    } as QueryPendingIbcAutoForwardsResponse;
    message.pending_ibc_auto_forwards = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pending_ibc_auto_forwards.push(
            PendingIbcAutoForward.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingIbcAutoForwardsResponse {
    const message = {
      ...baseQueryPendingIbcAutoForwardsResponse,
    } as QueryPendingIbcAutoForwardsResponse;
    message.pending_ibc_auto_forwards = [];
    if (
      object.pending_ibc_auto_forwards !== undefined &&
      object.pending_ibc_auto_forwards !== null
    ) {
      for (const e of object.pending_ibc_auto_forwards) {
        message.pending_ibc_auto_forwards.push(
          PendingIbcAutoForward.fromJSON(e)
        );
      }
    }
    return message;
  },

  toJSON(message: QueryPendingIbcAutoForwardsResponse): unknown {
    const obj: any = {};
    if (message.pending_ibc_auto_forwards) {
      obj.pending_ibc_auto_forwards = message.pending_ibc_auto_forwards.map(
        (e) => (e ? PendingIbcAutoForward.toJSON(e) : undefined)
      );
    } else {
      obj.pending_ibc_auto_forwards = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryPendingIbcAutoForwardsResponse>
  ): QueryPendingIbcAutoForwardsResponse {
    const message = {
      ...baseQueryPendingIbcAutoForwardsResponse,
    } as QueryPendingIbcAutoForwardsResponse;
    message.pending_ibc_auto_forwards = [];
    if (
      object.pending_ibc_auto_forwards !== undefined &&
      object.pending_ibc_auto_forwards !== null
    ) {
      for (const e of object.pending_ibc_auto_forwards) {
        message.pending_ibc_auto_forwards.push(
          PendingIbcAutoForward.fromPartial(e)
        );
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service */
export interface Query {
  /** Deployments queries deployments */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  CurrentValset(
    request: QueryCurrentValsetRequest
  ): Promise<QueryCurrentValsetResponse>;
  ValsetRequest(
    request: QueryValsetRequestRequest
  ): Promise<QueryValsetRequestResponse>;
  ValsetConfirm(
    request: QueryValsetConfirmRequest
  ): Promise<QueryValsetConfirmResponse>;
  ValsetConfirmsByNonce(
    request: QueryValsetConfirmsByNonceRequest
  ): Promise<QueryValsetConfirmsByNonceResponse>;
  LastValsetRequests(
    request: QueryLastValsetRequestsRequest
  ): Promise<QueryLastValsetRequestsResponse>;
  LastPendingValsetRequestByAddr(
    request: QueryLastPendingValsetRequestByAddrRequest
  ): Promise<QueryLastPendingValsetRequestByAddrResponse>;
  LastPendingBatchRequestByAddr(
    request: QueryLastPendingBatchRequestByAddrRequest
  ): Promise<QueryLastPendingBatchRequestByAddrResponse>;
  LastPendingLogicCallByAddr(
    request: QueryLastPendingLogicCallByAddrRequest
  ): Promise<QueryLastPendingLogicCallByAddrResponse>;
  LastEventNonceByAddr(
    request: QueryLastEventNonceByAddrRequest
  ): Promise<QueryLastEventNonceByAddrResponse>;
  BatchFees(request: QueryBatchFeeRequest): Promise<QueryBatchFeeResponse>;
  OutgoingTxBatches(
    request: QueryOutgoingTxBatchesRequest
  ): Promise<QueryOutgoingTxBatchesResponse>;
  OutgoingLogicCalls(
    request: QueryOutgoingLogicCallsRequest
  ): Promise<QueryOutgoingLogicCallsResponse>;
  BatchRequestByNonce(
    request: QueryBatchRequestByNonceRequest
  ): Promise<QueryBatchRequestByNonceResponse>;
  BatchConfirms(
    request: QueryBatchConfirmsRequest
  ): Promise<QueryBatchConfirmsResponse>;
  LogicConfirms(
    request: QueryLogicConfirmsRequest
  ): Promise<QueryLogicConfirmsResponse>;
  ERC20ToDenom(
    request: QueryERC20ToDenomRequest
  ): Promise<QueryERC20ToDenomResponse>;
  DenomToERC20(
    request: QueryDenomToERC20Request
  ): Promise<QueryDenomToERC20Response>;
  GetAttestations(
    request: QueryAttestationsRequest
  ): Promise<QueryAttestationsResponse>;
  GetDelegateKeyByValidator(
    request: QueryDelegateKeysByValidatorAddress
  ): Promise<QueryDelegateKeysByValidatorAddressResponse>;
  GetDelegateKeyByEth(
    request: QueryDelegateKeysByEthAddress
  ): Promise<QueryDelegateKeysByEthAddressResponse>;
  GetDelegateKeyByOrchestrator(
    request: QueryDelegateKeysByOrchestratorAddress
  ): Promise<QueryDelegateKeysByOrchestratorAddressResponse>;
  GetPendingSendToEth(
    request: QueryPendingSendToEth
  ): Promise<QueryPendingSendToEthResponse>;
  GetPendingIbcAutoForwards(
    request: QueryPendingIbcAutoForwards
  ): Promise<QueryPendingIbcAutoForwardsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  CurrentValset(
    request: QueryCurrentValsetRequest
  ): Promise<QueryCurrentValsetResponse> {
    const data = QueryCurrentValsetRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "CurrentValset", data);
    return promise.then((data) =>
      QueryCurrentValsetResponse.decode(new Reader(data))
    );
  }

  ValsetRequest(
    request: QueryValsetRequestRequest
  ): Promise<QueryValsetRequestResponse> {
    const data = QueryValsetRequestRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "ValsetRequest", data);
    return promise.then((data) =>
      QueryValsetRequestResponse.decode(new Reader(data))
    );
  }

  ValsetConfirm(
    request: QueryValsetConfirmRequest
  ): Promise<QueryValsetConfirmResponse> {
    const data = QueryValsetConfirmRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "ValsetConfirm", data);
    return promise.then((data) =>
      QueryValsetConfirmResponse.decode(new Reader(data))
    );
  }

  ValsetConfirmsByNonce(
    request: QueryValsetConfirmsByNonceRequest
  ): Promise<QueryValsetConfirmsByNonceResponse> {
    const data = QueryValsetConfirmsByNonceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "ValsetConfirmsByNonce",
      data
    );
    return promise.then((data) =>
      QueryValsetConfirmsByNonceResponse.decode(new Reader(data))
    );
  }

  LastValsetRequests(
    request: QueryLastValsetRequestsRequest
  ): Promise<QueryLastValsetRequestsResponse> {
    const data = QueryLastValsetRequestsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "LastValsetRequests",
      data
    );
    return promise.then((data) =>
      QueryLastValsetRequestsResponse.decode(new Reader(data))
    );
  }

  LastPendingValsetRequestByAddr(
    request: QueryLastPendingValsetRequestByAddrRequest
  ): Promise<QueryLastPendingValsetRequestByAddrResponse> {
    const data = QueryLastPendingValsetRequestByAddrRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "LastPendingValsetRequestByAddr",
      data
    );
    return promise.then((data) =>
      QueryLastPendingValsetRequestByAddrResponse.decode(new Reader(data))
    );
  }

  LastPendingBatchRequestByAddr(
    request: QueryLastPendingBatchRequestByAddrRequest
  ): Promise<QueryLastPendingBatchRequestByAddrResponse> {
    const data = QueryLastPendingBatchRequestByAddrRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "LastPendingBatchRequestByAddr",
      data
    );
    return promise.then((data) =>
      QueryLastPendingBatchRequestByAddrResponse.decode(new Reader(data))
    );
  }

  LastPendingLogicCallByAddr(
    request: QueryLastPendingLogicCallByAddrRequest
  ): Promise<QueryLastPendingLogicCallByAddrResponse> {
    const data = QueryLastPendingLogicCallByAddrRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "LastPendingLogicCallByAddr",
      data
    );
    return promise.then((data) =>
      QueryLastPendingLogicCallByAddrResponse.decode(new Reader(data))
    );
  }

  LastEventNonceByAddr(
    request: QueryLastEventNonceByAddrRequest
  ): Promise<QueryLastEventNonceByAddrResponse> {
    const data = QueryLastEventNonceByAddrRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "LastEventNonceByAddr",
      data
    );
    return promise.then((data) =>
      QueryLastEventNonceByAddrResponse.decode(new Reader(data))
    );
  }

  BatchFees(request: QueryBatchFeeRequest): Promise<QueryBatchFeeResponse> {
    const data = QueryBatchFeeRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "BatchFees", data);
    return promise.then((data) =>
      QueryBatchFeeResponse.decode(new Reader(data))
    );
  }

  OutgoingTxBatches(
    request: QueryOutgoingTxBatchesRequest
  ): Promise<QueryOutgoingTxBatchesResponse> {
    const data = QueryOutgoingTxBatchesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "OutgoingTxBatches",
      data
    );
    return promise.then((data) =>
      QueryOutgoingTxBatchesResponse.decode(new Reader(data))
    );
  }

  OutgoingLogicCalls(
    request: QueryOutgoingLogicCallsRequest
  ): Promise<QueryOutgoingLogicCallsResponse> {
    const data = QueryOutgoingLogicCallsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "OutgoingLogicCalls",
      data
    );
    return promise.then((data) =>
      QueryOutgoingLogicCallsResponse.decode(new Reader(data))
    );
  }

  BatchRequestByNonce(
    request: QueryBatchRequestByNonceRequest
  ): Promise<QueryBatchRequestByNonceResponse> {
    const data = QueryBatchRequestByNonceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "BatchRequestByNonce",
      data
    );
    return promise.then((data) =>
      QueryBatchRequestByNonceResponse.decode(new Reader(data))
    );
  }

  BatchConfirms(
    request: QueryBatchConfirmsRequest
  ): Promise<QueryBatchConfirmsResponse> {
    const data = QueryBatchConfirmsRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "BatchConfirms", data);
    return promise.then((data) =>
      QueryBatchConfirmsResponse.decode(new Reader(data))
    );
  }

  LogicConfirms(
    request: QueryLogicConfirmsRequest
  ): Promise<QueryLogicConfirmsResponse> {
    const data = QueryLogicConfirmsRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "LogicConfirms", data);
    return promise.then((data) =>
      QueryLogicConfirmsResponse.decode(new Reader(data))
    );
  }

  ERC20ToDenom(
    request: QueryERC20ToDenomRequest
  ): Promise<QueryERC20ToDenomResponse> {
    const data = QueryERC20ToDenomRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "ERC20ToDenom", data);
    return promise.then((data) =>
      QueryERC20ToDenomResponse.decode(new Reader(data))
    );
  }

  DenomToERC20(
    request: QueryDenomToERC20Request
  ): Promise<QueryDenomToERC20Response> {
    const data = QueryDenomToERC20Request.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "DenomToERC20", data);
    return promise.then((data) =>
      QueryDenomToERC20Response.decode(new Reader(data))
    );
  }

  GetAttestations(
    request: QueryAttestationsRequest
  ): Promise<QueryAttestationsResponse> {
    const data = QueryAttestationsRequest.encode(request).finish();
    const promise = this.rpc.request("gravity.Query", "GetAttestations", data);
    return promise.then((data) =>
      QueryAttestationsResponse.decode(new Reader(data))
    );
  }

  GetDelegateKeyByValidator(
    request: QueryDelegateKeysByValidatorAddress
  ): Promise<QueryDelegateKeysByValidatorAddressResponse> {
    const data = QueryDelegateKeysByValidatorAddress.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "GetDelegateKeyByValidator",
      data
    );
    return promise.then((data) =>
      QueryDelegateKeysByValidatorAddressResponse.decode(new Reader(data))
    );
  }

  GetDelegateKeyByEth(
    request: QueryDelegateKeysByEthAddress
  ): Promise<QueryDelegateKeysByEthAddressResponse> {
    const data = QueryDelegateKeysByEthAddress.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "GetDelegateKeyByEth",
      data
    );
    return promise.then((data) =>
      QueryDelegateKeysByEthAddressResponse.decode(new Reader(data))
    );
  }

  GetDelegateKeyByOrchestrator(
    request: QueryDelegateKeysByOrchestratorAddress
  ): Promise<QueryDelegateKeysByOrchestratorAddressResponse> {
    const data = QueryDelegateKeysByOrchestratorAddress.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "GetDelegateKeyByOrchestrator",
      data
    );
    return promise.then((data) =>
      QueryDelegateKeysByOrchestratorAddressResponse.decode(new Reader(data))
    );
  }

  GetPendingSendToEth(
    request: QueryPendingSendToEth
  ): Promise<QueryPendingSendToEthResponse> {
    const data = QueryPendingSendToEth.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "GetPendingSendToEth",
      data
    );
    return promise.then((data) =>
      QueryPendingSendToEthResponse.decode(new Reader(data))
    );
  }

  GetPendingIbcAutoForwards(
    request: QueryPendingIbcAutoForwards
  ): Promise<QueryPendingIbcAutoForwardsResponse> {
    const data = QueryPendingIbcAutoForwards.encode(request).finish();
    const promise = this.rpc.request(
      "gravity.Query",
      "GetPendingIbcAutoForwards",
      data
    );
    return promise.then((data) =>
      QueryPendingIbcAutoForwardsResponse.decode(new Reader(data))
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
