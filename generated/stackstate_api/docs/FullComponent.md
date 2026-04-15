# FullComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**Fields** | [**[]ComponentField**](ComponentField.md) |  | 
**Synced** | [**[]ExternalComponent**](ExternalComponent.md) |  | 
**Provisioning** | Pointer to [**ComponentProvisioning**](ComponentProvisioning.md) |  | [optional] 
**RelatedResources** | [**[]RelatedResource**](RelatedResource.md) | Resolved related resource definitions in display order. Backend populates from both legacy and new presentation definitions. | 
**Data** | [**ComponentData**](ComponentData.md) |  | 
**Highlights** | Pointer to [**LegacyComponentHighlights**](LegacyComponentHighlights.md) |  | [optional] 
**Actions** | [**[]ComponentAction**](ComponentAction.md) |  | 
**BoundMetrics** | [**[]BoundMetric**](BoundMetric.md) |  | 
**BoundTraces** | Pointer to [**BoundTraces**](BoundTraces.md) |  | [optional] 

## Methods

### NewFullComponent

`func NewFullComponent(typeName string, fields []ComponentField, synced []ExternalComponent, relatedResources []RelatedResource, data ComponentData, actions []ComponentAction, boundMetrics []BoundMetric, ) *FullComponent`

NewFullComponent instantiates a new FullComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullComponentWithDefaults

`func NewFullComponentWithDefaults() *FullComponent`

NewFullComponentWithDefaults instantiates a new FullComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *FullComponent) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FullComponent) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FullComponent) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetIcon

`func (o *FullComponent) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FullComponent) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FullComponent) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FullComponent) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetFields

`func (o *FullComponent) GetFields() []ComponentField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FullComponent) GetFieldsOk() (*[]ComponentField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FullComponent) SetFields(v []ComponentField)`

SetFields sets Fields field to given value.


### GetSynced

`func (o *FullComponent) GetSynced() []ExternalComponent`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *FullComponent) GetSyncedOk() (*[]ExternalComponent, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *FullComponent) SetSynced(v []ExternalComponent)`

SetSynced sets Synced field to given value.


### GetProvisioning

`func (o *FullComponent) GetProvisioning() ComponentProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *FullComponent) GetProvisioningOk() (*ComponentProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *FullComponent) SetProvisioning(v ComponentProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *FullComponent) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRelatedResources

`func (o *FullComponent) GetRelatedResources() []RelatedResource`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *FullComponent) GetRelatedResourcesOk() (*[]RelatedResource, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *FullComponent) SetRelatedResources(v []RelatedResource)`

SetRelatedResources sets RelatedResources field to given value.


### GetData

`func (o *FullComponent) GetData() ComponentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FullComponent) GetDataOk() (*ComponentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FullComponent) SetData(v ComponentData)`

SetData sets Data field to given value.


### GetHighlights

`func (o *FullComponent) GetHighlights() LegacyComponentHighlights`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *FullComponent) GetHighlightsOk() (*LegacyComponentHighlights, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *FullComponent) SetHighlights(v LegacyComponentHighlights)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *FullComponent) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetActions

`func (o *FullComponent) GetActions() []ComponentAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FullComponent) GetActionsOk() (*[]ComponentAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FullComponent) SetActions(v []ComponentAction)`

SetActions sets Actions field to given value.


### GetBoundMetrics

`func (o *FullComponent) GetBoundMetrics() []BoundMetric`

GetBoundMetrics returns the BoundMetrics field if non-nil, zero value otherwise.

### GetBoundMetricsOk

`func (o *FullComponent) GetBoundMetricsOk() (*[]BoundMetric, bool)`

GetBoundMetricsOk returns a tuple with the BoundMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundMetrics

`func (o *FullComponent) SetBoundMetrics(v []BoundMetric)`

SetBoundMetrics sets BoundMetrics field to given value.


### GetBoundTraces

`func (o *FullComponent) GetBoundTraces() BoundTraces`

GetBoundTraces returns the BoundTraces field if non-nil, zero value otherwise.

### GetBoundTracesOk

`func (o *FullComponent) GetBoundTracesOk() (*BoundTraces, bool)`

GetBoundTracesOk returns a tuple with the BoundTraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTraces

`func (o *FullComponent) SetBoundTraces(v BoundTraces)`

SetBoundTraces sets BoundTraces field to given value.

### HasBoundTraces

`func (o *FullComponent) HasBoundTraces() bool`

HasBoundTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


