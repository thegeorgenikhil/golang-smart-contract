package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	store "github.com/thegeorgenikhil/golang-smart-contract/api" // for demo
)

func main() {
	godotenv.Load()
	client, err := ethclient.Dial(os.Getenv("GOERLI_ALCHEMY"))

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	// read the value of the key "foo"
	key := [32]byte{}
	copy(key[:], []byte("foo"))
	value, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Foo:", string(value[:])) // "bar"
	fmt.Println("Version:", version)      // "1.0"
}
