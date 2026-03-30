# LegacyComponentHighlights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamePlural** | **string** |  | 
**Events** | [**ComponentTypeEvents**](ComponentTypeEvents.md) |  | 
**ShowLogs** | **bool** |  | 
**ShowLastChange** | **bool** |  | 
**ExternalComponent** | [**ComponentTypeExternalComponent**](ComponentTypeExternalComponent.md) |  | 
**RelatedResources** | [**[]ComponentTypeRelatedResources**](ComponentTypeRelatedResources.md) |  | 
**Metrics** | [**[]ComponentHighlightMetrics**](ComponentHighlightMetrics.md) |  | 

## Methods

### NewLegacyComponentHighlights

`func NewLegacyComponentHighlights(namePlural string, events ComponentTypeEvents, showLogs bool, showLastChange bool, externalComponent ComponentTypeExternalComponent, relatedResources []ComponentTypeRelatedResources, metrics []ComponentHighlightMetrics, ) *LegacyComponentHighlights`

NewLegacyComponentHighlights instantiates a new LegacyComponentHighlights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyComponentHighlightsWithDefaults

`func NewLegacyComponentHighlightsWithDefaults() *LegacyComponentHighlights`

NewLegacyComponentHighlightsWithDefaults instantiates a new LegacyComponentHighlights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamePlural

`func (o *LegacyComponentHighlights) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *LegacyComponentHighlights) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *LegacyComponentHighlights) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.


### GetEvents

`func (o *LegacyComponentHighlights) GetEvents() ComponentTypeEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LegacyComponentHighlights) GetEventsOk() (*ComponentTypeEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LegacyComponentHighlights) SetEvents(v ComponentTypeEvents)`

SetEvents sets Events field to given value.


### GetShowLogs

`func (o *LegacyComponentHighlights) GetShowLogs() bool`

GetShowLogs returns the ShowLogs field if non-nil, zero value otherwise.

### GetShowLogsOk

`func (o *LegacyComponentHighlights) GetShowLogsOk() (*bool, bool)`

GetShowLogsOk returns a tuple with the ShowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogs

`func (o *LegacyComponentHighlights) SetShowLogs(v bool)`

SetShowLogs sets ShowLogs field to given value.


### GetShowLastChange

`func (o *LegacyComponentHighlights) GetShowLastChange() bool`

GetShowLastChange returns the ShowLastChange field if non-nil, zero value otherwise.

### GetShowLastChangeOk

`func (o *LegacyComponentHighlights) GetShowLastChangeOk() (*bool, bool)`

GetShowLastChangeOk returns a tuple with the ShowLastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastChange

`func (o *LegacyComponentHighlights) SetShowLastChange(v bool)`

SetShowLastChange sets ShowLastChange field to given value.


### GetExternalComponent

`func (o *LegacyComponentHighlights) GetExternalComponent() ComponentTypeExternalComponent`

GetExternalComponent returns the ExternalComponent field if non-nil, zero value otherwise.

### GetExternalComponentOk

`func (o *LegacyComponentHighlights) GetExternalComponentOk() (*ComponentTypeExternalComponent, bool)`

GetExternalComponentOk returns a tuple with the ExternalComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalComponent

`func (o *LegacyComponentHighlights) SetExternalComponent(v ComponentTypeExternalComponent)`

SetExternalComponent sets ExternalComponent field to given value.


### GetRelatedResources

`func (o *LegacyComponentHighlights) GetRelatedResources() []ComponentTypeRelatedResources`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *LegacyComponentHighlights) GetRelatedResourcesOk() (*[]ComponentTypeRelatedResources, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *LegacyComponentHighlights) SetRelatedResources(v []ComponentTypeRelatedResources)`

SetRelatedResources sets RelatedResources field to given value.


### GetMetrics

`func (o *LegacyComponentHighlights) GetMetrics() []ComponentHighlightMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *LegacyComponentHighlights) GetMetricsOk() (*[]ComponentHighlightMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *LegacyComponentHighlights) SetMetrics(v []ComponentHighlightMetrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


