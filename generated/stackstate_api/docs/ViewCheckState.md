# ViewCheckState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStateId** | **string** |  | 
**HealthState** | [**HealthStateValue**](HealthStateValue.md) |  | 
**ComponentName** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | **string** |  | 
**ComponentType** | Pointer to **string** |  | [optional] 
**LastUpdateTimestamp** | **int64** |  | 

## Methods

### NewViewCheckState

`func NewViewCheckState(checkStateId string, healthState HealthStateValue, componentIdentifier string, lastUpdateTimestamp int64, ) *ViewCheckState`

NewViewCheckState instantiates a new ViewCheckState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCheckStateWithDefaults

`func NewViewCheckStateWithDefaults() *ViewCheckState`

NewViewCheckStateWithDefaults instantiates a new ViewCheckState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStateId

`func (o *ViewCheckState) GetCheckStateId() string`

GetCheckStateId returns the CheckStateId field if non-nil, zero value otherwise.

### GetCheckStateIdOk

`func (o *ViewCheckState) GetCheckStateIdOk() (*string, bool)`

GetCheckStateIdOk returns a tuple with the CheckStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStateId

`func (o *ViewCheckState) SetCheckStateId(v string)`

SetCheckStateId sets CheckStateId field to given value.


### GetHealthState

`func (o *ViewCheckState) GetHealthState() HealthStateValue`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *ViewCheckState) GetHealthStateOk() (*HealthStateValue, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *ViewCheckState) SetHealthState(v HealthStateValue)`

SetHealthState sets HealthState field to given value.


### GetComponentName

`func (o *ViewCheckState) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ViewCheckState) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ViewCheckState) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ViewCheckState) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ViewCheckState) GetComponentIdentifier() string`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ViewCheckState) GetComponentIdentifierOk() (*string, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ViewCheckState) SetComponentIdentifier(v string)`

SetComponentIdentifier sets ComponentIdentifier field to given value.


### GetComponentType

`func (o *ViewCheckState) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *ViewCheckState) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *ViewCheckState) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *ViewCheckState) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ViewCheckState) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ViewCheckState) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ViewCheckState) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


