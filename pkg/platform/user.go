package platform

type User interface {
	// should returns UUID of user
	ID() string
}
