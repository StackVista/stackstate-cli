# DurationProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StartDate** | **string** | Cel expression that returns a date | 
**EndDate** | Pointer to **string** | Cel expression that returns a date | [optional] 

## Methods

### NewDurationProjection

`func NewDurationProjection(type_ string, startDate string, ) *DurationProjection`

NewDurationProjection instantiates a new DurationProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationProjectionWithDefaults

`func NewDurationProjectionWithDefaults() *DurationProjection`

NewDurationProjectionWithDefaults instantiates a new DurationProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DurationProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DurationProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DurationProjection) SetType(v string)`

SetType sets Type field to given value.


### GetStartDate

`func (o *DurationProjection) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DurationProjection) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DurationProjection) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *DurationProjection) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DurationProjection) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DurationProjection) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DurationProjection) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


