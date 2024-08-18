package core

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type AuthManager struct {
	UserManager *UserManager
}

func (m *AuthManager) Validate(ctx context.Context, service platform.AuthService, credential any) (*platform.User, error) {
	email, err := service.Validate(ctx, credential)
	if err != nil {
		return nil, err
	}

	return m.UserManager.GetUserByEmail(ctx, email)
}

func (m *AuthManager) Middleware(service platform.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		credential, err := service.Credential(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		user, err := m.Validate(c, service, credential)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
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
