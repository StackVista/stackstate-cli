# TopologyStreamListItemWithErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**TopologyStreamListItem**](TopologyStreamListItem.md) |  | 
**ErrorDetails** | [**[]TopologyStreamError**](TopologyStreamError.md) |  | 
**Metrics** | Pointer to [**TopologyStreamMetrics**](TopologyStreamMetrics.md) |  | [optional] 

## Methods

### NewTopologyStreamListItemWithErrorDetails

`func NewTopologyStreamListItemWithErrorDetails(item TopologyStreamListItem, errorDetails []TopologyStreamError, ) *TopologyStreamListItemWithErrorDetails`

NewTopologyStreamListItemWithErrorDetails instantiates a new TopologyStreamListItemWithErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyStreamListItemWithErrorDetailsWithDefaults

`func NewTopologyStreamListItemWithErrorDetailsWithDefaults() *TopologyStreamListItemWithErrorDetails`

NewTopologyStreamListItemWithErrorDetailsWithDefaults instantiates a new TopologyStreamListItemWithErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *TopologyStreamListItemWithErrorDetails) GetItem() TopologyStreamListItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TopologyStreamListItemWithErrorDetails) GetItemOk() (*TopologyStreamListItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TopologyStreamListItemWithErrorDetails) SetItem(v TopologyStreamListItem)`

SetItem sets Item field to given value.


### GetErrorDetails

`func (o *TopologyStreamListItemWithErrorDetails) GetErrorDetails() []TopologyStreamError`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *TopologyStreamListItemWithErrorDetails) GetErrorDetailsOk() (*[]TopologyStreamError, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *TopologyStreamListItemWithErrorDetails) SetErrorDetails(v []TopologyStreamError)`

SetErrorDetails sets ErrorDetails field to given value.


### GetMetrics

`func (o *TopologyStreamListItemWithErrorDetails) GetMetrics() TopologyStreamMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TopologyStreamListItemWithErrorDetails) GetMetricsOk() (*TopologyStreamMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TopologyStreamListItemWithErrorDetails) SetMetrics(v TopologyStreamMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TopologyStreamListItemWithErrorDetails) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


