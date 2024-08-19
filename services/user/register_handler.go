package user

import (
	"github.com/gin-gonic/gin"
	"github.com/slainless/mock-fintech-platform/pkg/core"
)

type Register struct {
	Token string `json:"token" form:"token" binding:"required"`
}

func (s *Service) registerWithEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var register Register
		err := c.ShouldBind(&register)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		email, err := s.emailJwtAuth.Validate(c, register.Token)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		err = s.userManager.Register(c, email)
		if err != nil {
			switch err {
			case core.ErrUserAlreadyRegistered:
				c.String(409, err.Error())
			default:
				c.String(500, "Failed to register user")
				s.errorTracker.Report(c, err)
			}
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	}
}
