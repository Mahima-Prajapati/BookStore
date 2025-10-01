package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mahima-Prajapati/BookStore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("App running on port '9010'")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
