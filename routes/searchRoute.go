package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/controllers"
)

func SearchRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/:search", controllers.SearchName())
}
