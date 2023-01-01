package main

import (
	"log"
	"net/http"

	"github.com/Gift-py/go-student-portal/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}
