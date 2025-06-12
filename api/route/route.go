package route

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/controller"
	"pledge-backend-test/api/middleware"
	"pledge-backend-test/config"
)

func InitRoute(e *gin.Engine) *gin.Engine {
	group := e.Group("/api/v" + config.Config.Env.Version)

	multiSignPoolController := controller.MultiSignPoolController{}
	group.POST("/pool/setMultiSign", multiSignPoolController.SetMultiSign)
	group.POST("/pool/getMultiSign", multiSignPoolController.GetMultiSign)

	poolController := controller.PoolController{}
	group.POST("/poolBaseInfo", poolController.SelectPoolBaseInfo)
	group.POST("/poolDataInfo", poolController.SelectPoolDataInfo)
	group.POST("/tokenList", poolController.SelectTokenList)
	group.POST("/pool/search", middleware.CheckToken(), poolController.Search)
	group.POST("/pool/debtTokenInfo", middleware.CheckToken(), poolController.DebtTokenList)

	userController := controller.UserController{}
	group.POST("/user/login", userController.Login)
	group.POST("/user/logout", middleware.CheckToken(), userController.Logout)

	priceController := controller.PriceController{}
	group.POST("/price", priceController.NewPrice)
	return e
}
