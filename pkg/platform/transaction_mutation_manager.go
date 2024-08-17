package platform

type TransactionMutationManager interface {
	MonetaryServices() map[string]MonetaryService

	// balance mutation
	Send(service MonetaryService, from, to User, amount MonetaryAmount) (TransactionHistory, error)
	Withdraw(service MonetaryService, from User, amount MonetaryAmount) (TransactionHistory, error)
}
