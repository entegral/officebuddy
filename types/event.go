package types

import (
	"context"
	"fmt"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/types"
)

// Event struct represents an event with necessary fields.
type Event struct {
	dynamo.Row                    // Embedding dynamo.Row to inherit its fields and methods.
	CreatedByEmail string         `json:"createdByEmail"` // The email of the user who created the event.
	GUID           string         `json:"guid"`           // The unique identifier for the event.
	Title          string         `json:"title"`          // The title of the event.
	Description    string         `json:"description"`    // The description of the event.
	Start          types.DateTime `json:"start"`          // The start time of the event.
	End            types.DateTime `json:"end"`            // The end time of the event.
}

// EventInput struct is used for creating a new event.
type EventInput struct {
	Event // Embedding Event to inherit its fields and methods.
}

// Type method returns the type of the struct, in this case "event".
func (e *Event) Type() string {
	return "event"
}

// Keys method generates and returns the primary and secondary keys for the event based on the given Global Secondary Index (gsi).
func (e *Event) Keys(gsi int) (string, string) {
	if e.GUID == "" {
		e.GUID = uuid.UUIDv4() // Generate a new UUID if it doesn't exist.
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

// NewEventOpts struct is used to pass options when creating a new event.
type NewEventOpts struct {
	GUID        *string
	Title       *string
	Description *string
	Start       *types.DateTime
	End         *types.DateTime
}

// NewEvent function creates a new event and returns it. If createdByEmail is empty, it returns an error.
func NewEvent(ctx context.Context, createdByEmail string, opts *NewEventOpts) (*Event, error) {
	if createdByEmail == "" {
		return nil, fmt.Errorf("createdByEmail must not be empty")
	}
	event := &Event{
		GUID:           uuid.UUIDv4(), // Generate a new UUID for the event.
		CreatedByEmail: createdByEmail,
	}
	if opts.GUID != nil && *opts.GUID != "" {
		event.GUID = *opts.GUID // If a GUID is provided in the options, use it instead of the generated one.
	}
	if opts.Title != nil && *opts.Title != "" {
		event.Title = *opts.Title // If a title is provided in the options, use it.
	}
	if opts.Description != nil && *opts.Description != "" {
		event.Description = *opts.Description // If a description is provided in the options, use it.
	}
	if opts.Start != nil {
		event.Start = *opts.Start // If a start time is provided in the options, use it.
	}
	if opts.End != nil {
		event.End = *opts.End // If an end time is provided in the options, use it.
	}

	return event, nil
}
