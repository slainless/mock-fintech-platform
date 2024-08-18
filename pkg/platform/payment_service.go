package platform

import (
	"context"
)

type PaymentService interface {
	// Send sends money from one account to another.
	Send(ctx context.Context, source, des *PaymentAccount) (*TransactionHistory, error)

	// Withdraw money.
	Withdraw(ctx context.Context, account *PaymentAccount, amount int64, callbackData string) (*TransactionHistory, error)

	// Get matching history from destination service/account
	GetMatchingHistory(ctx context.Context, account *PaymentAccount, history *TransactionHistory) (*TransactionHistory, error)
}
