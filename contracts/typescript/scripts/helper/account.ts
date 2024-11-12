import * as bip39 from "bip39";
import crypto from "crypto";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { EnglishMnemonic } from "@cosmjs/crypto";
import { ethers } from "ethers";
import { fromBech32, toHex, toBech32, fromHex } from "@cosmjs/encoding";


export function generateCosmosMnemonic(): string {
  const mnemonic = bip39.generateMnemonic(256);
  const englishMnemonic = new EnglishMnemonic(mnemonic);
  if (!englishMnemonic) {
    throw new Error("Failed to generate a valid Cosmos mnemonic");
  }
  return mnemonic;
}

export async function addressFromMnemonic(prefix: string, mnemonic: string) {
  const hdWallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: prefix,
  });

  const addressInfo = await hdWallet.getAccounts().then((res) => {
    return res;
  });

  return addressInfo;
}

export function cosmosToEthereumAddress(cosmosAddress: string): string {
  // Step 1: Convert Cosmos Bech32 address to bytes
  const { data } = fromBech32(cosmosAddress);

  // Step 2: Convert bytes to hex and apply EIP-55 checksum formatting
  const ethereumAddress = ethers.getAddress(toHex(data));
  
  return ethereumAddress;
}


export function ethereumToBech32(prefix: string, ethAddress: string): string {
  // Remove "0x" prefix if present
  const cleanHexAddress = ethAddress.startsWith("0x") ? ethAddress.slice(2) : ethAddress;

  // Convert the hex string to bytes
  const addressBytes = fromHex(cleanHexAddress);

  // Encode the bytes as a Bech32 address with the given prefix
  const bech32Address = toBech32(prefix, addressBytes);

  return bech32Address;
}