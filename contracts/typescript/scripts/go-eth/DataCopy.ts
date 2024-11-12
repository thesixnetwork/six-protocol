import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const IDENTITY_ADDRESS = "0x0000000000000000000000000000000000000004";

async function TestIdentity() {
    // Input data to copy
    const input = "Hello, Ethereum!";

    // Convert input to bytes, then to hex string
    const inputBytes = ethers.toUtf8Bytes(input);
    const inputHex = ethers.hexlify(inputBytes);

    try {
        // Call the Identity precompiled contract
        const result = await provider.call({
            to: IDENTITY_ADDRESS,
            data: inputHex
        });

        console.log('Input:', input);
        console.log('Input (hex):', inputHex);
        console.log('Output from precompile:', result);

        // Convert result back to string for human-readable comparison
        const outputString = ethers.toUtf8String(result);
        console.log('Output (decoded):', outputString);

        // Compare the results
        console.log('Identity test ' + (inputHex === result ? 'PASSED' : 'FAILED'));

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: IDENTITY_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling Identity contract:', error);
    }
}

// Call the function
TestIdentity().catch(console.error);