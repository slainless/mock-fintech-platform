package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
)

func GetUserMonetaryAccounts(ctx context.Context, db *sql.DB, userUuid string) ([]model.MonetaryAccounts, error) {
	var accounts []model.MonetaryAccounts

	stmt := SELECT(
		table.MonetaryAccounts.UUID,
		table.MonetaryAccounts.Balance,
		table.MonetaryAccounts.Currency,
		table.MonetaryAccounts.ServiceID,
	).
		FROM(table.MonetaryAccounts).
		WHERE(table.MonetaryAccounts.UserUUID.EQ(String(userUuid)))

	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// func MutateMonetaryAccount(ctx context.Context, db *sql.DB, accountUuid string, amount int64) error {}

// func SetMonetaryAccountBalance(
// 	ctx context.Context,
// 	db *sql.DB,
// 	accountUuid string,
// 	amount int64,
// 	currency string,
// ) error {
// }

// func InsertMonetaryAccount(ctx context.Context, db *sql.DB, account *MonetaryAccountModel) error {}
