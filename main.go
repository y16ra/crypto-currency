package main

import (
	"fmt"

	"github.com/y16ra/crypto-currency/gmo"
)

func main() {
	// GMO Coin client
	gmoClient := new(gmo.GmoCoin).NewClient()

	// ticker := gmoClient.Ticker.All()
	// fmt.Println(ticker)
	// assetsResp := gmoClient.Account.Assets()
	// fmt.Printf("%+v \n", assetsResp)

	// merginResp := gmoClient.Account.Margin()
	// fmt.Printf("%+v \n", merginResp)

	// orderParam := "orderId=1145804966"
	// orderResp := gmoClient.Order.OrderList(orderParam)
	// fmt.Printf("%+v \n", orderResp)

	aOrders := gmoClient.Order.ActiveOrderList("BTC")
	fmt.Printf("%+v \n", aOrders)
}
