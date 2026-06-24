# \PerspectivesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPerspectives**](PerspectivesApi.md#GetPerspectives) | **Get** /perspectives/{presentationOrViewUrn} | Get the perspectives for a view



## GetPerspectives

> Perspectives GetPerspectives(ctx, presentationOrViewUrn).Execute()

Get the perspectives for a view



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
    presentationOrViewUrn := "presentationOrViewUrn_example" // string | A Component Presentation Identifier, legacy View (QueryView, ViewType) URNs are supported for backward compatibility

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerspectivesApi.GetPerspectives(context.Background(), presentationOrViewUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerspectivesApi.GetPerspectives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPerspectives`: Perspectives
    fmt.Fprintf(os.Stdout, "Response from `PerspectivesApi.GetPerspectives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presentationOrViewUrn** | **string** | A Component Presentation Identifier, legacy View (QueryView, ViewType) URNs are supported for backward compatibility | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPerspectivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Perspectives**](Perspectives.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

