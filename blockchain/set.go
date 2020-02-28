package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
)

// SetAsset query the chaincode to get the state of Asset
func (setup *FabricSetup) SetAsset(args []string) (string, error) {

	if setup.initialized {
		return "sdk already initialized for MRF", nil
	}
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return "failed to create SDK for MRF", nil
	}
	setup.sdk = sdk
	// fmt.Println("SDK created for MRF!!")
	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return "failed to create new channel client by Gabriel", nil
	}
	// fmt.Println("Channel client created by Gabriel")
	// Creation of the client which will enables access to our channel events
	setup.event, err = event.New(clientContext)
	if err != nil {
		return "", fmt.Errorf("failed to create new event client by Gabriel %v", err)
	}
	// fmt.Println("Event client created by Gabriel")

	// Prepare arguments

	args = append(args, "set")

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in SetAsset")

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[len(args)-1], Args: [][]byte{[]byte(args[0]), []byte(args[1])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to edit asset value: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 25):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.TransactionID), nil

}