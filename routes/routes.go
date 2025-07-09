package routes

import (
	"github.com/arbeeorlar/controllers/module/auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	public := router.Group("/auth")
	{

		public.POST("/register", controllers.SignUpController)
	}

	//private := router.Group("/user")
	//private.Use(controllers.AuthenticationMiddleware())
	//{
	//	private.GET("/premium", controllers.PremiumController)
	//	private.GET("/signout", controllers.SignOutController)
	//	private.GET("/home", controllers.HomeController)
	//}

}
