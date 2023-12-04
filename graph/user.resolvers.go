package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

// Users is the resolver for the Users field.
func (r *mutationResolver) Users(ctx context.Context, input []*types.UserSaver) ([]*types.User, error) {
	ret := []*types.User{}
	for _, u := range input {
		// newUser, err := types.NewUser(ctx, u.GUID, u.Email, u.FirstName, u.LastName)
		// if err != nil {
		// 	return nil, err
		// }
		_, err := helpers.PutItem(ctx, &u.User)
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
		user := u.User
		loaded, err := helpers.GetItem(ctx, &user)
		if err != nil {
			return nil, err
		}
		if loaded {
			ret = append(ret, &user)
		}
	}
	return ret, nil
}

// Memberships is the resolver for the memberships field.
func (r *userResolver) Memberships(ctx context.Context, obj *types.User, roles []types.Role) ([]*types.Membership, error) {
	memberships, err := obj.Memberships(ctx, r.Clients)
	if err != nil {
		return nil, err
	}
	withRoles := []*types.Membership{}
	if len(roles) == 0 {
		return memberships, nil
	}
	for _, m := range memberships {
		for _, r := range roles {
			if m.Role == r {
				withRoles = append(withRoles, m)
			}
		}
	}
	return withRoles, nil
}

// Invites is the resolver for the Invites field.
func (r *userResolver) Invites(ctx context.Context, obj *types.User, status []types.InviteStatus) ([]*types.Invite, error) {
	invites, err := dynamo.FindCustomLinksByEntity1[*types.Event, *types.User, *types.Invite](ctx, r.Clients, obj)
	if err != nil {
		return nil, err
	}
	if len(status) == 0 {
		return invites, nil
	}
	withStatus := []*types.Invite{}
	for _, i := range invites {
		for _, s := range status {
			if i.Status == s {
				withStatus = append(withStatus, i)
			}
		}
	}
	return withStatus, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
