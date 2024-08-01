# StackElementNotFound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **string** |  | 
**Message** | **string** |  | 
**ExistedEarlierMs** | Pointer to **int64** |  | [optional] 
**ExistsLaterMs** | Pointer to **int64** |  | [optional] 

## Methods

### NewStackElementNotFound

`func NewStackElementNotFound(type_ string, objectType string, objectId string, message string, ) *StackElementNotFound`

NewStackElementNotFound instantiates a new StackElementNotFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackElementNotFoundWithDefaults

`func NewStackElementNotFoundWithDefaults() *StackElementNotFound`

NewStackElementNotFoundWithDefaults instantiates a new StackElementNotFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StackElementNotFound) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackElementNotFound) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackElementNotFound) SetType(v string)`

SetType sets Type field to given value.


### GetObjectType

`func (o *StackElementNotFound) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StackElementNotFound) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StackElementNotFound) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *StackElementNotFound) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *StackElementNotFound) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *StackElementNotFound) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetMessage

`func (o *StackElementNotFound) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StackElementNotFound) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StackElementNotFound) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExistedEarlierMs

`func (o *StackElementNotFound) GetExistedEarlierMs() int64`

GetExistedEarlierMs returns the ExistedEarlierMs field if non-nil, zero value otherwise.

### GetExistedEarlierMsOk

`func (o *StackElementNotFound) GetExistedEarlierMsOk() (*int64, bool)`

GetExistedEarlierMsOk returns a tuple with the ExistedEarlierMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistedEarlierMs

`func (o *StackElementNotFound) SetExistedEarlierMs(v int64)`

SetExistedEarlierMs sets ExistedEarlierMs field to given value.

### HasExistedEarlierMs

`func (o *StackElementNotFound) HasExistedEarlierMs() bool`

HasExistedEarlierMs returns a boolean if a field has been set.

### GetExistsLaterMs

`func (o *StackElementNotFound) GetExistsLaterMs() int64`

GetExistsLaterMs returns the ExistsLaterMs field if non-nil, zero value otherwise.

### GetExistsLaterMsOk

`func (o *StackElementNotFound) GetExistsLaterMsOk() (*int64, bool)`

GetExistsLaterMsOk returns a tuple with the ExistsLaterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsLaterMs

`func (o *StackElementNotFound) SetExistsLaterMs(v int64)`

SetExistsLaterMs sets ExistsLaterMs field to given value.

### HasExistsLaterMs

`func (o *StackElementNotFound) HasExistsLaterMs() bool`

HasExistsLaterMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


