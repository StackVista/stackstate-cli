# DurationHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]DurationHistogramBucket**](DurationHistogramBucket.md) | List of duration buckets | 
**Quantiles** | [**DurationQuantiles**](DurationQuantiles.md) |  | 

## Methods

### NewDurationHistogram

`func NewDurationHistogram(buckets []DurationHistogramBucket, quantiles DurationQuantiles, ) *DurationHistogram`

NewDurationHistogram instantiates a new DurationHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationHistogramWithDefaults

`func NewDurationHistogramWithDefaults() *DurationHistogram`

NewDurationHistogramWithDefaults instantiates a new DurationHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *DurationHistogram) GetBuckets() []DurationHistogramBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *DurationHistogram) GetBucketsOk() (*[]DurationHistogramBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *DurationHistogram) SetBuckets(v []DurationHistogramBucket)`

SetBuckets sets Buckets field to given value.


### GetQuantiles

`func (o *DurationHistogram) GetQuantiles() DurationQuantiles`

GetQuantiles returns the Quantiles field if non-nil, zero value otherwise.

### GetQuantilesOk

`func (o *DurationHistogram) GetQuantilesOk() (*DurationQuantiles, bool)`

GetQuantilesOk returns a tuple with the Quantiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantiles

`func (o *DurationHistogram) SetQuantiles(v DurationQuantiles)`

SetQuantiles sets Quantiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


