package platform

type AuthService interface {
	Authenticate(credential any) (userID string, err error)
}
