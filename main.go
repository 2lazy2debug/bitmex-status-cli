// Copyright (c) 2019 Manuel Cabras
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"bitmex-status-cli/bitmex"
	"bitmex-status-cli/math"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

const BITMEXREFLINK = "https://www.bitmex.com/register/1TTdTj"

func main() {
	fmt.Printf("\n")
	fmt.Printf("#########################################################################\n")
	fmt.Printf("# BitMEX Status (CLI) - Author : Manuel Cabras - cabrasm@protonmail.com #\n")
	fmt.Printf("#          Join BitMEX: https://www.bitmex.com/register/1TTdTj          #\n")
	fmt.Printf("#########################################################################\n\n")

	readPositions(bitmex.GetPositions())
	readOrders(bitmex.GetOrders())

	fmt.Println("Join BitMEX and spare 10% fees: ", BITMEXREFLINK)
}

func readPositions(positions bitmex.Positions) {
	fmt.Printf("Open positions : \n--------------------------------------------- \n\n")
	for _, position := range positions {
		if position.CurrentQty != 0 {

			fmt.Printf("Symbol : " + position.Symbol + "\n")
			fmt.Printf("Qty : " + strconv.Itoa(position.CurrentQty) + "\n")
			fmt.Printf("Leverage : " + strconv.Itoa(position.Leverage) + "\n")
			fmt.Printf("Mark Value : " + fmt.Sprintf("%.8f", math.SatoshisToXBT(-1*position.MarkValue)) + "\n")
			fmt.Printf("Last Price : " + fmt.Sprintf("%.2f", position.LastPrice) + "\n")
			fmt.Printf("Average Cost Price : " + strconv.Itoa(position.AvgCostPrice) + "\n")
			fmt.Printf("Liquidation Price : " + strconv.Itoa(position.LiquidationPrice) + "\n")
			fmt.Printf("Unrealised PNL : " + fmt.Sprintf("%.8f", math.SatoshisToXBT(position.UnrealisedPnl)) + "\n")
			fmt.Printf("\n\n")
		}
	}
}

func readOrders(orders bitmex.Orders) {
	fmt.Printf("Open orders : \n--------------------------------------------- \n\n")
	for _, order := range orders {

		if order.OrdType == "StopLimit" {
			color.Red("STOP LOSS POSITION")
		} else {
			color.Green("TAKE PROFIT POSITION")
		}

		fmt.Printf("Order id : " + order.OrderID + "\n")
		fmt.Printf("Symbol : " + order.Symbol + "\n")
		fmt.Printf("Current Quantity : " + strconv.Itoa(order.OrderQty) + "\n")
		fmt.Printf("Price : " + fmt.Sprintf("%.2f", order.Price) + "\n")
		fmt.Printf("Type : " + order.OrdType + "\n")
		fmt.Printf("Status : " + order.OrdStatus + "\n")
		fmt.Printf("Triggered : " + order.Triggered + "\n")
		fmt.Printf("\n\n")
	}
}
