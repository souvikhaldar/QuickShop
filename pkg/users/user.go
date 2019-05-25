package users

import (
	"QuickShop/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--Running in AddNewUser--")
	var u user
	er := json.NewDecoder(r.Body).Decode(&u)
	//Allow CORS here By * or specific origin
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if er != nil {
		fmt.Println("Error in unmarshalling the adduser request body ", er)
		http.Error(w, fmt.Sprintln("Error in unmarshalling the adduser request body ", er), 500)
		return
	}
	var userid int32
	if err := db.DB.QueryRow(db.AddUser, u.FirstName, u.LastName, u.Email, u.Password, u.Age, u.Gender, u.QR).Scan(&userid); err != nil {
		er := fmt.Sprintf("Unable to add a new user: %s", err)
		fmt.Println(err)
		http.Error(w, er, 500)
		return
	}
	fmt.Println("Successfully inserted user with userid:", userid)
	json.NewEncoder(w).Encode(&response{userid})
	return
}
