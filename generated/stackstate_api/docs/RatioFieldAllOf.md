# RatioFieldAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Numerator** | **float32** |  | 
**Denominator** | **float32** |  | 
**Status** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 

## Methods

### NewRatioFieldAllOf

`func NewRatioFieldAllOf(type_ string, numerator float32, denominator float32, ) *RatioFieldAllOf`

NewRatioFieldAllOf instantiates a new RatioFieldAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatioFieldAllOfWithDefaults

`func NewRatioFieldAllOfWithDefaults() *RatioFieldAllOf`

NewRatioFieldAllOfWithDefaults instantiates a new RatioFieldAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RatioFieldAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatioFieldAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatioFieldAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetNumerator

`func (o *RatioFieldAllOf) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *RatioFieldAllOf) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *RatioFieldAllOf) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.


### GetDenominator

`func (o *RatioFieldAllOf) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *RatioFieldAllOf) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *RatioFieldAllOf) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.


### GetStatus

`func (o *RatioFieldAllOf) GetStatus() HealthStateValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RatioFieldAllOf) GetStatusOk() (*HealthStateValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RatioFieldAllOf) SetStatus(v HealthStateValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RatioFieldAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


