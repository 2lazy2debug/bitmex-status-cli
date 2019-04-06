package utils

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getApiCall(endpoint string) (call string) {
	var sb strings.Builder
	sb.WriteString(APIURL)
	sb.WriteString(endpoint)
	call = sb.String()

	return
}

func performApiCall(call string) (result string) {
	req, err := http.NewRequest(http.MethodGet, call, nil)
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

func generateSignature() {

}
