# GetKubernetesLogsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]KubernetesLogRecord**](KubernetesLogRecord.md) | List of log records ordered by timestamps in ascending order. | 
**PageSize** | **int32** | Maximum number of the log lines in the result. | 
**Page** | **int32** | The requested logs page. | 
**MatchesTotal** | **int64** | The total amount of matching logs for the requested search criteria. | 

## Methods

### NewGetKubernetesLogsResult

`func NewGetKubernetesLogsResult(logs []KubernetesLogRecord, pageSize int32, page int32, matchesTotal int64, ) *GetKubernetesLogsResult`

NewGetKubernetesLogsResult instantiates a new GetKubernetesLogsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesLogsResultWithDefaults

`func NewGetKubernetesLogsResultWithDefaults() *GetKubernetesLogsResult`

NewGetKubernetesLogsResultWithDefaults instantiates a new GetKubernetesLogsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetKubernetesLogsResult) GetLogs() []KubernetesLogRecord`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetKubernetesLogsResult) GetLogsOk() (*[]KubernetesLogRecord, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetKubernetesLogsResult) SetLogs(v []KubernetesLogRecord)`

SetLogs sets Logs field to given value.


### GetPageSize

`func (o *GetKubernetesLogsResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetKubernetesLogsResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetKubernetesLogsResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetKubernetesLogsResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetKubernetesLogsResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetKubernetesLogsResult) SetPage(v int32)`

SetPage sets Page field to given value.


### GetMatchesTotal

`func (o *GetKubernetesLogsResult) GetMatchesTotal() int64`

GetMatchesTotal returns the MatchesTotal field if non-nil, zero value otherwise.

### GetMatchesTotalOk

`func (o *GetKubernetesLogsResult) GetMatchesTotalOk() (*int64, bool)`

GetMatchesTotalOk returns a tuple with the MatchesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesTotal

`func (o *GetKubernetesLogsResult) SetMatchesTotal(v int64)`

SetMatchesTotal sets MatchesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


