# Envelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to [**Data**](Data.md) |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnvelope

`func NewEnvelope(status string, ) *Envelope`

NewEnvelope instantiates a new Envelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopeWithDefaults

`func NewEnvelopeWithDefaults() *Envelope`

NewEnvelopeWithDefaults instantiates a new Envelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Envelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Envelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Envelope) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *Envelope) GetData() Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Envelope) GetDataOk() (*Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Envelope) SetData(v Data)`

SetData sets Data field to given value.

### HasData

`func (o *Envelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *Envelope) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *Envelope) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *Envelope) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *Envelope) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *Envelope) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Envelope) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Envelope) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Envelope) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *Envelope) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Envelope) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Envelope) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Envelope) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


