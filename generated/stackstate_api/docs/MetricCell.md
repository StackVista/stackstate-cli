# MetricCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IndividualQuery** | **string** | The frontend can use the individual query to refresh the metrics (at interval). Allows to keep the current behaviour of making individual calls for each row. | 

## Methods

### NewMetricCell

`func NewMetricCell(type_ string, individualQuery string, ) *MetricCell`

NewMetricCell instantiates a new MetricCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCellWithDefaults

`func NewMetricCellWithDefaults() *MetricCell`

NewMetricCellWithDefaults instantiates a new MetricCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricCell) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricCell) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricCell) SetType(v string)`

SetType sets Type field to given value.


### GetIndividualQuery

`func (o *MetricCell) GetIndividualQuery() string`

GetIndividualQuery returns the IndividualQuery field if non-nil, zero value otherwise.

### GetIndividualQueryOk

`func (o *MetricCell) GetIndividualQueryOk() (*string, bool)`

GetIndividualQueryOk returns a tuple with the IndividualQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualQuery

`func (o *MetricCell) SetIndividualQuery(v string)`

SetIndividualQuery sets IndividualQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


