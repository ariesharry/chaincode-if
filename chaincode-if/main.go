package main

import (
	"log"
	"chaincode-if/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	rantaiPasokChaincode, err := contractapi.NewChaincode(chaincode.NewPalmOilContract())

	if err != nil {
		log.Panicf("Error creating RantaiPasokChaincode: %v", err)
	}

	if err := rantaiPasokChaincode.Start(); err != nil {
		log.Panicf("Error starting RantaiPasokChaincode: %v", err)
	}
}