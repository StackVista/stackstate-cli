# KubernetesLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Message** | **string** |  | 
**Highlights** | [**[]KubernetesLogHighlight**](KubernetesLogHighlight.md) |  | 
**PodName** | **string** |  | 
**PodUID** | **string** |  | 
**ContainerName** | **string** |  | 

## Methods

### NewKubernetesLogRecord

`func NewKubernetesLogRecord(timestamp time.Time, message string, highlights []KubernetesLogHighlight, podName string, podUID string, containerName string, ) *KubernetesLogRecord`

NewKubernetesLogRecord instantiates a new KubernetesLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLogRecordWithDefaults

`func NewKubernetesLogRecordWithDefaults() *KubernetesLogRecord`

NewKubernetesLogRecordWithDefaults instantiates a new KubernetesLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *KubernetesLogRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *KubernetesLogRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *KubernetesLogRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMessage

`func (o *KubernetesLogRecord) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubernetesLogRecord) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubernetesLogRecord) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHighlights

`func (o *KubernetesLogRecord) GetHighlights() []KubernetesLogHighlight`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *KubernetesLogRecord) GetHighlightsOk() (*[]KubernetesLogHighlight, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *KubernetesLogRecord) SetHighlights(v []KubernetesLogHighlight)`

SetHighlights sets Highlights field to given value.


### GetPodName

`func (o *KubernetesLogRecord) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *KubernetesLogRecord) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *KubernetesLogRecord) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetPodUID

`func (o *KubernetesLogRecord) GetPodUID() string`

GetPodUID returns the PodUID field if non-nil, zero value otherwise.

### GetPodUIDOk

`func (o *KubernetesLogRecord) GetPodUIDOk() (*string, bool)`

GetPodUIDOk returns a tuple with the PodUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUID

`func (o *KubernetesLogRecord) SetPodUID(v string)`

SetPodUID sets PodUID field to given value.


### GetContainerName

`func (o *KubernetesLogRecord) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *KubernetesLogRecord) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *KubernetesLogRecord) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


