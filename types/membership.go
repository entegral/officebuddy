package types

import (
	"context"
	"time"

	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

// Membership is a type for defining a membership of a user to an office.
type Membership struct {
	dynamo.DiLink[*User, *Office]
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

// // Disconnect links the membership to the user and office.
// func (m *Membership) Disconnect(ctx context.Context, clients clients.Client) error {
// 	_, err := helpers.PutItem(ctx, m)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

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

// NewMembership simplifies the creation of a new membership.
// If either the user or office is not found, an error is returned.
func NewMembership(email, officeGUID string, role Role, modifyKeys func(*dynamo.Row, *User, *Office) error) (*Membership, error) {
	membership := &Membership{
		DiLink: dynamo.DiLink[*User, *Office]{
			Entity0: &User{Email: email},
			Entity1: &Office{GUID: officeGUID},
		},
		Role:      role,
		CreatedAt: time.Now(),
	}
	if modifyKeys != nil {
		membership.ModifyKeys = modifyKeys
	}
	loaded0, err := helpers.GetItem(context.Background(), membership.Entity0)
	if err != nil {
		return nil, err
	}
	if !loaded0 {
		return nil, NewMembershipError{message: ErrUserNotFound}
	}
	loaded1, err := helpers.GetItem(context.Background(), membership.Entity1)
	if err != nil {
		return nil, err
	}
	if !loaded1 {
		return nil, NewMembershipError{message: ErrorOfficeNotFound}
	}
	return membership, nil
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

func a() {

}
