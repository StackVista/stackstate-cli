# SnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | **string** | STQL query string | 
**QueryVersion** | **string** |  | 
**Metadata** | [**QueryMetadata**](QueryMetadata.md) |  | 

## Methods

### NewSnapshotRequest

`func NewSnapshotRequest(type_ string, query string, queryVersion string, metadata QueryMetadata, ) *SnapshotRequest`

NewSnapshotRequest instantiates a new SnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRequestWithDefaults

`func NewSnapshotRequestWithDefaults() *SnapshotRequest`

NewSnapshotRequestWithDefaults instantiates a new SnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SnapshotRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotRequest) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *SnapshotRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SnapshotRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SnapshotRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryVersion

`func (o *SnapshotRequest) GetQueryVersion() string`

GetQueryVersion returns the QueryVersion field if non-nil, zero value otherwise.

### GetQueryVersionOk

`func (o *SnapshotRequest) GetQueryVersionOk() (*string, bool)`

GetQueryVersionOk returns a tuple with the QueryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVersion

`func (o *SnapshotRequest) SetQueryVersion(v string)`

SetQueryVersion sets QueryVersion field to given value.


### GetMetadata

`func (o *SnapshotRequest) GetMetadata() QueryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnapshotRequest) GetMetadataOk() (*QueryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnapshotRequest) SetMetadata(v QueryMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


