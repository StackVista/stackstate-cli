# CompositeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Sources** | [**map[string]ComponentTypeValueExtractor**](ComponentTypeValueExtractor.md) |  | 

## Methods

### NewCompositeSource

`func NewCompositeSource(type_ string, sources map[string]ComponentTypeValueExtractor, ) *CompositeSource`

NewCompositeSource instantiates a new CompositeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeSourceWithDefaults

`func NewCompositeSourceWithDefaults() *CompositeSource`

NewCompositeSourceWithDefaults instantiates a new CompositeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompositeSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompositeSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompositeSource) SetType(v string)`

SetType sets Type field to given value.


### GetSources

`func (o *CompositeSource) GetSources() map[string]ComponentTypeValueExtractor`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CompositeSource) GetSourcesOk() (*map[string]ComponentTypeValueExtractor, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CompositeSource) SetSources(v map[string]ComponentTypeValueExtractor)`

SetSources sets Sources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


