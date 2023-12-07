package types

import (
	"context"
	"errors"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/toolbox/dynamo"
)

type User struct {
	dynamo.Row
	GUID      string `json:"guid,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	// AdminOf   []*Office `json:"adminOf,omitempty"`
}

func (u *User) Type() string {
	return "user"
}

// Keys returns the partition key and sort key for the given GSI.
func (u *User) Keys(gsi int) (partitionKey, sortKey string) {
	u.Pk = "user:" + u.Email
	u.Sk = "userinfo"
	u.Pk1 = "user:" + u.GUID
	u.Sk1 = u.Sk
	switch gsi {
	case 0:
		return u.Pk, u.Sk
	case 1:
		return u.Pk1, u.Sk1
	default:
		return "", ""
	}
}

// ErrInvalidEmail is returned when an email is invalid.
var ErrInvalidEmail = errors.New("invalid email")

// ErrInvalidFirstName is returned when a first name is invalid.
var ErrInvalidFirstName = errors.New("invalid first name")

// ErrInvalidLastName is returned when a last name is invalid.
var ErrInvalidLastName = errors.New("invalid last name")

// NewUser creates a new user.
func NewUser(ctx context.Context, guid, email, fname, lname string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if fname == "" {
		return nil, ErrInvalidFirstName
	}
	if lname == "" {
		return nil, ErrInvalidLastName
	}
	if guid == "" {
		guid = uuid.UUIDv4()
	}
	return &User{
		Email:     email,
		FirstName: fname,
		LastName:  lname,
		GUID:      guid,
	}, nil
}

// Memberships returns the office memberships for the user.
func (u *User) Memberships(ctx context.Context, roles []Role) ([]*Membership, error) {
	memberships, err := dynamo.FindCustomLinksByEntity0[*User, *Membership](ctx, u)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return memberships, nil
	}
	withRoles := []*Membership{}
	if len(roles) == 0 {
		return memberships, nil
	}
	for _, m := range memberships {
		for _, r := range roles {
			if m.Role == r {
				withRoles = append(withRoles, m)
			}
		}
	}
	return withRoles, nil
}

// Invites returns the invites for the user.
func (u *User) Invites(ctx context.Context, status []InviteStatus) ([]*Invite, error) {
	invites, err := dynamo.FindCustomLinksByEntity1[*User, *Invite](
		ctx,
		u,
	)
	if err != nil {
		return nil, err
	}
	if len(status) == 0 {
		return invites, nil
	}
	withStatus := []*Invite{}
	for _, i := range invites {
		for _, s := range status {
			if i.Status == s {
				withStatus = append(withStatus, i)
			}
		}
	}
	return withStatus, nil
}
