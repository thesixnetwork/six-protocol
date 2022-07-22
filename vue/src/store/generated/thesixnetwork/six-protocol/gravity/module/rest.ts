/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/**
* The actual content of the claims is passed in with the transaction making the claim
and then passed through the call stack alongside the attestation while it is processed
the key in which the attestation is stored is keyed on the exact details of the claim
but there is no reason to store those exact details becuause the next message sender
will kindly provide you with them.
*/
export interface GravityAttestation {
  observed?: boolean;
  votes?: string[];

  /** @format uint64 */
  height?: string;

  /**
   * `Any` contains an arbitrary serialized protocol buffer message along with a
   * URL that describes the type of the serialized message.
   *
   * Protobuf library provides support to pack/unpack Any values in the form
   * of utility functions or additional generated methods of the Any type.
   *
   * Example 1: Pack and unpack a message in C++.
   *
   *     Foo foo = ...;
   *     Any any;
   *     any.PackFrom(foo);
   *     ...
   *     if (any.UnpackTo(&foo)) {
   *       ...
   *     }
   *
   * Example 2: Pack and unpack a message in Java.
   *
   *     Foo foo = ...;
   *     Any any = Any.pack(foo);
   *     ...
   *     if (any.is(Foo.class)) {
   *       foo = any.unpack(Foo.class);
   *     }
   *
   *  Example 3: Pack and unpack a message in Python.
   *
   *     foo = Foo(...)
   *     any = Any()
   *     any.Pack(foo)
   *     ...
   *     if any.Is(Foo.DESCRIPTOR):
   *       any.Unpack(foo)
   *       ...
   *
   *  Example 4: Pack and unpack a message in Go
   *
   *      foo := &pb.Foo{...}
   *      any, err := anypb.New(foo)
   *      if err != nil {
   *        ...
   *      }
   *      ...
   *      foo := &pb.Foo{}
   *      if err := any.UnmarshalTo(foo); err != nil {
   *        ...
   *      }
   *
   * The pack methods provided by protobuf library will by default use
   * 'type.googleapis.com/full.type.name' as the type URL and the unpack
   * methods only use the fully qualified type name after the last '/'
   * in the type URL, for example "foo.bar.com/x/y.z" will yield type
   * name "y.z".
   *
   *
   * JSON
   * ====
   * The JSON representation of an `Any` value uses the regular
   * representation of the deserialized, embedded message, with an
   * additional field `@type` which contains the type URL. Example:
   *
   *     package google.profile;
   *     message Person {
   *       string first_name = 1;
   *       string last_name = 2;
   *     }
   *
   *     {
   *       "@type": "type.googleapis.com/google.profile.Person",
   *       "firstName": <string>,
   *       "lastName": <string>
   *     }
   *
   * If the embedded message type is well-known and has a custom JSON
   * representation, that representation will be embedded adding a field
   * `value` which holds the custom JSON in addition to the `@type`
   * field. Example (for message [google.protobuf.Duration][]):
   *
   *     {
   *       "@type": "type.googleapis.com/google.protobuf.Duration",
   *       "value": "1.212s"
   *     }
   */
  claim?: ProtobufAny;
}

export interface GravityBatchFees {
  token?: string;
  total_fees?: string;

  /** @format uint64 */
  tx_count?: string;
}

export interface GravityBridgeValidator {
  /** @format uint64 */
  power?: string;
  ethereum_address?: string;
}

export interface GravityERC20Token {
  contract?: string;
  amount?: string;
}

export type GravityMsgBatchSendToEthClaimResponse = object;

export type GravityMsgCancelSendToEthResponse = object;

export interface GravityMsgConfirmBatch {
  /** @format uint64 */
  nonce?: string;
  token_contract?: string;
  eth_signer?: string;
  orchestrator?: string;
  signature?: string;
}

export type GravityMsgConfirmBatchResponse = object;

export interface GravityMsgConfirmLogicCall {
  invalidation_id?: string;

  /** @format uint64 */
  invalidation_nonce?: string;
  eth_signer?: string;
  orchestrator?: string;
  signature?: string;
}

export type GravityMsgConfirmLogicCallResponse = object;

export type GravityMsgERC20DeployedClaimResponse = object;

export type GravityMsgExecuteIbcAutoForwardsResponse = object;

