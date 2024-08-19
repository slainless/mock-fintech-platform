package payment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
)

type Withdraw struct {
	AccountUUID  string `json:"account_id" form:"account_id" binding:"required,uuid"`
	Amount       int64  `json:"amount" form:"amount" binding:"required"`
	CallbackData string `json:"callback" form:"callback" binding:"required"`
}

func (s *Service) withdraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var withdraw Withdraw
		err := c.ShouldBind(&withdraw)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		account, err := s.accountManager.GetAccount(c, withdraw.AccountUUID)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound):
				c.String(400, err.Error())
			default:
				s.errorTracker.Report(c, err)
				c.String(500, "Failed to get account")
			}
			return
		}

		if account.UserUUID != user.UUID {
			c.String(400, core.ErrAccountNotFound.Error())
			return
		}

		history, err := s.paymentManager.Withdraw(c, account, withdraw.Amount, withdraw.CallbackData)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrPaymentServiceNotSupported):
				c.String(501, err.Error())
			default:
				c.String(500, "Failed to withdraw")
				s.errorTracker.Report(c, err)
			}
			return
		}

		c.JSON(200, gin.H{"history": history})
	}
}
