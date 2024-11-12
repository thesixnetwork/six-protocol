import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const ECRECOVER_ADDRESS = "0x0000000000000000000000000000000000000001";

async function TestECrecover() {
    const message = "Hello, Ethereum!";

    // Hash the message
    const messageHash = ethers.hashMessage(message);

    // Sign the message
    const signature = await wallet.signMessage(message);

    // Split the signature
    const sig = ethers.Signature.from(signature);

    // Prepare the input for ECRECOVER
    const input = ethers.concat([
        messageHash,
        ethers.zeroPadValue(ethers.toBeHex(sig.v), 32),
        sig.r,
        sig.s
    ]);

    try {
        // Call the ECRECOVER precompiled contract
        const result = await provider.call({
            to: ECRECOVER_ADDRESS,
            data: input
        });

        // The result is a 32-byte value, where the last 20 bytes are the recovered address
        const recoveredAddress = ethers.getAddress('0x' + result.slice(-40));

        console.log('Original address:', wallet.address);
        console.log('Recovered address:', recoveredAddress);
        console.log('ECRECOVER test ' + (wallet.address.toLowerCase() === recoveredAddress.toLowerCase() ? 'PASSED' : 'FAILED'));

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: ECRECOVER_ADDRESS,
            data: input
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling ECRECOVER:', error);
    }
}

// Call the function
TestECrecover().then((res) => {
    console.log
}).catch(console.error);