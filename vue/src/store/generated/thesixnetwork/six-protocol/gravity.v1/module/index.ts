// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestBatch } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgConfirmLogicCall } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgERC20DeployedClaim } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgBatchSendToEthClaim } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgValsetConfirm } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgValsetUpdatedClaim } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgLogicCallExecutedClaim } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgSubmitBadSignatureEvidence } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgSendToCosmosClaim } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgCancelSendToEth } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgSetOrchestratorAddress } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgSendToEth } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";
import { MsgConfirmBatch } from "./types/thesixnetwork/six-protocol/gravity/v1/msgs";


const types = [
  ["/gravity.v1.MsgRequestBatch", MsgRequestBatch],
  ["/gravity.v1.MsgConfirmLogicCall", MsgConfirmLogicCall],
  ["/gravity.v1.MsgERC20DeployedClaim", MsgERC20DeployedClaim],
  ["/gravity.v1.MsgBatchSendToEthClaim", MsgBatchSendToEthClaim],
  ["/gravity.v1.MsgValsetConfirm", MsgValsetConfirm],
  ["/gravity.v1.MsgValsetUpdatedClaim", MsgValsetUpdatedClaim],
  ["/gravity.v1.MsgLogicCallExecutedClaim", MsgLogicCallExecutedClaim],
  ["/gravity.v1.MsgSubmitBadSignatureEvidence", MsgSubmitBadSignatureEvidence],
  ["/gravity.v1.MsgSendToCosmosClaim", MsgSendToCosmosClaim],
  ["/gravity.v1.MsgCancelSendToEth", MsgCancelSendToEth],
  ["/gravity.v1.MsgSetOrchestratorAddress", MsgSetOrchestratorAddress],
  ["/gravity.v1.MsgSendToEth", MsgSendToEth],
  ["/gravity.v1.MsgConfirmBatch", MsgConfirmBatch],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRequestBatch: (data: MsgRequestBatch): EncodeObject => ({ typeUrl: "/gravity.v1.MsgRequestBatch", value: MsgRequestBatch.fromPartial( data ) }),
    msgConfirmLogicCall: (data: MsgConfirmLogicCall): EncodeObject => ({ typeUrl: "/gravity.v1.MsgConfirmLogicCall", value: MsgConfirmLogicCall.fromPartial( data ) }),
    msgERC20DeployedClaim: (data: MsgERC20DeployedClaim): EncodeObject => ({ typeUrl: "/gravity.v1.MsgERC20DeployedClaim", value: MsgERC20DeployedClaim.fromPartial( data ) }),
    msgBatchSendToEthClaim: (data: MsgBatchSendToEthClaim): EncodeObject => ({ typeUrl: "/gravity.v1.MsgBatchSendToEthClaim", value: MsgBatchSendToEthClaim.fromPartial( data ) }),
    msgValsetConfirm: (data: MsgValsetConfirm): EncodeObject => ({ typeUrl: "/gravity.v1.MsgValsetConfirm", value: MsgValsetConfirm.fromPartial( data ) }),
    msgValsetUpdatedClaim: (data: MsgValsetUpdatedClaim): EncodeObject => ({ typeUrl: "/gravity.v1.MsgValsetUpdatedClaim", value: MsgValsetUpdatedClaim.fromPartial( data ) }),
    msgLogicCallExecutedClaim: (data: MsgLogicCallExecutedClaim): EncodeObject => ({ typeUrl: "/gravity.v1.MsgLogicCallExecutedClaim", value: MsgLogicCallExecutedClaim.fromPartial( data ) }),
    msgSubmitBadSignatureEvidence: (data: MsgSubmitBadSignatureEvidence): EncodeObject => ({ typeUrl: "/gravity.v1.MsgSubmitBadSignatureEvidence", value: MsgSubmitBadSignatureEvidence.fromPartial( data ) }),
    msgSendToCosmosClaim: (data: MsgSendToCosmosClaim): EncodeObject => ({ typeUrl: "/gravity.v1.MsgSendToCosmosClaim", value: MsgSendToCosmosClaim.fromPartial( data ) }),
    msgCancelSendToEth: (data: MsgCancelSendToEth): EncodeObject => ({ typeUrl: "/gravity.v1.MsgCancelSendToEth", value: MsgCancelSendToEth.fromPartial( data ) }),
    msgSetOrchestratorAddress: (data: MsgSetOrchestratorAddress): EncodeObject => ({ typeUrl: "/gravity.v1.MsgSetOrchestratorAddress", value: MsgSetOrchestratorAddress.fromPartial( data ) }),
    msgSendToEth: (data: MsgSendToEth): EncodeObject => ({ typeUrl: "/gravity.v1.MsgSendToEth", value: MsgSendToEth.fromPartial( data ) }),
    msgConfirmBatch: (data: MsgConfirmBatch): EncodeObject => ({ typeUrl: "/gravity.v1.MsgConfirmBatch", value: MsgConfirmBatch.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
