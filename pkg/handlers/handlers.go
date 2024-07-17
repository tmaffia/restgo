package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "up"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
