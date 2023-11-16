package types

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awstypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/entegral/toolbox/helpers"
	"github.com/entegral/toolbox/types"
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
	in := types.GSI1.String()
	pk1, sk1 := u.Keys(1)
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
	users, err := helpers.Query[User](ctx, i)
	if err != nil {
		return false, err
	}
	if len(users) == 0 {
		return false, nil
	}
	u.User = users[0]
	return true, err
}

// LoadByEmail fetches the latest version of a user by email.
func (u *UserFinder) LoadByEmail(ctx context.Context) (bool, error) {
	return helpers.GetItem(ctx, &u.User)
}
