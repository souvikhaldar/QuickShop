package inventory

import (
	"QuickShop/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddToInventory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--Running in AddToInventory--")
	var i Inventory
	if er := json.NewDecoder(r.Body).Decode(&i); er != nil {
		err := fmt.Sprintln("Error in unmarshalling the inventory body: ", er.Error())
		fmt.Println(err)
		http.Error(w, err, 500)
		return
	}
	fmt.Println("The inventory details recieved: ", i)
	if _, er := db.DB.Exec(db.InsertIntoInventory, i.Rfid_no, i.Barcode); er != nil {
		err := fmt.Sprintln("Error in inserting the inventory details: ", er.Error())
		fmt.Println(err)
		http.Error(w, err, 500)
		return
	}
	fmt.Fprint(w, "Successfully inserted to inventory")
}
