package main
import
(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	// "github.com/Volkswagen_World/blockchain"
)

var data map[string]CDRModel

func SaveService(cdrModel CDRModel) (string, string){
	fmt.Println("Searching...")
	if data==nil {
		data=make(map[string]CDRModel)
	}
	data[cdrModel.CDRID]=cdrModel
	// blockchain.InvokeChaincode(webForm.Key,webForm.Value)
	return "true", "CDR is saved."
}

func SearchService() ([]CDRModel){
	fmt.Println("Searching...")
	var dataList []CDRModel
	dataList=make([]CDRModel, 0)
	for _, v := range data { 
		dataList=append(dataList,v)
	}
	// foundValue:=blockchain.QueryChaincode(key)
	return dataList
}