package db

const (
	CreateInventoryTable = "create table inventory(rfid_no text primary key,barcode text unique)"
	CreateUserTable      = "create table users(user_id serial primary key,first_name text,last_name text,email text unique,password text,age integer,gender text,qr bytea)"
	AddUser              = "INSERT INTO users(first_name,last_name,email,password,age,gender,qr) VALUES($1,$2,$3,$4,$5,$6,$7) returning user_id"
	GetCredentials       = "SELECT user_id,password from users where email=$1"
	InsertIntoInventory  = "INSERT INTO inventory(rfid_no,barcode) values($1,$2)"
)
