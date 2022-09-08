// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateBinding } from "./types/evmbind/tx";
import { MsgUpdateBinding } from "./types/evmbind/tx";
import { MsgDeleteBinding } from "./types/evmbind/tx";
import { MsgEthSend } from "./types/evmbind/tx";


const types = [
  ["/thesixnetwork.sixprotocol.evmbind.MsgCreateBinding", MsgCreateBinding],
  ["/thesixnetwork.sixprotocol.evmbind.MsgUpdateBinding", MsgUpdateBinding],
  ["/thesixnetwork.sixprotocol.evmbind.MsgDeleteBinding", MsgDeleteBinding],
  ["/thesixnetwork.sixprotocol.evmbind.MsgEthSend", MsgEthSend],
  
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
    msgCreateBinding: (data: MsgCreateBinding): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.evmbind.MsgCreateBinding", value: MsgCreateBinding.fromPartial( data ) }),
    msgUpdateBinding: (data: MsgUpdateBinding): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.evmbind.MsgUpdateBinding", value: MsgUpdateBinding.fromPartial( data ) }),
    msgDeleteBinding: (data: MsgDeleteBinding): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.evmbind.MsgDeleteBinding", value: MsgDeleteBinding.fromPartial( data ) }),
    msgEthSend: (data: MsgEthSend): EncodeObject => ({ typeUrl: "/thesixnetwork.sixprotocol.evmbind.MsgEthSend", value: MsgEthSend.fromPartial( data ) }),
    
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
