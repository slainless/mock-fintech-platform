package platform

import (
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
)

type TransactionHistory struct {
	model.TransactionHistories
	ServiceUUID string
	UserUUID    string
}

func (t *TransactionHistory) Clone() TransactionHistory {
	return TransactionHistory{
		TransactionHistories: t.TransactionHistories,
		ServiceUUID:          t.ServiceUUID,
		UserUUID:             t.UserUUID,
	}
}
