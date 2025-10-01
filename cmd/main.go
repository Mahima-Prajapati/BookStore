package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mahima-Prajapati/BookStore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	port := os.Getenv("APP_PORT")
	fmt.Println("App running on port " + port)

	domain := os.Getenv("DOMAIN")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(domain, r))
}
