import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteGroup } from "./types/protocoladmin/tx";
import { MsgRemoveAdminFromGroup } from "./types/protocoladmin/tx";
import { MsgUpdateGroup } from "./types/protocoladmin/tx";
import { MsgCreateGroup } from "./types/protocoladmin/tx";
import { MsgAddAdminToGroup } from "./types/protocoladmin/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgDeleteGroup: (data: MsgDeleteGroup) => EncodeObject;
    msgRemoveAdminFromGroup: (data: MsgRemoveAdminFromGroup) => EncodeObject;
    msgUpdateGroup: (data: MsgUpdateGroup) => EncodeObject;
    msgCreateGroup: (data: MsgCreateGroup) => EncodeObject;
    msgAddAdminToGroup: (data: MsgAddAdminToGroup) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
