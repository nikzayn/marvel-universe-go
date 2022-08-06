package api

import (
	"os"
	"strconv"
	"time"

	"github.com/nikzayn/marvel-universe-go/helpers"
	"github.com/nikzayn/marvel-universe-go/models"
)

func GetAllCharacter(name, searchName, timeStampParams, offset string) []models.MarvelData {

	characterUrl := os.Getenv("CHARACTER_URL")
	publicKey := os.Getenv("PUBLIC_KEY")

	// To get current timestamp
	currentTime := strconv.Itoa(int(time.Now().Unix()))

	// Checking if are getting valid hash value
	hashedValue := helpers.HashedString(currentTime, publicKey)

	return helpers.ResponseURL(characterUrl, searchName, name, timeStampParams, offset, currentTime, publicKey, hashedValue)

}
