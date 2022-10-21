rm -rf ~/.six/data
rm -rf ~/.six/wasm
rm ~/.six/config/addrbook.json
mkdir ~/.six/data/
touch ~/.six/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > ~/.six/data/priv_validator_state.json