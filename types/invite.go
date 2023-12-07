package types

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/entegral/toolbox/dynamo"
)

// Invite is a struct representing an invitation.
type Invite struct {
	dynamo.DiLink[*Event, *User]              // DiLink is a struct from the dynamo package that links two entities.
	Status                       InviteStatus `json:"status"` // Status represents the status of the invitation.
}

// Type method returns the type of the struct, in this case "invite".
func (i Invite) Type() string {
	return "invite"
}

// User method loads the user entity associated with the invite.
func (i *Invite) User(ctx context.Context) (*User, error) {
	_, err := i.LoadEntity1(ctx)
	if err != nil {
		return nil, err
	}
	return i.Entity1, nil
}

// Event method loads the event entity associated with the invite.
func (i *Invite) Event(ctx context.Context) (*Event, error) {
	_, err := i.LoadEntity0(ctx)
	if err != nil {
		return nil, err
	}
	return i.Entity0, nil
}

// NewInviteOpts is a struct that holds options for creating a new invite.
type NewInviteOpts struct {
	Status *InviteStatus // Status is a pointer to an InviteStatus value.
}

// NewInvite is a function that creates a new invite. It requires an event, a user, and options for the invite.
func NewInvite(ctx context.Context, event Event, user User, opts *NewInviteOpts) (*Invite, error) {
	if event.CreatedByEmail == "" {
		return nil, fmt.Errorf("event must have a CreatedByEmail")
	}
	dilink, err := dynamo.CheckLink[*Event, *User](&event, &user)
	invite := &Invite{
		DiLink: *dilink,
	}
	if opts != nil && opts.Status != nil {
		invite.Status = *opts.Status
	}
	_, _, err = invite.LoadEntities(ctx)
	return invite, err
}

// InviteStatus is a type representing the status of an invitation.
type InviteStatus string

// Constants representing possible statuses of an invitation.
const (
	InviteStatusAccepted InviteStatus = "ACCEPTED"
	InviteStatusDeclined InviteStatus = "DECLINED"
	InviteStatusPending  InviteStatus = "PENDING"
)

// AllInviteStatus is a slice containing all possible statuses of an invitation.
var AllInviteStatus = []InviteStatus{
	InviteStatusAccepted,
	InviteStatusDeclined,
	InviteStatusPending,
}

// IsValid method checks if the InviteStatus value is valid.
func (e InviteStatus) IsValid() bool {
	switch e {
	case InviteStatusAccepted, InviteStatusDeclined, InviteStatusPending:
		return true
	}
	return false
}

// String method returns the InviteStatus value as a string.
func (e InviteStatus) String() string {
	return string(e)
}

// UnmarshalGQL method is used to unmarshal a GraphQL value into an InviteStatus value.
func (e *InviteStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InviteStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InviteStatus", str)
	}
	return nil
}

// MarshalGQL method is used to marshal an InviteStatus value into a GraphQL value.
func (e InviteStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
