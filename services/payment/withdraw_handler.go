package payment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type WithdrawPayload struct {
	AccountUUID  string `json:"account_id" form:"account_id" binding:"required,uuid"`
	Amount       int64  `json:"amount" form:"amount" binding:"required,max=999999999999999,min=1"`
	CallbackData string `json:"callback" form:"callback" binding:"required"`
}

type WithdrawResponse struct {
	Transaction *platform.TransactionHistory `json:"transaction"`
}

func (s *Service) withdraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var withdraw WithdrawPayload
		err := c.ShouldBind(&withdraw)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		account, err := s.accountManager.GetAccountWithAccess(c, user, uuid.MustParse(withdraw.AccountUUID), core.AccountPermissionWithdraw)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound):
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to get account")
			}
			return
		}

		if account.UserUUID != user.UUID {
			c.String(400, core.ErrAccountNotFound.Error())
			return
		}

		history, err := s.paymentManager.Withdraw(c, user, account, withdraw.Amount, withdraw.CallbackData)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrPaymentServiceNotSupported):
				c.String(501, err.Error())
			default:
				c.String(500, "Failed to withdraw")
			}
			return
		}

		c.JSON(200, WithdrawResponse{Transaction: history})
	}
}
