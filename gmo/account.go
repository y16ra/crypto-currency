package gmo

import (
	"encoding/json"
	"time"
)

// Account has margin, assets
type Account struct {
	client *GmoCoin
}

// Margin retrieve account info about margin
func (a Account) Margin() MarginResponce {
	var responce MarginResponce
	json.Unmarshal(a.client.Request("GET", "/v1/account/margin", "", false), &responce)
	return responce
}

type MarginResponce struct {
	Status int `json:"status"`
	Data   struct {
		ActualProfitLoss string `json:"actualProfitLoss"`
		AvailableAmount  string `json:"availableAmount"`
		Margin           string `json:"margin"`
		ProfitLoss       string `json:"profitLoss"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

// Assets 資産残高を取得します
func (a Account) Assets() AssetsResponse {
	var assetsResp AssetsResponse
	json.Unmarshal(a.client.Request("GET", "/v1/account/assets", "", false), &assetsResp)
	return assetsResp
}

// AssetsResponse
type AssetsResponse struct {
	Status int `json:"status"`
	Data   []struct {
		Amount         string `json:"amount"`
		Available      string `json:"available"`
		ConversionRate string `json:"conversionRate"`
		Symbol         string `json:"symbol"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}
