package model

type UserModel struct {
	UserId       int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisterDate string `json:"registerDate"`
}
