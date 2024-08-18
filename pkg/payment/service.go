package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type PaymentService struct {
	db *sql.DB

	userManager             *core.UserManager
	accountManager          *core.MonetaryAccountManager
	historyManager          *core.TransactionHistoryManager
	moneyExchangeManager    *core.MoneyExchangeManager
	recurrentPaymentManager *core.RecurrentPaymentManager
}

func (s *PaymentService) UserManager() platform.UserManager {
	return s.userManager
}

func (s *PaymentService) AccountManager() platform.MonetaryAccountManager {
	return s.accountManager
}

func (s *PaymentService) TransactionHistoryManager() platform.TransactionHistoryManager {
	return s.historyManager
}

func (s *PaymentService) MoneyExchangeManager() platform.MoneyExchangeManager {
	return s.moneyExchangeManager
}

func (s *PaymentService) RecurrentPaymentManager() platform.RecurrentPaymentManager {
	return s.recurrentPaymentManager
}

func NewPaymentService(db *sql.DB) IPaymentService {
	return &PaymentService{}
}
