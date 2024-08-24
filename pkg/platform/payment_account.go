package platform

import "github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"

type SharedAccountAccess struct {
	model.SharedAccountAccess
}

type PaymentAccount struct {
	model.PaymentAccounts
	Permission int16
}

type PaymentAccountDetail struct {
	PaymentAccount
	Permissions []SharedAccountAccess
}
