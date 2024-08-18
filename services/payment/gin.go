package payment

import "github.com/gin-gonic/gin"

func (s *Service) Mount(r gin.IRouter) {
	my := r.Group("/")
	my.Use(s.authManager.Middleware(s.emailJwtAuth))

	my.GET("/send", s.send())
	my.GET("/withdraw", s.withdraw())
}
