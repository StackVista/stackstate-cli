# OverviewColumnProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ShowChart** | Pointer to **bool** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Query** | **string** | Individual metric query that returns a timeseries for a specific cell. | 
**Name** | **string** | Cel expression that returns a string that represents the name of the component to link to | 
**ComponentIdentifier** | **string** | Cel expression that returns a string that represents the componentIdentifier in order to build the link | 
**Value** | **string** | Cel expression that returns a number | 
**StartDate** | **string** | Cel expression that returns a date | 
**EndDate** | Pointer to **string** | Cel expression that returns a date | [optional] 
**ReadyNumber** | **string** | Cel expression that returns a number | 
**TotalNumber** | **string** | Cel expression that returns a number | 
**ReadyStatus** | Pointer to **string** | Cel expression that returns a string that represents a valid HealthState | [optional] 
**ImageId** | **string** | Cel expression that returns a string | 
**ImageName** | **string** | Cel expression that returns a string | 

## Methods

### NewOverviewColumnProjection

`func NewOverviewColumnProjection(type_ string, query string, name string, componentIdentifier string, value string, startDate string, readyNumber string, totalNumber string, imageId string, imageName string, ) *OverviewColumnProjection`

NewOverviewColumnProjection instantiates a new OverviewColumnProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewColumnProjectionWithDefaults

`func NewOverviewColumnProjectionWithDefaults() *OverviewColumnProjection`

NewOverviewColumnProjectionWithDefaults instantiates a new OverviewColumnProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewColumnProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewColumnProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewColumnProjection) SetType(v string)`

SetType sets Type field to given value.


### GetShowChart

`func (o *OverviewColumnProjection) GetShowChart() bool`

GetShowChart returns the ShowChart field if non-nil, zero value otherwise.

### GetShowChartOk

`func (o *OverviewColumnProjection) GetShowChartOk() (*bool, bool)`

GetShowChartOk returns a tuple with the ShowChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowChart

`func (o *OverviewColumnProjection) SetShowChart(v bool)`

SetShowChart sets ShowChart field to given value.

### HasShowChart

`func (o *OverviewColumnProjection) HasShowChart() bool`

HasShowChart returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *OverviewColumnProjection) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *OverviewColumnProjection) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *OverviewColumnProjection) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *OverviewColumnProjection) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetUnit

`func (o *OverviewColumnProjection) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OverviewColumnProjection) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OverviewColumnProjection) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OverviewColumnProjection) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *OverviewColumnProjection) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *OverviewColumnProjection) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetQuery

`func (o *OverviewColumnProjection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OverviewColumnProjection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OverviewColumnProjection) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetName

`func (o *OverviewColumnProjection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverviewColumnProjection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverviewColumnProjection) SetName(v string)`

SetName sets Name field to given value.


### GetComponentIdentifier

`func (o *OverviewColumnProjection) GetComponentIdentifier() string`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *OverviewColumnProjection) GetComponentIdentifierOk() (*string, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *OverviewColumnProjection) SetComponentIdentifier(v string)`

SetComponentIdentifier sets ComponentIdentifier field to given value.


### GetValue

`func (o *OverviewColumnProjection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OverviewColumnProjection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OverviewColumnProjection) SetValue(v string)`

SetValue sets Value field to given value.


### GetStartDate

`func (o *OverviewColumnProjection) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OverviewColumnProjection) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OverviewColumnProjection) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *OverviewColumnProjection) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OverviewColumnProjection) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OverviewColumnProjection) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OverviewColumnProjection) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReadyNumber

`func (o *OverviewColumnProjection) GetReadyNumber() string`

GetReadyNumber returns the ReadyNumber field if non-nil, zero value otherwise.

### GetReadyNumberOk

`func (o *OverviewColumnProjection) GetReadyNumberOk() (*string, bool)`

GetReadyNumberOk returns a tuple with the ReadyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyNumber

`func (o *OverviewColumnProjection) SetReadyNumber(v string)`

SetReadyNumber sets ReadyNumber field to given value.


### GetTotalNumber

`func (o *OverviewColumnProjection) GetTotalNumber() string`

GetTotalNumber returns the TotalNumber field if non-nil, zero value otherwise.

### GetTotalNumberOk

`func (o *OverviewColumnProjection) GetTotalNumberOk() (*string, bool)`

GetTotalNumberOk returns a tuple with the TotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumber

`func (o *OverviewColumnProjection) SetTotalNumber(v string)`

SetTotalNumber sets TotalNumber field to given value.


### GetReadyStatus

`func (o *OverviewColumnProjection) GetReadyStatus() string`

GetReadyStatus returns the ReadyStatus field if non-nil, zero value otherwise.

### GetReadyStatusOk

`func (o *OverviewColumnProjection) GetReadyStatusOk() (*string, bool)`

GetReadyStatusOk returns a tuple with the ReadyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyStatus

`func (o *OverviewColumnProjection) SetReadyStatus(v string)`

SetReadyStatus sets ReadyStatus field to given value.

### HasReadyStatus

`func (o *OverviewColumnProjection) HasReadyStatus() bool`

HasReadyStatus returns a boolean if a field has been set.

### GetImageId

`func (o *OverviewColumnProjection) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *OverviewColumnProjection) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *OverviewColumnProjection) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetImageName

`func (o *OverviewColumnProjection) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *OverviewColumnProjection) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *OverviewColumnProjection) SetImageName(v string)`

SetImageName sets ImageName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


