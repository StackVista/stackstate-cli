# ComponentTypeHighlights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamePlural** | **string** |  | 
**Fields** | [**[]ComponentTypeField**](ComponentTypeField.md) |  | 
**About** | [**ComponentTypeAbout**](ComponentTypeAbout.md) |  | 
**Events** | [**ComponentTypeEvents**](ComponentTypeEvents.md) |  | 
**ShowLogs** | **bool** |  | 
**ShowLastChange** | **bool** |  | 
**ExternalComponent** | [**ComponentTypeExternalComponent**](ComponentTypeExternalComponent.md) |  | 
**RelatedResources** | [**[]ComponentTypeRelatedResources**](ComponentTypeRelatedResources.md) |  | 
**Metrics** | [**[]ComponentHighlightMetrics**](ComponentHighlightMetrics.md) |  | 

## Methods

### NewComponentTypeHighlights

`func NewComponentTypeHighlights(namePlural string, fields []ComponentTypeField, about ComponentTypeAbout, events ComponentTypeEvents, showLogs bool, showLastChange bool, externalComponent ComponentTypeExternalComponent, relatedResources []ComponentTypeRelatedResources, metrics []ComponentHighlightMetrics, ) *ComponentTypeHighlights`

NewComponentTypeHighlights instantiates a new ComponentTypeHighlights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTypeHighlightsWithDefaults

`func NewComponentTypeHighlightsWithDefaults() *ComponentTypeHighlights`

NewComponentTypeHighlightsWithDefaults instantiates a new ComponentTypeHighlights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamePlural

`func (o *ComponentTypeHighlights) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *ComponentTypeHighlights) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *ComponentTypeHighlights) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.


### GetFields

`func (o *ComponentTypeHighlights) GetFields() []ComponentTypeField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ComponentTypeHighlights) GetFieldsOk() (*[]ComponentTypeField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ComponentTypeHighlights) SetFields(v []ComponentTypeField)`

SetFields sets Fields field to given value.


### GetAbout

`func (o *ComponentTypeHighlights) GetAbout() ComponentTypeAbout`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ComponentTypeHighlights) GetAboutOk() (*ComponentTypeAbout, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ComponentTypeHighlights) SetAbout(v ComponentTypeAbout)`

SetAbout sets About field to given value.


### GetEvents

`func (o *ComponentTypeHighlights) GetEvents() ComponentTypeEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ComponentTypeHighlights) GetEventsOk() (*ComponentTypeEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ComponentTypeHighlights) SetEvents(v ComponentTypeEvents)`

SetEvents sets Events field to given value.


### GetShowLogs

`func (o *ComponentTypeHighlights) GetShowLogs() bool`

GetShowLogs returns the ShowLogs field if non-nil, zero value otherwise.

### GetShowLogsOk

`func (o *ComponentTypeHighlights) GetShowLogsOk() (*bool, bool)`

GetShowLogsOk returns a tuple with the ShowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogs

`func (o *ComponentTypeHighlights) SetShowLogs(v bool)`

SetShowLogs sets ShowLogs field to given value.


### GetShowLastChange

`func (o *ComponentTypeHighlights) GetShowLastChange() bool`

GetShowLastChange returns the ShowLastChange field if non-nil, zero value otherwise.

### GetShowLastChangeOk

`func (o *ComponentTypeHighlights) GetShowLastChangeOk() (*bool, bool)`

GetShowLastChangeOk returns a tuple with the ShowLastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastChange

`func (o *ComponentTypeHighlights) SetShowLastChange(v bool)`

SetShowLastChange sets ShowLastChange field to given value.


### GetExternalComponent

`func (o *ComponentTypeHighlights) GetExternalComponent() ComponentTypeExternalComponent`

GetExternalComponent returns the ExternalComponent field if non-nil, zero value otherwise.

### GetExternalComponentOk

`func (o *ComponentTypeHighlights) GetExternalComponentOk() (*ComponentTypeExternalComponent, bool)`

GetExternalComponentOk returns a tuple with the ExternalComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalComponent

`func (o *ComponentTypeHighlights) SetExternalComponent(v ComponentTypeExternalComponent)`

SetExternalComponent sets ExternalComponent field to given value.


### GetRelatedResources

`func (o *ComponentTypeHighlights) GetRelatedResources() []ComponentTypeRelatedResources`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *ComponentTypeHighlights) GetRelatedResourcesOk() (*[]ComponentTypeRelatedResources, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *ComponentTypeHighlights) SetRelatedResources(v []ComponentTypeRelatedResources)`

SetRelatedResources sets RelatedResources field to given value.


### GetMetrics

`func (o *ComponentTypeHighlights) GetMetrics() []ComponentHighlightMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ComponentTypeHighlights) GetMetricsOk() (*[]ComponentHighlightMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ComponentTypeHighlights) SetMetrics(v []ComponentHighlightMetrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


