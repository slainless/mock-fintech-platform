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

func (m *TransactionHistoryManager) GetHistories(ctx context.Context, user *platform.User, account *platform.PaymentAccount, from, to time.Time) (histories []platform.TransactionHistory, err error) {
	if account != nil {
		histories, err = query.GetHistoriesOfAccountWithAccess(ctx, m.db, user.UUID, account.UUID, from, to, AccountPermissionHistory)
	} else {
		histories, err = query.GetHistoriesWithAccess(ctx, m.db, user.UUID, from, to, AccountPermissionHistory)
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
