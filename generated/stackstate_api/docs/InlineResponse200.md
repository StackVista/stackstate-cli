# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**StreamSnaphots** | Pointer to [**[]LatestTelemetryStreamMetrics**](LatestTelemetryStreamMetrics.md) |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStreamSnaphots

`func (o *InlineResponse200) GetStreamSnaphots() []LatestTelemetryStreamMetrics`

GetStreamSnaphots returns the StreamSnaphots field if non-nil, zero value otherwise.

### GetStreamSnaphotsOk

`func (o *InlineResponse200) GetStreamSnaphotsOk() (*[]LatestTelemetryStreamMetrics, bool)`

GetStreamSnaphotsOk returns a tuple with the StreamSnaphots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSnaphots

`func (o *InlineResponse200) SetStreamSnaphots(v []LatestTelemetryStreamMetrics)`

SetStreamSnaphots sets StreamSnaphots field to given value.

### HasStreamSnaphots

`func (o *InlineResponse200) HasStreamSnaphots() bool`

HasStreamSnaphots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


