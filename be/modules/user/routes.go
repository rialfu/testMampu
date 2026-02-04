package user

import (
	"rialfu/wallet/middlewares"
	"rialfu/wallet/modules/auth/service"
	"rialfu/wallet/modules/user/controller"
	"rialfu/wallet/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	userController := do.MustInvoke[controller.UserController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("", userController.GetAllUser)
		userRoutes.GET("/me", middlewares.Authenticate(jwtService), userController.Me)
		userRoutes.PUT("/:id", middlewares.Authenticate(jwtService), userController.Update)
	}
}
