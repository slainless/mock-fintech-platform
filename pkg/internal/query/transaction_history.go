package query

import (
	"context"
	"database/sql"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func historySelection() SelectStatement {
	return SELECT(
		table.TransactionHistories.AllColumns,
		table.PaymentAccounts.ServiceID,
		table.PaymentAccounts.UserUUID,
	).
		FROM(
			table.TransactionHistories.
				INNER_JOIN(table.PaymentAccounts, table.TransactionHistories.AccountUUID.EQ(table.PaymentAccounts.UUID)).
				INNER_JOIN(table.SharedAccountAccess, table.TransactionHistories.AccountUUID.EQ(table.SharedAccountAccess.AccountUUID)),
		)
}

func GetHistoriesOfAccountWithAccess(ctx context.Context, db *sql.DB, userUUID, accountUUID uuid.UUID, from, to time.Time, access AccountPermission) ([]platform.TransactionHistory, error) {
	stmt := historySelection().
		WHERE(
			table.TransactionHistories.AccountUUID.EQ(UUID(accountUUID)).
				AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))).
				AND(OR(
					table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)),
					table.SharedAccountAccess.UserUUID.EQ(UUID(userUUID)).
						AND(table.SharedAccountAccess.Permission.BIT_AND(Int32(int32(access))).EQ(Int32(int32(access)))),
				)),
		).
		GROUP_BY(table.TransactionHistories.UUID)

	histories := make([]platform.TransactionHistory, 0)
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func GetHistoriesWithAccess(ctx context.Context, db *sql.DB, userUUID uuid.UUID, from, to time.Time, access AccountPermission) ([]platform.TransactionHistory, error) {
	stmt := historySelection().
		WHERE(
			OR(
				table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)),
				table.SharedAccountAccess.UserUUID.EQ(UUID(userUUID)).
					AND(table.SharedAccountAccess.Permission.BIT_AND(Int32(int32(access))).EQ(Int32(int32(access)))),
			).AND(table.TransactionHistories.TransactionDate.GT_EQ(TimestampT(from))).
				AND(table.TransactionHistories.TransactionDate.LT_EQ(TimestampT(to))),
		).
		GROUP_BY(table.TransactionHistories.UUID)

	histories := make([]platform.TransactionHistory, 0)
	err := stmt.QueryContext(ctx, db, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func InsertHistories(ctx context.Context, db *sql.DB, histories []platform.TransactionHistory) error {
	stmt := table.TransactionHistories.INSERT(table.TransactionHistories.MutableColumns).
		MODELS(histories).
		RETURNING(table.TransactionHistories.AllColumns)

	_, err := stmt.ExecContext(ctx, db)
	return err
}
