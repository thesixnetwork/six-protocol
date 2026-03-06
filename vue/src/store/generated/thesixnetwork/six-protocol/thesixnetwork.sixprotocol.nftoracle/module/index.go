// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSyncActionSigner } from "./types/nftoracle/tx";
import { MsgSubmitSyncActionSigner } from "./types/nftoracle/tx";
import { MsgCreateActionRequest } from "./types/nftoracle/tx";
import { MsgDeleteActionSignerConfig } from "./types/nftoracle/tx";
import { MsgCreateVerifyCollectionOwnerRequest } from "./types/nftoracle/tx";
import { MsgSubmitActionResponse } from "./types/nftoracle/tx";
import { MsgCreateActionSigner } from "./types/nftoracle/tx";
import { MsgDeleteActionSigner } from "./types/nftoracle/tx";
import { MsgCreateMintRequest } from "./types/nftoracle/tx";
import { MsgSubmitVerifyCollectionOwner } from "./types/nftoracle/tx";
import { MsgSubmitMintResponse } from "./types/nftoracle/tx";
import { MsgUpdateActionSignerConfig } from "./types/nftoracle/tx";
import { MsgSetMinimumConfirmation } from "./types/nftoracle/tx";
import { MsgUpdateActionSigner } from "./types/nftoracle/tx";
import { MsgCreateActionSignerConfig } from "./types/nftoracle/tx";


const types = [
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateSyncActionSigner", MsgCreateSyncActionSigner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgSubmitSyncActionSigner", MsgSubmitSyncActionSigner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionRequest", MsgCreateActionRequest],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgDeleteActionSignerConfig", MsgDeleteActionSignerConfig],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateVerifyCollectionOwnerRequest", MsgCreateVerifyCollectionOwnerRequest],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgSubmitActionResponse", MsgSubmitActionResponse],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionSigner", MsgCreateActionSigner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgDeleteActionSigner", MsgDeleteActionSigner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateMintRequest", MsgCreateMintRequest],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgSubmitVerifyCollectionOwner", MsgSubmitVerifyCollectionOwner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgSubmitMintResponse", MsgSubmitMintResponse],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgUpdateActionSignerConfig", MsgUpdateActionSignerConfig],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgSetMinimumConfirmation", MsgSetMinimumConfirmation],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgUpdateActionSigner", MsgUpdateActionSigner],
  ["/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionSignerConfig", MsgCreateActionSignerConfig],
  
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
    msgCreateSyncActionSigner: (data: MsgCreateSyncActionSigner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateSyncActionSigner", value: MsgCreateSyncActionSigner.fromPartial( data ) }),
    msgSubmitSyncActionSigner: (data: MsgSubmitSyncActionSigner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgSubmitSyncActionSigner", value: MsgSubmitSyncActionSigner.fromPartial( data ) }),
    msgCreateActionRequest: (data: MsgCreateActionRequest): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionRequest", value: MsgCreateActionRequest.fromPartial( data ) }),
    msgDeleteActionSignerConfig: (data: MsgDeleteActionSignerConfig): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgDeleteActionSignerConfig", value: MsgDeleteActionSignerConfig.fromPartial( data ) }),
    msgCreateVerifyCollectionOwnerRequest: (data: MsgCreateVerifyCollectionOwnerRequest): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateVerifyCollectionOwnerRequest", value: MsgCreateVerifyCollectionOwnerRequest.fromPartial( data ) }),
    msgSubmitActionResponse: (data: MsgSubmitActionResponse): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgSubmitActionResponse", value: MsgSubmitActionResponse.fromPartial( data ) }),
    msgCreateActionSigner: (data: MsgCreateActionSigner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionSigner", value: MsgCreateActionSigner.fromPartial( data ) }),
    msgDeleteActionSigner: (data: MsgDeleteActionSigner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgDeleteActionSigner", value: MsgDeleteActionSigner.fromPartial( data ) }),
    msgCreateMintRequest: (data: MsgCreateMintRequest): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateMintRequest", value: MsgCreateMintRequest.fromPartial( data ) }),
    msgSubmitVerifyCollectionOwner: (data: MsgSubmitVerifyCollectionOwner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgSubmitVerifyCollectionOwner", value: MsgSubmitVerifyCollectionOwner.fromPartial( data ) }),
    msgSubmitMintResponse: (data: MsgSubmitMintResponse): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgSubmitMintResponse", value: MsgSubmitMintResponse.fromPartial( data ) }),
    msgUpdateActionSignerConfig: (data: MsgUpdateActionSignerConfig): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgUpdateActionSignerConfig", value: MsgUpdateActionSignerConfig.fromPartial( data ) }),
    msgSetMinimumConfirmation: (data: MsgSetMinimumConfirmation): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgSetMinimumConfirmation", value: MsgSetMinimumConfirmation.fromPartial( data ) }),
    msgUpdateActionSigner: (data: MsgUpdateActionSigner): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgUpdateActionSigner", value: MsgUpdateActionSigner.fromPartial( data ) }),
    msgCreateActionSignerConfig: (data: MsgCreateActionSignerConfig): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.nftoracle.MsgCreateActionSignerConfig", value: MsgCreateActionSignerConfig.fromPartial( data ) }),
    
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
