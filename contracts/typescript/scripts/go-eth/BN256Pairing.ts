import { ethers } from "ethers";
import { provider, wallet } from "../../client";
import dotenv from "dotenv";
dotenv.config();

const BN256PAIRING_ADDRESS = "0x0000000000000000000000000000000000000008";

function padTo32Bytes(value: string): string {
    return ethers.zeroPadValue(`0x${value}`, 32);
}

async function TestBN256Pairing() {
    // Define input points for the pairing check
    // This is a valid pairing input that should return true
    const input = [
        "1c76476f4def4bb94541d57ebba1193381ffa7aa76ada664dd31c16024c43f59",
        "3034dd2920f673e204fee2811c678745fc819b55d3e9d294e45c9b03a76aef41",
        "209dd15ebff5d46c4bd888e51a93cf99a7329636c63514396b4a452003a35bf7",
        "04bf11ca01483bfa8b34b43561848d28905960114c8ac04049af4b6315a41678",
        "2bb8324af6cfc93537a2ad1a445cfd0ca2a71acd7ac41fadbf933c2a51be344d",
        "120a2a4cf30c1bf9845f20c6fe39e07ea2cce61f0c9bb048165fe5e4de877550",
        "111e129f1cf1097710d41c4ac70fcdfa5ba2023c6ff1cbeac322de49d1b6df7c",
        "2032c61a830e3c17286de9462bf242fca2883585b93870a73853face6a6bf411",
        "198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2",
        "1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed",
        "090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b",
        "12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
    ];

    // Prepare the input data
    const inputHex = ethers.concat(input.map(padTo32Bytes));

    try {
        // Call the BN256Pairing precompiled contract
        const result = await provider.call({
            to: BN256PAIRING_ADDRESS,
            data: inputHex
        });

        console.log('Input: Valid pairing input (12 32-byte values)');

        // The result should be a 32-byte value
        // If it's 1, the pairing check succeeded. If it's 0, the check failed.
        const pairingResult = BigInt(result);

        console.log('Output from precompile:');
        console.log(`  Pairing check result: ${pairingResult}`);
        console.log(`  Pairing ${pairingResult === 1n ? 'succeeded' : 'failed'}`);

        // Estimate gas for the transaction
        const gasEstimate = await provider.estimateGas({
            to: BN256PAIRING_ADDRESS,
            data: inputHex
        });

        console.log(`Estimated gas: ${gasEstimate.toString()}`);

        // Get the transaction count (nonce) for the wallet
        const nonce = await provider.getTransactionCount(wallet.address);
        console.log(`Current nonce: ${nonce}`);

    } catch (error) {
        console.error('Error calling BN256Pairing contract:', error);
    }
}

// Call the function
TestBN256Pairing().catch(console.error);