# EventRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**EventType** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Source** | [**EventComponent**](EventComponent.md) |  | 
**Target** | [**EventComponent**](EventComponent.md) |  | 
**DependencyDirection** | [**DependencyDirection**](DependencyDirection.md) |  | 

## Methods

### NewEventRelation

`func NewEventRelation(type_ string, id int64, source EventComponent, target EventComponent, dependencyDirection DependencyDirection, ) *EventRelation`

NewEventRelation instantiates a new EventRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRelationWithDefaults

`func NewEventRelationWithDefaults() *EventRelation`

NewEventRelationWithDefaults instantiates a new EventRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventRelation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventRelation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventRelation) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *EventRelation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRelation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRelation) SetId(v int64)`

SetId sets Id field to given value.


### GetEventType

`func (o *EventRelation) GetEventType() int64`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventRelation) GetEventTypeOk() (*int64, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventRelation) SetEventType(v int64)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventRelation) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *EventRelation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventRelation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventRelation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventRelation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *EventRelation) GetSource() EventComponent`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventRelation) GetSourceOk() (*EventComponent, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventRelation) SetSource(v EventComponent)`

SetSource sets Source field to given value.


### GetTarget

`func (o *EventRelation) GetTarget() EventComponent`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EventRelation) GetTargetOk() (*EventComponent, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EventRelation) SetTarget(v EventComponent)`

SetTarget sets Target field to given value.


### GetDependencyDirection

`func (o *EventRelation) GetDependencyDirection() DependencyDirection`

GetDependencyDirection returns the DependencyDirection field if non-nil, zero value otherwise.

### GetDependencyDirectionOk

`func (o *EventRelation) GetDependencyDirectionOk() (*DependencyDirection, bool)`

GetDependencyDirectionOk returns a tuple with the DependencyDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDirection

`func (o *EventRelation) SetDependencyDirection(v DependencyDirection)`

SetDependencyDirection sets DependencyDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

