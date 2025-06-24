# EventElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**TypeName** | **string** |  | 
**Name** | **string** |  | 
**Identifiers** | **[]string** |  | 
**Iconbase64** | Pointer to **string** |  | [optional] 
**Source** | [**EventComponent**](EventComponent.md) |  | 
**Target** | [**EventComponent**](EventComponent.md) |  | 
**DependencyDirection** | [**DependencyDirection**](DependencyDirection.md) |  | 

## Methods

### NewEventElement

`func NewEventElement(type_ string, id int64, typeName string, name string, identifiers []string, source EventComponent, target EventComponent, dependencyDirection DependencyDirection, ) *EventElement`

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


### GetTypeName

`func (o *EventElement) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EventElement) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EventElement) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


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


### GetIdentifiers

`func (o *EventElement) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *EventElement) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *EventElement) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetIconbase64

`func (o *EventElement) GetIconbase64() string`

GetIconbase64 returns the Iconbase64 field if non-nil, zero value otherwise.

### GetIconbase64Ok

`func (o *EventElement) GetIconbase64Ok() (*string, bool)`

GetIconbase64Ok returns a tuple with the Iconbase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconbase64

`func (o *EventElement) SetIconbase64(v string)`

SetIconbase64 sets Iconbase64 field to given value.

### HasIconbase64

`func (o *EventElement) HasIconbase64() bool`

HasIconbase64 returns a boolean if a field has been set.

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


