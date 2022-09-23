# ExemplarData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesLabels** | **map[string]string** |  | 
**Exemplars** | [**[]Exemplar**](Exemplar.md) |  | 

## Methods

### NewExemplarData

`func NewExemplarData(seriesLabels map[string]string, exemplars []Exemplar, ) *ExemplarData`

NewExemplarData instantiates a new ExemplarData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExemplarDataWithDefaults

`func NewExemplarDataWithDefaults() *ExemplarData`

NewExemplarDataWithDefaults instantiates a new ExemplarData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesLabels

`func (o *ExemplarData) GetSeriesLabels() map[string]string`

GetSeriesLabels returns the SeriesLabels field if non-nil, zero value otherwise.

### GetSeriesLabelsOk

`func (o *ExemplarData) GetSeriesLabelsOk() (*map[string]string, bool)`

GetSeriesLabelsOk returns a tuple with the SeriesLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLabels

`func (o *ExemplarData) SetSeriesLabels(v map[string]string)`

SetSeriesLabels sets SeriesLabels field to given value.


### GetExemplars

`func (o *ExemplarData) GetExemplars() []Exemplar`

GetExemplars returns the Exemplars field if non-nil, zero value otherwise.

### GetExemplarsOk

`func (o *ExemplarData) GetExemplarsOk() (*[]Exemplar, bool)`

GetExemplarsOk returns a tuple with the Exemplars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemplars

`func (o *ExemplarData) SetExemplars(v []Exemplar)`

SetExemplars sets Exemplars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


