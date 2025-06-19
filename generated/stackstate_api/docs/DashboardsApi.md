# \DashboardsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDashboard**](DashboardsApi.md#CloneDashboard) | **Post** /dashboards/{dashboardIdOrUrn}/clone | Clone a dashboard
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /dashboards | Create a new dashboard
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /dashboards/{dashboardIdOrUrn} | Delete a dashboard
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /dashboards/{dashboardIdOrUrn} | Get a single dashboard
[**GetDashboards**](DashboardsApi.md#GetDashboards) | **Get** /dashboards | Get a list of dashboards
[**PatchDashboard**](DashboardsApi.md#PatchDashboard) | **Patch** /dashboards/{dashboardIdOrUrn} | Patch a dashboard



## CloneDashboard

> DashboardReadFullSchema CloneDashboard(ctx, dashboardIdOrUrn).DashboardCloneSchema(dashboardCloneSchema).Execute()

Clone a dashboard



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
    dashboardIdOrUrn := "dashboardIdOrUrn_example" // string | The identifier of a dashboard
    dashboardCloneSchema := *openapiclient.NewDashboardCloneSchema("Name_example") // DashboardCloneSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.CloneDashboard(context.Background(), dashboardIdOrUrn).DashboardCloneSchema(dashboardCloneSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CloneDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDashboard`: DashboardReadFullSchema
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CloneDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardIdOrUrn** | **string** | The identifier of a dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardCloneSchema** | [**DashboardCloneSchema**](DashboardCloneSchema.md) |  | 

### Return type

[**DashboardReadFullSchema**](DashboardReadFullSchema.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboard

> DashboardReadFullSchema CreateDashboard(ctx).DashboardWriteSchema(dashboardWriteSchema).Execute()

Create a new dashboard



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
    dashboardWriteSchema := *openapiclient.NewDashboardWriteSchema("Name_example", "Description_example", openapiclient.DashboardScope("publicDashboard"), *openapiclient.NewPersesDashboard()) // DashboardWriteSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.CreateDashboard(context.Background()).DashboardWriteSchema(dashboardWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: DashboardReadFullSchema
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardWriteSchema** | [**DashboardWriteSchema**](DashboardWriteSchema.md) |  | 

### Return type

[**DashboardReadFullSchema**](DashboardReadFullSchema.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx, dashboardIdOrUrn).Execute()

Delete a dashboard



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
    dashboardIdOrUrn := "dashboardIdOrUrn_example" // string | The identifier of a dashboard

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.DeleteDashboard(context.Background(), dashboardIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardIdOrUrn** | **string** | The identifier of a dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GetDashboard

> DashboardReadSchema GetDashboard(ctx, dashboardIdOrUrn).LoadFullDashboard(loadFullDashboard).Execute()

Get a single dashboard



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
    dashboardIdOrUrn := "dashboardIdOrUrn_example" // string | The identifier of a dashboard
    loadFullDashboard := true // bool | If true, includes the full dashboard content in the response. Defaults to false (only metadata is returned) for retrieval endpoints. Defaults to true for create/update operations. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.GetDashboard(context.Background(), dashboardIdOrUrn).LoadFullDashboard(loadFullDashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: DashboardReadSchema
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardIdOrUrn** | **string** | The identifier of a dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadFullDashboard** | **bool** | If true, includes the full dashboard content in the response. Defaults to false (only metadata is returned) for retrieval endpoints. Defaults to true for create/update operations. | [default to false]

### Return type

[**DashboardReadSchema**](DashboardReadSchema.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboards

> DashboardList GetDashboards(ctx).LoadFullDashboard(loadFullDashboard).Execute()

Get a list of dashboards



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
    loadFullDashboard := true // bool | If true, includes the full dashboard content in the response. Defaults to false (only metadata is returned) for retrieval endpoints. Defaults to true for create/update operations. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.GetDashboards(context.Background()).LoadFullDashboard(loadFullDashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboards`: DashboardList
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadFullDashboard** | **bool** | If true, includes the full dashboard content in the response. Defaults to false (only metadata is returned) for retrieval endpoints. Defaults to true for create/update operations. | [default to false]

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDashboard

> DashboardReadFullSchema PatchDashboard(ctx, dashboardIdOrUrn).DashboardPatchSchema(dashboardPatchSchema).Execute()

Patch a dashboard



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
    dashboardIdOrUrn := "dashboardIdOrUrn_example" // string | The identifier of a dashboard
    dashboardPatchSchema := *openapiclient.NewDashboardPatchSchema() // DashboardPatchSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.PatchDashboard(context.Background(), dashboardIdOrUrn).DashboardPatchSchema(dashboardPatchSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.PatchDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDashboard`: DashboardReadFullSchema
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.PatchDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardIdOrUrn** | **string** | The identifier of a dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardPatchSchema** | [**DashboardPatchSchema**](DashboardPatchSchema.md) |  | 

### Return type

[**DashboardReadFullSchema**](DashboardReadFullSchema.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

