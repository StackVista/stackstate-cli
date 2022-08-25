# GetCausingEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Items** | [**[]TopologyEvent**](TopologyEvent.md) |  | 

## Methods

### NewGetCausingEventsResult

`func NewGetCausingEventsResult(type_ string, items []TopologyEvent, ) *GetCausingEventsResult`

NewGetCausingEventsResult instantiates a new GetCausingEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCausingEventsResultWithDefaults

`func NewGetCausingEventsResultWithDefaults() *GetCausingEventsResult`

NewGetCausingEventsResultWithDefaults instantiates a new GetCausingEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCausingEventsResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCausingEventsResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCausingEventsResult) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *GetCausingEventsResult) GetItems() []TopologyEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetCausingEventsResult) GetItemsOk() (*[]TopologyEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetCausingEventsResult) SetItems(v []TopologyEvent)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


