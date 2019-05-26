package db

const (
	AddUser             = "INSERT INTO users(first_name,last_name,email,password,age,gender,qr) VALUES($1,$2,$3,$4,$5,$6,$7) returning user_id"
	GetCredentials      = "SELECT user_id,password from users where email=$1"
	InsertIntoInventory = "INSERT INTO inventory(rfid_no,barcode) values($1,$2)"
	GetUserDetails      = "SELECT user_id,first_name,last_name,email,age,gender,qr from users where user_id=$1"
)
