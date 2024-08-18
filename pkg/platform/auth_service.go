package platform

type AuthService interface {
	Authenticate(credential any) (userUUID string, err error)
}
