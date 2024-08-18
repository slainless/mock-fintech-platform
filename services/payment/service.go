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
	recurrentPaymentManager *core.RecurrentPaymentManager
	paymentManager          *core.PaymentManager

	supabaseJwtAuth *auth.SupabaseJWTAuthService

	errorTracker platform.ErrorTracker
}
