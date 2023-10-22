package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KulwinderSingh07/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running on port 3001")
	log.Fatal(http.ListenAndServe("localhost:3001", r))
}
