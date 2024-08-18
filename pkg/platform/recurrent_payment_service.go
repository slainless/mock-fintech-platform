package platform

type RecurrentPaymentService interface {
	ID() string

	// Charging callback. This should be called
	// by recurrent transaction manager.
	Charge(user User) (TransactionHistory, error)
}
