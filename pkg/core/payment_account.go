package core

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var ErrInvalidTransferDestination = errors.New("invalid transfer destination")
var ErrAccountNotFound = errors.New("account not found or no permission")
var ErrAccountAlreadyRegistered = errors.New("account already registered")

type AccountPermission = query.AccountPermission

const (
	AccountPermissionBase         = query.AccountPermissionBase
	AccountPermissionRead         = query.AccountPermissionRead
	AccountPermissionHistory      = query.AccountPermissionHistory
	AccountPermissionSend         = query.AccountPermissionSend
	AccountPermissionWithdraw     = query.AccountPermissionWithdraw
	AccountPermissionSubscription = query.AccountPermissionSubscription
	AccountPermissionAll          = query.AccountPermissionAll
)

type PaymentAccountManager struct {
	services     map[string]platform.PaymentService
	errorTracker platform.ErrorTracker

	db *sql.DB
}

func NewPaymentAccountManager(db *sql.DB, svc map[string]platform.PaymentService, tracker platform.ErrorTracker) *PaymentAccountManager {
	return &PaymentAccountManager{
		db: db,

		services:     svc,
		errorTracker: tracker,
	}
}

func (m *PaymentAccountManager) GetAccountsWithAccess(ctx context.Context, user *platform.User, access AccountPermission) ([]platform.PaymentAccount, error) {
	accounts, err := query.GetAllAccountsWithAccess(ctx, m.db, user.UUID, access)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return accounts, nil
}

func (m *PaymentAccountManager) GetAccount(ctx context.Context, accountUUID uuid.UUID) (*platform.PaymentAccount, error) {
	account, err := query.GetAccount(ctx, m.db, accountUUID)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return account, nil
}

func (m *PaymentAccountManager) GetAccountWithAccess(ctx context.Context, user *platform.User, accountUUID uuid.UUID, access AccountPermission) (*platform.PaymentAccount, error) {
	account, err := query.GetAccountWithAccess(ctx, m.db, user.UUID, accountUUID, access)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return account, nil
}

func (m *PaymentAccountManager) CheckOwner(ctx context.Context, user *platform.User, accountUUID uuid.UUID) error {
	err := query.CheckOwner(ctx, m.db, user.UUID, accountUUID)
	if err != nil {
		if err == qrm.ErrNoRows {
			return ErrAccountNotFound
		}
		m.errorTracker.Report(ctx, err)
		return err
	}

	return nil
}

func (m *PaymentAccountManager) GetBalance(ctx context.Context, account *platform.PaymentAccount) (*platform.MonetaryAmount, error) {
	service := m.services[account.ServiceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	balance, err := service.Balance(ctx, account)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return balance, nil
}

func (m *PaymentAccountManager) Register(ctx context.Context, user *platform.User, serviceID string, name string, accountForeignID string, CallbackData string) (*platform.PaymentAccount, error) {
	service := m.services[serviceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	err := service.Validate(ctx, user, accountForeignID, CallbackData)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	account := &platform.PaymentAccount{
		PaymentAccounts: model.PaymentAccounts{
			UserUUID:  user.UUID,
			ServiceID: serviceID,
			ForeignID: accountForeignID,
			Name:      &name,
		},
	}
	err = query.InsertAccount(ctx, m.db, account)
	if err != nil {
		if err := util.PQErrorCode(err); err != "" {
			switch err {
			case "unique_violation":
				return nil, ErrAccountAlreadyRegistered
			}
		}
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return account, nil
}
