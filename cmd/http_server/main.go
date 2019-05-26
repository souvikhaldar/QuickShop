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
	router.HandleFunc("/signup", users.AddNewUser).Methods("POST")
	router.HandleFunc("/signin", users.Signin).Methods("POST")
	router.HandleFunc("/inventory", inventory.AddToInventory).Methods("POST")
	router.HandleFunc("/user/{user_id}/", users.GetUserDetails).Methods("GET")
	log.Fatal(http.ListenAndServe(":8192", router))
}