export type GravityMsgLogicCallExecutedClaimResponse = object;

export type GravityMsgRequestBatchResponse = object;

export type GravityMsgSendToCosmosClaimResponse = object;

export type GravityMsgSendToEthResponse = object;

export type GravityMsgSetOrchestratorAddressResponse = object;

export type GravityMsgSubmitBadSignatureEvidenceResponse = object;

/**
* MsgValsetConfirm
this is the message sent by the validators when they wish to submit their
signatures over the validator set at a given block height. A validator must
first call MsgSetEthAddress to set their Ethereum address to be used for
signing. Then someone (anyone) must make a ValsetRequest, the request is
essentially a messaging mechanism to determine which block all validators
should submit signatures over. Finally validators sign the validator set,
powers, and Ethereum addresses of the entire validator set at the height of a
ValsetRequest and submit that signature with this message.

If a sufficient number of validators (66% of voting power) (A) have set
Ethereum addresses and (B) submit ValsetConfirm messages with their
signatures it is then possible for anyone to view these signatures in the
chain store and submit them to Ethereum to update the validator set
-------------
*/
export interface GravityMsgValsetConfirm {
  /** @format uint64 */
  nonce?: string;
  orchestrator?: string;
  eth_address?: string;
  signature?: string;
}

export type GravityMsgValsetConfirmResponse = object;

export type GravityMsgValsetUpdatedClaimResponse = object;

export interface GravityOutgoingLogicCall {
  transfers?: GravityERC20Token[];
  fees?: GravityERC20Token[];
  logic_contract_address?: string;

  /** @format byte */
  payload?: string;

  /** @format uint64 */
  timeout?: string;

  /** @format byte */
  invalidation_id?: string;

  /** @format uint64 */
  invalidation_nonce?: string;

  /** @format uint64 */
  block?: string;
}

export interface GravityOutgoingTransferTx {
  /** @format uint64 */
  id?: string;
  sender?: string;
  dest_address?: string;
  erc20_token?: GravityERC20Token;
  erc20_fee?: GravityERC20Token;
}

export interface GravityOutgoingTxBatch {
  /** @format uint64 */
  batch_nonce?: string;

  /** @format uint64 */
  batch_timeout?: string;
  transactions?: GravityOutgoingTransferTx[];
  token_contract?: string;

  /** @format uint64 */
  block?: string;
}

/**
* unbond_slashing_valsets_window

The unbond slashing valsets window is used to determine how many blocks after starting to unbond
a validator needs to continue signing blocks. The goal of this paramater is that when a validator leaves
the set, if their leaving creates enough change in the validator set to justify an update they will sign
a validator set update for the Ethereum bridge that does not include themselves. Allowing us to remove them
from the Ethereum bridge and replace them with the new set gracefully.

valset_reward

These parameters allow for the bridge oracle to resolve a fork on the Ethereum chain without halting
the chain. Once set reset bridge state will roll back events to the nonce provided in reset_bridge_nonce
if and only if those events have not yet been observed (executed on the Cosmos chain). This allows for easy
handling of cases where for example an Ethereum hardfork has occured and more than 1/3 of the vlaidtor set
disagrees with the rest. Normally this would require a chain halt, manual genesis editing and restar to resolve
with this feature a governance proposal can be used instead

bridge_active

This boolean flag can be used by governance to temporarily halt the bridge due to a vulnerability or other issue
In this context halting the bridge means prevent the execution of any oracle events from Ethereum and preventing
the creation of new batches that may be relayed to Ethereum.
This does not prevent the creation of validator sets
or slashing for not submitting validator set signatures as either of these might allow key signers to leave the validator
set and steal funds on Ethereum without consequence.
The practical outcome of this flag being set to 'false' is that deposits from Ethereum will not show up and withdraws from
Cosmos will not execute on Ethereum.
*/
export interface GravityParams {
  gravity_id?: string;
  contract_source_hash?: string;
  bridge_ethereum_address?: string;

  /** @format uint64 */
  bridge_chain_id?: string;

  /** @format uint64 */
  signed_valsets_window?: string;

  /** @format uint64 */
  signed_batches_window?: string;

  /** @format uint64 */
  signed_logic_calls_window?: string;

