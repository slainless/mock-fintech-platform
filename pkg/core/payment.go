package core

import (
	"context"
	"errors"

	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var (
	ErrPaymentServiceNotSupported = errors.New("payment service not supported")
	ErrInvalidAmount              = errors.New("invalid amount")
)

type PaymentManager struct {
	services map[string]platform.PaymentService

	errorTracker platform.ErrorTracker

	accountManager *PaymentAccountManager
	historyManager *TransactionHistoryManager
}

func NewPaymentManager(account *PaymentAccountManager, history *TransactionHistoryManager, svc map[string]platform.PaymentService, tracker platform.ErrorTracker) *PaymentManager {
	return &PaymentManager{
		accountManager: account,
		historyManager: history,
		services:       svc,
		errorTracker:   tracker,
	}
}

func (m *PaymentManager) Send(ctx context.Context, user *platform.User, from, to *platform.PaymentAccount, amount int64, callbackData string) (*platform.TransactionHistory, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	service := m.services[from.ServiceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	sourceHistory, err := service.Send(ctx, user, from, to, amount, callbackData)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	var destHistory *platform.TransactionHistory
	destHistory, err = service.GetMatchingHistory(ctx, to, sourceHistory)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		m.historyManager.Records(ctx, []platform.TransactionHistory{*sourceHistory})
	} else {
		m.historyManager.Records(ctx, []platform.TransactionHistory{
			*sourceHistory,
			*destHistory,
		})
	}

	return sourceHistory, nil
}

func (m *PaymentManager) Withdraw(ctx context.Context, user *platform.User, from *platform.PaymentAccount, amount int64, callbackData string) (*platform.TransactionHistory, error) {
	service := m.services[from.ServiceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	history, err := service.Withdraw(ctx, user, from, amount, callbackData)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	m.historyManager.Records(ctx, []platform.TransactionHistory{*history})

	return history, nil
}
