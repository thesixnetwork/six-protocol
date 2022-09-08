import { DirectSecp256k1HdWallet , Registry ,OfflineSigner , EncodeObject} from "@cosmjs/proto-signing";
import { Secp256k1HdWallet , createMultisigThresholdPubkey, encodeSecp256k1Pubkey ,makeCosmoshubPath, pubkeyToAddress, AminoMsg, } from "@cosmjs/amino";
import { TxRaw } from "cosmjs-types/cosmos/tx/v1beta1/tx";
import {
  SigningStargateClient,
  StargateClient,
  makeMultisignedTx,
  SignerData,
  AminoTypes,
} from "@cosmjs/stargate";
import { SigningCosmWasmClient, CosmWasmClient ,} from "@cosmjs/cosmwasm-stargate";
import { coins } from "@cosmjs/proto-signing";
import { encode } from "uint8-to-base64";
import { MsgCreateBinding, createBindingAdminoTypes } from "./evmbindign";


const alice_mnemonic =
  "history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind";
const bob_mnemonic =
  "limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture";
const list_mnemonic = [alice_mnemonic, bob_mnemonic];
const rpc = "http://localhost:26657";

const multisigAccountAddress = "6x169u786wn20yz8j88wd0g6l5r2rv3dz49004u0l";
const customRegistry = new Registry();
    customRegistry.register(
      "/thesixnetwork.sixprotocol.evmbind.MsgCreateBinding",
      MsgCreateBinding
);

const signingInstruction = (async () => {
  const client = await StargateClient.connect(rpc);
  const accountOnChain = await client.getAccount(multisigAccountAddress);
  const MsgSend = {
    fromAddress: multisigAccountAddress,
    toAddress: "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv",
    amount: coins("10000000", "usix"),
  };

  const msgSend = {
    typeUrl: "/cosmos.bank.v1beta1.MsgSend",
    value: MsgSend,
  };
  const gasLimit = 200000;
  const fee = {
    amount: coins("250000", "usix"),
    gas: gasLimit.toString(),
  };

  return {
    accountNumber: accountOnChain.accountNumber,
    sequence: accountOnChain.sequence,
    chainId: await client.getChainId(),
    msgs: [msgSend],
    fee: fee,
    memo: "Use your tokens wisely",
  };
})();

const signatureSend = async() =>{
  const [
    [pubkey0, signature0, bodyBytes],
    [pubkey1, signature1],
  ] = await Promise.all(
    list_mnemonic.map(async (i) => {
      // Signing environment
      const wallet:OfflineSigner = await Secp256k1HdWallet.fromMnemonic(i, {prefix: "6x"});
      const pubkey = await encodeSecp256k1Pubkey(await (wallet.getAccounts())[0].pubkey);
      const address = (await wallet.getAccounts())[0].address;
      const signingClient = await SigningStargateClient.offline(wallet);
      const signerData: SignerData = {
        accountNumber: (await signingInstruction).accountNumber,
        sequence: (await signingInstruction).sequence,
        chainId: (await signingInstruction).chainId,
      };
      console.log("signerData", signerData);
      const { bodyBytes: bb, signatures } = await signingClient.sign(
        address,
        (await signingInstruction).msgs,
        (await signingInstruction).fee,
        (await signingInstruction).memo,
        signerData,
      );
      return [pubkey, signatures[0], bb] as const;
    }),
  );
  // console.log(bodyBytes);
  // console.log(pubkey0, signature0, pubkeyToAddress(pubkey0, "6x")); //alice
  // console.log(pubkey1, signature1, pubkeyToAddress(pubkey1, "6x")); 
  return {bodyBytes, pubkey0, signature0, pubkey1, signature1};
}

