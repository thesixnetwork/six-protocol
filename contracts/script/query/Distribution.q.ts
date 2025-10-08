import { ethers } from "ethers";
import IDistABI from "../../out/IDistribution.sol/IDistr.json";
import { JsonRpcProvider, Wallet, formatUnits } from "ethers";

const provider = new JsonRpcProvider("http://localhost:8545");

async function main() {
  const distributionPrecompileContract = "0x0000000000000000000000000000000000001007";
  const delegator = "6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp";
  const validator = "6xvaloper13g50hqdqsjk85fmgqz2h5xdxq49lsmjdz3mr76";

  const DISPLAY_DECIMALS = 18;
  const USIX_DECIMALS = 6;
  const ASIX_DECIMALS = 18;

  // Function to format token amounts for display
  function formatTokenAmount(amount: bigint, denom: string): string {
    return formatUnits(amount, DISPLAY_DECIMALS);
  }

  function calculateActualTokenValue(amount: bigint, denom: string): number {
    if (denom === "usix") {
      return Number(formatUnits(amount, USIX_DECIMALS));
    } else if (denom === "asix") {
      return Number(formatUnits(amount, ASIX_DECIMALS));
    } else {
      return Number(formatUnits(amount, 18));
    }
  }

  const distributionContract = new ethers.Contract(distributionPrecompileContract, IDistABI.abi, provider);

  try {
    console.log("\n=== VALIDATOR SPECIFIC REWARDS ===");
    try {
      const validatorRewards = await distributionContract.rewards(validator, delegator);
      console.log("Rewards from validator", validator, ":");
      console.log("-----------------");

      let totalConvertedValue = 0;

      if (Array.isArray(validatorRewards[0])) {
        validatorRewards[0].forEach((coin: any, index: number) => {
          const denom = coin[1];
          const amount = coin[0];

          const formattedAmount = formatTokenAmount(amount, denom);
          console.log(`${index + 1}. ${formattedAmount} ${denom}`);

          const actualValue = calculateActualTokenValue(amount, denom);
          totalConvertedValue += actualValue;
        });
      }

      console.log("-----------------");
      console.log("Validator Address:", validatorRewards[1]);
      console.log("Total after conversion:", totalConvertedValue.toFixed(6), "tokens");
    } catch (error) {
      console.error("Error querying validator rewards:", error);
    }

    console.log("\n=== ALL DELEGATOR REWARDS ===");
    const rawRewards = await distributionContract.allRewards(delegator);

    if (Array.isArray(rawRewards[0])) {
      console.log("Rewards by validator:");
      console.log("-----------------");

      let totalByValidatorMap = new Map<string, number>();

      rawRewards[0].forEach((reward: any) => {
        const validatorAddr = reward[1];
        console.log(`Validator: ${validatorAddr}`);

        let validatorTotal = 0;

        if (Array.isArray(reward[0])) {
          reward[0].forEach((coin: any) => {
            const denom = coin[1];
            const amount = coin[0];

            const formattedAmount = formatTokenAmount(amount, denom);
            console.log(`  ${formattedAmount} ${denom}`);

            const actualValue = calculateActualTokenValue(amount, denom);
            validatorTotal += actualValue;
          });
        }

        totalByValidatorMap.set(validatorAddr, validatorTotal);
        console.log(`  Total after conversion: ${validatorTotal.toFixed(6)} tokens`);
      });
    }

    if (Array.isArray(rawRewards[1])) {
      console.log("\nTotal rewards across all validators:");
      console.log("-----------------");

      let grandTotal = 0;

      rawRewards[1].forEach((coin: any) => {
        const denom = coin[1];
        const amount = coin[0];

        const formattedAmount = formatTokenAmount(amount, denom);
        console.log(`  ${formattedAmount} ${denom}`);

        const actualValue = calculateActualTokenValue(amount, denom);
        grandTotal += actualValue;
      });

      console.log("-----------------");
      console.log(`GRAND TOTAL after conversion: ${grandTotal.toFixed(6)} tokens`);
    }
  } catch (error) {
    console.error("Error in distribution operations:", error);
  }
}

main()
  .then(() => {
    console.log("\nScript completed successfully");
  })
  .catch((err) => {
    console.error("Script failed:", err);
  });
