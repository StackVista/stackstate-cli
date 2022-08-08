# EventTelemetryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Conditions** | [**[]TelemetryQueryCondition**](TelemetryQueryCondition.md) |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewEventTelemetryQuery

`func NewEventTelemetryQuery(type_ string, conditions []TelemetryQueryCondition, ) *EventTelemetryQuery`

NewEventTelemetryQuery instantiates a new EventTelemetryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTelemetryQueryWithDefaults

`func NewEventTelemetryQueryWithDefaults() *EventTelemetryQuery`

NewEventTelemetryQueryWithDefaults instantiates a new EventTelemetryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventTelemetryQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventTelemetryQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventTelemetryQuery) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *EventTelemetryQuery) GetConditions() []TelemetryQueryCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EventTelemetryQuery) GetConditionsOk() (*[]TelemetryQueryCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EventTelemetryQuery) SetConditions(v []TelemetryQueryCondition)`

SetConditions sets Conditions field to given value.


### GetId

`func (o *EventTelemetryQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTelemetryQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTelemetryQuery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventTelemetryQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *EventTelemetryQuery) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *EventTelemetryQuery) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *EventTelemetryQuery) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *EventTelemetryQuery) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


