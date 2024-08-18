package user

import "github.com/gin-gonic/gin"

type Create struct {
	AccountID    string `json:"account_id" form:"account_id" binding:"required"`
	ServiceID    string `json:"service_id" form:"service_id" binding:"required"`
	Name         string `json:"name" form:"name" binding:""`
	CallbackData string `json:"callback" form:"callback" binding:"required"`
}

func (s *Service) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := s.authManager.GetUser(c)

		var create Create
		err := c.Bind(&create)
		if err != nil {
			return
		}

		account, err := s.accountManager.Register(c, user, create.AccountID, create.ServiceID, create.Name, create.CallbackData)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		balance, err := s.accountManager.GetBalance(c, account)
		if err != nil {
			s.errorTracker.Report(c, err)
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"account": account, "balance": balance})
	}
}
