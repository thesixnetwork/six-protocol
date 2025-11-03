import { ethers } from "ethers";
import INFTABI from "../../out/INFTManager.sol/INFTMNGR.json";

import { JsonRpcProvider, Wallet, formatEther } from "ethers";
const provider = new JsonRpcProvider("http://localhost:8545");

async function main() {
  const nftPrecompileContract = "0x0000000000000000000000000000000000001055";

  const nftSchemaName = "six-protocol.example";
  const tokenId = "1";
  const attributeName = "points";
  const ownerAddress = "0xd907f36f7D83344057a619b6D83A45B3288c3c21";

  const nftManagerContract = new ethers.Contract(nftPrecompileContract, INFTABI.abi, provider);

  const attributeValue = await nftManagerContract.getAttributeValue(nftSchemaName, tokenId, attributeName);
  console.log(`Value of ${attributeName}: ${attributeValue}`);

  const isExecutor = await nftManagerContract.isActionExecutor(nftSchemaName, ownerAddress);
  console.log(`Is ${ownerAddress} an executor?: ${isExecutor}`);

  const isSchemaOwner = await nftManagerContract.isSchemaOwner(nftSchemaName, ownerAddress);
  console.log(`Is ${ownerAddress} an owner of ${nftSchemaName}?: ${isSchemaOwner}`);

  return;
}

main()
  .then(() => {
    console.log;
  })
  .catch((err) => {
    console.log(err);
  });
