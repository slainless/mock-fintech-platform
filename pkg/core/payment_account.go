package core

import (
	"context"
	"database/sql"
	"errors"

	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var ErrInvalidTransferDestination = errors.New("invalid transfer destination")
var ErrInvalidAccount = errors.New("invalid account")

type PaymentAccountManager struct {
	db *sql.DB
}

func (m *PaymentAccountManager) GetAccounts(ctx context.Context, user *platform.User) ([]platform.PaymentAccount, error) {
	accounts, err := query.GetAllAccounts(ctx, m.db, user.UUID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (m *PaymentAccountManager) GetAccount(ctx context.Context, accountUUID string) (*platform.PaymentAccount, error) {
	account, err := query.GetAccount(ctx, m.db, accountUUID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (m *PaymentAccountManager) PrepareTransfer(ctx context.Context, fromUUID, toUUID string) (*platform.PaymentAccount, *platform.PaymentAccount, error) {
	from, to, err := query.GetTwoAccounts(ctx, m.db, fromUUID, toUUID)
	if err != nil {
		return nil, nil, err
	}

	if from == nil {
		return nil, nil, ErrInvalidAccount
	}

	if to == nil {
		return nil, nil, ErrInvalidTransferDestination
	}

	return from, to, nil
}

func (m *PaymentAccountManager) CheckOwner(ctx context.Context, user *platform.User, accountUUID string) error {
	return query.CheckOwner(ctx, m.db, user.UUID, accountUUID)
}
