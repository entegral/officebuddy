package types

import (
	"context"
	"errors"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/gobox/dynamo"
	"github.com/sirupsen/logrus"
)

type UserCache struct {
	dynamo.MonoLink[*User]
	HairStyle string `json:"hairStyle,omitempty"`
}

// Type returns the type of the entity.
func (u *UserCache) Type() string {
	return "usercache"
}

// User is a user.
type User struct {
	dynamo.Row
	GUID      string `json:"guid,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	// AdminOf   []*Office `json:"adminOf,omitempty"`
}

// Type returns the type of the entity.
func (u *User) Type() string {
	return "user"
}

// Cache returns the cache for the user.
func (u *User) Cache(ctx context.Context) (*UserCache, error) {
	caches, err := dynamo.FindByEntity0[*User, *UserCache](ctx, u)
	if err != nil {
		return nil, err
	}
	if len(caches) == 0 {
		logrus.Warn("no cache found for user, creating")
		cache := &UserCache{
			MonoLink:  dynamo.NewMonoLink(u),
			HairStyle: "bald",
		}
		err = cache.Put(ctx, cache)
		if err != nil {
			return nil, err
		}
		return cache, nil
	}
	return caches[0], nil
}

// Keys returns the partition key and sort key for the given GSI.
func (u *User) Keys(gsi int) (partitionKey, sortKey string) {
	u.Pk = "email:" + u.Email
	u.Sk = "userinfo"
	u.Pk1 = "guid:" + u.GUID
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
	memberships, err := dynamo.FindByEntity0[*User, *Membership](ctx, u)
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
	invites, err := dynamo.FindByEntity1[*User, *Invite](
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
