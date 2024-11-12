import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import { createHash } from "crypto";
import dotenv from "dotenv";
dotenv.config();

const RIPEMD160_ADDRESS = "0x0000000000000000000000000000000000000003";

async function TestRIPEMD160() {
    // Input data to hash
    const input = "Hello, Ethereum!";

    // Convert input to bytes, then to hex string
    const inputBytes = ethers.toUtf8Bytes(input);
    const inputHex = ethers.hexlify(inputBytes);

    try {
        // Call the RIPEMD-160 precompiled contract
        const result = await provider.call({
            to: RIPEMD160_ADDRESS,
            data: inputHex
        });

        console.log('Input:', input);
        console.log('RIPEMD-160 hash (from precompile):', result);

        // Compute the hash locally for comparison
        const localHash = createHash('ripemd160').update(Buffer.from(inputBytes)).digest('hex');
        console.log('RIPEMD-160 hash (computed locally):', '0x' + localHash);

        // Extract the actual hash from the precompile result (last 40 characters)
        const precompileHash = '0x' + result.slice(-40);
        console.log('Extracted hash from precompile:', precompileHash);

        // Compare the results
        console.log('RIPEMD-160 test ' + (precompileHash === '0x' + localHash ? 'PASSED' : 'FAILED'));

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: RIPEMD160_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling RIPEMD-160:', error);
    }
}

// Call the function
TestRIPEMD160().catch(console.error);