# OtelMappingStatusItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**RelationCount** | **int64** |  | 
**ComponentCount** | **int64** |  | 

## Methods

### NewOtelMappingStatusItem

`func NewOtelMappingStatusItem(name string, relationCount int64, componentCount int64, ) *OtelMappingStatusItem`

NewOtelMappingStatusItem instantiates a new OtelMappingStatusItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelMappingStatusItemWithDefaults

`func NewOtelMappingStatusItemWithDefaults() *OtelMappingStatusItem`

NewOtelMappingStatusItemWithDefaults instantiates a new OtelMappingStatusItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OtelMappingStatusItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelMappingStatusItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelMappingStatusItem) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *OtelMappingStatusItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OtelMappingStatusItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OtelMappingStatusItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *OtelMappingStatusItem) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRelationCount

`func (o *OtelMappingStatusItem) GetRelationCount() int64`

GetRelationCount returns the RelationCount field if non-nil, zero value otherwise.

### GetRelationCountOk

`func (o *OtelMappingStatusItem) GetRelationCountOk() (*int64, bool)`

GetRelationCountOk returns a tuple with the RelationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationCount

`func (o *OtelMappingStatusItem) SetRelationCount(v int64)`

SetRelationCount sets RelationCount field to given value.


### GetComponentCount

`func (o *OtelMappingStatusItem) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *OtelMappingStatusItem) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *OtelMappingStatusItem) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


