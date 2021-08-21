package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var password string
var path string
var file string
var address string
var url string

func init() {
	godotenv.Load()
	password = os.Getenv("PASSWORD")
	path = os.Getenv("KEYSTORE_PATH")
	address = os.Getenv("ADDRESS")
	url = os.Getenv("INFURA_URL")
}

func getBalance(client *ethclient.Client, address string) (*big.Int, error) {
	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func toFloat64(value *big.Int) float64 {
	numerator := new(big.Float).SetInt64(value.Int64())
	denominator := new(big.Float).SetInt64(1e+18)
	res, _ := new(big.Float).Quo(numerator, denominator).Float64()

	return res
}

func main() {
	// address := keystore.CreateKeyStore(password, "./tmp")
	// fmt.Println(address)

	// file = "./tmp/UTC--2021-08-21T13-48-59.942696000Z--663ed5184aa98203da9b7b7418d57e43efe406d3"
	// account := keystore.ImportKeyStore(file, path, password, password)
	// fmt.Println(account.Address.Hex())

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := getBalance(client, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%f\n", toFloat64(balance))
}
