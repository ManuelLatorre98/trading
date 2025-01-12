package userRepository

import "manulatorre98/trading/graph/model"

type UserRepository interface {
	InsertUser(user *model.UserInput) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUserName(nickName string) (*model.User, error)
}
