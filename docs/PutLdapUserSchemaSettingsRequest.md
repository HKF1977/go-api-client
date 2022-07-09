# PutLdapUserSchemaSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserBase** | **string** | Base DN from which to search for Users | 
**UserIdAttribute** | **string** | Attribute type for the Users | 
**UserListFilter** | Pointer to **string** | Search filter for Users | [optional] 

## Methods

### NewPutLdapUserSchemaSettingsRequest

`func NewPutLdapUserSchemaSettingsRequest(userBase string, userIdAttribute string, ) *PutLdapUserSchemaSettingsRequest`

NewPutLdapUserSchemaSettingsRequest instantiates a new PutLdapUserSchemaSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLdapUserSchemaSettingsRequestWithDefaults

`func NewPutLdapUserSchemaSettingsRequestWithDefaults() *PutLdapUserSchemaSettingsRequest`

NewPutLdapUserSchemaSettingsRequestWithDefaults instantiates a new PutLdapUserSchemaSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserBase

`func (o *PutLdapUserSchemaSettingsRequest) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *PutLdapUserSchemaSettingsRequest) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *PutLdapUserSchemaSettingsRequest) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.


### GetUserIdAttribute

`func (o *PutLdapUserSchemaSettingsRequest) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *PutLdapUserSchemaSettingsRequest) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *PutLdapUserSchemaSettingsRequest) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.


### GetUserListFilter

`func (o *PutLdapUserSchemaSettingsRequest) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *PutLdapUserSchemaSettingsRequest) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *PutLdapUserSchemaSettingsRequest) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *PutLdapUserSchemaSettingsRequest) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


