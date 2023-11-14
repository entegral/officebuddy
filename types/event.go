package types

type Event struct {
	GUID        string    `json:"guid"`
	Title       string    `json:"title"`
	Schedule    *Schedule `json:"schedule"`
	Description string    `json:"description"`
	Offices     []*Office `json:"offices,omitempty"`
}
