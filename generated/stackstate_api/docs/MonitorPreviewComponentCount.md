# MonitorPreviewComponentCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueIdentifiers** | **int32** | The number of unique identifiers for the total health state count. Only available when there non-empty identifiers. | 
**MatchingComponents** | **int32** | The number of components for the identifiers: matchingComponents &lt;&#x3D; uniqueIdentifiers. Only available when there non-empty identifiers. | 

## Methods

### NewMonitorPreviewComponentCount

`func NewMonitorPreviewComponentCount(uniqueIdentifiers int32, matchingComponents int32, ) *MonitorPreviewComponentCount`

NewMonitorPreviewComponentCount instantiates a new MonitorPreviewComponentCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorPreviewComponentCountWithDefaults

`func NewMonitorPreviewComponentCountWithDefaults() *MonitorPreviewComponentCount`

NewMonitorPreviewComponentCountWithDefaults instantiates a new MonitorPreviewComponentCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueIdentifiers

`func (o *MonitorPreviewComponentCount) GetUniqueIdentifiers() int32`

GetUniqueIdentifiers returns the UniqueIdentifiers field if non-nil, zero value otherwise.

### GetUniqueIdentifiersOk

`func (o *MonitorPreviewComponentCount) GetUniqueIdentifiersOk() (*int32, bool)`

GetUniqueIdentifiersOk returns a tuple with the UniqueIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifiers

`func (o *MonitorPreviewComponentCount) SetUniqueIdentifiers(v int32)`

SetUniqueIdentifiers sets UniqueIdentifiers field to given value.


### GetMatchingComponents

`func (o *MonitorPreviewComponentCount) GetMatchingComponents() int32`

GetMatchingComponents returns the MatchingComponents field if non-nil, zero value otherwise.

### GetMatchingComponentsOk

`func (o *MonitorPreviewComponentCount) GetMatchingComponentsOk() (*int32, bool)`

GetMatchingComponentsOk returns a tuple with the MatchingComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingComponents

`func (o *MonitorPreviewComponentCount) SetMatchingComponents(v int32)`

SetMatchingComponents sets MatchingComponents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


