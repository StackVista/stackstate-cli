# StreamListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urn** | **string** |  | 
**ConsistencyModel** | **string** |  | 
**SubStreams** | **int32** |  | 

## Methods

### NewStreamListItem

`func NewStreamListItem(urn string, consistencyModel string, subStreams int32, ) *StreamListItem`

NewStreamListItem instantiates a new StreamListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamListItemWithDefaults

`func NewStreamListItemWithDefaults() *StreamListItem`

NewStreamListItemWithDefaults instantiates a new StreamListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrn

`func (o *StreamListItem) GetUrn() string`

GetUrn returns the Urn field if non-nil, zero value otherwise.

### GetUrnOk

`func (o *StreamListItem) GetUrnOk() (*string, bool)`

GetUrnOk returns a tuple with the Urn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrn

`func (o *StreamListItem) SetUrn(v string)`

SetUrn sets Urn field to given value.


### GetConsistencyModel

`func (o *StreamListItem) GetConsistencyModel() string`

GetConsistencyModel returns the ConsistencyModel field if non-nil, zero value otherwise.

### GetConsistencyModelOk

`func (o *StreamListItem) GetConsistencyModelOk() (*string, bool)`

GetConsistencyModelOk returns a tuple with the ConsistencyModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyModel

`func (o *StreamListItem) SetConsistencyModel(v string)`

SetConsistencyModel sets ConsistencyModel field to given value.


### GetSubStreams

`func (o *StreamListItem) GetSubStreams() int32`

GetSubStreams returns the SubStreams field if non-nil, zero value otherwise.

### GetSubStreamsOk

`func (o *StreamListItem) GetSubStreamsOk() (*int32, bool)`

GetSubStreamsOk returns a tuple with the SubStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreams

`func (o *StreamListItem) SetSubStreams(v int32)`

SetSubStreams sets SubStreams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


