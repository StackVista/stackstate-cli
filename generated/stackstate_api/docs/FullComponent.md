# FullComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | **string** |  | 
**Iconbase64** | Pointer to **string** |  | [optional] 
**Data** | [**ComponentData**](ComponentData.md) |  | 
**Highlights** | Pointer to [**ComponentTypeHighlights**](ComponentTypeHighlights.md) |  | [optional] 
**Actions** | [**[]ComponentAction**](ComponentAction.md) |  | 
**BoundMetrics** | [**[]BoundMetric**](BoundMetric.md) |  | 
**BoundTraces** | Pointer to [**BoundTraces**](BoundTraces.md) |  | [optional] 

## Methods

### NewFullComponent

`func NewFullComponent(typeName string, data ComponentData, actions []ComponentAction, boundMetrics []BoundMetric, ) *FullComponent`

NewFullComponent instantiates a new FullComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullComponentWithDefaults

`func NewFullComponentWithDefaults() *FullComponent`

NewFullComponentWithDefaults instantiates a new FullComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *FullComponent) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FullComponent) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FullComponent) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetIconbase64

`func (o *FullComponent) GetIconbase64() string`

GetIconbase64 returns the Iconbase64 field if non-nil, zero value otherwise.

### GetIconbase64Ok

`func (o *FullComponent) GetIconbase64Ok() (*string, bool)`

GetIconbase64Ok returns a tuple with the Iconbase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconbase64

`func (o *FullComponent) SetIconbase64(v string)`

SetIconbase64 sets Iconbase64 field to given value.

### HasIconbase64

`func (o *FullComponent) HasIconbase64() bool`

HasIconbase64 returns a boolean if a field has been set.

### GetData

`func (o *FullComponent) GetData() ComponentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FullComponent) GetDataOk() (*ComponentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FullComponent) SetData(v ComponentData)`

SetData sets Data field to given value.


### GetHighlights

`func (o *FullComponent) GetHighlights() ComponentTypeHighlights`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *FullComponent) GetHighlightsOk() (*ComponentTypeHighlights, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *FullComponent) SetHighlights(v ComponentTypeHighlights)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *FullComponent) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetActions

`func (o *FullComponent) GetActions() []ComponentAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FullComponent) GetActionsOk() (*[]ComponentAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FullComponent) SetActions(v []ComponentAction)`

SetActions sets Actions field to given value.


### GetBoundMetrics

`func (o *FullComponent) GetBoundMetrics() []BoundMetric`

GetBoundMetrics returns the BoundMetrics field if non-nil, zero value otherwise.

### GetBoundMetricsOk

`func (o *FullComponent) GetBoundMetricsOk() (*[]BoundMetric, bool)`

GetBoundMetricsOk returns a tuple with the BoundMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundMetrics

`func (o *FullComponent) SetBoundMetrics(v []BoundMetric)`

SetBoundMetrics sets BoundMetrics field to given value.


### GetBoundTraces

`func (o *FullComponent) GetBoundTraces() BoundTraces`

GetBoundTraces returns the BoundTraces field if non-nil, zero value otherwise.

### GetBoundTracesOk

`func (o *FullComponent) GetBoundTracesOk() (*BoundTraces, bool)`

GetBoundTracesOk returns a tuple with the BoundTraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTraces

`func (o *FullComponent) SetBoundTraces(v BoundTraces)`

SetBoundTraces sets BoundTraces field to given value.

### HasBoundTraces

`func (o *FullComponent) HasBoundTraces() bool`

HasBoundTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


