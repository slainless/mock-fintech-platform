package platform

type User interface {
	ID() string
	PaymentAccounts() map[string]PaymentAccount
	// TODO: add return signature
	PaymentHistory()
}
