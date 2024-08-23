package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type HistoriesResponse struct {
	Histories []platform.TransactionHistory `json:"histories"`
}

func (s *Service) histories() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		from, to, accountUUID := s.historyManager.GetHistoryParams(c)
		var account *platform.PaymentAccount
		if accountUUID != "" {
			acc, err := s.accountManager.GetAccountWithAccess(c, user, uuid.MustParse(accountUUID), core.AccountPermissionHistory)
			if err != nil {
				switch {
				case errors.Is(err, core.ErrAccountNotFound):
					c.String(400, core.ErrAccountNotFound.Error())
				default:
					c.String(500, "Failed to get account")
				}
				return
			}

			account = acc
		}

		// TODO: introduce atomicity to this operation
		histories, err := s.historyManager.GetHistories(c, user, account, *from, *to)
		if err != nil {
			c.String(500, "Failed to get histories")
			return
		}

		c.JSON(200, HistoriesResponse{Histories: histories})
	}
}
