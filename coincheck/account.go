package coincheck

// Account
type Account struct {
	client *CoinCheck
}

// Balance get a balance.
func (a Account) Balance() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}

// LeverageBalance get a leverage balance.
func (a Account) LeverageBalance() string {
	return a.client.Request("GET", "api/accounts/leverage_balance", "")
}

// Info get an account information.
func (a Account) Info() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}
