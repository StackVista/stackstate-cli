# MonitorCheckStatusRelatedFailures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**TopologyTime** | **int32** |  | 
**CheckStatuses** | [**[]MonitorCheckStatusRelatedFailuresCheckStatus**](MonitorCheckStatusRelatedFailuresCheckStatus.md) |  | 

## Methods

### NewMonitorCheckStatusRelatedFailures

`func NewMonitorCheckStatusRelatedFailures(identifier string, topologyTime int32, checkStatuses []MonitorCheckStatusRelatedFailuresCheckStatus, ) *MonitorCheckStatusRelatedFailures`

NewMonitorCheckStatusRelatedFailures instantiates a new MonitorCheckStatusRelatedFailures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusRelatedFailuresWithDefaults

`func NewMonitorCheckStatusRelatedFailuresWithDefaults() *MonitorCheckStatusRelatedFailures`

NewMonitorCheckStatusRelatedFailuresWithDefaults instantiates a new MonitorCheckStatusRelatedFailures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MonitorCheckStatusRelatedFailures) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MonitorCheckStatusRelatedFailures) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MonitorCheckStatusRelatedFailures) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTopologyTime

`func (o *MonitorCheckStatusRelatedFailures) GetTopologyTime() int32`

GetTopologyTime returns the TopologyTime field if non-nil, zero value otherwise.

### GetTopologyTimeOk

`func (o *MonitorCheckStatusRelatedFailures) GetTopologyTimeOk() (*int32, bool)`

GetTopologyTimeOk returns a tuple with the TopologyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyTime

`func (o *MonitorCheckStatusRelatedFailures) SetTopologyTime(v int32)`

SetTopologyTime sets TopologyTime field to given value.


### GetCheckStatuses

`func (o *MonitorCheckStatusRelatedFailures) GetCheckStatuses() []MonitorCheckStatusRelatedFailuresCheckStatus`

GetCheckStatuses returns the CheckStatuses field if non-nil, zero value otherwise.

### GetCheckStatusesOk

`func (o *MonitorCheckStatusRelatedFailures) GetCheckStatusesOk() (*[]MonitorCheckStatusRelatedFailuresCheckStatus, bool)`

GetCheckStatusesOk returns a tuple with the CheckStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatuses

`func (o *MonitorCheckStatusRelatedFailures) SetCheckStatuses(v []MonitorCheckStatusRelatedFailuresCheckStatus)`

SetCheckStatuses sets CheckStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


