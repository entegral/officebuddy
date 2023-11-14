package types

type Schedule struct {
	DayNumber int       `json:"dayNumber"`
	Day       DayOfWeek `json:"day,omitempty"`
	Active    bool      `json:"active"`
	Coworkers []*User   `json:"coworkers,omitempty"`
}

type DayOfWeek string

const (
	DayOfWeekMon DayOfWeek = "Mon"
	DayOfWeekTue DayOfWeek = "Tue"
	DayOfWeekWed DayOfWeek = "Wed"
	DayOfWeekThu DayOfWeek = "Thu"
	DayOfWeekFri DayOfWeek = "Fri"
	DayOfWeekSat DayOfWeek = "Sat"
	DayOfWeekSun DayOfWeek = "Sun"
)
