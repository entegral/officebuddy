package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/helpers"
)

// Schedules is the resolver for the schedules field.
func (r *officeResolver) Schedules(ctx context.Context, obj *types.Office, scheduleGUID *string) ([]*types.Schedule, error) {
	schedule := types.Schedule{
		OfficeGUID: obj.GUID,
	}
	if scheduleGUID != nil && *scheduleGUID != "" {
		schedule.ScheduleGUID = *scheduleGUID
		loaded, err := helpers.GetItem(ctx, &schedule)
		if err != nil {
			return nil, err
		}
		if !loaded {
			return nil, nil
		}
		return []*types.Schedule{&schedule}, nil
	}
	pk, _ := schedule.Keys(0)
	schedules, err := helpers.QueryByGSI[*types.Schedule](ctx, 0, pk, "schedule:")
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// Office is the resolver for the office field.
func (r *queryResolver) Office(ctx context.Context, officeGUID string) (*types.Office, error) {
	office := types.Office{
		GUID: officeGUID,
	}
	loaded, err := helpers.GetItem(ctx, &office)
	if err != nil {
		return nil, err
	}
	if !loaded {
		return nil, nil
	}
	return &office, nil
}

// Offices is the resolver for the offices field.
func (r *queryResolver) Offices(ctx context.Context, userGUID string) ([]*types.Office, error) {
	// TODO use GSI0 to query the office memberships for the user
	// TODO Use that list to query the offices' info
	return nil, nil
}

// Office returns OfficeResolver implementation.
func (r *Resolver) Office() OfficeResolver { return &officeResolver{r} }

type officeResolver struct{ *Resolver }
