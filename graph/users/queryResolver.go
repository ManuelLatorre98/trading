package users

import (
	"context"
	"database/sql"
)

type UserQueryResolver struct {
	DB *sql.DB
}

func (r *UserQueryResolver) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := r.DB.QueryRow(getUserByEmailPSQLQuery(), email)
	user := &User{}
	err := row.Scan(&user.userId, &user.username, &user.email, &user.registerDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserQueryResolver) GetUserById(ctx context.Context, userName string) (*User, error) {
	row := r.DB.QueryRow(getUserByIdPSQLQuery(), userName)
	user := &User{}
	err := row.Scan(&user.userId, &user.username, &user.email, &user.registerDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
