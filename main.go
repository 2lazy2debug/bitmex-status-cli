// Copyright (c) 2019 Manuel Cabras
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"bitmex-status-cli/bitmex"
	"bitmex-status-cli/math"
	"fmt"
	"strconv"
)

const BITMEXREFLINK = "https://www.bitmex.com/register/1TTdTj"

func main() {
	position := bitmex.GetPosition()
	fmt.Printf("\n")
	fmt.Printf("#########################################################################\n")
	fmt.Printf("# BitMEX Status (CLI) - Author : Manuel Cabras - cabrasm@protonmail.com #\n")
	fmt.Printf("#          Join BitMEX: https://www.bitmex.com/register/1TTdTj          #\n")
	fmt.Printf("#########################################################################\n\n")

	fmt.Printf("Open position : \n--------------------------------------------- \n\n")
	fmt.Printf("Symbol : " + position[0].Symbol + "\n")
	fmt.Printf("Current Quantity : " + strconv.Itoa(position[0].CurrentQty) + "\n")
	fmt.Printf("Leverage : " + strconv.Itoa(position[0].Leverage) + "\n")
	fmt.Printf("Mark Value : " + fmt.Sprintf("%.8f", math.SatoshisToXBT(-1*position[0].MarkValue)) + "\n")
	fmt.Printf("Last Price : " + fmt.Sprintf("%.2f", position[0].LastPrice) + "\n")
	fmt.Printf("Average Cost Price : " + strconv.Itoa(position[0].AvgCostPrice) + "\n")
	fmt.Printf("Liquidation Price : " + strconv.Itoa(position[0].LiquidationPrice) + "\n")
	fmt.Printf("Unrealised PNL : " + fmt.Sprintf("%.8f", math.SatoshisToXBT(position[0].UnrealisedPnl)) + "\n")
	fmt.Printf("\n\n")

	fmt.Println("Join BitMEX and spare 10% fees: ", BITMEXREFLINK)
}
