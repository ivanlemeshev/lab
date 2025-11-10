package users

type Role int

const (
	RoleUnknown Role = iota
	RoleUser
	RoleAdmin
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "User"
	case RoleAdmin:
		return "Admin"
	default:
		return "Unknown"
	}
}
