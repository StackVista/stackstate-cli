# SettingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]Setting**](Setting.md) |  | [optional] 

## Methods

### NewSettingList

`func NewSettingList() *SettingList`

NewSettingList instantiates a new SettingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingListWithDefaults

`func NewSettingListWithDefaults() *SettingList`

NewSettingListWithDefaults instantiates a new SettingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *SettingList) GetNodes() []Setting`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SettingList) GetNodesOk() (*[]Setting, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SettingList) SetNodes(v []Setting)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *SettingList) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


