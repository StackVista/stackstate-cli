# PromQLMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**Query** | **string** |  | 
**Unit** | **string** |  | 
**AliasTemplate** | **string** |  | 

## Methods

### NewPromQLMetric

`func NewPromQLMetric(type_ string, query string, unit string, aliasTemplate string, ) *PromQLMetric`

NewPromQLMetric instantiates a new PromQLMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromQLMetricWithDefaults

`func NewPromQLMetricWithDefaults() *PromQLMetric`

NewPromQLMetricWithDefaults instantiates a new PromQLMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PromQLMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromQLMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromQLMetric) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PromQLMetric) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromQLMetric) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromQLMetric) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PromQLMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuery

`func (o *PromQLMetric) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PromQLMetric) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PromQLMetric) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUnit

`func (o *PromQLMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PromQLMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PromQLMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAliasTemplate

`func (o *PromQLMetric) GetAliasTemplate() string`

GetAliasTemplate returns the AliasTemplate field if non-nil, zero value otherwise.

### GetAliasTemplateOk

`func (o *PromQLMetric) GetAliasTemplateOk() (*string, bool)`

GetAliasTemplateOk returns a tuple with the AliasTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTemplate

`func (o *PromQLMetric) SetAliasTemplate(v string)`

SetAliasTemplate sets AliasTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


