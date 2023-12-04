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

func NewVenue(ctx context.Context, event *Event, office *Office, room, instructions *string) (*Venue, error) {
	venue := &Venue{
		Room:         room,
		Instructions: instructions,
		DiLink: dynamo.DiLink[*Event, *Office]{
			Entity0:  event,
			Entity1:  office,
			Relation: dynamo.OneToMany,
		},
	}
	_, _, err := venue.LoadEntities(ctx, *clients.GetDefaultClient(ctx))
	return venue, err
}
