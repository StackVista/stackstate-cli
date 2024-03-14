# EventBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total logs record count in the bucket. | 
**StartTime** | **int32** | The bucket initial timestamp. | 
**EndTime** | **int32** | The bucket final timestamp. | 
**CategoryBuckets** | [**[]EventCategoryBucket**](EventCategoryBucket.md) |  | 

## Methods

### NewEventBucket

`func NewEventBucket(count int64, startTime int32, endTime int32, categoryBuckets []EventCategoryBucket, ) *EventBucket`

NewEventBucket instantiates a new EventBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBucketWithDefaults

`func NewEventBucketWithDefaults() *EventBucket`

NewEventBucketWithDefaults instantiates a new EventBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *EventBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetStartTime

`func (o *EventBucket) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EventBucket) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EventBucket) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *EventBucket) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EventBucket) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EventBucket) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetCategoryBuckets

`func (o *EventBucket) GetCategoryBuckets() []EventCategoryBucket`

GetCategoryBuckets returns the CategoryBuckets field if non-nil, zero value otherwise.

### GetCategoryBucketsOk

`func (o *EventBucket) GetCategoryBucketsOk() (*[]EventCategoryBucket, bool)`

GetCategoryBucketsOk returns a tuple with the CategoryBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryBuckets

`func (o *EventBucket) SetCategoryBuckets(v []EventCategoryBucket)`

SetCategoryBuckets sets CategoryBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


