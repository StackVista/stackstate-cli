# RatioProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Numerator** | **string** | Cel expression that returns a number | 
**Denominator** | **string** | Cel expression that returns a number | 
**Status** | Pointer to **string** | Cel expression that returns a string that represents a valid HealthState | [optional] 

## Methods

### NewRatioProjection

`func NewRatioProjection(type_ string, numerator string, denominator string, ) *RatioProjection`

NewRatioProjection instantiates a new RatioProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatioProjectionWithDefaults

`func NewRatioProjectionWithDefaults() *RatioProjection`

NewRatioProjectionWithDefaults instantiates a new RatioProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RatioProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatioProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatioProjection) SetType(v string)`

SetType sets Type field to given value.


### GetNumerator

`func (o *RatioProjection) GetNumerator() string`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *RatioProjection) GetNumeratorOk() (*string, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *RatioProjection) SetNumerator(v string)`

SetNumerator sets Numerator field to given value.


### GetDenominator

`func (o *RatioProjection) GetDenominator() string`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *RatioProjection) GetDenominatorOk() (*string, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *RatioProjection) SetDenominator(v string)`

SetDenominator sets Denominator field to given value.


### GetStatus

`func (o *RatioProjection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RatioProjection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RatioProjection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RatioProjection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


