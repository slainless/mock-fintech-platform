package core

import (
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
)

type User struct {
	user model.Users
}

func (u *User) ID() string {
	return u.user.UUID
}
