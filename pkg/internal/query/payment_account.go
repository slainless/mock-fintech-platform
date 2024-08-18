package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
)

func GetAllAccounts(ctx context.Context, db *sql.DB, userUUID string) ([]model.PaymentAccounts, error) {
	stmt := SELECT(
		table.PaymentAccounts.UUID,
		table.PaymentAccounts.ServiceID,
		table.PaymentAccounts.Balance,
		table.PaymentAccounts.Currency,
	).
		FROM(table.PaymentAccounts).
		WHERE(table.PaymentAccounts.UserUUID.EQ(String(userUUID)))

	var accounts []model.PaymentAccounts
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
