package io

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC(secret string, data string) (result string) {

	hmacObj := hmac.New(sha256.New, []byte(secret))
	hmacObj.Write([]byte(data))
	result = hex.EncodeToString(hmacObj.Sum(nil))
	return
}
