# NumericProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | Cel expression that returns a number | 
**Unit** | Pointer to **NullableString** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 

## Methods

### NewNumericProjection

`func NewNumericProjection(type_ string, value string, ) *NumericProjection`

NewNumericProjection instantiates a new NumericProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericProjectionWithDefaults

`func NewNumericProjectionWithDefaults() *NumericProjection`

NewNumericProjectionWithDefaults instantiates a new NumericProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NumericProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumericProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumericProjection) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *NumericProjection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumericProjection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumericProjection) SetValue(v string)`

SetValue sets Value field to given value.


### GetUnit

`func (o *NumericProjection) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *NumericProjection) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *NumericProjection) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *NumericProjection) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *NumericProjection) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *NumericProjection) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDecimalPlaces

`func (o *NumericProjection) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *NumericProjection) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *NumericProjection) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *NumericProjection) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


