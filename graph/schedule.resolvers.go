package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/entegral/officebuddy/types"
	"github.com/entegral/toolbox/helpers"
)

// CreateSchedule is the resolver for the createSchedule field.
func (r *mutationResolver) CreateSchedule(ctx context.Context, schedule types.ScheduleSaver) (*types.Schedule, error) {
	_, err := helpers.PutItem(ctx, &schedule)
	if err != nil {
		return nil, err
	}
	return &schedule.Schedule, nil
}

// DeleteSchedule is the resolver for the deleteSchedule field.
func (r *mutationResolver) DeleteSchedule(ctx context.Context, scheduleGUID string, officeGUID string) (bool, error) {
	_, err := helpers.DeleteItem(ctx, &types.Schedule{ScheduleGUID: scheduleGUID, OfficeGUID: officeGUID})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetSchedule is the resolver for the getSchedule field.
func (r *queryResolver) GetSchedule(ctx context.Context, scheduleGUID string, officeGUID string) (*types.Schedule, error) {
	schedule := &types.Schedule{ScheduleGUID: scheduleGUID, OfficeGUID: officeGUID}
	_, err := helpers.GetItem(ctx, schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

// GetOfficeSchedules is the resolver for the getOfficeSchedules field.
func (r *queryResolver) GetOfficeSchedules(ctx context.Context, input types.ScheduleFinder) ([]*types.Schedule, error) {
	s := input.Schedule
	schedules, err := s.LoadSchedulesForOffice(ctx)
	if err != nil {
		return nil, err
	}
	ret := []*types.Schedule{}
	for _, schedule := range schedules {
		ret = append(ret, &schedule)
	}
	return ret, nil
}

// Start is the resolver for the start field.
func (r *scheduleResolver) Start(ctx context.Context, obj *types.Schedule) (string, error) {
	return obj.Start.String(), nil
}

// End is the resolver for the end field.
func (r *scheduleResolver) End(ctx context.Context, obj *types.Schedule) (string, error) {
	return obj.End.String(), nil
}

// Offices is the resolver for the offices field.
func (r *scheduleResolver) Offices(ctx context.Context, obj *types.Schedule) ([]*types.Office, error) {
	panic(fmt.Errorf("not implemented: Offices - offices"))
}

// Start is the resolver for the start field.
func (r *scheduleSaverResolver) Start(ctx context.Context, obj *types.ScheduleSaver, data string) error {
	panic(fmt.Errorf("not implemented: Start - start"))
}

// End is the resolver for the end field.
func (r *scheduleSaverResolver) End(ctx context.Context, obj *types.ScheduleSaver, data string) error {
	panic(fmt.Errorf("not implemented: End - end"))
}

// Schedule returns ScheduleResolver implementation.
func (r *Resolver) Schedule() ScheduleResolver { return &scheduleResolver{r} }

// ScheduleSaver returns ScheduleSaverResolver implementation.
func (r *Resolver) ScheduleSaver() ScheduleSaverResolver { return &scheduleSaverResolver{r} }

type scheduleResolver struct{ *Resolver }
type scheduleSaverResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetOfficeSchedule(ctx context.Context, scheduleGUID string) ([]*types.Schedule, error) {
	panic(fmt.Errorf("not implemented: GetOfficeSchedule - getOfficeSchedule"))
}
