package types

type UserRole string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
)

func (r UserRole) String() string {
	return string(r)
}
