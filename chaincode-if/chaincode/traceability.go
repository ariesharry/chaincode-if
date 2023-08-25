package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Commodity represents the structure for a commodity
type Commodity struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Quantity     float64      `json:"quantity"`
	DateHarvested string      `json:"dateHarvested"`
	Traceability Traceability `json:"traceability"`
}

// Traceability represents the traceability structure for a commodity
type Traceability struct {
	ID      string   `json:"id"`
	Status  []string `json:"status"`
	Location []string `json:"location"`
	PIC     []string `json:"pic"`
}

// ProcessedCommodity represents the structure for a processed commodity
type ProcessedCommodity struct {
	ID         string    `json:"id"`
	Quantity   float64   `json:"quantity"`
	Material   []string  `json:"material"`
	BatchNumber string   `json:"batchNumber"`
	Quality    string    `json:"quality"`
}

func (pc *PalmOilContract) Harvest(ctx contractapi.TransactionContextInterface, commodityID string, name string, quantity float64, dateHarvested string, traceabilityID string) error {
	traceability := Traceability{
		ID: traceabilityID,
		Status: []string{"harvested"},
		Location: []string{},
		PIC: []string{},
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

func (pc *PalmOilContract) Collect(ctx contractapi.TransactionContextInterface, commodityID string) error {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	// Update the traceability's status
	commodity.Traceability.Status = append(commodity.Traceability.Status, "collected")

	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(commodityID, commodityJSON)
}

func (pc *PalmOilContract) Transport(ctx contractapi.TransactionContextInterface, commodityID string) error {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	// Update the traceability's status
	commodity.Traceability.Status = append(commodity.Traceability.Status, "in transport")

	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(commodityID, commodityJSON)
}

func (pc *PalmOilContract) UpdateTransport(ctx contractapi.TransactionContextInterface, commodityID string, location string, hasArrived bool) error {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	// Update the traceability's location
	commodity.Traceability.Location = append(commodity.Traceability.Location, location)

	if hasArrived {
		commodity.Traceability.Status = append(commodity.Traceability.Status, "delivered")
	}

	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(commodityID, commodityJSON)
}

func (pc *PalmOilContract) Process(ctx contractapi.TransactionContextInterface, commodityID string, processedID string, quantity float64, material []string, batchNumber string, quality string) error {
	commodityJSON, err := ctx.GetStub().GetState(commodityID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if commodityJSON == nil {
		return fmt.Errorf("the commodity with ID %s does not exist", commodityID)
	}

	var commodity Commodity
	json.Unmarshal(commodityJSON, &commodity)

	// Update the traceability's status
	commodity.Traceability.Status = append(commodity.Traceability.Status, "processed")

	commodityJSON, err = json.Marshal(commodity)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(commodityID, commodityJSON)

	processedCommodity := ProcessedCommodity{
		ID:         processedID,
		Quantity:   quantity,
		Material:   material,
		BatchNumber: batchNumber,
		Quality:    quality,
	}

	processedJSON, err := json.Marshal(processedCommodity)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(processedID, processedJSON)
}
