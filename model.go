package main
type CDRModel struct {
	CallingNo string `json:"callingNo"`
	CalledNo string `json:"calledNo"`
	Duration string `json:"duration"`
	EquipmentID string `json:"equipmentID"`
	CallResult string `json:"callResult"`
	Timestamp string `json:"timestamp"`
	CDRID string `json:"cdrID"`
}