# PaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of items matching the ComponentPresentation before applying column filters or search.  | 
**Filtered** | Pointer to **NullableInt32** | Total number of items after filters/search are applied, but before pagination.  | [optional] 
**ElementName** | **string** |  | 
**ElementNamePlural** | **string** |  | 
**StartCursor** | Pointer to **NullableString** |  | [optional] 
**EndCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginationResponse

`func NewPaginationResponse(total int32, elementName string, elementNamePlural string, ) *PaginationResponse`

NewPaginationResponse instantiates a new PaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseWithDefaults

`func NewPaginationResponseWithDefaults() *PaginationResponse`

NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PaginationResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetFiltered

`func (o *PaginationResponse) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *PaginationResponse) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *PaginationResponse) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *PaginationResponse) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### SetFilteredNil

`func (o *PaginationResponse) SetFilteredNil(b bool)`

 SetFilteredNil sets the value for Filtered to be an explicit nil

### UnsetFiltered
`func (o *PaginationResponse) UnsetFiltered()`

UnsetFiltered ensures that no value is present for Filtered, not even an explicit nil
### GetElementName

`func (o *PaginationResponse) GetElementName() string`

GetElementName returns the ElementName field if non-nil, zero value otherwise.

### GetElementNameOk

`func (o *PaginationResponse) GetElementNameOk() (*string, bool)`

GetElementNameOk returns a tuple with the ElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementName

`func (o *PaginationResponse) SetElementName(v string)`

SetElementName sets ElementName field to given value.


### GetElementNamePlural

`func (o *PaginationResponse) GetElementNamePlural() string`

GetElementNamePlural returns the ElementNamePlural field if non-nil, zero value otherwise.

### GetElementNamePluralOk

`func (o *PaginationResponse) GetElementNamePluralOk() (*string, bool)`

GetElementNamePluralOk returns a tuple with the ElementNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementNamePlural

`func (o *PaginationResponse) SetElementNamePlural(v string)`

SetElementNamePlural sets ElementNamePlural field to given value.


### GetStartCursor

`func (o *PaginationResponse) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *PaginationResponse) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *PaginationResponse) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *PaginationResponse) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### SetStartCursorNil

`func (o *PaginationResponse) SetStartCursorNil(b bool)`

 SetStartCursorNil sets the value for StartCursor to be an explicit nil

### UnsetStartCursor
`func (o *PaginationResponse) UnsetStartCursor()`

UnsetStartCursor ensures that no value is present for StartCursor, not even an explicit nil
### GetEndCursor

`func (o *PaginationResponse) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *PaginationResponse) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *PaginationResponse) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *PaginationResponse) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### SetEndCursorNil

`func (o *PaginationResponse) SetEndCursorNil(b bool)`

 SetEndCursorNil sets the value for EndCursor to be an explicit nil

### UnsetEndCursor
`func (o *PaginationResponse) UnsetEndCursor()`

UnsetEndCursor ensures that no value is present for EndCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


