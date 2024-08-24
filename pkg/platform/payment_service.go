package platform

import (
	"context"
	"errors"
)

var (
	ErrInvalidAccountData  = errors.New("invalid account data")
	ErrTransactionRejected = errors.New("transaction rejected")
)

type PaymentService interface {
	// Send sends money from one account to another.
	Send(ctx context.Context, user *User, source, des *PaymentAccount, amount int64, callbackData string) (*TransactionHistory, error)

	// Withdraw money.
	Withdraw(ctx context.Context, user *User, account *PaymentAccount, amount int64, callbackData string) (*TransactionHistory, error)

	// Get matching history from destination service/account
	GetMatchingHistory(ctx context.Context, account *PaymentAccount, history *TransactionHistory) (*TransactionHistory, error)

	// Get balance
	Balance(ctx context.Context, account *PaymentAccount) (*MonetaryAmount, error)

	// Validate user
	Validate(ctx context.Context, user *User, accountForeignID string, callbackData string) error
}
