# BoundMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**BoundQueries** | [**[]BoundMetricQuery**](BoundMetricQuery.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**ChartType** | [**ChartType**](ChartType.md) |  | 
**Alias** | Pointer to **string** |  | [optional] 
**Valuation** | Pointer to [**MetricValuation**](MetricValuation.md) |  | [optional] 
**Tags** | **map[string]string** |  | 
**Layout** | Pointer to [**MetricBindingLayout**](MetricBindingLayout.md) |  | [optional] 
**Dummy** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoundMetric

`func NewBoundMetric(name string, boundQueries []BoundMetricQuery, chartType ChartType, tags map[string]string, ) *BoundMetric`

NewBoundMetric instantiates a new BoundMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundMetricWithDefaults

`func NewBoundMetricWithDefaults() *BoundMetric`

NewBoundMetricWithDefaults instantiates a new BoundMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BoundMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundMetric) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *BoundMetric) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BoundMetric) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BoundMetric) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BoundMetric) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetBoundQueries

`func (o *BoundMetric) GetBoundQueries() []BoundMetricQuery`

GetBoundQueries returns the BoundQueries field if non-nil, zero value otherwise.

### GetBoundQueriesOk

`func (o *BoundMetric) GetBoundQueriesOk() (*[]BoundMetricQuery, bool)`

GetBoundQueriesOk returns a tuple with the BoundQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundQueries

`func (o *BoundMetric) SetBoundQueries(v []BoundMetricQuery)`

SetBoundQueries sets BoundQueries field to given value.


### GetDescription

`func (o *BoundMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BoundMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BoundMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BoundMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *BoundMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BoundMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BoundMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BoundMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetChartType

`func (o *BoundMetric) GetChartType() ChartType`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *BoundMetric) GetChartTypeOk() (*ChartType, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *BoundMetric) SetChartType(v ChartType)`

SetChartType sets ChartType field to given value.


### GetAlias

`func (o *BoundMetric) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BoundMetric) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BoundMetric) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *BoundMetric) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetValuation

`func (o *BoundMetric) GetValuation() MetricValuation`

GetValuation returns the Valuation field if non-nil, zero value otherwise.

### GetValuationOk

`func (o *BoundMetric) GetValuationOk() (*MetricValuation, bool)`

GetValuationOk returns a tuple with the Valuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuation

`func (o *BoundMetric) SetValuation(v MetricValuation)`

SetValuation sets Valuation field to given value.

### HasValuation

`func (o *BoundMetric) HasValuation() bool`

HasValuation returns a boolean if a field has been set.

### GetTags

`func (o *BoundMetric) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BoundMetric) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BoundMetric) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetLayout

`func (o *BoundMetric) GetLayout() MetricBindingLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *BoundMetric) GetLayoutOk() (*MetricBindingLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *BoundMetric) SetLayout(v MetricBindingLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *BoundMetric) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDummy

`func (o *BoundMetric) GetDummy() bool`

GetDummy returns the Dummy field if non-nil, zero value otherwise.

### GetDummyOk

`func (o *BoundMetric) GetDummyOk() (*bool, bool)`

GetDummyOk returns a tuple with the Dummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummy

`func (o *BoundMetric) SetDummy(v bool)`

SetDummy sets Dummy field to given value.

### HasDummy

`func (o *BoundMetric) HasDummy() bool`

HasDummy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


