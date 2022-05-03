# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | **string** |  | 
**Id** | **int64** |  | 
**LastUpdateTimestamp** | **int64** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Manual** | Pointer to **bool** |  | [optional] 
**IsSettingsNode** | Pointer to **bool** |  | [optional] 

## Methods

### NewNode

`func NewNode(typeName string, id int64, lastUpdateTimestamp int64, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *Node) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Node) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Node) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetId

`func (o *Node) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v int64)`

SetId sets Id field to given value.


### GetLastUpdateTimestamp

`func (o *Node) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Node) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Node) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.


### GetIdentifier

`func (o *Node) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Node) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Node) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Node) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Node) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Node) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Node) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Node) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnedBy

`func (o *Node) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Node) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Node) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *Node) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetManual

`func (o *Node) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *Node) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *Node) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *Node) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetIsSettingsNode

`func (o *Node) GetIsSettingsNode() bool`

GetIsSettingsNode returns the IsSettingsNode field if non-nil, zero value otherwise.

### GetIsSettingsNodeOk

`func (o *Node) GetIsSettingsNodeOk() (*bool, bool)`

GetIsSettingsNodeOk returns a tuple with the IsSettingsNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSettingsNode

`func (o *Node) SetIsSettingsNode(v bool)`

SetIsSettingsNode sets IsSettingsNode field to given value.

### HasIsSettingsNode

`func (o *Node) HasIsSettingsNode() bool`

HasIsSettingsNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


