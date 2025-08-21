MONIKER=$1
if [ -z "$MONIKER" ]; then
  MONIKER="mynode"
fi
export CHAIN_ID=testnet
export KEYRING=test
export VALKEY=val1
export SIX_HOME=./build/six_home
export KEYALGO="secp256k1"

ALICE_MNEMONIC="history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
BOB_MNEMONIC="limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture"
VAL1_MNEMONIC="note base stone list envelope tail start forget alarm acoustic cook occur divert giant bike curtain chase shuffle fade glow capital slot file provide"
VAL2_MNEMONIC="strike tower consider despair bridge diesel clay celery violin base hello ride they weather tunnel elite truth oblige spot hen wise flag pet battle"
VAL3_MNEMONIC="canvas human require month loan oak december blame grit palm slice error absorb total spice autumn trouble soda repeat shove quit bid forward organ"
VAL4_MNEMONIC="grant raw marine drink text dove flat waste wish buzz output hand merge cluster civil clog stay alert silent reunion idea cake village almost"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"

rm -Rf ${SIX_HOME}

# Set client config
sixd config set client chain-id $CHAIN_ID --home ${SIX_HOME}
sixd config set client keyring-backend $KEYRING --home ${SIX_HOME}

# mint to validator
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ALICE_MNEMONIC | sixd keys add alice --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $BOB_MNEMONIC | sixd keys add bob --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL1_MNEMONIC | sixd keys add val1 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL2_MNEMONIC | sixd keys add val2 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL3_MNEMONIC | sixd keys add val3 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL4_MNEMONIC | sixd keys add val4 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}

sixd init $MONIKER --chain-id $CHAIN_ID --home ${SIX_HOME}

sixd genesis add-genesis-account $(sixd keys show -a val1 --keyring-backend=test --home ${SIX_HOME}) 11000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a val2 --keyring-backend=test --home ${SIX_HOME}) 11000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a val3 --keyring-backend=test --home ${SIX_HOME}) 11000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a val4 --keyring-backend=test --home ${SIX_HOME}) 11000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a alice --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a bob --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd genesis add-genesis-account $(sixd keys show -a super-admin --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}

sixd genesis gentx ${VALKEY} 1000000000000usix --min-self-delegation="10000000000" --validator-mode=0 --min-delegation="10000000000" --enable-redelegation=false --keyring-backend $KEYRING --chain-id $CHAIN_ID --home ${SIX_HOME}
sixd genesis collect-gentxs --home ${SIX_HOME}