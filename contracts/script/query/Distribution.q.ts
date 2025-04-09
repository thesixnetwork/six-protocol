import { ethers } from "ethers";
import IDistABI from "../../out/IDistribuion.sol/IDistr.json";
import { JsonRpcProvider, Wallet, formatEther } from "ethers";

const provider = new JsonRpcProvider("http://localhost:8545");

interface Coin {
  amount: bigint;
  decimal: bigint;
  denom: string;
}

interface Reward {
  coin: Array<Coin>;
  validator_address: string;
}

interface Rewards {
  reward: Array<Reward>;
  coin: Array<Coin>;
}
async function main() {
  const distributionPrecompileContract =
    "0x0000000000000000000000000000000000001007";
  const delegator = "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv";
  const validator = "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh";

  const distributionContract = new ethers.Contract(
    distributionPrecompileContract,
    IDistABI.abi,
    provider,
  );

  try {
    const rawReward = await distributionContract.rewards(validator, delegator);
    console.log(rawReward);
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
