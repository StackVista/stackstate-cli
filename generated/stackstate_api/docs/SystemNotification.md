# SystemNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | **string** |  | 
**Title** | **string** |  | 
**Severity** | [**SystemNotificationSeverity**](SystemNotificationSeverity.md) |  | 
**NotificationTimeEpochMs** | **int64** |  | 
**Content** | **string** |  | 
**Toast** | **bool** |  | 

## Methods

### NewSystemNotification

`func NewSystemNotification(notificationId string, title string, severity SystemNotificationSeverity, notificationTimeEpochMs int64, content string, toast bool, ) *SystemNotification`

NewSystemNotification instantiates a new SystemNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNotificationWithDefaults

`func NewSystemNotificationWithDefaults() *SystemNotification`

NewSystemNotificationWithDefaults instantiates a new SystemNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationId

`func (o *SystemNotification) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *SystemNotification) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *SystemNotification) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.


### GetTitle

`func (o *SystemNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SystemNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SystemNotification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSeverity

`func (o *SystemNotification) GetSeverity() SystemNotificationSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SystemNotification) GetSeverityOk() (*SystemNotificationSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SystemNotification) SetSeverity(v SystemNotificationSeverity)`

SetSeverity sets Severity field to given value.


### GetNotificationTimeEpochMs

`func (o *SystemNotification) GetNotificationTimeEpochMs() int64`

GetNotificationTimeEpochMs returns the NotificationTimeEpochMs field if non-nil, zero value otherwise.

### GetNotificationTimeEpochMsOk

`func (o *SystemNotification) GetNotificationTimeEpochMsOk() (*int64, bool)`

GetNotificationTimeEpochMsOk returns a tuple with the NotificationTimeEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTimeEpochMs

`func (o *SystemNotification) SetNotificationTimeEpochMs(v int64)`

SetNotificationTimeEpochMs sets NotificationTimeEpochMs field to given value.


### GetContent

`func (o *SystemNotification) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SystemNotification) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SystemNotification) SetContent(v string)`

SetContent sets Content field to given value.


### GetToast

`func (o *SystemNotification) GetToast() bool`

GetToast returns the Toast field if non-nil, zero value otherwise.

### GetToastOk

`func (o *SystemNotification) GetToastOk() (*bool, bool)`

GetToastOk returns a tuple with the Toast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToast

`func (o *SystemNotification) SetToast(v bool)`

SetToast sets Toast field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


