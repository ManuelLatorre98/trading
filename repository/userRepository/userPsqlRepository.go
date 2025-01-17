package userRepository

import (
	"database/sql"
	"manulatorre98/trading/graph/model"
)

type UserPsqlRepository struct {
	db *sql.DB
}

func NewUserPsqlRepository(db *sql.DB) *UserPsqlRepository {
	return &UserPsqlRepository{db: db}
}

func (r *UserPsqlRepository) InsertUser(user *model.UserInput) (*model.User, error) {
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
		return nil, err
	}
	return user, nil
}

func (r *UserPsqlRepository) GetUserByUserName(nickName string) (*model.User, error) {
	row := r.db.QueryRow(getUserByIdPSQLQuery(), nickName)
	user := &model.User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
