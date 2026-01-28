# OtelComponentMappingOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**Name** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**TypeName** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**TypeIdentifier** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | [optional] 
**LayerName** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**LayerIdentifier** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | [optional] 
**DomainName** | **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | 
**DomainIdentifier** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | [optional] 
**Required** | Pointer to [**OtelComponentMappingFieldMapping**](OtelComponentMappingFieldMapping.md) |  | [optional] 
**Optional** | Pointer to [**OtelComponentMappingFieldMapping**](OtelComponentMappingFieldMapping.md) |  | [optional] 

## Methods

### NewOtelComponentMappingOutput

`func NewOtelComponentMappingOutput(identifier string, name string, typeName string, layerName string, domainName string, ) *OtelComponentMappingOutput`

NewOtelComponentMappingOutput instantiates a new OtelComponentMappingOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelComponentMappingOutputWithDefaults

`func NewOtelComponentMappingOutputWithDefaults() *OtelComponentMappingOutput`

NewOtelComponentMappingOutputWithDefaults instantiates a new OtelComponentMappingOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *OtelComponentMappingOutput) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OtelComponentMappingOutput) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OtelComponentMappingOutput) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *OtelComponentMappingOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelComponentMappingOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelComponentMappingOutput) SetName(v string)`

SetName sets Name field to given value.


### GetTypeName

`func (o *OtelComponentMappingOutput) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *OtelComponentMappingOutput) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *OtelComponentMappingOutput) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetTypeIdentifier

`func (o *OtelComponentMappingOutput) GetTypeIdentifier() string`

GetTypeIdentifier returns the TypeIdentifier field if non-nil, zero value otherwise.

### GetTypeIdentifierOk

`func (o *OtelComponentMappingOutput) GetTypeIdentifierOk() (*string, bool)`

GetTypeIdentifierOk returns a tuple with the TypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIdentifier

`func (o *OtelComponentMappingOutput) SetTypeIdentifier(v string)`

SetTypeIdentifier sets TypeIdentifier field to given value.

### HasTypeIdentifier

`func (o *OtelComponentMappingOutput) HasTypeIdentifier() bool`

HasTypeIdentifier returns a boolean if a field has been set.

### GetLayerName

`func (o *OtelComponentMappingOutput) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *OtelComponentMappingOutput) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *OtelComponentMappingOutput) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.


### GetLayerIdentifier

`func (o *OtelComponentMappingOutput) GetLayerIdentifier() string`

GetLayerIdentifier returns the LayerIdentifier field if non-nil, zero value otherwise.

### GetLayerIdentifierOk

`func (o *OtelComponentMappingOutput) GetLayerIdentifierOk() (*string, bool)`

GetLayerIdentifierOk returns a tuple with the LayerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerIdentifier

`func (o *OtelComponentMappingOutput) SetLayerIdentifier(v string)`

SetLayerIdentifier sets LayerIdentifier field to given value.

### HasLayerIdentifier

`func (o *OtelComponentMappingOutput) HasLayerIdentifier() bool`

HasLayerIdentifier returns a boolean if a field has been set.

### GetDomainName

`func (o *OtelComponentMappingOutput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *OtelComponentMappingOutput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *OtelComponentMappingOutput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetDomainIdentifier

`func (o *OtelComponentMappingOutput) GetDomainIdentifier() string`

GetDomainIdentifier returns the DomainIdentifier field if non-nil, zero value otherwise.

### GetDomainIdentifierOk

`func (o *OtelComponentMappingOutput) GetDomainIdentifierOk() (*string, bool)`

GetDomainIdentifierOk returns a tuple with the DomainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainIdentifier

`func (o *OtelComponentMappingOutput) SetDomainIdentifier(v string)`

SetDomainIdentifier sets DomainIdentifier field to given value.

### HasDomainIdentifier

`func (o *OtelComponentMappingOutput) HasDomainIdentifier() bool`

HasDomainIdentifier returns a boolean if a field has been set.

### GetRequired

`func (o *OtelComponentMappingOutput) GetRequired() OtelComponentMappingFieldMapping`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OtelComponentMappingOutput) GetRequiredOk() (*OtelComponentMappingFieldMapping, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OtelComponentMappingOutput) SetRequired(v OtelComponentMappingFieldMapping)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *OtelComponentMappingOutput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOptional

`func (o *OtelComponentMappingOutput) GetOptional() OtelComponentMappingFieldMapping`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *OtelComponentMappingOutput) GetOptionalOk() (*OtelComponentMappingFieldMapping, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *OtelComponentMappingOutput) SetOptional(v OtelComponentMappingFieldMapping)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *OtelComponentMappingOutput) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


