package types

import (
	"context"
	"errors"

	"log/slog"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awstypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/entegral/gobox/clients"
	"github.com/entegral/gobox/dynamo"
)

type UserFinder struct {
	User
}

// Load attempts to load the user using data from any of the fields.
// func (u *UserFinder) Load(ctx context.Context) error {
// 	var user *User
// 	if u.GUID != "" {

// LoadByGUID fetches the latest version of a user by GUID.
func (u *UserFinder) LoadByGUID(ctx context.Context) (bool, error) {
	tn := u.TableName(ctx)
	in := dynamo.GSI1.String()
	pk1, sk1, err := u.Keys(1)
	if err != nil {
		return false, err
	}
	kce := "pk1 = :pk1 and sk1 = :sk1"
	i := dynamodb.QueryInput{
		TableName:              &tn,
		IndexName:              &in,
		KeyConditionExpression: &kce,
		ExpressionAttributeValues: map[string]awstypes.AttributeValue{
			":pk1": &awstypes.AttributeValueMemberS{Value: pk1},
			":sk1": &awstypes.AttributeValueMemberS{Value: sk1},
		},
	}
	client := clients.GetDefaultClient(ctx)
	out, err := client.Dynamo().Query(ctx, &i)
	if err != nil {
		return false, err
	}
	if out.Count == 0 {
		return false, nil
	}
	users := make([]User, len(out.Items))
	err = attributevalue.UnmarshalListOfMaps(out.Items, &users)
	if err != nil {
		return false, err
	}
	if len(users) > 1 {
		slog.With("users", users).Error("more than one user found, returning the first")
		return true, errors.New("more than one user found")
	}
	u.User = users[0]
	return true, err
}

// LoadByEmail fetches the latest version of a user by email.
func (u *UserFinder) LoadByEmail(ctx context.Context) (bool, error) {
	return u.Get(ctx, &u.User)
}
