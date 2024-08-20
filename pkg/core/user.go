package core

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var ErrUserAlreadyRegistered = errors.New("user already registered")

type UserManager struct {
	db *sql.DB

	errorTracker platform.ErrorTracker
}

func NewUserManager(db *sql.DB, errorTracker platform.ErrorTracker) *UserManager {
	return &UserManager{
		db: db,

		errorTracker: errorTracker,
	}
}

func (m *UserManager) GetUserByEmail(ctx context.Context, email string) (*platform.User, error) {
	user, err := query.GetUser(ctx, m.db, email)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrUserNotRegistered
		}
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return user, nil
}

func (m *UserManager) Register(ctx context.Context, email string) error {
	err := query.InsertFreshUser(ctx, m.db, email)
	if err != nil {
		if err := util.PQErrorCode(err); err != "" {
			switch err {
			case "unique_violation":
				return ErrUserAlreadyRegistered
			}
		}

		m.errorTracker.Report(ctx, err)
		return err
	}

	return nil
}
