package core

import (
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type User struct {
	model *model.Users
}

func (u *User) ID() string {
	return u.model.UUID
}

func (u *User) Email() string {
	return u.model.Email
}

func UserFrom(model *model.Users) platform.User {
	return &User{model: model}
}
