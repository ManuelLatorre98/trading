package users

import (
	"manulatorre98/trading/graph/model"
	"manulatorre98/trading/repository/userRepository"
)

func UserByIDResolver(repo userRepository.UserRepository, userId string) (*model.User, error) {
	return repo.GetUserById(userId)

}
func UserByEmailResolver(repo userRepository.UserRepository, email string) (*model.User, error) {
	return repo.GetUserByEmail(email)
}
