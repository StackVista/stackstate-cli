# GenericAnnotationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Annotation** | **map[string]interface{}** |  | 

## Methods

### NewGenericAnnotationData

`func NewGenericAnnotationData(type_ string, annotation map[string]interface{}, ) *GenericAnnotationData`

NewGenericAnnotationData instantiates a new GenericAnnotationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericAnnotationDataWithDefaults

`func NewGenericAnnotationDataWithDefaults() *GenericAnnotationData`

NewGenericAnnotationDataWithDefaults instantiates a new GenericAnnotationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericAnnotationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericAnnotationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericAnnotationData) SetType(v string)`

SetType sets Type field to given value.


### GetAnnotation

`func (o *GenericAnnotationData) GetAnnotation() map[string]interface{}`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *GenericAnnotationData) GetAnnotationOk() (*map[string]interface{}, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *GenericAnnotationData) SetAnnotation(v map[string]interface{})`

SetAnnotation sets Annotation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


