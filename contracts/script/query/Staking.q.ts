import { ethers } from "ethers";
import IStakingABI from "../../out/IStaking.sol/IStaking.json";
import { JsonRpcProvider, Wallet, formatEther } from "ethers";

const provider = new JsonRpcProvider("https://rpc-evm.fivenet.sixprotocol.net");

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
  const delegator = "6x1l0wacktp0q8y6wvcp7vwk5ujcc8lf9u3quaxl2";
  const validator = "6xvaloper1fl9ypcr9al7w2294adtla42njc0qnws66gdv73";

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
