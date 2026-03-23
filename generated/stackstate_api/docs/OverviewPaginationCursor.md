# OverviewPaginationCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **NullableString** | Opaque cursor returned from previous request | 
**Direction** | [**OverviewPaginationDirection**](OverviewPaginationDirection.md) |  | 

## Methods

### NewOverviewPaginationCursor

`func NewOverviewPaginationCursor(cursor NullableString, direction OverviewPaginationDirection, ) *OverviewPaginationCursor`

NewOverviewPaginationCursor instantiates a new OverviewPaginationCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewPaginationCursorWithDefaults

`func NewOverviewPaginationCursorWithDefaults() *OverviewPaginationCursor`

NewOverviewPaginationCursorWithDefaults instantiates a new OverviewPaginationCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *OverviewPaginationCursor) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OverviewPaginationCursor) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OverviewPaginationCursor) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### SetCursorNil

`func (o *OverviewPaginationCursor) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *OverviewPaginationCursor) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetDirection

`func (o *OverviewPaginationCursor) GetDirection() OverviewPaginationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OverviewPaginationCursor) GetDirectionOk() (*OverviewPaginationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OverviewPaginationCursor) SetDirection(v OverviewPaginationDirection)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


