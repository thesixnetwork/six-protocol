import { GeneratedType } from "@cosmjs/proto-signing";
import { Params } from "./types/sixprotocol/protocoladmin/params";
import { QueryParamsResponse } from "./types/sixprotocol/protocoladmin/query";
import { GenesisState } from "./types/sixprotocol/protocoladmin/genesis";
import { MsgUpdateParamsResponse } from "./types/sixprotocol/protocoladmin/tx";
import { MsgUpdateParams } from "./types/sixprotocol/protocoladmin/tx";
import { QueryParamsRequest } from "./types/sixprotocol/protocoladmin/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixprotocol.protocoladmin.Params", Params],
    ["/sixprotocol.protocoladmin.QueryParamsResponse", QueryParamsResponse],
    ["/sixprotocol.protocoladmin.GenesisState", GenesisState],
    ["/sixprotocol.protocoladmin.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/sixprotocol.protocoladmin.MsgUpdateParams", MsgUpdateParams],
    ["/sixprotocol.protocoladmin.QueryParamsRequest", QueryParamsRequest],
    
];

export { msgTypes }