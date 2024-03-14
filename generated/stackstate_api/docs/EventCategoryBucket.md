# EventCategoryBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total events count for a particular category in the bucket. | 
**Category** | [**EventCategory**](EventCategory.md) |  | 

## Methods

### NewEventCategoryBucket

`func NewEventCategoryBucket(count int64, category EventCategory, ) *EventCategoryBucket`

NewEventCategoryBucket instantiates a new EventCategoryBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCategoryBucketWithDefaults

`func NewEventCategoryBucketWithDefaults() *EventCategoryBucket`

NewEventCategoryBucketWithDefaults instantiates a new EventCategoryBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *EventCategoryBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventCategoryBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventCategoryBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetCategory

`func (o *EventCategoryBucket) GetCategory() EventCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventCategoryBucket) GetCategoryOk() (*EventCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventCategoryBucket) SetCategory(v EventCategory)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


