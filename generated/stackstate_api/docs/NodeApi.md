# \NodeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Clone**](NodeApi.md#Clone) | **Post** /node/{nodeType}/{nodeIdOrUrn}/clone | Clone a node with a new name
[**Delete**](NodeApi.md#Delete) | **Delete** /node/{nodeType}/{nodeId} | Node deletion API
[**Lock**](NodeApi.md#Lock) | **Get** /node/{nodeType}/{nodeId}/islocked | Retrieve if a node is locked
[**NodeListTypes**](NodeApi.md#NodeListTypes) | **Get** /node | Node API
[**TypeList**](NodeApi.md#TypeList) | **Get** /node/{nodeType} | Node type API
[**Unlock**](NodeApi.md#Unlock) | **Post** /node/{nodeType}/{nodeId}/unlock | Node unlock API



## Clone

> Node Clone(ctx, nodeType, nodeIdOrUrn).NodeName(nodeName).Execute()

Clone a node with a new name



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
    nodeIdOrUrn := "nodeIdOrUrn_example" // string | 
    nodeName := *openapiclient.NewNodeName("Name_example") // NodeName | New name for the node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodeApi.Clone(context.Background(), nodeType, nodeIdOrUrn).NodeName(nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.Clone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Clone`: Node
    fmt.Fprintf(os.Stdout, "Response from `NodeApi.Clone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeType** | **string** |  | 
**nodeIdOrUrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeName** | [**NodeName**](NodeName.md) | New name for the node | 

### Return type

[**Node**](Node.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, nodeType, nodeId).TimeoutSeconds(timeoutSeconds).Execute()

Node deletion API



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
    nodeId := int64(789) // int64 | 
    timeoutSeconds := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodeApi.Delete(context.Background(), nodeType, nodeId).TimeoutSeconds(timeoutSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeType** | **string** |  | 
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeoutSeconds** | **int64** |  | 

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


## Lock

> LockedResponse Lock(ctx, nodeType, nodeId).Execute()

Retrieve if a node is locked



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
    nodeId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodeApi.Lock(context.Background(), nodeType, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.Lock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Lock`: LockedResponse
    fmt.Fprintf(os.Stdout, "Response from `NodeApi.Lock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeType** | **string** |  | 
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LockedResponse**](LockedResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unlock

> Node Unlock(ctx, nodeType, nodeId).Execute()

Node unlock API



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
    nodeId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodeApi.Unlock(context.Background(), nodeType, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeApi.Unlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unlock`: Node
    fmt.Fprintf(os.Stdout, "Response from `NodeApi.Unlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeType** | **string** |  | 
**nodeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Node**](Node.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

