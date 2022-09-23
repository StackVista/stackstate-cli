# PromExemplarData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesLabels** | **map[string]string** |  | 
**Exemplars** | [**[]PromExemplar**](PromExemplar.md) |  | 

## Methods

### NewPromExemplarData

`func NewPromExemplarData(seriesLabels map[string]string, exemplars []PromExemplar, ) *PromExemplarData`

NewPromExemplarData instantiates a new PromExemplarData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromExemplarDataWithDefaults

`func NewPromExemplarDataWithDefaults() *PromExemplarData`

NewPromExemplarDataWithDefaults instantiates a new PromExemplarData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesLabels

`func (o *PromExemplarData) GetSeriesLabels() map[string]string`

GetSeriesLabels returns the SeriesLabels field if non-nil, zero value otherwise.

### GetSeriesLabelsOk

`func (o *PromExemplarData) GetSeriesLabelsOk() (*map[string]string, bool)`

GetSeriesLabelsOk returns a tuple with the SeriesLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLabels

`func (o *PromExemplarData) SetSeriesLabels(v map[string]string)`

SetSeriesLabels sets SeriesLabels field to given value.


### GetExemplars

`func (o *PromExemplarData) GetExemplars() []PromExemplar`

GetExemplars returns the Exemplars field if non-nil, zero value otherwise.

### GetExemplarsOk

`func (o *PromExemplarData) GetExemplarsOk() (*[]PromExemplar, bool)`

GetExemplarsOk returns a tuple with the Exemplars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemplars

`func (o *PromExemplarData) SetExemplars(v []PromExemplar)`

SetExemplars sets Exemplars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


