package types

import (
	"context"
	"fmt"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/types"
)

// Event is a type for defining an event
type Event struct {
	dynamo.Row
	CreatedByEmail string         `json:"createdByEmail"`
	GUID           string         `json:"guid"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Start          types.DateTime `json:"start"`
	End            types.DateTime `json:"end"`
}

type EventInput struct {
	Event
}

func (e *Event) Type() string {
	return "event"
}

func (e *Event) Keys(gsi int) (string, string) {
	if e.GUID == "" {
		e.GUID = uuid.UUIDv4()
	}
	e.Pk = "event:" + e.GUID
	e.Sk = "info"
	e.Pk1 = "event:" + e.GUID
	e.Sk1 = "start:" + e.Start.String()
	e.Pk2 = e.Pk1
	e.Sk2 = "end:" + e.End.String()
	switch gsi {
	case 0:
		return e.Pk, e.Sk
	case 1:
		return e.Pk1, e.Sk1
	case 2:
		return e.Pk2, e.Sk2
	}
	return "", ""
}

type NewEventOpts struct {
	GUID        *string
	Title       *string
	Description *string
	Start       *types.DateTime
	End         *types.DateTime
}

// NewEvent creates a new event.
func NewEvent(ctx context.Context, createdByEmail string, opts *NewEventOpts) (*Event, error) {
	if createdByEmail == "" {
		return nil, fmt.Errorf("createdByEmail must not be empty")
	}
	event := &Event{
		GUID:           uuid.UUIDv4(),
		CreatedByEmail: createdByEmail,
	}
	if opts.GUID != nil && *opts.GUID != "" {
		event.GUID = *opts.GUID
	}
	if opts.Title != nil && *opts.Title != "" {
		event.Title = *opts.Title
	}
	if opts.Description != nil && *opts.Description != "" {
		event.Description = *opts.Description
	}
	if opts.Start != nil {
		event.Start = *opts.Start
	}
	if opts.End != nil {
		event.End = *opts.End
	}

	return event, nil
}
