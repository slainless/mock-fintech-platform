package platform

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UserManager interface {
	Login(ctx context.Context, service AuthService, credential any) (User, error)

	Logout(ctx context.Context, service AuthService, user User) error
	LogoutByID(ctx context.Context, service AuthService, userID string) error

	AuthWall() gin.HandlerFunc
}
