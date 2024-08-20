package platform

import (
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
)

type TransactionHistory struct {
	model.TransactionHistories
	ServiceUUID string
	UserUUID    uuid.UUID
}
