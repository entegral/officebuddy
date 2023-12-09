package types

import (
	"context"

	"github.com/entegral/toolbox/dynamo"
)

type Venue struct {
	dynamo.DiLink[*Event, *Office]
	Room         *string `json:"room,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
}

func (v *Venue) Type() string {
	return "venue"
}

func (v *Venue) Office(ctx context.Context) (*Office, error) {
	_, err := v.LoadEntity1(ctx)
	if err != nil {
		return nil, err
	}
	return v.Entity1, nil
}

func (v *Venue) Events(ctx context.Context) (*Event, error) {
	_, err := v.LoadEntity0(ctx)
	if err != nil {
		return nil, err
	}
	return v.Entity0, nil
}

func (v *Venue) Link(ctx context.Context) error {
	return v.Put(ctx, v)
}

type NewVenueOpts struct {
	Room         *string
	Instructions *string
}

// NewVenue creates a new venue.
func NewVenue(ctx context.Context, event Event, office Office, opts *NewVenueOpts) (*Venue, error) {
	link, err := dynamo.CheckDiLink[*Event, *Office](&event, &office)
	venue := &Venue{
		DiLink: *link,
	}
	if opts != nil {
		venue.Room = opts.Room
		venue.Instructions = opts.Instructions
	}
	switch err.(type) {
	case nil:
		return venue, nil
	case dynamo.ErrLinkNotFound:
		_, _, err := venue.LoadEntities(ctx)
		return venue, err
	default:
		return nil, err
	}
}
