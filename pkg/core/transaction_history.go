package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type TransactionHistoryManager struct {
	db *sql.DB
}

func (m *TransactionHistoryManager) GetHistories(ctx context.Context, user *platform.User, from, to time.Time) ([]platform.TransactionHistory, error) {
	histories, err := query.GetHistories(ctx, m.db, user.UUID, from, to)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (m *TransactionHistoryManager) GetRange(ctx *gin.Context) (*time.Time, *time.Time, error) {

}

func (m *TransactionHistoryManager) Records(ctx context.Context, histories []platform.TransactionHistory) error {
	err := query.InsertHistories(ctx, m.db, histories)
	if err != nil {
		return err
	}

	return nil
}

func (m *TransactionHistoryManager) CreateMakeshiftMatchingTransferHistory(ctx context.Context, history *platform.TransactionHistory, note string) *platform.TransactionHistory {
	match := *history
	match.TransactionNote = &note
	match.AccountUUID = *history.DestUUID
	match.Mutation = history.Mutation * -1
	match.Address = nil

	return &match
}
