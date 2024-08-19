package user

import "github.com/gin-gonic/gin"

func (s *Service) Mount(r gin.IRouter) {
	r.POST("/register", s.registerWithEmail())

	my := r.Group("/")
	my.Use(s.authManager.Middleware(s.emailJwtAuth))
	my.GET("/account", s.accounts())
	my.GET("/history", s.histories())

	my.POST("/account", s.create())
	my.GET("/account/:uuid", s.account())
}
