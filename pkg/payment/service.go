package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/manager"
)

type Service struct {
	db *sql.DB

	AuthManager             *manager.AuthManager
	UserManager             *manager.UserManager
	AccountManager          *manager.PaymentAccountManager
	HistoryManager          *manager.TransactionHistoryManager
	RecurrentPaymentManager *manager.RecurrentPaymentManager
	PaymentManager          *manager.PaymentManager
}
