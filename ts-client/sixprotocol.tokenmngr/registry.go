import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/sixprotocol/tokenmngr/query";
import { QueryParamsResponse } from "./types/sixprotocol/tokenmngr/query";
import { Params } from "./types/sixprotocol/tokenmngr/params";
import { MsgMigrateDelegation } from "./types/sixprotocol/tokenmngr/tx";
import { MsgUpdateParams } from "./types/sixprotocol/tokenmngr/tx";
import { MsgUpdateParamsResponse } from "./types/sixprotocol/tokenmngr/tx";
import { MsgMint } from "./types/sixprotocol/tokenmngr/tx";
import { MsgMintResponse } from "./types/sixprotocol/tokenmngr/tx";
import { MsgMigrateDelegationResponse } from "./types/sixprotocol/tokenmngr/tx";
import { GenesisState } from "./types/sixprotocol/tokenmngr/genesis";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixprotocol.tokenmngr.QueryParamsRequest", QueryParamsRequest],
    ["/sixprotocol.tokenmngr.QueryParamsResponse", QueryParamsResponse],
    ["/sixprotocol.tokenmngr.Params", Params],
    ["/sixprotocol.tokenmngr.MsgMigrateDelegation", MsgMigrateDelegation],
    ["/sixprotocol.tokenmngr.MsgUpdateParams", MsgUpdateParams],
    ["/sixprotocol.tokenmngr.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/sixprotocol.tokenmngr.MsgMint", MsgMint],
    ["/sixprotocol.tokenmngr.MsgMintResponse", MsgMintResponse],
    ["/sixprotocol.tokenmngr.MsgMigrateDelegationResponse", MsgMigrateDelegationResponse],
    ["/sixprotocol.tokenmngr.GenesisState", GenesisState],
    
];

export { msgTypes }