  /** @format uint64 */
  target_batch_timeout?: string;

  /** @format uint64 */
  average_block_time?: string;

  /** @format uint64 */
  average_ethereum_block_time?: string;

  /** @format byte */
  slash_fraction_valset?: string;

  /** @format byte */
  slash_fraction_batch?: string;

  /** @format byte */
  slash_fraction_logic_call?: string;

  /** @format uint64 */
  unbond_slashing_valsets_window?: string;

  /** @format byte */
  slash_fraction_bad_eth_signature?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  valset_reward?: V1Beta1Coin;
  bridge_active?: boolean;
  ethereum_blacklist?: string[];
}

export interface GravityPendingIbcAutoForward {
  foreign_receiver?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  token?: V1Beta1Coin;
  ibc_channel?: string;

  /** @format uint64 */
  event_nonce?: string;
}

export interface GravityQueryAttestationsResponse {
  attestations?: GravityAttestation[];
}

export interface GravityQueryBatchConfirmsResponse {
  confirms?: GravityMsgConfirmBatch[];
}

export interface GravityQueryBatchFeeResponse {
  batch_fees?: GravityBatchFees[];
}

export interface GravityQueryBatchRequestByNonceResponse {
  batch?: GravityOutgoingTxBatch;
}

export interface GravityQueryCurrentValsetResponse {
  valset?: GravityValset;
}

export interface GravityQueryDelegateKeysByEthAddressResponse {
  validator_address?: string;
  orchestrator_address?: string;
}

export interface GravityQueryDelegateKeysByOrchestratorAddressResponse {
  validator_address?: string;
  eth_address?: string;
}

export interface GravityQueryDelegateKeysByValidatorAddressResponse {
  eth_address?: string;
  orchestrator_address?: string;
}

export interface GravityQueryDenomToERC20Response {
  erc20?: string;
  cosmos_originated?: boolean;
}

export interface GravityQueryERC20ToDenomResponse {
  denom?: string;
  cosmos_originated?: boolean;
}

export interface GravityQueryLastEventNonceByAddrResponse {
  /** @format uint64 */
  event_nonce?: string;
}

export interface GravityQueryLastPendingBatchRequestByAddrResponse {
  batch?: GravityOutgoingTxBatch[];
}

export interface GravityQueryLastPendingLogicCallByAddrResponse {
  call?: GravityOutgoingLogicCall[];
}

export interface GravityQueryLastPendingValsetRequestByAddrResponse {
  valsets?: GravityValset[];
}

export interface GravityQueryLastValsetRequestsResponse {
  valsets?: GravityValset[];
}

export interface GravityQueryLogicConfirmsResponse {
  confirms?: GravityMsgConfirmLogicCall[];
}

export interface GravityQueryOutgoingLogicCallsResponse {
  calls?: GravityOutgoingLogicCall[];
}

export interface GravityQueryOutgoingTxBatchesResponse {
  batches?: GravityOutgoingTxBatch[];
}

export interface GravityQueryParamsResponse {
  /**
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
  params?: GravityParams;
}

export interface GravityQueryPendingIbcAutoForwardsResponse {
  pending_ibc_auto_forwards?: GravityPendingIbcAutoForward[];
}

export interface GravityQueryPendingSendToEthResponse {
  transfers_in_batches?: GravityOutgoingTransferTx[];
  unbatched_transfers?: GravityOutgoingTransferTx[];
}

export interface GravityQueryValsetConfirmResponse {
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
  confirm?: GravityMsgValsetConfirm;
}

export interface GravityQueryValsetConfirmsByNonceResponse {
  confirms?: GravityMsgValsetConfirm[];
}

export interface GravityQueryValsetRequestResponse {
  valset?: GravityValset;
}

export interface GravityValset {
  /** @format uint64 */
  nonce?: string;
  members?: GravityBridgeValidator[];

