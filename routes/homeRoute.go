package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/controllers"
)

func HomeRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", controllers.HomePage())
}
