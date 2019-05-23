package main

import (
	"QuickShop/pkg/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", users.AddNewUser)
	log.Fatal(http.ListenAndServe(":8192", router))
}
