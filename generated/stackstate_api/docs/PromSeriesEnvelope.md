# PromSeriesEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to **[]map[string]string** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromSeriesEnvelope

`func NewPromSeriesEnvelope(status string, ) *PromSeriesEnvelope`

NewPromSeriesEnvelope instantiates a new PromSeriesEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromSeriesEnvelopeWithDefaults

`func NewPromSeriesEnvelopeWithDefaults() *PromSeriesEnvelope`

NewPromSeriesEnvelopeWithDefaults instantiates a new PromSeriesEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PromSeriesEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PromSeriesEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PromSeriesEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *PromSeriesEnvelope) GetData() []map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromSeriesEnvelope) GetDataOk() (*[]map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromSeriesEnvelope) SetData(v []map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PromSeriesEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *PromSeriesEnvelope) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PromSeriesEnvelope) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PromSeriesEnvelope) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *PromSeriesEnvelope) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *PromSeriesEnvelope) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PromSeriesEnvelope) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PromSeriesEnvelope) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PromSeriesEnvelope) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *PromSeriesEnvelope) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PromSeriesEnvelope) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PromSeriesEnvelope) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PromSeriesEnvelope) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


