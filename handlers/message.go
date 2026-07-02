package handlers

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string
}

// request + response with controlled json file
func MessageReturner(w http.ResponseWriter, r *http.Request) {
	// takes message then returns recived and length
	var message Message
	json.NewDecoder(r.Body).Decode(&message)

	// return the Body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// return the same message
	json.NewEncoder(w).Encode(message)
}
