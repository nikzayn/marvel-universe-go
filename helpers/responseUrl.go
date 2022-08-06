package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/nikzayn/marvel-universe-go/models"
)

// var cache = memcache.New("localhost:11211")

func ResponseURL(urlVal, searchName, name, timeStampParams, offset, currentTime, publicKey, hashedValue string) []models.MarvelData {

	responseURL, err := http.Get(
		urlVal +
			searchName +
			url.QueryEscape(name) +
			timeStampParams +
			currentTime +
			"&apikey=" + publicKey +
			"&hash=" + hashedValue +
			"&limit=" + "8" +
			"&offset=" + offset)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer responseURL.Body.Close()
	responseData, err := ioutil.ReadAll(responseURL.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Trying to cache api here using memcached
	// cacheErr := cache.Set(&memcache.Item{Key: "", Value: responseData, Expiration: 10})
	// if cacheErr != nil {
	// 	panic(cacheErr)
	// }

	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Data.MarvelData
}
