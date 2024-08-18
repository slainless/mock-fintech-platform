package user

import "github.com/gin-gonic/gin"

func (s *Service) Mount(r gin.IRouter) {
	r.POST("/register", register(s.supabaseJwtAuth, s.userManager))

	my := r.Group("/my")
	my.Use(s.authManager.Middleware(s.supabaseJwtAuth))
	my.GET("/accounts", accounts(s.authManager, s.accountManager))
	my.GET("/history", histories(s.authManager, s.historyManager))
}
