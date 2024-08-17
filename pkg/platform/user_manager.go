package platform

type UserManager interface {
	Login(service AuthService, credential any) (User, error)

	Logout(service AuthService, user User) error
	LogoutByID(service AuthService, userID string) error
}
