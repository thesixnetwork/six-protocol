import { ethers } from "ethers";
import NFTMngrABI from "../../out/INFTManager.sol/INFTMNGR.json";
import { provider, wallet } from "../client";
import dotenv from "dotenv";
dotenv.config();

import exmapleSchema from "../../resources/nft-schema.json"
let schema_name = exmapleSchema.code;
const split_schema = schema_name.split(".")
const _name = split_schema[1]
const org_name = process.env.ORG_NAME
let schemaCode: string;
schemaCode = `${org_name}.${_name}`;

console.log(schemaCode);

exmapleSchema.code = schemaCode;

async function main() {
    
  const precompileContract = "0x0000000000000000000000000000000000001055";
  const nftContractAddress = "0x0ecE60ba8131f92edDF9BaC35aEd6c2EBaE95713";

  exmapleSchema.origin_data.origin_contract_address = nftContractAddress

  let encodeBase64Schema = Buffer.from(JSON.stringify(exmapleSchema)).toString(
    "base64"
  );

  const nftmngr = new ethers.Contract(
    precompileContract,
    NFTMngrABI.abi,
    provider,
  );


  // @ts-ignore
  const tx = await nftmngr.connect(wallet).createSchema(encodeBase64Schema)

  console.log("Transaction hash:", tx.hash);

  // Wait for the transaction to be mined
  const receipt = await tx.wait();

  // Get gas used
  const gasUsed = receipt.gasUsed;

  console.log(`Gas used: ${gasUsed.toString()}`);
  console.log("TX Info", tx);
}

main()
  .then(() => {
    console.log("Transaction complete.");
  })
  .catch((err) => {
    console.error(err);
  });
