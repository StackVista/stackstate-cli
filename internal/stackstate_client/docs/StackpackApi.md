# \StackpackApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StackpackList**](StackpackApi.md#StackpackList) | **Get** /stackpack | StackPack API
[**StackpackUpload**](StackpackApi.md#StackpackUpload) | **Post** /stackpack | StackPack API



## StackpackList

> []Stackpack StackpackList(ctx).Execute()

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
    // response from `StackpackList`: []Stackpack
    fmt.Fprintf(os.Stdout, "Response from `StackpackApi.StackpackList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStackpackListRequest struct via the builder pattern


### Return type

[**[]Stackpack**](Stackpack.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StackpackUpload

> Stackpack StackpackUpload(ctx).StackPack(stackPack).Execute()

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
    // response from `StackpackUpload`: Stackpack
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

[**Stackpack**](Stackpack.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

