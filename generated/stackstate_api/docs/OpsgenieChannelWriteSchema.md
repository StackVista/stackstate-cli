# OpsgenieChannelWriteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**OpsgenieRegion**](OpsgenieRegion.md) |  | 
**GenieKey** | **string** |  | 
**Responders** | [**[]OpsgenieResponder**](OpsgenieResponder.md) |  | 
**Priority** | [**OpsgeniePriority**](OpsgeniePriority.md) |  | 

## Methods

### NewOpsgenieChannelWriteSchema

`func NewOpsgenieChannelWriteSchema(region OpsgenieRegion, genieKey string, responders []OpsgenieResponder, priority OpsgeniePriority, ) *OpsgenieChannelWriteSchema`

NewOpsgenieChannelWriteSchema instantiates a new OpsgenieChannelWriteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsgenieChannelWriteSchemaWithDefaults

`func NewOpsgenieChannelWriteSchemaWithDefaults() *OpsgenieChannelWriteSchema`

NewOpsgenieChannelWriteSchemaWithDefaults instantiates a new OpsgenieChannelWriteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *OpsgenieChannelWriteSchema) GetRegion() OpsgenieRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OpsgenieChannelWriteSchema) GetRegionOk() (*OpsgenieRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OpsgenieChannelWriteSchema) SetRegion(v OpsgenieRegion)`

SetRegion sets Region field to given value.


### GetGenieKey

`func (o *OpsgenieChannelWriteSchema) GetGenieKey() string`

GetGenieKey returns the GenieKey field if non-nil, zero value otherwise.

### GetGenieKeyOk

`func (o *OpsgenieChannelWriteSchema) GetGenieKeyOk() (*string, bool)`

GetGenieKeyOk returns a tuple with the GenieKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenieKey

`func (o *OpsgenieChannelWriteSchema) SetGenieKey(v string)`

SetGenieKey sets GenieKey field to given value.


### GetResponders

`func (o *OpsgenieChannelWriteSchema) GetResponders() []OpsgenieResponder`

GetResponders returns the Responders field if non-nil, zero value otherwise.

### GetRespondersOk

`func (o *OpsgenieChannelWriteSchema) GetRespondersOk() (*[]OpsgenieResponder, bool)`

GetRespondersOk returns a tuple with the Responders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponders

`func (o *OpsgenieChannelWriteSchema) SetResponders(v []OpsgenieResponder)`

SetResponders sets Responders field to given value.


### GetPriority

`func (o *OpsgenieChannelWriteSchema) GetPriority() OpsgeniePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OpsgenieChannelWriteSchema) GetPriorityOk() (*OpsgeniePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OpsgenieChannelWriteSchema) SetPriority(v OpsgeniePriority)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


