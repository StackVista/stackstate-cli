# \IngestionApiKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIngestionApiKey**](IngestionApiKeyApi.md#DeleteIngestionApiKey) | **Delete** /security/ingestion/api_keys/{ingestionApiKeyId} | Delete Ingestion Api Key
[**GenerateIngestionApiKey**](IngestionApiKeyApi.md#GenerateIngestionApiKey) | **Post** /security/ingestion/api_keys | Generate a new Ingestion Api Key
[**GetIngestionApiKeys**](IngestionApiKeyApi.md#GetIngestionApiKeys) | **Get** /security/ingestion/api_keys | List Ingestion Api Keys



## DeleteIngestionApiKey

> DeleteIngestionApiKey(ctx, ingestionApiKeyId).Execute()

Delete Ingestion Api Key



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
    ingestionApiKeyId := int64(789) // int64 | The identifier of a key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionApiKeyApi.DeleteIngestionApiKey(context.Background(), ingestionApiKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeyApi.DeleteIngestionApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ingestionApiKeyId** | **int64** | The identifier of a key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateIngestionApiKey

> GeneratedIngestionApiKeyResponse GenerateIngestionApiKey(ctx).GenerateIngestionApiKeyRequest(generateIngestionApiKeyRequest).Execute()

Generate a new Ingestion Api Key



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
    generateIngestionApiKeyRequest := *openapiclient.NewGenerateIngestionApiKeyRequest("Name_example") // GenerateIngestionApiKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IngestionApiKeyApi.GenerateIngestionApiKey(context.Background()).GenerateIngestionApiKeyRequest(generateIngestionApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeyApi.GenerateIngestionApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateIngestionApiKey`: GeneratedIngestionApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeyApi.GenerateIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateIngestionApiKeyRequest** | [**GenerateIngestionApiKeyRequest**](GenerateIngestionApiKeyRequest.md) |  | 

### Return type

[**GeneratedIngestionApiKeyResponse**](GeneratedIngestionApiKeyResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestionApiKeys

> []IngestionApiKey GetIngestionApiKeys(ctx).Execute()

List Ingestion Api Keys



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
    resp, r, err := apiClient.IngestionApiKeyApi.GetIngestionApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeyApi.GetIngestionApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestionApiKeys`: []IngestionApiKey
    fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeyApi.GetIngestionApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionApiKeysRequest struct via the builder pattern


### Return type

[**[]IngestionApiKey**](IngestionApiKey.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

