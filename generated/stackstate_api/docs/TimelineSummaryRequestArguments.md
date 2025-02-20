# TimelineSummaryRequestArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComponentIdentifier** | **string** |  | 
**QueryTime** | Pointer to **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | [optional] 
**Query** | **string** |  | 

## Methods

### NewTimelineSummaryRequestArguments

`func NewTimelineSummaryRequestArguments(type_ string, componentIdentifier string, query string, ) *TimelineSummaryRequestArguments`

NewTimelineSummaryRequestArguments instantiates a new TimelineSummaryRequestArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSummaryRequestArgumentsWithDefaults

`func NewTimelineSummaryRequestArgumentsWithDefaults() *TimelineSummaryRequestArguments`

NewTimelineSummaryRequestArgumentsWithDefaults instantiates a new TimelineSummaryRequestArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimelineSummaryRequestArguments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineSummaryRequestArguments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineSummaryRequestArguments) SetType(v string)`

SetType sets Type field to given value.


### GetComponentIdentifier

`func (o *TimelineSummaryRequestArguments) GetComponentIdentifier() string`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *TimelineSummaryRequestArguments) GetComponentIdentifierOk() (*string, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *TimelineSummaryRequestArguments) SetComponentIdentifier(v string)`

SetComponentIdentifier sets ComponentIdentifier field to given value.


### GetQueryTime

`func (o *TimelineSummaryRequestArguments) GetQueryTime() int32`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *TimelineSummaryRequestArguments) GetQueryTimeOk() (*int32, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *TimelineSummaryRequestArguments) SetQueryTime(v int32)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *TimelineSummaryRequestArguments) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetQuery

`func (o *TimelineSummaryRequestArguments) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TimelineSummaryRequestArguments) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TimelineSummaryRequestArguments) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


