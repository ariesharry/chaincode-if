// farm.go

package palmoil

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Farm represents the structure for a farm
type Farm struct {
	ID            string  `json:"id"`
	Owner         string  `json:"owner"`
	PlantedYear   int     `json:"plantedYear"`
	SeedVarieties string  `json:"seedVarieties"`
	Area          float64 `json:"area"`
	Address       string  `json:"address"`
	Coordinate    string  `json:"coordinate"`
	Capacity      float64 `json:"capacity"`
	Legality      string  `json:"legality"`
	Certificate   string  `json:"certificate"`
}

// AddFarm adds a new farm to the ledger
func (pc *PalmOilContract) AddFarm(ctx contractapi.TransactionContextInterface, id string, owner string, plantedYear int, seedVarieties string, area float64, address string, coordinate string, capacity float64, legality string, certificate string) error {
	existingFarmJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existingFarmJSON != nil {
		return fmt.Errorf("a farm with ID %s already exists", id)
	}

	farm := Farm{
		ID:            id,
		Owner:         owner,
		PlantedYear:   plantedYear,
		SeedVarieties: seedVarieties,
		Area:          area,
		Address:       address,
		Coordinate:    coordinate,
		Capacity:      capacity,
		Legality:      legality,
		Certificate:   certificate,
	}

	farmJSON, err := json.Marshal(farm)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmJSON)
}

// UpdateFarm updates an existing farm on the ledger
func (pc *PalmOilContract) UpdateFarm(ctx contractapi.TransactionContextInterface, id string, owner string, plantedYear int, seedVarieties string, area float64, address string, coordinate string, capacity float64, legality string, certificate string) error {
	farmJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if farmJSON == nil {
		return fmt.Errorf("the farm with ID %s does not exist", id)
	}

	var farm Farm
	json.Unmarshal(farmJSON, &farm)

	// Update the farm's attributes
	farm.Owner = owner
	farm.PlantedYear = plantedYear
	farm.SeedVarieties = seedVarieties
	farm.Area = area
	farm.Address = address
	farm.Coordinate = coordinate
	farm.Capacity = capacity
	farm.Legality = legality
	farm.Certificate = certificate

	farmJSON, err = json.Marshal(farm)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmJSON)
}

// QueryFarmByID retrieves a farm by its ID from the ledger
func (pc *PalmOilContract) QueryFarmByID(ctx contractapi.TransactionContextInterface, id string) (*Farm, error) {
	farmJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if farmJSON == nil {
		return nil, fmt.Errorf("the farm with ID %s does not exist", id)
	}

	var farm Farm
	json.Unmarshal(farmJSON, &farm)

	return &farm, nil
}

// QueryAllFarms retrieves all farms from the ledger
func (pc *PalmOilContract) QueryAllFarms(ctx contractapi.TransactionContextInterface) ([]*Farm, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var farms []*Farm
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var farm Farm
		json.Unmarshal(queryResponse.Value, &farm)
		farms = append(farms, &farm)
	}

	return farms, nil
}
