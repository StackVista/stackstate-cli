# CellValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IndividualQuery** | **string** | The frontend can use the individual query to refresh the metrics (at interval). Allows to keep the current behaviour of making individual calls for each row. | 
**Name** | **string** |  | 
**Url** | **string** |  | 
**Component** | Pointer to [**ComponentLink**](ComponentLink.md) |  | [optional] 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 
**Value** | **float32** |  | 
**StartDate** | **int32** |  | 
**EndDate** | Pointer to **NullableInt32** |  | [optional] 
**Ready** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 

## Methods

### NewCellValue

`func NewCellValue(type_ string, individualQuery string, name string, url string, state HealthStateValue, value float32, startDate int32, ) *CellValue`

NewCellValue instantiates a new CellValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellValueWithDefaults

`func NewCellValueWithDefaults() *CellValue`

NewCellValueWithDefaults instantiates a new CellValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CellValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CellValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CellValue) SetType(v string)`

SetType sets Type field to given value.


### GetIndividualQuery

`func (o *CellValue) GetIndividualQuery() string`

GetIndividualQuery returns the IndividualQuery field if non-nil, zero value otherwise.

### GetIndividualQueryOk

`func (o *CellValue) GetIndividualQueryOk() (*string, bool)`

GetIndividualQueryOk returns a tuple with the IndividualQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualQuery

`func (o *CellValue) SetIndividualQuery(v string)`

SetIndividualQuery sets IndividualQuery field to given value.


### GetName

`func (o *CellValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CellValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CellValue) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CellValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CellValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CellValue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetComponent

`func (o *CellValue) GetComponent() ComponentLink`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CellValue) GetComponentOk() (*ComponentLink, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CellValue) SetComponent(v ComponentLink)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CellValue) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetState

`func (o *CellValue) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CellValue) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CellValue) SetState(v HealthStateValue)`

SetState sets State field to given value.


### GetValue

`func (o *CellValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CellValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CellValue) SetValue(v float32)`

SetValue sets Value field to given value.


### GetStartDate

`func (o *CellValue) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CellValue) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CellValue) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CellValue) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CellValue) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CellValue) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CellValue) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CellValue) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CellValue) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetReady

`func (o *CellValue) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *CellValue) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *CellValue) SetReady(v int32)`

SetReady sets Ready field to given value.

### HasReady

`func (o *CellValue) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetTotal

`func (o *CellValue) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CellValue) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CellValue) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CellValue) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetStatus

`func (o *CellValue) GetStatus() HealthStateValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CellValue) GetStatusOk() (*HealthStateValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CellValue) SetStatus(v HealthStateValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CellValue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


