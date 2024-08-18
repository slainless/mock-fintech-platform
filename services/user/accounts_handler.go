package user

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) accounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		accounts, err := s.accountManager.GetAccounts(c, user)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, accounts)
	}
}
