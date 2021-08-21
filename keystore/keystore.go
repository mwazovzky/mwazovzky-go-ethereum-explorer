package keystore

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// CreateKeyStore creates new wallet and stores it's data to a file
func CreateKeyStore(password string, path string) string {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	return account.Address.Hex()
}

// ImportKeyStore imports account information from a file and creates a new keystore
//
// @param file - file to read account info from
// @param path - path to store keystore to
// @param password - password to decode existing keystore
// @param passwordNew - password to encode new keystore
func ImportKeyStore(file string, path string, password string, passwordNew string) accounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)

	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, password, passwordNew)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}

	return account
}
