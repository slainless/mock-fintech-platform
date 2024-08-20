package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func GetRecurringPayment(ctx context.Context, db *sql.DB, uuid string) (*platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns.Except(table.RecurringPayments.ID)).
		FROM(table.RecurringPayments).
		WHERE(table.RecurringPayments.UUID.EQ(String(uuid)))

	var account platform.RecurringPayment
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetRecurringPaymentWhereUser(ctx context.Context, db *sql.DB, userUUID, uuid string) (*platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns.Except(table.RecurringPayments.ID)).
		FROM(
			table.RecurringPayments.
				INNER_JOIN(table.PaymentAccounts, table.PaymentAccounts.UUID.EQ(table.RecurringPayments.AccountUUID)),
		).
		WHERE(
			table.PaymentAccounts.UserUUID.EQ(String(userUUID)).
				AND(table.RecurringPayments.UUID.EQ(String(uuid))),
		)

	var account platform.RecurringPayment
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func InsertRecurringPayment(ctx context.Context, db *sql.DB, account *platform.RecurringPayment) error {
	stmt := table.RecurringPayments.INSERT(
		table.RecurringPayments.AllColumns.Except(table.RecurringPayments.ID),
	).
		MODEL(account)

	_, err := stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRecurringPayment(ctx context.Context, db *sql.DB, uuid string) error {
	stmt := table.RecurringPayments.DELETE().
		WHERE(table.RecurringPayments.UUID.EQ(String(uuid)))

	_, err := stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
