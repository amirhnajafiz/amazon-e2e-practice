package routes

type Routes struct {
	Auth   *AuthGroup
	Health *HealthGroup
	Users  *UsersGroup
}

func NewRoutes() *Routes {
	return &Routes{
		Auth:   &AuthGroup{},
		Health: &HealthGroup{},
		Users:  &UsersGroup{},
	}
}
