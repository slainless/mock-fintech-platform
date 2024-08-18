package platform

import "context"

type MonetaryAccountManager interface {
	Accounts(ctx context.Context, user User) ([]MonetaryAccount, error)

	// balance mutation
	Send(service MonetaryService, from, to User, amount MonetaryAmount) (TransactionHistory, error)
	Withdraw(service MonetaryService, from User, amount MonetaryAmount) (TransactionHistory, error)
}
