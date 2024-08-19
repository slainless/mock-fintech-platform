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

func NewTransactionHistoryManager(db *sql.DB) *TransactionHistoryManager {
	return &TransactionHistoryManager{
		db: db,
	}
}

func (m *TransactionHistoryManager) GetHistories(ctx context.Context, user *platform.User, accountUUID string, from, to time.Time) ([]platform.TransactionHistory, error) {
	if accountUUID != "" {
		return query.GetHistoriesOfAccount(ctx, m.db, accountUUID, from, to)
	} else {
		return query.GetHistories(ctx, m.db, user.UUID, from, to)
	}
}

type DateRange struct {
	From      *time.Time `form:"from" time_format:"2006-01-02"`
	To        *time.Time `form:"to" time_format:"2006-01-02"`
	AccountID string     `form:"account_id"`
}

// TODO: Fix this brahh
func (m *TransactionHistoryManager) GetHistoryParams(ctx *gin.Context) (*time.Time, *time.Time, string) {
	var rg DateRange
	ctx.BindQuery(&rg)

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
