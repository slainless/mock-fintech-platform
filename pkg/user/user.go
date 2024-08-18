package user

import "github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"

type User struct {
	model *model.Users
}

func (u *User) ID() string {
	return u.model.UUID
}

func (u *User) Email() string {
	return u.model.Email
}

func NewUser(model *model.Users) *User {
	return &User{model: model}
}
