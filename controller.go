package main

import
(
	//"fmt"
	"log"
	"strings"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func SaveController(w http.ResponseWriter, r *http.Request){
		// Setting return type
		w.Header().Set("Content-Type", "application/json")

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		var inputForm CDRModel
		err = json.Unmarshal(reqBody, &inputForm)
		//fmt.Println(string(reqBody))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		err=nil
		success, message:=SaveService(inputForm)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
	
		// Status code
		w.WriteHeader(http.StatusOK)
		
		// Body
		data := simplejson.New()
		data.Set("success", success)
		data.Set("message", message)
		payload, err := data.MarshalJSON()
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		// Return response
		w.Write(payload)
		return
}

func SaveContractController(w http.ResponseWriter, r *http.Request){
	// Setting return type
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	var inputForm ContractModel
	err = json.Unmarshal(reqBody, &inputForm)
	//fmt.Println(string(reqBody))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	err=nil
	success, message:=SaveContractService(inputForm)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}

	// Status code
	w.WriteHeader(http.StatusOK)
	
	// Body
	data := simplejson.New()
	data.Set("success", success)
	data.Set("message", message)
	payload, err := data.MarshalJSON()
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	// Return response
	w.Write(payload)
	return
}

func GetController(w http.ResponseWriter, r *http.Request) {

	// Setting return type
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channelName := vars["channelName"]
	// value := vars["value"]
	
	x:=SearchService(channelName)
	
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write(getErrorPayload(message))
	// 	return
	// }

	// Status code
	w.WriteHeader(http.StatusOK)
	
	// Body
	data := simplejson.New()
	data.Set("data",x)
	// data.Set("success",success)
	// data.Set("message", message)
	payload, err := data.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	// Return response
	w.Write(payload)
	return
}

func GetContractController(w http.ResponseWriter, r *http.Request) {

	// Setting return type
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channelName := vars["channelName"]
	// value := vars["value"]
	
	x:=GetContractService(channelName)
	
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write(getErrorPayload(message))
	// 	return
	// }

	// Status code
	w.WriteHeader(http.StatusOK)
	
	// Body
	data := simplejson.New()
	data.Set("data",x)
	// data.Set("success",success)
	// data.Set("message", message)
	payload, err := data.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	// Return response
	w.Write(payload)
	return
}

func getErrorPayload(err error) []byte {
	var data = simplejson.New()
	errorParts := strings.Split(string(err.Error()), "Description: ")

	// Body
	data.Set("success", false)
	data.Set("error", errorParts[len(errorParts)-1])

	payload, err := data.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	return payload
}