const multisigSend = async () => {
  const {pubkey0, signature0, pubkey1, signature1, bodyBytes} = await signatureSend();
  // console.log(bodyBytes);
  // console.log(pubkey0, signature0, pubkeyToAddress(pubkey0, "6x")); //alice
  // console.log(pubkey1, signature1, pubkeyToAddress(pubkey1, "6x")); 
  const multisigPubkey = await createMultisigThresholdPubkey(
    [pubkey0, pubkey1],
    2,
  );
  const multisigAddress = pubkeyToAddress(multisigPubkey, "6x");
  const address0 = pubkeyToAddress(pubkey0, "6x");
  const address1 = pubkeyToAddress(pubkey1, "6x");
  
  // const wallet = await Secp256k1HdWallet.fromMnemonic(
  //   alice_mnemonic,
  //   { prefix: "6x" }, // Replace with your own Bech32 address prefix
  // );


  //  const broadcaster = await SigningStargateClient.connectWithSigner(
  //   rpc,
  //   wallet,
  //   {
  //     registry: customRegistry,
  //     aminoTypes: new AminoTypes(createBindingAdminoTypes())
  //   }
  // )

  const broadcaster = await SigningStargateClient.connect(rpc);
  // console.log("####################",broadcaster,"########################");
  const signedTx = await makeMultisignedTx(
    multisigPubkey,
    (await signingInstruction).sequence,
    (await signingInstruction).fee,
    bodyBytes,
    new Map<string, Uint8Array>([
      [address0, signature0],
      [address1, signature1],
    ]),
  );
  // ensure signature is valid
  const result = await broadcaster.broadcastTx(Uint8Array.from(TxRaw.encode(signedTx).finish()));
  console.log(result);
}

const signingBindInstruction = (async () => {
  const client = await StargateClient.connect(rpc);
  const accountOnChain = await client.getAccount(multisigAccountAddress);

  const MsgCreateBinding: MsgCreateBinding = {
    creator: multisigAccountAddress,
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

  return {
    accountNumber: accountOnChain.accountNumber,
    sequence: accountOnChain.sequence,
    chainId: await client.getChainId(),
    msgs: [msgCreateBind],
    fee: fee,
    memo: "bindign with multisig",
  };
})();

const signatureBind = async() =>{
  const [
    [pubkey0, signature0, bodyBytes],
    [pubkey1, signature1],
  ] = await Promise.all(
    list_mnemonic.map(async (i) => {
      // Signing environment
      const wallet:OfflineSigner = await Secp256k1HdWallet.fromMnemonic(i, {prefix: "6x"});
      const pubkey = await encodeSecp256k1Pubkey((await wallet.getAccounts())[0].pubkey);
      const address = (await wallet.getAccounts())[0].address;
      const signingClient = await SigningStargateClient.offline(wallet, {
        registry: customRegistry,
        aminoTypes: new AminoTypes(createBindingAdminoTypes())
      });
      const signerData: SignerData = {
        accountNumber: (await signingBindInstruction).accountNumber,
        sequence: (await signingBindInstruction).sequence,
        chainId: (await signingBindInstruction).chainId,
      };
      // console.log("signerData", signerData);
      const { bodyBytes: bb, signatures } = await signingClient.sign(
        address,
        (await signingBindInstruction).msgs,
        (await signingBindInstruction).fee,
        (await signingBindInstruction).memo,
        signerData,
      );
      return [pubkey, signatures[0], bb] as const;
    }),
  );
  // console.log(bodyBytes);
  // console.log(pubkey0, signature0, pubkeyToAddress(pubkey0, "6x")); //alice
  // console.log(pubkey1, signature1, pubkeyToAddress(pubkey1, "6x")); 
  return {bodyBytes, pubkey0, signature0, pubkey1, signature1};
}

const multisigBind = async () => {
  const {pubkey0, signature0, pubkey1, signature1, bodyBytes} = await signatureBind();
  const multisigPubkey = await createMultisigThresholdPubkey(
    [pubkey0, pubkey1],
    2,
  );
  const multisigAddress = pubkeyToAddress(multisigPubkey, "6x");
  const address0 = pubkeyToAddress(pubkey0, "6x");
  const address1 = pubkeyToAddress(pubkey1, "6x");
  
  const wallet = await Secp256k1HdWallet.fromMnemonic(
    alice_mnemonic,
    { prefix: "6x" }, // Replace with your own Bech32 address prefix
  );


   const broadcaster = await SigningStargateClient.connectWithSigner(
    rpc,
    wallet,
    {
      registry: customRegistry,
      aminoTypes: new AminoTypes(createBindingAdminoTypes())
    }
  )

  // const broadcaster = await SigningCosmWasmClient.connect(rpc);
  // console.log("####################",broadcaster,"########################");
  const signedTx = await makeMultisignedTx(
    multisigPubkey,
    (await signingBindInstruction).sequence,
    (await signingBindInstruction).fee,
    bodyBytes,
    new Map<string, Uint8Array>([
      [address0, signature0],
      [address1, signature1],
    ]),
  );
  
  // ensure signature is valid
  const result = await broadcaster.broadcastTx(Uint8Array.from(TxRaw.encode(signedTx).finish()));
  console.log(result);
}

// multisigSend();
multisigBind();