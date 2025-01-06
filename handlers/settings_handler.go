package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"project/models"
	"project/utils"
)

func GetSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", utils.Origin)

	file, err := os.Open(utils.SettingsFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening settings file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading settings file: %v", err), http.StatusInternalServerError)
		return
	}

	var settings models.Settings
	err = json.Unmarshal(fileContent, &settings)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding settings JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", utils.Origin)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var settings models.Settings
	err := json.NewDecoder(r.Body).Decode(&settings)
	if err != nil {
		http.Error(w, "Failed to parse settings", http.StatusBadRequest)
		return
	}

	fileData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		http.Error(w, "Failed to serialize settings", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(utils.SettingsFilePath, fileData, 0644)
	if err != nil {
		http.Error(w, "Failed to save settings to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Settings saved successfully"}`)
}
