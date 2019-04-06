package utils

import (
	"fmt"
)

const APIURL = "https://www.bitmex.com/api/v1/"

type Position struct {
	//something
}

func getPosition() {
	endpoint := "position"

	result := performApiCall(endpoint, APIURL)

	fmt.Printf("%v", result)

}
