package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/nikzayn/marvel-universe-go/controllers"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
)

func TestPingRoute(t *testing.T) {
	mockResponse := `<!doctype html>
    <html lang="en">
      <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Marvel</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js" integrity="sha512-K1qjQ+NcF2TYO/eI3M6v8EiNYZfA95pQumfvcVrTHtwQVDG+aHRqLi/ETn2uB+1JqwYqVG3LIvdm9lj6imS/pQ==" crossorigin="anonymous"></script>
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-A3rJD856KowSb7dwlZdYEkO39Gagi7vIsF0jrRAoQmDKKtQBHUuLZ9AsSv4jD4Xa" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js" integrity="sha512-894YE6QWD5I59HgZOGReFYm4dnWc1Qt5NtvYSaNcOP+u1T9qYdvdihz0PPSiiqn/+/3e7Jo4EaG7TubfWGUrMQ==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
      </head>
      <style>
        .heading {
          display: flex;
          justify-content: center;
          top: 20px;
        }
    
        .card {
          margin: 10px;
        }
      </style>
      <body>
        <h1 class="heading">Marvel Freeverse</h1>
        <section class="container text-center">
          <form action="/search" method="post">
            <input type="search" name="search">
            <input type="submit" value="Search">
          </form>
          <br>
          <div class="flex-row">
          <form action="/search" method="post">
            <input type="hidden" id="prev" value="0"/>
            <input type="button" id="prevButton" onclick="prevPageValue(this)" value="Prev" disabled/>
            <input type="submit" id="submitId" value="0" name="submitId">
            <input type="hidden" id="next" value="0"/>
            <input type="button" id="button" onclick="nextPageValue()" value="Next" />
          </form>
          </div>
        </form>
          <br>
          <div class="row justify-content-center">
            
        </div>
        <br>
        <br>
        </section>
      </body>
        <script>
    
          
          
          
          
          
          limit = 8;
    page_no = 1;
    
    let value;
    let buttonVal;
    
    function prevPageValue(btn) {
      if (parseInt(value) == 0) {
        document.getElementById(btn.id).disabled = true;
      }
      value = parseInt(document.getElementById('prev').value, 10);
      value = ((page_no - 1) - 1) * limit;
      page_no -= 1
      if (parseInt(value) == 0) {
        document.getElementById(btn.id).disabled = true;
        buttonVal = btn;
      }
      document.getElementById('prev').value = value;
      document.getElementById('submitId').value = value;
      console.log(document.getElementById('prev').value);
    }
    
    function nextPageValue() {
      document.getElementById('prevButton').disabled = false;
      if (parseInt(value) === 0) {
        document.getElementById(buttonVal.id).disabled = false;
      }
      value = parseInt(document.getElementById('next').value, 10);
      value = ((page_no + 1) - 1) * limit;
      page_no += 1;
      document.getElementById('next').value = value;
      document.getElementById('submitId').value = value;
      console.log(document.getElementById('next').value);
    }
    
    
      </script>
    </html>%`
	r := SetUpRouter()
	r.GET("/", controllers.HomePage())
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Marvel Freeverse",
		})
	})

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, responseData)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRespondsWithLove(t *testing.T) {

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	resource, err := pool.Run("docker-marvel", "latest", []string{})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8080/tcp"), "/"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.Contains(t, string(body), "<3", "does not respond with love?")
}
