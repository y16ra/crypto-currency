package gmo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/y16ra/crypto-currency/config"
)

// GmoCoin client
type GmoCoin struct {
	baseURL   string
	accessKey string
	secretKey string
	// APIs
	Ticker  Ticker
	Account Account
	Order   Order
}

// NewClient create client for GMO Coin APIs.
func (c GmoCoin) NewClient() GmoCoin {
	config := config.GmoConfig
	c.baseURL = config.BaseURL
	c.accessKey = config.ApiKey
	c.secretKey = config.ApiSecret
	// public api
	c.Ticker = Ticker{&c}
	// private api
	c.Account = Account{&c}
	c.Order = Order{&c}

	return c
}

// Request connect to GMO Coin APIs.
func (c GmoCoin) Request(method string, path string, param string, public bool) []byte {
	url := ""
	if param != "" && method == "GET" {
		url = createURL(c.baseURL, path+"?"+param, public)
		param = ""
	} else {
		url = createURL(c.baseURL, path, public)
	}
	req := &http.Request{}
	if method == "POST" {
		payload := strings.NewReader(param)
		req, _ = http.NewRequest(method, url, payload)
	} else {
		log.Printf("url=%s", url)
		req, _ = http.NewRequest(method, url, nil)
	}
	timestamp := createTimestamp()
	message := timestamp + method + path + param
	sign := computeHmac256(message, c.secretKey)
	req.Header.Set("API-KEY", c.accessKey)
	req.Header.Set("API-TIMESTAMP", timestamp)
	req.Header.Set("API-SIGN", sign)

	client := &http.Client{}
	res, err := client.Do(req)
	fmt.Printf("status code= %d \n", res.StatusCode)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func createTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

//ComputeHmac256 create signature
func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func createURL(baseURL, path string, public bool) string {
	if public {
		baseURL += "/public" + path
	} else {
		baseURL += "/private" + path
	}
	return baseURL
}
