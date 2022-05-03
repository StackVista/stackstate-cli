# ProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StackPackVersion** | Pointer to **string** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewProvisionResponse

`func NewProvisionResponse() *ProvisionResponse`

NewProvisionResponse instantiates a new ProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionResponseWithDefaults

`func NewProvisionResponseWithDefaults() *ProvisionResponse`

NewProvisionResponseWithDefaults instantiates a new ProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisionResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *ProvisionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackPackVersion

`func (o *ProvisionResponse) GetStackPackVersion() string`

GetStackPackVersion returns the StackPackVersion field if non-nil, zero value otherwise.

### GetStackPackVersionOk

`func (o *ProvisionResponse) GetStackPackVersionOk() (*string, bool)`

GetStackPackVersionOk returns a tuple with the StackPackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPackVersion

`func (o *ProvisionResponse) SetStackPackVersion(v string)`

SetStackPackVersion sets StackPackVersion field to given value.

### HasStackPackVersion

`func (o *ProvisionResponse) HasStackPackVersion() bool`

HasStackPackVersion returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ProvisionResponse) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ProvisionResponse) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ProvisionResponse) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ProvisionResponse) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


