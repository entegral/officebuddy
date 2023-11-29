package types

import (
	"context"

	"github.com/entegral/toolbox/dynamo"
)

// Office is a type for defining an office.
type Office struct {
	dynamo.Row
	GUID        string   `json:"guid"`
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

// Events returns the events for the office
func (o Office) Events(ctx context.Context) ([]*Event, error) {
	// load the events for the office by querying the office_events GSI
	return nil, nil
}

// Members returns the members for the office
func (o Office) Members(ctx context.Context) ([]*User, error) {
	// load the members for the office by querying the office_members GSI
	return nil, nil
}

// Admins returns the admins for the office
func (o Office) Admins(ctx context.Context) ([]*User, error) {
	// load the admins for the office by querying the office_admins GSI
	return nil, nil
}

// Address is a type for defining an address.
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
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
