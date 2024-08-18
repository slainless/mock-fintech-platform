package user

import "github.com/gin-gonic/gin"

func (s *Service) account() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		account, err := s.accountManager.GetAccount(c, c.Param("uuid"))
		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
			return
		}

		if account.UserUUID != user.UUID {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid account"})
			return
		}

		balance, err := s.accountManager.GetBalance(c, account)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"account": account, "balance": balance})
	}
}
