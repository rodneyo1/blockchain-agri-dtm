package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// Replace with your contract ABI and address
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
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getTransactionCount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "_previousHash",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "_timestamp",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_nonce",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_farmerName",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_farmerPhoneNumber",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "_customerName",
        "type": "string"
      },
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_customerPhoneNumber",
        "type": "string"
      }
    ],
    "name": "logTransaction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "transactionCount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "transactions",
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
    "stateMutability": "view",
    "type": "function"
  }]`
const contractAddress = "0xcC4072ed9C652fF91b7581A4f8EB58b78f4b698D"

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the Infura URL and private key
	infuraURL := os.Getenv("INFURA_URL")
	privateKey := os.Getenv("PRIVATE_KEY")

	// Connect to Ethereum network
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Load contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// Create a new call message
	contractAddr := common.HexToAddress(contractAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: parsedABI.Methods["myFunction"].ID,
	}

	// Call the contract function
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract function: %v", err)
	}

	// Decode the result
	var output *big.Int
	err = parsedABI.UnpackIntoInterface(&output, "myFunction", result)
	if err != nil {
		log.Fatalf("Failed to unpack contract result: %v", err)
	}

	fmt.Printf("Contract result: %s\n", output.String())
}
