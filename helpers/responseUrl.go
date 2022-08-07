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

// This is our helper function to get the data from api
func ResponseURL(urlVal,
	nameSearchParams,
	inputName,
	timeStampParams,
	currentTime,
	publicKey,
	hashedValue,
	pageLimit,
	currentOffset string) []models.MarvelData {

	responseURL, err := http.Get(
		urlVal +
			nameSearchParams +
			url.QueryEscape(inputName) +
			timeStampParams +
			currentTime +
			"&apikey=" + publicKey +
			"&hash=" + hashedValue +
			"&limit=" + pageLimit +
			"&offset=" + currentOffset)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Closing all calls from http.Get function after getting data successfully.
	// It will run after the function completion.
	defer responseURL.Body.Close()
	responseData, err := ioutil.ReadAll(responseURL.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.Response
	/* We are using json unmarshal to decode the data which is coming in the form
	of JSON type and converting it to struct
	*/
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Data.MarvelData
}
