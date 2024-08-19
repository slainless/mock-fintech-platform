package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func GetAllAccounts(ctx context.Context, db *sql.DB, userUUID string) ([]platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(table.PaymentAccounts.UserUUID.EQ(String(userUUID)))

	accounts := make([]platform.PaymentAccount, 0)
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccount(ctx context.Context, db *sql.DB, accountUUID string) (*platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(table.PaymentAccounts.UUID.EQ(String(accountUUID)))

	var account platform.PaymentAccount
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetTwoAccounts(ctx context.Context, db *sql.DB, FirstUUID, SecondUUID string) (*platform.PaymentAccount, *platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(
			table.PaymentAccounts.UUID.EQ(String(FirstUUID)).
				OR(table.PaymentAccounts.UUID.EQ(String(SecondUUID))),
		)

	var accounts []platform.PaymentAccount
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, nil, err
	}

	var first, second *platform.PaymentAccount
	if accounts[0].UUID == FirstUUID {
		first = &accounts[0]
		second = &accounts[1]
	} else {
		first = &accounts[1]
		second = &accounts[0]
	}

	return first, second, nil
}

func CheckOwner(ctx context.Context, db *sql.DB, userUUID, accountUUID string) error {
	stmt := SELECT(Bool(true)).
		FROM(table.PaymentAccounts).
		WHERE(
			table.PaymentAccounts.UUID.EQ(String(accountUUID)).
				AND(table.PaymentAccounts.UserUUID.EQ(String(userUUID))),
		)

	var exists bool
	err := stmt.QueryContext(ctx, db, &exists)
	if err != nil {
		return err
	}

	return nil
}

func InsertAccount(ctx context.Context, db *sql.DB, account *platform.PaymentAccount) error {
	stmt := table.PaymentAccounts.INSERT(
		table.PaymentAccounts.AllColumns.Except(table.PaymentAccounts.ID),
	).
		VALUES(account)

	_, err := stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
