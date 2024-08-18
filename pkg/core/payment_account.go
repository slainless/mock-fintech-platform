package core

import (
	"context"
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type PaymentAccountManager struct {
	db *sql.DB
}

func (m *PaymentAccountManager) GetAccounts(ctx context.Context, user platform.User) ([]platform.PaymentAccount, error) {
	models, err := query.GetAllAccounts(ctx, m.db, user.ID())
	if err != nil {
		return nil, err
	}

	return PaymentAccountsFrom(models), nil
}
