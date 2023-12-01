package types

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awstypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/entegral/toolbox/dynamo"
	"github.com/entegral/toolbox/helpers"
)

// Schedule is a type for defining a schedule for a given office.
type Schedule struct {
	dynamo.Row
	ScheduleGUID string    `json:"scheduleGUID"`
	OfficeGUID   string    `json:"officeGUID"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Active       bool      `json:"active"`
	Coworkers    []*User   `json:"coworkers,omitempty"`
}

func (s *Schedule) Type() string {
	return "schedule"
}

type ScheduleSaver struct {
	Schedule
}

type ScheduleFinder struct {
	Schedule
}

func (s *Schedule) LoadSchedule(ctx context.Context) (loaded bool, err error) {
	if s.ScheduleGUID == "" && s.OfficeGUID == "" {
		return false, nil
	}
	if s.ScheduleGUID != "" && s.OfficeGUID != "" {
		return helpers.GetItem(ctx, s)
	}
	if s.OfficeGUID != "" {
		schedules, err := s.LoadSchedulesForOffice(ctx)
		if err != nil {
			return false, err
		}
		if len(schedules) != 0 {
			*s = schedules[0]
			return true, nil
		}
	}
	if s.ScheduleGUID != "" {
		schedules, err := s.LoadScheduleAssignments(ctx)
		if err != nil {
			return false, err
		}
		if len(schedules) != 0 {
			*s = schedules[0]
			return true, nil
		}
	}
	return false, nil
}

// LoadSchedulesForOffice loads all schedules for a given office
func (s *Schedule) LoadSchedulesForOffice(ctx context.Context) ([]Schedule, error) {
	if s.ScheduleGUID == "" {
		return nil, fmt.Errorf("office GUID is required to load schedules for office")
	}
	pk, _ := s.Keys(0)
	tablename := s.TableName(ctx)
	exp := "pk = :pk"
	i := dynamodb.QueryInput{
		TableName:              &tablename,
		KeyConditionExpression: &exp,
		ExpressionAttributeValues: map[string]awstypes.AttributeValue{
			":pk1": &awstypes.AttributeValueMemberS{
				Value: pk,
			},
		},
	}
	schedules, err := helpers.Query[Schedule](ctx, i)
	if err != nil {
		return nil, err
	}
	if len(schedules) == 0 {
		return nil, nil
	}
	return schedules, nil
}

// LoadScheduleAssignments loads all office assignments for a given schedule
func (s *Schedule) LoadScheduleAssignments(ctx context.Context) ([]Schedule, error) {
	if s.OfficeGUID == "" {
		return nil, fmt.Errorf("office GUID is required to load schedules for office")
	}
	tablename := s.TableName(ctx)
	gsi := helpers.GSI1.String()
	exp := "pk1 = :pk1"
	pk1, _ := s.Keys(1)
	i := dynamodb.QueryInput{
		TableName:              &tablename,
		IndexName:              &gsi,
		KeyConditionExpression: &exp,
		ExpressionAttributeValues: map[string]awstypes.AttributeValue{
			":pk1": &awstypes.AttributeValueMemberS{
				Value: pk1,
			},
		},
	}
	schedules, err := helpers.Query[Schedule](ctx, i)
	if err != nil {
		return nil, err
	}
	if len(schedules) == 0 {
		return nil, nil
	}
	return schedules, nil
}

// Keys returns the partition key and sort key for the row
func (s *Schedule) Keys(gsi int) (string, string) {
	// For this example, assuming GUID is the partition key and Email is the sort key.
	// Additional logic can be added to handle different GSIs if necessary.
	s.Pk = "office:" + s.OfficeGUID
	s.Sk = "schedule:" + s.ScheduleGUID
	s.Pk1 = s.Sk
	s.Sk1 = s.Pk
	switch gsi {
	case 0: // Primary keys
		return s.Pk, s.Sk
	case 1: // GSI1
		return s.Pk1, s.Sk1
	default:
		// Handle other GSIs or return an error
		return "", ""
	}
}
