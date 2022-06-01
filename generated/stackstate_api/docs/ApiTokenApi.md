# \ApiTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentUserApiTokens**](ApiTokenApi.md#GetCurrentUserApiTokens) | **Get** /user/profile/tokens | Get current user&#39;s API tokens



## GetCurrentUserApiTokens

> []ApiToken GetCurrentUserApiTokens(ctx).Execute()

Get current user's API tokens



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiTokenApi.GetCurrentUserApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenApi.GetCurrentUserApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserApiTokens`: []ApiToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokenApi.GetCurrentUserApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserApiTokensRequest struct via the builder pattern


### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

