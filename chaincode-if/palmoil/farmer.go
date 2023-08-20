package palmoil

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Farmer represents the structure for a farmer
type Farmer struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	NIK     string   `json:"nik"`
	Address string   `json:"address"`
	Email   string   `json:"email"`
	NoHP    string   `json:"noHP"`
	Farm    []string `json:"farm"`
}

// AddFarmer adds a new farmer to the ledger
func (pc *PalmOilContract) AddFarmer(ctx contractapi.TransactionContextInterface, id string, name string, nik string, address string, email string, noHP string, farmsInput string) error {
	// Check if a farmer with the given ID already exists
	existingFarmerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existingFarmerJSON != nil {
		return fmt.Errorf("a farmer with ID %s already exists", id)
	}

	// Parse the farmsInput into a []string
	var farms []string
	err = json.Unmarshal([]byte(farmsInput), &farms)
	if err != nil {
		return fmt.Errorf("failed to parse farm attribute: %v", err)
	}

	farmer := Farmer{
		ID:      id,
		Name:    name,
		NIK:     nik,
		Address: address,
		Email:   email,
		NoHP:    noHP,
		Farm:    farms,
	}

	farmerJSON, err := json.Marshal(farmer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmerJSON)
}

// UpdateFarmer updates an existing farmer on the ledger
func (pc *PalmOilContract) UpdateFarmer(ctx contractapi.TransactionContextInterface, id string, name string, nik string, address string, email string, noHP string, farmsInput string) error {
	farmerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if farmerJSON == nil {
		return fmt.Errorf("the farmer with ID %s does not exist", id)
	}

	// Parse the farmsInput into a []string
	var farms []string
	err = json.Unmarshal([]byte(farmsInput), &farms)
	if err != nil {
		return fmt.Errorf("failed to parse farm attribute: %v", err)
	}

	var farmer Farmer
	json.Unmarshal(farmerJSON, &farmer)

	// Update the farmer's attributes
	farmer.Name = name
	farmer.NIK = nik
	farmer.Address = address
	farmer.Email = email
	farmer.NoHP = noHP
	farmer.Farm = farms

	farmerJSON, err = json.Marshal(farmer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmerJSON)
}

// QueryFarmerByID retrieves a farmer by its ID from the ledger
func (pc *PalmOilContract) QueryFarmerByID(ctx contractapi.TransactionContextInterface, id string) (*Farmer, error) {
	farmerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if farmerJSON == nil {
		return nil, fmt.Errorf("the farmer with ID %s does not exist", id)
	}

	var farmer Farmer
	json.Unmarshal(farmerJSON, &farmer)

	return &farmer, nil
}

// QueryAllFarmers retrieves all farmers from the ledger
func (pc *PalmOilContract) QueryAllFarmers(ctx contractapi.TransactionContextInterface) ([]*Farmer, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var farmers []*Farmer
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var farmer Farmer
		json.Unmarshal(queryResponse.Value, &farmer)
		farmers = append(farmers, &farmer)
	}

	return farmers, nil
}


