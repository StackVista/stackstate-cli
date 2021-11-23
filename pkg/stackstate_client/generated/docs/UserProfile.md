# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**StarredViews** | Pointer to **[]int64** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile(name string, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *UserProfile) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *UserProfile) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *UserProfile) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *UserProfile) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *UserProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UserProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *UserProfile) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserProfile) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserProfile) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UserProfile) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetStarredViews

`func (o *UserProfile) GetStarredViews() []int64`

GetStarredViews returns the StarredViews field if non-nil, zero value otherwise.

### GetStarredViewsOk

`func (o *UserProfile) GetStarredViewsOk() (*[]int64, bool)`

GetStarredViewsOk returns a tuple with the StarredViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredViews

`func (o *UserProfile) SetStarredViews(v []int64)`

SetStarredViews sets StarredViews field to given value.

### HasStarredViews

`func (o *UserProfile) HasStarredViews() bool`

HasStarredViews returns a boolean if a field has been set.

### GetOwnedBy

`func (o *UserProfile) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *UserProfile) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *UserProfile) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *UserProfile) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


