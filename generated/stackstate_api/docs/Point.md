# Point

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **int64** |  | 
**V** | **float64** |  | 

## Methods

### NewPoint

`func NewPoint(ts int64, v float64, ) *Point`

NewPoint instantiates a new Point object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointWithDefaults

`func NewPointWithDefaults() *Point`

NewPointWithDefaults instantiates a new Point object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *Point) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *Point) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *Point) SetTs(v int64)`

SetTs sets Ts field to given value.


### GetV

`func (o *Point) GetV() float64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Point) GetVOk() (*float64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Point) SetV(v float64)`

SetV sets V field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


