package userRepository

import (
	"database/sql"
	"log"
	"manulatorre98/trading/graph/users"
	"manulatorre98/trading/model"
)

type UserLocalRepository struct {
	db *sql.DB
}

func NewUserLocalRepository(db *sql.DB) *UserLocalRepository {
	return &UserLocalRepository{db: db}
}

func (r *UserLocalRepository) InsertUser(user model.User) (int, error) {
	var id int
	err := r.db.QueryRow(users.insertUserPSQLQuery(), user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Printf("userRepository: Error inserting user: %v", err)
		return 0, err
	}
	return id, nil
}

func (r *UserLocalRepository) GetUserByEmail(email string) (*model.User, error) {
	row := r.db.QueryRow(users.getUserByEmailPSQLQuery(), email)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserLocalRepository) GetUserByUserName(nickName string) (*model.User, error) {
	row := r.db.QueryRow(users.getUserByUserNamePSQLQuery(), nickName)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.RegisterDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
