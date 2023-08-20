package main

func main() {
	chaincode, err := contractapi.NewChaincode(new(PalmOilContract))
	if err != nil {
		fmt.Printf("Error create palm oil chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting palm oil chaincode: %s", err.Error())
	}
}