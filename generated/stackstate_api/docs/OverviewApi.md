# \OverviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOverview**](OverviewApi.md#GetOverview) | **Post** /overview/{presentationOrViewUrn} | Get overview data for components



## GetOverview

> OverviewPageResponse GetOverview(ctx, presentationOrViewUrn).OverviewRequest(overviewRequest).Execute()

Get overview data for components

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
    presentationOrViewUrn := "presentationOrViewUrn_example" // string | A Component Presentation Identifier, legacy View (QueryView, ViewType) URNs are be supported for backward compatibility
    overviewRequest := *openapiclient.NewOverviewRequest() // OverviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOverview(context.Background(), presentationOrViewUrn).OverviewRequest(overviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOverview`: OverviewPageResponse
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presentationOrViewUrn** | **string** | A Component Presentation Identifier, legacy View (QueryView, ViewType) URNs are be supported for backward compatibility | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overviewRequest** | [**OverviewRequest**](OverviewRequest.md) |  | 

### Return type

[**OverviewPageResponse**](OverviewPageResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

