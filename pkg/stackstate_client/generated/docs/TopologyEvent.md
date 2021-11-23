# TopologyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**SourceIdentifier** | Pointer to **string** |  | [optional] 
**ElementIdentifiers** | **[]string** |  | 
**Elements** | [**[]EventElement**](EventElement.md) |  | 
**CausingEvents** | [**[]EventRef**](EventRef.md) |  | 
**Source** | **string** |  | 
**Category** | [**EventCategory**](EventCategory.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**SourceLinks** | [**[]SourceLink**](SourceLink.md) |  | 
**Data** | **map[string]interface{}** |  | 
**EventType** | **string** |  | 
**EventTime** | **int64** |  | 
**ProcessedTime** | **int64** |  | 
**Tags** | [**[]EventTag**](EventTag.md) |  | 

## Methods

### NewTopologyEvent

`func NewTopologyEvent(identifier string, elementIdentifiers []string, elements []EventElement, causingEvents []EventRef, source string, category EventCategory, name string, sourceLinks []SourceLink, data map[string]interface{}, eventType string, eventTime int64, processedTime int64, tags []EventTag, ) *TopologyEvent`

NewTopologyEvent instantiates a new TopologyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyEventWithDefaults

`func NewTopologyEventWithDefaults() *TopologyEvent`

NewTopologyEventWithDefaults instantiates a new TopologyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *TopologyEvent) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TopologyEvent) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TopologyEvent) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSourceIdentifier

`func (o *TopologyEvent) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *TopologyEvent) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *TopologyEvent) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *TopologyEvent) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetElementIdentifiers

`func (o *TopologyEvent) GetElementIdentifiers() []string`

GetElementIdentifiers returns the ElementIdentifiers field if non-nil, zero value otherwise.

### GetElementIdentifiersOk

`func (o *TopologyEvent) GetElementIdentifiersOk() (*[]string, bool)`

GetElementIdentifiersOk returns a tuple with the ElementIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIdentifiers

`func (o *TopologyEvent) SetElementIdentifiers(v []string)`

SetElementIdentifiers sets ElementIdentifiers field to given value.


### GetElements

`func (o *TopologyEvent) GetElements() []EventElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *TopologyEvent) GetElementsOk() (*[]EventElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *TopologyEvent) SetElements(v []EventElement)`

SetElements sets Elements field to given value.


### GetCausingEvents

`func (o *TopologyEvent) GetCausingEvents() []EventRef`

GetCausingEvents returns the CausingEvents field if non-nil, zero value otherwise.

### GetCausingEventsOk

`func (o *TopologyEvent) GetCausingEventsOk() (*[]EventRef, bool)`

GetCausingEventsOk returns a tuple with the CausingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausingEvents

`func (o *TopologyEvent) SetCausingEvents(v []EventRef)`

SetCausingEvents sets CausingEvents field to given value.


### GetSource

`func (o *TopologyEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopologyEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopologyEvent) SetSource(v string)`

SetSource sets Source field to given value.


### GetCategory

`func (o *TopologyEvent) GetCategory() EventCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TopologyEvent) GetCategoryOk() (*EventCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TopologyEvent) SetCategory(v EventCategory)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *TopologyEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TopologyEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TopologyEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TopologyEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TopologyEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyEvent) SetName(v string)`

SetName sets Name field to given value.


### GetSourceLinks

`func (o *TopologyEvent) GetSourceLinks() []SourceLink`

GetSourceLinks returns the SourceLinks field if non-nil, zero value otherwise.

### GetSourceLinksOk

`func (o *TopologyEvent) GetSourceLinksOk() (*[]SourceLink, bool)`

GetSourceLinksOk returns a tuple with the SourceLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLinks

`func (o *TopologyEvent) SetSourceLinks(v []SourceLink)`

SetSourceLinks sets SourceLinks field to given value.


### GetData

`func (o *TopologyEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TopologyEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TopologyEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetEventType

`func (o *TopologyEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TopologyEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TopologyEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventTime

`func (o *TopologyEvent) GetEventTime() int64`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *TopologyEvent) GetEventTimeOk() (*int64, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *TopologyEvent) SetEventTime(v int64)`

SetEventTime sets EventTime field to given value.


### GetProcessedTime

`func (o *TopologyEvent) GetProcessedTime() int64`

GetProcessedTime returns the ProcessedTime field if non-nil, zero value otherwise.

### GetProcessedTimeOk

`func (o *TopologyEvent) GetProcessedTimeOk() (*int64, bool)`

GetProcessedTimeOk returns a tuple with the ProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTime

`func (o *TopologyEvent) SetProcessedTime(v int64)`

SetProcessedTime sets ProcessedTime field to given value.


### GetTags

`func (o *TopologyEvent) GetTags() []EventTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TopologyEvent) GetTagsOk() (*[]EventTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TopologyEvent) SetTags(v []EventTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


