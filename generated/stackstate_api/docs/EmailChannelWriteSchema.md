# EmailChannelWriteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** |  | 
**Cc** | **[]string** |  | 
**SubjectPrefix** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailChannelWriteSchema

`func NewEmailChannelWriteSchema(to []string, cc []string, ) *EmailChannelWriteSchema`

NewEmailChannelWriteSchema instantiates a new EmailChannelWriteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailChannelWriteSchemaWithDefaults

`func NewEmailChannelWriteSchemaWithDefaults() *EmailChannelWriteSchema`

NewEmailChannelWriteSchemaWithDefaults instantiates a new EmailChannelWriteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *EmailChannelWriteSchema) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailChannelWriteSchema) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailChannelWriteSchema) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *EmailChannelWriteSchema) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailChannelWriteSchema) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailChannelWriteSchema) SetCc(v []string)`

SetCc sets Cc field to given value.


### GetSubjectPrefix

`func (o *EmailChannelWriteSchema) GetSubjectPrefix() string`

GetSubjectPrefix returns the SubjectPrefix field if non-nil, zero value otherwise.

### GetSubjectPrefixOk

`func (o *EmailChannelWriteSchema) GetSubjectPrefixOk() (*string, bool)`

GetSubjectPrefixOk returns a tuple with the SubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPrefix

`func (o *EmailChannelWriteSchema) SetSubjectPrefix(v string)`

SetSubjectPrefix sets SubjectPrefix field to given value.

### HasSubjectPrefix

`func (o *EmailChannelWriteSchema) HasSubjectPrefix() bool`

HasSubjectPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


