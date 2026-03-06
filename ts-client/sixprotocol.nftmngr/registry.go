import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/sixprotocol/nftmngr/tx";
import { MsgUpdateParamsResponse } from "./types/sixprotocol/nftmngr/tx";
import { Params } from "./types/sixprotocol/nftmngr/params";
import { QueryParamsRequest } from "./types/sixprotocol/nftmngr/query";
import { QueryParamsResponse } from "./types/sixprotocol/nftmngr/query";
import { GenesisState } from "./types/sixprotocol/nftmngr/genesis";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixprotocol.nftmngr.MsgUpdateParams", MsgUpdateParams],
    ["/sixprotocol.nftmngr.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/sixprotocol.nftmngr.Params", Params],
    ["/sixprotocol.nftmngr.QueryParamsRequest", QueryParamsRequest],
    ["/sixprotocol.nftmngr.QueryParamsResponse", QueryParamsResponse],
    ["/sixprotocol.nftmngr.GenesisState", GenesisState],
    
];

export { msgTypes }