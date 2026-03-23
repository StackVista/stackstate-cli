# OverviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | STQL query string to retrieve and project using the presentationOrViewUrn (helpful for related resources) | [optional] 
**TopologyTime** | Pointer to **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | [optional] 
**Filters** | Pointer to [**OverviewFilters**](OverviewFilters.md) |  | [optional] 
**Pagination** | Pointer to [**OverviewPaginationRequest**](OverviewPaginationRequest.md) |  | [optional] 
**Sorting** | Pointer to [**[]OverviewSorting**](OverviewSorting.md) |  | [optional] 

## Methods

### NewOverviewRequest

`func NewOverviewRequest() *OverviewRequest`

NewOverviewRequest instantiates a new OverviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewRequestWithDefaults

`func NewOverviewRequestWithDefaults() *OverviewRequest`

NewOverviewRequestWithDefaults instantiates a new OverviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *OverviewRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OverviewRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OverviewRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *OverviewRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTopologyTime

`func (o *OverviewRequest) GetTopologyTime() int32`

GetTopologyTime returns the TopologyTime field if non-nil, zero value otherwise.

### GetTopologyTimeOk

`func (o *OverviewRequest) GetTopologyTimeOk() (*int32, bool)`

GetTopologyTimeOk returns a tuple with the TopologyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyTime

`func (o *OverviewRequest) SetTopologyTime(v int32)`

SetTopologyTime sets TopologyTime field to given value.

### HasTopologyTime

`func (o *OverviewRequest) HasTopologyTime() bool`

HasTopologyTime returns a boolean if a field has been set.

### GetFilters

`func (o *OverviewRequest) GetFilters() OverviewFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OverviewRequest) GetFiltersOk() (*OverviewFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OverviewRequest) SetFilters(v OverviewFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OverviewRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPagination

`func (o *OverviewRequest) GetPagination() OverviewPaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OverviewRequest) GetPaginationOk() (*OverviewPaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OverviewRequest) SetPagination(v OverviewPaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *OverviewRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSorting

`func (o *OverviewRequest) GetSorting() []OverviewSorting`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *OverviewRequest) GetSortingOk() (*[]OverviewSorting, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *OverviewRequest) SetSorting(v []OverviewSorting)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *OverviewRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


