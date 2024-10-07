package entity

type XxxRequest struct {
}

type XxxResponse struct {
}

type XxxResult struct {
}

const (
	RoleSuperAdmin = "super-admin"
	RoleAdmin      = "admin"
)

type Role string

type User struct {
	ID       int
	Username string
	Password string
	RoleID   int
	Role 	Role
}
