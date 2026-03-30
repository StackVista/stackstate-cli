# DurationProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StartTime** | **string** | Cel expression that returns an ISO8601 formatted timestamp | 
**EndTime** | Pointer to **string** | Cel expression that returns an ISO8601 formatted timestamp | [optional] 

## Methods

### NewDurationProjection

`func NewDurationProjection(type_ string, startTime string, ) *DurationProjection`

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


### GetStartTime

`func (o *DurationProjection) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DurationProjection) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DurationProjection) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *DurationProjection) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DurationProjection) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DurationProjection) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DurationProjection) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


