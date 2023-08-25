package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PalmOilContract struct {
	contractapi.Contract
}

func NewPalmOilContract() contractapi.ContractInterface {
	return &PalmOilContract{
		Contract: contractapi.Contract{},
	}
}
