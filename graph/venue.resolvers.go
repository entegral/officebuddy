package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/clients"
)

// PutVenue is the resolver for the putVenue field.
func (r *mutationResolver) PutVenue(ctx context.Context, officeGUID string, eventGUID string, room *string, instructions *string) (*types.Venue, error) {
	venue, err := types.NewVenue(ctx, types.Event{GUID: eventGUID}, types.Office{GUID: officeGUID}, &types.NewVenueOpts{Room: room, Instructions: instructions})
	if err != nil {
		return nil, err
	}
	return venue, venue.Link(ctx, *clients.GetDefaultClient(ctx))
}

// Office is the resolver for the Office field.
func (r *venueResolver) Office(ctx context.Context, obj *types.Venue) (*types.Office, error) {
	_, err := obj.LoadEntity1(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	return obj.Entity1, nil
}

// Events is the resolver for the Events field.
func (r *venueResolver) Events(ctx context.Context, obj *types.Venue) (*types.Event, error) {
	_, err := obj.LoadEntity0(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	return obj.Entity0, nil
}

// Venue returns VenueResolver implementation.
func (r *Resolver) Venue() VenueResolver { return &venueResolver{r} }

type venueResolver struct{ *Resolver }
