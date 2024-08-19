package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
	"golang.org/x/crypto/bcrypt"
)

func GetUserInto(ctx context.Context, db *sql.DB, email string, user *platform.User) error {
	stmt := SELECT(
		table.Users.UUID,
		table.Users.FullName,
		table.Users.UserName,
	).
		FROM(table.Users).
		WHERE(table.Users.Email.EQ(String(email)))

	err := stmt.QueryContext(ctx, db, user)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(ctx context.Context, db *sql.DB, email string) (*platform.User, error) {
	var user platform.User
	err := GetUserInto(ctx, db, email, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Authenticate(ctx context.Context, db *sql.DB, email string, password []byte) (string, error) {
	stmt := SELECT(
		table.Users.PasswordHash,
		table.Users.UUID,
	).
		FROM(table.Users).
		WHERE(table.Users.Email.EQ(String(email)))

	var user model.Users
	err := stmt.QueryContext(ctx, db, &user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), password)
	if err != nil {
		return "", err
	}

	return user.UUID, nil
}

func InsertFreshUser(ctx context.Context, db *sql.DB, email string) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	stmt := table.Users.INSERT(
		table.Users.Email,
		table.Users.UUID,
	).
		VALUES(String(email), String(uuid.String()))

	_, err = stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
