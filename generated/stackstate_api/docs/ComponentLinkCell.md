# ComponentLinkCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Component** | Pointer to [**ComponentLink**](ComponentLink.md) |  | [optional] 

## Methods

### NewComponentLinkCell

`func NewComponentLinkCell(type_ string, ) *ComponentLinkCell`

NewComponentLinkCell instantiates a new ComponentLinkCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentLinkCellWithDefaults

`func NewComponentLinkCellWithDefaults() *ComponentLinkCell`

NewComponentLinkCellWithDefaults instantiates a new ComponentLinkCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentLinkCell) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentLinkCell) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentLinkCell) SetType(v string)`

SetType sets Type field to given value.


### GetComponent

`func (o *ComponentLinkCell) GetComponent() ComponentLink`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentLinkCell) GetComponentOk() (*ComponentLink, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentLinkCell) SetComponent(v ComponentLink)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ComponentLinkCell) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


