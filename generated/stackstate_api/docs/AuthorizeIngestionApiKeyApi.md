# \AuthorizeIngestionApiKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeIngestionApiKey**](AuthorizeIngestionApiKeyApi.md#AuthorizeIngestionApiKey) | **Post** /security/ingestion/authorize | Check authorization for an Ingestion Api Key



## AuthorizeIngestionApiKey

> AuthorizeIngestionApiKey(ctx).AuthorizeIngestionApiKeyRequest(authorizeIngestionApiKeyRequest).WithReceiverKey(withReceiverKey).Execute()

Check authorization for an Ingestion Api Key



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
    authorizeIngestionApiKeyRequest := *openapiclient.NewAuthorizeIngestionApiKeyRequest("ApiKey_example") // AuthorizeIngestionApiKeyRequest | 
    withReceiverKey := true // bool | By default, the endpoint uses only Ingestion API Keys, true value - to verify also Receiver API Key (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeIngestionApiKeyApi.AuthorizeIngestionApiKey(context.Background()).AuthorizeIngestionApiKeyRequest(authorizeIngestionApiKeyRequest).WithReceiverKey(withReceiverKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeIngestionApiKeyApi.AuthorizeIngestionApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizeIngestionApiKeyRequest** | [**AuthorizeIngestionApiKeyRequest**](AuthorizeIngestionApiKeyRequest.md) |  | 
 **withReceiverKey** | **bool** | By default, the endpoint uses only Ingestion API Keys, true value - to verify also Receiver API Key | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

