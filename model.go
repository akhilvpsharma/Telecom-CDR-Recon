package main
type CDRModel struct {

	ChannelName string `json:"channelName"`
	
	OrgName string `json:"orgName"`
	CallingNo string `json:"callingNo"`
	CalledNo string `json:"calledNo"`
	Duration string `json:"duration"`
	EquipmentID string `json:"equipmentID"`
	CallResult string `json:"callResult"`
	Timestamp string `json:"timestamp, omitempty"`
	CDRID string `json:"cdrID"`
}

type ContractModel struct {

	Rate string `json:"rate"`	
	ChannelName string `json:"channelName"`
	Cycle string `json:"cycle"`
	Timestamp string `json:"timestamp"`
}