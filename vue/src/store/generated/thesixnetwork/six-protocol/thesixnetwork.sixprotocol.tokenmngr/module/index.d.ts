import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
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
    msgUpdateToken: (data: MsgUpdateToken) => EncodeObject;
    msgDeleteMintperm: (data: MsgDeleteMintperm) => EncodeObject;
    msgDeleteToken: (data: MsgDeleteToken) => EncodeObject;
    msgCreateOptions: (data: MsgCreateOptions) => EncodeObject;
    msgCreateMintperm: (data: MsgCreateMintperm) => EncodeObject;
    msgMint: (data: MsgMint) => EncodeObject;
    msgCreateToken: (data: MsgCreateToken) => EncodeObject;
    msgUpdateMintperm: (data: MsgUpdateMintperm) => EncodeObject;
    msgUpdateOptions: (data: MsgUpdateOptions) => EncodeObject;
    msgDeleteOptions: (data: MsgDeleteOptions) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
