package platform

type RecurrentBillingService interface {
	ID() string

	Charge(user User) (TransactionHistory, error)
}
