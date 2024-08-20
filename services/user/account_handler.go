package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type AccountResponse struct {
	Account *platform.PaymentAccount `json:"account"`
	Balance *platform.MonetaryAmount `json:"balance"`
}

func (s *Service) account() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		account, err := s.accountManager.GetAccountWhereUser(c, user.UUID, c.Param("uuid"))
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound):
				c.String(404, err.Error())
			default:
				c.String(500, "Failed to get account")
			}
			return
		}

		if account.UserUUID != user.UUID {
			c.String(404, core.ErrAccountNotFound.Error())
			// c.String(403, "Forbidden")
			return
		}

		balance, err := s.accountManager.GetBalance(c, account)
		if err != nil {
			c.String(500, "Failed to get account")
			s.errorTracker.Report(c, err)
			return
		}

		c.JSON(200, AccountResponse{Account: account, Balance: balance})
	}
}
