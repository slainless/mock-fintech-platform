package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
)

func histories(authManager *core.AuthManager, historyManager *core.TransactionHistoryManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := authManager.GetUser(c)

		from, to, err := historyManager.GetRange(c)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		histories, err := historyManager.GetHistories(c, user, *from, *to)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, histories)
	}
}
