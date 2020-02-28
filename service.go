package main
import
(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	// "github.com/Volkswagen_World/blockchain"
)

// var data map[string]CDRModel
var channel12Data []CDRModel
var channel13Data []CDRModel
var channel23Data []CDRModel

var channel12ContractData []ContractModel
var channel13ContractData []ContractModel
var channel23ContractData []ContractModel

func SaveService(cdrData CDRModel) (string, string){
	fmt.Println("Saving...")
	if(channel12Data==nil) {
		channel12Data=make([]CDRModel, 0)
	}
	if(channel13Data==nil) {
		channel13Data=make([]CDRModel, 0)
	}
	if(channel23Data==nil) {
		channel23Data=make([]CDRModel, 0)
	}
	
	if (cdrData.ChannelName == "channel12") {
		channel12Data=append(channel12Data, cdrData)
	}
	if (cdrData.ChannelName == "channel13") {
		channel13Data=append(channel13Data, cdrData)
	}
	if (cdrData.ChannelName == "channel23") {
		channel23Data=append(channel23Data, cdrData)
	}
	// blockchain.InvokeChaincode(webForm.Key,webForm.Value)
	return "true", "CDR is saved."
}

func SearchService(channelName string) ([]CDRModel){
	fmt.Println("Searching...")
	var dataList []CDRModel
	dataList=make([]CDRModel, 0)

	if(channelName == "channel12") {
		for _, v := range channel12Data { 
			dataList=append(dataList,v)
		}
	}
	if(channelName == "channel13") {
		for _, v := range channel13Data { 
			dataList=append(dataList,v)
		}
	}
	if(channelName == "channel23") {
		for _, v := range channel23Data { 
			dataList=append(dataList,v)
		}
	}
	// foundValue:=blockchain.QueryChaincode(key)
	return dataList
}

func SaveContractService(contractData ContractModel) (string, string){
	fmt.Println("Saving Contract...")
	if(channel12ContractData==nil) {
		channel12ContractData=make([]ContractModel, 0)
	}
	if(channel13ContractData==nil) {
		channel13ContractData=make([]ContractModel, 0)
	}
	if(channel23Data==nil) {
		channel23ContractData=make([]ContractModel, 0)
	}
	
	if (contractData.ChannelName == "channel12") {
		channel12ContractData=append(channel12ContractData, contractData)
	}
	if (contractData.ChannelName == "channel13") {
		channel13ContractData=append(channel13ContractData, contractData)
	}
	if (contractData.ChannelName == "channel23") {
		channel23ContractData=append(channel23ContractData, contractData)
	}
	// blockchain.InvokeChaincode(webForm.Key,webForm.Value)
	return "true", "Contract is saved."
}

func GetContractService(channelName string) ([]ContractModel){
	fmt.Println("Searching Contract...")
	var dataList []ContractModel
	dataList=make([]ContractModel, 0)

	if(channelName == "channel12") {
		for _, v := range channel12ContractData { 
			dataList=append(dataList,v)
		}
	}
	if(channelName == "channel13") {
		for _, v := range channel13ContractData { 
			dataList=append(dataList,v)
		}
	}
	if(channelName == "channel23") {
		for _, v := range channel23ContractData { 
			dataList=append(dataList,v)
		}
	}
	// foundValue:=blockchain.QueryChaincode(key)
	return dataList
}