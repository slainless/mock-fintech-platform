package payment

import "github.com/gin-gonic/gin"

func (s *Service) Mount(r gin.IRouter) {
	my := r.Group("/")
	my.Use(s.authManager.Middleware(s.emailJwtAuth))

	my.POST("/send", s.send())
	my.POST("/withdraw", s.withdraw())
}
