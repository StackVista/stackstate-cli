# OtelComponentMappingRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specificity** | **float64** | Determines how much of a \&quot;specialization\&quot; this mapping is. Higher number means more specific. Used during component merge to pick the winning typeName / name when multiple mappings contribute to the same component identifier. Recommended ranges: 1–99 environment &amp; infrastructure; 100–199 platform &amp; orchestration (k8s, otel base); 200–299 application / service; 300–399 runtime / SDK / language; 400+ user overrides. | 

## Methods

### NewOtelComponentMappingRank

`func NewOtelComponentMappingRank(specificity float64, ) *OtelComponentMappingRank`

NewOtelComponentMappingRank instantiates a new OtelComponentMappingRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelComponentMappingRankWithDefaults

`func NewOtelComponentMappingRankWithDefaults() *OtelComponentMappingRank`

NewOtelComponentMappingRankWithDefaults instantiates a new OtelComponentMappingRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecificity

`func (o *OtelComponentMappingRank) GetSpecificity() float64`

GetSpecificity returns the Specificity field if non-nil, zero value otherwise.

### GetSpecificityOk

`func (o *OtelComponentMappingRank) GetSpecificityOk() (*float64, bool)`

GetSpecificityOk returns a tuple with the Specificity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificity

`func (o *OtelComponentMappingRank) SetSpecificity(v float64)`

SetSpecificity sets Specificity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


