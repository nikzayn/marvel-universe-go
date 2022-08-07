package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePage function to show html blob on browser
func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Marvel Freeverse",
		})

	}
}
