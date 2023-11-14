package types

import (
	"context"

	"github.com/entegral/toolbox/clients"
	"github.com/entegral/toolbox/dynamo"
)

type UserSaver struct {
	*User
}

func (u *UserSaver) Put(ctx context.Context, clients clients.Client) ([]*User, error) {
	_, err := dynamo.Put(ctx, clients, u)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
