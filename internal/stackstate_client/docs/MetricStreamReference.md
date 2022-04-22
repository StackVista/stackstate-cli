# MetricStreamReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ElementIdentifiers** | **[]string** |  | 
**StreamNodeId** | **int64** |  | 

## Methods

### NewMetricStreamReference

`func NewMetricStreamReference(type_ string, elementIdentifiers []string, streamNodeId int64, ) *MetricStreamReference`

NewMetricStreamReference instantiates a new MetricStreamReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStreamReferenceWithDefaults

`func NewMetricStreamReferenceWithDefaults() *MetricStreamReference`

NewMetricStreamReferenceWithDefaults instantiates a new MetricStreamReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricStreamReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricStreamReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricStreamReference) SetType(v string)`

SetType sets Type field to given value.


### GetElementIdentifiers

`func (o *MetricStreamReference) GetElementIdentifiers() []string`

GetElementIdentifiers returns the ElementIdentifiers field if non-nil, zero value otherwise.

### GetElementIdentifiersOk

`func (o *MetricStreamReference) GetElementIdentifiersOk() (*[]string, bool)`

GetElementIdentifiersOk returns a tuple with the ElementIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIdentifiers

`func (o *MetricStreamReference) SetElementIdentifiers(v []string)`

SetElementIdentifiers sets ElementIdentifiers field to given value.


### GetStreamNodeId

`func (o *MetricStreamReference) GetStreamNodeId() int64`

GetStreamNodeId returns the StreamNodeId field if non-nil, zero value otherwise.

### GetStreamNodeIdOk

`func (o *MetricStreamReference) GetStreamNodeIdOk() (*int64, bool)`

GetStreamNodeIdOk returns a tuple with the StreamNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamNodeId

`func (o *MetricStreamReference) SetStreamNodeId(v int64)`

SetStreamNodeId sets StreamNodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


