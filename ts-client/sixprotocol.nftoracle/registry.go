import { GeneratedType } from "@cosmjs/proto-signing";
import { MintRequest } from "./types/sixprotocol/nftoracle/mint_request";
import { MsgUpdateParamsResponse } from "./types/sixprotocol/nftoracle/tx";
import { MsgUpdateParams } from "./types/sixprotocol/nftoracle/tx";
import { QueryParamsResponse } from "./types/sixprotocol/nftoracle/query";
import { QueryGetMintRequestRequest } from "./types/sixprotocol/nftoracle/query";
import { QueryAllMintRequestResponse } from "./types/sixprotocol/nftoracle/query";
import { GenesisState } from "./types/sixprotocol/nftoracle/genesis";
import { QueryParamsRequest } from "./types/sixprotocol/nftoracle/query";
import { QueryGetMintRequestResponse } from "./types/sixprotocol/nftoracle/query";
import { QueryAllMintRequestRequest } from "./types/sixprotocol/nftoracle/query";
import { Params } from "./types/sixprotocol/nftoracle/params";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixprotocol.nftoracle.MintRequest", MintRequest],
    ["/sixprotocol.nftoracle.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/sixprotocol.nftoracle.MsgUpdateParams", MsgUpdateParams],
    ["/sixprotocol.nftoracle.QueryParamsResponse", QueryParamsResponse],
    ["/sixprotocol.nftoracle.QueryGetMintRequestRequest", QueryGetMintRequestRequest],
    ["/sixprotocol.nftoracle.QueryAllMintRequestResponse", QueryAllMintRequestResponse],
    ["/sixprotocol.nftoracle.GenesisState", GenesisState],
    ["/sixprotocol.nftoracle.QueryParamsRequest", QueryParamsRequest],
    ["/sixprotocol.nftoracle.QueryGetMintRequestResponse", QueryGetMintRequestResponse],
    ["/sixprotocol.nftoracle.QueryAllMintRequestRequest", QueryAllMintRequestRequest],
    ["/sixprotocol.nftoracle.Params", Params],
    
];

export { msgTypes }