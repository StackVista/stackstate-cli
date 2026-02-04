# QuerySnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field | 
**SnapshotResponse** | **map[string]interface{}** | The query result or error response. This is opaque at the API level but contains one of the following DTO types: - ViewSnapshot (success) - ViewSnapshotFetchTimeout (error) - ViewSnapshotTooManyActiveQueries (error) - ViewSnapshotTopologySizeOverflow (error) - ViewSnapshotDataUnavailable (error) Use the _type field to discriminate between types.  | 

## Methods

### NewQuerySnapshotResult

`func NewQuerySnapshotResult(type_ string, snapshotResponse map[string]interface{}, ) *QuerySnapshotResult`

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


### GetSnapshotResponse

`func (o *QuerySnapshotResult) GetSnapshotResponse() map[string]interface{}`

GetSnapshotResponse returns the SnapshotResponse field if non-nil, zero value otherwise.

### GetSnapshotResponseOk

`func (o *QuerySnapshotResult) GetSnapshotResponseOk() (*map[string]interface{}, bool)`

GetSnapshotResponseOk returns a tuple with the SnapshotResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotResponse

`func (o *QuerySnapshotResult) SetSnapshotResponse(v map[string]interface{})`

SetSnapshotResponse sets SnapshotResponse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


