package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goropencho/relay/controller"
)

func Routes() *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome User",
		})
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	routes := service.Group("/api")
	// authRoutes := routes.Group("/auth")
	// authRoutes.POST("/signup", authController.SignUp)
	// authRoutes.POST("/login", authController.Login)
	// authRoutes.POST("/logout", authController.Logout)
	user := controller.UserController{}
	userRoutes := routes.Group("/user")
	userRoutes.GET("", user.Me)
	userRoutes.PUT("", user.UpdateProfile)
	// mailRoutes := routes.Group("/mail")
	// mailRoutes.POST("", mailController.SendMail)

	return service
}
