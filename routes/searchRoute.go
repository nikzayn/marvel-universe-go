package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/controllers"
)

//Route to HomePage - http://localhost:8080/search
func SearchRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/:search", controllers.SearchName())
}
