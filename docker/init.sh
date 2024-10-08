MONIKER=$1
if [ -z "$MONIKER" ]; then
  MONIKER="mynode"
fi
export CHAIN_ID=testnet
export VALKEY=val1 # should be: export as docker env var
export SIX_HOME=./build/six_home
ALICE_MNEMONIC="history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
BOB_MNEMONIC="limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture"
VAL1_MNEMONIC="note base stone list envelope tail start forget alarm acoustic cook occur divert giant bike curtain chase shuffle fade glow capital slot file provide"
VAL2_MNEMONIC="strike tower consider despair bridge diesel clay celery violin base hello ride they weather tunnel elite truth oblige spot hen wise flag pet battle"
VAL3_MNEMONIC="canvas human require month loan oak december blame grit palm slice error absorb total spice autumn trouble soda repeat shove quit bid forward organ"
VAL4_MNEMONIC="grant raw marine drink text dove flat waste wish buzz output hand merge cluster civil clog stay alert silent reunion idea cake village almost"
ORACLE1_MNEMONIC="list split future remain scene cheap pledge forum siren purse bright ivory split morning swing dumb fabric rapid remove worth diary task island donkey"
ORACLE2_MNEMONIC="achieve rice anger junk delay glove slam find poem feed emerge next core twice kitchen road proof remain notice slice walk super piece father"
ORACLE3_MNEMONIC="hint expose mix lemon leave genuine host fiction peasant daughter enable region mixture bean soda auction armed turtle iron become bracket wasp drama front"
ORACLE4_MNEMONIC="clown cabbage clean design mosquito surround citizen virus kite castle sponsor wife lesson coffee alien panel hand together good crazy fabric mouse hat town"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

# mint to validator
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ALICE_MNEMONIC | sixd keys add alice --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $BOB_MNEMONIC | sixd keys add bob --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL1_MNEMONIC | sixd keys add val1 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL2_MNEMONIC | sixd keys add val2 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL3_MNEMONIC | sixd keys add val3 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL4_MNEMONIC | sixd keys add val4 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE1_MNEMONIC | sixd keys add oracle1 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE2_MNEMONIC | sixd keys add oracle2 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE3_MNEMONIC | sixd keys add oracle3 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE4_MNEMONIC | sixd keys add oracle4 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118

sixd add-genesis-account $(sixd keys show -a val1 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val2 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val3 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val4 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle1 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle2 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle3 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle4 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a alice --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a bob --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a super-admin --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}

# eth_secp256k1 address
echo $ALICE_MNEMONIC | sixd keys add alice_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1 --coin-type 60
echo $BOB_MNEMONIC | sixd keys add bob_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1 --coin-type 60

sixd gentx ${VALKEY} 100000000usix --chain-id=${CHAIN_ID} --keyring-backend=test --home ${SIX_HOME}
sixd collect-gentxs --home ${SIX_HOME}