package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var mockData []ClientContainer = []ClientContainer{}

func main() {
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/create/client", createContainer).Methods("POST")
	r.HandleFunc("/get/devices", getDevices).Methods("GET")
	r.HandleFunc("/delete/device/{id}", deleteDevice).Methods("DELETE")
	r.HandleFunc("/update/device/{id}", updateDevice).Methods("UPDATE")

	// starts server
	fmt.Println("Server running....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
