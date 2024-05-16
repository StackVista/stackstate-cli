# \StackpackAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmManualSteps**](StackpackAPI.md#ConfirmManualSteps) | **Post** /stackpack/{stackPackName}/confirm-manual-steps/{stackPackInstanceId} | Confirm manual steps
[**ProvisionDetails**](StackpackAPI.md#ProvisionDetails) | **Post** /stackpack/{stackPackName}/provision | Provision API
[**ProvisionUninstall**](StackpackAPI.md#ProvisionUninstall) | **Post** /stackpack/{stackPackName}/deprovision/{stackPackInstanceId} | Provision API
[**StackPackList**](StackpackAPI.md#StackPackList) | **Get** /stackpack | StackPack API
[**StackPackUpload**](StackpackAPI.md#StackPackUpload) | **Post** /stackpack | StackPack API
[**UpgradeStackPack**](StackpackAPI.md#UpgradeStackPack) | **Post** /stackpack/{stackPackName}/upgrade | Upgrade API



## ConfirmManualSteps

> string ConfirmManualSteps(ctx, stackPackName, stackPackInstanceId).Execute()

Confirm manual steps



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
    stackPackName := "stackPackName_example" // string | 
    stackPackInstanceId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackAPI.ConfirmManualSteps(context.Background(), stackPackName, stackPackInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.ConfirmManualSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmManualSteps`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.ConfirmManualSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 
**stackPackInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmManualStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionDetails

> ProvisionResponse ProvisionDetails(ctx, stackPackName).Unlocked(unlocked).RequestBody(requestBody).Execute()

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
    stackPackName := "stackPackName_example" // string | 
    unlocked := "unlocked_example" // string | 
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackAPI.ProvisionDetails(context.Background(), stackPackName).Unlocked(unlocked).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.ProvisionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionDetails`: ProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.ProvisionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unlocked** | **string** |  | 
 **requestBody** | **map[string]string** |  | 

### Return type

[**ProvisionResponse**](ProvisionResponse.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionUninstall

> string ProvisionUninstall(ctx, stackPackName, stackPackInstanceId).Execute()

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
    stackPackName := "stackPackName_example" // string | 
    stackPackInstanceId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackAPI.ProvisionUninstall(context.Background(), stackPackName, stackPackInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.ProvisionUninstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionUninstall`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.ProvisionUninstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 
**stackPackInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionUninstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackPackList

> []FullStackPack StackPackList(ctx).Execute()

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
    resp, r, err := apiClient.StackpackAPI.StackPackList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.StackPackList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackList`: []FullStackPack
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.StackPackList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStackPackListRequest struct via the builder pattern


### Return type

[**[]FullStackPack**](FullStackPack.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackPackUpload

> StackPack StackPackUpload(ctx).StackPack(stackPack).Execute()

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
    resp, r, err := apiClient.StackpackAPI.StackPackUpload(context.Background()).StackPack(stackPack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.StackPackUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackUpload`: StackPack
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.StackPackUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStackPackUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackPack** | ***os.File** |  | 

### Return type

[**StackPack**](StackPack.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeStackPack

> string UpgradeStackPack(ctx, stackPackName).Unlocked(unlocked).Execute()

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
    stackPackName := "stackPackName_example" // string | 
    unlocked := "unlocked_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackAPI.UpgradeStackPack(context.Background(), stackPackName).Unlocked(unlocked).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackAPI.UpgradeStackPack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeStackPack`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackAPI.UpgradeStackPack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeStackPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unlocked** | **string** |  | 

### Return type

**string**

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

