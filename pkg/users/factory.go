package users

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int32  `json:"age"`
	Gender    string `json:"gender"`
	QR        []byte `json:"qr"`
}
type response struct {
	UserID int32 `json:"user_id"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
