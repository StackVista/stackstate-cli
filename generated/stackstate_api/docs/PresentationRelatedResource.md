# PresentationRelatedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | Stable identity key for merging across presentations, analogous to fieldId and filterId. | 
**Title** | **string** | Section heading displayed in the UI for this related resource. | 
**Order** | **float64** | Display order. Higher value means it shows first in UI. | 
**TopologyQuery** | Pointer to **string** | STQL query scoping the related components. Supports template interpolation with ${*} placeholders (e.g. ${identifiers[0]}, ${tags[&#39;namespace&#39;]}). The backend intersects this query with the referenced presentation&#39;s binding query.  | [optional] 
**PresentationIdentifier** | Pointer to **string** | References a ComponentPresentation by identifier to reuse its overview spec (columns, name, filters) for rendering this related resource section. Must be within the same stackpack.  | [optional] 

## Methods

### NewPresentationRelatedResource

`func NewPresentationRelatedResource(resourceId string, title string, order float64, ) *PresentationRelatedResource`

NewPresentationRelatedResource instantiates a new PresentationRelatedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationRelatedResourceWithDefaults

`func NewPresentationRelatedResourceWithDefaults() *PresentationRelatedResource`

NewPresentationRelatedResourceWithDefaults instantiates a new PresentationRelatedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *PresentationRelatedResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PresentationRelatedResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PresentationRelatedResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetTitle

`func (o *PresentationRelatedResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationRelatedResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationRelatedResource) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOrder

`func (o *PresentationRelatedResource) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationRelatedResource) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationRelatedResource) SetOrder(v float64)`

SetOrder sets Order field to given value.


### GetTopologyQuery

`func (o *PresentationRelatedResource) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *PresentationRelatedResource) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *PresentationRelatedResource) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.

### HasTopologyQuery

`func (o *PresentationRelatedResource) HasTopologyQuery() bool`

HasTopologyQuery returns a boolean if a field has been set.

### GetPresentationIdentifier

`func (o *PresentationRelatedResource) GetPresentationIdentifier() string`

GetPresentationIdentifier returns the PresentationIdentifier field if non-nil, zero value otherwise.

### GetPresentationIdentifierOk

`func (o *PresentationRelatedResource) GetPresentationIdentifierOk() (*string, bool)`

GetPresentationIdentifierOk returns a tuple with the PresentationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentationIdentifier

`func (o *PresentationRelatedResource) SetPresentationIdentifier(v string)`

SetPresentationIdentifier sets PresentationIdentifier field to given value.

### HasPresentationIdentifier

`func (o *PresentationRelatedResource) HasPresentationIdentifier() bool`

HasPresentationIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


