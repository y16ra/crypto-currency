# \PublicApi

All URIs are relative to *https://api.coin.z.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicV1OrderbooksGet**](PublicApi.md#PublicV1OrderbooksGet) | **Get** /public/v1/orderbooks | Get orderbooks
[**PublicV1StatusGet**](PublicApi.md#PublicV1StatusGet) | **Get** /public/v1/status | Get an exchange status
[**PublicV1TickerGet**](PublicApi.md#PublicV1TickerGet) | **Get** /public/v1/ticker | Get latest rate
[**PublicV1TradesGet**](PublicApi.md#PublicV1TradesGet) | **Get** /public/v1/trades | Get trade histories



## PublicV1OrderbooksGet

> Orderbooks PublicV1OrderbooksGet(ctx).Symbol(symbol).Execute()

Get orderbooks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := openapiclient.Symbols("BTC") // Symbols | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.PublicV1OrderbooksGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.PublicV1OrderbooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicV1OrderbooksGet`: Orderbooks
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.PublicV1OrderbooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicV1OrderbooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**Symbols**](Symbols.md) |  | 

### Return type

[**Orderbooks**](Orderbooks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicV1StatusGet

> Status PublicV1StatusGet(ctx).Execute()

Get an exchange status

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.PublicV1StatusGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.PublicV1StatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicV1StatusGet`: Status
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.PublicV1StatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPublicV1StatusGetRequest struct via the builder pattern


### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicV1TickerGet

> Tickers PublicV1TickerGet(ctx).Symbol(symbol).Execute()

Get latest rate

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := openapiclient.Symbols("BTC") // Symbols |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.PublicV1TickerGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.PublicV1TickerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicV1TickerGet`: Tickers
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.PublicV1TickerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicV1TickerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**Symbols**](Symbols.md) |  | 

### Return type

[**Tickers**](Tickers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicV1TradesGet

> Trades PublicV1TradesGet(ctx).Symbol(symbol).Page(page).Count(count).Execute()

Get trade histories

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    symbol := openapiclient.Symbols("BTC") // Symbols | 
    page := int32(56) // int32 | page number. (default 1) (optional)
    count := int32(56) // int32 | Max count per request. (default 100) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.PublicV1TradesGet(context.Background()).Symbol(symbol).Page(page).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.PublicV1TradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicV1TradesGet`: Trades
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.PublicV1TradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicV1TradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**Symbols**](Symbols.md) |  | 
 **page** | **int32** | page number. (default 1) | 
 **count** | **int32** | Max count per request. (default 100) | 

### Return type

[**Trades**](Trades.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

