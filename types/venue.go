package types

import (
	"context"

	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

type Venue struct {
	dynamo.DiLink[*Event, *Office]
	Room         *string `json:"room,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
}

func (v *Venue) Type() string {
	return "venue"
}

func (v *Venue) Link(ctx context.Context, clients clients.Client) error {
	_, err := helpers.PutItem(ctx, v)
	return err
}

type NewVenueOpts struct {
	Room         *string
	Instructions *string
}

// NewVenue creates a new venue.
func NewVenue(ctx context.Context, event Event, office Office, opts *NewVenueOpts) (*Venue, error) {
	link, err := dynamo.CheckLink[*Event, *Office](&event, &office, dynamo.OneToMany)
	venue := &Venue{
		DiLink: *link,
	}
	switch err.(type) {
	case nil:
		return venue, nil
	case dynamo.ErrLinkNotFound:
		if opts != nil {
			venue.Room = opts.Room
			venue.Instructions = opts.Instructions
		}
		_, _, err := venue.LoadEntities(ctx, *clients.GetDefaultClient(ctx))
		return venue, err
	default:
		return nil, err
	}
}
