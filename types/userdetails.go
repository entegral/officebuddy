package types

import (
	"github.com/entegral/gobox/dynamo"
)

// UserDetails contains some profile-related info about a user
type UserDetails struct {
	dynamo.MonoLink[*User]
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// Type returns the type of the entity.
func (u *UserDetails) Type() string {
	return "UserDetails"
}
