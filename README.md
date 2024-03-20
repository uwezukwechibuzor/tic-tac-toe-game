# TIC-CHAIN
**Ticchain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

A TIC TAC TOE Game implementation on the blockchain


# Developing

For setting up your environment to develop `tic-chain`. Shows how to build, run, send transactions and querying

# TODO - Docker

### [docker](https://www.docker.com/)

Docker is used to help make release and static builds locally.


- [Golang](https://go.dev/dl/): you can download it from the linked page or:
  - Linux: Use your distribution's package manager.
  - Mac: Use `macports` or `brew`.
- Ensure that `$GOPATH` and `$PATH` have been set properly. On a Mac that uses the Z shell, you may have to run the following:

```zsh
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.zprofile
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.zprofile
echo "export GO111MODULE=on" >> ~/.zprofile
source ~/.zprofile
```

### [make](https://www.gnu.org/software/make/)

We use GNU Make to help us built, lint, fmt, and etc for this project.

- Linux:
  - Your distro likely already comes with this. You can check by running `which make`.
  - Otherwise please see your distro specific installation tool(i.e `apt`) and use that to install it.
- Macos:
  - You can check if it's already installed by `which make`.
  - Otherwise use [brew](https://brew.sh/) or [macports](https://www.macports.org/) to install it.


### [Protobuf](https://protobuf.dev/)

A necessary tool for generating protobuf code.

- Linux:
  - Please see your distro specific installation tool(i.e `apt`) and use that to install it.
- Macos:
  - Using [brew](https://brew.sh/): `brew install protobuf`
  - Using [macports](https://www.macports.org/): `sudo port install protobuf-cpp`


## Building using Make

To build the protobuf(only necessary if you change the protobuf) you will need to run,:

```bash
make proto-gen
```

To build, run:

```bash
make build
```

To install (builds and moves the executable to `$GOPATH/bin`, which should be in `$PATH`), run:

```bash
make install
```

## Running a Single-node Local Testnet

To run a single-node testnet locally:

```bash
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
```
 
OR Simply Run

```
chmod +x start.sh
./start.sh
```

## Linting & Formatting

To lint and format the protobuf(only necessary if you mess with protobuf):

```bash
make proto-format
```

Run format and run linter for go sided:

```bash
make lint
```

## Running

After running the `make install` command you should be able to use `tic-chaind --help`.

## Testing

To run all unit tests:

```bash
make test
```

To see test coverage:

```bash
make test-cover
```

Below is a flowchart representing the flow of logic in tic-chain:

<img width="1309" alt="Screenshot 2024-02-26 at 21 53 38" src="https://github.com/InterviewingTask/recruitment-blockchain-developer-uwezukwechibuzor/assets/66339097/e6cff9c6-1af6-48d9-9b56-2a9b7d0d6ee4">


# TX COMMANDS

# Query COMMANDS

