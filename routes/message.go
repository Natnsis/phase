package routes

import (
	"pases/handlers"

	"github.com/gorilla/mux"
)

func MessageRoutes(router *mux.Router) {
	router.HandleFunc("/respond",
		handlers.MessageReturner).Methods("POST")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
}
