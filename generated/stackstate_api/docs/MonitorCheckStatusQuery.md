# MonitorCheckStatusQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**ComponentIdentifierTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitorCheckStatusQuery

`func NewMonitorCheckStatusQuery(query string, ) *MonitorCheckStatusQuery`

NewMonitorCheckStatusQuery instantiates a new MonitorCheckStatusQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusQueryWithDefaults

`func NewMonitorCheckStatusQueryWithDefaults() *MonitorCheckStatusQuery`

NewMonitorCheckStatusQueryWithDefaults instantiates a new MonitorCheckStatusQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *MonitorCheckStatusQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorCheckStatusQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorCheckStatusQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetAlias

`func (o *MonitorCheckStatusQuery) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *MonitorCheckStatusQuery) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *MonitorCheckStatusQuery) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *MonitorCheckStatusQuery) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetComponentIdentifierTemplate

`func (o *MonitorCheckStatusQuery) GetComponentIdentifierTemplate() string`

GetComponentIdentifierTemplate returns the ComponentIdentifierTemplate field if non-nil, zero value otherwise.

### GetComponentIdentifierTemplateOk

`func (o *MonitorCheckStatusQuery) GetComponentIdentifierTemplateOk() (*string, bool)`

GetComponentIdentifierTemplateOk returns a tuple with the ComponentIdentifierTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifierTemplate

`func (o *MonitorCheckStatusQuery) SetComponentIdentifierTemplate(v string)`

SetComponentIdentifierTemplate sets ComponentIdentifierTemplate field to given value.

### HasComponentIdentifierTemplate

`func (o *MonitorCheckStatusQuery) HasComponentIdentifierTemplate() bool`

HasComponentIdentifierTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


