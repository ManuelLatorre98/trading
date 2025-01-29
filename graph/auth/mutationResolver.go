package auth

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"manulatorre98/trading/customErrors"
	"manulatorre98/trading/graph/model"
	"manulatorre98/trading/repository/userRepository"
	"strings"
)

func SignUpResolver(repo userRepository.UserRepository, input *model.SignUpInput) (*model.User,
	error) {
	_, err := repo.GetUserByEmail(input.Email)
	if err != nil && err.Error() != customErrors.ErrUserEmailNotFound {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	if err == nil {
		return nil, fmt.Errorf(customErrors.ErrUserEmailAlreadyExists)
	}

	_, err = repo.GetUserByUserName(input.Username)
	if err != nil && err.Error() != customErrors.ErrUserNameNotFound {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	if err == nil {
		return nil, fmt.Errorf(customErrors.ErrUserNameAlreadyExists)
	}
	hashPass, err := hashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	input.Password = string(hashPass)
	return repo.InsertUser(input)
}

func LoginResolver(repo userRepository.UserRepository, input *model.LoginInput) (*model.AuthPayload, error) {
	//Validate email
	userStored, err := repo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	//Validate password
	storedPass, err := repo.GetPassword(input.Email)
	parts := strings.Split(*storedPass, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}

	encodedSalt := parts[0]
	storedEncodedHash := parts[1]

	salt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	hash := argon2.IDKey([]byte(input.Password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	inputEncodedHash := base64.RawStdEncoding.EncodeToString(hash)
	if inputEncodedHash != storedEncodedHash {
		return nil, fmt.Errorf(customErrors.ErrEmailOrPassWrong)
	}

	//Generate TOKEN
	token, err := generateToken(*userStored.UserID)
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	return &model.AuthPayload{Token: token, User: userStored}, nil
}
