package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	palmoil "github.com/ariesharry/chaincode-if/tree/main/chaincode-if/palmoil"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&palmoil.PalmOilContract{})
	if err != nil {
		fmt.Printf("Error create palm oil chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting palm oil chaincode: %s", err.Error())
	}
}