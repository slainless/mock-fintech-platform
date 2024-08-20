package payment

import (
	"context"
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

func NewService(
	authSecret string,
	db *sql.DB,
	paymentServices map[string]platform.PaymentService,
	recurringPaymentServices map[string]platform.RecurringPaymentService,
	tracker platform.ErrorTracker,
) *Service {
	emailJwtAuth := auth.NewEmailJWTAuthService([]byte(authSecret))

	user := core.NewUserManager(db, tracker)
	auth := core.NewAuthManager(user)
	account := core.NewPaymentAccountManager(db, paymentServices, tracker)
	history := core.NewTransactionHistoryManager(db, tracker)
	recurringPayment := core.NewRecurringPaymentManager(db, recurringPaymentServices, history, tracker)
	err := recurringPayment.InitScheduler()
	if err != nil {
		tracker.Report(context.TODO(), err)
		panic(err)
	}

	payment := core.NewPaymentManager(account, history, paymentServices, tracker)

	return &Service{
		db: db,

		authManager:             auth,
		userManager:             user,
		accountManager:          account,
		historyManager:          history,
		recurringPaymentManager: recurringPayment,
		paymentManager:          payment,

		emailJwtAuth: emailJwtAuth,

		errorTracker: tracker,
	}
}
