# Deploying and Interacting with the Smart Contract using Golang

This repo has been created to demonstrate how to deploy and interact with a smart contract using Golang. Read more about here: [Go Ethereum Book](https://goethereumbook.org/smart-contracts/)

## Setup

- Add in your `.env` file the following variables:

```bash
GOERLI_ALCHEMY=
PRIVATE_KEY=
CONTRACT_ADDRESS=
```

- Install the dependencies

```bash
go mod download
```

### Generating the Smart Contract ABI

```bash
solc --abi contract/Store.sol -o out
```

### Generating the Smart Contract Bytecode

```bash
solc --bin contract/Store.sol -o out
```

### Converting ABI to Go 

Now let's convert the ABI to a Go file that we can import. This new file will contain all the available methods the we can use to interact with the smart contract from our Go application.

```bash
abigen --bin=out/Store.bin --abi=out/Store.abi --pkg=store --out=api/Store.go
```

### Deploying the Smart Contract

```bash
go run cmd/deploy/main.go
```

### Interacting with the Smart Contract

- Reading data

```bash
go run cmd/read/main.go
```

- Writing data

```bash
go run cmd/write/main.go
```

### Getting the Bytecode

```bash
go run cmd/bytecode/main.go
``` 