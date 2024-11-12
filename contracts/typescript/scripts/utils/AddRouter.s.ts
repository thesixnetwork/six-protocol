import { SixDataChainConnector } from "../../sdk/client";
import { fee, ITxNFTmngr } from "../.././sdk";
import { EncodeObject } from "@cosmjs/proto-signing";
import { GasPrice, calculateFee } from "@cosmjs/stargate/build/fee";
import { getConnectorConfig } from "../../client";
import { ethereumToBech32 } from "../helper/account";
import dotenv from "dotenv";
dotenv.config();

const NETOWRK = process.argv[2]!;

const ROUTER_ADDDRES = "0xE18A45EE9e297e3fe1948B4B67E754d691936552"

import exmapleSchema from "../../../resources/nft-schema.json"
let schema_name = exmapleSchema.code;
const split_schema = schema_name.split(".")
const _name = split_schema[1]
const org_name = process.env.ORG_NAME
let schemaCode: string;
schemaCode = `${org_name}.${_name}`;

console.log(schemaCode);

exmapleSchema.code = schemaCode;

const addRouter = async () => {
  if (!NETOWRK) {
    throw new Error("INPUT NETWORK BY RUNNING: bun run ./scripts/deploy.ts fivenet || yarn ts-node ./scripts/deploy.ts fivenet");
  }

  const { rpcUrl, apiUrl, mnemonic } = await getConnectorConfig(NETOWRK);
  const sixConnector = new SixDataChainConnector();
  sixConnector.rpcUrl = rpcUrl;
  sixConnector.apiUrl = apiUrl;

  const accountSigner = await sixConnector.accounts.mnemonicKeyToAccount(mnemonic);

  const address = (await accountSigner.getAccounts())[0].address;
  const rpcClient = await sixConnector.connectRPCClient(accountSigner, {
    gasPrice: GasPrice.fromString("1.25usix"),
  });


  const routerBec32 = ethereumToBech32('6x', ROUTER_ADDDRES)

  let msgArray: Array<EncodeObject> = [];

  const msgAddRouter: ITxNFTmngr.MsgCreateActionExecutor = {
    creator: address,
    nftSchemaCode: schemaCode,
    executorAddress: routerBec32
  }

  const msg = await rpcClient.nftmngrModule.msgCreateActionExecutor(
    msgAddRouter
  )

  msgArray.push(msg);
  const txResponse = await rpcClient.nftmngrModule.signAndBroadcast(msgArray, {
    fee: "auto",
  });

  if (txResponse.code) {
    console.log(txResponse.rawLog);
  }
  console.log(
    `gasUsed: ${txResponse.gasUsed}\ngasWanted:${txResponse.gasWanted}\n`
  );
  return txResponse;
}



// ask to enter confirmmation
const readline = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
});

readline.question(
  `Are you sure you want to chnage owner of  ${schemaCode} to ${NETOWRK} (y/n)?`,
  (answer) => {
    if (
      answer === "y" ||
      answer === "Y" ||
      answer === "yes" ||
      answer === "Yes"
    ) {
      console.log("RUNNING...");
      addRouter();
    } else {
      console.log("aborting...");
      process.exit(1);
    }
    readline.close();
  }
);
