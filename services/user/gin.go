package user

import "github.com/gin-gonic/gin"

func (s *UserService) Mount(r gin.IRouter) {
	r.POST("/register", register(s.supabaseJwtAuth, s.userManager))

	my := r.Group("/my")
	my.Use(s.authManager.Middleware(s.supabaseJwtAuth))
}
