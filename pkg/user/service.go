package user

import (
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type UserService struct {
	db *sql.DB

	userManager    *core.UserManager
	accountManager *core.MonetaryAccountManager
	historyManager *core.TransactionHistoryManager

	services map[string]platform.MonetaryService
}

func (s *UserService) MonetaryServices() map[string]platform.MonetaryService {
	return s.services
}

func (s *UserService) UserManager() platform.UserManager {
	return s.userManager
}

func (s *UserService) AccountManager() platform.MonetaryAccountManager {
	return s.accountManager
}

func (s *UserService) TransactionHistoryManager() platform.TransactionHistoryManager {
	return s.historyManager
}

func NewUserService(db *sql.DB) IUserService {
	return &UserService{}
}
