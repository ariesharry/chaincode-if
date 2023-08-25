package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Transporter represents the structure for a transporter
type Transporter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	NIK     string `json:"nik"`
	NoHP    string `json:"noHP"`
	NumShip int    `json:"numShip"`
}

// AddTransporter adds a new transporter to the ledger
func (pc *PalmOilContract) AddTransporter(ctx contractapi.TransactionContextInterface, id string, name string, nik string, noHP string, numShip int) error {
	// Check if a transporter with the given ID already exists
	existingTransporterJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existingTransporterJSON != nil {
		return fmt.Errorf("a transporter with ID %s already exists", id)
	}

	transporter := Transporter{
		ID:      id,
		Name:    name,
		NIK:     nik,
		NoHP:    noHP,
		NumShip: numShip,
	}

	transporterJSON, err := json.Marshal(transporter)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, transporterJSON)
}

// UpdateTransporter updates an existing transporter on the ledger
func (pc *PalmOilContract) UpdateTransporter(ctx contractapi.TransactionContextInterface, id string, name string, nik string, noHP string, numShip int) error {
	transporterJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if transporterJSON == nil {
		return fmt.Errorf("the transporter with ID %s does not exist", id)
	}

	var transporter Transporter
	json.Unmarshal(transporterJSON, &transporter)

	// Update the transporter's attributes
	transporter.Name = name
	transporter.NIK = nik
	transporter.NoHP = noHP
	transporter.NumShip = numShip

	transporterJSON, err = json.Marshal(transporter)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, transporterJSON)
}

// QueryTransporterByID retrieves a transporter by its ID from the ledger
func (pc *PalmOilContract) QueryTransporterByID(ctx contractapi.TransactionContextInterface, id string) (*Transporter, error) {
	transporterJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if transporterJSON == nil {
		return nil, fmt.Errorf("the transporter with ID %s does not exist", id)
	}

	var transporter Transporter
	json.Unmarshal(transporterJSON, &transporter)

	return &transporter, nil
}

// QueryAllTransporters retrieves all transporters from the ledger
func (pc *PalmOilContract) QueryAllTransporters(ctx contractapi.TransactionContextInterface) ([]*Transporter, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("TRA_", "TRA_zzzzzzzzzz")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transporters []*Transporter
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transporter Transporter
		json.Unmarshal(queryResponse.Value, &transporter)
		transporters = append(transporters, &transporter)
	}

	return transporters, nil
}

