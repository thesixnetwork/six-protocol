// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateGroup } from "./types/protocoladmin/tx";
import { MsgRemoveAdminFromGroup } from "./types/protocoladmin/tx";
import { MsgAddAdminToGroup } from "./types/protocoladmin/tx";
import { MsgDeleteGroup } from "./types/protocoladmin/tx";
import { MsgUpdateGroup } from "./types/protocoladmin/tx";
const types = [
    ["/thesixnetwork.sixprotocol.protocoladmin.MsgCreateGroup", MsgCreateGroup],
    ["/thesixnetwork.sixprotocol.protocoladmin.MsgRemoveAdminFromGroup", MsgRemoveAdminFromGroup],
    ["/thesixnetwork.sixprotocol.protocoladmin.MsgAddAdminToGroup", MsgAddAdminToGroup],
    ["/thesixnetwork.sixprotocol.protocoladmin.MsgDeleteGroup", MsgDeleteGroup],
    ["/thesixnetwork.sixprotocol.protocoladmin.MsgUpdateGroup", MsgUpdateGroup],
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
        msgCreateGroup: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.protocoladmin.MsgCreateGroup", value: MsgCreateGroup.fromPartial(data) }),
        msgRemoveAdminFromGroup: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.protocoladmin.MsgRemoveAdminFromGroup", value: MsgRemoveAdminFromGroup.fromPartial(data) }),
        msgAddAdminToGroup: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.protocoladmin.MsgAddAdminToGroup", value: MsgAddAdminToGroup.fromPartial(data) }),
        msgDeleteGroup: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.protocoladmin.MsgDeleteGroup", value: MsgDeleteGroup.fromPartial(data) }),
        msgUpdateGroup: (data) => ({ typeUrl: "/thesixnetwork.sixprotocol.protocoladmin.MsgUpdateGroup", value: MsgUpdateGroup.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
