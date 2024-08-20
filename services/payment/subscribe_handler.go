package payment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type SubscribePayload struct {
	AccountUUID  string `json:"account" form:"account_id" binding:"required,uuid"`
	ServiceID    string `json:"service" form:"service_id" binding:"required"`
	BillingID    string `json:"billing" form:"billing_id" binding:"required"`
	CallbackData string `json:"callback_data" form:"callback" binding:"required"`
}

type SubscribeResponse struct {
	Payment     *platform.RecurringPayment   `json:"payment"`
	Transaction *platform.TransactionHistory `json:"transaction"`
}

func (s *Service) subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		var payload SubscribePayload
		err := c.ShouldBind(&payload)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		account, err := s.accountManager.GetAccountWhereUser(c, user, payload.AccountUUID)
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

		payment, history, err := s.recurringPaymentManager.Subscribe(c, account, payload.ServiceID, payload.BillingID, payload.CallbackData)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrPaymentServiceNotSupported):
				c.String(501, err.Error())
			default:
				c.String(500, "Failed to subscribe")
			}
			return
		}

		c.JSON(201, SubscribeResponse{
			Payment:     payment,
			Transaction: history,
		})
	}
}
