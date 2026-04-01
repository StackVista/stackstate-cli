# OtelComponentMappingFieldMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalIdentifiers** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | An expression that must produce a string. It must be one of these formats:   - A plain string, for example &#x60;\&quot;this is a plain string\&quot;&#x60;   - A cel expression that must return a string, for example: &#x60;resource.attributes[&#39;service.namespace&#39;]&#x60; | [optional] 
**Tags** | Pointer to [**[]OtelTagMapping**](OtelTagMapping.md) |  | [optional] 
**Configuration** | Pointer to **string** | An expression that can produce any type.  Variables use it to store any type of value. For example, to store a boolean in a variable named  &#x60;inTestNamespace&#x60; assign it the expression &#x60;resource.attributes[&#39;service.namespace&#39;] &#x3D;&#x3D; &#39;test&#39;&#x60;. The variable can now be used directly in the conditions like this: &#x60;vars.inTestNamespace&#x60;. | [optional] 
**Status** | Pointer to **string** | An expression that can produce any type.  Variables use it to store any type of value. For example, to store a boolean in a variable named  &#x60;inTestNamespace&#x60; assign it the expression &#x60;resource.attributes[&#39;service.namespace&#39;] &#x3D;&#x3D; &#39;test&#39;&#x60;. The variable can now be used directly in the conditions like this: &#x60;vars.inTestNamespace&#x60;. | [optional] 

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

### GetConfiguration

`func (o *OtelComponentMappingFieldMapping) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OtelComponentMappingFieldMapping) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OtelComponentMappingFieldMapping) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OtelComponentMappingFieldMapping) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetStatus

`func (o *OtelComponentMappingFieldMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OtelComponentMappingFieldMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OtelComponentMappingFieldMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OtelComponentMappingFieldMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


