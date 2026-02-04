package wallet

import (
	"rialfu/wallet/middlewares"
	"rialfu/wallet/modules/auth/service"
	"rialfu/wallet/modules/wallet/controller"
	"rialfu/wallet/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	controller := do.MustInvoke[controller.WalletController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	routes := server.Group("/api/wallet")
	{
		routes.GET("/balance", middlewares.Authenticate(jwtService), controller.CheckBalancee)
		routes.POST("/withdraw", middlewares.Authenticate(jwtService), controller.WithdrawProcess)
		routes.POST("/deposit", controller.StoreProcess)
	}
}
