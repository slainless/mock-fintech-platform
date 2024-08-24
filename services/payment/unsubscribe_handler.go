package payment

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type UnsubscribePayload struct {
	PaymentUUID string `json:"payment_id" form:"payment_id" binding:"required,uuid"`
}

type UnsubscribeResponse struct {
	Transaction *platform.TransactionHistory `json:"transaction"`
	Status      string                       `json:"status"`
}

func (s *Service) unsubscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		var payload UnsubscribePayload
		err := c.ShouldBind(&payload)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		payment, err := s.recurringPaymentManager.GetPaymentWithAccess(c, user, uuid.MustParse(payload.PaymentUUID))
		if err != nil {
			switch {
			case errors.Is(err, core.ErrRecurringPaymentNotFound):
				c.String(400, err.Error())
			default:
				c.String(500, "Failed to get payment")
			}
			return
		}

		history, err := s.recurringPaymentManager.Unsubscribe(c, payment)
		if err != nil {
			c.String(500, "Failed to unsubscribe")
			return
		}

		c.JSON(200, UnsubscribeResponse{Transaction: history, Status: "ok"})
	}
}
