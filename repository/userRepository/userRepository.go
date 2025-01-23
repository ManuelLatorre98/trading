package userRepository

import "manulatorre98/trading/graph/model"

type UserRepository interface {
	InsertUser(user *model.SignUpInput) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUserName(nickName string) (*model.User, error)
	GetUserById(userId string) (*model.User, error)
	GetPassword(email string) (*string, error)
}
