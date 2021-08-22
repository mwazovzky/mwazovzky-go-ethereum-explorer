package explorer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Explorer struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *Explorer {
	return &Explorer{client}
}

func (ex *Explorer) GetBlockTransactions(blockNumber *big.Int) []*types.Transaction {
	transactions := []*types.Transaction{}

	block, err := ex.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		transactions = append(transactions, tx)
	}

	return transactions
}

func (ex *Explorer) GetTransactions(address string) []*types.Transaction {
	balance, err := ex.GetBalance(address)
	if err != nil {
		log.Fatal(err)
	}

	// header, err := e.client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// blockNumber := header.Number

	blockNumber := new(big.Int).SetInt64(10881305)

	transactions := []*types.Transaction{}
	value := new(big.Int).SetInt64(0)

	// history starts when balance was 0
	for balance.Cmp(value) == 1 {
		block, err := ex.client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		for _, tx := range block.Transactions() {
			toAddress := tx.To()
			if toAddress != nil && toAddress.Hex() == address {
				transactions = append(transactions, tx)
				value = value.Add(value, tx.Value())
				fmt.Printf("%#v\n", tx.Value())
				timestamp := int64(block.Time())
				fmt.Println(blockNumber, time.Unix(timestamp, 0))
			}
		}

		blockNumber = blockNumber.Sub(blockNumber, new(big.Int).SetInt64(1))
	}

	return transactions
}

func (ex *Explorer) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)

	balance, err := ex.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (ex *Explorer) ShowTransaction(tx *types.Transaction) {
	fmt.Println("Hash:", tx.Hash().Hex())
	fmt.Println("Value:", tx.Value().String(), "wei")

	from, err := ex.getAddressFrom(tx)
	if err == nil {
		fmt.Println("From:", from.Hex())
	}

	fmt.Println("To", tx.To().Hex())
	fmt.Println("Gas:", tx.Gas())
	fmt.Println("Gas Price:", tx.GasPrice().Uint64(), "wei")
	fmt.Println("Nonce", tx.Nonce())
	fmt.Println("Data:", tx.Data())
}

func toFloat64(value *big.Int) float64 {
	numerator := new(big.Float).SetInt64(value.Int64())
	denominator := new(big.Float).SetInt64(1e+18)
	res, _ := new(big.Float).Quo(numerator, denominator).Float64()

	return res
}

func (ex *Explorer) getAddressFrom(tx *types.Transaction) (common.Address, error) {
	chainID, err := ex.client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)

	if err != nil {
		return common.Address{}, err
	}

	return msg.From(), nil
}
