# EventsHistogramRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopologyTimestamp** | Pointer to **int32** |  | [optional] 
**StartTimestampMs** | **int32** |  | 
**EndTimestampMs** | **int32** |  | 
**TopologyQuery** | **string** |  | 
**IncludeConnectedComponents** | Pointer to **bool** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**EventCategories** | Pointer to [**[]EventCategory**](EventCategory.md) |  | [optional] 
**HistogramBucketCount** | **int32** |  | 

## Methods

### NewEventsHistogramRequest

`func NewEventsHistogramRequest(startTimestampMs int32, endTimestampMs int32, topologyQuery string, histogramBucketCount int32, ) *EventsHistogramRequest`

NewEventsHistogramRequest instantiates a new EventsHistogramRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsHistogramRequestWithDefaults

`func NewEventsHistogramRequestWithDefaults() *EventsHistogramRequest`

NewEventsHistogramRequestWithDefaults instantiates a new EventsHistogramRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopologyTimestamp

`func (o *EventsHistogramRequest) GetTopologyTimestamp() int32`

GetTopologyTimestamp returns the TopologyTimestamp field if non-nil, zero value otherwise.

### GetTopologyTimestampOk

`func (o *EventsHistogramRequest) GetTopologyTimestampOk() (*int32, bool)`

GetTopologyTimestampOk returns a tuple with the TopologyTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyTimestamp

`func (o *EventsHistogramRequest) SetTopologyTimestamp(v int32)`

SetTopologyTimestamp sets TopologyTimestamp field to given value.

### HasTopologyTimestamp

`func (o *EventsHistogramRequest) HasTopologyTimestamp() bool`

HasTopologyTimestamp returns a boolean if a field has been set.

### GetStartTimestampMs

`func (o *EventsHistogramRequest) GetStartTimestampMs() int32`

GetStartTimestampMs returns the StartTimestampMs field if non-nil, zero value otherwise.

### GetStartTimestampMsOk

`func (o *EventsHistogramRequest) GetStartTimestampMsOk() (*int32, bool)`

GetStartTimestampMsOk returns a tuple with the StartTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMs

`func (o *EventsHistogramRequest) SetStartTimestampMs(v int32)`

SetStartTimestampMs sets StartTimestampMs field to given value.


### GetEndTimestampMs

`func (o *EventsHistogramRequest) GetEndTimestampMs() int32`

GetEndTimestampMs returns the EndTimestampMs field if non-nil, zero value otherwise.

### GetEndTimestampMsOk

`func (o *EventsHistogramRequest) GetEndTimestampMsOk() (*int32, bool)`

GetEndTimestampMsOk returns a tuple with the EndTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMs

`func (o *EventsHistogramRequest) SetEndTimestampMs(v int32)`

SetEndTimestampMs sets EndTimestampMs field to given value.


### GetTopologyQuery

`func (o *EventsHistogramRequest) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *EventsHistogramRequest) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *EventsHistogramRequest) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.


### GetIncludeConnectedComponents

`func (o *EventsHistogramRequest) GetIncludeConnectedComponents() bool`

GetIncludeConnectedComponents returns the IncludeConnectedComponents field if non-nil, zero value otherwise.

### GetIncludeConnectedComponentsOk

`func (o *EventsHistogramRequest) GetIncludeConnectedComponentsOk() (*bool, bool)`

GetIncludeConnectedComponentsOk returns a tuple with the IncludeConnectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeConnectedComponents

`func (o *EventsHistogramRequest) SetIncludeConnectedComponents(v bool)`

SetIncludeConnectedComponents sets IncludeConnectedComponents field to given value.

### HasIncludeConnectedComponents

`func (o *EventsHistogramRequest) HasIncludeConnectedComponents() bool`

HasIncludeConnectedComponents returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventsHistogramRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventsHistogramRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventsHistogramRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventsHistogramRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEventCategories

`func (o *EventsHistogramRequest) GetEventCategories() []EventCategory`

GetEventCategories returns the EventCategories field if non-nil, zero value otherwise.

### GetEventCategoriesOk

`func (o *EventsHistogramRequest) GetEventCategoriesOk() (*[]EventCategory, bool)`

GetEventCategoriesOk returns a tuple with the EventCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategories

`func (o *EventsHistogramRequest) SetEventCategories(v []EventCategory)`

SetEventCategories sets EventCategories field to given value.

### HasEventCategories

`func (o *EventsHistogramRequest) HasEventCategories() bool`

HasEventCategories returns a boolean if a field has been set.

### GetHistogramBucketCount

`func (o *EventsHistogramRequest) GetHistogramBucketCount() int32`

GetHistogramBucketCount returns the HistogramBucketCount field if non-nil, zero value otherwise.

### GetHistogramBucketCountOk

`func (o *EventsHistogramRequest) GetHistogramBucketCountOk() (*int32, bool)`

GetHistogramBucketCountOk returns a tuple with the HistogramBucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBucketCount

`func (o *EventsHistogramRequest) SetHistogramBucketCount(v int32)`

SetHistogramBucketCount sets HistogramBucketCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


