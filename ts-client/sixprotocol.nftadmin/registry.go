import { GeneratedType } from "@cosmjs/proto-signing";
import { Params } from "./types/sixprotocol/nftadmin/params";
import { QueryParamsRequest } from "./types/sixprotocol/nftadmin/query";
import { QueryParamsResponse } from "./types/sixprotocol/nftadmin/query";
import { GenesisState } from "./types/sixprotocol/nftadmin/genesis";
import { MsgUpdateParamsResponse } from "./types/sixprotocol/nftadmin/tx";
import { MsgUpdateParams } from "./types/sixprotocol/nftadmin/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixprotocol.nftadmin.Params", Params],
    ["/sixprotocol.nftadmin.QueryParamsRequest", QueryParamsRequest],
    ["/sixprotocol.nftadmin.QueryParamsResponse", QueryParamsResponse],
    ["/sixprotocol.nftadmin.GenesisState", GenesisState],
    ["/sixprotocol.nftadmin.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/sixprotocol.nftadmin.MsgUpdateParams", MsgUpdateParams],
    
];

export { msgTypes }