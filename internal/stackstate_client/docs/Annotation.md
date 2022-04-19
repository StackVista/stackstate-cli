# Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AnnotationType** | [**AnnotationType**](AnnotationType.md) |  | 
**CreatedTime** | **int64** |  | 
**Data** | Pointer to [**AnnotationData**](AnnotationData.md) |  | [optional] 
**Description** | **string** |  | 
**EventTimeInterval** | [**TimeRange**](TimeRange.md) |  | 
**Id** | **string** |  | 
**Identifiers** | **[]string** |  | 
**Name** | **string** |  | 
**ProcessedTime** | **int64** |  | 
**Reference** | [**Reference**](Reference.md) |  | 
**Tags** | **[]string** |  | 

## Methods

### NewAnnotation

`func NewAnnotation(type_ string, annotationType AnnotationType, createdTime int64, description string, eventTimeInterval TimeRange, id string, identifiers []string, name string, processedTime int64, reference Reference, tags []string, ) *Annotation`

NewAnnotation instantiates a new Annotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationWithDefaults

`func NewAnnotationWithDefaults() *Annotation`

NewAnnotationWithDefaults instantiates a new Annotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Annotation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Annotation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Annotation) SetType(v string)`

SetType sets Type field to given value.


### GetAnnotationType

`func (o *Annotation) GetAnnotationType() AnnotationType`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *Annotation) GetAnnotationTypeOk() (*AnnotationType, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *Annotation) SetAnnotationType(v AnnotationType)`

SetAnnotationType sets AnnotationType field to given value.


### GetCreatedTime

`func (o *Annotation) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Annotation) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Annotation) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.


### GetData

`func (o *Annotation) GetData() AnnotationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Annotation) GetDataOk() (*AnnotationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Annotation) SetData(v AnnotationData)`

SetData sets Data field to given value.

### HasData

`func (o *Annotation) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *Annotation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Annotation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Annotation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventTimeInterval

`func (o *Annotation) GetEventTimeInterval() TimeRange`

GetEventTimeInterval returns the EventTimeInterval field if non-nil, zero value otherwise.

### GetEventTimeIntervalOk

`func (o *Annotation) GetEventTimeIntervalOk() (*TimeRange, bool)`

GetEventTimeIntervalOk returns a tuple with the EventTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeInterval

`func (o *Annotation) SetEventTimeInterval(v TimeRange)`

SetEventTimeInterval sets EventTimeInterval field to given value.


### GetId

`func (o *Annotation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Annotation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Annotation) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifiers

`func (o *Annotation) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *Annotation) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *Annotation) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetName

`func (o *Annotation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Annotation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Annotation) SetName(v string)`

SetName sets Name field to given value.


### GetProcessedTime

`func (o *Annotation) GetProcessedTime() int64`

GetProcessedTime returns the ProcessedTime field if non-nil, zero value otherwise.

### GetProcessedTimeOk

`func (o *Annotation) GetProcessedTimeOk() (*int64, bool)`

GetProcessedTimeOk returns a tuple with the ProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTime

`func (o *Annotation) SetProcessedTime(v int64)`

SetProcessedTime sets ProcessedTime field to given value.


### GetReference

`func (o *Annotation) GetReference() Reference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Annotation) GetReferenceOk() (*Reference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Annotation) SetReference(v Reference)`

SetReference sets Reference field to given value.


### GetTags

`func (o *Annotation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Annotation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Annotation) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


