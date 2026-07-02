package main

import (
	"fmt"
	"log"
	"net/http"

	"pases/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.MessageRoutes(router)
	routes.NewRoutes(router)
	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
