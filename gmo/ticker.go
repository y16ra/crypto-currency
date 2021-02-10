package gmo

type Ticker struct {
	client *GmoCoin
}

func (t Ticker) All() []byte {
	return t.client.Request("GET", "/v1/ticker", "", true)
}
