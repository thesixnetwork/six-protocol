import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const KZG_POINT_EVALUATION_ADDRESS = "0x0000000000000000000000000000000000000014";

function padTo32Bytes(value: string): string {
    return ethers.zeroPadValue(`0x${value}`, 32);
}

async function TestKZGPointEvaluation() {
    // Sample input data (this is just an example and may not be a valid KZG proof)
    const versioned_hash = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef";
    const z = "1234567890123456789012345678901234567890123456789012345678901234";
    const y = "2345678901234567890123456789012345678901234567890123456789012345";
    const commitment = "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001";
    const proof = "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002";

    // Prepare the input data
    const inputHex = ethers.concat([
        padTo32Bytes(versioned_hash),
        padTo32Bytes(z),
        padTo32Bytes(y),
        `0x${commitment}`,  // Use 0x prefix for these larger values
        `0x${proof}`        // Use 0x prefix for these larger values
    ]);

    try {
        // Call the KZG Point Evaluation precompiled contract
        const result = await provider.call({
            to: KZG_POINT_EVALUATION_ADDRESS,
            data: inputHex
        });

        console.log('Input:');
        console.log(`  Versioned Hash: 0x${versioned_hash}`);
        console.log(`  z: 0x${z}`);
        console.log(`  y: 0x${y}`);
        console.log(`  Commitment: 0x${commitment}`);
        console.log(`  Proof: 0x${proof}`);

        console.log('Output from precompile:');
        console.log(`  Result: ${result}`);

        // Interpret the result
        if (result === '0x01') {
            console.log('KZG Point Evaluation verified successfully');
        } else if (result === '0x00') {
            console.log('KZG Point Evaluation failed verification');
        } else {
            console.log('Unexpected result from KZG Point Evaluation');
        }

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: KZG_POINT_EVALUATION_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling KZG Point Evaluation contract:', error);
    }
}

// Call the function
TestKZGPointEvaluation().catch(console.error);