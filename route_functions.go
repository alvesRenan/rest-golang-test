package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var c ClientContainer

// adds a new device to the slice
func createContainer(w http.ResponseWriter, r *http.Request) {
	c.ID = rand.Intn(100)

	json.NewDecoder(r.Body).Decode(&c)
	mockData = append(mockData, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Success!"))
}

// retunrs all devices in the slice
func getDevices(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockData)
}

// deletes a device of the slice by name
func deleteDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for index, device := range mockData {
		// check id
		if strconv.Itoa(device.ID) == mux.Vars(r)["id"] {
			// if the same, delete from slice
			mockData = append(mockData[:index], mockData[index+1:]...)
			json.NewEncoder(w).Encode(mockData)

			//return to break the loop
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Device not found!"))
}

func updateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(r.Body).Decode(&c)

	for _, device := range mockData {
		if strconv.Itoa(device.ID) == mux.Vars(r)["id"] {
			device.Name = c.Name
			device.Network = c.Network
			device.AdbPort = c.AdbPort
			device.VNCPort = c.VNCPort

			json.NewEncoder(w).Encode(device)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Device not found!"))
}
