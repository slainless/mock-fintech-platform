package platform

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UserManager interface {
	Authenticate(ctx context.Context, service AuthService, credential any) (User, error)
	Register(ctx context.Context, email string) (User, error)

	Revoke(ctx context.Context, service AuthService, credential any) error

	AuthWall() gin.HandlerFunc
}
