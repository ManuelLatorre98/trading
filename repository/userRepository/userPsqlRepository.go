package userRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"manulatorre98/trading/customErrors"
	"manulatorre98/trading/graph/model"
)

type UserPsqlRepository struct {
	db *sql.DB
}

func NewUserPsqlRepository(db *sql.DB) *UserPsqlRepository {
	return &UserPsqlRepository{db: db}
}

func (r *UserPsqlRepository) InsertUser(user *model.SignUpInput) (*model.User, error) {
	var newUser model.User
	err := r.db.QueryRow(insertUserPSQLQuery(), user.Username, user.Email, user.Password).Scan(
		&newUser.UserID,
		&newUser.Username,
		&newUser.Email, &newUser.RegisterDate,
	)
	return &newUser, err
}

func (r *UserPsqlRepository) GetUserByEmail(email string) (*model.User, error) {
	row := r.db.QueryRow(getUserByEmailPSQLQuery(), email)
	user := &model.User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(customErrors.ErrUserEmailNotFound)
		}
		return nil, err
	}

	return user, nil
}

func (r *UserPsqlRepository) GetUserByUserName(nickName string) (*model.User, error) {
	row := r.db.QueryRow(getUserByUserNamePSQLQuery(), nickName)
	user := &model.User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(customErrors.ErrUserNameNotFound)
		}
		return nil, err
	}
	return user, nil
}
func (r *UserPsqlRepository) GetUserById(userId string) (*model.User, error) {
	row := r.db.QueryRow(getUserByIdPSQLQuery(), userId)
	user := &model.User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(customErrors.ErrUserIdNotFound)
		}
		return nil, err
	}
	return user, nil
}

func (r *UserPsqlRepository) GetPassword(email string) (*string, error) {
	row := r.db.QueryRow(getEmailPasswordPSQLQuery(), email)
	var pass string
	err := row.Scan(&pass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(customErrors.ErrEmailOrPassWrong)
		}
		return nil, err
	}
	return &pass, nil
}
