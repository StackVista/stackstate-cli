# MonitorCheckStatusMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Query** | [**PromqlMetricQuery**](PromqlMetricQuery.md) |  | 

## Methods

### NewMonitorCheckStatusMetric

`func NewMonitorCheckStatusMetric(type_ string, name string, query PromqlMetricQuery, ) *MonitorCheckStatusMetric`

NewMonitorCheckStatusMetric instantiates a new MonitorCheckStatusMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusMetricWithDefaults

`func NewMonitorCheckStatusMetricWithDefaults() *MonitorCheckStatusMetric`

NewMonitorCheckStatusMetricWithDefaults instantiates a new MonitorCheckStatusMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorCheckStatusMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorCheckStatusMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorCheckStatusMetric) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *MonitorCheckStatusMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorCheckStatusMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorCheckStatusMetric) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *MonitorCheckStatusMetric) GetQuery() PromqlMetricQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorCheckStatusMetric) GetQueryOk() (*PromqlMetricQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorCheckStatusMetric) SetQuery(v PromqlMetricQuery)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


