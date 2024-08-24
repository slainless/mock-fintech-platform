package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func GetRecurringPayment(ctx context.Context, db *sql.DB, uuid uuid.UUID) (*platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns).
		FROM(table.RecurringPayments).
		WHERE(table.RecurringPayments.UUID.EQ(UUID(uuid)))

	var account platform.RecurringPayment
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetRecurringPaymentWithAccess(ctx context.Context, db *sql.DB, userUUID, uuid uuid.UUID, access AccountPermission) (*platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns).
		FROM(
			table.RecurringPayments.
				INNER_JOIN(table.PaymentAccounts, table.PaymentAccounts.UUID.EQ(table.RecurringPayments.AccountUUID)).
				INNER_JOIN(table.SharedAccountAccess, table.RecurringPayments.AccountUUID.EQ(table.SharedAccountAccess.AccountUUID)),
		).
		WHERE(
			table.RecurringPayments.UUID.EQ(UUID(uuid)).
				AND(OR(
					table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)),
					table.SharedAccountAccess.UserUUID.EQ(UUID(userUUID)).
						AND(table.SharedAccountAccess.Permission.BIT_AND(Int32(int32(access))).EQ(Int32(int32(access)))),
				)),
		).
		GROUP_BY(table.RecurringPayments.UUID)

	var account platform.RecurringPayment
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetRecurringPayments(ctx context.Context, db *sql.DB, accountUUID uuid.UUID) ([]platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns).
		FROM(table.RecurringPayments).
		WHERE(table.RecurringPayments.AccountUUID.EQ(UUID(accountUUID)))

	accounts := make([]platform.RecurringPayment, 0)
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetRecurringPaymentsOfUser(ctx context.Context, db *sql.DB, userUUID uuid.UUID) ([]platform.RecurringPayment, error) {
	stmt := SELECT(table.RecurringPayments.AllColumns).
		FROM(
			table.RecurringPayments.
				INNER_JOIN(table.PaymentAccounts, table.PaymentAccounts.UUID.EQ(table.RecurringPayments.AccountUUID)),
		).
		WHERE(
			table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)),
		)

	accounts := make([]platform.RecurringPayment, 0)
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func InsertRecurringPayment(ctx context.Context, db *sql.DB, account *platform.RecurringPayment) error {
	stmt := table.RecurringPayments.INSERT(table.RecurringPayments.MutableColumns).
		MODEL(account).
		RETURNING(table.RecurringPayments.AllColumns)

	err := stmt.QueryContext(ctx, db, account)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRecurringPayment(ctx context.Context, db *sql.DB, uuid uuid.UUID) error {
	stmt := table.RecurringPayments.DELETE().
		WHERE(table.RecurringPayments.UUID.EQ(UUID(uuid)))

	_, err := stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
