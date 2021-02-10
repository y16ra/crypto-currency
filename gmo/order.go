package gmo

import (
	"encoding/json"
	"time"
)

type Order struct {
	client *GmoCoin
}

type OrderResponce struct {
	Status int `json:"status"`
	Data   struct {
		List []struct {
			OrderID       int       `json:"orderId"`
			RootOrderID   int       `json:"rootOrderId"`
			Symbol        string    `json:"symbol"`
			Side          string    `json:"side"`
			OrderType     string    `json:"orderType"`
			ExecutionType string    `json:"executionType"`
			SettleType    string    `json:"settleType"`
			Size          string    `json:"size"`
			ExecutedSize  string    `json:"executedSize"`
			Price         string    `json:"price"`
			LosscutPrice  string    `json:"losscutPrice"`
			Status        string    `json:"status"`
			TimeInForce   string    `json:"timeInForce"`
			Timestamp     time.Time `json:"timestamp"`
			CancelType    string    `json:"cancelType,omitempty"`
		} `json:"list"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

// OrderList retrive order list
func (o Order) OrderList(orderID string) OrderResponce {
	var responce OrderResponce
	json.Unmarshal(o.client.Request("GET", "/v1/orders", orderID, false), &responce)
	return responce
}

type ActiveOrderResponse struct {
	Status int `json:"status"`
	Data   struct {
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			Count       int `json:"count"`
		} `json:"pagination"`
		List []struct {
			RootOrderID   int       `json:"rootOrderId"`
			OrderID       int       `json:"orderId"`
			Symbol        string    `json:"symbol"`
			Side          string    `json:"side"`
			OrderType     string    `json:"orderType"`
			ExecutionType string    `json:"executionType"`
			SettleType    string    `json:"settleType"`
			Size          string    `json:"size"`
			ExecutedSize  string    `json:"executedSize"`
			Price         string    `json:"price"`
			LosscutPrice  string    `json:"losscutPrice"`
			Status        string    `json:"status"`
			TimeInForce   string    `json:"timeInForce"`
			Timestamp     time.Time `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

// ActiveOrderList retrive active orders.
func (o Order) ActiveOrderList(symbol string) ActiveOrderResponse {
	var responce ActiveOrderResponse
	json.Unmarshal(o.client.Request("GET", "/v1/activeOrders", "symbol="+symbol, false), &responce)
	return responce
}
