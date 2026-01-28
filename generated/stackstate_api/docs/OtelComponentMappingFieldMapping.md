# OtelComponentMappingFieldMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalIdentifiers** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A string containing a CEL expression within curly braces &#x60;${}&#x60;, for example \&quot;a string with a cel expression: &#x60;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60; A string with only a cel expression is also valid as long as it is within a &#x60;${}&#x60; section, for example &#x60;\&quot;${resource.attributes[&#39;service.namespace&#39;]}\&quot;&#x60;. | [optional] 
**Tags** | Pointer to [**[]OtelTagMapping**](OtelTagMapping.md) |  | [optional] 

## Methods

### NewOtelComponentMappingFieldMapping

`func NewOtelComponentMappingFieldMapping() *OtelComponentMappingFieldMapping`

NewOtelComponentMappingFieldMapping instantiates a new OtelComponentMappingFieldMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelComponentMappingFieldMappingWithDefaults

`func NewOtelComponentMappingFieldMappingWithDefaults() *OtelComponentMappingFieldMapping`

NewOtelComponentMappingFieldMappingWithDefaults instantiates a new OtelComponentMappingFieldMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalIdentifiers

`func (o *OtelComponentMappingFieldMapping) GetAdditionalIdentifiers() []string`

GetAdditionalIdentifiers returns the AdditionalIdentifiers field if non-nil, zero value otherwise.

### GetAdditionalIdentifiersOk

`func (o *OtelComponentMappingFieldMapping) GetAdditionalIdentifiersOk() (*[]string, bool)`

GetAdditionalIdentifiersOk returns a tuple with the AdditionalIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIdentifiers

`func (o *OtelComponentMappingFieldMapping) SetAdditionalIdentifiers(v []string)`

SetAdditionalIdentifiers sets AdditionalIdentifiers field to given value.

### HasAdditionalIdentifiers

`func (o *OtelComponentMappingFieldMapping) HasAdditionalIdentifiers() bool`

HasAdditionalIdentifiers returns a boolean if a field has been set.

### GetVersion

`func (o *OtelComponentMappingFieldMapping) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OtelComponentMappingFieldMapping) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OtelComponentMappingFieldMapping) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OtelComponentMappingFieldMapping) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTags

`func (o *OtelComponentMappingFieldMapping) GetTags() []OtelTagMapping`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OtelComponentMappingFieldMapping) GetTagsOk() (*[]OtelTagMapping, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OtelComponentMappingFieldMapping) SetTags(v []OtelTagMapping)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OtelComponentMappingFieldMapping) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


