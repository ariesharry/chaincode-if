package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Commodity represents the structure for a commodity
type Traceability struct {
	ID       string   `json:"id"`
	Status   []string `json:"status"`
	Location []string `json:"location"`
	PIC      []string `json:"pic"`
}

type Commodity struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Quantity     float64      `json:"quantity"`
	DateHarvested string      `json:"dateHarvested"`
	Traceability Traceability `json:"traceability"`
}

type ProcessedCommodity struct {
	ID         string   `json:"id"`
	Quantity   float64  `json:"quantity"`
	Material   []string `json:"material"`
	BatchNumber string  `json:"batchNumber"`
	Quality    string   `json:"quality"`
}

func (pc *PalmOilContract) Harvest(ctx contractapi.TransactionContextInterface, commodityID string, name string, quantity float64, dateHarvested string, traceabilityID string, pic string, location string) error {
	traceability := Traceability{
		ID: traceabilityID,
		Status: []string{"harvested"},
		Location: []string{location},
		PIC: []string{pic},
	}

	commodity := Commodity{
		ID:           commodityID,
		Name:         name,
		Quantity:     quantity,
		DateHarvested: dateHarvested,
		Traceability: traceability,
	}

	commodityJSON, err := json.Marshal(commodity)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(commodityID, commodityJSON)
}


func (pc *PalmOilContract) Collect(ctx contractapi.TransactionContextInterface, commodityID string, pic string, location string) error {
	// Fetch the commodity data from the ledger
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	err = json.Unmarshal(commodityJSON, &commodity)
	if err != nil {
		return fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
	}

	// Update the traceability's status, location, and PIC
	commodity.Traceability.Status = append(commodity.Traceability.Status, "collected")
	commodity.Traceability.PIC = append(commodity.Traceability.PIC, pic)
	commodity.Traceability.Location = append(commodity.Traceability.Location, location)

	// Marshal the updated commodity back to JSON
	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return fmt.Errorf("failed to marshal commodity: %v", err)
	}

	// Update the commodity in the ledger
	return ctx.GetStub().PutState(commodityID, commodityJSON)
}


func (pc *PalmOilContract) Transport(ctx contractapi.TransactionContextInterface, commodityID string, pic string, location string) error {
	// Fetch the commodity data from the ledger
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	err = json.Unmarshal(commodityJSON, &commodity)
	if err != nil {
		return fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
	}

	// Update the traceability's status, location, and PIC
	commodity.Traceability.Status = append(commodity.Traceability.Status, "in transport")
	commodity.Traceability.PIC = append(commodity.Traceability.PIC, pic)
	commodity.Traceability.Location = append(commodity.Traceability.Location, location)

	// Marshal the updated commodity back to JSON
	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return fmt.Errorf("failed to marshal commodity: %v", err)
	}

	// Update the commodity in the ledger
	return ctx.GetStub().PutState(commodityID, commodityJSON)
}


func (pc *PalmOilContract) Transported(ctx contractapi.TransactionContextInterface, commodityID string, location string, pic string) error {
	// Fetch the commodity data from the ledger
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	err = json.Unmarshal(commodityJSON, &commodity)
	if err != nil {
		return fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
	}

	// Update the traceability's status, location, and PIC
	commodity.Traceability.Status = append(commodity.Traceability.Status, "delivered")
	commodity.Traceability.PIC = append(commodity.Traceability.PIC, pic)
	commodity.Traceability.Location = append(commodity.Traceability.Location, location)

	// Marshal the updated commodity back to JSON
	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return fmt.Errorf("failed to marshal commodity: %v", err)
	}

	// Update the commodity in the ledger
	return ctx.GetStub().PutState(commodityID, commodityJSON)
}


func (pc *PalmOilContract) Process(ctx contractapi.TransactionContextInterface, commodityID string, processedID string, quantity float64, material []string, batchNumber string, quality string, pic string, location string) error {
	// Fetch the commodity data from the ledger
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	err = json.Unmarshal(commodityJSON, &commodity)
	if err != nil {
		return fmt.Errorf("failed to unmarshal commodity JSON: %v", err)
	}

	// Update the traceability's status, location, and PIC for the commodity
	commodity.Traceability.Status = append(commodity.Traceability.Status, "processed")
	commodity.Traceability.PIC = append(commodity.Traceability.PIC, pic)
	commodity.Traceability.Location = append(commodity.Traceability.Location, location)

	// Marshal the updated commodity back to JSON
	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return fmt.Errorf("failed to marshal commodity: %v", err)
	}
	// Update the commodity in the ledger
	err = ctx.GetStub().PutState(commodityID, commodityJSON)
	if err != nil {
		return fmt.Errorf("failed to update commodity in ledger: %v", err)
	}

	// Create a new processed commodity
	processedCommodity := ProcessedCommodity{
		ID:         processedID,
		Quantity:   quantity,
		Material:   material,
		BatchNumber: batchNumber,
		Quality:    quality,
	}

	processedJSON, err := json.Marshal(processedCommodity)
	if err != nil {
		return fmt.Errorf("failed to marshal processed commodity: %v", err)
	}

	// Add the processed commodity to the ledger
	return ctx.GetStub().PutState(processedID, processedJSON)
}


// QueryCommodityAndTraceabilityByID retrieves a commodity and its traceability by commodity ID from the ledger
func (pc *PalmOilContract) QueryCommodityByID(ctx contractapi.TransactionContextInterface, commodityID string) (*Commodity, *Traceability, error) {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return nil, nil, fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	traceabilityJSON, err := ctx.GetStub().GetState(commodity.Traceability.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read traceability from world state: %v", err)
	}
	if traceabilityJSON == nil {
		return nil, nil, fmt.Errorf("the traceability with ID %s does not exist", commodity.Traceability.ID)
	}

	var traceability Traceability
	json.Unmarshal(traceabilityJSON, &traceability)

	return &commodity, &traceability, nil
}

// QueryAllCommoditiesAndTraceabilities retrieves all commodities and their traceabilities from the ledger
func (pc *PalmOilContract) QueryAllCommodities(ctx contractapi.TransactionContextInterface) ([]*Commodity, []*Traceability, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, nil, err
	}
	defer resultsIterator.Close()

	var commodities []*Commodity
	var traceabilities []*Traceability
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, nil, err
		}

		var commodity Commodity
		json.Unmarshal(queryResponse.Value, &commodity)
		commodities = append(commodities, &commodity)

		traceabilityJSON, err := ctx.GetStub().GetState(commodity.Traceability.ID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read traceability from world state: %v", err)
		}
		if traceabilityJSON != nil {
			var traceability Traceability
			json.Unmarshal(traceabilityJSON, &traceability)
			traceabilities = append(traceabilities, &traceability)
		}
	}

	return commodities, traceabilities, nil
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
