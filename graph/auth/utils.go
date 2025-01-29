package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/argon2"
	"manulatorre98/trading/customErrors"
	"os"
	"strconv"
	"time"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func hashPassword(password string) (hashedPass string, err error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return fmt.Sprintf("%s.%s", encodedSalt, encodedHash), nil
}
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateToken(userID string) (*string, error) {
	expirationTime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	return &tokenString, nil
}
