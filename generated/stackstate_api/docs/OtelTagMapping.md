# OtelTagMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | An expression that can produce any type. It uses the CEL expression within curly braces &#x60;${}&#x60; syntax. Variables use it to store any type of value. For example, to store a boolean in a variable named  &#x60;inTestNamespace&#x60; assign  it the expression &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;] &#x3D;&#x3D; &#39;test&#39;}\&quot;&#x60;. The variable can now be used directly in the conditions like this: &#x60;vars.inTestNamespace&#x60;. | 
**Target** | **string** | Name of the target tag key to which the value should be mapped. | 
**Pattern** | Pointer to **string** | Optional regex pattern applied to the source value. Capturing groups can be referenced in the target (e.g., ${1}). | [optional] 

## Methods

### NewOtelTagMapping

`func NewOtelTagMapping(source string, target string, ) *OtelTagMapping`

NewOtelTagMapping instantiates a new OtelTagMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelTagMappingWithDefaults

`func NewOtelTagMappingWithDefaults() *OtelTagMapping`

NewOtelTagMappingWithDefaults instantiates a new OtelTagMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *OtelTagMapping) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OtelTagMapping) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OtelTagMapping) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *OtelTagMapping) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *OtelTagMapping) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *OtelTagMapping) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetPattern

`func (o *OtelTagMapping) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *OtelTagMapping) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *OtelTagMapping) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *OtelTagMapping) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


