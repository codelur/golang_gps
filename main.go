package main

import (
	"fmt"
	"net/http"

	"project/handlers"
)

func main() {
	http.HandleFunc("/api/getDevicesInfo", handlers.ProxyExternalRequest)
	http.HandleFunc("/api/map", handlers.MapHandler)
	http.HandleFunc("/api/savesettings", handlers.SaveSettings)
	http.HandleFunc("/api/getsettings", handlers.GetSettings)

	fmt.Println("Go server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
