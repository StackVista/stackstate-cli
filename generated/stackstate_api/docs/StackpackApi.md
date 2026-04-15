# \StackpackApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmManualSteps**](StackpackApi.md#ConfirmManualSteps) | **Post** /stackpack/{stackPackName}/confirm-manual-steps/{stackPackInstanceId} | Confirm manual steps
[**ProvisionDetails**](StackpackApi.md#ProvisionDetails) | **Post** /stackpack/{stackPackName}/provision | Provision API
[**ProvisionUninstall**](StackpackApi.md#ProvisionUninstall) | **Post** /stackpack/{stackPackName}/deprovision/{stackPackInstanceId} | Provision API
[**StackPackDeleteVersion**](StackpackApi.md#StackPackDeleteVersion) | **Delete** /stackpack/{stackPackName}/versions/{version} | Delete a StackPack version
[**StackPackDeleteVersions**](StackpackApi.md#StackPackDeleteVersions) | **Delete** /stackpack/{stackPackName}/versions | Delete StackPack versions
[**StackPackList**](StackpackApi.md#StackPackList) | **Get** /stackpack | StackPack API
[**StackPackListVersions**](StackpackApi.md#StackPackListVersions) | **Get** /stackpack/{stackPackName}/versions | List StackPack versions
[**StackPackUpload**](StackpackApi.md#StackPackUpload) | **Post** /stackpack | StackPack API
[**StackPackValidate**](StackpackApi.md#StackPackValidate) | **Post** /stackpack/validate | Validate API
[**UpgradeStackPack**](StackpackApi.md#UpgradeStackPack) | **Post** /stackpack/{stackPackName}/upgrade | Upgrade API



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
    resp, r, err := apiClient.StackpackApi.ConfirmManualSteps(context.Background(), stackPackName, stackPackInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.ConfirmManualSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmManualSteps`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.ConfirmManualSteps`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.StackpackApi.ProvisionDetails(context.Background(), stackPackName).Unlocked(unlocked).RequestBody(requestBody).Execute()
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.StackpackApi.ProvisionUninstall(context.Background(), stackPackName, stackPackInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.ProvisionUninstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionUninstall`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.ProvisionUninstall`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackPackDeleteVersion

> StackPackDeleteVersion(ctx, stackPackName, version).Execute()

Delete a StackPack version



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
    version := "version_example" // string | Version string (e.g. '1.2.3')

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.StackPackDeleteVersion(context.Background(), stackPackName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackDeleteVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 
**version** | **string** | Version string (e.g. &#39;1.2.3&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStackPackDeleteVersionRequest struct via the builder pattern


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


## StackPackDeleteVersions

> DeleteVersionsResult StackPackDeleteVersions(ctx, stackPackName).Before(before).From(from).To(to).Execute()

Delete StackPack versions



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
    before := "before_example" // string | Delete all versions strictly less than this version (optional)
    from := "from_example" // string | Inclusive start of version range to delete (optional)
    to := "to_example" // string | Inclusive end of version range to delete (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.StackPackDeleteVersions(context.Background(), stackPackName).Before(before).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackDeleteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackDeleteVersions`: DeleteVersionsResult
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackPackDeleteVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStackPackDeleteVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **string** | Delete all versions strictly less than this version | 
 **from** | **string** | Inclusive start of version range to delete | 
 **to** | **string** | Inclusive end of version range to delete | 

### Return type

[**DeleteVersionsResult**](DeleteVersionsResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.StackpackApi.StackPackList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackList`: []FullStackPack
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackPackList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStackPackListRequest struct via the builder pattern


### Return type

[**[]FullStackPack**](FullStackPack.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackPackListVersions

> []StackPackVersionInfo StackPackListVersions(ctx, stackPackName).Execute()

List StackPack versions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StackpackApi.StackPackListVersions(context.Background(), stackPackName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackListVersions`: []StackPackVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackPackListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackPackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStackPackListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StackPackVersionInfo**](StackPackVersionInfo.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.StackpackApi.StackPackUpload(context.Background()).StackPack(stackPack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackUpload`: StackPack
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackPackUpload`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackPackValidate

> string StackPackValidate(ctx).StackPack(stackPack).Execute()

Validate API



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
    resp, r, err := apiClient.StackpackApi.StackPackValidate(context.Background()).StackPack(stackPack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackpackApi.StackPackValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StackPackValidate`: string
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackPackValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStackPackValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackPack** | ***os.File** |  | 

### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.StackpackApi.UpgradeStackPack(context.Background(), stackPackName).Unlocked(unlocked).Execute()
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
**stackPackName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeStackPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unlocked** | **string** |  | 

### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

