package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// QueryCommodityByID retrieves a commodity by its ID from the ledger
func (pc *PalmOilContract) QueryCommodityByID(ctx contractapi.TransactionContextInterface, id string) (*Commodity, error) {
	commodityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return nil, fmt.Errorf("the commodity with ID %s does not exist", id)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	return &commodity, nil
}

// QueryAllCommodities retrieves all commodities from the ledger
func (pc *PalmOilContract) QueryAllCommodities(ctx contractapi.TransactionContextInterface) ([]*Commodity, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var commodities []*Commodity
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var commodity Commodity
		json.Unmarshal(queryResponse.Value, &commodity)
		commodities = append(commodities, &commodity)
	}

	return commodities, nil
}

// QueryTraceabilityByID retrieves a traceability by its ID from the ledger
func (pc *PalmOilContract) QueryTraceabilityByID(ctx contractapi.TransactionContextInterface, id string) (*Traceability, error) {
	traceabilityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if traceabilityJSON == nil {
		return nil, fmt.Errorf("the traceability with ID %s does not exist", id)
	}

	var traceability Traceability
	json.Unmarshal(traceabilityJSON, &traceability)

	return &traceability, nil
}

// QueryAllTraceabilities retrieves all traceabilities from the ledger
func (pc *PalmOilContract) QueryAllTraceabilities(ctx contractapi.TransactionContextInterface) ([]*Traceability, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var traceabilities []*Traceability
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var traceability Traceability
		json.Unmarshal(queryResponse.Value, &traceability)
		traceabilities = append(traceabilities, &traceability)
	}

	return traceabilities, nil
}

// QueryProcessedCommodityByID retrieves a processed commodity by its ID from the ledger
func (pc *PalmOilContract) QueryProcessedCommodityByID(ctx contractapi.TransactionContextInterface, id string) (*ProcessedCommodity, error) {
	processedCommodityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if processedCommodityJSON == nil {
		return nil, fmt.Errorf("the processed commodity with ID %s does not exist", id)
	}

	var processedCommodity ProcessedCommodity
	json.Unmarshal(processedCommodityJSON, &processedCommodity)

	return &processedCommodity, nil
}

// QueryAllProcessedCommodities retrieves all processed commodities from the ledger
func (pc *PalmOilContract) QueryAllProcessedCommodities(ctx contractapi.TransactionContextInterface) ([]*ProcessedCommodity, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var processedCommodities []*ProcessedCommodity
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var processedCommodity ProcessedCommodity
		json.Unmarshal(queryResponse.Value, &processedCommodity)
		processedCommodities = append(processedCommodities, &processedCommodity)
	}

	return processedCommodities, nil
}
