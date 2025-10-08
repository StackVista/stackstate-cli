# OtelMappingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**OtelMappingStatusItem**](OtelMappingStatusItem.md) |  | 
**ErrorDetails** | [**[]OtelMappingError**](OtelMappingError.md) |  | 
**Metrics** | Pointer to [**OtelMappingMetrics**](OtelMappingMetrics.md) |  | [optional] 

## Methods

### NewOtelMappingStatus

`func NewOtelMappingStatus(item OtelMappingStatusItem, errorDetails []OtelMappingError, ) *OtelMappingStatus`

NewOtelMappingStatus instantiates a new OtelMappingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelMappingStatusWithDefaults

`func NewOtelMappingStatusWithDefaults() *OtelMappingStatus`

NewOtelMappingStatusWithDefaults instantiates a new OtelMappingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *OtelMappingStatus) GetItem() OtelMappingStatusItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OtelMappingStatus) GetItemOk() (*OtelMappingStatusItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OtelMappingStatus) SetItem(v OtelMappingStatusItem)`

SetItem sets Item field to given value.


### GetErrorDetails

`func (o *OtelMappingStatus) GetErrorDetails() []OtelMappingError`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *OtelMappingStatus) GetErrorDetailsOk() (*[]OtelMappingError, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *OtelMappingStatus) SetErrorDetails(v []OtelMappingError)`

SetErrorDetails sets ErrorDetails field to given value.


### GetMetrics

`func (o *OtelMappingStatus) GetMetrics() OtelMappingMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OtelMappingStatus) GetMetricsOk() (*OtelMappingMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OtelMappingStatus) SetMetrics(v OtelMappingMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OtelMappingStatus) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


