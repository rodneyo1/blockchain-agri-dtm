package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MpesaResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}

type SmartContract interface {
	ConfirmPayment(opts *bind.TransactOpts, transactionID string) (*big.Int, error)
}

func handleMpesaPayment(w http.ResponseWriter, r *http.Request) {
	// Parse MPesa payment request
	var mpesaResponse MpesaResponse
	err := json.NewDecoder(r.Body).Decode(&mpesaResponse)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Connect to Ethereum blockchain
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	// Load smart contract
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	// Replace with your smart contract address and ABI
	address := "SMART_CONTRACT_ADDRESS"
	instance, err := NewSmartContract(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal(err)
	}

	// Call smart contract to confirm payment
	tx, err := instance.ConfirmPayment(auth, mpesaResponse.TransactionID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Payment confirmed: %s", tx.Hash().Hex())
}

// func main() {
// 	http.HandleFunc("/mpesa-payment", handleMpesaPayment)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
