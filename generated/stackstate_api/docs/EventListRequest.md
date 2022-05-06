# EventListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestampMs** | **int32** |  | 
**EndTimestampMs** | **int32** |  | 
**TopologyQuery** | **string** |  | 
**Limit** | **int32** |  | 
**PlayHeadTimestampMs** | Pointer to **int32** |  | [optional] 
**RootCauseMode** | Pointer to [**RootCauseMode**](RootCauseMode.md) |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**EventTags** | Pointer to **[]string** |  | [optional] 
**EventCategories** | Pointer to [**[]EventCategory**](EventCategory.md) |  | [optional] 
**EventSources** | Pointer to **[]string** |  | [optional] 
**Cursor** | Pointer to [**EventCursor**](EventCursor.md) |  | [optional] 

## Methods

### NewEventListRequest

`func NewEventListRequest(startTimestampMs int32, endTimestampMs int32, topologyQuery string, limit int32, ) *EventListRequest`

NewEventListRequest instantiates a new EventListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventListRequestWithDefaults

`func NewEventListRequestWithDefaults() *EventListRequest`

NewEventListRequestWithDefaults instantiates a new EventListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimestampMs

`func (o *EventListRequest) GetStartTimestampMs() int32`

GetStartTimestampMs returns the StartTimestampMs field if non-nil, zero value otherwise.

### GetStartTimestampMsOk

`func (o *EventListRequest) GetStartTimestampMsOk() (*int32, bool)`

GetStartTimestampMsOk returns a tuple with the StartTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMs

`func (o *EventListRequest) SetStartTimestampMs(v int32)`

SetStartTimestampMs sets StartTimestampMs field to given value.


### GetEndTimestampMs

`func (o *EventListRequest) GetEndTimestampMs() int32`

GetEndTimestampMs returns the EndTimestampMs field if non-nil, zero value otherwise.

### GetEndTimestampMsOk

`func (o *EventListRequest) GetEndTimestampMsOk() (*int32, bool)`

GetEndTimestampMsOk returns a tuple with the EndTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMs

`func (o *EventListRequest) SetEndTimestampMs(v int32)`

SetEndTimestampMs sets EndTimestampMs field to given value.


### GetTopologyQuery

`func (o *EventListRequest) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *EventListRequest) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *EventListRequest) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.


### GetLimit

`func (o *EventListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EventListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EventListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPlayHeadTimestampMs

`func (o *EventListRequest) GetPlayHeadTimestampMs() int32`

GetPlayHeadTimestampMs returns the PlayHeadTimestampMs field if non-nil, zero value otherwise.

### GetPlayHeadTimestampMsOk

`func (o *EventListRequest) GetPlayHeadTimestampMsOk() (*int32, bool)`

GetPlayHeadTimestampMsOk returns a tuple with the PlayHeadTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayHeadTimestampMs

`func (o *EventListRequest) SetPlayHeadTimestampMs(v int32)`

SetPlayHeadTimestampMs sets PlayHeadTimestampMs field to given value.

### HasPlayHeadTimestampMs

`func (o *EventListRequest) HasPlayHeadTimestampMs() bool`

HasPlayHeadTimestampMs returns a boolean if a field has been set.

### GetRootCauseMode

`func (o *EventListRequest) GetRootCauseMode() RootCauseMode`

GetRootCauseMode returns the RootCauseMode field if non-nil, zero value otherwise.

### GetRootCauseModeOk

`func (o *EventListRequest) GetRootCauseModeOk() (*RootCauseMode, bool)`

GetRootCauseModeOk returns a tuple with the RootCauseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseMode

`func (o *EventListRequest) SetRootCauseMode(v RootCauseMode)`

SetRootCauseMode sets RootCauseMode field to given value.

### HasRootCauseMode

`func (o *EventListRequest) HasRootCauseMode() bool`

HasRootCauseMode returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventListRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventListRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventListRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventListRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEventTags

`func (o *EventListRequest) GetEventTags() []string`

GetEventTags returns the EventTags field if non-nil, zero value otherwise.

### GetEventTagsOk

`func (o *EventListRequest) GetEventTagsOk() (*[]string, bool)`

GetEventTagsOk returns a tuple with the EventTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTags

`func (o *EventListRequest) SetEventTags(v []string)`

SetEventTags sets EventTags field to given value.

### HasEventTags

`func (o *EventListRequest) HasEventTags() bool`

HasEventTags returns a boolean if a field has been set.

### GetEventCategories

`func (o *EventListRequest) GetEventCategories() []EventCategory`

GetEventCategories returns the EventCategories field if non-nil, zero value otherwise.

### GetEventCategoriesOk

`func (o *EventListRequest) GetEventCategoriesOk() (*[]EventCategory, bool)`

GetEventCategoriesOk returns a tuple with the EventCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategories

`func (o *EventListRequest) SetEventCategories(v []EventCategory)`

SetEventCategories sets EventCategories field to given value.

### HasEventCategories

`func (o *EventListRequest) HasEventCategories() bool`

HasEventCategories returns a boolean if a field has been set.

### GetEventSources

`func (o *EventListRequest) GetEventSources() []string`

GetEventSources returns the EventSources field if non-nil, zero value otherwise.

### GetEventSourcesOk

`func (o *EventListRequest) GetEventSourcesOk() (*[]string, bool)`

GetEventSourcesOk returns a tuple with the EventSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSources

`func (o *EventListRequest) SetEventSources(v []string)`

SetEventSources sets EventSources field to given value.

### HasEventSources

`func (o *EventListRequest) HasEventSources() bool`

HasEventSources returns a boolean if a field has been set.

### GetCursor

`func (o *EventListRequest) GetCursor() EventCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *EventListRequest) GetCursorOk() (*EventCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *EventListRequest) SetCursor(v EventCursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *EventListRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


