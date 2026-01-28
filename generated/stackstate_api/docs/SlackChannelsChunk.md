# SlackChannelsChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**[]SlackChannel**](SlackChannel.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewSlackChannelsChunk

`func NewSlackChannelsChunk(channels []SlackChannel, ) *SlackChannelsChunk`

NewSlackChannelsChunk instantiates a new SlackChannelsChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackChannelsChunkWithDefaults

`func NewSlackChannelsChunkWithDefaults() *SlackChannelsChunk`

NewSlackChannelsChunkWithDefaults instantiates a new SlackChannelsChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *SlackChannelsChunk) GetChannels() []SlackChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *SlackChannelsChunk) GetChannelsOk() (*[]SlackChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *SlackChannelsChunk) SetChannels(v []SlackChannel)`

SetChannels sets Channels field to given value.


### GetNextCursor

`func (o *SlackChannelsChunk) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SlackChannelsChunk) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SlackChannelsChunk) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *SlackChannelsChunk) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


