import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteOptions } from "./types/tokenmngr/tx";
import { MsgDeleteToken } from "./types/tokenmngr/tx";
import { MsgCreateToken } from "./types/tokenmngr/tx";
import { MsgMint } from "./types/tokenmngr/tx";
import { MsgUpdateMintperm } from "./types/tokenmngr/tx";
import { MsgUpdateToken } from "./types/tokenmngr/tx";
import { MsgDeleteMintperm } from "./types/tokenmngr/tx";
import { MsgCreateMintperm } from "./types/tokenmngr/tx";
import { MsgCreateOptions } from "./types/tokenmngr/tx";
import { MsgUpdateOptions } from "./types/tokenmngr/tx";
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
    msgDeleteOptions: (data: MsgDeleteOptions) => EncodeObject;
    msgDeleteToken: (data: MsgDeleteToken) => EncodeObject;
    msgCreateToken: (data: MsgCreateToken) => EncodeObject;
    msgMint: (data: MsgMint) => EncodeObject;
    msgUpdateMintperm: (data: MsgUpdateMintperm) => EncodeObject;
    msgUpdateToken: (data: MsgUpdateToken) => EncodeObject;
    msgDeleteMintperm: (data: MsgDeleteMintperm) => EncodeObject;
    msgCreateMintperm: (data: MsgCreateMintperm) => EncodeObject;
    msgCreateOptions: (data: MsgCreateOptions) => EncodeObject;
    msgUpdateOptions: (data: MsgUpdateOptions) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
