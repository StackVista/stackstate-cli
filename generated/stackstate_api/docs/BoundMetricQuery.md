# BoundMetricQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** |  | 
**Alias** | **string** |  | 
**ComponentIdentifierTemplate** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoundMetricQuery

`func NewBoundMetricQuery(expression string, alias string, ) *BoundMetricQuery`

NewBoundMetricQuery instantiates a new BoundMetricQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundMetricQueryWithDefaults

`func NewBoundMetricQueryWithDefaults() *BoundMetricQuery`

NewBoundMetricQueryWithDefaults instantiates a new BoundMetricQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *BoundMetricQuery) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BoundMetricQuery) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BoundMetricQuery) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetAlias

`func (o *BoundMetricQuery) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BoundMetricQuery) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BoundMetricQuery) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetComponentIdentifierTemplate

`func (o *BoundMetricQuery) GetComponentIdentifierTemplate() string`

GetComponentIdentifierTemplate returns the ComponentIdentifierTemplate field if non-nil, zero value otherwise.

### GetComponentIdentifierTemplateOk

`func (o *BoundMetricQuery) GetComponentIdentifierTemplateOk() (*string, bool)`

GetComponentIdentifierTemplateOk returns a tuple with the ComponentIdentifierTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifierTemplate

`func (o *BoundMetricQuery) SetComponentIdentifierTemplate(v string)`

SetComponentIdentifierTemplate sets ComponentIdentifierTemplate field to given value.

### HasComponentIdentifierTemplate

`func (o *BoundMetricQuery) HasComponentIdentifierTemplate() bool`

HasComponentIdentifierTemplate returns a boolean if a field has been set.

### GetPrimary

`func (o *BoundMetricQuery) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *BoundMetricQuery) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *BoundMetricQuery) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *BoundMetricQuery) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


