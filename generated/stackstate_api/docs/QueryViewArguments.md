# QueryViewArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | **string** |  | 
**QueryTime** | Pointer to **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | [optional] 

## Methods

### NewQueryViewArguments

`func NewQueryViewArguments(type_ string, query string, ) *QueryViewArguments`

NewQueryViewArguments instantiates a new QueryViewArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryViewArgumentsWithDefaults

`func NewQueryViewArgumentsWithDefaults() *QueryViewArguments`

NewQueryViewArgumentsWithDefaults instantiates a new QueryViewArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryViewArguments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryViewArguments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryViewArguments) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *QueryViewArguments) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryViewArguments) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryViewArguments) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryTime

`func (o *QueryViewArguments) GetQueryTime() int32`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *QueryViewArguments) GetQueryTimeOk() (*int32, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *QueryViewArguments) SetQueryTime(v int32)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *QueryViewArguments) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


