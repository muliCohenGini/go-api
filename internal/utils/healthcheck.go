package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(status)
}

func HealthCheckHandler(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods("GET")
}
