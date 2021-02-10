package coincheck

import "fmt"

// Order 取引所の注文情報
type Order struct {
	client *CoinCheck
}

type Pair string

const (
	BTC = Pair("btc_jpy")
)

type OrderType string

const (
	Buy  = OrderType("buy")
	Sell = OrderType("sell")
)

// Rate 取引所の注文を元にレートを取得する
func (a Order) Rate(orderType OrderType, pair Pair, amount float64) string {
	param := fmt.Sprintf("?order_type=%s&amount=%f&pair=%s", orderType, amount, pair)
	return a.client.Request("GET", "api/exchange/orders/rate"+param, "")
}
