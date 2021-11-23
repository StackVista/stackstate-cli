# EventRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**EventId** | **string** |  | 
**EventTimestamp** | **int64** |  | 
**EventType** | **string** |  | 

## Methods

### NewEventRef

`func NewEventRef(title string, eventId string, eventTimestamp int64, eventType string, ) *EventRef`

NewEventRef instantiates a new EventRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRefWithDefaults

`func NewEventRefWithDefaults() *EventRef`

NewEventRefWithDefaults instantiates a new EventRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *EventRef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventRef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventRef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEventId

`func (o *EventRef) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventRef) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventRef) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventTimestamp

`func (o *EventRef) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *EventRef) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *EventRef) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetEventType

`func (o *EventRef) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventRef) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventRef) SetEventType(v string)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


