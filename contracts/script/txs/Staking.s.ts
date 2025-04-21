import { ethers } from "ethers";
import IStakingABI from "../../out/IStaking.sol/IStaking.json";
import { JsonRpcProvider, Wallet, formatEther, parseEther } from "ethers";
import dotenv from "dotenv";
dotenv.config();

const provider = new JsonRpcProvider("http://localhost:8545");
if (!process.env.PRIVATE_KEY) {
  console.error("TEST_ACCOUNT private key undefined in .env");
  process.exit();
}

export const wallet = new Wallet(process.env.PRIVATE_KEY, provider);

async function main() {
  const stakingPrecompileContract =
    "0x0000000000000000000000000000000000001005";
  const validator = "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh";
  const amount = parseEther("10");
  console.log(amount)

  const stakingContract = new ethers.Contract(
    stakingPrecompileContract,
    IStakingABI.abi,
    provider,
  );

  try {
    const tx = await stakingContract
      .connect(wallet)
      // @ts-ignore
      .delegate(validator, amount);
    // Wait for the transaction to be mined
    const receipt = await tx.wait();

    // Get gas used
    const gasUsed = receipt.gasUsed;

    console.log(`Gas used: ${gasUsed.toString()}`);
    console.log("TX Info", tx);
    return;
  } catch (error) {
    console.error("Error querying delegation:", error);
  }

  return;
}

main()
  .then(() => {
    console.log("Script completed successfully");
  })
  .catch((err) => {
    console.error("Script failed:", err);
  });
