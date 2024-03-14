# MonitorIdentifierSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | [**[]MonitorIdentifierSuggestion**](MonitorIdentifierSuggestion.md) |  | 
**ComponentsInspected** | **int64** |  | 
**TimeseriesInspected** | **int64** |  | 

## Methods

### NewMonitorIdentifierSuggestions

`func NewMonitorIdentifierSuggestions(suggestions []MonitorIdentifierSuggestion, componentsInspected int64, timeseriesInspected int64, ) *MonitorIdentifierSuggestions`

NewMonitorIdentifierSuggestions instantiates a new MonitorIdentifierSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorIdentifierSuggestionsWithDefaults

`func NewMonitorIdentifierSuggestionsWithDefaults() *MonitorIdentifierSuggestions`

NewMonitorIdentifierSuggestionsWithDefaults instantiates a new MonitorIdentifierSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *MonitorIdentifierSuggestions) GetSuggestions() []MonitorIdentifierSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *MonitorIdentifierSuggestions) GetSuggestionsOk() (*[]MonitorIdentifierSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *MonitorIdentifierSuggestions) SetSuggestions(v []MonitorIdentifierSuggestion)`

SetSuggestions sets Suggestions field to given value.


### GetComponentsInspected

`func (o *MonitorIdentifierSuggestions) GetComponentsInspected() int64`

GetComponentsInspected returns the ComponentsInspected field if non-nil, zero value otherwise.

### GetComponentsInspectedOk

`func (o *MonitorIdentifierSuggestions) GetComponentsInspectedOk() (*int64, bool)`

GetComponentsInspectedOk returns a tuple with the ComponentsInspected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsInspected

`func (o *MonitorIdentifierSuggestions) SetComponentsInspected(v int64)`

SetComponentsInspected sets ComponentsInspected field to given value.


### GetTimeseriesInspected

`func (o *MonitorIdentifierSuggestions) GetTimeseriesInspected() int64`

GetTimeseriesInspected returns the TimeseriesInspected field if non-nil, zero value otherwise.

### GetTimeseriesInspectedOk

`func (o *MonitorIdentifierSuggestions) GetTimeseriesInspectedOk() (*int64, bool)`

GetTimeseriesInspectedOk returns a tuple with the TimeseriesInspected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesInspected

`func (o *MonitorIdentifierSuggestions) SetTimeseriesInspected(v int64)`

SetTimeseriesInspected sets TimeseriesInspected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


