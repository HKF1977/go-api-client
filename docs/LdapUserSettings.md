# LdapUserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserBase** | Pointer to **string** | Base DN from which to search for Users | [optional] 
**UserIdAttribute** | Pointer to **int32** | Attribute type for the Users | [optional] 
**UserListFilter** | Pointer to **string** | Search filter for Users | [optional] 

## Methods

### NewLdapUserSettings

`func NewLdapUserSettings() *LdapUserSettings`

NewLdapUserSettings instantiates a new LdapUserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapUserSettingsWithDefaults

`func NewLdapUserSettingsWithDefaults() *LdapUserSettings`

NewLdapUserSettingsWithDefaults instantiates a new LdapUserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserBase

`func (o *LdapUserSettings) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *LdapUserSettings) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *LdapUserSettings) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.

### HasUserBase

`func (o *LdapUserSettings) HasUserBase() bool`

HasUserBase returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *LdapUserSettings) GetUserIdAttribute() int32`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *LdapUserSettings) GetUserIdAttributeOk() (*int32, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *LdapUserSettings) SetUserIdAttribute(v int32)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *LdapUserSettings) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserListFilter

`func (o *LdapUserSettings) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *LdapUserSettings) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *LdapUserSettings) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *LdapUserSettings) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


