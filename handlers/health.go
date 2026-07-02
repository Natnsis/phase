package handlers

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Status string
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := Status{
		Status: "ok",
	}
	json.NewEncoder(w).Encode(response)
}
