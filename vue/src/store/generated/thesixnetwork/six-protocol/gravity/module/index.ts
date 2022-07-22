// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestBatch } from "./types/gravity/msgs";
import { MsgSendToCosmosClaim } from "./types/gravity/msgs";
import { MsgSubmitBadSignatureEvidence } from "./types/gravity/msgs";
import { MsgERC20DeployedClaim } from "./types/gravity/msgs";
import { MsgConfirmBatch } from "./types/gravity/msgs";
import { MsgValsetConfirm } from "./types/gravity/msgs";
import { MsgLogicCallExecutedClaim } from "./types/gravity/msgs";
import { MsgSendToEth } from "./types/gravity/msgs";
import { MsgBatchSendToEthClaim } from "./types/gravity/msgs";
import { MsgCancelSendToEth } from "./types/gravity/msgs";
import { MsgConfirmLogicCall } from "./types/gravity/msgs";
import { MsgSetOrchestratorAddress } from "./types/gravity/msgs";
import { MsgValsetUpdatedClaim } from "./types/gravity/msgs";


const types = [
  ["/gravity.MsgRequestBatch", MsgRequestBatch],
  ["/gravity.MsgSendToCosmosClaim", MsgSendToCosmosClaim],
  ["/gravity.MsgSubmitBadSignatureEvidence", MsgSubmitBadSignatureEvidence],
  ["/gravity.MsgERC20DeployedClaim", MsgERC20DeployedClaim],
  ["/gravity.MsgConfirmBatch", MsgConfirmBatch],
  ["/gravity.MsgValsetConfirm", MsgValsetConfirm],
  ["/gravity.MsgLogicCallExecutedClaim", MsgLogicCallExecutedClaim],
  ["/gravity.MsgSendToEth", MsgSendToEth],
  ["/gravity.MsgBatchSendToEthClaim", MsgBatchSendToEthClaim],
  ["/gravity.MsgCancelSendToEth", MsgCancelSendToEth],
  ["/gravity.MsgConfirmLogicCall", MsgConfirmLogicCall],
  ["/gravity.MsgSetOrchestratorAddress", MsgSetOrchestratorAddress],
  ["/gravity.MsgValsetUpdatedClaim", MsgValsetUpdatedClaim],
  
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
    msgRequestBatch: (data: MsgRequestBatch): EncodeObject => ({ typeUrl: "/gravity.MsgRequestBatch", value: MsgRequestBatch.fromPartial( data ) }),
    msgSendToCosmosClaim: (data: MsgSendToCosmosClaim): EncodeObject => ({ typeUrl: "/gravity.MsgSendToCosmosClaim", value: MsgSendToCosmosClaim.fromPartial( data ) }),
    msgSubmitBadSignatureEvidence: (data: MsgSubmitBadSignatureEvidence): EncodeObject => ({ typeUrl: "/gravity.MsgSubmitBadSignatureEvidence", value: MsgSubmitBadSignatureEvidence.fromPartial( data ) }),
    msgERC20DeployedClaim: (data: MsgERC20DeployedClaim): EncodeObject => ({ typeUrl: "/gravity.MsgERC20DeployedClaim", value: MsgERC20DeployedClaim.fromPartial( data ) }),
    msgConfirmBatch: (data: MsgConfirmBatch): EncodeObject => ({ typeUrl: "/gravity.MsgConfirmBatch", value: MsgConfirmBatch.fromPartial( data ) }),
    msgValsetConfirm: (data: MsgValsetConfirm): EncodeObject => ({ typeUrl: "/gravity.MsgValsetConfirm", value: MsgValsetConfirm.fromPartial( data ) }),
    msgLogicCallExecutedClaim: (data: MsgLogicCallExecutedClaim): EncodeObject => ({ typeUrl: "/gravity.MsgLogicCallExecutedClaim", value: MsgLogicCallExecutedClaim.fromPartial( data ) }),
    msgSendToEth: (data: MsgSendToEth): EncodeObject => ({ typeUrl: "/gravity.MsgSendToEth", value: MsgSendToEth.fromPartial( data ) }),
    msgBatchSendToEthClaim: (data: MsgBatchSendToEthClaim): EncodeObject => ({ typeUrl: "/gravity.MsgBatchSendToEthClaim", value: MsgBatchSendToEthClaim.fromPartial( data ) }),
    msgCancelSendToEth: (data: MsgCancelSendToEth): EncodeObject => ({ typeUrl: "/gravity.MsgCancelSendToEth", value: MsgCancelSendToEth.fromPartial( data ) }),
    msgConfirmLogicCall: (data: MsgConfirmLogicCall): EncodeObject => ({ typeUrl: "/gravity.MsgConfirmLogicCall", value: MsgConfirmLogicCall.fromPartial( data ) }),
    msgSetOrchestratorAddress: (data: MsgSetOrchestratorAddress): EncodeObject => ({ typeUrl: "/gravity.MsgSetOrchestratorAddress", value: MsgSetOrchestratorAddress.fromPartial( data ) }),
    msgValsetUpdatedClaim: (data: MsgValsetUpdatedClaim): EncodeObject => ({ typeUrl: "/gravity.MsgValsetUpdatedClaim", value: MsgValsetUpdatedClaim.fromPartial( data ) }),
    
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
