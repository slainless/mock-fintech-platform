package platform

// entrypoint of the platform
// separation of concerns should also occurs here
type TransactionManager interface {
	// service ID
	ID() string

	// history manager
	HistoryManager() TransactionHistoryManager

	// mutation manager
	MutationManager() TransactionMutationManager
}
