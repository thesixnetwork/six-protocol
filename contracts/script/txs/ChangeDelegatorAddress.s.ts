/**
 * ChangeDelegatorAddress Test Script
 * 
 * This script tests the changeDelegatorAddress function in the TokenFactory precompile.
 * The function allows changing the delegator address for staking operations.
 * 
 * Prerequisites:
 * 1. Set PRIVATE_KEY in your .env file
 * 2. Optionally set OLD_DELEGATOR_ADDRESS and NEW_DELEGATOR_ADDRESS in .env
 * 3. Ensure your local blockchain is running on http://localhost:8545
 * 
 * Usage:
 * npm run ts-node contracts/script/txs/ChangeDelegatorAddress.s.ts
 * 
 * Environment Variables:
 * - PRIVATE_KEY: Required - Private key for the wallet to execute the transaction
 * - OLD_DELEGATOR_ADDRESS: Optional - Address to change from (defaults to example address)
 * - NEW_DELEGATOR_ADDRESS: Optional - Address to change to (defaults to example address)
 */

import { ethers } from "ethers";
import { JsonRpcProvider, Wallet } from "ethers";
import ITokenFactoryABI from "../../out/ITokenFactory.sol/ITokenFactory.json";
import dotenv from "dotenv";
dotenv.config();

// TokenFactory ABI - only includes the changeDelegatorAddress function
const TokenFactoryABI = ITokenFactoryABI.abi

const provider = new JsonRpcProvider("http://localhost:8545");
if (!process.env.PRIVATE_KEY) {
  console.error("PRIVATE_KEY undefined in .env");
  process.exit();
}

export const wallet = new Wallet(process.env.PRIVATE_KEY, provider);

async function main() {
  const tokenFactoryPrecompileContract = "0x0000000000000000000000000000000000001069";
  
  // You can customize these addresses for your testing needs
  // Option 1: Use specific test addresses
  // const oldDelegatorAddress = process.env.OLD_DELEGATOR_ADDRESS || "0x8D4Bb008181fd32E7bDac45666D6d2066781B974";
  // const newDelegatorAddress = process.env.NEW_DELEGATOR_ADDRESS || "0x25271cfc8C7EFaFBc33a357b2246369783AEae80";
  
  // Option 2: Use wallet address as one of the addresses for testing
  const oldDelegatorAddress = wallet.address;
  const newDelegatorAddress = "0x25271cfc8C7EFaFBc33a357b2246369783AEae80";
  
  console.log("=== ChangeDelegatorAddress Test ===");
  console.log("TokenFactory Precompile:", tokenFactoryPrecompileContract);
  console.log("Old Delegator Address:", oldDelegatorAddress);
  console.log("New Delegator Address:", newDelegatorAddress);
  console.log("Wallet Address:", wallet.address);

  const tokenFactoryContract = new ethers.Contract(
    tokenFactoryPrecompileContract,
    TokenFactoryABI,
    provider,
  );

  // Validate addresses
  if (!ethers.isAddress(oldDelegatorAddress)) {
    throw new Error(`Invalid old delegator address: ${oldDelegatorAddress}`);
  }
  if (!ethers.isAddress(newDelegatorAddress)) {
    throw new Error(`Invalid new delegator address: ${newDelegatorAddress}`);
  }
  if (oldDelegatorAddress === newDelegatorAddress) {
    throw new Error("Old and new delegator addresses cannot be the same");
  }

  try {
    console.log("\n=== Executing changeDelegatorAddress ===");
    console.log("Estimating gas...");
    
    // Estimate gas first
    const gasEstimate = await tokenFactoryContract
      .connect(wallet)
      // @ts-ignore
      .changeDelegatorAddress.estimateGas(oldDelegatorAddress, newDelegatorAddress);
    
    console.log(`Estimated gas: ${gasEstimate.toString()}`);
    
    console.log("Sending transaction...");
    
    const tx = await tokenFactoryContract
      .connect(wallet)
      // @ts-ignore
      .changeDelegatorAddress(oldDelegatorAddress, newDelegatorAddress);
    
    console.log("Transaction sent:", tx.hash);
    console.log("Waiting for confirmation...");
    
    // Wait for the transaction to be mined
    const receipt = await tx.wait();

    // Get gas used
    const gasUsed = receipt.gasUsed;

    console.log("\n=== Transaction Results ===");
    console.log(`Status: ${receipt.status === 1 ? 'SUCCESS' : 'FAILED'}`);
    console.log(`Gas used: ${gasUsed.toString()}`);
    console.log(`Gas estimate: ${gasEstimate.toString()}`);
    console.log(`Gas efficiency: ${((Number(gasUsed) / Number(gasEstimate)) * 100).toFixed(2)}%`);
    console.log(`Block number: ${receipt.blockNumber}`);
    console.log(`Transaction hash: ${receipt.hash}`);
    
    return receipt.status === 1;
  } catch (error: any) {
    console.error("\n=== Error Details ===");
    console.error("Error changing delegator address:", error.message);
    if (error.code) {
      console.error("Error code:", error.code);
    }
    if (error.reason) {
      console.error("Error reason:", error.reason);
    }
    if (error.transaction) {
      console.error("Failed transaction data:", error.transaction);
    }
    return false;
  }
}

main()
  .then((success) => {
    console.log("\n=== Script Completion ===");
    if (success) {
      console.log("✅ ChangeDelegatorAddress script completed successfully");
      process.exit(0);
    } else {
      console.log("❌ ChangeDelegatorAddress script failed");
      process.exit(1);
    }
  })
  .catch((err) => {
    console.error("\n=== Unexpected Error ===");
    console.error("Script failed with unexpected error:", err);
    process.exit(1);
  });
