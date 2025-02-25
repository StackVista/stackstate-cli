# TimelineSummaryEventBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Count** | **int64** |  | 
**StartTimeEpochMillis** | **int64** |  | 
**EndTimeEpochMillis** | Pointer to **int64** |  | [optional] 

## Methods

### NewTimelineSummaryEventBucket

`func NewTimelineSummaryEventBucket(type_ string, count int64, startTimeEpochMillis int64, ) *TimelineSummaryEventBucket`

NewTimelineSummaryEventBucket instantiates a new TimelineSummaryEventBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSummaryEventBucketWithDefaults

`func NewTimelineSummaryEventBucketWithDefaults() *TimelineSummaryEventBucket`

NewTimelineSummaryEventBucketWithDefaults instantiates a new TimelineSummaryEventBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimelineSummaryEventBucket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineSummaryEventBucket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineSummaryEventBucket) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *TimelineSummaryEventBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TimelineSummaryEventBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TimelineSummaryEventBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetStartTimeEpochMillis

`func (o *TimelineSummaryEventBucket) GetStartTimeEpochMillis() int64`

GetStartTimeEpochMillis returns the StartTimeEpochMillis field if non-nil, zero value otherwise.

### GetStartTimeEpochMillisOk

`func (o *TimelineSummaryEventBucket) GetStartTimeEpochMillisOk() (*int64, bool)`

GetStartTimeEpochMillisOk returns a tuple with the StartTimeEpochMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeEpochMillis

`func (o *TimelineSummaryEventBucket) SetStartTimeEpochMillis(v int64)`

SetStartTimeEpochMillis sets StartTimeEpochMillis field to given value.


### GetEndTimeEpochMillis

`func (o *TimelineSummaryEventBucket) GetEndTimeEpochMillis() int64`

GetEndTimeEpochMillis returns the EndTimeEpochMillis field if non-nil, zero value otherwise.

### GetEndTimeEpochMillisOk

`func (o *TimelineSummaryEventBucket) GetEndTimeEpochMillisOk() (*int64, bool)`

GetEndTimeEpochMillisOk returns a tuple with the EndTimeEpochMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeEpochMillis

`func (o *TimelineSummaryEventBucket) SetEndTimeEpochMillis(v int64)`

SetEndTimeEpochMillis sets EndTimeEpochMillis field to given value.

### HasEndTimeEpochMillis

`func (o *TimelineSummaryEventBucket) HasEndTimeEpochMillis() bool`

HasEndTimeEpochMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


