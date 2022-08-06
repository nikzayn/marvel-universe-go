package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nikzayn/marvel-universe-go/routes"
)

func main() {

	// To load environment variables
	err := godotenv.Load()

	// Sanity check if there are not environmental variables
	if err != nil {
		fmt.Println(err)
	}

	// To initialize router with gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.LoadHTMLGlob("templates/*")
	router.Use(gin.Logger())

	routes.HomeRoute(router)
	routes.SearchRoute(router)

	router.Run()
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
