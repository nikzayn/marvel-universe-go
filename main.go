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

	// Set gin mode to release to avoid logs clutter in terminal
	gin.SetMode(gin.ReleaseMode)

	// To initialize router with gin
	router := gin.New()

	// Trusting all proxies on default host
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Loading all HTML glob in template directory
	router.LoadHTMLGlob("templates/*")

	// Initialized gin logger
	router.Use(gin.Logger())

	// Route for main page
	routes.HomeRoute(router)

	// Route where you can search for data
	routes.SearchRoute(router)

	// Gin router is running on port 80 by default
	router.Run()
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
