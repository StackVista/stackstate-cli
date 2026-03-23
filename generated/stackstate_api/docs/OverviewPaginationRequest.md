# OverviewPaginationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**OverviewPaginationCursor**](OverviewPaginationCursor.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 50]

## Methods

### NewOverviewPaginationRequest

`func NewOverviewPaginationRequest() *OverviewPaginationRequest`

NewOverviewPaginationRequest instantiates a new OverviewPaginationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewPaginationRequestWithDefaults

`func NewOverviewPaginationRequestWithDefaults() *OverviewPaginationRequest`

NewOverviewPaginationRequestWithDefaults instantiates a new OverviewPaginationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *OverviewPaginationRequest) GetCursor() OverviewPaginationCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OverviewPaginationRequest) GetCursorOk() (*OverviewPaginationCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OverviewPaginationRequest) SetCursor(v OverviewPaginationCursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *OverviewPaginationRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetLimit

`func (o *OverviewPaginationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OverviewPaginationRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OverviewPaginationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *OverviewPaginationRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


