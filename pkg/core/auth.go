package core

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var (
	ErrUserNotRegistered = errors.New("user not registered")
)

type AuthManager struct {
	UserManager *UserManager
}

func NewAuthManager(userManager *UserManager) *AuthManager {
	return &AuthManager{
		UserManager: userManager,
	}
}

func (m *AuthManager) Validate(ctx context.Context, service platform.AuthService, credential any) (*platform.User, error) {
	email, err := service.Validate(ctx, credential)
	if err != nil {
		return nil, err
	}

	user, err := m.UserManager.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *AuthManager) Middleware(service platform.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		credential, err := service.Credential(c)
		if err != nil {
			c.String(401, err.Error())
			c.Abort()
			return
		}

		user, err := m.Validate(c, service, credential)
		if err != nil {
			switch err {
			case ErrUserNotRegistered:
				c.String(401, err.Error()+"\nPlease register your account first at /register.")
			default:
				c.String(401, err.Error())
			}
			c.Abort()
			return
		}

		m.SetUser(c, user)
	}
}

func (m *AuthManager) SetUser(c *gin.Context, user *platform.User) {
	c.Set("__auth_manager_user", user)
}

func (m *AuthManager) GetUser(c *gin.Context) *platform.User {
	user, ok := c.Get("__auth_manager_user")
	if !ok {
		return nil
	}

	if u, ok := user.(*platform.User); ok {
		return u
	} else {
		return nil
	}
}
