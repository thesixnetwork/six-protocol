import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/consensus/v1/query";
import { QueryParamsResponse } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/consensus/v1/query";
import { MsgUpdateParamsResponse } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/consensus/v1/tx";
import { MsgUpdateParams } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/consensus/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cosmos.consensus.v1.QueryParamsRequest", QueryParamsRequest],
    ["/cosmos.consensus.v1.QueryParamsResponse", QueryParamsResponse],
    ["/cosmos.consensus.v1.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/cosmos.consensus.v1.MsgUpdateParams", MsgUpdateParams],
    
];

export { msgTypes }