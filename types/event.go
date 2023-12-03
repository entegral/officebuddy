package types

import (
	"context"

	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
	"github.com/entegral/toolbox/types"
)

// Event is a type for defining an event
type Event struct {
	dynamo.DiLink[*User, *Office]
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Start       types.DateTime `json:"start"`
	End         types.DateTime `json:"end"`
	Address     Address        `json:"address"`
}

func (e *Event) Type() string {
	return "event"
}

// Link we will override the default Link method to add the extra fields to the row.
func (e *Event) Link(ctx context.Context, clients clients.Client) error {
	_, err := helpers.PutItem(ctx, e)
	return err
}

// User loads the user for the event.
func (e *Event) User(ctx context.Context) (*User, error) {
	loaded, err := e.LoadEntity0(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, dynamo.ErrEntityNotFound[*User]{Entity: e.Entity0}
	}
	return e.Entity0, nil
}

// Office loads the office for the event.
func (e *Event) Office(ctx context.Context) (*Office, error) {
	loaded, err := e.LoadEntity1(ctx, *clients.GetDefaultClient(ctx))
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, dynamo.ErrEntityNotFound[*User]{Entity: e.Entity0}
	}
	return e.Entity1, nil
}

// NewEvent creates a new event.
func NewEvent(ctx context.Context, userEmail string, officeGUID string, title string, description string, start, end types.DateTime) (*Event, error) {
	link, err := dynamo.CheckLink[*User, *Office](&User{Email: userEmail}, &Office{GUID: officeGUID}, dynamo.OneToMany)
	switch err.(type) {
	case nil, dynamo.ErrLinkNotFound:
		event := &Event{
			DiLink:      *link,
			Title:       title,
			Description: description,
			Start:       start,
			End:         end,
		}
		err = event.Link(ctx, *clients.GetDefaultClient(ctx))
		if err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, err
	}
}
