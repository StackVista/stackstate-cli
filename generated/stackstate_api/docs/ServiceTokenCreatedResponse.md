# ServiceTokenCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**Token** | **string** |  | [readonly] 
**Expiration** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewServiceTokenCreatedResponse

`func NewServiceTokenCreatedResponse(name string, token string, ) *ServiceTokenCreatedResponse`

NewServiceTokenCreatedResponse instantiates a new ServiceTokenCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenCreatedResponseWithDefaults

`func NewServiceTokenCreatedResponseWithDefaults() *ServiceTokenCreatedResponse`

NewServiceTokenCreatedResponseWithDefaults instantiates a new ServiceTokenCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceTokenCreatedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTokenCreatedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTokenCreatedResponse) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *ServiceTokenCreatedResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ServiceTokenCreatedResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ServiceTokenCreatedResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiration

`func (o *ServiceTokenCreatedResponse) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ServiceTokenCreatedResponse) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ServiceTokenCreatedResponse) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ServiceTokenCreatedResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


