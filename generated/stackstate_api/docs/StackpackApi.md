# \StackpackApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisionDetails**](StackpackApi.md#ProvisionDetails) | **Post** /stackpack/{stackName}/provision | Provision API
[**StackpackList**](StackpackApi.md#StackpackList) | **Get** /stackpack | StackPack API
[**StackpackUpload**](StackpackApi.md#StackpackUpload) | **Post** /stackpack | StackPack API
[**UpgradeStackPack**](StackpackApi.md#UpgradeStackPack) | **Post** /stackpack/{stackName}/upgrade | Upgrade API



## ProvisionDetails

> ProvisionResponse ProvisionDetails(ctx, stackName).RequestBody(requestBody).Execute()

Provision API



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
    stackName := "stackName_example" // string | 
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.ProvisionDetails(context.Background(), stackName).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.ProvisionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionDetails`: ProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.ProvisionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]string** |  | 

### Return type

[**ProvisionResponse**](ProvisionResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackpackList

> []Sstackpack StackpackList(ctx).Execute()

StackPack API



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
    resp, r, err := apiClient.StackpackApi.StackpackList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackpackList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackpackList`: []Sstackpack
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackpackList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStackpackListRequest struct via the builder pattern


### Return type

[**[]Sstackpack**](Sstackpack.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackpackUpload

> StackPack StackpackUpload(ctx).StackPack(stackPack).Execute()

StackPack API



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
    stackPack := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.StackpackUpload(context.Background()).StackPack(stackPack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackpackUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackpackUpload`: StackPack
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackpackUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStackpackUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackPack** | ***os.File** |  | 

### Return type

[**StackPack**](StackPack.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeStackPack

> string UpgradeStackPack(ctx, stackName).Unlocked(unlocked).Execute()

Upgrade API



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
    stackName := "stackName_example" // string | 
    unlocked := "unlocked_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.UpgradeStackPack(context.Background(), stackName).Unlocked(unlocked).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.UpgradeStackPack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeStackPack`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.UpgradeStackPack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeStackPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unlocked** | **string** |  | 

### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

