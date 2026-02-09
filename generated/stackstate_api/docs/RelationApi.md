# \RelationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFullRelation**](RelationApi.md#GetFullRelation) | **Get** /relations/{relationId} | Get full relation



## GetFullRelation

> FullRelation GetFullRelation(ctx, relationId).TopologyTime(topologyTime).Execute()

Get full relation



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
    relationId := int64(789) // int64 | The id of a relation
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.GetFullRelation(context.Background(), relationId).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetFullRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFullRelation`: FullRelation
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetFullRelation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The id of a relation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**FullRelation**](FullRelation.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

