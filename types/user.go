package types

import (
	"context"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/toolbox/clients"
)

type User struct {
	GUID      string `json:"guid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	// AdminOf   []*Office `json:"adminOf,omitempty"`
}

// These satisfy the Keyable interface which is used in both Put and Update calls of dynamo

// Pk is used to satisfy the Keyable interface
func (u *User) Pk(ctx context.Context, clients clients.Client) (string, error) {
	return "EML/" + u.Email, nil
}

// Sk is used to satisfy the Keyable interface
func (u *User) Sk(ctx context.Context, clients clients.Client) (string, error) {
	return "data/profileinfo", nil
}

// These satisfy the Hookable interface

// Before satisfies the Hookable interface
func (u *User) Before(ctx context.Context, clients clients.Client) error {
	if u.GUID == "" {
		u.GUID = uuid.UUIDv4()
	}
	return nil
}

// After satisfies the Hookable interface
func (u *User) After(ctx context.Context, clients clients.Client) error { return nil }

func (u *User) AdminOf(ctx context.Context, clients clients.Client) ([]*Office, error) {
	return nil, nil
}
