# OtelVariableMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** | An expression that can produce any type. It uses the CEL expression within curly braces &#x60;${}&#x60; syntax. Variables use it to store any type of value. For example, to store a boolean in a variable named  &#x60;inTestNamespace&#x60; assign  it the expression &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;] &#x3D;&#x3D; &#39;test&#39;}\&quot;&#x60;. The variable can now be used directly in the conditions like this: &#x60;vars.inTestNamespace&#x60;. | 

## Methods

### NewOtelVariableMapping

`func NewOtelVariableMapping(name string, value string, ) *OtelVariableMapping`

NewOtelVariableMapping instantiates a new OtelVariableMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelVariableMappingWithDefaults

`func NewOtelVariableMappingWithDefaults() *OtelVariableMapping`

NewOtelVariableMappingWithDefaults instantiates a new OtelVariableMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OtelVariableMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelVariableMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelVariableMapping) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *OtelVariableMapping) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OtelVariableMapping) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OtelVariableMapping) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


