# PromEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to [**PromData**](PromData.md) |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromEnvelope

`func NewPromEnvelope(status string, ) *PromEnvelope`

NewPromEnvelope instantiates a new PromEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromEnvelopeWithDefaults

`func NewPromEnvelopeWithDefaults() *PromEnvelope`

NewPromEnvelopeWithDefaults instantiates a new PromEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PromEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PromEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PromEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *PromEnvelope) GetData() PromData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromEnvelope) GetDataOk() (*PromData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromEnvelope) SetData(v PromData)`

SetData sets Data field to given value.

### HasData

`func (o *PromEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *PromEnvelope) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PromEnvelope) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PromEnvelope) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *PromEnvelope) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *PromEnvelope) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PromEnvelope) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PromEnvelope) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PromEnvelope) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *PromEnvelope) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PromEnvelope) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PromEnvelope) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PromEnvelope) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


