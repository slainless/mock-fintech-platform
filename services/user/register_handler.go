package user

import (
	"github.com/gin-gonic/gin"
)

type Register struct {
	Token string `json:"token" form:"token" binding:"required"`
}

func (s *Service) registerWithSupabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		var register Register
		err := c.Bind(&register)
		if err != nil {
			return
		}

		email, err := s.supabaseJwtAuth.Validate(c, register.Token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		err = s.userManager.Register(c, email)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}
