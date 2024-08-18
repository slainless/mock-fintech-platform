package manager

import (
	"database/sql"
)

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

// func (m *UserManager) Login(ctx context.Context, service platform.AuthService, credential any) (platform.User, error) {
// 	uuid, err := service.Authenticate(credential)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user User
// 	err = query.GetUserInto(ctx, m.db, uuid, &user.user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
