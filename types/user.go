package types

import (
	"context"
	"errors"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/gobox/dynamo"
	"github.com/sirupsen/logrus"
)

// User is a user.
type User struct {
	dynamo.Row
	GUID  string `json:"guid,omitempty"`
	Email string `json:"email,omitempty"`
	// AdminOf   []*Office `json:"adminOf,omitempty"`

	// MonoLinks
	Details *UserDetails `json:"details,omitempty" dynamoav:"-"`
}

// Type returns the type of the entity.
func (u *User) Type() string {
	return "User"
}

// Keys returns the partition key and sort key for the given GSI.
func (u *User) Keys(gsi int) (partitionKey, sortKey string, err error) {
	u.Pk = "email:" + u.Email
	u.Sk = "userinfo"
	u.Pk1 = "guid:" + u.GUID
	u.Sk1 = u.Sk
	switch gsi {
	case 0:
		return u.Pk, u.Sk, nil
	case 1:
		return u.Pk1, u.Sk1, nil
	default:
		return "", "", nil
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
	user := &User{
		Email: email,
		GUID:  guid,
	}
	details := &UserDetails{
		MonoLink:  *dynamo.NewMonoLink(user),
		FirstName: fname,
		LastName:  lname,
	}
	user.Details = details
	return user, nil
}

// GetDetails returns the user details, if it isnt present, it loads them from dynamo.
func (u *User) GetDetails(ctx context.Context) (*UserDetails, error) {
	if u.Details != nil {
		return u.Details, nil
	}
	u.Details = &UserDetails{
		MonoLink: *dynamo.NewMonoLink(u),
	}
	loaded, err := u.Details.Get(ctx, u.Details)
	if err != nil {
		logrus.WithError(err).Error("failed to get user cache")
	}
	if !loaded {
		return nil, err
	}
	return u.Details, nil
}

// Memberships returns the office memberships for the user.
func (u *User) Memberships(ctx context.Context, roles []Role) ([]*Membership, error) {
	memberships, err := dynamo.FindLinksByEntity0[*User, *Membership](ctx, u, "Membership")
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
	invites, err := dynamo.FindLinksByEntity1[*User, *Invite](
		ctx,
		u,
		"Invite",
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
