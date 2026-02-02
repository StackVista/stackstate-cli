# QuerySnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field | 
**ViewSnapshotResponse** | **map[string]interface{}** | The query result or error response. This is opaque at the API level but contains one of: - ViewSnapshot (success) - ViewSnapshotFetchTimeout (error) - ViewSnapshotTooManyActiveQueries (error) - ViewSnapshotTopologySizeOverflow (error) - ViewSnapshotDataUnavailable (error) Use the _type field to discriminate between types.  | 
**ViewId** | Pointer to **NullableInt64** | Optional view ID associated with this result | [optional] 

## Methods

### NewQuerySnapshotResult

`func NewQuerySnapshotResult(type_ string, viewSnapshotResponse map[string]interface{}, ) *QuerySnapshotResult`

NewQuerySnapshotResult instantiates a new QuerySnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySnapshotResultWithDefaults

`func NewQuerySnapshotResultWithDefaults() *QuerySnapshotResult`

NewQuerySnapshotResultWithDefaults instantiates a new QuerySnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QuerySnapshotResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuerySnapshotResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuerySnapshotResult) SetType(v string)`

SetType sets Type field to given value.


### GetViewSnapshotResponse

`func (o *QuerySnapshotResult) GetViewSnapshotResponse() map[string]interface{}`

GetViewSnapshotResponse returns the ViewSnapshotResponse field if non-nil, zero value otherwise.

### GetViewSnapshotResponseOk

`func (o *QuerySnapshotResult) GetViewSnapshotResponseOk() (*map[string]interface{}, bool)`

GetViewSnapshotResponseOk returns a tuple with the ViewSnapshotResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSnapshotResponse

`func (o *QuerySnapshotResult) SetViewSnapshotResponse(v map[string]interface{})`

SetViewSnapshotResponse sets ViewSnapshotResponse field to given value.


### GetViewId

`func (o *QuerySnapshotResult) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *QuerySnapshotResult) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *QuerySnapshotResult) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *QuerySnapshotResult) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *QuerySnapshotResult) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *QuerySnapshotResult) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


