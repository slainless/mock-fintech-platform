package core

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var (
	ErrInvalidTransferDestination = errors.New("invalid transfer destination")
	ErrAccountNotFound            = errors.New("account not found or no permission")
	ErrAccountAlreadyRegistered   = errors.New("account already registered")
	ErrInvalidPermission          = errors.New("invalid permission")
)

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

func (m *PaymentAccountManager) Register(ctx context.Context, user *platform.User, serviceID string, name string, accountForeignID string, CallbackData string) (*platform.PaymentAccountDetail, error) {
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

	access := platform.SharedAccountAccess{}
	access.AccountUUID = account.UUID
	access.UserUUID = user.UUID
	access.Permission = int32(AccountPermissionAll)

	acc := &platform.PaymentAccountDetail{
		PaymentAccount: *account,
		Permissions:    []platform.SharedAccountAccess{access},
	}

	return acc, nil
}

func (m *PaymentAccountManager) ParsePermission(permission []string) (AccountPermission, error) {
	perm := AccountPermissionBase
	for _, p := range permission {
		switch p {
		case "read":
			perm |= AccountPermissionRead
		case "history":
			perm |= AccountPermissionHistory
		case "send":
			perm |= AccountPermissionSend
		case "withdraw":
			perm |= AccountPermissionWithdraw
		case "subscription":
			perm |= AccountPermissionSubscription
		case "all":
			perm |= AccountPermissionAll
		default:
			return 0, ErrInvalidPermission
		}
	}

	return perm, nil
}

func (m *PaymentAccountManager) SetPermission(ctx context.Context, userUUID, accountUUID uuid.UUID, permission AccountPermission) error {
	err := query.SetPermission(ctx, m.db, userUUID, accountUUID, permission)
	if err != nil {
		if err := util.PQError(err); err != nil {
			switch err.Code.Name() {
			case "foreign_key_violation":
				switch {
				case strings.HasSuffix(err.Constraint, "account_uuid_fkey"):
					return ErrAccountNotFound
				case strings.HasSuffix(err.Constraint, "user_uuid_fkey"):
					return ErrUserNotRegistered
				}
			}
		}
		m.errorTracker.Report(ctx, err)
		return err
	}
	return nil
}

func (m *PaymentAccountManager) GetAccountDetail(ctx context.Context, user *platform.User, accountUUID uuid.UUID) (*platform.PaymentAccountDetail, error) {
	account, err := query.GetAccountDetail(ctx, m.db, user.UUID, accountUUID)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	if account.UserUUID == user.UUID {
		if account.Permissions == nil || len(account.Permissions) == 0 {
			access := platform.SharedAccountAccess{}
			access.AccountUUID = accountUUID
			access.UserUUID = user.UUID
			access.Permission = int32(AccountPermissionAll)
			account.Permissions = []platform.SharedAccountAccess{access}
		} else {
			for _, access := range account.Permissions {
				if access.UserUUID == uuid.Nil {
					access.UserUUID = user.UUID
					access.AccountUUID = accountUUID
					access.Permission = int32(AccountPermissionAll)
				}
			}
		}
		account.Permission = int32(AccountPermissionAll)
	} else {
		for _, access := range account.Permissions {
			if access.UserUUID == user.UUID {
				account.Permission = access.Permission
				break
			}
		}
	}

	return account, nil
}
