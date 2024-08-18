package platform

// entrypoint of the platform
// separation of concerns should also occurs here

// NOTE: looks like we are building microservices
// instead of monolithic platform...
// this will be ignored.
type Platform interface {
	// history manager
	HistoryManager() TransactionHistoryManager

	// mutation manager
	MutationManager() TransactionMutationManager

	// user manager
	UserManager() UserManager

	// money exchange manager
	MoneyExchangeManager() MoneyExchangeManager

	// recurrent transaction manager
	RecurrentTransactionManager() RecurrentTransactionManager
}
