# ComponentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Message** | **string** |  | 
**Query** | **string** | The frontend can use the query to refresh the metrics (at interval). Allows to keep the current behaviour of making individual calls for each row. | 
**Unit** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Url** | **string** |  | 
**Identifier** | **string** |  | 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 
**Value** | **float32** |  | 
**AsTag** | **bool** |  | [default to false]
**StartTime** | **int32** |  | 
**EndTime** | Pointer to **NullableInt32** |  | [optional] 
**Numerator** | **float32** |  | 
**Denominator** | **float32** |  | 
**Status** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 
**Values** | **map[string]string** |  | 

## Methods

### NewComponentField

`func NewComponentField(fieldId string, title string, type_ string, message string, query string, name string, url string, identifier string, state HealthStateValue, value float32, asTag bool, startTime int32, numerator float32, denominator float32, values map[string]string, ) *ComponentField`

NewComponentField instantiates a new ComponentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentFieldWithDefaults

`func NewComponentFieldWithDefaults() *ComponentField`

NewComponentFieldWithDefaults instantiates a new ComponentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *ComponentField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ComponentField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ComponentField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *ComponentField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ComponentField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ComponentField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ComponentField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ComponentField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentField) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ComponentField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ComponentField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ComponentField) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetQuery

`func (o *ComponentField) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ComponentField) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ComponentField) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUnit

`func (o *ComponentField) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ComponentField) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ComponentField) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ComponentField) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *ComponentField) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *ComponentField) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *ComponentField) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *ComponentField) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetName

`func (o *ComponentField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentField) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *ComponentField) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComponentField) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComponentField) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIdentifier

`func (o *ComponentField) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ComponentField) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ComponentField) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetState

`func (o *ComponentField) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComponentField) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComponentField) SetState(v HealthStateValue)`

SetState sets State field to given value.


### GetValue

`func (o *ComponentField) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentField) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentField) SetValue(v float32)`

SetValue sets Value field to given value.


### GetAsTag

`func (o *ComponentField) GetAsTag() bool`

GetAsTag returns the AsTag field if non-nil, zero value otherwise.

### GetAsTagOk

`func (o *ComponentField) GetAsTagOk() (*bool, bool)`

GetAsTagOk returns a tuple with the AsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTag

`func (o *ComponentField) SetAsTag(v bool)`

SetAsTag sets AsTag field to given value.


### GetStartTime

`func (o *ComponentField) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ComponentField) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ComponentField) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ComponentField) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ComponentField) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ComponentField) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ComponentField) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ComponentField) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ComponentField) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetNumerator

`func (o *ComponentField) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *ComponentField) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *ComponentField) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.


### GetDenominator

`func (o *ComponentField) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *ComponentField) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *ComponentField) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.


### GetStatus

`func (o *ComponentField) GetStatus() HealthStateValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComponentField) GetStatusOk() (*HealthStateValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComponentField) SetStatus(v HealthStateValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComponentField) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValues

`func (o *ComponentField) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ComponentField) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ComponentField) SetValues(v map[string]string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


