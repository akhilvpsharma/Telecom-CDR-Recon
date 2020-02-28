package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
)

// QueryAsset query the chaincode to get the state of asset
func (setup *FabricSetup) QueryAsset(asset string) (string, error) {

	fmt.Println("Querying Asset:"+setup.ChainCodeID)
	// fmt.Println(setup.client)
	
	if setup.initialized {
		return "sdk already initialized for MRF", nil
	}
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return "failed to create SDK for MRF", nil
	}
	setup.sdk = sdk
	fmt.Println("SDK created for MRF!!")
	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return "failed to create new channel client by Gabriel", nil
	}
	fmt.Println("Channel client created by Gabriel")
	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, asset)
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}