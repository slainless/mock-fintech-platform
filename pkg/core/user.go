package core

import (
	"context"
	"database/sql"

	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) GetUserByEmail(ctx context.Context, email string) (platform.User, error) {
	model, err := query.GetUser(ctx, m.db, email)
	if err != nil {
		return nil, err
	}

	return UserFrom(model), nil
}

func (m *UserManager) Register(ctx context.Context, email string) error {
	return query.InsertFreshUser(ctx, m.db, email)
}
