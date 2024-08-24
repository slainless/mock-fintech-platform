package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/core"
)

type AccountPermissionPayload struct {
	UserUUID   string   `json:"user_id" form:"user_id" binding:"required,uuid"`
	Permission []string `json:"permission" form:"permission" binding:"required,dive,oneof=all read history send withdraw subscription"`
}

func (s *Service) account_permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)
		var payload AccountPermissionPayload
		err := c.ShouldBind(&payload)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		perm, err := s.accountManager.ParsePermission(payload.Permission)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		accountUUID, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			c.String(400, err.Error())
			return
		}

		userUUID, err := uuid.Parse(payload.UserUUID)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		if user.UUID == userUUID {
			c.String(400, "You can't target yourself")
			return
		}

		err = s.accountManager.CheckOwner(c, user, accountUUID)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound):
				c.String(404, err.Error())
			default:
				c.String(500, "Failed to get account")
			}
			return
		}

		err = s.accountManager.SetPermission(c, uuid.MustParse(payload.UserUUID), uuid.MustParse(payload.UserUUID), perm)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrAccountNotFound), errors.Is(err, core.ErrUserNotRegistered):
				c.String(404, err.Error())
			default:
				c.String(500, "Failed to set permission")
			}
			return
		}

		c.Status(200)
	}
}
