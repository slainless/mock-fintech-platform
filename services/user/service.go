package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/auth"
	"github.com/slainless/mock-fintech-platform/pkg/manager"
)

type UserService struct {
	db *sql.DB

	authManager    *manager.AuthManager
	userManager    *manager.UserManager
	accountManager *manager.PaymentAccountManager
	historyManager *manager.TransactionHistoryManager

	supabaseJwtAuth *auth.SupabaseJWTAuthService
}
