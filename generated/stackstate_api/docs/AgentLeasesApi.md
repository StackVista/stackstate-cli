# \AgentLeasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCheckLease**](AgentLeasesApi.md#AgentCheckLease) | **Post** /agents/{agentId}/checkLease | Check the lease of an agent.



## AgentCheckLease

> AgentRegistration AgentCheckLease(ctx, agentId).CheckLeaseRequest(checkLeaseRequest).Execute()

Check the lease of an agent.



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
    agentId := "agentId_example" // string | The identifier of an agent
    checkLeaseRequest := *openapiclient.NewCheckLeaseRequest("ApiKey_example") // CheckLeaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentLeasesApi.AgentCheckLease(context.Background(), agentId).CheckLeaseRequest(checkLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentLeasesApi.AgentCheckLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentCheckLease`: AgentRegistration
    fmt.Fprintf(os.Stdout, "Response from `AgentLeasesApi.AgentCheckLease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | The identifier of an agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCheckLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkLeaseRequest** | [**CheckLeaseRequest**](CheckLeaseRequest.md) |  | 

### Return type

[**AgentRegistration**](AgentRegistration.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

