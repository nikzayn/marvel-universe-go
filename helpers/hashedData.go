package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

/* Hashed string function provides a hash to use the
marvel api with timestamp, private key and public key
*/
func HashedString(currentTime string, publicKey string) string {
	//Private key generated from Marvel developers account
	privateKey := os.Getenv("PRIVATE_KEY")

	/* So, we are using crypto package to hash our api keys and timestamp via using MD5 checksum
	and at last we are converting it to string to get the desired hash to call the api successfully
	*/
	hashMD5 := md5.New()
	hashMD5.Write([]byte(currentTime + privateKey + publicKey))
	return hex.EncodeToString(hashMD5.Sum(nil))
}
