package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"golang.org/x/crypto/bcrypt"
)

func GetUserInto(ctx context.Context, db *sql.DB, email string, user *model.Users) error {
	stmt := SELECT(
		table.Users.UUID,
		table.Users.FullName,
		table.Users.UserName,
	).
		FROM(table.Users).
		WHERE(table.Users.Email.EQ(String(email)))

	err := stmt.QueryContext(ctx, db, &user)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(ctx context.Context, db *sql.DB, email string) (*model.Users, error) {
	var user model.Users

	err := GetUserInto(ctx, db, email, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Authenticate(
	ctx context.Context,
	db *sql.DB,
	email string,
	password []byte,
) (string, error) {
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

// func InsertUser(ctx context.Context, db *sql.DB, user *UserModel) error {}
