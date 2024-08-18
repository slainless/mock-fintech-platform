package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var ErrPaymentServiceNotSupported = errors.New("payment service not supported")

type PaymentManager struct {
	services map[string]platform.PaymentService

	errorTracker platform.ErrorTracker

	accountManager *PaymentAccountManager
	historyManager *TransactionHistoryManager
}

func (m *PaymentManager) Send(ctx context.Context, from, to *platform.PaymentAccount, amount int64) (*platform.TransactionHistory, error) {
	service := m.services[from.ServiceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	sourceHistory, err := service.Send(ctx, from, to)
	if err != nil {
		return nil, err
	}

	var destHistory *platform.TransactionHistory
	destHistory, err = service.GetMatchingHistory(ctx, to, sourceHistory)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		destHistory = m.historyManager.CreateMakeshiftMatchingTransferHistory(ctx, sourceHistory, fmt.Sprintf("Transfer from %s", from.UUID))
		return nil, err
	}

	m.errorTracker.Report(ctx,
		m.historyManager.Records(ctx, []platform.TransactionHistory{
			*sourceHistory,
			*destHistory,
		}))

	return sourceHistory, nil
}

func (m *PaymentManager) Withdraw(ctx context.Context, from *platform.PaymentAccount, amount int64, callbackData string) (*platform.TransactionHistory, error) {
	service := m.services[from.ServiceID]
	if service == nil {
		return nil, ErrPaymentServiceNotSupported
	}

	history, err := service.Withdraw(ctx, from, amount, callbackData)
	if err != nil {
		return nil, err
	}

	m.errorTracker.Report(ctx,
		m.historyManager.Records(ctx, []platform.TransactionHistory{*history}))

	return history, nil
}
