# CheckState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 
**LastHealthStateChangeTimestamp** | **int64** |  | 
**ActiveErrors** | [**[]CheckError**](CheckError.md) |  | 
**Source** | **int64** |  | 
**StateAcknowledgement** | Pointer to [**CheckStateAcknowledgement**](CheckStateAcknowledgement.md) |  | [optional] 
**DetailedMessage** | Pointer to **string** |  | [optional] 
**ShortMessage** | Pointer to **string** |  | [optional] 
**CausingEvents** | [**[]EventRef**](EventRef.md) |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCheckState

`func NewCheckState(type_ string, state HealthStateValue, lastHealthStateChangeTimestamp int64, activeErrors []CheckError, source int64, causingEvents []EventRef, ) *CheckState`

NewCheckState instantiates a new CheckState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckStateWithDefaults

`func NewCheckStateWithDefaults() *CheckState`

NewCheckStateWithDefaults instantiates a new CheckState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckState) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CheckState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckState) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *CheckState) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *CheckState) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *CheckState) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *CheckState) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetState

`func (o *CheckState) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CheckState) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CheckState) SetState(v HealthStateValue)`

SetState sets State field to given value.


### GetLastHealthStateChangeTimestamp

`func (o *CheckState) GetLastHealthStateChangeTimestamp() int64`

GetLastHealthStateChangeTimestamp returns the LastHealthStateChangeTimestamp field if non-nil, zero value otherwise.

### GetLastHealthStateChangeTimestampOk

`func (o *CheckState) GetLastHealthStateChangeTimestampOk() (*int64, bool)`

GetLastHealthStateChangeTimestampOk returns a tuple with the LastHealthStateChangeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHealthStateChangeTimestamp

`func (o *CheckState) SetLastHealthStateChangeTimestamp(v int64)`

SetLastHealthStateChangeTimestamp sets LastHealthStateChangeTimestamp field to given value.


### GetActiveErrors

`func (o *CheckState) GetActiveErrors() []CheckError`

GetActiveErrors returns the ActiveErrors field if non-nil, zero value otherwise.

### GetActiveErrorsOk

`func (o *CheckState) GetActiveErrorsOk() (*[]CheckError, bool)`

GetActiveErrorsOk returns a tuple with the ActiveErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveErrors

`func (o *CheckState) SetActiveErrors(v []CheckError)`

SetActiveErrors sets ActiveErrors field to given value.


### GetSource

`func (o *CheckState) GetSource() int64`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CheckState) GetSourceOk() (*int64, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CheckState) SetSource(v int64)`

SetSource sets Source field to given value.


### GetStateAcknowledgement

`func (o *CheckState) GetStateAcknowledgement() CheckStateAcknowledgement`

GetStateAcknowledgement returns the StateAcknowledgement field if non-nil, zero value otherwise.

### GetStateAcknowledgementOk

`func (o *CheckState) GetStateAcknowledgementOk() (*CheckStateAcknowledgement, bool)`

GetStateAcknowledgementOk returns a tuple with the StateAcknowledgement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateAcknowledgement

`func (o *CheckState) SetStateAcknowledgement(v CheckStateAcknowledgement)`

SetStateAcknowledgement sets StateAcknowledgement field to given value.

### HasStateAcknowledgement

`func (o *CheckState) HasStateAcknowledgement() bool`

HasStateAcknowledgement returns a boolean if a field has been set.

### GetDetailedMessage

`func (o *CheckState) GetDetailedMessage() string`

GetDetailedMessage returns the DetailedMessage field if non-nil, zero value otherwise.

### GetDetailedMessageOk

`func (o *CheckState) GetDetailedMessageOk() (*string, bool)`

GetDetailedMessageOk returns a tuple with the DetailedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedMessage

`func (o *CheckState) SetDetailedMessage(v string)`

SetDetailedMessage sets DetailedMessage field to given value.

### HasDetailedMessage

`func (o *CheckState) HasDetailedMessage() bool`

HasDetailedMessage returns a boolean if a field has been set.

### GetShortMessage

`func (o *CheckState) GetShortMessage() string`

GetShortMessage returns the ShortMessage field if non-nil, zero value otherwise.

### GetShortMessageOk

`func (o *CheckState) GetShortMessageOk() (*string, bool)`

GetShortMessageOk returns a tuple with the ShortMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMessage

`func (o *CheckState) SetShortMessage(v string)`

SetShortMessage sets ShortMessage field to given value.

### HasShortMessage

`func (o *CheckState) HasShortMessage() bool`

HasShortMessage returns a boolean if a field has been set.

### GetCausingEvents

`func (o *CheckState) GetCausingEvents() []EventRef`

GetCausingEvents returns the CausingEvents field if non-nil, zero value otherwise.

### GetCausingEventsOk

`func (o *CheckState) GetCausingEventsOk() (*[]EventRef, bool)`

GetCausingEventsOk returns a tuple with the CausingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausingEvents

`func (o *CheckState) SetCausingEvents(v []EventRef)`

SetCausingEvents sets CausingEvents field to given value.


### GetData

`func (o *CheckState) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckState) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckState) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CheckState) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


