# ComponentLinkProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | Cel expression that returns a string that represents the name of the component to link to | 
**Identifier** | **string** | Cel expression that returns a string that represents the identifier in order to build the link | 

## Methods

### NewComponentLinkProjection

`func NewComponentLinkProjection(type_ string, name string, identifier string, ) *ComponentLinkProjection`

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


### GetIdentifier

`func (o *ComponentLinkProjection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ComponentLinkProjection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ComponentLinkProjection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


