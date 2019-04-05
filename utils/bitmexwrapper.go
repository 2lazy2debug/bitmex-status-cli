package bitmexwrapper

import "strings"

const APIURL = "https://www.bitmex.com/api/v1/"

type Position struct {
	//something
}

func getPosition() {

	var sb strings.Builder
	endpoint := "position"

	sb.WriteString(APIURL)
	sb.WriteString(endpoint)

}
