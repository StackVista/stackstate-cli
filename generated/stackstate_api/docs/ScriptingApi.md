# \ScriptingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScriptExecute**](ScriptingAPI.md#ScriptExecute) | **Post** /script/execute | Execute script



## ScriptExecute

> ExecuteScriptResponse ScriptExecute(ctx).ExecuteScriptRequest(executeScriptRequest).Execute()

Execute script



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
    executeScriptRequest := *openapiclient.NewExecuteScriptRequest("Script_example") // ExecuteScriptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScriptingAPI.ScriptExecute(context.Background()).ExecuteScriptRequest(executeScriptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScriptingAPI.ScriptExecute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScriptExecute`: ExecuteScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ScriptingAPI.ScriptExecute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScriptExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeScriptRequest** | [**ExecuteScriptRequest**](ExecuteScriptRequest.md) |  | 

### Return type

[**ExecuteScriptResponse**](ExecuteScriptResponse.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

