# WebhookChannelWriteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Token** | **string** |  | 
**VerifySsl** | **bool** |  | 
**Metadata** | **map[string]string** |  | 

## Methods

### NewWebhookChannelWriteSchema

`func NewWebhookChannelWriteSchema(url string, token string, verifySsl bool, metadata map[string]string, ) *WebhookChannelWriteSchema`

NewWebhookChannelWriteSchema instantiates a new WebhookChannelWriteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookChannelWriteSchemaWithDefaults

`func NewWebhookChannelWriteSchemaWithDefaults() *WebhookChannelWriteSchema`

NewWebhookChannelWriteSchemaWithDefaults instantiates a new WebhookChannelWriteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookChannelWriteSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookChannelWriteSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookChannelWriteSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *WebhookChannelWriteSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebhookChannelWriteSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebhookChannelWriteSchema) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerifySsl

`func (o *WebhookChannelWriteSchema) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *WebhookChannelWriteSchema) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *WebhookChannelWriteSchema) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.


### GetMetadata

`func (o *WebhookChannelWriteSchema) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookChannelWriteSchema) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookChannelWriteSchema) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


