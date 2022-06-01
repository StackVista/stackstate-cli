# \NodeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NodeListTypes**](NodeApi.md#NodeListTypes) | **Get** /node | Node API
[**TypeList**](NodeApi.md#TypeList) | **Get** /node/{nodeType} | Node type API



## NodeListTypes

> NodeTypes NodeListTypes(ctx).Execute()

Node API



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
    resp, r, err := apiClient.NodeApi.NodeListTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.NodeListTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NodeListTypes`: NodeTypes
    fmt.Fprintf(os.Stdout, "Response from `NodeApi.NodeListTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNodeListTypesRequest struct via the builder pattern


### Return type

[**NodeTypes**](NodeTypes.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypeList

> []Node TypeList(ctx, nodeType).Namespace(namespace).OwnedBy(ownedBy).Execute()

Node type API



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
    nodeType := "nodeType_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    ownedBy := "ownedBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodeApi.TypeList(context.Background(), nodeType).Namespace(namespace).OwnedBy(ownedBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.TypeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypeList`: []Node
    fmt.Fprintf(os.Stdout, "Response from `NodeApi.TypeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **ownedBy** | **string** |  | 

### Return type

[**[]Node**](Node.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

