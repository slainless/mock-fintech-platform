package platform

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	ServiceID() string
	Validate(ctx context.Context, credential any) (email string, err error)
	Credential(ctx *gin.Context) (any, error)
}
