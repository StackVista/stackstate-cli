# \KubernetesLogsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesLogs**](KubernetesLogsApi.md#GetKubernetesLogs) | **Get** /k8s/logs | Get Kubernetes logs
[**GetKubernetesLogsAutocomplete**](KubernetesLogsApi.md#GetKubernetesLogsAutocomplete) | **Get** /k8s/logs/autocomplete | Get Kubernetes logs autocomplete values
[**GetKubernetesLogsHistogram**](KubernetesLogsApi.md#GetKubernetesLogsHistogram) | **Get** /k8s/logs/histogram | Get Kubernetes logs histogram



## GetKubernetesLogs

> GetKubernetesLogsResult GetKubernetesLogs(ctx).From(from).To(to).PodUID(podUID).Cluster(cluster).PageSize(pageSize).Page(page).Query(query).ContainerNames(containerNames).Direction(direction).Severity(severity).TopologyTime(topologyTime).Execute()

Get Kubernetes logs

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
    from := int32(56) // int32 | Logs initial timestamp.
    to := int32(56) // int32 | Logs last timestamp.
    podUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Find only logs for the given pod UID.
    cluster := "cluster_example" // string | Cluster name which identifies the scope of the RBAC data
    pageSize := int32(56) // int32 | Maximum number of the log lines in the result. (optional) (default to 25)
    page := int32(56) // int32 | The page for which the log lines of pageSize must be returned. (optional) (default to 1)
    query := "query_example" // string | Find only logs containing query text. (optional)
    containerNames := []string{"Inner_example"} // []string | Find only logs for the given container names. (optional)
    direction := openapiclient.LogsDirection("NEWEST") // LogsDirection | Fetch Oldest or Newest first. (optional) (default to "NEWEST")
    severity := []openapiclient.LogSeverity{openapiclient.LogSeverity("WARNING")} // []LogSeverity | Search a specific log severity WARN, ERROR, OTHER. (optional)
    topologyTime := int32(56) // int32 | A timestamp at which the pod existed. If not given the pod is queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesLogsApi.GetKubernetesLogs(context.Background()).From(from).To(to).PodUID(podUID).Cluster(cluster).PageSize(pageSize).Page(page).Query(query).ContainerNames(containerNames).Direction(direction).Severity(severity).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesLogsApi.GetKubernetesLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesLogs`: GetKubernetesLogsResult
    fmt.Fprintf(os.Stdout, "Response from `KubernetesLogsApi.GetKubernetesLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Logs initial timestamp. | 
 **to** | **int32** | Logs last timestamp. | 
 **podUID** | **string** | Find only logs for the given pod UID. | 
 **cluster** | **string** | Cluster name which identifies the scope of the RBAC data | 
 **pageSize** | **int32** | Maximum number of the log lines in the result. | [default to 25]
 **page** | **int32** | The page for which the log lines of pageSize must be returned. | [default to 1]
 **query** | **string** | Find only logs containing query text. | 
 **containerNames** | **[]string** | Find only logs for the given container names. | 
 **direction** | [**LogsDirection**](LogsDirection.md) | Fetch Oldest or Newest first. | [default to &quot;NEWEST&quot;]
 **severity** | [**[]LogSeverity**](LogSeverity.md) | Search a specific log severity WARN, ERROR, OTHER. | 
 **topologyTime** | **int32** | A timestamp at which the pod existed. If not given the pod is queried at current time. | 

### Return type

[**GetKubernetesLogsResult**](GetKubernetesLogsResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesLogsAutocomplete

> GetKubernetesLogsAutocompleteResult GetKubernetesLogsAutocomplete(ctx).From(from).To(to).PodUID(podUID).Cluster(cluster).TopologyTime(topologyTime).Execute()

Get Kubernetes logs autocomplete values

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
    from := int32(56) // int32 | Logs initial timestamp.
    to := int32(56) // int32 | Logs last timestamp.
    podUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Find only logs for the given pod UID.
    cluster := "cluster_example" // string | Cluster name which identifies the scope of the RBAC data
    topologyTime := int32(56) // int32 | A timestamp at which the pod existed. If not given the pod is queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesLogsApi.GetKubernetesLogsAutocomplete(context.Background()).From(from).To(to).PodUID(podUID).Cluster(cluster).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesLogsApi.GetKubernetesLogsAutocomplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesLogsAutocomplete`: GetKubernetesLogsAutocompleteResult
    fmt.Fprintf(os.Stdout, "Response from `KubernetesLogsApi.GetKubernetesLogsAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesLogsAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Logs initial timestamp. | 
 **to** | **int32** | Logs last timestamp. | 
 **podUID** | **string** | Find only logs for the given pod UID. | 
 **cluster** | **string** | Cluster name which identifies the scope of the RBAC data | 
 **topologyTime** | **int32** | A timestamp at which the pod existed. If not given the pod is queried at current time. | 

### Return type

[**GetKubernetesLogsAutocompleteResult**](GetKubernetesLogsAutocompleteResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesLogsHistogram

> GetKubernetesLogsHistogramResult GetKubernetesLogsHistogram(ctx).From(from).To(to).PodUID(podUID).Cluster(cluster).BucketsCount(bucketsCount).Query(query).ContainerNames(containerNames).Severity(severity).TopologyTime(topologyTime).Execute()

Get Kubernetes logs histogram

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
    from := int32(56) // int32 | Logs initial timestamp.
    to := int32(56) // int32 | Logs last timestamp.
    podUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Find only logs for the given pod UID.
    cluster := "cluster_example" // string | Cluster name which identifies the scope of the RBAC data
    bucketsCount := int32(56) // int32 | The number of histogram buckets.
    query := "query_example" // string | Find only logs containing query text. (optional)
    containerNames := []string{"Inner_example"} // []string | Find only logs for the given container names. (optional)
    severity := []openapiclient.LogSeverity{openapiclient.LogSeverity("WARNING")} // []LogSeverity | Search a specific log severity WARN, ERROR, OTHER. (optional)
    topologyTime := int32(56) // int32 | A timestamp at which the pod existed. If not given the pod is queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesLogsApi.GetKubernetesLogsHistogram(context.Background()).From(from).To(to).PodUID(podUID).Cluster(cluster).BucketsCount(bucketsCount).Query(query).ContainerNames(containerNames).Severity(severity).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesLogsApi.GetKubernetesLogsHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesLogsHistogram`: GetKubernetesLogsHistogramResult
    fmt.Fprintf(os.Stdout, "Response from `KubernetesLogsApi.GetKubernetesLogsHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesLogsHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Logs initial timestamp. | 
 **to** | **int32** | Logs last timestamp. | 
 **podUID** | **string** | Find only logs for the given pod UID. | 
 **cluster** | **string** | Cluster name which identifies the scope of the RBAC data | 
 **bucketsCount** | **int32** | The number of histogram buckets. | 
 **query** | **string** | Find only logs containing query text. | 
 **containerNames** | **[]string** | Find only logs for the given container names. | 
 **severity** | [**[]LogSeverity**](LogSeverity.md) | Search a specific log severity WARN, ERROR, OTHER. | 
 **topologyTime** | **int32** | A timestamp at which the pod existed. If not given the pod is queried at current time. | 

### Return type

[**GetKubernetesLogsHistogramResult**](GetKubernetesLogsHistogramResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

