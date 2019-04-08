package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func generateHMAC(secret string, data string) (result string) {

	hmacObj := hmac.New(sha256.New, []byte(secret))
	hmacObj.Write([]byte(data))
	result = hex.EncodeToString(hmacObj.Sum(nil))
	return
}
