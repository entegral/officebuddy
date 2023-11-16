package types

type Office struct {
	GUID        string   `json:"guid"`
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Admins      []*User  `json:"admins,omitempty"`
	Members     []*User  `json:"members,omitempty"`
	Events      []*Event `json:"events,omitempty"`
}
