# DashboardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | [**[]DashboardReadSchema**](DashboardReadSchema.md) | List of dashboards, which can be either metadata or full representations. | 

## Methods

### NewDashboardList

`func NewDashboardList(dashboards []DashboardReadSchema, ) *DashboardList`

NewDashboardList instantiates a new DashboardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListWithDefaults

`func NewDashboardListWithDefaults() *DashboardList`

NewDashboardListWithDefaults instantiates a new DashboardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *DashboardList) GetDashboards() []DashboardReadSchema`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardList) GetDashboardsOk() (*[]DashboardReadSchema, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardList) SetDashboards(v []DashboardReadSchema)`

SetDashboards sets Dashboards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


