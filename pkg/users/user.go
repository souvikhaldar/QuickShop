package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	var u user
	er := json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	if er != nil {
		fmt.Println("Error in unmarshalling the adduser request body ", er)
		json.NewEncoder(w).Encode(&response{"Error: Unable to unamrshall the request body"})
		return
	}
}
