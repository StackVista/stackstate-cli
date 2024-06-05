# SpanComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Identifier** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**HealthState** | [**HealthStateValue**](HealthStateValue.md) |  | 

## Methods

### NewSpanComponent

`func NewSpanComponent(id int64, identifier string, name string, type_ string, healthState HealthStateValue, ) *SpanComponent`

NewSpanComponent instantiates a new SpanComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanComponentWithDefaults

`func NewSpanComponentWithDefaults() *SpanComponent`

NewSpanComponentWithDefaults instantiates a new SpanComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpanComponent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpanComponent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpanComponent) SetId(v int64)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *SpanComponent) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SpanComponent) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SpanComponent) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *SpanComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanComponent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SpanComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpanComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpanComponent) SetType(v string)`

SetType sets Type field to given value.


### GetHealthState

`func (o *SpanComponent) GetHealthState() HealthStateValue`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *SpanComponent) GetHealthStateOk() (*HealthStateValue, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *SpanComponent) SetHealthState(v HealthStateValue)`

SetHealthState sets HealthState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


