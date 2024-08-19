package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type AccountsResponse struct {
	Accounts []platform.PaymentAccount `json:"accounts"`
}

func (s *Service) accounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		accounts, err := s.accountManager.GetAccounts(c, user)
		if err != nil {
			c.String(500, "Failed to load user accounts")
			s.errorTracker.Report(c, err)
			return
		}

		c.JSON(200, AccountsResponse{Accounts: accounts})
	}
}
