# NotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NotificationConfigurationId** | Pointer to **int64** |  | [optional] 
**Status** | [**NotificationChannelStatus**](NotificationChannelStatus.md) |  | 
**Type** | **string** |  | 
**SlackWorkspace** | **string** |  | 
**SlackChannel** | Pointer to **string** |  | [optional] 
**SlackChannelId** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Token** | **string** |  | 
**VerifySsl** | **bool** |  | 
**Metadata** | **map[string]string** |  | 
**Region** | [**OpsgenieRegion**](OpsgenieRegion.md) |  | 
**GenieKey** | **string** |  | 
**Responders** | [**[]OpsgenieResponder**](OpsgenieResponder.md) |  | 
**Priority** | [**OpsgeniePriority**](OpsgeniePriority.md) |  | 
**To** | **[]string** |  | 
**Cc** | **[]string** |  | 
**SubjectPrefix** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationChannel

`func NewNotificationChannel(id int64, status NotificationChannelStatus, type_ string, slackWorkspace string, url string, token string, verifySsl bool, metadata map[string]string, region OpsgenieRegion, genieKey string, responders []OpsgenieResponder, priority OpsgeniePriority, to []string, cc []string, ) *NotificationChannel`

NewNotificationChannel instantiates a new NotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelWithDefaults

`func NewNotificationChannelWithDefaults() *NotificationChannel`

NewNotificationChannelWithDefaults instantiates a new NotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationChannel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationChannel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationChannel) SetId(v int64)`

SetId sets Id field to given value.


### GetNotificationConfigurationId

`func (o *NotificationChannel) GetNotificationConfigurationId() int64`

GetNotificationConfigurationId returns the NotificationConfigurationId field if non-nil, zero value otherwise.

### GetNotificationConfigurationIdOk

`func (o *NotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool)`

GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfigurationId

`func (o *NotificationChannel) SetNotificationConfigurationId(v int64)`

SetNotificationConfigurationId sets NotificationConfigurationId field to given value.

### HasNotificationConfigurationId

`func (o *NotificationChannel) HasNotificationConfigurationId() bool`

HasNotificationConfigurationId returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationChannel) GetStatus() NotificationChannelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationChannel) SetStatus(v NotificationChannelStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *NotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannel) SetType(v string)`

SetType sets Type field to given value.


### GetSlackWorkspace

`func (o *NotificationChannel) GetSlackWorkspace() string`

GetSlackWorkspace returns the SlackWorkspace field if non-nil, zero value otherwise.

### GetSlackWorkspaceOk

`func (o *NotificationChannel) GetSlackWorkspaceOk() (*string, bool)`

GetSlackWorkspaceOk returns a tuple with the SlackWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWorkspace

`func (o *NotificationChannel) SetSlackWorkspace(v string)`

SetSlackWorkspace sets SlackWorkspace field to given value.


### GetSlackChannel

`func (o *NotificationChannel) GetSlackChannel() string`

GetSlackChannel returns the SlackChannel field if non-nil, zero value otherwise.

### GetSlackChannelOk

`func (o *NotificationChannel) GetSlackChannelOk() (*string, bool)`

GetSlackChannelOk returns a tuple with the SlackChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannel

`func (o *NotificationChannel) SetSlackChannel(v string)`

SetSlackChannel sets SlackChannel field to given value.

### HasSlackChannel

`func (o *NotificationChannel) HasSlackChannel() bool`

HasSlackChannel returns a boolean if a field has been set.

### GetSlackChannelId

`func (o *NotificationChannel) GetSlackChannelId() string`

GetSlackChannelId returns the SlackChannelId field if non-nil, zero value otherwise.

### GetSlackChannelIdOk

`func (o *NotificationChannel) GetSlackChannelIdOk() (*string, bool)`

GetSlackChannelIdOk returns a tuple with the SlackChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannelId

`func (o *NotificationChannel) SetSlackChannelId(v string)`

SetSlackChannelId sets SlackChannelId field to given value.

### HasSlackChannelId

`func (o *NotificationChannel) HasSlackChannelId() bool`

HasSlackChannelId returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationChannel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *NotificationChannel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NotificationChannel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NotificationChannel) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerifySsl

`func (o *NotificationChannel) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *NotificationChannel) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *NotificationChannel) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.


### GetMetadata

`func (o *NotificationChannel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationChannel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationChannel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetRegion

`func (o *NotificationChannel) GetRegion() OpsgenieRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NotificationChannel) GetRegionOk() (*OpsgenieRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NotificationChannel) SetRegion(v OpsgenieRegion)`

SetRegion sets Region field to given value.


### GetGenieKey

`func (o *NotificationChannel) GetGenieKey() string`

GetGenieKey returns the GenieKey field if non-nil, zero value otherwise.

### GetGenieKeyOk

`func (o *NotificationChannel) GetGenieKeyOk() (*string, bool)`

GetGenieKeyOk returns a tuple with the GenieKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenieKey

`func (o *NotificationChannel) SetGenieKey(v string)`

SetGenieKey sets GenieKey field to given value.


### GetResponders

`func (o *NotificationChannel) GetResponders() []OpsgenieResponder`

GetResponders returns the Responders field if non-nil, zero value otherwise.

### GetRespondersOk

`func (o *NotificationChannel) GetRespondersOk() (*[]OpsgenieResponder, bool)`

GetRespondersOk returns a tuple with the Responders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponders

`func (o *NotificationChannel) SetResponders(v []OpsgenieResponder)`

SetResponders sets Responders field to given value.


### GetPriority

`func (o *NotificationChannel) GetPriority() OpsgeniePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationChannel) GetPriorityOk() (*OpsgeniePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationChannel) SetPriority(v OpsgeniePriority)`

SetPriority sets Priority field to given value.


### GetTo

`func (o *NotificationChannel) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NotificationChannel) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NotificationChannel) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *NotificationChannel) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *NotificationChannel) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *NotificationChannel) SetCc(v []string)`

SetCc sets Cc field to given value.


### GetSubjectPrefix

`func (o *NotificationChannel) GetSubjectPrefix() string`

GetSubjectPrefix returns the SubjectPrefix field if non-nil, zero value otherwise.

### GetSubjectPrefixOk

`func (o *NotificationChannel) GetSubjectPrefixOk() (*string, bool)`

GetSubjectPrefixOk returns a tuple with the SubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPrefix

`func (o *NotificationChannel) SetSubjectPrefix(v string)`

SetSubjectPrefix sets SubjectPrefix field to given value.

### HasSubjectPrefix

`func (o *NotificationChannel) HasSubjectPrefix() bool`

HasSubjectPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


