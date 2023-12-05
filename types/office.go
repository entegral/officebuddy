package types

import (
	"context"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
)

// Office is a type for defining an office.
type Office struct {
	dynamo.Row
	GUID        string   `json:"guid"`
	Name        string   `json:"name"`
	CreatedBy   string   `json:"createdBy"`
	Description *string  `json:"description,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

func (o *Office) Type() string {
	return "office"
}

// NewOffice simplifies the creation of a new office.
func NewOffice(ctx context.Context, name, createdBy string, guid, description *string, address Address) Office {
	o := Office{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		Address:     &address,
		GUID:        uuid.UUIDv4(),
	}
	if guid != nil && *guid != "" {
		o.GUID = *guid
	}
	return o
}

// Events returns the events for the office
func (o Office) Events(ctx context.Context) ([]*Event, error) {
	// load the events for the office by querying the office_events GSI
	return nil, nil
}

// Memberships returns the members for the office
func (o Office) Memberships(ctx context.Context) ([]*Membership, error) {
	return dynamo.FindCustomLinksByEntity1[*User, *Office, *Membership](ctx, *clients.GetDefaultClient(ctx), &o)
}

// Venue returns the venue for an office
func (o Office) Venue(ctx context.Context, clients clients.Client) ([]*Venue, error) {
	venues, err := dynamo.FindCustomLinksByEntity1[*Event, *Office, *Venue](
		ctx,
		clients,
		&o,
	)
	if err != nil {
		return nil, err
	}
	return venues, nil
}

type AddressInput struct {
	Address
}

// Address is a type for defining an address.
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// Keys returns the partition key and sort key for the row
func (o Office) Keys(gsi int) (string, string) {
	// For this example, assuming GUID is the partition key and Email is the sort key.
	// Additional logic can be added to handle different GSIs if necessary.

	o.Pk = "office:" + o.GUID
	o.Sk = "info"
	switch gsi {
	case 0: // Primary keys
		return o.Pk, o.Sk
	default:
		// Handle other GSIs or return an error
		return "", ""
	}
}
