package payment

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/auth"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type Service struct {
	db *sql.DB

	authManager             *core.AuthManager
	userManager             *core.UserManager
	accountManager          *core.PaymentAccountManager
	historyManager          *core.TransactionHistoryManager
	recurringPaymentManager *core.RecurringPaymentManager
	paymentManager          *core.PaymentManager

	emailJwtAuth *auth.EmailJWTAuthService

	errorTracker platform.ErrorTracker
}
