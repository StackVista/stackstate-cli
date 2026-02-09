# ExternalComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** |  | [optional] 
**Identifiers** | **[]string** |  | 
**Data** | **map[string]interface{}** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**ElementTypeTag** | Pointer to **string** |  | [optional] 
**SyncName** | **string** |  | 
**SourceProperties** | **map[string]interface{}** |  | 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExternalComponent

`func NewExternalComponent(identifiers []string, data map[string]interface{}, syncName string, sourceProperties map[string]interface{}, ) *ExternalComponent`

NewExternalComponent instantiates a new ExternalComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalComponentWithDefaults

`func NewExternalComponentWithDefaults() *ExternalComponent`

NewExternalComponentWithDefaults instantiates a new ExternalComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ExternalComponent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExternalComponent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExternalComponent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ExternalComponent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdentifiers

`func (o *ExternalComponent) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ExternalComponent) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ExternalComponent) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetData

`func (o *ExternalComponent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalComponent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalComponent) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetId

`func (o *ExternalComponent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalComponent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalComponent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ExternalComponent) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ExternalComponent) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ExternalComponent) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ExternalComponent) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetElementTypeTag

`func (o *ExternalComponent) GetElementTypeTag() string`

GetElementTypeTag returns the ElementTypeTag field if non-nil, zero value otherwise.

### GetElementTypeTagOk

`func (o *ExternalComponent) GetElementTypeTagOk() (*string, bool)`

GetElementTypeTagOk returns a tuple with the ElementTypeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementTypeTag

`func (o *ExternalComponent) SetElementTypeTag(v string)`

SetElementTypeTag sets ElementTypeTag field to given value.

### HasElementTypeTag

`func (o *ExternalComponent) HasElementTypeTag() bool`

HasElementTypeTag returns a boolean if a field has been set.

### GetSyncName

`func (o *ExternalComponent) GetSyncName() string`

GetSyncName returns the SyncName field if non-nil, zero value otherwise.

### GetSyncNameOk

`func (o *ExternalComponent) GetSyncNameOk() (*string, bool)`

GetSyncNameOk returns a tuple with the SyncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncName

`func (o *ExternalComponent) SetSyncName(v string)`

SetSyncName sets SyncName field to given value.


### GetSourceProperties

`func (o *ExternalComponent) GetSourceProperties() map[string]interface{}`

GetSourceProperties returns the SourceProperties field if non-nil, zero value otherwise.

### GetSourcePropertiesOk

`func (o *ExternalComponent) GetSourcePropertiesOk() (*map[string]interface{}, bool)`

GetSourcePropertiesOk returns a tuple with the SourceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProperties

`func (o *ExternalComponent) SetSourceProperties(v map[string]interface{})`

SetSourceProperties sets SourceProperties field to given value.


### GetStatus

`func (o *ExternalComponent) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalComponent) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalComponent) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalComponent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


