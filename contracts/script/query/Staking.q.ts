import { ethers } from "ethers";
import IStakingABI from "../../out/IStaking.sol/IStaking.json";
import { JsonRpcProvider, Wallet, formatEther } from "ethers";

const provider = new JsonRpcProvider("http://localhost:8545");

interface Balance {
  amount: bigint;
  denom: string;
}

interface DelegationDetails {
  delegator_address: string;
  shares: bigint;
  decimals: bigint;
  validator_address: string;
}

interface Delegation {
  balance: Balance;
  delegation: DelegationDetails;
}

async function main() {
  const stakingPrecompileContract = "0x0000000000000000000000000000000000001005";
  const delegator = "6x1kch0sdjr5tuvjh0h3a55c6l5sr6m0phjeag9f2";
  const validator = "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh";

  const stakingContract = new ethers.Contract(stakingPrecompileContract, IStakingABI.abi, provider);

  try {
    const delegation = await stakingContract.delegation(delegator, validator);

    // Cast to our TypeScript interface
    const typedDelegation = delegation as Delegation;

    // Format the output nicely
    console.log("Delegation Details:");
    console.log("-----------------");
    console.log(`Amount: ${typedDelegation.balance.amount} (${typedDelegation.balance.denom})`);
    console.log(`Delegator: ${typedDelegation.delegation.delegator_address}`);
    console.log(`Validator: ${typedDelegation.delegation.validator_address}`);
    console.log(`Shares: ${typedDelegation.delegation.shares}`);
    console.log(`Decimals: ${typedDelegation.delegation.decimals}`);
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
