# PromBatchQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]PromBatchQueryItem**](PromBatchQueryItem.md) | List of queries to execute. | 
**Timeout** | Pointer to **string** |  | [optional] 

## Methods

### NewPromBatchQueryRequest

`func NewPromBatchQueryRequest(queries []PromBatchQueryItem, ) *PromBatchQueryRequest`

NewPromBatchQueryRequest instantiates a new PromBatchQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromBatchQueryRequestWithDefaults

`func NewPromBatchQueryRequestWithDefaults() *PromBatchQueryRequest`

NewPromBatchQueryRequestWithDefaults instantiates a new PromBatchQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *PromBatchQueryRequest) GetQueries() []PromBatchQueryItem`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *PromBatchQueryRequest) GetQueriesOk() (*[]PromBatchQueryItem, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *PromBatchQueryRequest) SetQueries(v []PromBatchQueryItem)`

SetQueries sets Queries field to given value.


### GetTimeout

`func (o *PromBatchQueryRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PromBatchQueryRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PromBatchQueryRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PromBatchQueryRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


