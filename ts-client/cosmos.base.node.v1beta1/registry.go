import { GeneratedType } from "@cosmjs/proto-signing";
import { StatusRequest } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/base/node/v1beta1/query";
import { StatusResponse } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/base/node/v1beta1/query";
import { ConfigRequest } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/base/node/v1beta1/query";
import { ConfigResponse } from "./types/../../../../../../../../go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.10/proto/cosmos/base/node/v1beta1/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cosmos.base.node.v1beta1.StatusRequest", StatusRequest],
    ["/cosmos.base.node.v1beta1.StatusResponse", StatusResponse],
    ["/cosmos.base.node.v1beta1.ConfigRequest", ConfigRequest],
    ["/cosmos.base.node.v1beta1.ConfigResponse", ConfigResponse],
    
];

export { msgTypes }