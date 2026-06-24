# PresentationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to [**PresentationOverview**](PresentationOverview.md) |  | [optional] 
**Highlight** | Pointer to [**PresentationHighlight**](PresentationHighlight.md) |  | [optional] 
**Summary** | Pointer to [**PresentationSummary**](PresentationSummary.md) |  | [optional] 
**Filters** | Pointer to [**[]ComponentPresentationFilter**](ComponentPresentationFilter.md) |  | [optional] 
**MetricPerspective** | Pointer to [**PresentationMetricPerspective**](PresentationMetricPerspective.md) |  | [optional] 
**Topology** | Pointer to [**TopologySettings**](TopologySettings.md) |  | [optional] 

## Methods

### NewPresentationDefinition

`func NewPresentationDefinition() *PresentationDefinition`

NewPresentationDefinition instantiates a new PresentationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationDefinitionWithDefaults

`func NewPresentationDefinitionWithDefaults() *PresentationDefinition`

NewPresentationDefinitionWithDefaults instantiates a new PresentationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *PresentationDefinition) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PresentationDefinition) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PresentationDefinition) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PresentationDefinition) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetOverview

`func (o *PresentationDefinition) GetOverview() PresentationOverview`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *PresentationDefinition) GetOverviewOk() (*PresentationOverview, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *PresentationDefinition) SetOverview(v PresentationOverview)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *PresentationDefinition) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetHighlight

`func (o *PresentationDefinition) GetHighlight() PresentationHighlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *PresentationDefinition) GetHighlightOk() (*PresentationHighlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *PresentationDefinition) SetHighlight(v PresentationHighlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *PresentationDefinition) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetSummary

`func (o *PresentationDefinition) GetSummary() PresentationSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PresentationDefinition) GetSummaryOk() (*PresentationSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PresentationDefinition) SetSummary(v PresentationSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PresentationDefinition) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetFilters

`func (o *PresentationDefinition) GetFilters() []ComponentPresentationFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PresentationDefinition) GetFiltersOk() (*[]ComponentPresentationFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PresentationDefinition) SetFilters(v []ComponentPresentationFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PresentationDefinition) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetricPerspective

`func (o *PresentationDefinition) GetMetricPerspective() PresentationMetricPerspective`

GetMetricPerspective returns the MetricPerspective field if non-nil, zero value otherwise.

### GetMetricPerspectiveOk

`func (o *PresentationDefinition) GetMetricPerspectiveOk() (*PresentationMetricPerspective, bool)`

GetMetricPerspectiveOk returns a tuple with the MetricPerspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPerspective

`func (o *PresentationDefinition) SetMetricPerspective(v PresentationMetricPerspective)`

SetMetricPerspective sets MetricPerspective field to given value.

### HasMetricPerspective

`func (o *PresentationDefinition) HasMetricPerspective() bool`

HasMetricPerspective returns a boolean if a field has been set.

### GetTopology

`func (o *PresentationDefinition) GetTopology() TopologySettings`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *PresentationDefinition) GetTopologyOk() (*TopologySettings, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *PresentationDefinition) SetTopology(v TopologySettings)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *PresentationDefinition) HasTopology() bool`

HasTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


