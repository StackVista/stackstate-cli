# Exemplar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **map[string]string** |  | 
**Value** | **string** |  | 
**Timestamp** | **float32** |  | 

## Methods

### NewExemplar

`func NewExemplar(labels map[string]string, value string, timestamp float32, ) *Exemplar`

NewExemplar instantiates a new Exemplar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExemplarWithDefaults

`func NewExemplarWithDefaults() *Exemplar`

NewExemplarWithDefaults instantiates a new Exemplar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *Exemplar) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Exemplar) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Exemplar) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetValue

`func (o *Exemplar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Exemplar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Exemplar) SetValue(v string)`

SetValue sets Value field to given value.


### GetTimestamp

`func (o *Exemplar) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Exemplar) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Exemplar) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


