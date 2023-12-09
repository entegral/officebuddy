package types

import (
	"context"
	"time"

	"github.com/entegral/toolbox/dynamo"
)

// Membership is a struct representing a membership.
type Membership struct {
	dynamo.DiLink[*User, *Office]           // DiLink is a struct from the dynamo package that links two entities.
	Role                          Role      `json:"role"`      // Role represents the role of the member.
	CreatedAt                     time.Time `json:"createdAt"` // CreatedAt represents the time when the membership was created.
}

// Type method returns the type of the struct, in this case "membership".
func (m *Membership) Type() string {
	return "membership"
}

// NewMembershipErrors is a type representing errors that can occur when creating a new membership.
type NewMembershipErrors string

// Constants representing possible errors when creating a new membership.
const (
	ErrUserNotFound     NewMembershipErrors = "user not found"
	ErrorOfficeNotFound NewMembershipErrors = "office not found"
)

// NewMembershipError is a struct representing an error when creating a new membership.
type NewMembershipError struct {
	message NewMembershipErrors // Message is the error message.
}

// Error method returns the error message as a string.
func (e NewMembershipError) Error() string {
	return string(e.message)
}

// User method loads the user entity associated with the membership.
func (m *Membership) User(ctx context.Context) (*User, error) {
	loaded, err := m.LoadEntity0(ctx)
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, NewMembershipError{ErrUserNotFound}
	}
	return m.Entity0, nil
}

// Office method loads the office entity associated with the membership.
func (m *Membership) Office(ctx context.Context) (*Office, error) {
	loaded, err := m.LoadEntity1(ctx)
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, NewMembershipError{ErrorOfficeNotFound}
	}
	return m.Entity1, nil
}

// NewMembership is a function that creates a new membership. It requires an email, an office GUID, and a role.
func NewMembership(ctx context.Context, email, officeGUID string, role Role) (*Membership, error) {
	dilink, newErr := dynamo.CheckDiLink[*User, *Office](&User{Email: email}, &Office{GUID: officeGUID})
	membership := &Membership{
		DiLink:    *dilink,
		Role:      role,
		CreatedAt: time.Now(),
	}
	switch newErr.(type) {
	case nil:
		return membership, nil
	case dynamo.ErrLinkNotFound:
		_, _, err := membership.LoadEntities(ctx)
		return membership, err
	default:
		return nil, newErr
	}
}

// Role is a type representing the role of a member.
type Role string

// String method returns the Role value as a string.
func (r Role) String() string {
	return string(r)
}

// Constants representing possible roles of a member.
const (
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
)
