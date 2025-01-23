package users

/*type UserQueryResolver struct {
	DB *sql.DB
}
type QueryResolver struct{ DB *sql.DB }
type UserResolver struct{}

func (r *QueryResolver) UserByID(ctx context.Context, email string) (*model.UserModel, error) {
	row := r.DB.QueryRow(userRepository.getUserByEmailPSQLQuery(), email)
	user := &model.UserModel{}
	err := row.Scan(&user.userId, &user.username, &user.email, &user.registerDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *QueryResolver) UserByEmail(ctx context.Context, userId string) (*model.UserModel, error) {
	row := r.DB.QueryRow(userRepository.getUserByIdPSQLQuery(), userId)
	user := &model.UserModel{}
	err := row.Scan(&user.userId, &user.username, &user.email, &user.registerDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserID is the resolver for the userId field.
func (r *UserResolver) UserID(ctx context.Context, obj *model.UserModel) (*string, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// Username is the resolver for the username field.
func (r *UserResolver) Username(ctx context.Context, obj *model.UserModel) (*string, error) {
	panic(fmt.Errorf("not implemented: Username - username"))
}

// Email is the resolver for the email field.
func (r *UserResolver) Email(ctx context.Context, obj *model.UserModel) (*string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// RegisterDate is the resolver for the registerDate field.
func (r *UserResolver) RegisterDate(ctx context.Context, obj *model.UserModel) (*string, error) {
	panic(fmt.Errorf("not implemented: RegisterDate - registerDate"))
}
*/
