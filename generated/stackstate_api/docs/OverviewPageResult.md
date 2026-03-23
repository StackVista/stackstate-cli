# OverviewPageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Metadata** | [**OverviewMetadata**](OverviewMetadata.md) |  | 
**Data** | [**[]OverviewRow**](OverviewRow.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewOverviewPageResult

`func NewOverviewPageResult(type_ string, metadata OverviewMetadata, data []OverviewRow, pagination PaginationResponse, ) *OverviewPageResult`

NewOverviewPageResult instantiates a new OverviewPageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewPageResultWithDefaults

`func NewOverviewPageResultWithDefaults() *OverviewPageResult`

NewOverviewPageResultWithDefaults instantiates a new OverviewPageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewPageResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewPageResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewPageResult) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *OverviewPageResult) GetMetadata() OverviewMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OverviewPageResult) GetMetadataOk() (*OverviewMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OverviewPageResult) SetMetadata(v OverviewMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *OverviewPageResult) GetData() []OverviewRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OverviewPageResult) GetDataOk() (*[]OverviewRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OverviewPageResult) SetData(v []OverviewRow)`

SetData sets Data field to given value.


### GetPagination

`func (o *OverviewPageResult) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OverviewPageResult) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OverviewPageResult) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


