import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const BN256SCALARMUL_ADDRESS = "0x0000000000000000000000000000000000000007";

function padTo32Bytes(value: bigint): string {
    return ethers.zeroPadValue(ethers.toBeHex(value), 32);
}

async function TestBN256ScalarMul() {
    // Define a point on the curve and a scalar
    const x = 1n;
    const y = 2n;
    const scalar = 3n;

    // Prepare the input data
    const inputHex = ethers.concat([
        padTo32Bytes(x),
        padTo32Bytes(y),
        padTo32Bytes(scalar)
    ]);

    try {
        // Call the BN256ScalarMul precompiled contract
        const result = await provider.call({
            to: BN256SCALARMUL_ADDRESS,
            data: inputHex
        });

        console.log('Input:');
        console.log(`  Point: (${x}, ${y})`);
        console.log(`  Scalar: ${scalar}`);

        // The result should be a 64-byte value representing the x and y coordinates of the resulting point
        const resultX = BigInt(`0x${result.slice(2, 66)}`);
        const resultY = BigInt(`0x${result.slice(66)}`);

        console.log('Output from precompile:');
        console.log(`  Resulting Point: (${resultX}, ${resultY})`);

        // Simple (but incorrect) local computation for demonstration
        const simpleX = x * scalar;
        const simpleY = y * scalar;

        console.log('Output from simple local computation (not accurate):');
        console.log(`  Resulting Point: (${simpleX}, ${simpleY})`);

        console.log('Comparison (Note: local computation is not accurate):');
        console.log(`  Precompile X: ${resultX}`);
        console.log(`  Simple X:     ${simpleX}`);
        console.log(`  Precompile Y: ${resultY}`);
        console.log(`  Simple Y:     ${simpleY}`);

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: BN256SCALARMUL_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling BN256ScalarMul contract:', error);
    }
}

// Call the function
TestBN256ScalarMul().catch(console.error);