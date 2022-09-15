# LicensedSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Subscription** | [**Subscription**](Subscription.md) |  | 

## Methods

### NewLicensedSubscription

`func NewLicensedSubscription(type_ string, subscription Subscription, ) *LicensedSubscription`

NewLicensedSubscription instantiates a new LicensedSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensedSubscriptionWithDefaults

`func NewLicensedSubscriptionWithDefaults() *LicensedSubscription`

NewLicensedSubscriptionWithDefaults instantiates a new LicensedSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LicensedSubscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicensedSubscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicensedSubscription) SetType(v string)`

SetType sets Type field to given value.


### GetSubscription

`func (o *LicensedSubscription) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *LicensedSubscription) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *LicensedSubscription) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


