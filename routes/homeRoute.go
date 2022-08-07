package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/controllers"
)

//Route to HomePage - http://localhost:8080
func HomeRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", controllers.HomePage())
}
