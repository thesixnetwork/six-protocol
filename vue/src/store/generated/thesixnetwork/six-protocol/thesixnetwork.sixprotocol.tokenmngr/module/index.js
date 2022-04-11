// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateToken } from "./types/tokenmngr/tx";
import { MsgDeleteMintperm } from "./types/tokenmngr/tx";
import { MsgDeleteToken } from "./types/tokenmngr/tx";
import { MsgCreateOptions } from "./types/tokenmngr/tx";
import { MsgCreateMintperm } from "./types/tokenmngr/tx";
import { MsgMint } from "./types/tokenmngr/tx";
import { MsgCreateToken } from "./types/tokenmngr/tx";
import { MsgUpdateMintperm } from "./types/tokenmngr/tx";
import { MsgUpdateOptions } from "./types/tokenmngr/tx";
import { MsgDeleteOptions } from "./types/tokenmngr/tx";
const types = [
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateToken", MsgUpdateToken],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteMintperm", MsgDeleteMintperm],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteToken", MsgDeleteToken],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgCreateOptions", MsgCreateOptions],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgCreateMintperm", MsgCreateMintperm],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgMint", MsgMint],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgCreateToken", MsgCreateToken],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateMintperm", MsgUpdateMintperm],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateOptions", MsgUpdateOptions],
    ["/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteOptions", MsgDeleteOptions],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateToken: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateToken", value: MsgUpdateToken.fromPartial(data) }),
        msgDeleteMintperm: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteMintperm", value: MsgDeleteMintperm.fromPartial(data) }),
        msgDeleteToken: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteToken", value: MsgDeleteToken.fromPartial(data) }),
        msgCreateOptions: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateOptions", value: MsgCreateOptions.fromPartial(data) }),
        msgCreateMintperm: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateMintperm", value: MsgCreateMintperm.fromPartial(data) }),
        msgMint: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgMint", value: MsgMint.fromPartial(data) }),
        msgCreateToken: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateToken", value: MsgCreateToken.fromPartial(data) }),
        msgUpdateMintperm: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateMintperm", value: MsgUpdateMintperm.fromPartial(data) }),
        msgUpdateOptions: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgUpdateOptions", value: MsgUpdateOptions.fromPartial(data) }),
        msgDeleteOptions: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgDeleteOptions", value: MsgDeleteOptions.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
