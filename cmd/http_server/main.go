package main

import (
	"QuickShop/inventory"
	"QuickShop/pkg/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", users.AddNewUser)
	router.HandleFunc("/signin", users.Signin)
	router.HandleFunc("/inventory", inventory.AddToInventory)
	log.Fatal(http.ListenAndServe(":8192", router))
}
