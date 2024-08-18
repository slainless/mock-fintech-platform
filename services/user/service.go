package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/auth"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type Service struct {
	db *sql.DB

	authManager    *core.AuthManager
	userManager    *core.UserManager
	accountManager *core.PaymentAccountManager
	historyManager *core.TransactionHistoryManager

	emailJwtAuth *auth.EmailJWTAuthService

	errorTracker platform.ErrorTracker
}

func NewService(
	authSecret string,
	db *sql.DB,
	services map[string]platform.PaymentService,
	tracker platform.ErrorTracker,
) *Service {
	emailJwtAuth := auth.NewEmailJWTAuthService([]byte(authSecret))

	user := core.NewUserManager(db)
	auth := core.NewAuthManager(user)
	account := core.NewPaymentAccountManager(db, services, tracker)
	history := core.NewTransactionHistoryManager(db)

	return &Service{
		db: db,

		authManager:    auth,
		userManager:    user,
		accountManager: account,
		historyManager: history,

		emailJwtAuth: emailJwtAuth,

		errorTracker: tracker,
	}
}
