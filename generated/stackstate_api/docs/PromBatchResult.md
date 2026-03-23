# PromBatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ResultType** | **string** |  | 
**Data** | [**PromEnvelope**](PromEnvelope.md) |  | 
**Error** | **string** |  | 

## Methods

### NewPromBatchResult

`func NewPromBatchResult(id string, resultType string, data PromEnvelope, error_ string, ) *PromBatchResult`

NewPromBatchResult instantiates a new PromBatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromBatchResultWithDefaults

`func NewPromBatchResultWithDefaults() *PromBatchResult`

NewPromBatchResultWithDefaults instantiates a new PromBatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromBatchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromBatchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromBatchResult) SetId(v string)`

SetId sets Id field to given value.


### GetResultType

`func (o *PromBatchResult) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *PromBatchResult) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *PromBatchResult) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetData

`func (o *PromBatchResult) GetData() PromEnvelope`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromBatchResult) GetDataOk() (*PromEnvelope, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromBatchResult) SetData(v PromEnvelope)`

SetData sets Data field to given value.


### GetError

`func (o *PromBatchResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PromBatchResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PromBatchResult) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


