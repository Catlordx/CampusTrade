package role

const (
	User  uint = 1
	Admin uint = 2
)

func String(roleID uint) string {
	switch roleID {
	case User:
		return "user"
	case Admin:
		return "admin"
	}
	return ""
}

func ID(roleName string) uint {
	switch roleName {
	case "user":
		return User
	case "admin":
		return Admin
	}
	return 0
}
