package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type SendPayload struct {
	AccountUUID  string `json:"account_id" form:"account_id" binding:"required,uuid"`
	DestUUID     string `json:"dest_id" form:"dest_id" binding:"required,uuid"`
	Amount       int64  `json:"amount" form:"amount" binding:"required,max=999999999999999,min=1"`
	CallbackData string `json:"callback" form:"callback"`
}

type SendResponse struct {
	Transaction *platform.TransactionHistory `json:"transaction"`
}

func (s *Service) send() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var send SendPayload
		err := c.ShouldBind(&send)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		if send.Amount <= 0 {
			c.String(400, core.ErrInvalidAmount.Error())
			return
		}

		if send.AccountUUID == send.DestUUID {
			c.String(400, core.ErrInvalidTransferDestination.Error()+"\nUnable to send to the same account!")
			return
		}

		sourceUUID := uuid.MustParse(send.AccountUUID)
		destUUID := uuid.MustParse(send.DestUUID)

		from, err := s.accountManager.GetAccountWithAccess(c, user, sourceUUID, core.AccountPermissionSend)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound:
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to check account")
			}
			return
		}

		to, err := s.accountManager.GetAccount(c, destUUID)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound:
				c.String(400, core.ErrInvalidTransferDestination.Error())
			default:
				c.String(500, "Failed to get destination account")
			}
			return
		}

		history, err := s.paymentManager.Send(c, user, from, to, send.Amount, send.CallbackData)
		if err != nil {
			switch err {
			case core.ErrPaymentServiceNotSupported:
				c.String(501, err.Error())
			default:
				c.String(500, "Failed to send")
			}
			return
		}

		c.JSON(200, SendResponse{Transaction: history})
	}
}
