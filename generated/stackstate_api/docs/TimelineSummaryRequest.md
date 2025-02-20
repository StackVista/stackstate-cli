# TimelineSummaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | [**TimelineSummaryRequestArguments**](TimelineSummaryRequestArguments.md) |  | 
**StartTime** | **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | 
**EndTime** | Pointer to **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | [optional] 
**HistogramBucketCount** | **int32** |  | 
**EventFilters** | Pointer to [**EventFilters**](EventFilters.md) |  | [optional] 

## Methods

### NewTimelineSummaryRequest

`func NewTimelineSummaryRequest(arguments TimelineSummaryRequestArguments, startTime int32, histogramBucketCount int32, ) *TimelineSummaryRequest`

NewTimelineSummaryRequest instantiates a new TimelineSummaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSummaryRequestWithDefaults

`func NewTimelineSummaryRequestWithDefaults() *TimelineSummaryRequest`

NewTimelineSummaryRequestWithDefaults instantiates a new TimelineSummaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *TimelineSummaryRequest) GetArguments() TimelineSummaryRequestArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *TimelineSummaryRequest) GetArgumentsOk() (*TimelineSummaryRequestArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *TimelineSummaryRequest) SetArguments(v TimelineSummaryRequestArguments)`

SetArguments sets Arguments field to given value.


### GetStartTime

`func (o *TimelineSummaryRequest) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimelineSummaryRequest) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimelineSummaryRequest) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *TimelineSummaryRequest) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimelineSummaryRequest) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimelineSummaryRequest) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimelineSummaryRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHistogramBucketCount

`func (o *TimelineSummaryRequest) GetHistogramBucketCount() int32`

GetHistogramBucketCount returns the HistogramBucketCount field if non-nil, zero value otherwise.

### GetHistogramBucketCountOk

`func (o *TimelineSummaryRequest) GetHistogramBucketCountOk() (*int32, bool)`

GetHistogramBucketCountOk returns a tuple with the HistogramBucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBucketCount

`func (o *TimelineSummaryRequest) SetHistogramBucketCount(v int32)`

SetHistogramBucketCount sets HistogramBucketCount field to given value.


### GetEventFilters

`func (o *TimelineSummaryRequest) GetEventFilters() EventFilters`

GetEventFilters returns the EventFilters field if non-nil, zero value otherwise.

### GetEventFiltersOk

`func (o *TimelineSummaryRequest) GetEventFiltersOk() (*EventFilters, bool)`

GetEventFiltersOk returns a tuple with the EventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilters

`func (o *TimelineSummaryRequest) SetEventFilters(v EventFilters)`

SetEventFilters sets EventFilters field to given value.

### HasEventFilters

`func (o *TimelineSummaryRequest) HasEventFilters() bool`

HasEventFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


