package user

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) histories() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		from, to, err := s.historyManager.GetRange(c)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		histories, err := s.historyManager.GetHistories(c, user, *from, *to)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, histories)
	}
}
