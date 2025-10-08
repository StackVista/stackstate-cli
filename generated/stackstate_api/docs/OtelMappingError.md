# OtelMappingError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**MessageLevel**](MessageLevel.md) |  | 
**Message** | **string** |  | 
**IssueId** | Pointer to **string** |  | [optional] 

## Methods

### NewOtelMappingError

`func NewOtelMappingError(level MessageLevel, message string, ) *OtelMappingError`

NewOtelMappingError instantiates a new OtelMappingError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelMappingErrorWithDefaults

`func NewOtelMappingErrorWithDefaults() *OtelMappingError`

NewOtelMappingErrorWithDefaults instantiates a new OtelMappingError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *OtelMappingError) GetLevel() MessageLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OtelMappingError) GetLevelOk() (*MessageLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *OtelMappingError) SetLevel(v MessageLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *OtelMappingError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OtelMappingError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OtelMappingError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetIssueId

`func (o *OtelMappingError) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *OtelMappingError) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *OtelMappingError) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *OtelMappingError) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


