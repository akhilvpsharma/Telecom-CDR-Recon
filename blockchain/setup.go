package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
	"os"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile      string
	OrgID           string
	OrdererID 		string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	client          *channel.Client
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	event           *event.Client
}

func (setup *FabricSetup) Initialize() error {

	if setup.initialized {
		return errors.New("sdk already initialized.")
	}

	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK for")
	}
	setup.sdk = sdk
	fmt.Println("SDK created.")

	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity for MRF")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity for MRF")
	}
	setup.admin = resMgmtClient
	fmt.Println("Resource management client created for MRF")

	// The MSP client allow us to retrieve user information from their identity, like its signing identity which we will need to save the channel
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to create MSP client for MRF")
	}
	adminIdentity, err := mspClient.GetSigningIdentity(setup.OrgAdmin)
	if err != nil {
		return errors.WithMessage(err, "failed to get admin signing identity for MRF")
	}
	req := resmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfigPath: setup.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := setup.admin.SaveChannel(req, resmgmt.WithOrdererEndpoint(setup.OrdererID))
	if err != nil || txID.TransactionID == "" {
		return errors.WithMessage(err, "failed to save channel for MRF")
	}
	fmt.Println("Channel created by MRF")

	// Make admin user join the previously created channel
	if err = setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
		return errors.WithMessage(err, "failed to make admin join channel for MRF")
	}
	fmt.Println("Channel joined for MRF")

	fmt.Println("Initialization Successful for MRF")
	setup.initialized = true
	return nil
}


func (setup *FabricSetup) JoinTheChannelAndInstallCC() error {

	if setup.initialized {
		return errors.New("sdk already initialized for Gabriel")
	}

	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK for Gabriel")
	}
	setup.sdk = sdk
	fmt.Println("SDK created for Gabriel !!")

	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity for Gabriel")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity for Gabriel")
	}
	setup.admin = resMgmtClient
	fmt.Println("Resource management client created for Gabriel")

	// Make admin user join the previously created channel
	if err := setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
		return errors.WithMessage(err, "failed to make admin join channel for Gabriel")
	}
	fmt.Println("Channel joined for Gabriel")

	fmt.Println("Initialization Successful for Gabriel")
	setup.initialized = true

	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
	if err != nil {
		return errors.WithMessage(err, "failed to create chaincode package for Gabriel")
	}
	fmt.Println("ccPkg created for Gabriel")

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
	_, err = setup.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.WithMessage(err, "failed to install chaincode for Gabriel")
	}
	fmt.Println("Chaincode installed for Gabriel")

	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new channel client by Gabriel")
	}
	fmt.Println("Channel client created by Gabriel")

	// Creation of the client which will enables access to our channel events
	setup.event, err = event.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new event client by Gabriel")
	}
	fmt.Println("Event client created by Gabriel")
	
	
	return nil
}


func (setup *FabricSetup) InstallAndInstantiateCC() error {

	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
	if err != nil {
		return errors.WithMessage(err, "failed to create chaincode package by MRF")
	}
	fmt.Println("ccPkg created by MRF")

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
	_, err = setup.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.WithMessage(err, "failed to install chaincode by MRF")
	}
	fmt.Println("Chaincode installed by MRF")

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"mrf.volkswagenag.com"})

	resp, err := setup.admin.InstantiateCC(setup.ChannelID, resmgmt.InstantiateCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodeGoPath, Version: "0", Args: [][]byte{[]byte("Polo"), []byte("100000")}, Policy: ccPolicy})
	if err != nil || resp.TransactionID == "" {
		return errors.WithMessage(err, "failed to instantiate the chaincode by MRF")
	}
	fmt.Println("Chaincode instantiated by MRF")

	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new channel client by MRF")
	}
	fmt.Println("Channel client created by MRF")

	// Creation of the client which will enables access to our channel events
	setup.event, err = event.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new event client by MRF")
	}
	fmt.Println("Event client created by MRF")

	fmt.Println("Chaincode Installation & Instantiation Successful by MRF")
	return nil
}

func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}

func InvokeChaincode( key string,value string ) (string, error) {
	fSetupOrg2 := FabricSetup{
		// Network parameters 
		OrdererID: "orderer.volkswagenag.com",

		// Channel parameters
		ChannelID:     "indianchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/Volkswagen_World/fixtures/artifacts/volkswagen.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "volkswagen_world",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/Volkswagen_World/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "mrf",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}
	// Invoke the chaincode
	txId, err := fSetupOrg2.SetAsset([]string{key,value})
	if err != nil {
		fmt.Printf("Unable to invoke asset on the chaincode: %v\n", err)
		return "Unable to invoke asset on the chaincode: %v\n", err
	} else {
		fmt.Printf("Successfully invoked the asset, transaction ID: %s\n", txId)
		return "Successfully invoked the asset", nil
	}
}

func QueryChaincode(key string) (string) {
	fSetupOrg1 := FabricSetup{
		// Network parameters 
		OrdererID: "orderer.volkswagenag.com",

		// Channel parameters
		ChannelID:     "indianchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/Volkswagen_World/fixtures/artifacts/volkswagen.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "volkswagen_world",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/Volkswagen_World/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "gabriel",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}
	fmt.Printf("\n-----------Querying -----------------"+fSetupOrg1.ChainCodeID)
	// Query the chaincode
	response, err := fSetupOrg1.QueryAsset(key)
	if err != nil {
		fmt.Printf("Unable to query asset on the chaincode: %v\n", err)
		return "Unable to query asset on the chaincode: %v\n"
	} else {
		fmt.Printf("Response from the asset query: %s\n", response)
		return response
	}
}
