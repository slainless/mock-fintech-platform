package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type AccountResponse struct {
	Account *platform.PaymentAccountDetail `json:"account"`
	Balance *platform.MonetaryAmount       `json:"balance"`
}

func (s *Service) account() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		accountUUID, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			c.String(400, err.Error())
			return
		}

		account, err := s.accountManager.GetAccountDetail(c, user, accountUUID)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound):
				c.String(404, err.Error())
			default:
				c.String(500, "Failed to get account")
			}
			return
		}

		balance, err := s.accountManager.GetBalance(c, &account.PaymentAccount)
		if err != nil {
			c.String(500, "Failed to get account")
			s.errorTracker.Report(c, err)
			return
		}

		c.JSON(200, AccountResponse{Account: account, Balance: balance})
	}
}
