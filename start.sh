#!/bin/bash

make install
BINARY=tic-chaind

$BINARY comet unsafe-reset-all
rm -rf ~/.tic-chaind

$BINARY init node0 --default-denom ttt --chain-id test

# Add admin key
$BINARY keys add admin

# Add firstplayer key
$BINARY keys add firstplayer

# Add secondplayer key
$BINARY keys add secondplayer

# Add genesis account
$BINARY genesis add-genesis-account $($BINARY keys show admin -a) 100000000000000ttt

$BINARY genesis add-genesis-account $($BINARY keys show firstplayer  -a) 1000000ttt

$BINARY genesis add-genesis-account $($BINARY keys show secondplayer  -a) 1000000ttt

$BINARY add-genesis-game-admin $($BINARY keys show admin -a)

# Generate genesis transaction
$BINARY genesis gentx admin 1000000ttt --chain-id test

# Collect genesis transactions
$BINARY genesis collect-gentxs

# Change settings
sed -i 's/^enable = .*/enable = true/' ~/.tic-chaind/config/app.toml

# Start the node
$BINARY start --minimum-gas-prices 0.025ttt 