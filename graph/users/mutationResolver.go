package users

import (
	"database/sql"
)

type UserMutationResolver struct{ DB *sql.DB }

func (r *UserMutationResolver) CreateUser(username string, email string, password string) (*User, error) {

	var newUser *User
	err := r.DB.QueryRow(insertUserPSQLQuery(), username, email, password).Scan(
		&newUser.userId,
		&newUser.username,
		&newUser.email,
		&newUser.registerDate)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
