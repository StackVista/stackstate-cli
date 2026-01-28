# \ComponentPresentationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponentPresentationByIdentifier**](ComponentPresentationApi.md#DeleteComponentPresentationByIdentifier) | **Delete** /component-presentations/{identifier} | Delete a component presentation by Identifier
[**GetComponentPresentationByIdentifier**](ComponentPresentationApi.md#GetComponentPresentationByIdentifier) | **Get** /component-presentations/{identifier} | Get a component presentation by Identifier
[**GetComponentPresentations**](ComponentPresentationApi.md#GetComponentPresentations) | **Get** /component-presentations | List all component presentations
[**UpsertComponentPresentations**](ComponentPresentationApi.md#UpsertComponentPresentations) | **Put** /component-presentations | Upserts (creates/updates) a component presentation



## DeleteComponentPresentationByIdentifier

> DeleteComponentPresentationByIdentifier(ctx, identifier).Execute()

Delete a component presentation by Identifier

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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentPresentationApi.DeleteComponentPresentationByIdentifier(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentPresentationApi.DeleteComponentPresentationByIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentPresentationByIdentifierRequest struct via the builder pattern


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


## GetComponentPresentationByIdentifier

> ComponentPresentation GetComponentPresentationByIdentifier(ctx, identifier).Execute()

Get a component presentation by Identifier

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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentPresentationApi.GetComponentPresentationByIdentifier(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentPresentationApi.GetComponentPresentationByIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentPresentationByIdentifier`: ComponentPresentation
    fmt.Fprintf(os.Stdout, "Response from `ComponentPresentationApi.GetComponentPresentationByIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentPresentationByIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentPresentation**](ComponentPresentation.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentPresentations

> []ComponentPresentation GetComponentPresentations(ctx).Execute()

List all component presentations

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
    resp, r, err := apiClient.ComponentPresentationApi.GetComponentPresentations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentPresentationApi.GetComponentPresentations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentPresentations`: []ComponentPresentation
    fmt.Fprintf(os.Stdout, "Response from `ComponentPresentationApi.GetComponentPresentations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentPresentationsRequest struct via the builder pattern


### Return type

[**[]ComponentPresentation**](ComponentPresentation.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertComponentPresentations

> ComponentPresentation UpsertComponentPresentations(ctx).ComponentPresentation(componentPresentation).Execute()

Upserts (creates/updates) a component presentation

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
    componentPresentation := *openapiclient.NewComponentPresentation("Identifier_example", "Name_example", *openapiclient.NewPresentationDefinition()) // ComponentPresentation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentPresentationApi.UpsertComponentPresentations(context.Background()).ComponentPresentation(componentPresentation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentPresentationApi.UpsertComponentPresentations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertComponentPresentations`: ComponentPresentation
    fmt.Fprintf(os.Stdout, "Response from `ComponentPresentationApi.UpsertComponentPresentations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertComponentPresentationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentPresentation** | [**ComponentPresentation**](ComponentPresentation.md) |  | 

### Return type

[**ComponentPresentation**](ComponentPresentation.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

