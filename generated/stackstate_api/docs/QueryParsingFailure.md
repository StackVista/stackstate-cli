# QueryParsingFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field | 
**Reason** | **string** | Error message describing why the query failed to parse | 

## Methods

### NewQueryParsingFailure

`func NewQueryParsingFailure(type_ string, reason string, ) *QueryParsingFailure`

NewQueryParsingFailure instantiates a new QueryParsingFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParsingFailureWithDefaults

`func NewQueryParsingFailureWithDefaults() *QueryParsingFailure`

NewQueryParsingFailureWithDefaults instantiates a new QueryParsingFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryParsingFailure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryParsingFailure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryParsingFailure) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *QueryParsingFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QueryParsingFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QueryParsingFailure) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


