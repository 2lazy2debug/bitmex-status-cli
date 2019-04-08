// Copyright (c) 2019 Manuel Cabras
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"bitmex-status-cli/utils"
	"fmt"
)

const BITMEXREFLINK = "https://www.bitmex.com/register/1TTdTj"

func main() {
	fmt.Println("Your actual BitMex position : \n", utils.GetPosition(), " \n")

	fmt.Println("Share this link with a friend if you want him to join BitMEX : \n", BITMEXREFLINK)
}
