package main

import (
	// "fmt"
	// "github.com/Telecom-CDR-Recon/blockchain"
	// "os"
)
func main() {
	// Definition of the Fabric SDK properties
	// fSetupOrg1 := blockchain.FabricSetup{
	// 	// Network parameters 
	// 	OrdererID: "orderer.telecom.com",

	// 	// Channel parameters
	// 	ChannelID:     "org12channel",
	// 	ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/Telecom-CDR-Recon/channel-artifacts/channel12.tx",

	// 	// Chaincode parameters
	// 	ChainCodeID:     "cdr",
	// 	ChaincodeGoPath: os.Getenv("GOPATH"),
	// 	ChaincodePath:   "github.com/Telecom-CDR-Recon/chaincode/",
	// 	OrgAdmin:        "Admin",
	// 	OrgName:         "org1",
	// 	ConfigFile:      "config.yaml",

	// 	// User parameters
	// 	UserName: "User1",
	// }

	// fSetupOrg2 := blockchain.FabricSetup{
	// 	// Network parameters 
	// 	OrdererID: "orderer.telecom.com",

	// 	// Channel parameters
	// 	ChannelID:     "org12channel",
	// 	ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/Telecom-CDR-Recon/channel-artifacts/channel12.tx",

	// 	ChainCodeID:     "cdr",
	// 	ChaincodeGoPath: os.Getenv("GOPATH"),
	// 	ChaincodePath:   "github.com/Telecom-CDR-Recon/chaincode/",
	// 	OrgAdmin:        "Admin",
	// 	OrgName:         "org2",
	// 	ConfigFile:      "config.yaml",

	// 	// User parameters
	// 	UserName: "User1",
	// }

	// fmt.Printf("\n-----------Initialization of ORG1 & Creating, Joining Org12 Channel-----------------\n")
	// // Initialization of the Fabric SDK from the previously set properties
	// err := fSetupMRF.Initialize()
	// if err != nil {
	// 	fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	// 	return
	// }

	// fmt.Printf("\n-----------Initialization of Gabriel ORG & \nJoining channel and Installing CC GabrielOrg-----------------\n")
	// err = fSetupGabriel.JoinTheChannelAndInstallCC()
	// if err != nil {
	// 	fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	// 	return
	// }

	// fmt.Printf("\n-----------Install and instantiate CC by MRF ORG-----------------\n")
	// // Install and instantiate the chaincode
	// err = fSetupMRF.InstallAndInstantiateCC()
	// if err != nil {
	// 	fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	// 	return
	// }
	
	// fmt.Printf("\n-----------------------------CC Operations-----------------------------\n")

	// fmt.Printf("\n-----------Querying Car Polo's Price(MRF)-----------------\n")
	// // Query the chaincode
	// response, err := fSetupMRF.QueryAsset("Polo")
	// if err != nil {
	// 	fmt.Printf("Unable to query asset on the chaincode: %v\n", err)
	// } else {
	// 	fmt.Printf("Response from the asset query: %s\n", response)
	// }

	// fmt.Printf("\n-----------Changing Car Polo's Price(MRF)-----------------\n")

	// // Invoke the chaincode
	// txID, err := fSetupMRF.SetAsset([]string{"New_Polo","300000"})
	// if err != nil {
	// 	fmt.Printf("Unable to invoke asset on the chaincode: %v\n", err)
	// } else {
	// 	fmt.Printf("Successfully invoked the asset, transaction ID: %s\n", txID)
	// }

	// fmt.Printf("\n-----------Querying Car Polo's Price(Gabriel)-----------------\n")
	// // Query the chaincode
	// response, err = fSetupGabriel.QueryAsset("New_Polo")
	// if err != nil {
	// 	fmt.Printf("Unable to query asset on the chaincode: %v\n", err)
	// } else {
	// 	fmt.Printf("Response from the asset query: %s\n", response)
	// }

	// // Close SDK
	// defer fSetupMRF.CloseSDK()	
	// defer fSetupGabriel.CloseSDK()
	Serve()	

}