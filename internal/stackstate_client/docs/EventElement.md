# EventElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**EventType** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Source** | [**EventComponent**](EventComponent.md) |  | 
**Target** | [**EventComponent**](EventComponent.md) |  | 
**DependencyDirection** | [**DependencyDirection**](DependencyDirection.md) |  | 

## Methods

### NewEventElement

`func NewEventElement(type_ string, id int64, name string, source EventComponent, target EventComponent, dependencyDirection DependencyDirection, ) *EventElement`

NewEventElement instantiates a new EventElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventElementWithDefaults

`func NewEventElementWithDefaults() *EventElement`

NewEventElementWithDefaults instantiates a new EventElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventElement) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *EventElement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventElement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventElement) SetId(v int64)`

SetId sets Id field to given value.


### GetEventType

`func (o *EventElement) GetEventType() int64`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventElement) GetEventTypeOk() (*int64, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventElement) SetEventType(v int64)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventElement) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *EventElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventElement) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *EventElement) GetSource() EventComponent`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventElement) GetSourceOk() (*EventComponent, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventElement) SetSource(v EventComponent)`

SetSource sets Source field to given value.


### GetTarget

`func (o *EventElement) GetTarget() EventComponent`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EventElement) GetTargetOk() (*EventComponent, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EventElement) SetTarget(v EventComponent)`

SetTarget sets Target field to given value.


### GetDependencyDirection

`func (o *EventElement) GetDependencyDirection() DependencyDirection`

GetDependencyDirection returns the DependencyDirection field if non-nil, zero value otherwise.

### GetDependencyDirectionOk

`func (o *EventElement) GetDependencyDirectionOk() (*DependencyDirection, bool)`

GetDependencyDirectionOk returns a tuple with the DependencyDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDirection

`func (o *EventElement) SetDependencyDirection(v DependencyDirection)`

SetDependencyDirection sets DependencyDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

