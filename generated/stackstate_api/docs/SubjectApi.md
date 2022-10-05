# \SubjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubject**](SubjectApi.md#CreateSubject) | **Put** /security/subjects/{subject} | Create a subject



## CreateSubject

> CreateSubject(ctx, subject).CreateSubject(createSubject).Execute()

Create a subject



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
    subject := "subject_example" // string |
    createSubject := *openapiclient.NewCreateSubject("Query_example", openapiclient.Version("0.0.1")) // CreateSubject |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectApi.CreateSubject(context.Background(), subject).CreateSubject(createSubject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectApi.CreateSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSubject** | [**CreateSubject**](CreateSubject.md) |  |

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
