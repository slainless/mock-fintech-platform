package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
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
		err := c.ShouldBind(&send)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		err = s.accountManager.CheckOwner(c, user, send.AccountUUID)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound:
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to check account")
				s.errorTracker.Report(c, err)
			}
			return
		}

		from, to, err := s.accountManager.PrepareTransfer(c, send.AccountUUID, send.DestUUID)
		if err != nil {
			switch err {
			case core.ErrAccountNotFound, core.ErrInvalidTransferDestination:
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to prepare transfer")
				s.errorTracker.Report(c, err)
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
				s.errorTracker.Report(c, err)
			}
			return
		}

		c.JSON(200, history)
	}
}
