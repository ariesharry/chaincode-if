package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ... [Your Other Struct Definitions Here] ...

type CommodityWithTraceability struct {
	Commodity    Commodity    `json:"commodity"`
	Traceability Traceability `json:"traceability"`
}

func (pc *PalmOilContract) QueryCommodityByID(ctx contractapi.TransactionContextInterface, commodityID string) (*CommodityWithTraceability, error) {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return nil, fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	err = json.Unmarshal(commodityJSON, &commodity)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
	}

	traceabilityJSON, err := ctx.GetStub().GetState(commodity.Traceability.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read traceability from world state: %v", err)
	}
	if traceabilityJSON == nil {
		return nil, fmt.Errorf("the traceability with ID %s does not exist", commodity.Traceability.ID)
	}

	var traceability Traceability
	err = json.Unmarshal(traceabilityJSON, &traceability)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal traceability JSON: %v", err)
	}

	return &CommodityWithTraceability{
		Commodity:    commodity,
		Traceability: traceability,
	}, nil
}

func (pc *PalmOilContract) QueryAllCommodities(ctx contractapi.TransactionContextInterface) ([]*CommodityWithTraceability, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var commoditiesWithTraceability []*CommodityWithTraceability
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var commodity Commodity
		err = json.Unmarshal(queryResponse.Value, &commodity)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
		}

		traceabilityJSON, err := ctx.GetStub().GetState(commodity.Traceability.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to read traceability from world state: %v", err)
		}
		if traceabilityJSON != nil {
			var traceability Traceability
			err = json.Unmarshal(traceabilityJSON, &traceability)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal traceability JSON: %v", err)
			}
			commoditiesWithTraceability = append(commoditiesWithTraceability, &CommodityWithTraceability{
				Commodity:    commodity,
				Traceability: traceability,
			})
		}
	}

	return commoditiesWithTraceability, nil
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