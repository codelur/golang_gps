package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/utils"
)

func MapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", utils.Origin)

	response := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/js?key=%s&libraries=marker",
		utils.GoogleMapsApiKey,
	)

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
