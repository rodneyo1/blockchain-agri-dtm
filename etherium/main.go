package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "math/big"
    "os"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"
)

const contractABI = `[{
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "previousHash",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "timestamp",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "nonce",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "farmerName",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "farmerPhoneNumber",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "customerName",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "customerPhoneNumber",
        "type": "string"
      }
    ],
    "name": "TransactionLogged",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "index",
        "type": "uint256"
      }
    ],
    "name": "getTransaction",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "previousHash",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "timestamp",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "nonce",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "farmerName",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "farmerPhoneNumber",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "customerName",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "customerPhoneNumber",
        "type": "string"
      }
    ],
    "name": "getTransaction",
    "type": "function"
  }]`

func encodeTransactionData(abi abi.ABI, index *big.Int) ([]byte, error) {
    return abi.Pack("getTransaction", index)
}

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
        return
    }

    client, err := ethclient.Dial(os.Getenv("SEPOLIA_URL"))
    if err != nil {
        fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
        return
    }

    privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    if err != nil {
        fmt.Printf("Failed to load private key: %v\n", err)
        return
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        fmt.Printf("Failed to cast public key to ECDSA\n")
        return
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        fmt.Printf("Failed to get nonce: %v\n", err)
        return
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        fmt.Printf("Failed to suggest gas price: %v\n", err)
        return
    }

    contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    parsedABI, err := abi.JSON(strings.NewReader(contractABI))
    if err != nil {
        fmt.Printf("Failed to parse contract ABI: %v\n", err)
        return
    }

    // Example index value for fetching the transaction
    index := big.NewInt(0) // Replace with the actual index value you want to query

    data, err := encodeTransactionData(parsedABI, index)
    if err != nil {
        fmt.Printf("Failed to pack transaction data: %v\n", err)
        return
    }

    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(300000), gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        fmt.Printf("Failed to get network ID: %v\n", err)
        return
    }
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %v\n", err)
        return
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        fmt.Printf("Failed to send transaction: %v\n", err)
        return
    }

    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

    txHash := signedTx.Hash()
    tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
    if err != nil {
        fmt.Printf("Failed to get transaction by hash: %v\n", err)
        return
    }

    fmt.Printf("Transaction details:\n")
    fmt.Printf("To: %s\n", tx.To().Hex())
    fmt.Printf("Value: %s\n", tx.Value().String())
    fmt.Printf("Gas: %d\n", tx.Gas())
    fmt.Printf("Gas Price: %s\n", tx.GasPrice().String())
    fmt.Printf("Nonce: %d\n", tx.Nonce())
    fmt.Printf("Data: %x\n", tx.Data())
    fmt.Printf("Is Pending: %v\n", isPending)

    receipt, err := client.TransactionReceipt(context.Background(), txHash)
    if err != nil {
        fmt.Printf("Failed to get transaction receipt: %v\n", err)
        return
    }

    if receipt == nil {
        fmt.Printf("Transaction receipt is nil\n")
        return
    }
    
    fmt.Printf("")
    fmt.Printf("Transaction receipt:\n")
    fmt.Printf("Status: %d\n", receipt.Status)
    fmt.Printf("Block Number: %d\n", receipt.BlockNumber.Uint64())
    fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
    fmt.Printf("Logs: %v\n", receipt.Logs)
}
