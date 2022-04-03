# Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]SettingParameter**](SettingParameter.md) |  | [optional] 
**Script** | Pointer to [**SettingScript**](SettingScript.md) |  | [optional] 

## Methods

### NewSetting

`func NewSetting() *Setting`

NewSetting instantiates a new Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingWithDefaults

`func NewSettingWithDefaults() *Setting`

NewSettingWithDefaults instantiates a new Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Setting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Setting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Setting) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Setting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Setting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Setting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Setting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Setting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *Setting) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Setting) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Setting) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Setting) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *Setting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Setting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Setting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Setting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnedBy

`func (o *Setting) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Setting) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Setting) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *Setting) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetParameters

`func (o *Setting) GetParameters() []SettingParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Setting) GetParametersOk() (*[]SettingParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Setting) SetParameters(v []SettingParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Setting) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetScript

`func (o *Setting) GetScript() SettingScript`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Setting) GetScriptOk() (*SettingScript, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Setting) SetScript(v SettingScript)`

SetScript sets Script field to given value.

### HasScript

`func (o *Setting) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


