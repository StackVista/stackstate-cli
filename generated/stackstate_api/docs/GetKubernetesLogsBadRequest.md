# GetKubernetesLogsBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**From** | **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | 
**To** | **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | 
**PageSize** | **int32** |  | 
**Page** | **int32** |  | 
**Query** | **string** |  | 

## Methods

### NewGetKubernetesLogsBadRequest

`func NewGetKubernetesLogsBadRequest(type_ string, message string, from int32, to int32, pageSize int32, page int32, query string, ) *GetKubernetesLogsBadRequest`

NewGetKubernetesLogsBadRequest instantiates a new GetKubernetesLogsBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesLogsBadRequestWithDefaults

`func NewGetKubernetesLogsBadRequestWithDefaults() *GetKubernetesLogsBadRequest`

NewGetKubernetesLogsBadRequestWithDefaults instantiates a new GetKubernetesLogsBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetKubernetesLogsBadRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetKubernetesLogsBadRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetKubernetesLogsBadRequest) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *GetKubernetesLogsBadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetKubernetesLogsBadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetKubernetesLogsBadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFrom

`func (o *GetKubernetesLogsBadRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetKubernetesLogsBadRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetKubernetesLogsBadRequest) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *GetKubernetesLogsBadRequest) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetKubernetesLogsBadRequest) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetKubernetesLogsBadRequest) SetTo(v int32)`

SetTo sets To field to given value.


### GetPageSize

`func (o *GetKubernetesLogsBadRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetKubernetesLogsBadRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetKubernetesLogsBadRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetKubernetesLogsBadRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetKubernetesLogsBadRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetKubernetesLogsBadRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetQuery

`func (o *GetKubernetesLogsBadRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetKubernetesLogsBadRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetKubernetesLogsBadRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


