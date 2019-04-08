package bitmex

import (
	"bitmex-status-cli/io"
	"bitmex-status-cli/network"
)

const APIURL = "https://www.bitmex.com"

func GetPosition() (result string) {
	endpoint := "/api/v1/position"

	keys := io.ReadKeyFile("bitmex.apikey")

	result = network.PerformAPICall(keys, "GET", endpoint, APIURL)
	return

}
