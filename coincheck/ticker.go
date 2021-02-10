package coincheck

type Ticker struct {
	client *CoinCheck
}

// All 各種最新情報を簡易に取得することができます(BTC only)。
func (a Ticker) All() string {
	return a.client.Request("GET", "api/ticker", "")
}
