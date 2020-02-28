package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Serve() {

	router := mux.NewRouter()
	
	router.HandleFunc("/save-cdr", SaveController).Methods("POST")
	router.HandleFunc("/get-cdr/{channelName}", GetController).Methods("GET")
	router.HandleFunc("/save-contract", SaveContractController).Methods("POST")
	router.HandleFunc("/get-contract/{channelName}", GetContractController).Methods("GET")
	
	fmt.Println("Listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}