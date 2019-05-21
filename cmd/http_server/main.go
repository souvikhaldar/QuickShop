package main

import (
	"QuickShop/pkg/users"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", users.AddNewUser)
}
