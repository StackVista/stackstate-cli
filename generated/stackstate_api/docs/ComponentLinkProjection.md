# ComponentLinkProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | Cel expression that returns a string that represents the name of the component to link to | 
**ComponentIdentifier** | **string** | Cel expression that returns a string that represents the componentIdentifier in order to build the link | 

## Methods

### NewComponentLinkProjection

`func NewComponentLinkProjection(type_ string, name string, componentIdentifier string, ) *ComponentLinkProjection`

NewComponentLinkProjection instantiates a new ComponentLinkProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentLinkProjectionWithDefaults

`func NewComponentLinkProjectionWithDefaults() *ComponentLinkProjection`

NewComponentLinkProjectionWithDefaults instantiates a new ComponentLinkProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentLinkProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentLinkProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentLinkProjection) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ComponentLinkProjection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentLinkProjection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentLinkProjection) SetName(v string)`

SetName sets Name field to given value.


### GetComponentIdentifier

`func (o *ComponentLinkProjection) GetComponentIdentifier() string`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ComponentLinkProjection) GetComponentIdentifierOk() (*string, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ComponentLinkProjection) SetComponentIdentifier(v string)`

SetComponentIdentifier sets ComponentIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


