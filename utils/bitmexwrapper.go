package utils

const APIURL = "https://www.bitmex.com"

type Position struct {
	//something
}

func GetPosition() (result string) {
	endpoint := "/api/v1/position"

	keys := readFileLines("bitmex.apikey")

	result = performAPICall(keys, "GET", endpoint, APIURL)
	return

}
