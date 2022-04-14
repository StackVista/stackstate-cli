# \ExportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSettings**](ExportApi.md#ExportSettings) | **Post** /export | Export settings



## ExportSettings

> string ExportSettings(ctx).Export(export).Execute()

Export settings



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
    export := *openapiclient.NewExport() // Export | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportSettings(context.Background()).Export(export).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSettings`: string
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export** | [**Export**](Export.md) |  | 

### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

