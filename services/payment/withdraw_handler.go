package payment

import (
	"github.com/gin-gonic/gin"
)

type Withdraw struct {
	AccountUUID  string `json:"account" form:"account" binding:"required,uuid"`
	Amount       int64  `json:"amount" form:"amount" binding:"required,max=9223372036854775807,min=-9223372036854775808"`
	CallbackData string `json:"callback" form:"callback" binding:"required"`
}

func (s *Service) withdraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var withdraw Withdraw
		err := c.Bind(&withdraw)
		if err != nil {
			return
		}

		account, err := s.accountManager.GetAccount(c, withdraw.AccountUUID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if account.UserUUID != user.UUID {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid account"})
			return
		}

		history, err := s.paymentManager.Withdraw(c, account, withdraw.Amount, withdraw.CallbackData)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, history)
	}
}
