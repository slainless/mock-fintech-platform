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

	errorTracker platform.ErrorTracker
}

func NewTransactionHistoryManager(db *sql.DB, errorTracker platform.ErrorTracker) *TransactionHistoryManager {
	return &TransactionHistoryManager{
		db: db,

		errorTracker: errorTracker,
	}
}

func (m *TransactionHistoryManager) GetHistories(ctx context.Context, user *platform.User, accountUUID string, from, to time.Time) ([]platform.TransactionHistory, error) {
	var histories []platform.TransactionHistory
	var err error
	if accountUUID != "" {
		histories, err = query.GetHistoriesOfAccount(ctx, m.db, accountUUID, from, to)
	} else {
		histories, err = query.GetHistories(ctx, m.db, user.UUID, from, to)
	}

	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return histories, nil
}

type HistoryParams struct {
	From      *time.Time `form:"from" time_format:"2006-01-02"`
	To        *time.Time `form:"to" time_format:"2006-01-02"`
	AccountID string     `form:"account_id" binding:"uuid"`
}

// TODO: Fix this brahh
func (m *TransactionHistoryManager) GetHistoryParams(ctx *gin.Context) (*time.Time, *time.Time, string) {
	var rg HistoryParams
	ctx.ShouldBindQuery(&rg)

	if rg.From == nil {
		rg.From = &time.Time{}
	}

	if rg.To == nil {
		date := time.Now()
		rg.To = &date
	}

	return rg.From, rg.To, rg.AccountID
}

func (m *TransactionHistoryManager) Records(ctx context.Context, histories []platform.TransactionHistory) error {
	err := query.InsertHistories(ctx, m.db, histories)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return err
	}

	return nil
}
