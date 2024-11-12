import { TokenboundClient, TBAccountParams } from "@tokenbound/sdk";
import { JsonRpcProvider, Wallet, formatEther } from "ethers";
import { CONSTANTS } from "./constants";
const {
  RPC,
} = CONSTANTS;

export const provider = new JsonRpcProvider(RPC);
import dotenv from "dotenv";
dotenv.config()

if (!process.env.TEST_ACCOUNT) {
  console.error("TEST_ACCOUNT private key undefined in .env");
  process.exit();
}

export const wallet = new Wallet(process.env.TEST_ACCOUNT, provider);


export async function getConnectorConfig(
  network: string,
): Promise<{ rpcUrl: string; apiUrl: string; mnemonic: string }> {
  switch (network) {
    case "local":
      return {
        rpcUrl: "http://localhost:26657",
        apiUrl: "http://localhost:1317",
        mnemonic: process.env.ALICE_MNEMONIC!,
      };
    case "fivenet":
      return {
        rpcUrl: "https://rpc1.fivenet.sixprotocol.net:443",
        apiUrl: "https://api1.fivenet.sixprotocol.net:443",
        mnemonic: process.env.TECHSAUCE_TESTNET_MNEMONIC!,
      };
    case "sixnet":
      return {
        rpcUrl: "https://sixnet-rpc.sixprotocol.net:443",
        apiUrl: "https://sixnet-api.sixprotocol.net:443",
        mnemonic: process.env.TECHSAUCE_MNEMONIC!,
      };
    default:
      throw new Error("Invalid network");
  }
}
