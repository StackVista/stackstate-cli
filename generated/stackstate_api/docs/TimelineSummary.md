# TimelineSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]TimelineSummaryEventBucket**](TimelineSummaryEventBucket.md) |  | 
**HealthHistory** | [**[]TimelineSummaryHealthChange**](TimelineSummaryHealthChange.md) |  | 
**FromTime** | **int64** |  | 

## Methods

### NewTimelineSummary

`func NewTimelineSummary(buckets []TimelineSummaryEventBucket, healthHistory []TimelineSummaryHealthChange, fromTime int64, ) *TimelineSummary`

NewTimelineSummary instantiates a new TimelineSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSummaryWithDefaults

`func NewTimelineSummaryWithDefaults() *TimelineSummary`

NewTimelineSummaryWithDefaults instantiates a new TimelineSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *TimelineSummary) GetBuckets() []TimelineSummaryEventBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *TimelineSummary) GetBucketsOk() (*[]TimelineSummaryEventBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *TimelineSummary) SetBuckets(v []TimelineSummaryEventBucket)`

SetBuckets sets Buckets field to given value.


### GetHealthHistory

`func (o *TimelineSummary) GetHealthHistory() []TimelineSummaryHealthChange`

GetHealthHistory returns the HealthHistory field if non-nil, zero value otherwise.

### GetHealthHistoryOk

`func (o *TimelineSummary) GetHealthHistoryOk() (*[]TimelineSummaryHealthChange, bool)`

GetHealthHistoryOk returns a tuple with the HealthHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthHistory

`func (o *TimelineSummary) SetHealthHistory(v []TimelineSummaryHealthChange)`

SetHealthHistory sets HealthHistory field to given value.


### GetFromTime

`func (o *TimelineSummary) GetFromTime() int64`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *TimelineSummary) GetFromTimeOk() (*int64, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *TimelineSummary) SetFromTime(v int64)`

SetFromTime sets FromTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


