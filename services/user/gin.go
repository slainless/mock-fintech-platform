package user

import "github.com/gin-gonic/gin"

func (s *UserService) Mount(r gin.IRouter) {

	my := r.Group("/my")
	my.Use(s.authManager.Middleware(s.supabaseJwtAuth))
}

// func MountService(r gin.IRouter, service IUserService) {
// 	r.Use(gin.Logger())
// 	r.Use(gin.Recovery())

// 	r.POST("/register", register())
// 	r.POST("/login")
// 	r.POST("/revoke")

// 	my := r.Group("/my")
// 	my.Use(service.UserManager().AuthWall())
// 	my.GET("/accounts")
// 	my.GET("/history")
// }
