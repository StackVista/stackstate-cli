# ComponentEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowEvents** | **bool** |  | 
**RelatedResourcesQuery** | Pointer to **string** | Topology query for resources related to the highlighted component | [optional] 

## Methods

### NewComponentEvents

`func NewComponentEvents(showEvents bool, ) *ComponentEvents`

NewComponentEvents instantiates a new ComponentEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentEventsWithDefaults

`func NewComponentEventsWithDefaults() *ComponentEvents`

NewComponentEventsWithDefaults instantiates a new ComponentEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowEvents

`func (o *ComponentEvents) GetShowEvents() bool`

GetShowEvents returns the ShowEvents field if non-nil, zero value otherwise.

### GetShowEventsOk

`func (o *ComponentEvents) GetShowEventsOk() (*bool, bool)`

GetShowEventsOk returns a tuple with the ShowEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEvents

`func (o *ComponentEvents) SetShowEvents(v bool)`

SetShowEvents sets ShowEvents field to given value.


### GetRelatedResourcesQuery

`func (o *ComponentEvents) GetRelatedResourcesQuery() string`

GetRelatedResourcesQuery returns the RelatedResourcesQuery field if non-nil, zero value otherwise.

### GetRelatedResourcesQueryOk

`func (o *ComponentEvents) GetRelatedResourcesQueryOk() (*string, bool)`

GetRelatedResourcesQueryOk returns a tuple with the RelatedResourcesQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourcesQuery

`func (o *ComponentEvents) SetRelatedResourcesQuery(v string)`

SetRelatedResourcesQuery sets RelatedResourcesQuery field to given value.

### HasRelatedResourcesQuery

`func (o *ComponentEvents) HasRelatedResourcesQuery() bool`

HasRelatedResourcesQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


