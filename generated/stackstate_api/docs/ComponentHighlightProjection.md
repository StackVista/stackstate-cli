# ComponentHighlightProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | Cel expression that returns a map&lt;string,dyn&gt; | 
**ShowChart** | Pointer to **bool** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Query** | **string** | Individual metric query that returns a timeseries for a specific cell. | 
**Name** | **string** | Cel expression that returns a string that represents the name of the component to link to | 
**Identifier** | **string** | Cel expression that returns a string that represents the identifier in order to build the link | 
**AsTag** | Pointer to **bool** | Should the value be rendered as a tag or as plain text | [optional] [default to false]
**StartTime** | **string** | Cel expression that returns an ISO8601 formatted timestamp | 
**EndTime** | Pointer to **string** | Cel expression that returns an ISO8601 formatted timestamp | [optional] 
**Numerator** | **string** | Cel expression that returns a number | 
**Denominator** | **string** | Cel expression that returns a number | 
**Status** | Pointer to **string** | Cel expression that returns a string that represents a valid HealthState | [optional] 
**ImageId** | **string** | Cel expression that returns a string | 
**ImageName** | **string** | Cel expression that returns a string | 
**AsTags** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewComponentHighlightProjection

`func NewComponentHighlightProjection(type_ string, value string, query string, name string, identifier string, startTime string, numerator string, denominator string, imageId string, imageName string, ) *ComponentHighlightProjection`

NewComponentHighlightProjection instantiates a new ComponentHighlightProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentHighlightProjectionWithDefaults

`func NewComponentHighlightProjectionWithDefaults() *ComponentHighlightProjection`

NewComponentHighlightProjectionWithDefaults instantiates a new ComponentHighlightProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentHighlightProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentHighlightProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentHighlightProjection) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ComponentHighlightProjection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentHighlightProjection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentHighlightProjection) SetValue(v string)`

SetValue sets Value field to given value.


### GetShowChart

`func (o *ComponentHighlightProjection) GetShowChart() bool`

GetShowChart returns the ShowChart field if non-nil, zero value otherwise.

### GetShowChartOk

`func (o *ComponentHighlightProjection) GetShowChartOk() (*bool, bool)`

GetShowChartOk returns a tuple with the ShowChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowChart

`func (o *ComponentHighlightProjection) SetShowChart(v bool)`

SetShowChart sets ShowChart field to given value.

### HasShowChart

`func (o *ComponentHighlightProjection) HasShowChart() bool`

HasShowChart returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *ComponentHighlightProjection) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *ComponentHighlightProjection) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *ComponentHighlightProjection) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *ComponentHighlightProjection) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetUnit

`func (o *ComponentHighlightProjection) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ComponentHighlightProjection) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ComponentHighlightProjection) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ComponentHighlightProjection) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *ComponentHighlightProjection) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *ComponentHighlightProjection) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetQuery

`func (o *ComponentHighlightProjection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ComponentHighlightProjection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ComponentHighlightProjection) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetName

`func (o *ComponentHighlightProjection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentHighlightProjection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentHighlightProjection) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *ComponentHighlightProjection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ComponentHighlightProjection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ComponentHighlightProjection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetAsTag

`func (o *ComponentHighlightProjection) GetAsTag() bool`

GetAsTag returns the AsTag field if non-nil, zero value otherwise.

### GetAsTagOk

`func (o *ComponentHighlightProjection) GetAsTagOk() (*bool, bool)`

GetAsTagOk returns a tuple with the AsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTag

`func (o *ComponentHighlightProjection) SetAsTag(v bool)`

SetAsTag sets AsTag field to given value.

### HasAsTag

`func (o *ComponentHighlightProjection) HasAsTag() bool`

HasAsTag returns a boolean if a field has been set.

### GetStartTime

`func (o *ComponentHighlightProjection) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ComponentHighlightProjection) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ComponentHighlightProjection) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ComponentHighlightProjection) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ComponentHighlightProjection) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ComponentHighlightProjection) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ComponentHighlightProjection) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetNumerator

`func (o *ComponentHighlightProjection) GetNumerator() string`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *ComponentHighlightProjection) GetNumeratorOk() (*string, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *ComponentHighlightProjection) SetNumerator(v string)`

SetNumerator sets Numerator field to given value.


### GetDenominator

`func (o *ComponentHighlightProjection) GetDenominator() string`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *ComponentHighlightProjection) GetDenominatorOk() (*string, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *ComponentHighlightProjection) SetDenominator(v string)`

SetDenominator sets Denominator field to given value.


### GetStatus

`func (o *ComponentHighlightProjection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComponentHighlightProjection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComponentHighlightProjection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComponentHighlightProjection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageId

`func (o *ComponentHighlightProjection) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ComponentHighlightProjection) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ComponentHighlightProjection) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetImageName

`func (o *ComponentHighlightProjection) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ComponentHighlightProjection) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ComponentHighlightProjection) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetAsTags

`func (o *ComponentHighlightProjection) GetAsTags() bool`

GetAsTags returns the AsTags field if non-nil, zero value otherwise.

### GetAsTagsOk

`func (o *ComponentHighlightProjection) GetAsTagsOk() (*bool, bool)`

GetAsTagsOk returns a tuple with the AsTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTags

`func (o *ComponentHighlightProjection) SetAsTags(v bool)`

SetAsTags sets AsTags field to given value.

### HasAsTags

`func (o *ComponentHighlightProjection) HasAsTags() bool`

HasAsTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


