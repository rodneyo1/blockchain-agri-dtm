package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "math/big"
    "net/http"
    "os"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/gorilla/mux"
)

const contractABI = `[{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getTransaction","outputs":[{"internalType":"bytes32","name":"previousHash","type":"bytes32"},{"internalType":"uint256","name":"timestamp","type":"uint256"},{"internalType":"uint256","name":"nonce","type":"uint256"},{"internalType":"string","name":"farmerName","type":"string"},{"internalType":"string","name":"farmerPhoneNumber","type":"string"},{"internalType":"string","name":"customerName","type":"string"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"string","name":"customerPhoneNumber","type":"string"}],"stateMutability":"view","type":"function"}]`

func handlePayment(w http.ResponseWriter, r *http.Request) {
    client, err := ethclient.Dial(os.Getenv("SEPOLIA_URL"))
    if err != nil {
        http.Error(w, "Failed to connect to Ethereum client", http.StatusInternalServerError)
        return
    }

    privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    if err != nil {
        http.Error(w, "Failed to load private key", http.StatusInternalServerError)
        return
    }

    contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    parsedABI, err := abi.JSON(strings.NewReader(contractABI))
    if err != nil {
        http.Error(w, "Failed to parse contract ABI", http.StatusInternalServerError)
        return
    }

    // Example: Handle payment request
    // You need to extract payment details from the request (amount, etc.)
    index := big.NewInt(0) // This should be dynamic based on request

    data, err := parsedABI.Pack("getTransaction", index)
    if err != nil {
        http.Error(w, "Failed to pack transaction data", http.StatusInternalServerError)
        return
    }

    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        http.Error(w, "Failed to get nonce", http.StatusInternalServerError)
        return
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        http.Error(w, "Failed to suggest gas price", http.StatusInternalServerError)
        return
    }

    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(300000), gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        http.Error(w, "Failed to get network ID", http.StatusInternalServerError)
        return
    }
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        http.Error(w, "Failed to sign transaction", http.StatusInternalServerError)
        return
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        http.Error(w, "Failed to send transaction", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Transaction sent: %s", signedTx.Hash().Hex())
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/payment", handlePayment).Methods("POST")

    http.ListenAndServe(":8080", r)
}
