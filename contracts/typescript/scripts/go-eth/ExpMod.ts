import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const EXPMOD_ADDRESS = "0x0000000000000000000000000000000000000005";

function padTo32Bytes(value: bigint): string {
    return ethers.zeroPadValue(ethers.toBeHex(value), 32);
}

async function TestExpMod() {
    // Define base, exponent, and modulus
    const base = 3n;
    const exponent = 7n;
    const modulus = 11n;

    // Prepare the input data
    const baseLength = padTo32Bytes(32n);
    const exponentLength = padTo32Bytes(32n);
    const modulusLength = padTo32Bytes(32n);

    const baseHex = padTo32Bytes(base);
    const exponentHex = padTo32Bytes(exponent);
    const modulusHex = padTo32Bytes(modulus);

    const inputHex = ethers.concat([
        baseLength,
        exponentLength,
        modulusLength,
        baseHex,
        exponentHex,
        modulusHex
    ]);

    try {
        // Call the ExpMod precompiled contract
        const result = await provider.call({
            to: EXPMOD_ADDRESS,
            data: inputHex
        });

        console.log('Input:');
        console.log('  Base:', base.toString());
        console.log('  Exponent:', exponent.toString());
        console.log('  Modulus:', modulus.toString());

        // Convert result to BigInt
        const resultBigInt = BigInt(result);
        console.log('Output from precompile:', resultBigInt.toString());

        // Compute locally for comparison
        const expectedResult = base ** exponent % modulus;
        console.log('Expected result:', expectedResult.toString());

        // Compare the results
        console.log('ExpMod test ' + (resultBigInt === expectedResult ? 'PASSED' : 'FAILED'));

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: EXPMOD_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling ExpMod contract:', error);
    }
}

// Call the function
TestExpMod().catch(console.error);