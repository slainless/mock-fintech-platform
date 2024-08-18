package user

import "github.com/gin-gonic/gin"

func (s *Service) Mount(r gin.IRouter) {
	r.POST("/register", register(s.supabaseJwtAuth, s.userManager))

	my := r.Group("/my")
	my.Use(s.authManager.Middleware(s.supabaseJwtAuth))
	my.GET("/account", s.accounts())
	my.GET("/history", s.histories())

	my.POST("/account", s.create())
	my.GET("/account/:uuid", s.account())
}
