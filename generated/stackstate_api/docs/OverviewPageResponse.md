# OverviewPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Metadata** | [**OverviewMetadata**](OverviewMetadata.md) |  | 
**Data** | [**[]OverviewRow**](OverviewRow.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 
**UsedTimeoutSeconds** | **int64** | Timeout used for the overview request (seconds). | 
**MaxSize** | **int32** | Maximum allowed topology size. | 
**UnavailableAtEpochMs** | **int64** | Time when data became unavailable (epoch ms). | 
**AvailableSinceEpochMs** | **int64** | Time when data became available again (epoch ms). | 

## Methods

### NewOverviewPageResponse

`func NewOverviewPageResponse(type_ string, metadata OverviewMetadata, data []OverviewRow, pagination PaginationResponse, usedTimeoutSeconds int64, maxSize int32, unavailableAtEpochMs int64, availableSinceEpochMs int64, ) *OverviewPageResponse`

NewOverviewPageResponse instantiates a new OverviewPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewPageResponseWithDefaults

`func NewOverviewPageResponseWithDefaults() *OverviewPageResponse`

NewOverviewPageResponseWithDefaults instantiates a new OverviewPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewPageResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewPageResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewPageResponse) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *OverviewPageResponse) GetMetadata() OverviewMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OverviewPageResponse) GetMetadataOk() (*OverviewMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OverviewPageResponse) SetMetadata(v OverviewMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *OverviewPageResponse) GetData() []OverviewRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OverviewPageResponse) GetDataOk() (*[]OverviewRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OverviewPageResponse) SetData(v []OverviewRow)`

SetData sets Data field to given value.


### GetPagination

`func (o *OverviewPageResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OverviewPageResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OverviewPageResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.


### GetUsedTimeoutSeconds

`func (o *OverviewPageResponse) GetUsedTimeoutSeconds() int64`

GetUsedTimeoutSeconds returns the UsedTimeoutSeconds field if non-nil, zero value otherwise.

### GetUsedTimeoutSecondsOk

`func (o *OverviewPageResponse) GetUsedTimeoutSecondsOk() (*int64, bool)`

GetUsedTimeoutSecondsOk returns a tuple with the UsedTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTimeoutSeconds

`func (o *OverviewPageResponse) SetUsedTimeoutSeconds(v int64)`

SetUsedTimeoutSeconds sets UsedTimeoutSeconds field to given value.


### GetMaxSize

`func (o *OverviewPageResponse) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *OverviewPageResponse) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *OverviewPageResponse) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### GetUnavailableAtEpochMs

`func (o *OverviewPageResponse) GetUnavailableAtEpochMs() int64`

GetUnavailableAtEpochMs returns the UnavailableAtEpochMs field if non-nil, zero value otherwise.

### GetUnavailableAtEpochMsOk

`func (o *OverviewPageResponse) GetUnavailableAtEpochMsOk() (*int64, bool)`

GetUnavailableAtEpochMsOk returns a tuple with the UnavailableAtEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableAtEpochMs

`func (o *OverviewPageResponse) SetUnavailableAtEpochMs(v int64)`

SetUnavailableAtEpochMs sets UnavailableAtEpochMs field to given value.


### GetAvailableSinceEpochMs

`func (o *OverviewPageResponse) GetAvailableSinceEpochMs() int64`

GetAvailableSinceEpochMs returns the AvailableSinceEpochMs field if non-nil, zero value otherwise.

### GetAvailableSinceEpochMsOk

`func (o *OverviewPageResponse) GetAvailableSinceEpochMsOk() (*int64, bool)`

GetAvailableSinceEpochMsOk returns a tuple with the AvailableSinceEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSinceEpochMs

`func (o *OverviewPageResponse) SetAvailableSinceEpochMs(v int64)`

SetAvailableSinceEpochMs sets AvailableSinceEpochMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


