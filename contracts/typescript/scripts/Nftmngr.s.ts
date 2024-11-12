import { ethers } from "ethers";
import NFTMngrABI from "../../out/INftmngr.sol/INftmngr.json";
import { provider, wallet } from "../client";

async function main() {
  const precompileContract = "0x0000000000000000000000000000000000000055";
  const nftContractAddress = "0x3753C81072A56072840990D3D02f354Efb7425A3";
  const actionName = "attend_stage";
  const tokenId = "4";
  const nftSchema = "TechSauce.GlobalSummit2024";
  const refId = "4";
  const jsonParams = '[{"name":"stage","value":"stage_2"}]';

  const nftmngr = new ethers.Contract(
    precompileContract,
    NFTMngrABI.abi,
    provider,
  );

  // @ts-ignore
  const tx = await nftmngr.connect(wallet).actionByAdmin(nftContractAddress, nftSchema, tokenId, actionName, refId, jsonParams)

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
