# EventCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastEventTimestampMs** | **int64** |  | 
**LastEventId** | **string** |  | 

## Methods

### NewEventCursor

`func NewEventCursor(lastEventTimestampMs int64, lastEventId string, ) *EventCursor`

NewEventCursor instantiates a new EventCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCursorWithDefaults

`func NewEventCursorWithDefaults() *EventCursor`

NewEventCursorWithDefaults instantiates a new EventCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastEventTimestampMs

`func (o *EventCursor) GetLastEventTimestampMs() int64`

GetLastEventTimestampMs returns the LastEventTimestampMs field if non-nil, zero value otherwise.

### GetLastEventTimestampMsOk

`func (o *EventCursor) GetLastEventTimestampMsOk() (*int64, bool)`

GetLastEventTimestampMsOk returns a tuple with the LastEventTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTimestampMs

`func (o *EventCursor) SetLastEventTimestampMs(v int64)`

SetLastEventTimestampMs sets LastEventTimestampMs field to given value.


### GetLastEventId

`func (o *EventCursor) GetLastEventId() string`

GetLastEventId returns the LastEventId field if non-nil, zero value otherwise.

### GetLastEventIdOk

`func (o *EventCursor) GetLastEventIdOk() (*string, bool)`

GetLastEventIdOk returns a tuple with the LastEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventId

`func (o *EventCursor) SetLastEventId(v string)`

SetLastEventId sets LastEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


