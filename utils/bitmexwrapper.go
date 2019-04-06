package utils

import (
	"fmt"
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
