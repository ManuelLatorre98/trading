package users

type User struct {
	userId       int    `json:"id"`
	username     string `json:"username"`
	email        string `json:"email"`
	password     string `json:"password"`
	registerDate string `json:"registerDate"`
}
