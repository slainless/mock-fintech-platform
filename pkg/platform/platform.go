package platform

// entrypoint of the platform
// separation of concerns should also occurs here
type Platform interface {
	// history manager
	HistoryManager() TransactionHistoryManager

	// mutation manager
	MutationManager() TransactionMutationManager

	// user manager
	UserManager() UserManager
}
