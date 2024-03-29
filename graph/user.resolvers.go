package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/gobox/dynamo"
	"github.com/entegral/officebuddy/types"
)

// Users is the resolver for the Users field.
func (r *mutationResolver) Users(ctx context.Context, input []*types.UserSaver) ([]*types.User, error) {
	ret := []*types.User{}
	for _, u := range input {
		err := u.Put(ctx, &u.User)
		if err != nil {
			return nil, err
		}
		u.UserDetails.MonoLink = *dynamo.NewMonoLink(&u.User)
		err = u.Put(ctx, &u.UserDetails)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &u.User)
	}
	return ret, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, input []*types.UserFinder) ([]*types.User, error) {
	ret := []*types.User{}
	for _, u := range input {
		loaded, err := u.User.Get(ctx, &u.User)
		if err != nil {
			return nil, err
		}
		if loaded {
			ret = append(ret, &u.User)
		}
	}
	return ret, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
