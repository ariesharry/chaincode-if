package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Processor represents the structure for a processor
type Processor struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	NIB      string  `json:"nib"`
	NIK      string  `json:"nik"`
	NoHP     string  `json:"noHP"`
	Email    string  `json:"email"`
	Address  string  `json:"address"`
	Capacity float64 `json:"capacity"`
}

// AddProcessor adds a new processor to the ledger
func (pc *PalmOilContract) AddProcessor(ctx contractapi.TransactionContextInterface, id string, name string, nib string, nik string, noHP string, email string, address string, capacity float64) error {
	// Check if a processor with the given ID already exists
	existingProcessorJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existingProcessorJSON != nil {
		return fmt.Errorf("a processor with ID %s already exists", id)
	}

	processor := Processor{
		ID:       id,
		Name:     name,
		NIB:      nib,
		NIK:      nik,
		NoHP:     noHP,
		Email:    email,
		Address:  address,
		Capacity: capacity,
	}

	processorJSON, err := json.Marshal(processor)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, processorJSON)
}

// UpdateProcessor updates an existing processor on the ledger
func (pc *PalmOilContract) UpdateProcessor(ctx contractapi.TransactionContextInterface, id string, name string, nib string, nik string, noHP string, email string, address string, capacity float64) error {
	processorJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if processorJSON == nil {
		return fmt.Errorf("the processor with ID %s does not exist", id)
	}

	var processor Processor
	json.Unmarshal(processorJSON, &processor)

	// Update the processor's attributes
	processor.Name = name
	processor.NIB = nib
	processor.NIK = nik
	processor.NoHP = noHP
	processor.Email = email
	processor.Address = address
	processor.Capacity = capacity

	processorJSON, err = json.Marshal(processor)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, processorJSON)
}

// QueryProcessorByID retrieves a processor by its ID from the ledger
func (pc *PalmOilContract) QueryProcessorByID(ctx contractapi.TransactionContextInterface, id string) (*Processor, error) {
	processorJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if processorJSON == nil {
		return nil, fmt.Errorf("the processor with ID %s does not exist", id)
	}

	var processor Processor
	json.Unmarshal(processorJSON, &processor)

	return &processor, nil
}

// QueryAllProcessors retrieves all processors from the ledger
func (pc *PalmOilContract) QueryAllProcessors(ctx contractapi.TransactionContextInterface) ([]*Processor, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("PRO_", "PRO_zzzzzzzzzz")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var processors []*Processor
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var processor Processor
		json.Unmarshal(queryResponse.Value, &processor)
		processors = append(processors, &processor)
	}

	return processors, nil
}
