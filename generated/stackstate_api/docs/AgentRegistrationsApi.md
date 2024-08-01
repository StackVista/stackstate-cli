# \AgentRegistrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllAgentRegistrations**](AgentRegistrationsApi.md#AllAgentRegistrations) | **Get** /agents | Overview of registered agents



## AllAgentRegistrations

> AgentRegistrations AllAgentRegistrations(ctx).Execute()

Overview of registered agents



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
    resp, r, err := apiClient.AgentRegistrationsApi.AllAgentRegistrations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationsApi.AllAgentRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllAgentRegistrations`: AgentRegistrations
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationsApi.AllAgentRegistrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAllAgentRegistrationsRequest struct via the builder pattern


### Return type

[**AgentRegistrations**](AgentRegistrations.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

