package routes

import (
	"pases/handlers"

	"github.com/gorilla/mux"
)

func NewRoutes(router *mux.Router) {
	router.HandleFunc("/new", handlers.SecondaryMessage).Methods("POST")
}
