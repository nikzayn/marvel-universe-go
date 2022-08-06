package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/api"
)

func SearchName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Request.FormValue("search")
		currentOffset := c.Request.PostFormValue("submitId")
		nameSearchParams := "?name="
		timestampParams := "&ts="
		if len(name) == 0 {
			nameSearchParams = ""
			timestampParams = "?ts="
		}

		characterData := api.GetAllCharacter(name, nameSearchParams, timestampParams, currentOffset)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Marvel Freeverse",
			"data":  characterData,
		})
	}
}
