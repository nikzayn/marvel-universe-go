package api

import (
	"os"
	"strconv"
	"time"

	"github.com/nikzayn/marvel-universe-go/helpers"
	"github.com/nikzayn/marvel-universe-go/models"
)

// Function to get all characters data
func GetAllCharacter(nameSearchParams, inputName, timeStampParams, currentOffset string) []models.MarvelData {

	// URL set to character url
	characterUrl := os.Getenv("CHARACTER_URL")
	// Public key
	publicKey := os.Getenv("PUBLIC_KEY")
	// Page limit by default set to 8
	pageLimit := "8"

	// To get current timestamp
	currentTime := strconv.Itoa(int(time.Now().Unix()))

	// To get the hashed value from current timestamp, private key and public key
	hashedValue := helpers.HashedString(currentTime, publicKey)

	return helpers.ResponseURL(characterUrl,
		nameSearchParams,
		inputName,
		timeStampParams,
		currentTime,
		publicKey,
		hashedValue,
		pageLimit,
		currentOffset)

}
