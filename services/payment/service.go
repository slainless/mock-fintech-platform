package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/core"
)

type Service struct {
	db *sql.DB

	AuthManager             *core.AuthManager
	UserManager             *core.UserManager
	AccountManager          *core.PaymentAccountManager
	HistoryManager          *core.TransactionHistoryManager
	RecurrentPaymentManager *core.RecurrentPaymentManager
	PaymentManager          *core.PaymentManager
}
