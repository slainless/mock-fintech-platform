package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type SendPayload struct {
	AccountUUID string `json:"account" form:"account_id" binding:"required,uuid"`
	DestUUID    string `json:"dest" form:"dest_id" binding:"required,uuid"`
	Amount      int64  `json:"amount" form:"amount" binding:"required,max=999999999999999,min=1"`
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

		err = s.accountManager.CheckOwner(c, user, sourceUUID)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound:
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to check account")
			}
			return
		}

		from, to, err := s.accountManager.PrepareTransfer(c, sourceUUID, destUUID)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound, core.ErrInvalidTransferDestination:
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to prepare transfer")
			}
			return
		}

		history, err := s.paymentManager.Send(c, from, to, send.Amount)
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
