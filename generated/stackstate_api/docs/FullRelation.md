# FullRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RelationData**](RelationData.md) |  | 
**Source** | [**FullComponent**](FullComponent.md) |  | 
**Target** | [**FullComponent**](FullComponent.md) |  | 
**TypeName** | **string** |  | 

## Methods

### NewFullRelation

`func NewFullRelation(data RelationData, source FullComponent, target FullComponent, typeName string, ) *FullRelation`

NewFullRelation instantiates a new FullRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullRelationWithDefaults

`func NewFullRelationWithDefaults() *FullRelation`

NewFullRelationWithDefaults instantiates a new FullRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FullRelation) GetData() RelationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FullRelation) GetDataOk() (*RelationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FullRelation) SetData(v RelationData)`

SetData sets Data field to given value.


### GetSource

`func (o *FullRelation) GetSource() FullComponent`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FullRelation) GetSourceOk() (*FullComponent, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FullRelation) SetSource(v FullComponent)`

SetSource sets Source field to given value.


### GetTarget

`func (o *FullRelation) GetTarget() FullComponent`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FullRelation) GetTargetOk() (*FullComponent, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FullRelation) SetTarget(v FullComponent)`

SetTarget sets Target field to given value.


### GetTypeName

`func (o *FullRelation) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FullRelation) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FullRelation) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


