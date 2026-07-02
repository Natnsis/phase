package handlers

import (
	"encoding/json"
	"net/http"
)

type TheMessage struct {
	Message string
}

type TheResponse struct {
	Response TheMessage
	Length   int
}

func SecondaryMessage(w http.ResponseWriter, r *http.Request) {
	var theMessage TheMessage
	json.NewDecoder(r.Body).Decode(&theMessage)

	response := TheResponse{
		Response: theMessage,
		Length:   len(theMessage.Message),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
