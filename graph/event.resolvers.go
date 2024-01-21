package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/entegral/gobox/dynamo"
	"github.com/entegral/officebuddy/types"
)

// Title is the resolver for the title field.
func (r *eventResolver) Title(ctx context.Context, obj *types.Event) (string, error) {
	if obj.Details != nil {
		return obj.Details.Title, nil
	}
	detail := &types.EventDetails{
		MonoLink: *dynamo.NewMonoLink(obj),
	}
	loaded, err := detail.Get(ctx, detail)
	if err != nil {
		return "", err
	}
	if loaded {
		obj.Details = detail
		return detail.Title, nil
	}
	return "", fmt.Errorf("event details not found")
}

// Description is the resolver for the description field.
func (r *eventResolver) Description(ctx context.Context, obj *types.Event) (string, error) {
	if obj.Details != nil {
		return obj.Details.Description, nil
	}
	detail := &types.EventDetails{
		MonoLink: *dynamo.NewMonoLink(obj),
	}
	loaded, err := detail.Get(ctx, detail)
	if err != nil {
		return "", err
	}
	if loaded {
		obj.Details = detail
		return detail.Description, nil
	}
	return "", fmt.Errorf("event details not found")
}

// Invites is the resolver for the Invites field.
func (r *eventResolver) Invites(ctx context.Context, obj *types.Event) ([]*types.Invite, error) {
	return dynamo.FindLinksByEntity0[*types.Event, *types.Invite](ctx, obj, "Invite")
}

// Venue is the resolver for the Venue field.
func (r *eventResolver) Venue(ctx context.Context, obj *types.Event) ([]*types.Venue, error) {
	venues, err := dynamo.FindLinksByEntity0[*types.Event, *types.Venue](ctx, obj, "Venue")
	if err != nil {
		return nil, err
	}
	return venues, nil
}

// PutEvents is the resolver for the putEvents field.
func (r *mutationResolver) PutEvents(ctx context.Context, events []*types.EventInput) ([]*types.Event, error) {
	ret := make([]*types.Event, len(events))
	for i, event := range events {
		// put the event
		err := event.Put(ctx, &event.Event)
		if err != nil {
			return nil, err
		}
		// put the event details, the fields are already handled by gqlgen, so we just need to set the MonoLink
		// so the details can leverage the Event's Keys() method.
		event.EventDetails.MonoLink = *dynamo.NewMonoLink(&event.Event)
		err = event.Put(ctx, &event.EventDetails)
		if err != nil {
			return nil, err
		}
		event.Event.Details = &event.EventDetails
		ret[i] = &event.Event
	}
	return ret, nil
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, userEmail string, eventGUID string) (bool, error) {
	event := &types.Event{GUID: eventGUID}
	err := event.Delete(ctx, event)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Event returns EventResolver implementation.
func (r *Resolver) Event() EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
