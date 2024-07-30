package main


import (
	"crypto/sha256"
"encoding/json"
"fmt"
"time"
"strings"
"log"

	"github.com/oyamo/daraja-go"
)


const (
	consumerKey = "E22yMhs"
	consumerSecret = "zAFGe5cWKv3U1HQ7"
)
daraja := darajago.NewDarajaApi(consumerKey, consumerSecret, darajago.ENVIRONMENT_SANDBOX)

lnmPayload := darajago.LipaNaMpesaPayload{
	BusinessShortCode: "",
	Password:          "",
	Amount:            "",
	PartyA:            "",
	PartyB:            "",
	PhoneNumber:       "",
	CallBackURL:       "",
	AccountReference:  "",
	TransactionDesc:   "",
  }
  
  paymentResponse, err := daraja.MakeSTKPushRequest(lnmPayload)
  if err != nil {
	// Handle error
   }

   qrPayload := darajago.QRPayload{
	MerchantName:          "",
	RefNo:                 "",
	Amount:                34,
	TransactionType:       darajago.TransactionTypeBuyGoods,
	CreditPartyIdentifier: "",
  }
	  
  // Make a QR code request
  qrResponse, err := daraja.MakeQRCodeRequest(qrPayload)
  if err != nil {
	// Handle error
  }


  for _, ch := range lnmPayload {
	type Block struct {
		nonce int
		previousHash string
		timestamp int64
		transactions []string
	ch.PartyB string
	ch.PhoneNumber string
	ch.AccountReference string
	ch.TransactionDesc string
		}




  }

  func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

  }

  

