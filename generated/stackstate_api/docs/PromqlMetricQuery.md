# PromqlMetricQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**EndTime** | Pointer to **int32** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**IdentifierFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewPromqlMetricQuery

`func NewPromqlMetricQuery(query string, ) *PromqlMetricQuery`

NewPromqlMetricQuery instantiates a new PromqlMetricQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromqlMetricQueryWithDefaults

`func NewPromqlMetricQueryWithDefaults() *PromqlMetricQuery`

NewPromqlMetricQueryWithDefaults instantiates a new PromqlMetricQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PromqlMetricQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PromqlMetricQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PromqlMetricQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetAlias

`func (o *PromqlMetricQuery) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PromqlMetricQuery) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PromqlMetricQuery) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PromqlMetricQuery) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *PromqlMetricQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PromqlMetricQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PromqlMetricQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PromqlMetricQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *PromqlMetricQuery) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PromqlMetricQuery) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PromqlMetricQuery) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PromqlMetricQuery) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *PromqlMetricQuery) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *PromqlMetricQuery) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *PromqlMetricQuery) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *PromqlMetricQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStep

`func (o *PromqlMetricQuery) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PromqlMetricQuery) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PromqlMetricQuery) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *PromqlMetricQuery) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetUnit

`func (o *PromqlMetricQuery) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PromqlMetricQuery) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PromqlMetricQuery) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PromqlMetricQuery) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetIdentifierFormat

`func (o *PromqlMetricQuery) GetIdentifierFormat() string`

GetIdentifierFormat returns the IdentifierFormat field if non-nil, zero value otherwise.

### GetIdentifierFormatOk

`func (o *PromqlMetricQuery) GetIdentifierFormatOk() (*string, bool)`

GetIdentifierFormatOk returns a tuple with the IdentifierFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierFormat

`func (o *PromqlMetricQuery) SetIdentifierFormat(v string)`

SetIdentifierFormat sets IdentifierFormat field to given value.

### HasIdentifierFormat

`func (o *PromqlMetricQuery) HasIdentifierFormat() bool`

HasIdentifierFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


