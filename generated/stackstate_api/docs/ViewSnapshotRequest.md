# ViewSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | **string** | STQL query string | 
**QueryVersion** | **string** |  | 
**Metadata** | [**QueryMetadata**](QueryMetadata.md) |  | 
**ViewId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewViewSnapshotRequest

`func NewViewSnapshotRequest(type_ string, query string, queryVersion string, metadata QueryMetadata, ) *ViewSnapshotRequest`

NewViewSnapshotRequest instantiates a new ViewSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSnapshotRequestWithDefaults

`func NewViewSnapshotRequestWithDefaults() *ViewSnapshotRequest`

NewViewSnapshotRequestWithDefaults instantiates a new ViewSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ViewSnapshotRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewSnapshotRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewSnapshotRequest) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *ViewSnapshotRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ViewSnapshotRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ViewSnapshotRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryVersion

`func (o *ViewSnapshotRequest) GetQueryVersion() string`

GetQueryVersion returns the QueryVersion field if non-nil, zero value otherwise.

### GetQueryVersionOk

`func (o *ViewSnapshotRequest) GetQueryVersionOk() (*string, bool)`

GetQueryVersionOk returns a tuple with the QueryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVersion

`func (o *ViewSnapshotRequest) SetQueryVersion(v string)`

SetQueryVersion sets QueryVersion field to given value.


### GetMetadata

`func (o *ViewSnapshotRequest) GetMetadata() QueryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ViewSnapshotRequest) GetMetadataOk() (*QueryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ViewSnapshotRequest) SetMetadata(v QueryMetadata)`

SetMetadata sets Metadata field to given value.


### GetViewId

`func (o *ViewSnapshotRequest) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ViewSnapshotRequest) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ViewSnapshotRequest) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ViewSnapshotRequest) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ViewSnapshotRequest) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ViewSnapshotRequest) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


