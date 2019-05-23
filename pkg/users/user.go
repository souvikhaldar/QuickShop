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
	w.Header().Set("Content-Type", "application/json")
	if er != nil {
		fmt.Println("Error in unmarshalling the adduser request body ", er)
		json.NewEncoder(w).Encode(&response{"Error: Unable to unamrshall the request body"})
		return
	}
	var user_id int32
	if err := db.DB.QueryRow(db.AddUser, u.FirstName, u.LastName, u.Email, u.Password, u.Age, u.Gender).Scan(&user_id); err != nil {
		er := fmt.Sprintf("Unable to add a new user: %s", err)
		fmt.Println(err)
		http.Error(w, er, 500)
		return
	}
	fmt.Println("Successfully inserted user with userid:", user_id)
	fmt.Fprintf(w, "Successfully inserted user with userid:%d\n", user_id)
}
