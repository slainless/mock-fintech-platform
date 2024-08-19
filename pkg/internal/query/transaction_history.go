package query

import (
	"context"
	"database/sql"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

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

func GetHistoriesOfAccount(ctx context.Context, db *sql.DB, accountUUID string, from, to time.Time) ([]platform.TransactionHistory, error) {
	stmt := historySelection().
		WHERE(
			table.TransactionHistories.AccountUUID.EQ(String(accountUUID)).
				AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))),
		)

	var histories []platform.TransactionHistory
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func GetHistories(ctx context.Context, db *sql.DB, userUUID string, from, to time.Time) ([]platform.TransactionHistory, error) {
	stmt := historySelection().
		WHERE(
			table.Users.UUID.EQ(String(userUUID)).
				AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))),
		)

	var histories []platform.TransactionHistory
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func InsertHistories(ctx context.Context, db *sql.DB, histories []platform.TransactionHistory) error {
	stmt := table.TransactionHistories.INSERT(
		table.TransactionHistories.AllColumns.Except(table.TransactionHistories.ID),
	).
		MODELS(histories)

	_, err := stmt.ExecContext(ctx, db)
	return err
}
