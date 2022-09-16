# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | **string** |  | 
**ExpiryTimestampMs** | Pointer to **int64** |  | [optional] 
**Plan** | **string** |  | 

## Methods

### NewSubscription

`func NewSubscription(tenant string, plan string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *Subscription) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Subscription) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Subscription) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetExpiryTimestampMs

`func (o *Subscription) GetExpiryTimestampMs() int64`

GetExpiryTimestampMs returns the ExpiryTimestampMs field if non-nil, zero value otherwise.

### GetExpiryTimestampMsOk

`func (o *Subscription) GetExpiryTimestampMsOk() (*int64, bool)`

GetExpiryTimestampMsOk returns a tuple with the ExpiryTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestampMs

`func (o *Subscription) SetExpiryTimestampMs(v int64)`

SetExpiryTimestampMs sets ExpiryTimestampMs field to given value.

### HasExpiryTimestampMs

`func (o *Subscription) HasExpiryTimestampMs() bool`

HasExpiryTimestampMs returns a boolean if a field has been set.

### GetPlan

`func (o *Subscription) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscription) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscription) SetPlan(v string)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


