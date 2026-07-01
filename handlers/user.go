package handlers

import (
	"encoding/json"
	"net/http"

	"pases/models"
)

// request + response with controlled json file
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	// result :=
}
