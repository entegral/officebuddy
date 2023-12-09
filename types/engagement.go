package types

import (
	"context"

	"github.com/entegral/toolbox/dynamo"
)

// Engagement is a type that allows you to link a user to an office and an event.
type Engagement struct {
	dynamo.TriLink[*User, *Office, *Event]
}

// User is the resolver for the user field.
func (e *Engagement) User(ctx context.Context) (*User, error) {
	loaded, err := e.LoadEntity0(ctx)
	if err != nil {
		return nil, err
	}
	if loaded {
		return e.Entity0, nil
	}
	return nil, nil
}

// Office is the resolver for the office field.
func (e *Engagement) Office(ctx context.Context) (*Office, error) {
	loaded, err := e.LoadEntity1(ctx)
	if err != nil {
		return nil, err
	}
	if loaded {
		return e.Entity1, nil
	}
	return nil, nil
}

// Event is the resolver for the event field.
func (e *Engagement) Event(ctx context.Context) (*Event, error) {
	loaded, err := e.LoadEntity2(ctx)
	if err != nil {
		return nil, err
	}
	if loaded {
		return e.Entity2, nil
	}
	return nil, nil
}
