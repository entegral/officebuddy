package types

import (
	"time"

	"github.com/entegral/toolbox/types"
)

// Schedule is a type for defining a schedule for a given office.
type Schedule struct {
	types.Row
	ScheduleGUID string    `json:"scheduleGUID"`
	OfficeGUID   string    `json:"OfficeGUID"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Active       bool      `json:"active"`
	Coworkers    []*User   `json:"coworkers,omitempty"`
}

// Keys returns the partition key and sort key for the row
func (s *Schedule) Keys(gsi int) (string, string) {
	// For this example, assuming GUID is the partition key and Email is the sort key.
	// Additional logic can be added to handle different GSIs if necessary.
	s.Pk = "schedule:" + s.ScheduleGUID
	s.Sk = "office:" + s.OfficeGUID
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
