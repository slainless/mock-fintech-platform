package platform

type User interface {
	ID() string

	PaymentAccounts() map[string]PaymentAccount
}
