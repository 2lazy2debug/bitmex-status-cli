package network

import (
	"bitmex-status-cli/io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PerformAPICall(keys map[string]string, reqType string, endpoint string, url string) (result string) {
	req, err := http.NewRequest(http.MethodGet, getAPICall(endpoint, url), nil)
	if err != nil {
		panic(err)
	}

	expires := strconv.FormatInt(int64(int32(time.Now().Unix())+10000), 10)
	//expires := strconv.FormatInt(int64(1518064236), 10)
	sig := generateSignature(keys["secret"], reqType, endpoint, expires)

	req.Header.Add("api-expires", expires)
	req.Header.Add("api-key", keys["id"])
	req.Header.Add("api-signature", sig)
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

func getAPICall(endpoint string, url string) (call string) {
	var sb strings.Builder
	sb.WriteString(url)
	sb.WriteString(endpoint)
	call = sb.String()

	return
}

func generateSignature(secret string, reqType string, endpoint string, expires string) (signature string) {

	var sb strings.Builder
	sb.WriteString(reqType)
	sb.WriteString(endpoint)
	sb.WriteString(expires)

	signature = io.GenerateHMAC(secret, sb.String())
	return
}
