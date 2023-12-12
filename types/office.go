package types

import (
	"context"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/gobox/dynamo"
)

// Office is a struct representing an office.
type Office struct {
	dynamo.Row           // Row is a struct from the dynamo package that represents a row in a table.
	GUID        string   `json:"guid"`                  // GUID is the unique identifier for the office.
	Name        string   `json:"name"`                  // Name is the name of the office.
	CreatedBy   string   `json:"createdBy"`             // CreatedBy is the identifier of the entity that created the office.
	Description *string  `json:"description,omitempty"` // Description is an optional description of the office.
	Address     *Address `json:"address,omitempty"`     // Address is an optional address of the office.
}

// Type method returns the type of the struct, in this case "office".
func (o *Office) Type() string {
	return "office"
}

// NewOffice is a function that creates a new office. It requires a name, a createdBy identifier, a guid, a description, and an address.
func NewOffice(ctx context.Context, name, createdBy string, guid, description *string, address Address) Office {
	o := Office{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		Address:     &address,
		GUID:        uuid.UUIDv4(), // Generate a new UUID for the office.
	}
	if guid != nil && *guid != "" {
		o.GUID = *guid // If a GUID is provided, use it instead of the generated one.
	}
	return o
}

// Memberships method returns the memberships associated with the office.
func (o Office) Memberships(ctx context.Context) ([]*Membership, error) {
	return dynamo.FindCustomLinksByEntity1[*Office, *Membership](ctx, &o)
}

// Venue method returns the venues associated with the office.
func (o Office) Venue(ctx context.Context) ([]*Venue, error) {
	venues, err := dynamo.FindCustomLinksByEntity1[*Office, *Venue](
		ctx,
		&o,
	)
	if err != nil {
		return nil, err
	}
	return venues, nil
}

// AddressInput is a struct representing an input for an address.
type AddressInput struct {
	Address // Address is a struct representing an address.
}

// Address is a struct representing an address.
type Address struct {
	Street  string `json:"street"`  // Street is the street of the address.
	City    string `json:"city"`    // City is the city of the address.
	State   string `json:"state"`   // State is the state of the address.
	Zip     string `json:"zip"`     // Zip is the zip code of the address.
	Country string `json:"country"` // Country is the country of the address.
}

// Keys method returns the primary key and the sort key of the office.
func (o Office) Keys(gsi int) (string, string) {
	o.Pk = "guid:" + o.GUID // Pk is the primary key of the office.
	o.Sk = "info"           // Sk is the sort key of the office.
	switch gsi {
	case 0:
		return o.Pk, o.Sk
	default:
		return "", "" // If the gsi is not 0, return empty strings.
	}
}
