import {
  DirectSecp256k1HdWallet,
  Registry,
  OfflineSigner,
  EncodeObject,
} from "@cosmjs/proto-signing";
import {
  Secp256k1HdWallet,
  createMultisigThresholdPubkey,
  encodeSecp256k1Pubkey,
  makeCosmoshubPath,
  pubkeyToAddress,
  AminoMsg,
} from "@cosmjs/amino";
import { TxRaw } from "cosmjs-types/cosmos/tx/v1beta1/tx";
import {
  SigningStargateClient,
  StargateClient,
  makeMultisignedTx,
  SignerData,
  AminoTypes,
} from "@cosmjs/stargate";
import { SigningCosmWasmClient, CosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { coins } from "@cosmjs/proto-signing";
import { encode } from "uint8-to-base64";
import { MsgCreateBinding, createBindingAdminoTypes } from "./evmbindign";

const alice_mnemonic =
  "history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind";
const bob_mnemonic =
  "limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture";
const rpc = "http://localhost:26657";

const customRegistry = new Registry();
customRegistry.register(
  "/thesixnetwork.sixprotocol.evmbind.MsgCreateBinding",
  MsgCreateBinding
);

const alice_binding = async () => { const alice_wallet = await DirectSecp256k1HdWallet.fromMnemonic(alice_mnemonic, {
    prefix: "6x",
  });
  const [alice_account] = await alice_wallet.getAccounts();

  const MsgCreateBinding: MsgCreateBinding = {
    creator: alice_account.address,
    ethAddress: "0xa3121CDcc7d89be798BF8BE6efb75ab02B761a96",
    ethSignature: "0xeaddadeb5aeadeef509585794043bfd9919cc90ac3859d9796a595ee56fab34317dfd843a26666373c72a2a7a052e46cba1330a4039a929dc5e9e6de516e004f1b",
    signMessage: "register"
  }

  const msgCreateBind = {
    typeUrl:"/thesixnetwork.sixprotocol.evmbind.MsgCreateBinding",
    value: MsgCreateBinding,
  };


  const gasLimit = 200000;
  const fee = {
    amount: coins("250000", "usix"),
    gas: gasLimit.toString(),
  };
  
  const alice_client = await SigningCosmWasmClient.connectWithSigner(
    rpc,
    alice_wallet
  );
  const alice_info = await alice_client.getAccount(alice_account.address);
  // const alice_chainId = await alice_client.getChainId();

  // const alice_signerData = {
  //   accountNumber: alice_info.accountNumber,
  //   sequence: alice_info.sequence,
  //   chainId: alice_chainId,
  // };
const client = await SigningCosmWasmClient.connectWithSigner(rpc, alice_wallet,{
    registry: customRegistry,
    aminoTypes: new AminoTypes(createBindingAdminoTypes()),
});
const sendData = await client.signAndBroadcast(
  alice_info.address,
  [msgCreateBind],
  fee,
  ""
);
console.log(sendData);
return sendData;
};

const sending = async () => {
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(alice_mnemonic, {
    prefix: "6x",
  });

  const [firstAccount] = await wallet.getAccounts();
  const msgSend = {
    fromAddress: firstAccount.address,
    toAddress: "6x1kcgtdfl703a2xtuarpglrrre84mvhmxh7dsgml",
    amount: coins("10000000", "usix"),
  };
  const msg = {
    typeUrl: "/cosmos.bank.v1beta1.MsgSend",
    value: msgSend,
  };
  const gasLimit = 200000;
  const fee = {
    amount: coins("250000", "usix"),
    gas: gasLimit.toString(),
  };
  const client = await SigningStargateClient.connectWithSigner(rpc, wallet);
  const sendData = await client.signAndBroadcast(
    firstAccount.address,
    [msg],
    fee,
    "",
  );
  console.log(sendData);
  return sendData;
};

alice_binding();