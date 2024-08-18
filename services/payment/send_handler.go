package payment

import (
	"github.com/gin-gonic/gin"
)

type Send struct {
	AccountUUID string `json:"account" form:"account" binding:"required,uuid"`
	DestUUID    string `json:"dest" form:"dest" binding:"required,uuid"`
	Amount      int64  `json:"amount" form:"amount" binding:"required,max=9223372036854775807,min=-9223372036854775808"`
}

func (s *Service) send() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var send Send
		err := c.Bind(&send)
		if err != nil {
			return
		}

		err = s.accountManager.CheckOwner(c, user, send.AccountUUID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		from, to, err := s.accountManager.PrepareTransfer(c, send.AccountUUID, send.DestUUID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		history, err := s.paymentManager.Send(c, from, to, send.Amount)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, history)
	}
}
