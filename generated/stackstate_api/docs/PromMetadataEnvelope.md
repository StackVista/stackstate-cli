# PromMetadataEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to [**map[string][]map[string]string**](array.md) |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromMetadataEnvelope

`func NewPromMetadataEnvelope(status string, ) *PromMetadataEnvelope`

NewPromMetadataEnvelope instantiates a new PromMetadataEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromMetadataEnvelopeWithDefaults

`func NewPromMetadataEnvelopeWithDefaults() *PromMetadataEnvelope`

NewPromMetadataEnvelopeWithDefaults instantiates a new PromMetadataEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PromMetadataEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PromMetadataEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PromMetadataEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *PromMetadataEnvelope) GetData() map[string][]map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromMetadataEnvelope) GetDataOk() (*map[string][]map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromMetadataEnvelope) SetData(v map[string][]map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PromMetadataEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *PromMetadataEnvelope) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PromMetadataEnvelope) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PromMetadataEnvelope) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *PromMetadataEnvelope) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *PromMetadataEnvelope) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PromMetadataEnvelope) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PromMetadataEnvelope) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PromMetadataEnvelope) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *PromMetadataEnvelope) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PromMetadataEnvelope) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PromMetadataEnvelope) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PromMetadataEnvelope) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


