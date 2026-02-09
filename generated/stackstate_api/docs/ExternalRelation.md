# ExternalRelation

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

## Methods

### NewExternalRelation

`func NewExternalRelation(identifiers []string, data map[string]interface{}, syncName string, ) *ExternalRelation`

NewExternalRelation instantiates a new ExternalRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRelationWithDefaults

`func NewExternalRelationWithDefaults() *ExternalRelation`

NewExternalRelationWithDefaults instantiates a new ExternalRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ExternalRelation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExternalRelation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExternalRelation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ExternalRelation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdentifiers

`func (o *ExternalRelation) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ExternalRelation) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ExternalRelation) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetData

`func (o *ExternalRelation) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalRelation) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalRelation) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetId

`func (o *ExternalRelation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalRelation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalRelation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalRelation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ExternalRelation) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ExternalRelation) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ExternalRelation) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ExternalRelation) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetElementTypeTag

`func (o *ExternalRelation) GetElementTypeTag() string`

GetElementTypeTag returns the ElementTypeTag field if non-nil, zero value otherwise.

### GetElementTypeTagOk

`func (o *ExternalRelation) GetElementTypeTagOk() (*string, bool)`

GetElementTypeTagOk returns a tuple with the ElementTypeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementTypeTag

`func (o *ExternalRelation) SetElementTypeTag(v string)`

SetElementTypeTag sets ElementTypeTag field to given value.

### HasElementTypeTag

`func (o *ExternalRelation) HasElementTypeTag() bool`

HasElementTypeTag returns a boolean if a field has been set.

### GetSyncName

`func (o *ExternalRelation) GetSyncName() string`

GetSyncName returns the SyncName field if non-nil, zero value otherwise.

### GetSyncNameOk

`func (o *ExternalRelation) GetSyncNameOk() (*string, bool)`

GetSyncNameOk returns a tuple with the SyncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncName

`func (o *ExternalRelation) SetSyncName(v string)`

SetSyncName sets SyncName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


