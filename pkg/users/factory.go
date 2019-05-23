package users

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int32  `json:"age"`
	Gender    string `json:"gender"`
}
type response struct {
	Response string `json:"response"`
}