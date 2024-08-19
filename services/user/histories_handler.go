package user

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) histories() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		from, to, accountUUID := s.historyManager.GetHistoryParams(c)
		histories, err := s.historyManager.GetHistories(c, user, accountUUID, *from, *to)
		if err != nil {
			c.String(500, "Failed to get histories")
			s.errorTracker.Report(c, err)
			return
		}

		c.JSON(200, gin.H{"histories": histories})
	}
}
