package utils

import (
	"fmt"
)

const APIURL = "https://www.bitmex.com"

type Position struct {
	//something
}

func getPosition() {
	endpoint := "/api/v1/position"

	keys := readFileLines("bitmex.apikey")

	result := performAPICall(keys, "GET", endpoint, APIURL)

	fmt.Printf("%v", result)

}
