package main

import (
	"log"
	"net/http"

	"github.com/deankinane/go-bookstore-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
		return
	}
}
