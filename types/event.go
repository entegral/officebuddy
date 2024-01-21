package types

import (
	"context"
	"fmt"

	"github.com/dgryski/trifles/uuid"
	"github.com/entegral/gobox/dynamo"
	"github.com/entegral/gobox/types"
)

// Event struct represents an event with necessary fields.
type Event struct {
	dynamo.Row                // Embedding dynamo.Row to inherit its fields and methods.
	GUID       string         `json:"guid"`                 // The unique identifier for the event.
	Start      types.DateTime `json:"start"`                // The start time of the event.
	End        types.DateTime `json:"end"`                  // The end time of the event.
	Details    *EventDetails  `json:"details" dynamoav:"-"` // The details of the event. Should be marshalled to json but not to dynamodb.
}

// EventInput struct is used for creating a new event.
type EventInput struct {
	Event // Embedding Event to inherit its fields and methods.
	EventDetails
}

// Type method returns the type of the struct, in this case "event".
func (e *Event) Type() string {
	return "event"
}

// Keys method generates and returns the primary and secondary keys for the event based on the given Global Secondary Index (gsi).
func (e *Event) Keys(gsi int) (string, string, error) {
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
		return e.Pk, e.Sk, nil
	case 1:
		return e.Pk1, e.Sk1, nil
	case 2:
		return e.Pk2, e.Sk2, nil
	}
	return "", "", nil
}

// NewEventOpts struct is used to pass options when creating a new event.
type NewEventOpts struct {
	GUID *string
}

// NewEvent function creates a new event and returns it. If createdByEmail is empty, it returns an error.
func NewEvent(ctx context.Context, title, description string, start, stop types.DateTime, opts *NewEventOpts) (*Event, error) {
	event := &Event{
		GUID: uuid.UUIDv4(), // Generate a new UUID for the event.
		// CreatedByEmail: createdByEmail,
	}
	if opts.GUID != nil && *opts.GUID != "" {
		event.GUID = *opts.GUID // If a GUID is provided in the options, use it instead of the generated one.
	}
	if start.IsZero() {
		return nil, fmt.Errorf("event start time must not be zero")
	}
	if stop.IsZero() {
		return nil, fmt.Errorf("event stop time must not be zero")
	}
	if opts == nil {
		return event, nil
	}
	details := &EventDetails{}
	if title == "" {
		return nil, fmt.Errorf("event title must not be empty")
	}
	details.Title = title
	if description == "" {
		return nil, fmt.Errorf("event description must not be empty")
	}
	details.Description = description
	event.Details = details
	return event, nil
}

// EventDetails struct represents the details of an event.
type EventDetails struct {
	dynamo.MonoLink[*Event]
	Title       string `json:"title"`       // The title of the event.
	Description string `json:"description"` // The description of the event.
}

// Type method satisfies the Typeable interface.
func (ed *EventDetails) Type() string {
	return "EventDetails"
}
