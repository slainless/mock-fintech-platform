package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func register(service platform.AuthService, manager *core.UserManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := service.Validate(c, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		err = manager.Register(c, email)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}
