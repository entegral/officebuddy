package types

import (
	"context"
	"time"

	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

// Membership is a type for defining a membership of a user to an office.
type Membership struct {
	dynamo.DiLink[*User, *Office]
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

func (m *Membership) Type() string {
	return "membership"
}

// Link we will override the default Link method to add the role and createdAt fields.
func (m *Membership) Link(ctx context.Context, clients clients.Client) error {
	_, err := helpers.PutItem(ctx, m)
	return err
}

type NewMembershipErrors string

const (
	// ErrUserNotFound is the error for when a user is not found.
	ErrUserNotFound NewMembershipErrors = "user not found"
	// ErrorOfficeNotFound is the error for when an office is not found.
	ErrorOfficeNotFound NewMembershipErrors = "office not found"
)

// NewMembershipError is the error type for when a new membership cannot be created.
type NewMembershipError struct {
	message NewMembershipErrors
}

func (e NewMembershipError) Error() string {
	return string(e.message)
}

// User loads the user for the membership.
func (m *Membership) User(ctx context.Context) (*User, error) {
	loaded, err := m.LoadEntity0(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, NewMembershipError{ErrUserNotFound}
	}
	return m.Entity0, nil
}

// Office loads the office for the membership.
func (m *Membership) Office(ctx context.Context) (*Office, error) {
	loaded, err := m.LoadEntity1(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, NewMembershipError{ErrorOfficeNotFound}
	}
	return m.Entity1, nil
}

// LoadMembership simplifies the loading of a membership from dynamo.
// If either the user or office is not found, an error is returned.
func LoadMembership(email, officeGUID string, role Role) (*Membership, error) {
	dilink, newErr := dynamo.CheckLink[*User, *Office](&User{Email: email}, &Office{GUID: officeGUID}, dynamo.OneToMany)
	membership := &Membership{
		DiLink:    *dilink,
		Role:      role,
		CreatedAt: time.Now(),
	}
	return membership, newErr
}

// Role is a type for defining the role of a user in an office.
type Role string

// String returns the string representation of the role.
func (r Role) String() string {
	return string(r)
}

const (
	// RoleAdmin is the admin role.
	RoleAdmin Role = "admin"
	// RoleMember is the member role.
	RoleMember Role = "member"
)
