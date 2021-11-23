# EventComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**EventType** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewEventComponent

`func NewEventComponent(type_ string, id int64, name string, ) *EventComponent`

NewEventComponent instantiates a new EventComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventComponentWithDefaults

`func NewEventComponentWithDefaults() *EventComponent`

NewEventComponentWithDefaults instantiates a new EventComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventComponent) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *EventComponent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventComponent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventComponent) SetId(v int64)`

SetId sets Id field to given value.


### GetEventType

`func (o *EventComponent) GetEventType() int64`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventComponent) GetEventTypeOk() (*int64, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventComponent) SetEventType(v int64)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventComponent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *EventComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventComponent) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

