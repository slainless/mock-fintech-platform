package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/auth"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type Service struct {
	db *sql.DB

	authManager       *core.AuthManager
	userManager       *core.UserManager
	accountManager    *core.PaymentAccountManager
	historyManager    *core.TransactionHistoryManager
	recurringPayments *core.RecurringPaymentManager

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

	user := core.NewUserManager(db, tracker)
	auth := core.NewAuthManager(user)
	account := core.NewPaymentAccountManager(db, services, tracker)
	history := core.NewTransactionHistoryManager(db, tracker)
	recurringPayments := core.NewRecurringPaymentManager(
		db,
		map[string]platform.RecurringPaymentService{},
		history,
		tracker,
	)

	return &Service{
		db: db,

		authManager:       auth,
		userManager:       user,
		accountManager:    account,
		historyManager:    history,
		recurringPayments: recurringPayments,

		emailJwtAuth: emailJwtAuth,

		errorTracker: tracker,
	}
}
