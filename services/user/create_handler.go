package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type Create struct {
	// foreign account id, different from internal account UUID.
	AccountID    string `json:"account_id" form:"account_id" binding:"required"`
	ServiceID    string `json:"service_id" form:"service_id" binding:"required"`
	Name         string `json:"name" form:"name" binding:""`
	CallbackData string `json:"callback" form:"callback" binding:"required"`
}

func (s *Service) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var create Create
		err := c.ShouldBind(&create)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		account, err := s.accountManager.Register(c, user, create.ServiceID, create.Name, create.AccountID, create.CallbackData)
		if err != nil {
			switch {
			case errors.Is(err, platform.ErrInvalidAccountData):
				// errors.Is(err, platform.ErrTransactionRejected):
				c.String(400, err.Error())
			case err == core.ErrAccountAlreadyRegistered:
				c.String(409, err.Error())
			case errors.Is(err, core.ErrPaymentServiceNotSupported):
				c.String(501, err.Error())
			default:
				s.errorTracker.Report(c, err)
				c.String(500, "Failed to register account")
			}
			return
		}

		balance, err := s.accountManager.GetBalance(c, account)
		if err != nil {
			s.errorTracker.Report(c, err)
			// c.String(500, "Failed to get account balance post-registration\nBut, don't worry, your account is successfully created")
			return
		}

		c.JSON(201, gin.H{"account": account, "balance": balance})
	}
}
