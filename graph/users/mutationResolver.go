package users

/*import (
	"context"
	"manulatorre98/trading/graph"
	"manulatorre98/trading/graph/model"
)

func CreateUser(ctx context.Context, input *model.UserInput, r graph.MutationResolver) (*UserModel, error) {
	var newUser *UserModel
	err := r.DB.QueryRow(insertUserPSQLQuery(), input.Username, input.Email, input.Password).Scan(
		&newUser.userId,
		&newUser.username,
		&newUser.email,
		&newUser.registerDate)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}*/
