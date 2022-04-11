# \ImportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportSettings**](ImportApi.md#ImportSettings) | **Post** /import | Import settings



## ImportSettings

> []map[string]interface{} ImportSettings(ctx).Body(body).TimeoutSeconds(timeoutSeconds).Namespace(namespace).Unlocked(unlocked).Execute()

Import settings



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
    body := "body_example" // string | 
    timeoutSeconds := int32(56) // int32 |  (optional)
    namespace := "namespace_example" // string |  (optional)
    unlocked := "unlocked_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.ImportSettings(context.Background()).Body(body).TimeoutSeconds(timeoutSeconds).Namespace(namespace).Unlocked(unlocked).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ImportSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSettings`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ImportSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 
 **timeoutSeconds** | **int32** |  | 
 **namespace** | **string** |  | 
 **unlocked** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: plain/text
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

