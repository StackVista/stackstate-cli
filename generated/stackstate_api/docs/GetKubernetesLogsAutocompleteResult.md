# GetKubernetesLogsAutocompleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerNames** | **[]string** |  | 
**Levels** | [**[]LogSeverity**](LogSeverity.md) |  | 

## Methods

### NewGetKubernetesLogsAutocompleteResult

`func NewGetKubernetesLogsAutocompleteResult(containerNames []string, levels []LogSeverity, ) *GetKubernetesLogsAutocompleteResult`

NewGetKubernetesLogsAutocompleteResult instantiates a new GetKubernetesLogsAutocompleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesLogsAutocompleteResultWithDefaults

`func NewGetKubernetesLogsAutocompleteResultWithDefaults() *GetKubernetesLogsAutocompleteResult`

NewGetKubernetesLogsAutocompleteResultWithDefaults instantiates a new GetKubernetesLogsAutocompleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerNames

`func (o *GetKubernetesLogsAutocompleteResult) GetContainerNames() []string`

GetContainerNames returns the ContainerNames field if non-nil, zero value otherwise.

### GetContainerNamesOk

`func (o *GetKubernetesLogsAutocompleteResult) GetContainerNamesOk() (*[]string, bool)`

GetContainerNamesOk returns a tuple with the ContainerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNames

`func (o *GetKubernetesLogsAutocompleteResult) SetContainerNames(v []string)`

SetContainerNames sets ContainerNames field to given value.


### GetLevels

`func (o *GetKubernetesLogsAutocompleteResult) GetLevels() []LogSeverity`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *GetKubernetesLogsAutocompleteResult) GetLevelsOk() (*[]LogSeverity, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *GetKubernetesLogsAutocompleteResult) SetLevels(v []LogSeverity)`

SetLevels sets Levels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


