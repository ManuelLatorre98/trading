package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"manulatorre98/trading/customErrors"
	"manulatorre98/trading/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	_, err := r.UserRepository.GetUserByEmail(input.Email)
	if err != nil && err.Error() != customErrors.ErrUserEmailNotFound {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	if err == nil {
		return nil, fmt.Errorf(customErrors.ErrUserEmailAlreadyExists)
	}

	_, err = r.UserRepository.GetUserByUserName(input.Username)
	if err != nil && err.Error() != customErrors.ErrUserNameNotFound {
		return nil, fmt.Errorf(customErrors.ErrInternalServer)
	}
	if err == nil {
		return nil, fmt.Errorf(customErrors.ErrUserNameAlreadyExists)
	}

	return r.UserRepository.InsertUser(input)
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByID - userById"))
}

// UserByEmail is the resolver for the userByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByEmail - userByEmail"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
