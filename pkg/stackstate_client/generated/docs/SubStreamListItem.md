# SubStreamListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubStreamId** | **string** |  | 
**SubStreamState** | [**HealthSubStreamConsistencyState**](HealthSubStreamConsistencyState.md) |  | 
**CheckStateCount** | **int32** |  | 

## Methods

### NewSubStreamListItem

`func NewSubStreamListItem(subStreamId string, subStreamState HealthSubStreamConsistencyState, checkStateCount int32, ) *SubStreamListItem`

NewSubStreamListItem instantiates a new SubStreamListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubStreamListItemWithDefaults

`func NewSubStreamListItemWithDefaults() *SubStreamListItem`

NewSubStreamListItemWithDefaults instantiates a new SubStreamListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubStreamId

`func (o *SubStreamListItem) GetSubStreamId() string`

GetSubStreamId returns the SubStreamId field if non-nil, zero value otherwise.

### GetSubStreamIdOk

`func (o *SubStreamListItem) GetSubStreamIdOk() (*string, bool)`

GetSubStreamIdOk returns a tuple with the SubStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreamId

`func (o *SubStreamListItem) SetSubStreamId(v string)`

SetSubStreamId sets SubStreamId field to given value.


### GetSubStreamState

`func (o *SubStreamListItem) GetSubStreamState() HealthSubStreamConsistencyState`

GetSubStreamState returns the SubStreamState field if non-nil, zero value otherwise.

### GetSubStreamStateOk

`func (o *SubStreamListItem) GetSubStreamStateOk() (*HealthSubStreamConsistencyState, bool)`

GetSubStreamStateOk returns a tuple with the SubStreamState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreamState

`func (o *SubStreamListItem) SetSubStreamState(v HealthSubStreamConsistencyState)`

SetSubStreamState sets SubStreamState field to given value.


### GetCheckStateCount

`func (o *SubStreamListItem) GetCheckStateCount() int32`

GetCheckStateCount returns the CheckStateCount field if non-nil, zero value otherwise.

### GetCheckStateCountOk

`func (o *SubStreamListItem) GetCheckStateCountOk() (*int32, bool)`

GetCheckStateCountOk returns a tuple with the CheckStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStateCount

`func (o *SubStreamListItem) SetCheckStateCount(v int32)`

SetCheckStateCount sets CheckStateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


