package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type SubscriptionParams struct {
	AccountUUID *string `json:"account_id" form:"account_id" binding:"uuid"`
}

type SubscriptionResponse struct {
	Payments []platform.RecurringPayment `json:"subscriptions"`
}

func (s *Service) subscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		var params SubscriptionParams
		err := c.ShouldBind(&params)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		var account *platform.PaymentAccount
		if params.AccountUUID != nil {
			account, err = s.accountManager.GetAccountWithAccess(c, user, uuid.MustParse(*params.AccountUUID), core.AccountPermissionRead)
			if err != nil {
				switch {
				case errors.Is(err, core.ErrAccountNotFound):
					c.String(404, err.Error())
				default:
					c.String(500, "Failed to get account")
				}
				return
			}
		}

		payments, err := s.recurringPayments.GetPayments(c, user, account)
		if err != nil {
			switch {
			default:
				c.String(500, "Failed to get subscription")
			}
			return
		}

		c.JSON(200, SubscriptionResponse{Payments: payments})
	}
}
