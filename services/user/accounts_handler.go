package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
)

func accounts(authManager *core.AuthManager, accountManager *core.PaymentAccountManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := authManager.GetUser(c)

		accounts, err := accountManager.GetAccounts(c, user)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, accounts)
	}
}
