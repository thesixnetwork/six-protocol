# Tokenmngr module

### __command list__

__Query__

```devnet-wasmd query tokenmngr```

```bash
Available Commands:
  list-mintperm list all mintperm
  list-token    list all token
  params        shows the parameters of the module
  show-mintperm shows a mintperm
  show-token    shows a token
  show-options  shows tokenmngr module options
```

list-mintperm

```bash
devnet-wasmd query tokenmngr list-mintperm
```

list-token

```bash
devnet-wasmd query tokenmngr list-token
```

show-mintperm

```bash
devnet-wasmd query tokenmngr show-mintperm [token] [address]
```

show-token

```bash
devnet-wasmd query tokenmngr show-token [name]
```

show-options

```bash
devnet-wasmd query tokenmngr show-options
```

__Tx__

```bash
Available Commands:
  create-mintperm Create a new mintperm
  create-token    Create a new token
  delete-mintperm Delete a mintperm
  delete-token    Delete a token
  mint            Broadcast message mint
  update-mintperm Update a mintperm
  update-token    Update a token
```

create-mintperm

msg sender have to be token admin

```bash
devnet-wasmd tx tokenmngr create-mintperm [token] [address]
```

```ts
// msg for cosmjs
const msgCreateMintperm: {
  creator: string
  token: string
  address: string
} = {
  creator: "6xsomething",
  token: "six",
  address: "6xminter"
};

const msg = {
  typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateMintperm",
  value: msgCreateMintperm
};
```

create-token

msg sender have to be token admin

```bash
devnet-wasmd tx tokenmngr create-token [name] [max-supply] [denom-metadata]
```

```ts
// msg for cosmjs
const denomMetaData = {
    description: "The native staking token of the SIX Protocol.",
    denom_units: [
      {
        denom: "usix",
        exponent: 12,
        aliases: [
          "microsix"
        ]
      },
      {
        denom: "six",
        exponent: 18
      }
    ],
    base: "usix",
    name: "SIX Token",
    symbol: "six",
    display: "six"
}

const msgCreateToken: {
  creator: string
  name: string
  maxSupply: string
  mintee: string
  denomMetaData: string
} = {
  creator: "6xsomething",
  name: "six",
  maxSupply: "10000000000000",
  mintee: "6xmintee",
  denomMetaData: JSON.stringify(denomMetaData)
};

const msg = {
  typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateToken",
  value: msgCreateToken
};
```

mint

msg sender have to have mint permission

```bash
devnet-wasmd tx tokenmngr mint [amount] [token] [receiver] [flags]
```

```ts
// msg for cosmjs
const msgMint: {
  creator: string
  amount: string
  token: string
} = {
  creator: "6xsomething",
  amount: "666666",
  token: "six"
};

const msg = {
  typeUrl: "/thesixnetwork.sixprotocol.tokenmngr.MsgMint",
  value: msgMint
};
```
