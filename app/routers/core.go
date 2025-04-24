package routers

import (
	"github.com/gin-gonic/gin"
	"kryptonim-interview/app"
	"kryptonim-interview/app/handlers"
)

func SetupCore(container *app.Container) *gin.Engine {
	r := gin.Default()

	handlers.NewCoreHandler(
		container.NewExchangeRatesService(),
	).RegisterRoutes(r)

	return r
}
