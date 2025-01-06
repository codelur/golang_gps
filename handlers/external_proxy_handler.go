package handlers

import (
	"fmt"
	"io"
	"net/http"
	"project/utils"
)

func ProxyExternalRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", utils.Origin)
	id := r.URL.Query().Get("id")
	externalURL := fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device%s?latest_point=true&api-key=%s", id, utils.OneStepGPSKey)

	resp, err := http.Get(externalURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making external request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
