package controllers

import (
	"fmt"
	"net/http"
	"time"

	// Gin middleware/handler to enable Cache

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/nikzayn/marvel-universe-go/api"
)

// To store the persisted in-memory via timestamp
var store = persistence.NewInMemoryStore(time.Second)

/* SearchName functionality is to get the input data from the user
and convert it to provide a new blob of HTML with meaningful data
*/
func SearchName() gin.HandlerFunc {

	// Added page cache to store the data with default timestamp
	// cache.CachePage(store, time.Minute, func(c *gin.Context) {})
	return func(c *gin.Context) {
		// Getting search value from web
		inputName := c.Request.FormValue("search")
		// Getting current offset from web
		currentOffset := c.Request.PostFormValue("submitId")

		/* A hacky way to enable search params or timeStamp params when there
		is no input coming from search name params or search bar
		*/
		nameSearchParams := "?name="
		timeStampParams := "&ts="
		if len(inputName) == 0 {
			nameSearchParams = ""
			timeStampParams = "?ts="
		}

		fmt.Println(currentOffset)

		characterData := api.GetAllCharacter(nameSearchParams, inputName, timeStampParams, currentOffset)

		// Updating view of HTML page with characterData
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Marvel Freeverse",
			"data":  characterData,
		})
	}
}
