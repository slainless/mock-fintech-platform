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

func (m *TransactionHistoryManager) GetHistories(ctx context.Context, user platform.User, from, to time.Time) ([]platform.TransactionHistory, error) {
	value, err := query.GetHistories(ctx, m.db, user.ID(), from, to)
	if err != nil {
		return nil, err
	}

	return HistoriesFrom(value), nil
}

func (m *TransactionHistoryManager) GetRange(ctx *gin.Context) (*time.Time, *time.Time, error) {

}
