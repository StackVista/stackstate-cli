# EventComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**TypeName** | **string** |  | 
**Name** | **string** |  | 
**Identifiers** | **[]string** |  | 
**Iconbase64** | Pointer to **string** |  | [optional] 

## Methods

### NewEventComponent

`func NewEventComponent(type_ string, id int64, typeName string, name string, identifiers []string, ) *EventComponent`

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


### GetTypeName

`func (o *EventComponent) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EventComponent) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EventComponent) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


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


### GetIdentifiers

`func (o *EventComponent) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *EventComponent) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *EventComponent) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetIconbase64

`func (o *EventComponent) GetIconbase64() string`

GetIconbase64 returns the Iconbase64 field if non-nil, zero value otherwise.

### GetIconbase64Ok

`func (o *EventComponent) GetIconbase64Ok() (*string, bool)`

GetIconbase64Ok returns a tuple with the Iconbase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconbase64

`func (o *EventComponent) SetIconbase64(v string)`

SetIconbase64 sets Iconbase64 field to given value.

### HasIconbase64

`func (o *EventComponent) HasIconbase64() bool`

HasIconbase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


