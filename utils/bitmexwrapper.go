package bitmexwrapper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const APIURL = "https://www.bitmex.com/api/v1/"

type Position struct {
	//something
}

func getPosition() {
	call := getApiCall("position")
	result := performApiCall(call)

	fmt.Printf("%v", result)

}

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
