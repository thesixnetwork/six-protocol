import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const BN256ADD_ADDRESS = "0x0000000000000000000000000000000000000006";

function padTo32Bytes(value: bigint): string {
    return ethers.zeroPadValue(ethers.toBeHex(value), 32);
}

async function TestBN256Add() {
    // Define two points on the curve
    // Point 1: (1, 2)
    const x1 = padTo32Bytes(1n);
    const y1 = padTo32Bytes(2n);
    // Point 2: (1, 2) // Adding the same point for simplicity
    const x2 = padTo32Bytes(1n);
    const y2 = padTo32Bytes(2n);

    // Prepare the input data
    const inputHex = ethers.concat([x1, y1, x2, y2]);

    try {
        // Call the BN256Add precompiled contract
        const result = await provider.call({
            to: BN256ADD_ADDRESS,
            data: inputHex
        });

        console.log('Input:');
        console.log('  Point 1: (1, 2)');
        console.log('  Point 2: (1, 2)');

        // The result should be a 64-byte value representing the x and y coordinates of the resulting point
        const resultX = `0x${result.slice(2, 66)}`;
        const resultY = `0x${result.slice(66)}`;

        console.log('Output from precompile:');
        console.log('  Resulting Point: (', resultX, ',', resultY, ')');

        // Note: Validating the result requires understanding of elliptic curve arithmetic
        // For this example, we're just printing the result

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: BN256ADD_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling BN256Add contract:', error);
    }
}

// Call the function
TestBN256Add().catch(console.error);