// Copyright (c) 2019 Manuel Cabras
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"bitmex-status-cli/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//var c config.Conf

const BITMEXREFLINK = "https://www.bitmex.com/register/1TTdTj"

type bitcoinTicker struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

func bitcoinToDollar() {
	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/btcusd")
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	valueBTC := bitcoinTicker{}
	err = json.Unmarshal(body, &valueBTC)
	if err != nil {
		log.Fatal(err)
	}

	bitcoinPrice, _ = strconv.ParseFloat(valueBTC.Bid, 64)
}

var bitcoinPrice float64

func init() {
	c.GetConf()
	bitcoinToDollar()
}

func main() {
	fmt.Println("%v", utils.GetPosition())

	fmt.Println("%v", BITMEXREFLINK)
}