  /** @format uint64 */
  height?: string;
  reward_amount?: string;
  reward_token?: string;
}

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   *
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   *
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   *
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title gravity/attestation.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Msg
   * @name MsgBatchSendToEthClaim
   * @request POST:/gravity/v1/batch_send_to_eth_claim
   */
  msgBatchSendToEthClaim = (
    query?: {
      event_nonce?: string;
      block_height?: string;
      batch_nonce?: string;
      token_contract?: string;
      orchestrator?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgBatchSendToEthClaimResponse, RpcStatus>({
      path: `/gravity/v1/batch_send_to_eth_claim`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgCancelSendToEth
   * @request POST:/gravity/v1/cancel_send_to_eth
   */
  msgCancelSendToEth = (query?: { transaction_id?: string; sender?: string }, params: RequestParams = {}) =>
    this.request<GravityMsgCancelSendToEthResponse, RpcStatus>({
      path: `/gravity/v1/cancel_send_to_eth`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgConfirmLogicCall
   * @request POST:/gravity/v1/confim_logic
   */
  msgConfirmLogicCall = (
    query?: {
      invalidation_id?: string;
      invalidation_nonce?: string;
      eth_signer?: string;
      orchestrator?: string;
      signature?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgConfirmLogicCallResponse, RpcStatus>({
      path: `/gravity/v1/confim_logic`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgConfirmBatch
   * @request POST:/gravity/v1/confirm_batch
   */
  msgConfirmBatch = (
    query?: { nonce?: string; token_contract?: string; eth_signer?: string; orchestrator?: string; signature?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgConfirmBatchResponse, RpcStatus>({
      path: `/gravity/v1/confirm_batch`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgErc20DeployedClaim
   * @request POST:/gravity/v1/erc20_deployed_claim
   */
  msgErc20DeployedClaim = (
    query?: {
      event_nonce?: string;
      block_height?: string;
      cosmos_denom?: string;
      token_contract?: string;
      name?: string;
      symbol?: string;
      decimals?: string;
      orchestrator?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgERC20DeployedClaimResponse, RpcStatus>({
      path: `/gravity/v1/erc20_deployed_claim`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgExecuteIbcAutoForwards
   * @request POST:/gravity/v1/execute_ibc_auto_forwards
   */
  msgExecuteIbcAutoForwards = (query?: { forwards_to_clear?: string; executor?: string }, params: RequestParams = {}) =>
    this.request<GravityMsgExecuteIbcAutoForwardsResponse, RpcStatus>({
      path: `/gravity/v1/execute_ibc_auto_forwards`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgLogicCallExecutedClaim
   * @request POST:/gravity/v1/logic_call_executed_claim
   */
  msgLogicCallExecutedClaim = (
    query?: {
      event_nonce?: string;
      block_height?: string;
      invalidation_id?: string;
      invalidation_nonce?: string;
      orchestrator?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgLogicCallExecutedClaimResponse, RpcStatus>({
      path: `/gravity/v1/logic_call_executed_claim`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgRequestBatch
   * @request POST:/gravity/v1/request_batch
   */
  msgRequestBatch = (query?: { sender?: string; denom?: string }, params: RequestParams = {}) =>
    this.request<GravityMsgRequestBatchResponse, RpcStatus>({
      path: `/gravity/v1/request_batch`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgSendToCosmosClaim
   * @request POST:/gravity/v1/send_to_cosmos_claim
   */
  msgSendToCosmosClaim = (
    query?: {
      event_nonce?: string;
      block_height?: string;
      token_contract?: string;
      amount?: string;
      ethereum_sender?: string;
      cosmos_receiver?: string;
      orchestrator?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgSendToCosmosClaimResponse, RpcStatus>({
      path: `/gravity/v1/send_to_cosmos_claim`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgSendToEth
   * @request POST:/gravity/v1/send_to_eth
   */
  msgSendToEth = (
    query?: {
      sender?: string;
      eth_dest?: string;
      "amount.denom"?: string;
      "amount.amount"?: string;
      "bridge_fee.denom"?: string;
      "bridge_fee.amount"?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgSendToEthResponse, RpcStatus>({
      path: `/gravity/v1/send_to_eth`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgSetOrchestratorAddress
   * @request POST:/gravity/v1/set_orchestrator_address
   */
  msgSetOrchestratorAddress = (
    query?: { validator?: string; orchestrator?: string; eth_address?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgSetOrchestratorAddressResponse, RpcStatus>({
      path: `/gravity/v1/set_orchestrator_address`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgSubmitBadSignatureEvidence
   * @request POST:/gravity/v1/submit_bad_signature_evidence
   */
  msgSubmitBadSignatureEvidence = (
    query?: { "subject.type_url"?: string; "subject.value"?: string; signature?: string; sender?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgSubmitBadSignatureEvidenceResponse, RpcStatus>({
      path: `/gravity/v1/submit_bad_signature_evidence`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgValsetConfirm
   * @request POST:/gravity/v1/valset_confirm
   */
  msgValsetConfirm = (
    query?: { nonce?: string; orchestrator?: string; eth_address?: string; signature?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgValsetConfirmResponse, RpcStatus>({
      path: `/gravity/v1/valset_confirm`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Msg
   * @name MsgValsetUpdateClaim
   * @request POST:/gravity/v1/valset_updated_claim
   */
  msgValsetUpdateClaim = (
    query?: {
      event_nonce?: string;
      valset_nonce?: string;
      block_height?: string;
      reward_amount?: string;
      reward_token?: string;
      orchestrator?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<GravityMsgValsetUpdatedClaimResponse, RpcStatus>({
      path: `/gravity/v1/valset_updated_claim`,
      method: "POST",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBatchConfirms
   * @request GET:/gravity/v1beta/batch/confirms
   */
  queryBatchConfirms = (query?: { nonce?: string; contract_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryBatchConfirmsResponse, RpcStatus>({
      path: `/gravity/v1beta/batch/confirms`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastPendingBatchRequestByAddr
   * @request GET:/gravity/v1beta/batch/last_pending_request_by_addr
   */
  queryLastPendingBatchRequestByAddr = (query?: { address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryLastPendingBatchRequestByAddrResponse, RpcStatus>({
      path: `/gravity/v1beta/batch/last_pending_request_by_addr`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutgoingLogicCalls
   * @request GET:/gravity/v1beta/batch/outgoinglogic
   */
  queryOutgoingLogicCalls = (params: RequestParams = {}) =>
    this.request<GravityQueryOutgoingLogicCallsResponse, RpcStatus>({
      path: `/gravity/v1beta/batch/outgoinglogic`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutgoingTxBatches
   * @request GET:/gravity/v1beta/batch/outgoingtx
   */
  queryOutgoingTxBatches = (params: RequestParams = {}) =>
    this.request<GravityQueryOutgoingTxBatchesResponse, RpcStatus>({
      path: `/gravity/v1beta/batch/outgoingtx`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBatchRequestByNonce
   * @request GET:/gravity/v1beta/batch/request_by_nonce
   */
  queryBatchRequestByNonce = (query?: { nonce?: string; contract_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryBatchRequestByNonceResponse, RpcStatus>({
      path: `/gravity/v1beta/batch/request_by_nonce`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBatchFees
   * @request GET:/gravity/v1beta/batchfees
   */
  queryBatchFees = (params: RequestParams = {}) =>
    this.request<GravityQueryBatchFeeResponse, RpcStatus>({
      path: `/gravity/v1beta/batchfees`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValsetConfirmsByNonce
   * @request GET:/gravity/v1beta/confirms/{nonce}
   */
  queryValsetConfirmsByNonce = (nonce: string, params: RequestParams = {}) =>
    this.request<GravityQueryValsetConfirmsByNonceResponse, RpcStatus>({
      path: `/gravity/v1beta/confirms/${nonce}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDenomToErc20
   * @request GET:/gravity/v1beta/cosmos_originated/denom_to_erc20
   */
  queryDenomToErc20 = (query?: { denom?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryDenomToERC20Response, RpcStatus>({
      path: `/gravity/v1beta/cosmos_originated/denom_to_erc20`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryErc20ToDenom
   * @request GET:/gravity/v1beta/cosmos_originated/erc20_to_denom
   */
  queryErc20ToDenom = (query?: { erc20?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryERC20ToDenomResponse, RpcStatus>({
      path: `/gravity/v1beta/cosmos_originated/erc20_to_denom`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLogicConfirms
   * @request GET:/gravity/v1beta/logic/confirms
   */
  queryLogicConfirms = (
    query?: { invalidation_id?: string; invalidation_nonce?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityQueryLogicConfirmsResponse, RpcStatus>({
      path: `/gravity/v1beta/logic/confirms`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastPendingLogicCallByAddr
   * @request GET:/gravity/v1beta/logic/{address}
   */
  queryLastPendingLogicCallByAddr = (address: string, params: RequestParams = {}) =>
    this.request<GravityQueryLastPendingLogicCallByAddrResponse, RpcStatus>({
      path: `/gravity/v1beta/logic/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastEventNonceByAddr
   * @request GET:/gravity/v1beta/oracle/eventnonce/{address}
   */
  queryLastEventNonceByAddr = (address: string, params: RequestParams = {}) =>
    this.request<GravityQueryLastEventNonceByAddrResponse, RpcStatus>({
      path: `/gravity/v1beta/oracle/eventnonce/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Deployments queries deployments
   * @request GET:/gravity/v1beta/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<GravityQueryParamsResponse, RpcStatus>({
      path: `/gravity/v1beta/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetAttestations
   * @request GET:/gravity/v1beta/query_attestations
   */
  queryGetAttestations = (
    query?: { limit?: string; order_by?: string; claim_type?: string; nonce?: string; height?: string },
    params: RequestParams = {},
  ) =>
    this.request<GravityQueryAttestationsResponse, RpcStatus>({
      path: `/gravity/v1beta/query_attestations`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetDelegateKeyByEth
   * @request GET:/gravity/v1beta/query_delegate_keys_by_eth
   */
  queryGetDelegateKeyByEth = (query?: { eth_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryDelegateKeysByEthAddressResponse, RpcStatus>({
      path: `/gravity/v1beta/query_delegate_keys_by_eth`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetDelegateKeyByOrchestrator
   * @request GET:/gravity/v1beta/query_delegate_keys_by_orchestrator
   */
  queryGetDelegateKeyByOrchestrator = (query?: { orchestrator_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryDelegateKeysByOrchestratorAddressResponse, RpcStatus>({
      path: `/gravity/v1beta/query_delegate_keys_by_orchestrator`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetDelegateKeyByValidator
   * @request GET:/gravity/v1beta/query_delegate_keys_by_validator
   */
  queryGetDelegateKeyByValidator = (query?: { validator_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryDelegateKeysByValidatorAddressResponse, RpcStatus>({
      path: `/gravity/v1beta/query_delegate_keys_by_validator`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPendingIbcAutoForwards
   * @request GET:/gravity/v1beta/query_pending_ibc_auto_forwards
   */
  queryGetPendingIbcAutoForwards = (query?: { limit?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryPendingIbcAutoForwardsResponse, RpcStatus>({
      path: `/gravity/v1beta/query_pending_ibc_auto_forwards`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPendingSendToEth
   * @request GET:/gravity/v1beta/query_pending_send_to_eth
   */
  queryGetPendingSendToEth = (query?: { sender_address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryPendingSendToEthResponse, RpcStatus>({
      path: `/gravity/v1beta/query_pending_send_to_eth`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValsetRequest
   * @request GET:/gravity/v1beta/valset
   */
  queryValsetRequest = (query?: { nonce?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryValsetRequestResponse, RpcStatus>({
      path: `/gravity/v1beta/valset`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValsetConfirm
   * @request GET:/gravity/v1beta/valset/confirm
   */
  queryValsetConfirm = (query?: { nonce?: string; address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryValsetConfirmResponse, RpcStatus>({
      path: `/gravity/v1beta/valset/confirm`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCurrentValset
   * @request GET:/gravity/v1beta/valset/current
   */
  queryCurrentValset = (params: RequestParams = {}) =>
    this.request<GravityQueryCurrentValsetResponse, RpcStatus>({
      path: `/gravity/v1beta/valset/current`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastPendingValsetRequestByAddr
   * @request GET:/gravity/v1beta/valset/last
   */
  queryLastPendingValsetRequestByAddr = (query?: { address?: string }, params: RequestParams = {}) =>
    this.request<GravityQueryLastPendingValsetRequestByAddrResponse, RpcStatus>({
      path: `/gravity/v1beta/valset/last`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastValsetRequests
   * @request GET:/gravity/v1beta/valset/requests
   */
  queryLastValsetRequests = (params: RequestParams = {}) =>
    this.request<GravityQueryLastValsetRequestsResponse, RpcStatus>({
      path: `/gravity/v1beta/valset/requests`,
      method: "GET",
      format: "json",
      ...params,
    });
}
