# OtelRelationMappingOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**TargetId** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**TypeName** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**TypeIdentifier** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | [optional] 

## Methods

### NewOtelRelationMappingOutput

`func NewOtelRelationMappingOutput(sourceId string, targetId string, typeName string, ) *OtelRelationMappingOutput`

NewOtelRelationMappingOutput instantiates a new OtelRelationMappingOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelRelationMappingOutputWithDefaults

`func NewOtelRelationMappingOutputWithDefaults() *OtelRelationMappingOutput`

NewOtelRelationMappingOutputWithDefaults instantiates a new OtelRelationMappingOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *OtelRelationMappingOutput) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *OtelRelationMappingOutput) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *OtelRelationMappingOutput) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetId

`func (o *OtelRelationMappingOutput) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *OtelRelationMappingOutput) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *OtelRelationMappingOutput) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetTypeName

`func (o *OtelRelationMappingOutput) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *OtelRelationMappingOutput) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *OtelRelationMappingOutput) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetTypeIdentifier

`func (o *OtelRelationMappingOutput) GetTypeIdentifier() string`

GetTypeIdentifier returns the TypeIdentifier field if non-nil, zero value otherwise.

### GetTypeIdentifierOk

`func (o *OtelRelationMappingOutput) GetTypeIdentifierOk() (*string, bool)`

GetTypeIdentifierOk returns a tuple with the TypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIdentifier

`func (o *OtelRelationMappingOutput) SetTypeIdentifier(v string)`

SetTypeIdentifier sets TypeIdentifier field to given value.

### HasTypeIdentifier

`func (o *OtelRelationMappingOutput) HasTypeIdentifier() bool`

HasTypeIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


