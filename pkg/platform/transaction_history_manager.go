package platform

import "time"

type TransactionHistoryManager interface {
	// get all transaction histories of the entire platform
	// this should not be implemented in production 💀
	AllHistory() ([]TransactionHistory, error)

	// get all transaction histories of a specific user
	// this should also not be implemented in production 💀
	// partial view should be used instead.
	AllHistoryOf(user User) ([]TransactionHistory, error)

	// get all transaction histories of a specific user
	// this should also not be implemented in production 💀
	// partial view should be used instead.
	AllHistoryByServiceOf(paymentService MonetaryService, user User) ([]TransactionHistory, error)

	LookupHistory(id string) (TransactionHistory, error)

	HistoryOf(user User, from, to time.Time) ([]TransactionHistory, error)
	HistoryByServiceOf(paymentService MonetaryService, user User, from, to time.Time) ([]TransactionHistory, error)
}
