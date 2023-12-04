package types

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

type Invite struct {
	dynamo.DiLink[*Event, *User]
	Status InviteStatus `json:"status"`
}

func (i Invite) Type() string {
	return "invite"
}

func (i *Invite) Link(ctx context.Context, clients clients.Client) error {
	_, err := helpers.PutItem(ctx, i)
	return err
}

// NewInviteOpts is a type for defining options for creating a new invite.
type NewInviteOpts struct {
	Status *InviteStatus
}

// NewInvite creates a new invite.
func NewInvite(ctx context.Context, event Event, user User, opts *NewInviteOpts) (*Invite, error) {
	if event.CreatedByEmail == "" {
		return nil, fmt.Errorf("event must have a CreatedByEmail")
	}
	invite := &Invite{
		DiLink: dynamo.DiLink[*Event, *User]{
			Entity0:  &event,
			Entity1:  &user,
			Relation: dynamo.OneToMany,
		},
	}
	if opts != nil && opts.Status != nil {
		invite.Status = *opts.Status
	}
	_, _, err := invite.LoadEntities(ctx, *clients.GetDefaultClient(ctx))
	return invite, err
}

type InviteStatus string

const (
	InviteStatusAccepted InviteStatus = "ACCEPTED"
	InviteStatusDeclined InviteStatus = "DECLINED"
	InviteStatusPending  InviteStatus = "PENDING"
)

var AllInviteStatus = []InviteStatus{
	InviteStatusAccepted,
	InviteStatusDeclined,
	InviteStatusPending,
}

func (e InviteStatus) IsValid() bool {
	switch e {
	case InviteStatusAccepted, InviteStatusDeclined, InviteStatusPending:
		return true
	}
	return false
}

func (e InviteStatus) String() string {
	return string(e)
}

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

func (e InviteStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
