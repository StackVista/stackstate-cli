# PromBatchInstantQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Query** | **string** |  | 
**Time** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **string** |  | [optional] 
**PostFilter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromBatchInstantQuery

`func NewPromBatchInstantQuery(id string, type_ string, query string, ) *PromBatchInstantQuery`

NewPromBatchInstantQuery instantiates a new PromBatchInstantQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromBatchInstantQueryWithDefaults

`func NewPromBatchInstantQueryWithDefaults() *PromBatchInstantQuery`

NewPromBatchInstantQueryWithDefaults instantiates a new PromBatchInstantQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromBatchInstantQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromBatchInstantQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromBatchInstantQuery) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromBatchInstantQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromBatchInstantQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromBatchInstantQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *PromBatchInstantQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PromBatchInstantQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PromBatchInstantQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTime

`func (o *PromBatchInstantQuery) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PromBatchInstantQuery) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PromBatchInstantQuery) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *PromBatchInstantQuery) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStep

`func (o *PromBatchInstantQuery) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PromBatchInstantQuery) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PromBatchInstantQuery) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *PromBatchInstantQuery) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTimeout

`func (o *PromBatchInstantQuery) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PromBatchInstantQuery) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PromBatchInstantQuery) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PromBatchInstantQuery) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPostFilter

`func (o *PromBatchInstantQuery) GetPostFilter() []string`

GetPostFilter returns the PostFilter field if non-nil, zero value otherwise.

### GetPostFilterOk

`func (o *PromBatchInstantQuery) GetPostFilterOk() (*[]string, bool)`

GetPostFilterOk returns a tuple with the PostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFilter

`func (o *PromBatchInstantQuery) SetPostFilter(v []string)`

SetPostFilter sets PostFilter field to given value.

### HasPostFilter

`func (o *PromBatchInstantQuery) HasPostFilter() bool`

HasPostFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


