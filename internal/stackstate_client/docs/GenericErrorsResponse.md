# GenericErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**TrackingKey** | **string** |  | 
**ServerTimestamp** | **int64** |  | 
**Errors** | [**[]GenericApiError**](GenericApiError.md) |  | 

## Methods

### NewGenericErrorsResponse

`func NewGenericErrorsResponse(trackingKey string, serverTimestamp int64, errors []GenericApiError, ) *GenericErrorsResponse`

NewGenericErrorsResponse instantiates a new GenericErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericErrorsResponseWithDefaults

`func NewGenericErrorsResponseWithDefaults() *GenericErrorsResponse`

NewGenericErrorsResponseWithDefaults instantiates a new GenericErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericErrorsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericErrorsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericErrorsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericErrorsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTrackingKey

`func (o *GenericErrorsResponse) GetTrackingKey() string`

GetTrackingKey returns the TrackingKey field if non-nil, zero value otherwise.

### GetTrackingKeyOk

`func (o *GenericErrorsResponse) GetTrackingKeyOk() (*string, bool)`

GetTrackingKeyOk returns a tuple with the TrackingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingKey

`func (o *GenericErrorsResponse) SetTrackingKey(v string)`

SetTrackingKey sets TrackingKey field to given value.


### GetServerTimestamp

`func (o *GenericErrorsResponse) GetServerTimestamp() int64`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *GenericErrorsResponse) GetServerTimestampOk() (*int64, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *GenericErrorsResponse) SetServerTimestamp(v int64)`

SetServerTimestamp sets ServerTimestamp field to given value.


### GetErrors

`func (o *GenericErrorsResponse) GetErrors() []GenericApiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GenericErrorsResponse) GetErrorsOk() (*[]GenericApiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GenericErrorsResponse) SetErrors(v []GenericApiError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


