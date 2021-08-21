package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mwazovzky/ethereum/keystore"
)

var password string
var path string

func init() {
	godotenv.Load()
	password = os.Getenv("PASSWORD")
	path = os.Getenv("KEYSTORE_PATH")
}

func main() {
	address := keystore.CreateKeyStore(password, "./tmp")
	fmt.Println(address)

	file := "./tmp/UTC--2021-08-22T09-24-52.009129000Z--0dd34627c97c3b539455d50f84c4710fb24b1fd7"
	account := keystore.ImportKeyStore(file, path, password, password)
	fmt.Println(account.Address.Hex())
}
