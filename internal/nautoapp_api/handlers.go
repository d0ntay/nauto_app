package nautoapp_api

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const inventoryFile = "./internal/src/inventory.json"

type InventoryQuery struct {
	Query string `json:"query"`
}

type Device struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	Model    string `json:"model"`
}

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
	return
}

func (app *Application) getInventory(w http.ResponseWriter, r *http.Request) {
	var devices []Device
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	f, err := os.Open(inventoryFile)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		app.AppLogger.Error("Error opening inv file")
		return
	}
	err = json.NewDecoder(f).Decode(&devices)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		app.AppLogger.Error("Error decoding into devices")
		return
	}

	response := map[string]map[string]any{
		"response": {
			"devices": devices,
		},
	}
	json.NewEncoder(w).Encode(response)
	return

}

func (app *Application) queryInventory(w http.ResponseWriter, r *http.Request) {
	//	handler to return map[string]map[string]any
	//	{
	// 		"response" : {
	//			"devices" : [
	//				"hostname" : device.hostname	type : string
	//			]
	//		}
	//	}
	//
	// 	accepts requests bodies like this
	//
	// 	{
	// 		"query" : query 	type : string
	// 	}
	var request InventoryQuery
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var devices []map[string]string
	if r.Body == nil {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]map[string]any{
			"response": {
				"devices": devices,
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var devicesFromFile []Device
	f, err := os.Open(inventoryFile)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		app.AppLogger.Error("Error opening inv file")
		return
	}
	err = json.NewDecoder(f).Decode(&devicesFromFile)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		app.AppLogger.Error("Error decoding into devices")
		return
	}

	for _, device := range devicesFromFile {
		if strings.Contains(device.Hostname, request.Query) {
			devices = append(devices, map[string]string{"hostname": device.Hostname})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]map[string]any{
		"response": {
			"devices": devices,
		},
	}
	json.NewEncoder(w).Encode(response)

}
