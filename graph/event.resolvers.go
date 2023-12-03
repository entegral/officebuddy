package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
	types1 "github.com/entegral/toolbox/types"
)

// PutEvent is the resolver for the putEvent field.
func (r *mutationResolver) PutEvent(ctx context.Context, userEmail string, officeGUID string, title string, description string, start types1.DateTime, end types1.DateTime) (*types.Event, error) {
	return types.NewEvent(ctx, userEmail, officeGUID, title, description, start, end)
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, userEmail string, officeGUID string) (bool, error) {
	event, err := dynamo.CheckLink[*types.User, *types.Office](&types.User{Email: userEmail}, &types.Office{GUID: officeGUID}, dynamo.OneToMany)
	switch err.(type) {
	case nil:
		return true, event.Unlink(ctx, *clients.GetDefaultClient(ctx))
	case dynamo.ErrLinkNotFound:
		return false, nil
	default:
		return false, err
	}
}
