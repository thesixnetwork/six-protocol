import { ethers } from "ethers";
import { JsonRpcProvider, Wallet } from "ethers";
import NFTMngrABI from "../../out/INFTManager.sol/INFTMNGR.json";
import proposal from "../../../resources/vitual_schema_prop.json";
import dotenv from "dotenv";
dotenv.config();


const provider = new JsonRpcProvider("http://localhost:8545");

if (!process.env.PRIVATE_KEY) {
  console.error("TEST_ACCOUNT private key undefined in .env");
  process.exit();
}

export const wallet = new Wallet(process.env.PRIVATE_KEY, provider);

async function main() {
  const precompileContract = "0x0000000000000000000000000000000000001055";
  const nftmngr = new ethers.Contract(
    precompileContract,
    NFTMngrABI.abi,
    provider,
  );

  const base64VirtualSchema = Buffer.from(JSON.stringify(proposal)).toString("base64");

  const tx = await nftmngr
    .connect(wallet)
    // @ts-ignore
    .virtualSchemaProposal(0, base64VirtualSchema);

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
