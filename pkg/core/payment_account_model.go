package core

import (
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type PaymentAccount struct {
	model *model.PaymentAccounts
}

func (a *PaymentAccount) Balance() platform.MonetaryAmount {
	return platform.MonetaryAmount{
		Currency: platform.Currency(a.model.Currency),
		Value:    a.model.Balance,
	}
}

// Currency implements platform.PaymentAccount.
func (a *PaymentAccount) Currency() platform.Currency {
	return platform.Currency(a.model.Currency)
}

// ID implements platform.PaymentAccount.
func (a *PaymentAccount) ID() string {
	return a.model.UUID
}

// ServiceID implements platform.PaymentAccount.
func (a *PaymentAccount) ServiceID() string {
	return a.model.ServiceID
}

func PaymentAccountFrom(model *model.PaymentAccounts) platform.PaymentAccount {
	return &PaymentAccount{model: model}
}

func PaymentAccountsFrom(models []model.PaymentAccounts) []platform.PaymentAccount {
	accounts := make([]platform.PaymentAccount, 0, len(models))
	for _, model := range models {
		accounts = append(accounts, PaymentAccountFrom(&model))
	}

	return accounts
}
