import { ethers } from "ethers";
import IBankABI from "../../out/IBank.sol/IBank.json";

import { JsonRpcProvider, Wallet, formatEther } from "ethers";
const provider = new JsonRpcProvider("http://localhost:8545");

async function main() {
  const bankPrecompileContract = "0x0000000000000000000000000000000000001001";

  const toCheckAddress = "0xB62ef83643A2F8c95dF78F694C6Bf480f5b786f2";
  const evmDenom = "asix";

  const bankContract = new ethers.Contract(bankPrecompileContract, IBankABI.abi, provider);

  const balance = await bankContract.balance(toCheckAddress, evmDenom);
  console.log(`Balance for asix: ${balance.toString()}`);

  const allBalances = await bankContract.all_balances(toCheckAddress);
  console.log(`All balance: ${allBalances}`);

  const name = await bankContract.name(evmDenom);
  console.log(`Name of EVM Denom: ${name}`);

  const symbol = await bankContract.symbol(evmDenom);
  console.log(`Symbol of EVM Denom: ${symbol}`);

  const decimal = await bankContract.decimals(evmDenom);
  console.log(`Decimal of EVM Denom: ${decimal}`);

  const supply = await bankContract.supply(evmDenom);
  console.log(`Supply of EVM Denom: ${supply}`);

  return;
}

main()
  .then(() => {
    console.log;
  })
  .catch((err) => {
    console.log(err);
  });
