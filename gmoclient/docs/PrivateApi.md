# \PrivateApi

All URIs are relative to *https://api.coin.z.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateV1AccountMarginGet**](PrivateApi.md#PrivateV1AccountMarginGet) | **Get** /private/v1/account/margin | Get account margin



## PrivateV1AccountMarginGet

> AccountMargin PrivateV1AccountMarginGet(ctx).Execute()

Get account margin

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
    resp, r, err := api_client.PrivateApi.PrivateV1AccountMarginGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1AccountMarginGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1AccountMarginGet`: AccountMargin
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1AccountMarginGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1AccountMarginGetRequest struct via the builder pattern


### Return type

[**AccountMargin**](AccountMargin.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

