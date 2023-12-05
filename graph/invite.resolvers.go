package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
)

// User is the resolver for the User field.
func (r *inviteResolver) User(ctx context.Context, obj *types.Invite) (*types.User, error) {
	return obj.User(ctx, *clients.GetDefaultClient(ctx))
}

// Event is the resolver for the Event field.
func (r *inviteResolver) Event(ctx context.Context, obj *types.Invite) (*types.Event, error) {
	return obj.Event(ctx, *clients.GetDefaultClient(ctx))
}

// PutInvite is the resolver for the putInvite field.
func (r *mutationResolver) PutInvite(ctx context.Context, userEmail string, eventGUID string, status types.InviteStatus) (*types.Invite, error) {
	dilink, newErr := dynamo.CheckLink[*types.Event, *types.User](&types.Event{CreatedByEmail: userEmail, GUID: eventGUID}, &types.User{Email: userEmail})
	invite := &types.Invite{
		DiLink: *dilink,
		Status: status,
	}
	switch newErr.(type) {
	case nil:
		// invite already exists
		return invite, nil
	case dynamo.ErrLinkNotFound:
		// invite does not exist
		return invite, invite.Put(ctx, invite)
	default:
		return nil, newErr
	}
}

// DeleteInvite is the resolver for the deleteInvite field.
func (r *mutationResolver) DeleteInvite(ctx context.Context, userEmail string, eventGUID string) (*types.Invite, error) {
	invite, err := types.NewInvite(ctx, types.Event{CreatedByEmail: userEmail, GUID: eventGUID}, types.User{Email: userEmail}, nil)
	if err != nil {
		return nil, err
	}
	return invite, invite.Delete(ctx, invite)
}

// Invite returns InviteResolver implementation.
func (r *Resolver) Invite() InviteResolver { return &inviteResolver{r} }

type inviteResolver struct{ *Resolver }
