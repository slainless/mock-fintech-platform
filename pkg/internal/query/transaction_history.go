package query

import (
	"context"
	"database/sql"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
)

type History struct {
	model.TransactionHistories
	ServiceID string
	UserUUID  string
}

func historySelection() SelectStatement {
	return SELECT(
		table.TransactionHistories.AllColumns,
		table.PaymentAccounts.ServiceID,
		table.Users.UUID,
	).
		FROM(
			table.TransactionHistories.
				INNER_JOIN(table.PaymentAccounts, table.TransactionHistories.AccountUUID.EQ(table.PaymentAccounts.UUID)).
				INNER_JOIN(table.Users, table.PaymentAccounts.UserUUID.EQ(table.Users.UUID)),
		)
}

func GetHistoriesOfAccount(ctx context.Context, db *sql.DB, accountUuid string, from, to time.Time) ([]History, error) {
	stmt := historySelection().
		WHERE(
			table.TransactionHistories.AccountUUID.EQ(String(accountUuid)).
				AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))),
		)

	var histories []History
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func GetHistories(ctx context.Context, db *sql.DB, userUuid string, from, to time.Time) ([]History, error) {
	stmt := historySelection().
		WHERE(
			table.Users.UUID.EQ(String(userUuid)).
				AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))),
		)

	var histories []History
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

// func GetUserHistories(
// 	ctx context.Context,
// 	db *sql.DB,
// 	userUuid string,
// 	from, to time.Time,
// ) ([]TransactionHistoryModel, error) {
// }

// func GetHistory(
// 	ctx context.Context,
// 	db *sql.DB,
// 	historyUuid string,
// ) (*TransactionHistoryModel, error) {
// }

// func InsertHistory(ctx context.Context, db *sql.DB, history *TransactionHistoryModel) error {}
