package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateHMAC(secret string, data string) (result string) {

	hmacObj := hmac.New(sha256.New, []byte(secret))
	fmt.Println("data : " + data)
	hmacObj.Write([]byte(data))
	result = hex.EncodeToString(hmacObj.Sum(nil))
	return
}
