package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type HistoriesResponse struct {
	Histories []platform.TransactionHistory `json:"histories"`
}

func (s *Service) histories() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		from, to, accountUUID := s.historyManager.GetHistoryParams(c)
		histories, err := s.historyManager.GetHistories(c, user, accountUUID, *from, *to)
		if err != nil {
			c.String(500, "Failed to get histories")
			return
		}

		c.JSON(200, HistoriesResponse{Histories: histories})
	}
}
