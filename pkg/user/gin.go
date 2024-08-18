package user

import "github.com/gin-gonic/gin"

func MountService(r gin.IRouter, service IUserService) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login")
	r.POST("/logout")

	my := r.Group("/my")
	my.Use(service.UserManager().AuthWall())
	my.GET("/accounts")
	my.GET("/history")
}
