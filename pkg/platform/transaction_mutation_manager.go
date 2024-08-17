package platform

type TransactionMutationManager interface {
	// balance mutation
	Send(service MonetaryService, from, to User, amount MonetaryAmount) (TransactionHistory, error)
	Withdraw(service MonetaryService, from User, amount MonetaryAmount) (TransactionHistory, error)
}
