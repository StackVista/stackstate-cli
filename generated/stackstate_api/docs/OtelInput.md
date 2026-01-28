# OtelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signal** | [**[]OtelInputSignal**](OtelInputSignal.md) |  | 
**Resource** | [**OtelInputResource**](OtelInputResource.md) |  | 

## Methods

### NewOtelInput

`func NewOtelInput(signal []OtelInputSignal, resource OtelInputResource, ) *OtelInput`

NewOtelInput instantiates a new OtelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelInputWithDefaults

`func NewOtelInputWithDefaults() *OtelInput`

NewOtelInputWithDefaults instantiates a new OtelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignal

`func (o *OtelInput) GetSignal() []OtelInputSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *OtelInput) GetSignalOk() (*[]OtelInputSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *OtelInput) SetSignal(v []OtelInputSignal)`

SetSignal sets Signal field to given value.


### GetResource

`func (o *OtelInput) GetResource() OtelInputResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OtelInput) GetResourceOk() (*OtelInputResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OtelInput) SetResource(v OtelInputResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


