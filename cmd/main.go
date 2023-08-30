package main

import (

	"net/http"
	"log"
	"github.com/gorilla/mux"

	"github.com/SEANYB4/go-bookstore-api/pkg/routes"
)


func main() {


	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))




}