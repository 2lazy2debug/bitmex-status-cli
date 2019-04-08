package utils

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func getApiCall(endpoint string, url string) (call string) {
	var sb strings.Builder
	sb.WriteString(url)
	sb.WriteString(endpoint)
	call = sb.String()

	return
}

func performApiCall(endpoint string, url string) (result string) {
	req, err := http.NewRequest(http.MethodGet, getApiCall(url, endpoint), nil)
	if err != nil {
		panic(err)
	}

	expires := int32(time.Now().Unix()) + 10000
	keys := readFileLines("bitmex.apikey")



	req.Header.Add("api-expires", strconv.FormatInt(int64(expires), 10))
	req.Header.Add("api-key", keys["id"])

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	result = string(body)
	return
}

func generateSignature(secret string, reqType string, endpoint string, expires string)(signature string) {

	var sb strings.Builder
	sb.WriteString(reqType)
	sb.WriteString("/")
	sb.WriteString(endpoint)
	sb.WriteString(expires)
	//HEX(HMAC_SHA256(apiSecret, 'GET/api/v1/instrument1518064236'))
}
