import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const SHA256_ADDRESS = "0x0000000000000000000000000000000000000002";

async function TestSHA256() {
    // Input data to hash
    const input = "Hello, Ethereum!";

    // Convert input to bytes
    const inputBytes = ethers.toUtf8Bytes(input);
    const inputHex: string = ethers.hexlify(ethers.toUtf8Bytes(input));

    try {
        // Call the SHA256 precompiled contract
        const result = await provider.call({
            to: SHA256_ADDRESS,
            data: inputHex,
        });

        console.log('Input:', input);
        console.log('SHA256 hash (from precompile):', result);

        // Compute the hash locally for comparison
        const localHash = ethers.sha256(inputBytes);
        console.log('SHA256 hash (computed locally):', localHash);

        // Compare the results
        console.log('SHA256 test ' + (result === localHash ? 'PASSED' : 'FAILED'));

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: SHA256_ADDRESS,
            data: inputHex,
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling SHA256:', error);
    }
}

// Call the function
TestSHA256().then(() => { console.log }).catch(console.error);