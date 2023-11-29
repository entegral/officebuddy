package types

import (
	"time"

	"github.com/entegral/toolbox/dynamo"
)

// OfficeMembership is a type for defining a membership of a user to an office.
type OfficeMembership struct {
	dynamo.Linker[*User, *Office]
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

// Role is a type for defining the role of a user in an office.
type Role string

// String returns the string representation of the role.
func (r Role) String() string {
	return string(r)
}

const (
	// RoleAdmin is the admin role.
	RoleAdmin Role = "admin"
	// RoleMember is the member role.
	RoleMember Role = "member"
)

func a() {

}
