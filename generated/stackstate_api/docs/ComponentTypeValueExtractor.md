# ComponentTypeValueExtractor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Sources** | [**map[string]ComponentTypeValueExtractor**](ComponentTypeValueExtractor.md) |  | 
**Value** | **string** |  | 
**Key** | **string** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**TagName** | **string** |  | 

## Methods

### NewComponentTypeValueExtractor

`func NewComponentTypeValueExtractor(type_ string, sources map[string]ComponentTypeValueExtractor, value string, key string, tagName string, ) *ComponentTypeValueExtractor`

NewComponentTypeValueExtractor instantiates a new ComponentTypeValueExtractor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTypeValueExtractorWithDefaults

`func NewComponentTypeValueExtractorWithDefaults() *ComponentTypeValueExtractor`

NewComponentTypeValueExtractorWithDefaults instantiates a new ComponentTypeValueExtractor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentTypeValueExtractor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentTypeValueExtractor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentTypeValueExtractor) SetType(v string)`

SetType sets Type field to given value.


### GetSources

`func (o *ComponentTypeValueExtractor) GetSources() map[string]ComponentTypeValueExtractor`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ComponentTypeValueExtractor) GetSourcesOk() (*map[string]ComponentTypeValueExtractor, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ComponentTypeValueExtractor) SetSources(v map[string]ComponentTypeValueExtractor)`

SetSources sets Sources field to given value.


### GetValue

`func (o *ComponentTypeValueExtractor) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentTypeValueExtractor) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentTypeValueExtractor) SetValue(v string)`

SetValue sets Value field to given value.


### GetKey

`func (o *ComponentTypeValueExtractor) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComponentTypeValueExtractor) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComponentTypeValueExtractor) SetKey(v string)`

SetKey sets Key field to given value.


### GetDefaultValue

`func (o *ComponentTypeValueExtractor) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ComponentTypeValueExtractor) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ComponentTypeValueExtractor) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ComponentTypeValueExtractor) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetTagName

`func (o *ComponentTypeValueExtractor) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ComponentTypeValueExtractor) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ComponentTypeValueExtractor) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


