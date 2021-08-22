package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/mwazovzky/ethereum/explorer"
)

var address string
var url string

func init() {
	godotenv.Load()
	address = os.Getenv("ADDRESS")
	url = os.Getenv("INFURA_URL")
}

func main() {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// blockNumber := header.Number

	ex := explorer.New(client)

	// balance, err := ex.GetBalance(address)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balance)

	transactions := ex.GetTransactions(address)

	// blockNumber := new(big.Int).SetInt64(10881303)
	// transactions := ex.GetBlockTransactions(blockNumber)

	for _, tx := range transactions {
		if tx.To() != nil && tx.To().Hex() == address {
			ex.ShowTransaction(tx)
		}
	}
}
