# Export

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodesWithIds** | Pointer to **[]int64** |  | [optional] 
**AllNodesOfTypes** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to **[]string** |  | [optional] 
**AllowReferences** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExport

`func NewExport() *Export`

NewExport instantiates a new Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportWithDefaults

`func NewExportWithDefaults() *Export`

NewExportWithDefaults instantiates a new Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodesWithIds

`func (o *Export) GetNodesWithIds() []int64`

GetNodesWithIds returns the NodesWithIds field if non-nil, zero value otherwise.

### GetNodesWithIdsOk

`func (o *Export) GetNodesWithIdsOk() (*[]int64, bool)`

GetNodesWithIdsOk returns a tuple with the NodesWithIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesWithIds

`func (o *Export) SetNodesWithIds(v []int64)`

SetNodesWithIds sets NodesWithIds field to given value.

### HasNodesWithIds

`func (o *Export) HasNodesWithIds() bool`

HasNodesWithIds returns a boolean if a field has been set.

### GetAllNodesOfTypes

`func (o *Export) GetAllNodesOfTypes() []string`

GetAllNodesOfTypes returns the AllNodesOfTypes field if non-nil, zero value otherwise.

### GetAllNodesOfTypesOk

`func (o *Export) GetAllNodesOfTypesOk() (*[]string, bool)`

GetAllNodesOfTypesOk returns a tuple with the AllNodesOfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNodesOfTypes

`func (o *Export) SetAllNodesOfTypes(v []string)`

SetAllNodesOfTypes sets AllNodesOfTypes field to given value.

### HasAllNodesOfTypes

`func (o *Export) HasAllNodesOfTypes() bool`

HasAllNodesOfTypes returns a boolean if a field has been set.

### GetNamespace

`func (o *Export) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Export) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Export) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Export) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetAllowReferences

`func (o *Export) GetAllowReferences() []string`

GetAllowReferences returns the AllowReferences field if non-nil, zero value otherwise.

### GetAllowReferencesOk

`func (o *Export) GetAllowReferencesOk() (*[]string, bool)`

GetAllowReferencesOk returns a tuple with the AllowReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReferences

`func (o *Export) SetAllowReferences(v []string)`

SetAllowReferences sets AllowReferences field to given value.

### HasAllowReferences

`func (o *Export) HasAllowReferences() bool`

HasAllowReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


